# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and
this project follows [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.12.0] - 2026-09-23

### Added

- Maintainer-dispatched GitHub Release workflow with changelog finalization and
  CI-owned tag creation.
- Release validation scripts, agent guidance (`AGENTS.md`), and
  [`docs/release-checklist.md`](docs/release-checklist.md).
- Internal HS256 JWT parsing for assertions (`internal/jwt`); public JWT types
  (`assertjson.JWTToken`, `JWTKeyFunc`, `JWTMapClaims`) and `SignHS256JWT` test helper.

### Changed

- JWT assertion callbacks use `assertjson` types instead of `github.com/golang-jwt/jwt/v5`.
- Module dependencies (`github.com/gofrs/uuid/v5`, `github.com/stretchr/testify`,
  `golang.org/x/net`) and minimum Go version (`1.26`).
- CI and local lint use golangci-lint `v2.13.2` (`gomodguard_v2`, test exclusions for
  `goconst`).

### Removed

- Direct dependency on `github.com/golang-jwt/jwt/v5`.
- Deprecated `assertjson` APIs: legacy `AssertNode` helpers (`EqualToTheString`,
  `IsNumberGreaterThan`, and similar), `Nodef` / `Atf`, UUID chain aliases
  (`Nil`, `Version`, …), and JSON Pointer strings passed as a single path
  argument to `Node` / `At`.

## [0.11.0] - 2026-02-07

### Added

- JSON Lines (NDJSON) assertions in `assertjson`.
- String assertions for prefix, suffix, integer, and number checks.
- Array-of-strings assertions in `assertjson`.

### Changed

- `Has` and `FileHas` in `assertjson` return `bool`.
- CI and golangci-lint updates; README improvements.

[Unreleased]: https://github.com/muonsoft/api-testing/compare/v0.12.0...HEAD
[0.12.0]: https://github.com/muonsoft/api-testing/releases/tag/v0.12.0
[0.11.0]: https://github.com/muonsoft/api-testing/releases/tag/v0.11.0
