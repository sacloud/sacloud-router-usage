// Copyright 2022 The sacloud/sacloud-router-usage Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/itchyny/gojq"
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/helper/api"
	"github.com/sacloud/iaas-api-go/search"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/sacloud-router-usage/version"
)

const (
	UNKNOWN = 3
	OK      = 0

	// CRITICAL = 2
	// WARNING  = 1
)

func main() {
	os.Exit(_main())
}

func _main() int {
	opts, err := parseOpts()
	if err != nil {
		log.Println(err)
		return UNKNOWN
	}
	if opts.Version {
		printVersion()
		return OK
	}

	client, err := routerClient()
	if err != nil {
		log.Println(err)
		return UNKNOWN
	}

	metrics, err := fetchMetrics(client, opts)
	if err != nil {
		log.Println(err)
		return UNKNOWN
	}

	if err := outputMetrics(os.Stdout, metrics, opts.Query); err != nil {
		log.Println(err)
		return UNKNOWN
	}
	return OK
}

type commandOpts struct {
	Time          uint     `long:"time" description:"Get average traffic for a specified amount of time" default:"3"`
	Item          string   `long:"item" description:"Item name" required:"true" choice:"in" choice:"out" default:"in"`
	Prefix        []string `long:"prefix" description:"prefix for router names. prefix accepts more than one." required:"true"`
	Zones         []string `long:"zone" description:"zone name" required:"true"`
	PercentileSet string   `long:"percentile-set" default:"99,95,90,75" description:"percentiles to dispaly"`
	Version       bool     `short:"v" long:"version" description:"Show version"`
	Query         string   `long:"query" description:"jq style query to result and display"`
	EnvFrom       string   `long:"env-from" description:"load environment values from this file"`
	percentiles   []percentile
}

type percentile struct {
	str   string
	float float64
}

func round(f float64) int64 {
	return int64(math.Round(f)) - 1
}

func routerClient() (iaasRouterAPI, error) {
	options := api.OptionsFromEnv()
	if options.AccessToken == "" {
		return nil, fmt.Errorf("environment variable %q is required", "SAKURACLOUD_ACCESS_TOKEN")
	}
	if options.AccessTokenSecret == "" {
		return nil, fmt.Errorf("environment variable %q is required", "SAKURACLOUD_ACCESS_TOKEN_SECRET")
	}

	if options.UserAgent == "" {
		options.UserAgent = fmt.Sprintf(
			"sacloud/sacloud-router-uage/v%s (%s/%s; +https://github.com/sacloud/sacloud-router-uage) %s",
			version.Version,
			runtime.GOOS,
			runtime.GOARCH,
			iaas.DefaultUserAgent,
		)
	}

	caller := api.NewCallerWithOptions(options)
	return iaas.NewInternetOp(caller), nil
}

type iaaSRouter struct {
	*iaas.Internet
	Zone string
}

type iaasRouterAPI interface {
	Find(ctx context.Context, zone string, conditions *iaas.FindCondition) (*iaas.InternetFindResult, error)
	MonitorRouter(ctx context.Context, zone string, id types.ID, condition *iaas.MonitorCondition) (*iaas.RouterActivity, error)
}

func findRouters(client iaasRouterAPI, opts *commandOpts) ([]*iaaSRouter, error) {
	var routers []*iaaSRouter
	for _, prefix := range opts.Prefix {
		for _, zone := range opts.Zones {
			condition := &iaas.FindCondition{
				Filter: map[search.FilterKey]interface{}{},
			}
			condition.Filter[search.Key("Name")] = search.PartialMatch(prefix)
			result, err := client.Find(
				context.Background(),
				zone,
				condition,
			)
			if err != nil {
				return nil, err
			}
			for _, r := range result.Internet {
				if strings.Index(r.Name, prefix) == 0 {
					routers = append(routers, &iaaSRouter{Internet: r, Zone: zone})
				}
			}
		}
	}
	return routers, nil
}

