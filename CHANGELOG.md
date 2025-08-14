# Changelog

## [v0.2.1](https://github.com/sacloud/sacloud-router-usage/compare/v0.2.0...v0.2.1) - 2025-08-14
- ci: bump actions/checkout from 4 to 5 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/82

## [v0.2.0](https://github.com/sacloud/sacloud-router-usage/compare/v0.1.0...v0.2.0) - 2025-08-13
- golangci-lint v2 by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/73
- update actions - use go-version-file by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/74
- ubuntu-20.04 -> 22.04, 24.04, latest by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/75
- ci: bump docker/build-push-action from 5 to 6 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/70
- go: bump github.com/sacloud/go-otelsetup from 0.0.6 to 0.5.0 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/71
- ci: bump goreleaser/goreleaser-action from 5 to 6 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/69
- update dependencies: sacloud/iaas-api-go and sacloud/sacloud-usage-lib by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/76
- opentelemetry-collector:0.132.0 by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/77
- tagpr and goreleaser by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/79
- textlint: ignore CHANGELOG.md by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/81

## [v0.1.0](https://github.com/sacloud/sacloud-router-usage/compare/v0.0.2...v0.1.0) - 2024-04-09
- go: bump golang.org/x/crypto from 0.0.0-20220214200702-86341886e292 to 0.1.0 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/21
- ci: bump actions/setup-go from 3 to 4 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/22
- ci: bump docker/build-push-action from 3 to 4 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/18
- ci: bump goreleaser/goreleaser-action from 3 to 4 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/11
- refactor: mainから各funcを切り出し by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/30
- GitHub ActionsでのCIパフォーマンスの改善 by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/31
- refactor: added tests on metrics fetch/usage calculation by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/32
- refactor: added tests for output proceccing by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/37
- fix: --percentile-setが空の場合の対応を追加 by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/39
- go: bump github.com/stretchr/testify from 1.8.0 to 1.8.4 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/40
- ci: bump crazy-max/ghaction-import-gpg from 5 to 6 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/38
- ci: bump docker/metadata-action from 4 to 5 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/36
- ci: bump docker/build-push-action from 4 to 5 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/35
- ci: bump actions/checkout from 3 to 4 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/34
- ci: bump goreleaser/goreleaser-action from 4 to 5 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/33
- go: bump github.com/sacloud/iaas-api-go from 1.6.0 to 1.11.1 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/29
- go: bump github.com/itchyny/gojq from 0.12.9 to 0.12.13 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/28
- go: bump github.com/joho/godotenv from 1.4.0 to 1.5.1 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/19
- refactor: outputMetrics()をcommandOpts非依存に by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/41
- refactor: リソース/アクティビティモニタ取得部分とメトリクス算出部分を分離 by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/42
- refactor: usageパッケージを切り出し by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/43
- refactor: sacloud-usage-libへ切り替え by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/44
- sacloud-usage-lib@v0.0.3 by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/45
- sacloud-usage-lib@v0.0.4 by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/49
- ci: bump docker/setup-qemu-action from 2 to 3 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/48
- ci: bump docker/setup-buildx-action from 2 to 3 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/47
- ci: bump docker/login-action from 2 to 3 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/46
- ci: bump actions/setup-go from 4 to 5 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/50
- go: bump github.com/sacloud/iaas-api-go from 1.11.1 to 1.11.2 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/51
- go: bump github.com/sacloud/sacloud-usage-lib from 0.0.4 to 0.0.5 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/52
- Trace with OpenTelemetry by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/53
- sacloud/sacloud-usage-lib@v0.0.7 by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/55
- OTel計装周りの修正 - 命名ルールの統一 by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/59
- go: bump github.com/sacloud/iaas-api-go from 1.11.2 to 1.12.0 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/64
- go: bump github.com/stretchr/testify from 1.8.4 to 1.9.0 by @dependabot[bot] in https://github.com/sacloud/sacloud-router-usage/pull/63

## [v0.0.2](https://github.com/sacloud/sacloud-router-usage/compare/v0.0.1...v0.0.2) - 2022-11-08
- 単位をbpsからMbpsに修正 by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/6
- docs: example更新 by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/7

## [v0.0.1](https://github.com/sacloud/sacloud-router-usage/commits/v0.0.1) - 2022-10-25
- sacloud/sacloud-cpu-usageからの移植 by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/1
- misc update by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/2
- Usageの更新 by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/3
- misc update by @yamamoto-febc in https://github.com/sacloud/sacloud-router-usage/pull/4
