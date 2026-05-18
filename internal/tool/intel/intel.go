// Package intel provides MCP tools for Cloudflare threat intelligence.
package intel

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// GetIPIntelInput holds parameters for retrieving IP threat intelligence.
type GetIPIntelInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	IP        string `json:"ip" jsonschema:"required,The IPv4 or IPv6 address to look up"`
}

func getIPIntel(ctx context.Context, _ *mcp.CallToolRequest, input GetIPIntelInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := fmt.Sprintf("%s/accounts/%s/intel/ip?ipv4=%s", cfapi.APIBase, input.AccountID, input.IP)

	cfResp, err := cfapi.DoRequest(ctx, http.MethodGet, url, apiToken, nil)
	if err != nil {
		return nil, nil, err
	}
	if !cfResp.Success {
		return cfapi.APIErrorResult(cfResp.Errors), nil, nil
	}

	result, err := cfapi.FormatResult(cfResp)
	if err != nil {
		return nil, nil, err
	}
	return result, nil, nil
}

// RegisterTools registers threat intelligence tools with the MCP server.
func RegisterTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_ip_intel",
		Description: "Get threat intelligence for an IP address. Returns geolocation, ASN, infrastructure type, and security threat categories. Useful for investigating suspicious IPs found in security events or access logs.",
	}, getIPIntel)
}
