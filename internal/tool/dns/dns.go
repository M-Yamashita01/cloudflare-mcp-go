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

// GetInput holds parameters for retrieving a single DNS record.
type GetInput struct {
	ZoneID   string `json:"zone_id"   jsonschema:"required,The ID of the zone"`
	RecordID string `json:"record_id" jsonschema:"required,The ID of the DNS record to retrieve"`
}

func get(ctx context.Context, _ *mcp.CallToolRequest, input GetInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_records/" + input.RecordID
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

// GetAnalyticsInput holds parameters for retrieving DNS analytics.
type GetAnalyticsInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
	Since  string `json:"since,omitempty" jsonschema:"Start date for the report in ISO 8601 format (e.g. 2026-05-01T00:00:00Z)"`
	Until  string `json:"until,omitempty" jsonschema:"End date for the report in ISO 8601 format"`
}

func getAnalytics(ctx context.Context, _ *mcp.CallToolRequest, input GetAnalyticsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_analytics/report"
	var params []string
	if input.Since != "" {
		params = append(params, fmt.Sprintf("since=%s", input.Since))
	}
	if input.Until != "" {
		params = append(params, fmt.Sprintf("until=%s", input.Until))
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

// GetAnalyticsByTimeInput holds parameters for retrieving DNS analytics grouped by time.
type GetAnalyticsByTimeInput struct {
	ZoneID    string `json:"zone_id" jsonschema:"required,The ID of the zone"`
	Since     string `json:"since,omitempty" jsonschema:"Start date for the report in ISO 8601 format (e.g. 2026-05-01T00:00:00Z)"`
	Until     string `json:"until,omitempty" jsonschema:"End date for the report in ISO 8601 format"`
	TimeDelta string `json:"time_delta,omitempty" jsonschema:"Unit of time to group data by: all, auto, year, quarter, month, week, day, hour, dekaminute, minute"`
}

func getAnalyticsByTime(ctx context.Context, _ *mcp.CallToolRequest, input GetAnalyticsByTimeInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_analytics/report/bytime"
	var params []string
	if input.Since != "" {
		params = append(params, fmt.Sprintf("since=%s", input.Since))
	}
	if input.Until != "" {
		params = append(params, fmt.Sprintf("until=%s", input.Until))
	}
	if input.TimeDelta != "" {
		params = append(params, fmt.Sprintf("time_delta=%s", input.TimeDelta))
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

// ExportInput holds parameters for exporting a zone's DNS records as a BIND file.
type ExportInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone whose DNS records to export"`
}

func export(ctx context.Context, _ *mcp.CallToolRequest, input ExportInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_records/export"
	body, status, err := cfapi.DoRawRequest(ctx, http.MethodGet, url, apiToken, nil)
	if err != nil {
		return nil, nil, err
	}
	if status != http.StatusOK {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Cloudflare API error: status %d: %s", status, string(body))}},
			IsError: true,
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: string(body)}},
	}, nil, nil
}

// GetUsageInput holds parameters for retrieving DNS record usage.
type GetUsageInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
}

func getUsage(ctx context.Context, _ *mcp.CallToolRequest, input GetUsageInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_records/usage"
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

// GetSettingsInput holds parameters for retrieving a zone's DNS settings.
type GetSettingsInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
}

func getSettings(ctx context.Context, _ *mcp.CallToolRequest, input GetSettingsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_settings"
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

// GetDNSSECInput holds parameters for retrieving DNSSEC details.
type GetDNSSECInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
}

func getDNSSEC(ctx context.Context, _ *mcp.CallToolRequest, input GetDNSSECInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dnssec"
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

// ListDNSSECZSKInput holds parameters for listing DNSSEC ZSKs.
type ListDNSSECZSKInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
}

func listDNSSECZSK(ctx context.Context, _ *mcp.CallToolRequest, input ListDNSSECZSKInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dnssec/zsk"
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

// ListScannedInput holds parameters for listing scanned DNS records.
type ListScannedInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
}

func listScanned(ctx context.Context, _ *mcp.CallToolRequest, input ListScannedInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_records/scan/review"
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

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_dns_record",
		Description: "Get a single DNS record in a Cloudflare zone by record ID. Returns the record's full detail (type, name, content, TTL, proxy status, comment).",
	}, get)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_dns_analytics",
		Description: "Get DNS query analytics report for a Cloudflare zone. Returns query counts, response codes, and query type distributions. Useful for detecting DNS anomalies and attack patterns.",
	}, getAnalytics)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_dns_analytics_bytime",
		Description: "Get DNS query analytics for a Cloudflare zone grouped by time interval. Returns time-series data points (grouped by time_delta) for spotting trends and spikes.",
	}, getAnalyticsByTime)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "export_dns_records",
		Description: "Export all DNS records for a Cloudflare zone as a BIND zone file. Returns the raw BIND-format text.",
	}, export)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_dns_record_usage",
		Description: "Get DNS record usage for a Cloudflare zone (counts toward plan limits). Returns current and allowed record counts.",
	}, getUsage)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_dns_settings",
		Description: "Get DNS settings for a Cloudflare zone (e.g. Foundation DNS, multi-provider, nameservers, zone mode). Returns the zone's DNS configuration.",
	}, getSettings)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_dnssec",
		Description: "Get DNSSEC details for a Cloudflare zone (status, DS record, digest, key tag, algorithm). Useful for verifying DNSSEC configuration.",
	}, getDNSSEC)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_dnssec_zsk",
		Description: "List DNSSEC Zone Signing Keys (ZSKs) for a Cloudflare zone (multi-signer DNSSEC). Returns the ZSK records.",
	}, listDNSSECZSK)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_scanned_dns_records",
		Description: "List DNS records discovered by a DNS record scan for a Cloudflare zone, pending review. Returns the scanned records.",
	}, listScanned)
}
