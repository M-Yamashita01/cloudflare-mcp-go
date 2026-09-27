# Release Flow

## Overview

This project uses an automated release flow driven by [release-please](https://github.com/googleapis/release-please)
and Conventional Commits. There is no manual tagging.

The pipeline is:

```
Conventional Commits merged to main
   |
   v
release-please keeps a Release PR up to date (version bump + CHANGELOG)
   |
   v  (you merge the Release PR when you decide to ship)
release-please creates the tag and the GitHub Release
   |
   v
GoReleaser builds cross-platform binaries and appends them to that release
```

For the decision policy (when to ship, how to pick the version), see
[versioning-policy.md](versioning-policy.md).

## Conventional Commits

release-please derives the next version from commit messages, so the prefix matters:

| Prefix | Effect on version (while in 0.x) |
|--------|----------------------------------|
| `feat:` | MINOR bump (new feature) |
| `fix:` | PATCH bump (bug fix) |
| `fix(deps):` | PATCH bump (dependency update, used by Renovate) |
| `feat!:` / `BREAKING CHANGE:` | MINOR bump (breaking changes stay MINOR until 1.0.0) |
| `docs:`, `refactor:`, `perf:` | Appear in the changelog; no version bump on their own |
| `chore:` | No version bump, hidden from the changelog |

If nothing since the last release bumps the version, no Release PR is created,
so weeks with only `chore:` changes produce no release.

## Releasing

### 1. Merge feature/fix PRs to main as usual

Each PR should use a Conventional Commit title. When such commits land on `main`,
the `Release Please` workflow opens or updates a single Release PR titled
`chore(main): release x.y.z`.

### 2. Review the Release PR (typically weekly)

Open the Release PR and check:

- The proposed version number matches the changes (see versioning-policy.md)
- The generated `CHANGELOG.md` entry reads correctly

If there is nothing worth shipping this week (for example, no user-facing change),
leave the PR open and revisit next week.

### 3. Merge the Release PR to release

Merging the Release PR is the single "ship it" action. It:

1. Updates `CHANGELOG.md` and `.release-please-manifest.json` on `main`
2. Creates the git tag (e.g. `v0.2.1`)
3. Creates the GitHub Release with the changelog as its body

### 4. Binaries are built automatically

The `goreleaser` job in the same workflow runs only when a release was created.
It builds binaries for:

- `linux/amd64`, `linux/arm64`
- `darwin/amd64`, `darwin/arm64`
- `windows/amd64`

and appends the archives (`.tar.gz` for Linux/macOS, `.zip` for Windows) to the
release that release-please created.

### 5. Verify the release

- Check the [Releases page](https://github.com/M-Yamashita01/cloudflare-mcp-go/releases)
- Confirm all platform binaries are attached
- Confirm the changelog reads correctly

## Release Infrastructure

| File | Purpose |
|------|---------|
| `.github/workflows/release-please.yml` | Runs release-please on push to `main`; runs GoReleaser when a release is created |
| `release-please-config.json` | release-please settings (release type, version bump rules, changelog sections) |
| `.release-please-manifest.json` | Tracks the current released version |
| `.goreleaser.yml` | GoReleaser config; `release.mode: append` so it adds binaries to the existing release |
| `main.go` (`version` variable) | Version injected via `-ldflags` at build time by GoReleaser |

## Renovate

Renovate is configured to use `fix(deps):` commit messages (see `renovate.json`),
so dependency updates become PATCH-level entries in the Release PR. This lets a
weekly dependency refresh ship as a single patch release when you merge the Release PR.

## Fixing a Bad Release

Prefer rolling forward: fix the issue on `main` and let the next Release PR ship a
new patch version (e.g. `v0.2.2`).

If a release must be pulled:

1. Delete the release from the GitHub Releases page
2. Delete the tag: `git push origin :refs/tags/v0.2.1`
3. Fix the issue on `main` and ship a new patch via the normal flow
