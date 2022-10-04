# Change Log
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).
## Unreleased

## [5.0.0] - 2022-10-06
### Changed
- **BREAKING:** - EmailDomains GlobalQuota deprecated.
- **BREAKING:** - EmailDomains EmailQuota struct 'Used' and 'Total' fields deprecated.
  Use 'UsedInMiB' and 'QuotaInGiB'
- **BREAKING:** - LoadBalancers 'AddtoBlacklist' 'RemoveFromBlacklist' deprecated.
  Use 'AddToBlocklist' and 'RemoveFromBlocklist' instead.
- **BREAKING:** - DNSDomains OrganizationNumber deprecated. Use 'NationalID'
  instead.
- Add `CloudConfig` and `CloudConfigParams` field to `CreateServerParams`
### Added
- Implement Server Templates endpoint.
- Implement Server PreviewCloudConfig endpoint.

## [4.0.1] - 2022-09-20
### Change
- Fix module version in go.mod
- Bump version number after release

## [4.0.0] - 2022-09-20
### Changed
- **BREAKING:** - server IsRunning() and IsLocked() functions deprecated.
- New fields in ServerDetails: IsRunning & IsLocked to match new fields returned
  by API.

## [3.0.0] - 2021-11-11
### Changed
- **BREAKING:** - Cost and Amount changed from int to float64. - @norrland
- Code now in base directory of project.

## [2.5.0] - 2020-12-15
### Changed
- `ObjectStorageService` resource (#32). - @norrland
- ServerDetails InitialTemplate to describe the template used during server creation (#33). - @norrland

## [2.4.2] - 2020-05-07
### Changed
- Use string type for CreateRecords. - @norrland
- Properly comment ServerIP struct - @norrland

## [2.4.1] - 2020-04-30
### Changed
- v2 dir for proper Go module versioning support. - @larsdunemark
- New `ServerIP` struct instead of using IP objects. - @larsdunemark

## [2.4.0] - 2020-02-13
### Changed
- `DNSDomainService` resource (#18). - @norrland
- Support `CreateServerParams.users` for KVM platform (#21). - @glesys-andreas
- `EmailDomainService` resource (#19). - @cromigon
- Additional `IP` fields and support for `SetPTR` and `ResetPTR` (#20). – @norrland
- `glesys/glesys-go` is now a Go module for Go >= 1.11 and works outside of $GOPATH (#16). - @norrland

## [2.3.1] - 2019-06-07
### Changed
- Bump version numbers after release.

## [2.3.0] - 2019-06-07
### Added
- `LoadBalancerService` resource. - @norrland

### Changed
- Reference the current URL for GleSYS Cloud. - @abergman

## [2.2.0] - 2018-11-04
### Added
- `Network.IsPublic()` helper. Thanks to @norrland.

## [2.1.0] - 2017-08-23
### Added
- `NetworkService` and `NetworkAdapterService` are now available with support
  for creating, editing and destroying networks and network adapters. Big thanks
  to @norrland for championing this.

- `ServerService.Edit()` allows for editing servers. Thanks to @norrland.

- `IP.IsIPv4()` and `IP.IsIPv6()` helpers. Thanks to @norrland.

### Changed
- The `ServerDetails` struct now contains `Bandwidth`, `Description` and
  `Template`. Thanks to @norrland.

## [2.0.0] - 2017-02-13
### Changed
- **BREAKING:** `NewClient()` now requires a user agent string.

## [1.0.0] - 2017-01-26
### Added
- Initial release
