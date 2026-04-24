# cloudflare-mcp-go

An MCP (Model Context Protocol) server for the Cloudflare API, written in Go.

## Prerequisites

- Go 1.23+
- Cloudflare API token

## Setup

```bash
go build -o cloudflare-mcp-go .
```

## Usage

Set the `CLOUDFLARE_API_TOKEN` environment variable and run the server.

```bash
export CLOUDFLARE_API_TOKEN="your-api-token"
./cloudflare-mcp-go
```

The server communicates over stdio transport and can be connected to from any MCP client.

## Available Tools

| Tool | Description |
|------|-------------|
| `list_zones` | List zones in your Cloudflare account |
| `get_zone` | Get details of a specific zone |
| `list_dns_records` | List DNS records for a zone |
| `list_accounts` | List Cloudflare accounts |
| `list_kv_namespaces` | List Workers KV namespaces |
| `list_ip_access_rules` | List IP access rules for a zone |
| `list_waf_managed_rulesets` | Get WAF managed rulesets configuration |
| `query_security_events` | Query security events via GraphQL Analytics API (auto-fallback for zones without Bot Management) |

## MCP Client Configuration

### Claude Code

Install directly from the repository (no manual download required):

```bash
claude mcp add cloudflare -- go run github.com/M-Yamashita01/cloudflare-mcp-go@v0.1.1
```

Or if you have the repository cloned locally:

```bash
claude mcp add cloudflare -- go run .
```

Or manually add to your MCP config (e.g., `~/.claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "cloudflare": {
      "command": "/path/to/cloudflare-mcp-go",
      "env": {
        "CLOUDFLARE_API_TOKEN": "your-api-token"
      }
    }
  }
}
```

### Claude Desktop

Add to `~/Library/Application Support/Claude/claude_desktop_config.json` (macOS) or `%APPDATA%\Claude\claude_desktop_config.json` (Windows):

```json
{
  "mcpServers": {
    "cloudflare": {
      "command": "/path/to/cloudflare-mcp-go",
      "env": {
        "CLOUDFLARE_API_TOKEN": "your-api-token"
      }
    }
  }
}
```

### Other MCP Clients

Any MCP-compatible client can connect via stdio transport using the same configuration pattern.

## Project Structure

```
main.go                  # Entry point
internal/
  cfapi/                 # Shared Cloudflare API client
  tool/
    zone/                # Zone management tools
    dns/                 # DNS record tools
    account/             # Account management tools
    kv/                  # Workers KV tools
    security/            # Firewall & WAF tools
doc/
  architecture.md        # Architecture documentation
```

See [doc/architecture.md](doc/architecture.md) for detailed design documentation.
