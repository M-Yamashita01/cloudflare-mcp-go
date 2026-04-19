# Release Flow

## Overview

This project uses a manual tag-based release flow. Releases are created only when features are sufficiently ready, not on every PR merge.

The release pipeline is: **tag push (`v*`) → CI (lint + test) → GoReleaser builds cross-platform binaries → GitHub Release created automatically**.

## When to Release

- A meaningful set of features or fixes has been merged to `main`
- All CI checks on `main` are passing
- The codebase is in a stable, tested state

There is no fixed schedule. Release when the changes are worth shipping.

## Versioning

Follow [Semantic Versioning](https://semver.org/):

- **MAJOR** (`v1.0.0` → `v2.0.0`): Breaking changes (e.g., tool renamed, input schema changed)
- **MINOR** (`v0.1.0` → `v0.2.0`): New features (e.g., new Cloudflare API tool added)
- **PATCH** (`v0.1.0` → `v0.1.1`): Bug fixes, documentation updates

While in `v0.x.x`, breaking changes may occur in MINOR versions.

## Release Steps

### 1. Verify main is stable

```bash
git checkout main
git pull
```

Confirm the latest CI run on `main` is green (lint + test passing).

### 2. Choose the version number

Check the latest tag:

```bash
git tag --sort=-v:refname | head -5
```

Decide the next version based on the changes since the last release.

### 3. Create and push the tag

```bash
git tag v0.2.0
git push origin v0.2.0
```

### 4. CI builds and releases automatically

The `release.yml` GitHub Actions workflow will:

1. Run lint (golangci-lint)
2. Run tests (`go test -race ./...`)
3. Build binaries via GoReleaser for:
   - `linux/amd64`, `linux/arm64`
   - `darwin/amd64`, `darwin/arm64`
   - `windows/amd64`
4. Create a GitHub Release with:
   - Downloadable archives (`.tar.gz` for Linux/macOS, `.zip` for Windows)
   - Auto-generated changelog from commits since the previous tag

### 5. Verify the release

- Check the [Releases page](https://github.com/M-Yamashita01/cloudflare-mcp-go/releases)
- Confirm all platform binaries are attached
- Review the auto-generated changelog and edit if needed

## Release Infrastructure

The release pipeline consists of:

| File | Purpose |
|------|---------|
| `.goreleaser.yml` | GoReleaser config (build targets, archive format, changelog) |
| `.github/workflows/release.yml` | GitHub Actions workflow triggered by `v*` tag push |
| `main.go` (`version` variable) | Version injected via `-ldflags` at build time |

### GoReleaser

- Builds static binaries (`CGO_ENABLED=0`)
- Injects git tag version into `main.version` via ldflags
- Generates changelog grouped by commit type (`feat:`, `fix:`, etc.)

## User Installation

After a release, users can download and use the binary without building from source:

```bash
# Download from GitHub Releases (example: macOS ARM64)
curl -LO https://github.com/M-Yamashita01/cloudflare-mcp-go/releases/latest/download/cloudflare-mcp-go_Darwin_arm64.tar.gz
tar xzf cloudflare-mcp-go_Darwin_arm64.tar.gz
chmod +x cloudflare-mcp-go

# Register with Claude Code
claude mcp add cloudflare -- ./cloudflare-mcp-go
```

## Fixing a Bad Release

If a release has a critical issue:

1. **Delete the release** from the GitHub Releases page
2. **Delete the tag**: `git tag -d v0.2.0 && git push origin :refs/tags/v0.2.0`
3. Fix the issue on `main`
4. Re-tag and push: `git tag v0.2.0 && git push origin v0.2.0`

Or simply release a patch version (`v0.2.1`) with the fix.
