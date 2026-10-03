// Package securitycenter provides MCP tools for Cloudflare Security Center.
package securitycenter

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// ListInsightsInput holds query parameters for listing security insights.
type ListInsightsInput struct {
	ZoneID     string `json:"zone_id" jsonschema:"required,The ID of the zone"`
	Severity   string `json:"severity,omitempty" jsonschema:"Filter by severity: critical, high, moderate, low, informational"`
	IssueType  string `json:"issue_type,omitempty" jsonschema:"Filter by issue type"`
	IssueClass string `json:"issue_class,omitempty" jsonschema:"Filter by issue class"`
}

func listInsights(ctx context.Context, _ *mcp.CallToolRequest, input ListInsightsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/security-center/insights"
	var params []string
	if input.Severity != "" {
		params = append(params, fmt.Sprintf("severity=%s", input.Severity))
	}
	if input.IssueType != "" {
		params = append(params, fmt.Sprintf("issue_type=%s", input.IssueType))
	}
	if input.IssueClass != "" {
		params = append(params, fmt.Sprintf("issue_class=%s", input.IssueClass))
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

// GetInsightCountsInput holds parameters for retrieving insight counts.
type GetInsightCountsInput struct {
	ZoneID    string `json:"zone_id" jsonschema:"required,The ID of the zone"`
	Dimension string `json:"dimension" jsonschema:"required,Aggregation dimension: severity or class or type"`
}

func getInsightCounts(ctx context.Context, _ *mcp.CallToolRequest, input GetInsightCountsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/security-center/insights/" + input.Dimension

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

// RegisterTools registers Security Center tools with the MCP server.
// doGet performs a GET against the standard Cloudflare REST API and formats the
// response. It is shared by the newer Security Center read tools.
func doGet(ctx context.Context, url, apiToken string) (*mcp.CallToolResult, error) {
	cfResp, err := cfapi.DoRequest(ctx, http.MethodGet, url, apiToken, nil)
	if err != nil {
		return nil, err
	}
	if !cfResp.Success {
		return cfapi.APIErrorResult(cfResp.Errors), nil
	}
	return cfapi.FormatResult(cfResp)
}

// ListAccountInsightsInput holds parameters for listing account Security Center insights.
type ListAccountInsightsInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func listAccountInsights(ctx context.Context, _ *mcp.CallToolRequest, input ListAccountInsightsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/security-center/insights", apiToken)
	return result, nil, err
}

func RegisterTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_security_insights",
		Description: "List Security Center insights for a Cloudflare zone. Returns security issues with severity, type, and classification. Useful for identifying misconfigurations and vulnerabilities.",
	}, listInsights)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_insight_counts",
		Description: "Get aggregated Security Center insight counts by dimension (severity, class, or type). Useful for quick security posture overview and prioritization.",
	}, getInsightCounts)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_account_security_insights",
		Description: "List Security Center insights for a Cloudflare account (account scope). Returns security issues with severity, type, and classification.",
	}, listAccountInsights)
}
