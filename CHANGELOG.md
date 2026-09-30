# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

## [1.6.0](https://github.com/dirk39/gossm/compare/v1.5.0...v1.6.0) (2026-09-30)


### Features

* add golangci-lint and govulncheck quality gates (Refs [#3](https://github.com/dirk39/gossm/issues/3)) ([eb38254](https://github.com/dirk39/gossm/commit/eb38254502a5ca165701bec75638d9dcb14fa274))
* automate releases, changelog, and semantic tagging (Refs [#9](https://github.com/dirk39/gossm/issues/9)) ([c050961](https://github.com/dirk39/gossm/commit/c050961e22608e98f5200f9a94f0ce82d7e750e2))
* migrate CI/CD from CircleCI to GitHub Actions (Refs [#1](https://github.com/dirk39/gossm/issues/1)) ([f4d87cd](https://github.com/dirk39/gossm/commit/f4d87cd7125ebe7e5a20062f8a7cdb07428ec683))
* upgrade Go runtime to 1.27.0 and update dependencies (Refs [#5](https://github.com/dirk39/gossm/issues/5)) ([ab0d7ec](https://github.com/dirk39/gossm/commit/ab0d7ecd95a4235a812c8ac9c72276e43934c545))


### Bug Fixes

* update session-manager-plugin assets with correct arm64 binary (Refs [#7](https://github.com/dirk39/gossm/issues/7)) ([6b34b75](https://github.com/dirk39/gossm/commit/6b34b757772fdfe54bdee7ae6346f03250d4a227))

## [1.5.0] - 2022-07-13

### Features
* Release v1.5.0.

## [1.1.0] - 2020-04-06

### Features
* Update `session-manager-plugin` from 1.1.54.0 to 1.1.61.0.
* Add `--version` flag support.

## [1.0.3] - 2020-02-14

### Features
* Support `AWS_PROFILE` environment variable.

## [1.0.2] - 2020-01-20

### Bug Fixes
* If target args exist, set to DNS variable in viper.

## [1.0.0] - 2020-01-15

### Initial Release
* Initial release of `gossm` interactive CLI for AWS Systems Manager Session Manager.
