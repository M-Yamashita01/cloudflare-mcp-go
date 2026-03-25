# cloudflare-mcp-go

An MCP (Model Context Protocol) server for the Cloudflare API, written in Go.

## Prerequisites

- Go 1.21+
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

### list_zones

Lists zones in your Cloudflare account.

**Parameters:**

| Parameter | Type | Required | Description |
|---|---|---|---|
| name | string | No | Filter by domain name |
| page | int | No | Page number of paginated results (default: 1) |
| per_page | int | No | Number of zones per page (default: 20, max: 50) |

## MCP Client Configuration Example

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
