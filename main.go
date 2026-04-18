package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/tool/account"
	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/tool/dns"
	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/tool/kv"
	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/tool/security"
	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/tool/zone"
)

func main() {
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "cloudflare-mcp-go",
			Version: "0.1.0",
		},
		nil,
	)

	zone.RegisterTools(server)
	dns.RegisterTools(server)
	account.RegisterTools(server)
	kv.RegisterTools(server)
	security.RegisterTools(server)

	log.Println("Starting Cloudflare MCP server (stdio)...")
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}
