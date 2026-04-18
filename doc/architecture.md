# Architecture

## Overview

cloudflare-mcp-go is an MCP (Model Context Protocol) server that exposes Cloudflare API v4 operations as tools. It runs over stdio transport and is designed for use with AI assistants.

## Directory Structure

```
cloudflare-mcp-go/
├── main.go                          # Entry point: server init + RegisterTools calls
├── internal/
│   ├── cfapi/
│   │   ├── cfapi.go                 # Shared Cloudflare API client (types + helpers)
│   │   └── cfapi_test.go
│   └── tool/
│       ├── zone/                    # Zone management (list_zones, get_zone)
│       ├── dns/                     # DNS records (list_dns_records)
│       ├── account/                 # Account management (list_accounts)
│       ├── kv/                      # Workers KV (list_kv_namespaces)
│       └── security/               # Firewall & WAF (list_ip_access_rules,
│                                    #   list_waf_managed_rulesets, query_security_events)
├── doc/
│   └── architecture.md             # This file
├── CLAUDE.md                        # AI assistant coding conventions
├── .golangci.yml                    # Linter configuration
├── go.mod
└── go.sum
```

## Design Principles

1. **Domain-based package organization** — Each Cloudflare API category has its own package under `internal/tool/`. This keeps related code together and limits the blast radius of changes.

2. **Thin entry point** — `main.go` only initializes the MCP server and calls `RegisterTools` for each domain package. No business logic lives here.

3. **Shared API client** — `internal/cfapi` provides the HTTP client, response types, and common helpers used by all domain packages. Type names use the package qualifier (e.g., `cfapi.Response`) to avoid redundant prefixes.

4. **`internal/` boundary** — All implementation packages are under `internal/` since this is a CLI application, not a library. This prevents accidental external imports.

5. **Co-located tests** — Each package has its own `*_test.go` file using white-box testing (same package).

## Package Dependencies

```
main.go
  ├── internal/tool/zone       ─┐
  ├── internal/tool/dns         │
  ├── internal/tool/account     ├── all depend on ──▶ internal/cfapi
  ├── internal/tool/kv          │
  └── internal/tool/security   ─┘
```

Domain packages (`internal/tool/*`) do NOT depend on each other. They only depend on `internal/cfapi` and the MCP SDK.

## Data Flow

```
AI Assistant
    │
    ▼ (stdio JSON-RPC)
┌─────────────────────────┐
│  MCP Server (main.go)   │
│  - mcp.NewServer()      │
│  - RegisterTools()      │
└────────┬────────────────┘
         │ dispatches to handler
         ▼
┌─────────────────────────┐
│  Domain Handler          │
│  (e.g., zone.list)      │
│  1. Read API token       │
│  2. Build request URL    │
│  3. Call cfapi.DoRequest │
│  4. Format response      │
└────────┬────────────────┘
         │ HTTP/HTTPS
         ▼
┌─────────────────────────┐
│  Cloudflare API v4       │
│  - REST: /client/v4/...  │
│  - GraphQL: /graphql     │
└─────────────────────────┘
```

## Adding a New Tool

To add a new Cloudflare API tool:

1. **Create a package** (or add to an existing one):
   ```
   internal/tool/<domain>/<domain>.go
   internal/tool/<domain>/<domain>_test.go
   ```

2. **Define the input struct**:
   ```go
   type ListInput struct {
       ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
       // ... other fields
   }
   ```

3. **Implement the handler** (unexported):
   ```go
   func list(ctx context.Context, _ *mcp.CallToolRequest, input ListInput) (*mcp.CallToolResult, any, error) {
       apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
       if result := cfapi.CheckToken(apiToken); result != nil {
           return result, nil, nil
       }
       // ... build URL, call cfapi.DoRequest, format result
   }
   ```

4. **Add `RegisterTools`** (or extend an existing one):
   ```go
   func RegisterTools(server *mcp.Server) {
       mcp.AddTool(server, &mcp.Tool{
           Name:        "tool_name",
           Description: "Tool description.",
       }, list)
   }
   ```

5. **Register in `main.go`**:
   ```go
   import "github.com/M-Yamashita01/cloudflare-mcp-go/internal/tool/<domain>"
   // ...
   <domain>.RegisterTools(server)
   ```

6. **Write tests** following t_wada principles (see CLAUDE.md).

## API Communication

- **REST API**: Uses `cfapi.DoRequest()` which handles HTTP request construction, Bearer token auth, and response parsing into `cfapi.Response`.
- **GraphQL API**: The `security` package implements GraphQL directly for `query_security_events`. If more GraphQL tools are needed in other domains, the GraphQL client logic should be extracted to `internal/cfapi`.

## Environment Variables

| Variable | Required | Description |
|----------|----------|-------------|
| `CLOUDFLARE_API_TOKEN` | Yes | Cloudflare API token for authentication |
| `CLOUDFLARE_ACCOUNT_ID` | No | Used by some tools (e.g., KV namespaces) passed as input |
