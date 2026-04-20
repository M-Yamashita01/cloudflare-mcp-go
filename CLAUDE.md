# CLAUDE.md

## Project Overview

Cloudflare MCP Go — A Go implementation of an MCP (Model Context Protocol) server that wraps Cloudflare API v4.

- Language: Go 1.23+
- MCP SDK: `github.com/modelcontextprotocol/go-sdk`
- Entry point: `main.go`
- Architecture: See [doc/architecture.md](doc/architecture.md) for detailed design docs

## Project Structure

```
main.go                          # Thin entry point (server init + RegisterTools)
internal/
  cfapi/                         # Shared Cloudflare API client (types + helpers)
  tool/
    zone/                        # list_zones, get_zone
    dns/                         # list_dns_records
    account/                     # list_accounts
    kv/                          # list_kv_namespaces
    security/                    # list_ip_access_rules, list_waf_managed_rulesets, query_security_events
doc/
  architecture.md                # Architecture documentation
```

### Adding a New Tool

1. Create or extend a package under `internal/tool/<domain>/`
2. Define Input struct + handler function + `RegisterTools()`
3. Add `<domain>.RegisterTools(server)` in `main.go`
4. Write co-located tests in `*_test.go`

## Build & Run

```bash
# Build
go build -o cloudflare-mcp-go .

# Run (requires CLOUDFLARE_API_TOKEN and CLOUDFLARE_ACCOUNT_ID env vars)
go run .

# Test
go test ./...

# Format
gofmt -w .

# Vet
go vet ./...
```

## Go Style Guide

Follow the [Google Go Style Guide](https://google.github.io/styleguide/go/index) as the baseline for all Go code in this project.

### Key Principles (in priority order)

1. **Clarity** — Purpose and rationale are clear to readers
2. **Simplicity** — Achieve the goal in the simplest way
3. **Concision** — High signal-to-noise ratio
4. **Maintainability** — Easy to change in the future
5. **Consistency** — Coherent with the rest of the codebase

### Naming

- Use `MixedCaps` / `mixedCaps`, never `snake_case` (except in test names, see Testing section)
- Constants use `MixedCaps`, not `ALL_CAPS`
- Package names: lowercase, single word, concise. Avoid `util`, `common`, `helper`
- Receiver names: 1-2 character abbreviation of the type name, consistent across methods
- Variable name length proportional to scope
- No `Get` prefix on getters: `Counts()` not `GetCounts()`
- Acronyms: all caps or all lower (`URL`, `ID`, `userID`)

### Formatting

- All code must pass `gofmt`
- Use field names in struct literals
- Handle errors first, then normal flow (no `else` after `if err != nil { return }`)

### Error Handling

- `error` is always the last return value
- Error strings: lowercase, no trailing period
- Use `errors.Is()` / `errors.As()`, never string matching
- Wrap with `%w` when callers need to inspect, `%v` at system boundaries
- Never `panic` in library code; return `error`
- Comment why when intentionally ignoring an error

### Package Design

- Import order: stdlib → third-party → protobuf → side-effect imports
- Define interfaces at the consumer side, keep them small
- Accept interfaces, return concrete types

### Other Rules

- `context.Context` is always the first parameter
- Prefer synchronous functions over async
- Pre-allocate memory only when benchmarks justify it

## Testing Guidelines — t_wada Style

Follow the testing principles advocated by t_wada (Takuto Wada).

### Core Principles

1. **Behavior-based test naming**
   - Name tests by the behavior under test: `Test_MethodName_description_of_behavior`
   - Example: `Test_ListDNSRecords_returns_error_when_zone_id_is_empty`

2. **Boundary value testing**
   - Always test boundary values: `0`, `-1`, empty, exact boundary, boundary+1, etc.
   - Think about edge cases at every boundary

3. **Equivalence partitioning**
   - Pick representative values from each partition
   - Don't test every value — test one from each equivalence class

4. **Arrange-Act-Assert (AAA pattern)**
   - Structure every test in three phases:
     - **Arrange** — Set up preconditions and inputs
     - **Act** — Execute the behavior under test
     - **Assert** — Verify the expected outcome
   - Separate each phase with a blank line for readability

5. **One assertion per test (principle)**
   - Each test should have exactly one reason to fail
   - If a test fails, the cause should be immediately obvious
   - Use table-driven tests to cover multiple cases without multiple assertions in one test

6. **Design for testability**
   - Use interfaces (Protocol) + Dependency Injection to make code mockable
   - Accept interfaces, return concrete types
   - Inject dependencies through constructor parameters

### Test Code Style

```go
func Test_MethodName_description_of_behavior(t *testing.T) {
    // Arrange
    input := "test-input"
    expected := "expected-output"

    // Act
    result := MethodName(input)

    // Assert
    if result != expected {
        t.Errorf("got %v, want %v", result, expected)
    }
}
```

### Table-Driven Tests

Use table-driven tests for covering equivalence partitions and boundary values:

```go
func Test_Validate_boundary_values(t *testing.T) {
    tests := []struct {
        name    string
        input   int
        wantErr bool
    }{
        {"exact_boundary_0", 0, false},
        {"boundary_minus_1", -1, true},
        {"boundary_plus_1", 1, false},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := Validate(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("Validate(%d) error = %v, wantErr %v", tt.input, err, tt.wantErr)
            }
        })
    }
}
```

### Testing Do's and Don'ts

- Use `cmp.Equal` / `cmp.Diff` for struct comparisons, not field-by-field checks
- Use `got`/`want` ordering in error messages: actual first, expected second
- Prefer `t.Error` over `t.Fatal` to report all failures
- Use `t.Fatal` only when subsequent assertions would be meaningless
- Do not call `t.Fatal` from goroutines — use `t.Error` and return

## Commit Messages

- Write commit messages in English
- PR descriptions in English
- Code comments in English
