package security

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// CreateIPAccessRuleInput holds parameters for creating an IP access rule.
type CreateIPAccessRuleInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
	Mode   string `json:"mode"    jsonschema:"required,Action to apply: block, challenge, whitelist, js_challenge, managed_challenge"`
	Target string `json:"target"  jsonschema:"required,Target type: ip, ip_range, asn, or country"`
	Value  string `json:"value"   jsonschema:"required,Target value (e.g. an IP, CIDR, AS number like AS13335, or 2-letter country code)"`
	Notes  string `json:"notes,omitempty" jsonschema:"Optional note describing the rule"`
}

func createIPAccessRule(ctx context.Context, _ *mcp.CallToolRequest, input CreateIPAccessRuleInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	body := map[string]any{
		"mode": input.Mode,
		"configuration": map[string]string{
			"target": input.Target,
			"value":  input.Value,
		},
	}
	if input.Notes != "" {
		body["notes"] = input.Notes
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return nil, nil, fmt.Errorf("marshaling request body: %w", err)
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/firewall/access_rules/rules"
	return sendWrite(ctx, http.MethodPost, url, apiToken, bytes.NewReader(payload))
}

// UpdateIPAccessRuleInput holds parameters for updating an IP access rule.
type UpdateIPAccessRuleInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
	RuleID string `json:"rule_id" jsonschema:"required,The ID of the IP access rule to update"`
	Mode   string `json:"mode,omitempty"  jsonschema:"Action to apply: block, challenge, whitelist, js_challenge, managed_challenge"`
	Notes  string `json:"notes,omitempty" jsonschema:"Note describing the rule"`
}

func updateIPAccessRule(ctx context.Context, _ *mcp.CallToolRequest, input UpdateIPAccessRuleInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	body := map[string]any{}
	if input.Mode != "" {
		body["mode"] = input.Mode
	}
	if input.Notes != "" {
		body["notes"] = input.Notes
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return nil, nil, fmt.Errorf("marshaling request body: %w", err)
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/firewall/access_rules/rules/" + input.RuleID
	return sendWrite(ctx, http.MethodPatch, url, apiToken, bytes.NewReader(payload))
}

// DeleteIPAccessRuleInput holds parameters for deleting an IP access rule.
type DeleteIPAccessRuleInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
	RuleID string `json:"rule_id" jsonschema:"required,The ID of the IP access rule to delete"`
}

func deleteIPAccessRule(ctx context.Context, _ *mcp.CallToolRequest, input DeleteIPAccessRuleInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/firewall/access_rules/rules/" + input.RuleID
	return sendWrite(ctx, http.MethodDelete, url, apiToken, nil)
}

// sendWrite executes a write request and formats the Cloudflare response.
func sendWrite(ctx context.Context, method, url, apiToken string, body io.Reader) (*mcp.CallToolResult, any, error) {
	cfResp, err := cfapi.DoRequest(ctx, method, url, apiToken, body)
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

// RegisterWriteTools registers security write (mutation) tools with the MCP server.
//
// It is called only when write mode is enabled via CLOUDFLARE_MCP_ENABLE_WRITE.
func RegisterWriteTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_ip_access_rule",
		Description: "Create an IP access rule for a Cloudflare zone. Specify mode (block, challenge, whitelist, js_challenge, managed_challenge) and a target (ip, ip_range, asn, country) with its value.",
	}, createIPAccessRule)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_ip_access_rule",
		Description: "Update an existing IP access rule by ID. Can change the mode and/or notes.",
	}, updateIPAccessRule)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_ip_access_rule",
		Description: "Delete an IP access rule from a Cloudflare zone by rule ID.",
	}, deleteIPAccessRule)
}