func fetchRouterMetrics(client iaasRouterAPI, opts *commandOpts, ss []*iaaSRouter) (map[string]interface{}, error) {
	b, _ := time.ParseDuration(fmt.Sprintf("-%dm", (opts.Time+3)*5))
	condition := &iaas.MonitorCondition{
		Start: time.Now().Add(b),
		End:   time.Now(),
	}

	valueFn := func(p *iaas.MonitorRouterValue) float64 {
		return p.GetOut()
	}
	if opts.Item == "in" {
		valueFn = func(p *iaas.MonitorRouterValue) float64 {
			return p.GetIn()
		}
	}

	var fs sort.Float64Slice
	routers := make([]interface{}, 0)
	total := float64(0)
	for _, t := range ss {
		activity, err := client.MonitorRouter(
			context.Background(),
			t.Zone,
			t.ID,
			condition,
		)
		if err != nil {
			return nil, err
		}
		usages := activity.GetValues()
		if len(usages) == 0 {
			continue
		}
		if len(usages) > int(opts.Time) {
			usages = usages[len(usages)-int(opts.Time):]
		}
		sum := float64(0)
		monitors := make([]interface{}, 0)
		for _, p := range usages {
			v := valueFn(p) / 1000 / 1000 // 単位変換: bps->Mbps
			m := map[string]interface{}{
				"traffic": v,
				"time":    p.GetTime().String(),
			}
			monitors = append(monitors, m)
			log.Printf("%s zone:%s traffic:%f time:%s", t.Name, t.Zone, v, p.GetTime().String())
			sum += v
		}
		avg := sum / float64(len(usages))
		log.Printf("%s average_traffic:%f", t.Name, avg)
		fs = append(fs, avg)
		total += avg

		routers = append(routers, map[string]interface{}{
			"name":     t.Name,
			"zone":     t.Zone,
			"avg":      avg,
			"monitors": monitors,
		})
	}

	if len(fs) == 0 {
		result := map[string]interface{}{}
		result["max"] = float64(0)
		result["avg"] = float64(0)
		result["min"] = float64(0)
		for _, p := range opts.percentiles {
			result[fmt.Sprintf("%spt", p.str)] = float64(0)
		}
		result["routers"] = routers
		return result, nil
	}

	sort.Sort(fs)
	fl := float64(len(fs))
	result := map[string]interface{}{}
	result["max"] = fs[len(fs)-1]
	result["avg"] = total / fl
	result["min"] = fs[0]
	for _, p := range opts.percentiles {
		result[fmt.Sprintf("%spt", p.str)] = fs[round(fl*(p.float))]
	}
	result["routers"] = routers
	return result, nil
}

func printVersion() {
	fmt.Printf(`%s %s
Compiler: %s %s
`,
		os.Args[0],
		version.Version,
		runtime.Compiler,
		runtime.Version())
}

func parseOpts() (*commandOpts, error) {
	opts := &commandOpts{}
	psr := flags.NewParser(opts, flags.HelpFlag|flags.PassDoubleDash)
	_, err := psr.Parse()
	if opts.Version {
		return opts, nil
	}

	if err != nil {
		return nil, err
	}

	if opts.Time < 1 {
		opts.Time = 1
	}

	if opts.EnvFrom != "" {
		if err := godotenv.Load(opts.EnvFrom); err != nil {
			return nil, err
		}
	}

	m := make(map[string]struct{})
	for _, z := range opts.Zones {
		if _, ok := m[z]; ok {
			return nil, fmt.Errorf("zone %q is duplicated", z)
		}
		m[z] = struct{}{}
	}

	var percentiles []percentile
	percentileStrings := strings.Split(opts.PercentileSet, ",")
	for _, s := range percentileStrings {
		if s == "" {
			continue
		}
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, fmt.Errorf("could not parse --percentile-set: %v", err)
		}
		f /= 100
		percentiles = append(percentiles, percentile{s, f})
	}
	opts.percentiles = percentiles

	return opts, nil
}

func fetchMetrics(client iaasRouterAPI, opts *commandOpts) (map[string]interface{}, error) {
	routers, err := findRouters(client, opts)
	if err != nil {
		return nil, err
	}

	return fetchRouterMetrics(client, opts, routers)
}

func outputMetrics(w io.Writer, metrics map[string]interface{}, query string) error {
	if query == "" {
		v, _ := json.Marshal(metrics)
		fmt.Fprintln(w, string(v))
		return nil
	}

	parsed, err := gojq.Parse(query)
	if err != nil {
		return err
	}
	iter := parsed.Run(metrics)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return err
		}
		if v == nil {
			return fmt.Errorf("%s not found in result", query)
		}
		j2, _ := json.Marshal(v)
		fmt.Fprintln(w, string(j2))
	}

	return nil
}
