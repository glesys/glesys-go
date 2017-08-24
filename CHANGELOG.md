# Change Log
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased]

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
