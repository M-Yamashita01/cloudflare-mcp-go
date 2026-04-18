// Package zone provides MCP tools for Cloudflare zone management.
package zone

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// ListInput holds query parameters for listing zones.
type ListInput struct {
	Name    string `json:"name,omitempty"     jsonschema:"A domain name to filter zones by"`
	Page    int    `json:"page,omitempty"     jsonschema:"Page number of paginated results (default: 1)"`
	PerPage int    `json:"per_page,omitempty" jsonschema:"Number of zones per page (default: 20, max: 50)"`
}

func list(ctx context.Context, _ *mcp.CallToolRequest, input ListInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones"
	var params []string
	if input.Name != "" {
		params = append(params, fmt.Sprintf("name=%s", input.Name))
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

// GetInput holds parameters for retrieving a single zone.
type GetInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone to retrieve"`
}

func get(ctx context.Context, _ *mcp.CallToolRequest, input GetInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID

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

// RegisterTools registers zone management tools with the MCP server.
func RegisterTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_zones",
		Description: "List zones in your Cloudflare account. Returns zone details such as ID, name, status, and plan.",
	}, list)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_zone",
		Description: "Get details of a specific Cloudflare zone. Returns zone details such as ID, name, status, and plan.",
	}, get)
}
