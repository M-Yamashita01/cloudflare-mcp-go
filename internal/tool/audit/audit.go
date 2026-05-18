// Package audit provides MCP tools for Cloudflare audit log investigation.
package audit

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// ListAuditLogsInput holds query parameters for listing audit logs.
type ListAuditLogsInput struct {
	AccountID  string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Since      string `json:"since,omitempty" jsonschema:"Filter by start date in ISO 8601 format (e.g. 2026-05-01T00:00:00Z)"`
	Before     string `json:"before,omitempty" jsonschema:"Filter by end date in ISO 8601 format"`
	ActorIP    string `json:"actor_ip,omitempty" jsonschema:"Filter by the IP address of the actor"`
	ActorEmail string `json:"actor_email,omitempty" jsonschema:"Filter by the email of the actor"`
	Direction  string `json:"direction,omitempty" jsonschema:"Sort direction: asc or desc (default: desc)"`
	Page       int    `json:"page,omitempty" jsonschema:"Page number of paginated results (default: 1)"`
	PerPage    int    `json:"per_page,omitempty" jsonschema:"Number of entries per page (default: 25, max: 1000)"`
}

func listAuditLogs(ctx context.Context, _ *mcp.CallToolRequest, input ListAuditLogsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/accounts/" + input.AccountID + "/audit_logs"
	var params []string
	if input.Since != "" {
		params = append(params, fmt.Sprintf("since=%s", input.Since))
	}
	if input.Before != "" {
		params = append(params, fmt.Sprintf("before=%s", input.Before))
	}
	if input.ActorIP != "" {
		params = append(params, fmt.Sprintf("actor.ip=%s", input.ActorIP))
	}
	if input.ActorEmail != "" {
		params = append(params, fmt.Sprintf("actor.email=%s", input.ActorEmail))
	}
	if input.Direction != "" {
		params = append(params, fmt.Sprintf("direction=%s", input.Direction))
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

// RegisterTools registers audit log tools with the MCP server.
func RegisterTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_audit_logs",
		Description: "List audit log entries for a Cloudflare account. Returns operation history including timestamps, actors, actions, and affected resources. Useful for investigating who changed what and when.",
	}, listAuditLogs)
}
