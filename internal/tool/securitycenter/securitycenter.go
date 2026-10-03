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

// GetAccountInsightCountsInput holds parameters for account insight counts by dimension.
type GetAccountInsightCountsInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Dimension string `json:"dimension"  jsonschema:"required,The dimension to group by: class, severity, or type"`
}

func getAccountInsightCounts(ctx context.Context, _ *mcp.CallToolRequest, input GetAccountInsightCountsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/security-center/insights/"+input.Dimension, apiToken)
	return result, nil, err
}

// GetAccountInsightsAuditLogInput holds parameters for the account insights audit log.
type GetAccountInsightsAuditLogInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func getAccountInsightsAuditLog(ctx context.Context, _ *mcp.CallToolRequest, input GetAccountInsightsAuditLogInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/security-center/insights/audit-log", apiToken)
	return result, nil, err
}

// GetAccountPartnerCountInput holds parameters for the account partner insight count.
type GetAccountPartnerCountInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func getAccountPartnerCount(ctx context.Context, _ *mcp.CallToolRequest, input GetAccountPartnerCountInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/security-center/insights/partner-count", apiToken)
	return result, nil, err
}

// GetAccountScansInput holds parameters for recent account scans.
type GetAccountScansInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func getAccountScans(ctx context.Context, _ *mcp.CallToolRequest, input GetAccountScansInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/security-center/insights/scans", apiToken)
	return result, nil, err
}

// GetAccountIssueAuditLogInput holds parameters for an account issue audit log.
type GetAccountIssueAuditLogInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	IssueID   string `json:"issue_id"   jsonschema:"required,The ID of the Security Center issue/insight"`
}

func getAccountIssueAuditLog(ctx context.Context, _ *mcp.CallToolRequest, input GetAccountIssueAuditLogInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/security-center/insights/"+input.IssueID+"/audit-log", apiToken)
	return result, nil, err
}

// GetAccountInsightContextInput holds parameters for an account insight context.
type GetAccountInsightContextInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	IssueID   string `json:"issue_id"   jsonschema:"required,The ID of the Security Center issue/insight"`
}

func getAccountInsightContext(ctx context.Context, _ *mcp.CallToolRequest, input GetAccountInsightContextInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/security-center/insights/"+input.IssueID+"/context", apiToken)
	return result, nil, err
}

// GetPartnerSettingsInput holds parameters for partner integration settings.
type GetPartnerSettingsInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Partner   string `json:"partner"    jsonschema:"required,The partner identifier"`
}

func getPartnerSettings(ctx context.Context, _ *mcp.CallToolRequest, input GetPartnerSettingsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/security-center/partners/"+input.Partner+"/settings", apiToken)
	return result, nil, err
}

// ListShadowZonesInput holds parameters for listing partner-discovered shadow zones.
type ListShadowZonesInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Partner   string `json:"partner"    jsonschema:"required,The partner identifier"`
}

func listShadowZones(ctx context.Context, _ *mcp.CallToolRequest, input ListShadowZonesInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/security-center/partners/"+input.Partner+"/shadow-zones", apiToken)
	return result, nil, err
}

// ListShadowZoneHostsInput holds parameters for listing hosts in a shadow zone.
type ListShadowZoneHostsInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Partner   string `json:"partner"    jsonschema:"required,The partner identifier"`
	Domain    string `json:"domain"     jsonschema:"required,The shadow zone domain"`
}

func listShadowZoneHosts(ctx context.Context, _ *mcp.CallToolRequest, input ListShadowZoneHostsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/security-center/partners/"+input.Partner+"/shadow-zones/"+input.Domain+"/hosts", apiToken)
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

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_account_insight_counts",
		Description: "Get aggregated Security Center insight counts for a Cloudflare account by dimension (class, severity, or type).",
	}, getAccountInsightCounts)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_account_insights_audit_log",
		Description: "Get the Security Center insights audit log for a Cloudflare account (changes to insights over time).",
	}, getAccountInsightsAuditLog)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_account_partner_insight_count",
		Description: "Get the count of partner-provided Security Center insights for a Cloudflare account.",
	}, getAccountPartnerCount)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_account_recent_scans",
		Description: "Get the recent Security Center scans for a Cloudflare account.",
	}, getAccountScans)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_account_issue_audit_log",
		Description: "Get the audit log for a specific Security Center insight (issue) in a Cloudflare account by issue ID.",
	}, getAccountIssueAuditLog)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_account_insight_context",
		Description: "Get the context (affected resources and details) for a specific Security Center insight in a Cloudflare account by issue ID.",
	}, getAccountInsightContext)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_partner_integration_settings",
		Description: "Get the Security Center partner integration settings for a Cloudflare account and partner.",
	}, getPartnerSettings)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_partner_shadow_zones",
		Description: "List partner-discovered shadow zones (unmanaged domains) for a Cloudflare account and partner.",
	}, listShadowZones)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_partner_shadow_zone_hosts",
		Description: "List the hosts in a partner-discovered shadow zone for a Cloudflare account, partner, and domain.",
	}, listShadowZoneHosts)
}
