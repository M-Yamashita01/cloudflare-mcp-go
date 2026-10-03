package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/tool/account"
	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/tool/audit"
	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/tool/dns"
	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/tool/intel"
	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/tool/kv"
	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/tool/logs"
	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/tool/security"
	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/tool/securitycenter"
	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/tool/zone"
)

var version = "dev"

func main() {
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "cloudflare-mcp-go",
			Version: version,
		},
		nil,
	)

	writeEnabled, warn := cfapi.WriteEnabled()
	if warn != "" {
		log.Println(warn)
	}

	// Read-only tools are always registered.
	zone.RegisterTools(server)
	dns.RegisterTools(server)
	account.RegisterTools(server)
	audit.RegisterTools(server)
	kv.RegisterTools(server)
	intel.RegisterTools(server)
	logs.RegisterTools(server)
	security.RegisterTools(server)
	securitycenter.RegisterTools(server)

	// Write (mutation) tools are registered only when enabled via
	// CLOUDFLARE_MCP_ENABLE_WRITE=true, so they stay out of tools/list by
	// default. No write tools exist yet; they will be wired here
	// (e.g. dns.RegisterWriteTools(server)). See issue #107.
	if writeEnabled {
		log.Println("Write tools enabled via " + cfapi.EnableWriteEnv)
	}

	log.Println("Starting Cloudflare MCP server (stdio)...")
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}
