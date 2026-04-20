# CLAUDE.md

## Project Overview

Cloudflare MCP Go — A Go implementation of an MCP (Model Context Protocol) server that wraps Cloudflare API v4.

- Language: Go 1.23+
- MCP SDK: `github.com/modelcontextprotocol/go-sdk`
- Entry point: `main.go`
- Architecture: See @doc/architecture.md for detailed design docs

## Build & Run

```bash
# Build
go build -o cloudflare-mcp-go .

# Run (requires CLOUDFLARE_API_TOKEN and CLOUDFLARE_ACCOUNT_ID env vars)
go run .

# Test
go test ./...

# Format & Vet
gofmt -w .
go vet ./...
```

## Adding a New Tool

1. Create or extend a package under `internal/tool/<domain>/`
2. Define Input struct + handler function + `RegisterTools()`
3. Add `<domain>.RegisterTools(server)` in `main.go`
4. Write co-located tests in `*_test.go`

## Code Style

Follow the [Google Go Style Guide](https://google.github.io/styleguide/go/index) as the baseline.

Project-specific conventions:
- Import order: stdlib, then third-party, then protobuf, then side-effect imports
- Error strings: lowercase, no trailing period
- Wrap errors with `%w` when callers need to inspect, `%v` at system boundaries

## Testing

Follow the testing principles advocated by t_wada (Takuto Wada).

- **Behavior-based naming**: `Test_MethodName_description_of_behavior` (e.g., `Test_ListDNSRecords_returns_error_when_zone_id_is_empty`)
- **Arrange-Act-Assert pattern**: separate each phase with a blank line
- **One assertion per test**: each test should have exactly one reason to fail
- **Boundary value testing**: always test `0`, `-1`, empty, exact boundary, boundary+1
- **Table-driven tests** for covering equivalence partitions and boundary values
- Use `cmp.Equal` / `cmp.Diff` for struct comparisons, not field-by-field checks
- Use `got`/`want` ordering in error messages
- Prefer `t.Error` over `t.Fatal` unless subsequent assertions would be meaningless

## Commit Messages

- Write commit messages in English
- PR descriptions in English
- Code comments in English
