# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and
this project follows [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Maintainer-dispatched GitHub Release workflow with changelog finalization and
  CI-owned tag creation.
- Release validation scripts, agent guidance (`AGENTS.md`), and
  [`docs/release-checklist.md`](docs/release-checklist.md).

## [0.11.0] - 2026-02-07

### Added

- JSON Lines (NDJSON) assertions in `assertjson`.
- String assertions for prefix, suffix, integer, and number checks.
- Array-of-strings assertions in `assertjson`.

### Changed

- `Has` and `FileHas` in `assertjson` return `bool`.
- CI and golangci-lint updates; README improvements.

[Unreleased]: https://github.com/muonsoft/api-testing/compare/v0.11.0...HEAD
[0.11.0]: https://github.com/muonsoft/api-testing/releases/tag/v0.11.0
