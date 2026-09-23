#!/usr/bin/env bash
# Prepare CHANGELOG.md for a SemVer release (Keep a Changelog).
#
# Supports an exact planned section (`## [X.Y.Z] — planned`), promotion of a
# non-empty [Unreleased] section, validation-only modes, and idempotent reruns.
#
# Usage:
#   scripts/prepare-release.sh 0.1.0 [YYYY-MM-DD]
#   scripts/prepare-release.sh 0.1.0 --check-only
#   scripts/prepare-release.sh 0.1.0 --require-final
#   scripts/prepare-release.sh 0.1.0 --semver-only

set -euo pipefail

version=${1:?version required (for example 0.1.0)}
changelog=${CHANGELOG_FILE:-CHANGELOG.md}
release_date=
check_only=false
require_final=false
semver_only=false

shift
while [[ $# -gt 0 ]]; do
  case "$1" in
    --check-only)
      check_only=true
      ;;
    --require-final)
      require_final=true
      ;;
    --semver-only)
      semver_only=true
      ;;
    --*)
      echo "unknown option: $1" >&2
      exit 1
      ;;
    *)
      if [[ -n "$release_date" ]]; then
        echo "unexpected extra argument: $1" >&2
        exit 1
      fi
      release_date=$1
      ;;
  esac
  shift
done

validate_numeric_identifier() {
  local identifier=$1
  local label=$2
  if [[ ! "$identifier" =~ ^(0|[1-9][0-9]*)$ ]]; then
    echo "invalid ${label} identifier: ${identifier}" >&2
    return 1
  fi
}

validate_semver() {
  local candidate=$1
  local core prerelease major minor patch identifier

  if [[ "$candidate" == *+* ]]; then
    echo "build metadata is not supported in release versions: ${candidate}" >&2
    return 1
  fi

  core=${candidate%%-*}
  prerelease=
  if [[ "$candidate" == *-* ]]; then
    prerelease=${candidate#*-}
  fi

  IFS=. read -r major minor patch <<<"$core"
  if [[ -z "${major:-}" || -z "${minor:-}" || -z "${patch:-}" || "$core" == *.*.*.* ]]; then
    echo "invalid SemVer core: ${candidate}" >&2
    return 1
  fi
  validate_numeric_identifier "$major" major || return 1
  validate_numeric_identifier "$minor" minor || return 1
  validate_numeric_identifier "$patch" patch || return 1

  if [[ "$candidate" == *-* ]]; then
    if [[ -z "$prerelease" || "$prerelease" == .* || "$prerelease" == *. || "$prerelease" == *..* ]]; then
      echo "invalid prerelease identifiers: ${candidate}" >&2
      return 1
    fi
    IFS=. read -ra prerelease_parts <<<"$prerelease"
    for identifier in "${prerelease_parts[@]}"; do
      if [[ "$identifier" =~ ^[0-9]+$ ]]; then
        validate_numeric_identifier "$identifier" prerelease || return 1
      elif [[ ! "$identifier" =~ ^[0-9A-Za-z-]+$ ]]; then
        echo "invalid prerelease identifier: ${identifier}" >&2
        return 1
      fi
    done
  fi
}

validate_semver "$version"
if [[ "$semver_only" == true ]]; then
  exit 0
fi

if [[ ! -f "$changelog" ]]; then
  echo "changelog file not found: $changelog" >&2
  exit 1
fi

if [[ -z "$release_date" ]]; then
  release_date=$(date -u +%Y-%m-%d)
fi
if [[ ! "$release_date" =~ ^[0-9]{4}-[0-9]{2}-[0-9]{2}$ ]]; then
  echo "invalid release date: $release_date (expected YYYY-MM-DD)" >&2
  exit 1
fi

final_prefix="## [${version}] - "
planned_header="## [${version}] — planned"

script_root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)

resolve_changelog_repo() {
  if [[ -n "${CHANGELOG_REPO:-}" ]]; then
    echo "$CHANGELOG_REPO"
    return
  fi
  if [[ -f "${script_root}/go.mod" ]]; then
    awk '/^module / { print $2; exit }' "${script_root}/go.mod"
    return
  fi
  echo "github.com/muonsoft/api-testing"
}

