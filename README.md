# cloudflare-mcp-go

[![Go Reference](https://pkg.go.dev/badge/github.com/M-Yamashita01/cloudflare-mcp-go.svg)](https://pkg.go.dev/github.com/M-Yamashita01/cloudflare-mcp-go)
[![Go Version](https://img.shields.io/github/go-mod/go-version/M-Yamashita01/cloudflare-mcp-go)](go.mod)
[![Latest Release](https://img.shields.io/github/v/release/M-Yamashita01/cloudflare-mcp-go)](https://github.com/M-Yamashita01/cloudflare-mcp-go/releases)
[![Glama score](https://glama.ai/mcp/servers/M-Yamashita01/cloudflare-mcp-go/badges/score.svg)](https://glama.ai/mcp/servers/M-Yamashita01/cloudflare-mcp-go)

An MCP (Model Context Protocol) server for the Cloudflare API v4, written in Go.

It lets any MCP client (Claude, and other MCP-compatible tools) inspect and investigate your Cloudflare account through natural language: zones, DNS, WAF and firewall rules, rate limits, HTTP request logs, Security Center insights, and threat intelligence. The server is read-focused and geared toward security triage and investigation workflows.

## Quick start

Requires a Cloudflare API token and Go 1.25+.

Add it to Claude Code with a single command:

```bash
claude mcp add cloudflare \
  -e CLOUDFLARE_API_TOKEN=your-api-token \
  -- go run github.com/M-Yamashita01/cloudflare-mcp-go@latest
```

That is all — no manual clone or build required. `CLOUDFLARE_ACCOUNT_ID` is optional and only needed by a few tools (accounts, audit logs, KV, intel).

## Tools

80 read-only tools grouped by domain, plus optional write tools (marked below) that are registered only when `CLOUDFLARE_MCP_ENABLE_WRITE=true`.

### Zones & DNS

| Tool | Description |
|------|-------------|
| `list_zones` | List zones in your account (ID, name, status, plan). |
| `get_zone` | Get details of a specific zone. |
| `get_zone_settings` | Get a zone's configuration settings (SSL mode, min TLS, cache level, security level). |
| `purge_cache` | Purge a zone's cache: everything, or by file/tag/host/prefix (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `list_dns_records` | List DNS records for a zone (type, name, content, TTL, proxy status). |
| `get_dns_record` | Get a single DNS record by ID (type, name, content, TTL, proxy status, comment). |
| `get_dns_analytics` | DNS query analytics: query counts, response codes, query type distribution. |
| `get_dns_analytics_bytime` | DNS query analytics grouped by time interval (time-series). |
| `export_dns_records` | Export all DNS records for a zone as a BIND zone file. |
| `get_dns_record_usage` | DNS record usage counts for a zone (toward plan limits). |
| `get_dns_settings` | Get a zone's DNS settings (Foundation DNS, multi-provider, nameservers, zone mode). |
| `get_dnssec` | Get DNSSEC details for a zone (status, DS record, digest, key tag, algorithm). |
| `list_dnssec_zsk` | List DNSSEC Zone Signing Keys (ZSKs) for a zone (multi-signer DNSSEC). |
| `list_scanned_dns_records` | List DNS records discovered by a scan, pending review. |
| `create_dns_record` | Create a DNS record (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `update_dns_record` | Update a DNS record, partial fields (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `delete_dns_record` | Delete a DNS record by ID (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `overwrite_dns_record` | Overwrite (fully replace) a DNS record by ID (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `batch_dns_records` | Apply a batch of DNS record changes atomically (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `import_dns_records` | Import DNS records from a BIND zone file (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `scan_dns_records` | Scan a zone for DNS records at common names (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `trigger_dns_record_scan` | Trigger an async DNS record scan (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `review_scanned_dns_records` | Accept/reject DNS records found by a scan (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `update_dns_settings` | Update a zone's DNS settings (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `edit_dnssec` | Edit a zone's DNSSEC status (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `delete_dnssec` | Delete a zone's DNSSEC records (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |

### Firewall, WAF & rate limiting

| Tool | Description |
|------|-------------|
| `list_ip_access_rules` | IP access rules that block, challenge, or allow IPs, CIDRs, ASNs, or countries. |
| `create_ip_access_rule` | Create an IP access rule (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `update_ip_access_rule` | Update an IP access rule's mode/notes (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `delete_ip_access_rule` | Delete an IP access rule by ID (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `list_waf_managed_rulesets` | WAF managed rulesets entrypoint and which managed rulesets are enabled. |
| `list_firewall_rules` | Custom firewall rules with expressions, actions, and priorities. |
| `get_firewall_rule` | Full configuration of a specific firewall rule by ID. |
| `list_rulesets` | All rulesets for a zone (metadata only). |
| `get_ruleset` | A specific ruleset with all its rules. |
| `create_ruleset` | Create a ruleset (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `update_ruleset` | Update a ruleset's name/description/rules (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `delete_ruleset` | Delete a ruleset by ID (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `list_rate_limits` | Rate limiting rules: thresholds, match criteria, actions. |
| `get_rate_limit` | Full configuration of a specific rate limiting rule by ID. |
| `create_rate_limit` | Create a rate limit rule (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `update_rate_limit` | Update a rate limit rule (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `delete_rate_limit` | Delete a rate limit rule by ID (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |

### Security events & logs

| Tool | Description |
|------|-------------|
| `query_security_events` | Query security events via the GraphQL Analytics API for triage. |
| `get_log_by_rayid` | Look up an HTTP request log entry by Cloudflare Ray ID. |
| `list_received_logs` | Retrieve HTTP request logs for a time range (max 1h, data ≥5min old). |
| `list_log_fields` | List available HTTP request log fields. |
| `list_logpush_jobs` | List Logpush jobs for a zone (dataset, destination, enabled). |
| `get_logpush_job` | Get a specific Logpush job in a zone by ID. |
| `list_logpush_dataset_jobs` | List Logpush jobs for a specific dataset in a zone. |
| `list_logpush_dataset_fields` | List available Logpush fields for a dataset in a zone. |
| `list_instant_logs_jobs` | List Instant Logs (edge) jobs for a zone. |
| `get_log_retention_flag` | Get a zone's log retention flag. |
| `list_log_datasets` | List Logs Explorer datasets for a zone. |
| `list_available_log_datasets` | List available Logs Explorer datasets for a zone. |
| `get_log_dataset` | Get a Logs Explorer dataset for a zone by ID. |
| `query_logs_sql` | Run a Logs Explorer SQL query for a zone (GET). |
| `list_account_logpush_jobs` | List Logpush jobs for an account. |
| `get_account_logpush_job` | Get a specific Logpush job in an account by ID. |
| `list_account_logpush_dataset_jobs` | List Logpush jobs for a dataset in an account. |
| `list_account_logpush_dataset_fields` | List available Logpush fields for a dataset in an account. |
| `list_account_logpush_transformers` | List Logpush transformers for an account. |
| `get_account_logpush_transformer` | Get a Logpush transformer in an account by ID. |
| `get_account_logpush_transformer_content` | Get a Logpush transformer's content in an account. |
| `list_account_logpush_transformer_versions` | List versions of a Logpush transformer in an account. |
| `get_account_audit_logs_v2` | Get account audit logs (v2) with since/before/limit. |
| `list_account_audit_log_product_categories` | List product categories for account audit logs (v2). |
| `get_account_audit_log_history` | Get change history for an account audit log entry (v2). |
| `get_cmb_config` | Get the Customer Metadata Boundary (CMB) config for an account. |
| `list_account_log_datasets` | List Logs Explorer datasets for an account. |
| `list_available_account_log_datasets` | List available Logs Explorer datasets for an account. |
| `get_account_log_dataset` | Get a Logs Explorer dataset for an account by ID. |
| `query_account_logs_sql` | Run a Logs Explorer SQL query for an account (GET). |
| `list_account_log_files` | List stored log files for an account (start/end/bucket). |
| `retrieve_account_logs` | Retrieve stored log entries (NDJSON) for an account. |
| `create_account_logpush_job` | Create an account Logpush job (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `update_account_logpush_job` | Update an account Logpush job (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `delete_account_logpush_job` | Delete an account Logpush job (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `get_account_logpush_ownership_challenge` | Request an account Logpush ownership challenge (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `validate_account_logpush_ownership` | Validate an account Logpush ownership challenge (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `create_account_logpush_transformer` | Create an account Logpush transformer (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `preview_account_logpush_transformer` | Preview an account Logpush transformer (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `update_account_logpush_transformer` | Update an account Logpush transformer (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `delete_account_logpush_transformer` | Delete an account Logpush transformer (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `validate_account_logpush_destination` | Validate an account Logpush destination (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `check_account_logpush_destination_exists` | Check an account Logpush destination exists (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `validate_account_logpush_origin` | Validate account Logpush origin options (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `update_cmb_config` | Update the Customer Metadata Boundary config (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `delete_cmb_config` | Delete the Customer Metadata Boundary config (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `create_account_log_dataset` | Create an account Logs Explorer dataset (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `update_account_log_dataset` | Update an account Logs Explorer dataset (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `delete_account_log_dataset` | Delete an account Logs Explorer dataset (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `run_account_logs_sql_query` | Run an account Logs Explorer SQL query (POST) (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `create_logpush_job` | Create a zone Logpush job (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `update_logpush_job` | Update a zone Logpush job (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `delete_logpush_job` | Delete a zone Logpush job (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `create_instant_logs_job` | Create a zone Instant Logs job (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `get_logpush_ownership_challenge` | Request a zone Logpush ownership challenge (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `validate_logpush_ownership` | Validate a zone Logpush ownership challenge (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `validate_logpush_destination` | Validate a zone Logpush destination (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `check_logpush_destination_exists` | Check a zone Logpush destination exists (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `validate_logpush_origin` | Validate zone Logpush origin options (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `update_log_retention_flag` | Enable/disable zone log retention (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `create_log_dataset` | Create a zone Logs Explorer dataset (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `update_log_dataset` | Update a zone Logs Explorer dataset (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `delete_log_dataset` | Delete a zone Logs Explorer dataset (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `run_logs_sql_query` | Run a zone Logs Explorer SQL query (POST) (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |

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
| `get_account_details` | Get details of a specific account by ID. |
| `list_account_members` | List members of an account (user, roles, status). |
| `get_account_member` | Get an account member by ID (user, roles, policies). |
| `list_account_roles` | List roles available in an account. |
| `get_account_role` | Get an account role by ID (name, permissions). |
| `list_account_subscriptions` | List subscriptions for an account. |
| `get_account_subscription` | Get an account subscription by ID. |
| `get_subscription_cancel_reason` | Get a subscription's cancellation reason. |
| `list_account_tokens` | List account-owned API tokens. |
| `list_audit_logs` | Account audit log entries (who changed what and when). |
| `list_user_audit_logs` | User-scoped audit log entries. |
| `list_kv_namespaces` | Workers KV namespaces in an account. |
| `get_kv_namespace` | Get a KV namespace by ID. |
| `get_kv_pairs_bulk` | Get multiple KV pairs at once (read-only; POST bulk/get). |
| `list_kv_keys` | List keys in a KV namespace (prefix filter, cursor pagination). |
| `get_kv_value` | Read the raw stored value for a key in a KV namespace. |
| `get_kv_metadata` | Read the metadata associated with a key in a KV namespace. |
| `write_kv_pair` | Write a key-value pair, optional TTL (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `delete_kv_pair` | Delete a key-value pair by key (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `create_kv_namespace` | Create a KV namespace (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `delete_kv_namespace` | Delete a KV namespace by ID (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `rename_kv_namespace` | Rename a KV namespace (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `write_kv_pairs_bulk` | Write multiple KV pairs at once (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `delete_kv_pairs_bulk` | Delete multiple KV pairs at once, DELETE bulk (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |
| `delete_kv_pairs_bulk_post` | Delete multiple KV pairs at once, POST bulk/delete (write; requires `CLOUDFLARE_MCP_ENABLE_WRITE=true`). |

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
| `CLOUDFLARE_MCP_ENABLE_WRITE` | No | Enables write (mutation) tools when set to exactly `true`. Unset or empty keeps them disabled (default); any other value is rejected (logged to stderr) and the server stays read-only. |

### Claude Code

```bash
claude mcp add cloudflare \
  -e CLOUDFLARE_API_TOKEN=your-api-token \
  -- go run github.com/M-Yamashita01/cloudflare-mcp-go@latest
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
