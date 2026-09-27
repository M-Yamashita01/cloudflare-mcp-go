# Versioning and Release Decision Policy

This document describes how to bump versions and when to cut a release (tag).
It is the shared reference for both humans and AI agents when making these decisions.
For the operational steps, see release.md.

## Context

This repository is an application (an MCP server), not a library.
Its users are people who run the server; there is little concern about other projects
importing it and breaking. Because of that, we prioritize being able to track what
changed over strict compatibility guarantees.

## Versioning Policy

Follow Semantic Versioning (SemVer). https://semver.org/

- MAJOR (v1.0.0 -> v2.0.0): Breaking changes, such as renaming a tool or changing an input schema
- MINOR (v0.2.0 -> v0.3.0): Backward-compatible new features, such as adding a new Cloudflare API tool
- PATCH (v0.2.0 -> v0.2.1): Bug fixes, dependency updates, documentation fixes

While in 0.x, breaking changes are allowed in MINOR versions.
Release 1.0.0 once the API is stable and backward compatibility can be promised.

### Linear versioning (no branching of version lines)

Versions only move forward. v0.2.1 is a continuation of the v0.2.x line;
it is not a separate thing from v0.2.

```
v0.2.0 -> v0.2.1 -> v0.2.2 -> ... -> v0.3.0 -> v0.3.1 -> ...
```

Shipping a patch as v0.2.1 is itself the result of patching v0.2,
so there is no need to apply the same change to a separate branch (backporting).

Backport maintenance, where an older line keeps receiving fixes in parallel
(for example, still patching v0.2.x after v0.3 is out), is something large OSS projects do
when many users are pinned to old versions. Given this project's size and 0.x stage,
we do not adopt it. Users are expected to move to the latest version.

## When to Cut a Tag

Decide using a combination of time-based and event-based triggers.

| Trigger | Action |
|---|---|
| Regular (weekly) | Ship accumulated fixes together as a PATCH |
| Security fix | Release immediately, without waiting for the regular cadence |
| New feature (feat) | Release as MINOR once the feature is usable |

### Release / no-release criteria

- Release: when there is at least one user-facing change worth shipping (bug fix, security, new feature)
- Do not release: when the changes are dependency updates only and do not affect user-facing behavior; it is fine to accumulate them and defer to the next week
- When in doubt: put it on the weekly cadence. Skip weeks with no changes

Dependency updates (such as Renovate) do not require a version bump every time.
Commits/merges and releases (version assignment) are decoupled;
bump the version when you cut a release, batching the changes.

## Release Flow

A tag push is the trigger; CI and the build run automatically.

```
main is stable (CI green)
   |
   v
Decide the next version (choose MAJOR/MINOR/PATCH per the policy above)
   |
   v
Create and push a tag (e.g. git tag v0.2.1 && git push origin v0.2.1)
   |
   v
release.yml fires -> lint + test -> GoReleaser builds per-OS binaries
   |
   v
A GitHub Release is created automatically
```

The concrete commands, the GoReleaser configuration, and how to fix a bad release
are documented in release.md.
