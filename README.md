# cloudflare-mcp-go

[![Go Reference](https://pkg.go.dev/badge/github.com/M-Yamashita01/cloudflare-mcp-go.svg)](https://pkg.go.dev/github.com/M-Yamashita01/cloudflare-mcp-go)
[![Go Version](https://img.shields.io/github/go-mod/go-version/M-Yamashita01/cloudflare-mcp-go)](go.mod)
[![Latest Release](https://img.shields.io/github/v/release/M-Yamashita01/cloudflare-mcp-go)](https://github.com/M-Yamashita01/cloudflare-mcp-go/releases)

An MCP (Model Context Protocol) server for the Cloudflare API v4, written in Go.

It lets any MCP client (Claude, and other MCP-compatible tools) inspect and investigate your Cloudflare account through natural language: zones, DNS, WAF and firewall rules, rate limits, HTTP request logs, Security Center insights, and threat intelligence. The server is read-focused and geared toward security triage and investigation workflows.

## Quick start

Requires a Cloudflare API token and Go 1.25+.

Add it to Claude Code with a single command:

```bash
claude mcp add cloudflare \
  -e CLOUDFLARE_API_TOKEN=your-api-token \
  -- go run github.com/M-Yamashita01/cloudflare-mcp-go@v0.2.0
```

That is all — no manual clone or build required. `CLOUDFLARE_ACCOUNT_ID` is optional and only needed by a few tools (accounts, audit logs, KV, intel).

## Tools

28 tools grouped by domain.

### Zones & DNS

| Tool | Description |
|------|-------------|
| `list_zones` | List zones in your account (ID, name, status, plan). |
| `get_zone` | Get details of a specific zone. |
| `list_dns_records` | List DNS records for a zone (type, name, content, TTL, proxy status). |
| `get_dns_analytics` | DNS query analytics: query counts, response codes, query type distribution. |

### Firewall, WAF & rate limiting

| Tool | Description |
|------|-------------|
| `list_ip_access_rules` | IP access rules that block, challenge, or allow IPs, CIDRs, ASNs, or countries. |
| `list_waf_managed_rulesets` | WAF managed rulesets entrypoint and which managed rulesets are enabled. |
| `list_firewall_rules` | Custom firewall rules with expressions, actions, and priorities. |
| `get_firewall_rule` | Full configuration of a specific firewall rule by ID. |
| `list_rulesets` | All rulesets for a zone (metadata only). |
| `get_ruleset` | A specific ruleset with all its rules. |
| `list_rate_limits` | Rate limiting rules: thresholds, match criteria, actions. |
| `get_rate_limit` | Full configuration of a specific rate limiting rule by ID. |

### Security events & logs

| Tool | Description |
|------|-------------|
| `query_security_events` | Query security events via the GraphQL Analytics API for triage. |
| `get_log_by_rayid` | Look up an HTTP request log entry by Cloudflare Ray ID. |
| `list_received_logs` | Retrieve HTTP request logs for a time range (max 1h, data ≥5min old). |
| `list_log_fields` | List available HTTP request log fields. |

### Security Center

| Tool | Description |
|------|-------------|
| `list_security_insights` | Security Center insights with severity, type, and classification. |
| `get_insight_counts` | Aggregated insight counts by severity, class, or type. |

### Threat intelligence

| Tool | Description |
|------|-------------|
| `get_ip_intel` | Threat intelligence for an IP: geolocation, ASN, infrastructure, threat categories. |
| `get_domain_intel` | Security intelligence for a domain: risk scores, categories, DNS info. |
| `get_domain_intel_bulk` | Threat intelligence for multiple domains at once. |
| `get_domain_history` | Historical threat data and classifications for a domain. |
| `list_passive_dns` | Domains that have resolved to a given IP (passive DNS). |
| `get_whois` | WHOIS registration data for a domain. |
| `get_asn_intel` | ASN overview and subnet allocations. |

### Account & audit

| Tool | Description |
|------|-------------|
| `list_accounts` | Cloudflare accounts accessible with the current token. |
| `list_audit_logs` | Account audit log entries (who changed what and when). |
| `list_kv_namespaces` | Workers KV namespaces in an account. |

## Usage examples

Once connected, ask your MCP client things like:

- "List the DNS records for example.com."
- "A request was blocked with Ray ID 8f1c2d3e4f5a6b7c — why?"
- "Show me the high-severity Security Center insights for this zone."
- "Is IP 203.0.113.10 known to be malicious? What domains resolve to it?"
- "Which firewall rules are currently challenging traffic?"

## Configuration

The server communicates over stdio and can be connected from any MCP client.

| Variable | Required | Purpose |
|----------|----------|---------|
| `CLOUDFLARE_API_TOKEN` | Yes | Authenticates all Cloudflare API calls. |
| `CLOUDFLARE_ACCOUNT_ID` | No | Needed by account, audit log, KV, and intel tools. |

### Claude Code

```bash
claude mcp add cloudflare \
  -e CLOUDFLARE_API_TOKEN=your-api-token \
  -- go run github.com/M-Yamashita01/cloudflare-mcp-go@v0.2.0
```

Or, if you have the repository cloned locally:

```bash
claude mcp add cloudflare -e CLOUDFLARE_API_TOKEN=your-api-token -- go run .
```

### Claude Desktop / other MCP clients

Build the binary and point your client's MCP config at it. On macOS the Claude Desktop config lives at `~/Library/Application Support/Claude/claude_desktop_config.json`:

```bash
go build -o cloudflare-mcp-go .
```

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

Any MCP-compatible client can connect via stdio transport using the same pattern.

## Build & run manually

```bash
go build -o cloudflare-mcp-go .
export CLOUDFLARE_API_TOKEN="your-api-token"
./cloudflare-mcp-go
```

## Project structure

```
main.go                  # Entry point
internal/
  cfapi/                 # Shared Cloudflare API client
  tool/
    zone/                # Zone management tools
    dns/                 # DNS record & analytics tools
    account/             # Account management tools
    audit/               # Audit log tools
    kv/                  # Workers KV tools
    security/            # Firewall, WAF & rate limiting tools
    securitycenter/      # Security Center insight tools
    intel/               # Threat intelligence tools
    logs/                # HTTP request log tools
doc/
  architecture.md        # Architecture documentation
```

See [doc/architecture.md](doc/architecture.md) for detailed design documentation.
