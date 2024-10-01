# Change Log
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).
## Unreleased

## [8.4.0] - 2024-10-01
### Added
- Implemented `server/listiso` and `server/mountiso` endpoints.
### Changed
- Add `ISOFile` attribute in `ServerDetails`.
- Add `Type` attribute for disk policy in `ServerDiskDetails`.

## [8.3.1] - 2024-09-19
### Changed
- Fix `privatenetwork/create` and `/details` request struct.

## [8.3.0] - 2024-09-11
### Added
- Servers - Implement `server/networkadapters` endpoint.
### Changed
- NetworkAdapters - New attributes `IsConnected`, `IsPrimary` & `MacAddress`

## [8.2.0] - 2024-09-04
### Added
- PrivateNetworks - Implement new endpoint `/privatenetwork/*`.
- NetworkCircuits - Implement new endpoint `/networkcircuit/*`.
### Changed
- Removed OpenVZ references in code.

## [8.1.0] - 2024-01-15
### Added
- ServerDisks - Implement `/serverdisk` endpoint.

## [8.0.0] - 2023-11-20
### Changed
- **BREAKING** - Server Templates cost float64
- Fixed rand.Seed deprecation
- Updated mergo dependency
## [7.1.0] - 2023-05-25
### Added
- Implement user/login
- Implement DNSDomains Export
- Implement DNSDomains GenerateAuthCode

### Changed
- Tweaking http functions to work with user/login
- Update dependencies
- Update Go version for GH Actions

## [7.0.1] - 2023-03-30
- Set correct major version in go.mod

## [7.0.0] - 2023-03-30
### Added
- Implement Server Console endpoint
- Implement Email ResetPassword, obtain a new password for a specified email account.
- Email `EmailAccount` type now contains `Password` for new accounts.
- ServerDetails now has `ServerBackupDetails`.

### Changed
- **BREAKING:** - Email EditAccount has no parameter `Password`
- **BREAKING:** - IPs `Reserved()` now requires `ReservedIPsParams{}` to allow
  filtering on `datacenter`, `platform`, `used` and `version`.
- Remove references to OpenVZ in Servers.
- Servers `WithDefaults()` now uses KVM platform as default.
- Bumped various dependencies..

## [6.1.0] - 2022-10-31
### Changed
- BaseURL is now exposed. With a helper method `SetBaseURL`

## [6.0.0] - 2022-10-28
### Changed
- **BREAKING:** - Go1.18 Required.
- Fix `CloudConfigParams` in Servers. (#67)
- Remove redundant WithContext call. (#69)
- Remove unused data struct from LoadBalancer `AddCertificate` function. (#70)
- Remove deprecated io/ioutil calls. (#72)

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
