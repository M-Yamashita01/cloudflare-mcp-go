// Package dns provides MCP tools for Cloudflare DNS record management.
package dns

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// ListInput holds query parameters for listing DNS records.
type ListInput struct {
	ZoneID  string `json:"zone_id"            jsonschema:"required,The ID of the zone"`
	Type    string `json:"type,omitempty"     jsonschema:"DNS record type to filter by (A, AAAA, CNAME, TXT, MX, etc.)"`
	Name    string `json:"name,omitempty"     jsonschema:"DNS record name to filter by"`
	Content string `json:"content,omitempty"  jsonschema:"DNS record content to filter by"`
	Page    int    `json:"page,omitempty"     jsonschema:"Page number of paginated results (default: 1)"`
	PerPage int    `json:"per_page,omitempty" jsonschema:"Number of records per page (default: 100, max: 5000)"`
}

func list(ctx context.Context, _ *mcp.CallToolRequest, input ListInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_records"
	var params []string
	if input.Type != "" {
		params = append(params, fmt.Sprintf("type=%s", input.Type))
	}
	if input.Name != "" {
		params = append(params, fmt.Sprintf("name=%s", input.Name))
	}
	if input.Content != "" {
		params = append(params, fmt.Sprintf("content=%s", input.Content))
	}
	if input.Page > 0 {
		params = append(params, fmt.Sprintf("page=%d", input.Page))
	}
	if input.PerPage > 0 {
		params = append(params, fmt.Sprintf("per_page=%d", input.PerPage))
	}
	if len(params) > 0 {
		url += "?" + strings.Join(params, "&")
	}

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

// RegisterTools registers DNS management tools with the MCP server.
func RegisterTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_dns_records",
		Description: "List DNS records for a Cloudflare zone. Returns record details such as ID, type, name, content, TTL, and proxy status.",
	}, list)
}
