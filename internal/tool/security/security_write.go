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

// rateLimitBody builds the request body shared by create and update rate limit.
func rateLimitBody(urlPattern string, methods []string, threshold, period int, mode string, timeout int, description string) map[string]any {
	request := map[string]any{"url_pattern": urlPattern}
	if len(methods) > 0 {
		request["methods"] = methods
	}
	action := map[string]any{"mode": mode}
	if timeout > 0 {
		action["timeout"] = timeout
	}
	body := map[string]any{
		"match":     map[string]any{"request": request},
		"threshold": threshold,
		"period":    period,
		"action":    action,
	}
	if description != "" {
		body["description"] = description
	}
	return body
}

// CreateRateLimitInput holds parameters for creating a rate limit rule.
type CreateRateLimitInput struct {
	ZoneID      string   `json:"zone_id"           jsonschema:"required,The ID of the zone"`
	URLPattern  string   `json:"url_pattern"       jsonschema:"required,URL pattern to match (e.g. example.com/api/*)"`
	Threshold   int      `json:"threshold"         jsonschema:"required,Number of requests before the action triggers"`
	Period      int      `json:"period"            jsonschema:"required,Time window in seconds over which requests are counted"`
	Mode        string   `json:"mode"              jsonschema:"required,Action mode: simulate, ban, challenge, js_challenge, or managed_challenge"`
	Methods     []string `json:"methods,omitempty" jsonschema:"HTTP methods to match (e.g. GET, POST); omit to match all"`
	Timeout     int      `json:"timeout,omitempty" jsonschema:"Seconds the action lasts once triggered (required for ban and challenge modes)"`
	Description string   `json:"description,omitempty" jsonschema:"Optional description of the rule"`
}

func createRateLimit(ctx context.Context, _ *mcp.CallToolRequest, input CreateRateLimitInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	body := rateLimitBody(input.URLPattern, input.Methods, input.Threshold, input.Period, input.Mode, input.Timeout, input.Description)
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, nil, fmt.Errorf("marshaling request body: %w", err)
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/rate_limits"
	return sendWrite(ctx, http.MethodPost, url, apiToken, bytes.NewReader(payload))
}

// UpdateRateLimitInput holds parameters for updating a rate limit rule.
//
// The Cloudflare rate limit update endpoint replaces the whole rule (PUT), so
// all core fields are required.
type UpdateRateLimitInput struct {
	ZoneID      string   `json:"zone_id"       jsonschema:"required,The ID of the zone"`
	RateLimitID string   `json:"rate_limit_id" jsonschema:"required,The ID of the rate limit rule to update"`
	URLPattern  string   `json:"url_pattern"   jsonschema:"required,URL pattern to match (e.g. example.com/api/*)"`
	Threshold   int      `json:"threshold"     jsonschema:"required,Number of requests before the action triggers"`
	Period      int      `json:"period"        jsonschema:"required,Time window in seconds over which requests are counted"`
	Mode        string   `json:"mode"          jsonschema:"required,Action mode: simulate, ban, challenge, js_challenge, or managed_challenge"`
	Methods     []string `json:"methods,omitempty" jsonschema:"HTTP methods to match (e.g. GET, POST); omit to match all"`
	Timeout     int      `json:"timeout,omitempty" jsonschema:"Seconds the action lasts once triggered (required for ban and challenge modes)"`
	Description string   `json:"description,omitempty" jsonschema:"Optional description of the rule"`
}

func updateRateLimit(ctx context.Context, _ *mcp.CallToolRequest, input UpdateRateLimitInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	body := rateLimitBody(input.URLPattern, input.Methods, input.Threshold, input.Period, input.Mode, input.Timeout, input.Description)
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, nil, fmt.Errorf("marshaling request body: %w", err)
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/rate_limits/" + input.RateLimitID
	return sendWrite(ctx, http.MethodPut, url, apiToken, bytes.NewReader(payload))
}

// DeleteRateLimitInput holds parameters for deleting a rate limit rule.
type DeleteRateLimitInput struct {
	ZoneID      string `json:"zone_id"       jsonschema:"required,The ID of the zone"`
	RateLimitID string `json:"rate_limit_id" jsonschema:"required,The ID of the rate limit rule to delete"`
}

func deleteRateLimit(ctx context.Context, _ *mcp.CallToolRequest, input DeleteRateLimitInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/rate_limits/" + input.RateLimitID
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

	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_rate_limit",
		Description: "Create a rate limit rule for a Cloudflare zone. Specify a URL pattern, threshold, period (seconds), and action mode (simulate, ban, challenge, js_challenge, managed_challenge).",
	}, createRateLimit)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_rate_limit",
		Description: "Update an existing rate limit rule by ID. Replaces the rule, so all core fields (url_pattern, threshold, period, mode) are required.",
	}, updateRateLimit)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_rate_limit",
		Description: "Delete a rate limit rule from a Cloudflare zone by rule ID.",
	}, deleteRateLimit)
}