update_changelog_footer() {
  local release_version=$1
  local changelog_file=$2
  local repo tag unreleased_line version_line
  repo=$(resolve_changelog_repo)
  tag="v${release_version}"
  unreleased_line="[Unreleased]: https://${repo}/compare/${tag}...HEAD"
  version_line="[${release_version}]: https://${repo}/releases/tag/${tag}"

  awk -v unreleased_line="$unreleased_line" -v version_line="$version_line" -v version="$release_version" '
    function is_version_link(line,    v) {
      return index(line, "[" v "]: ") == 1
    }
    BEGIN {
      have_unreleased = 0
      have_version_link = 0
      insert_version_after_unreleased = 0
    }
    /^\[Unreleased\]: / {
      print unreleased_line
      have_unreleased = 1
      if (!have_version_link) {
        insert_version_after_unreleased = 1
      }
      next
    }
    /^\[/ {
      if (is_version_link($0, version)) {
        print version_line
        have_version_link = 1
        insert_version_after_unreleased = 0
        next
      }
      if (insert_version_after_unreleased) {
        print version_line
        have_version_link = 1
        insert_version_after_unreleased = 0
      }
      print
      next
    }
    {
      if (insert_version_after_unreleased && $0 !~ /^$/) {
        print version_line
        have_version_link = 1
        insert_version_after_unreleased = 0
      }
      print
    }
    END {
      if (!have_unreleased) {
        print unreleased_line
      }
      if (!have_version_link) {
        print version_line
      }
    }
  ' "$changelog_file" >"${changelog_file}.footer.tmp"
  mv "${changelog_file}.footer.tmp" "$changelog_file"
}

is_finalized() {
  awk -v prefix="$final_prefix" '
    index($0, prefix) == 1 {
      value = substr($0, length(prefix) + 1)
      if (value ~ /^[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]$/) {
        found = 1
        exit
      }
    }
    END { exit(found ? 0 : 1) }
  ' "$changelog"
}

is_planned() {
  grep -qFx "$planned_header" "$changelog"
}

section_has_content() {
  local header=$1
  awk -v header="$header" '
    $0 == header { in_section = 1; next }
    in_section && /^## \[/ { exit }
    in_section && /^### / { found = 1; exit }
    END { exit(found ? 0 : 1) }
  ' "$changelog"
}

if is_finalized; then
  update_changelog_footer "$version" "$changelog"
  echo "changelog section [${version}] is finalized"
  exit 0
fi

if [[ "$require_final" == true ]]; then
  echo "changelog section [${version}] is not finalized" >&2
  exit 1
fi

source_section=
if is_planned && section_has_content "$planned_header"; then
  source_section=planned
elif grep -qFx '## [Unreleased]' "$changelog" && section_has_content '## [Unreleased]'; then
  source_section=unreleased
fi

if [[ -z "$source_section" ]]; then
  echo "no non-empty planned [${version}] or [Unreleased] section in $changelog" >&2
  exit 1
fi

if [[ "$check_only" == true ]]; then
  echo "${source_section} changelog section is ready for ${version}"
  exit 0
fi

temporary=$(mktemp "${TMPDIR:-/tmp}/api-testing-changelog.XXXXXX")
cleanup() {
  rm -f "$temporary"
}
trap cleanup EXIT INT HUP TERM

if [[ "$source_section" == planned ]]; then
  awk -v planned="$planned_header" -v final="${final_prefix}${release_date}" '
    $0 == planned { print final; next }
    { print }
  ' "$changelog" >"$temporary"
else
  awk -v final="${final_prefix}${release_date}" '
    $0 == "## [Unreleased]" {
      print "## [Unreleased]"
      print ""
      print final
      next
    }
    { print }
  ' "$changelog" >"$temporary"
fi

mv "$temporary" "$changelog"
trap - EXIT INT HUP TERM
update_changelog_footer "$version" "$changelog"
echo "finalized changelog section [${version}] - ${release_date}"
