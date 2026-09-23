# Project agent rules

This repository is `github.com/muonsoft/api-testing`.

## Sources of truth

- `README.md` — public API and usage notes.
- `CHANGELOG.md` — Keep a Changelog record of user-visible changes.
- `docs/release-checklist.md` — release procedure and verification.

## Release policy

- Update `CHANGELOG.md` `[Unreleased]` when behavior or public docs change.
- Local agents and scripts never create or push release tags. The
  maintainer-dispatched Release workflow is the only path authorized to push its
  changelog-only commit and create a release tag.

## Work discipline

- Preserve unrelated user changes.
- Keep changes focused and commits atomic.
- Run checks proportional to the change (`gofmt`, `go test`, `go test -race`) and
  `bash scripts/test-all.sh` before publication.
