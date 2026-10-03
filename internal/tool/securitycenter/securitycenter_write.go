package securitycenter

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// sendWrite performs a write request against the standard Cloudflare REST API
// and formats the response. It is shared by the Security Center write tools.
func sendWrite(ctx context.Context, method, url, apiToken string, body io.Reader) (*mcp.CallToolResult, error) {
	cfResp, err := cfapi.DoRequest(ctx, method, url, apiToken, body)
	if err != nil {
		return nil, err
	}
	if !cfResp.Success {
		return cfapi.APIErrorResult(cfResp.Errors), nil
	}
	return cfapi.FormatResult(cfResp)
}

// invalidJSON returns an error result when s is not valid JSON, or nil otherwise.
func invalidJSON(field, s string) *mcp.CallToolResult {
	if json.Valid([]byte(s)) {
		return nil
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: "Error: " + field + " must be valid JSON"}},
		IsError: true,
	}
}

// UpdateAccountClassificationInput holds parameters for updating an account insight classification.
type UpdateAccountClassificationInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	IssueID   string `json:"issue_id"   jsonschema:"required,The ID of the Security Center issue/insight"`
	Config    string `json:"config"     jsonschema:"required,JSON object with the new classification (e.g. severity)"`
}

func updateAccountClassification(ctx context.Context, _ *mcp.CallToolRequest, input UpdateAccountClassificationInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/accounts/" + input.AccountID + "/security-center/insights/" + input.IssueID + "/classification"
	result, err := sendWrite(ctx, http.MethodPatch, url, apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// DismissAccountInsightInput holds parameters for dismissing an account insight.
type DismissAccountInsightInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	IssueID   string `json:"issue_id"   jsonschema:"required,The ID of the Security Center issue/insight"`
	Config    string `json:"config"     jsonschema:"required,JSON object with the dismiss state, e.g. {\"dismiss\":true}"`
}

func dismissAccountInsight(ctx context.Context, _ *mcp.CallToolRequest, input DismissAccountInsightInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/accounts/" + input.AccountID + "/security-center/insights/" + input.IssueID + "/dismiss"
	result, err := sendWrite(ctx, http.MethodPut, url, apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// StartAccountScanInput holds parameters for starting an on-demand account scan.
type StartAccountScanInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func startAccountScan(ctx context.Context, _ *mcp.CallToolRequest, input StartAccountScanInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/accounts/" + input.AccountID + "/security-center/insights/scans"
	result, err := sendWrite(ctx, http.MethodPost, url, apiToken, nil)
	return result, nil, err
}

// UpdatePartnerSettingsInput holds parameters for updating partner integration settings.
type UpdatePartnerSettingsInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Partner   string `json:"partner"    jsonschema:"required,The partner identifier"`
	Config    string `json:"config"     jsonschema:"required,JSON object with the partner integration settings"`
}

func updatePartnerSettings(ctx context.Context, _ *mcp.CallToolRequest, input UpdatePartnerSettingsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/accounts/" + input.AccountID + "/security-center/partners/" + input.Partner + "/settings"
	result, err := sendWrite(ctx, http.MethodPost, url, apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// UpdateAccountStateInput holds parameters for updating the account Security Center state.
type UpdateAccountStateInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Config    string `json:"config"     jsonschema:"required,JSON object with the new Security Center state"`
}

func updateAccountState(ctx context.Context, _ *mcp.CallToolRequest, input UpdateAccountStateInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/accounts/" + input.AccountID + "/security-center/state"
	result, err := sendWrite(ctx, http.MethodPut, url, apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// StartZoneScanInput holds parameters for starting an on-demand zone scan.
type StartZoneScanInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
}

func startZoneScan(ctx context.Context, _ *mcp.CallToolRequest, input StartZoneScanInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/security-center/insights/scans"
	result, err := sendWrite(ctx, http.MethodPost, url, apiToken, nil)
	return result, nil, err
}

// UpdateZoneClassificationInput holds parameters for updating a zone insight classification.
type UpdateZoneClassificationInput struct {
	ZoneID  string `json:"zone_id"  jsonschema:"required,The ID of the zone"`
	IssueID string `json:"issue_id" jsonschema:"required,The ID of the Security Center issue/insight"`
	Config  string `json:"config"   jsonschema:"required,JSON object with the new classification (e.g. severity)"`
}

func updateZoneClassification(ctx context.Context, _ *mcp.CallToolRequest, input UpdateZoneClassificationInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/security-center/insights/" + input.IssueID + "/classification"
	result, err := sendWrite(ctx, http.MethodPatch, url, apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// DismissZoneInsightInput holds parameters for dismissing a zone insight.
type DismissZoneInsightInput struct {
	ZoneID  string `json:"zone_id"  jsonschema:"required,The ID of the zone"`
	IssueID string `json:"issue_id" jsonschema:"required,The ID of the Security Center issue/insight"`
	Config  string `json:"config"   jsonschema:"required,JSON object with the dismiss state, e.g. {\"dismiss\":true}"`
}

func dismissZoneInsight(ctx context.Context, _ *mcp.CallToolRequest, input DismissZoneInsightInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/security-center/insights/" + input.IssueID + "/dismiss"
	result, err := sendWrite(ctx, http.MethodPut, url, apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// RegisterWriteTools registers Security Center write (mutation) tools with the MCP server.
//
// It is called only when write mode is enabled via CLOUDFLARE_MCP_ENABLE_WRITE.
func RegisterWriteTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_account_insight_classification",
		Description: "Update the classification (e.g. severity) of a Security Center insight in a Cloudflare account. The config argument is a JSON object.",
	}, updateAccountClassification)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "dismiss_account_insight",
		Description: "Dismiss (archive) or un-dismiss a Security Center insight in a Cloudflare account. The config argument is a JSON object with the dismiss state.",
	}, dismissAccountInsight)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "start_account_scan",
		Description: "Start an on-demand Security Center scan for a Cloudflare account.",
	}, startAccountScan)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_partner_integration_settings",
		Description: "Update the Security Center partner integration settings for a Cloudflare account and partner. The config argument is a JSON object.",
	}, updatePartnerSettings)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_account_security_center_state",
		Description: "Update the Security Center state for a Cloudflare account. The config argument is a JSON object.",
	}, updateAccountState)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "start_zone_scan",
		Description: "Start an on-demand Security Center scan for a zone.",
	}, startZoneScan)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_zone_insight_classification",
		Description: "Update the classification (e.g. severity) of a Security Center insight in a zone. The config argument is a JSON object.",
	}, updateZoneClassification)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "dismiss_zone_insight",
		Description: "Dismiss (archive) or un-dismiss a Security Center insight in a zone. The config argument is a JSON object with the dismiss state.",
	}, dismissZoneInsight)
}
