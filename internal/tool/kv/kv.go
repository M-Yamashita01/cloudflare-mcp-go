// Package kv provides MCP tools for Cloudflare Workers KV management.
package kv

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// ListNamespacesInput holds query parameters for listing KV namespaces.
type ListNamespacesInput struct {
	AccountID string `json:"account_id"          jsonschema:"required,The ID of the Cloudflare account"`
	Page      int    `json:"page,omitempty"      jsonschema:"Page number of paginated results (default: 1)"`
	PerPage   int    `json:"per_page,omitempty"  jsonschema:"Number of namespaces per page (default: 20, max: 100)"`
	Order     string `json:"order,omitempty"     jsonschema:"Order results by field: id or title"`
	Direction string `json:"direction,omitempty" jsonschema:"Sort direction: asc or desc"`
}

func listNamespaces(ctx context.Context, _ *mcp.CallToolRequest, input ListNamespacesInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/accounts/" + input.AccountID + "/storage/kv/namespaces"
	var params []string
	if input.Page > 0 {
		params = append(params, fmt.Sprintf("page=%d", input.Page))
	}
	if input.PerPage > 0 {
		params = append(params, fmt.Sprintf("per_page=%d", input.PerPage))
	}
	if input.Order != "" {
		params = append(params, fmt.Sprintf("order=%s", input.Order))
	}
	if input.Direction != "" {
		params = append(params, fmt.Sprintf("direction=%s", input.Direction))
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

// RegisterTools registers KV management tools with the MCP server.
func RegisterTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_kv_namespaces",
		Description: "List Workers KV namespaces in a Cloudflare account. Returns namespace details such as ID and title.",
	}, listNamespaces)
}
