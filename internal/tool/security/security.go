// Package security provides MCP tools for Cloudflare security and firewall management.
package security

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// ListIPAccessRulesInput holds query parameters for listing IP access rules.
type ListIPAccessRulesInput struct {
	ZoneID  string `json:"zone_id"            jsonschema:"required,The ID of the zone"`
	IP      string `json:"ip,omitempty"       jsonschema:"Filter by specific IP address"`
	Mode    string `json:"mode,omitempty"     jsonschema:"Filter by mode: block, challenge, whitelist, js_challenge"`
	Page    int    `json:"page,omitempty"     jsonschema:"Page number of paginated results (default: 1)"`
	PerPage int    `json:"per_page,omitempty" jsonschema:"Number of rules per page (default: 20, max: 1000)"`
}

func listIPAccessRules(ctx context.Context, _ *mcp.CallToolRequest, input ListIPAccessRulesInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/firewall/access_rules/rules"
	var params []string
	if input.IP != "" {
		params = append(params, fmt.Sprintf("configuration.value=%s", input.IP))
	}
	if input.Mode != "" {
		params = append(params, fmt.Sprintf("mode=%s", input.Mode))
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

// ListWAFManagedRulesetsInput holds parameters for listing WAF managed rulesets.
type ListWAFManagedRulesetsInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
}

func listWAFManagedRulesets(ctx context.Context, _ *mcp.CallToolRequest, input ListWAFManagedRulesetsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/rulesets/phases/http_request_firewall_managed/entrypoint"

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

// QuerySecurityEventsInput holds query parameters for querying security events via GraphQL.
type QuerySecurityEventsInput struct {
	ZoneID      string `json:"zone_id" jsonschema:"required,The ID of the zone"`
	DateFrom    string `json:"date_from" jsonschema:"required,Start datetime in RFC3339 format (e.g. 2026-03-23T08:19:58Z)"`
	DateTo      string `json:"date_to" jsonschema:"required,End datetime in RFC3339 format"`
	Source      string `json:"source,omitempty" jsonschema:"Filter by mitigation source: firewallManaged, firewallCustom, firewallrules, waf, rateLimit, bic, hot, securitylevel, uablock, ip, iprange, asn, country, zonelockdown, l7ddos, botfight, botmanagement, apishield, apishieldschemavalidation, apishieldtokenvalidation, apishieldsequencemitigation, dlp, validation"`
	Action      string `json:"action,omitempty" jsonschema:"Filter by action taken: block, challenge, jschallenge, managedchallenge, log, allow, bypass, connectionclose"`
	ClientIP    string `json:"client_ip,omitempty" jsonschema:"Filter by client IP address"`
	Host        string `json:"host,omitempty" jsonschema:"Filter by requested HTTP hostname"`
	RuleID      string `json:"rule_id,omitempty" jsonschema:"Filter by the rule ID that triggered the event"`
	Country     string `json:"country,omitempty" jsonschema:"Filter by client country name (e.g. US, JP, CN)"`
	HTTPMethod  string `json:"http_method,omitempty" jsonschema:"Filter by HTTP request method (e.g. GET, POST)"`
	RequestPath string `json:"request_path,omitempty" jsonschema:"Filter by HTTP request path"`
	Limit       int    `json:"limit,omitempty" jsonschema:"Max number of events to return (default: 100, max: 10000)"`
}

type graphqlResponse struct {
	Data   json.RawMessage `json:"data"`
	Errors []graphqlError  `json:"errors"`
}

type graphqlError struct {
	Message string `json:"message"`
}

func querySecurityEvents(ctx context.Context, _ *mcp.CallToolRequest, input QuerySecurityEventsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	limit := input.Limit
	if limit <= 0 {
		limit = 100
	}

	type filterObj struct {
		DatetimeGeq                 string `json:"datetime_geq"`
		DatetimeLt                  string `json:"datetime_lt"`
		Source                      string `json:"source,omitempty"`
		Action                      string `json:"action,omitempty"`
		ClientIP                    string `json:"clientIP,omitempty"`
		ClientRequestHTTPHost       string `json:"clientRequestHTTPHost,omitempty"`
		RuleID                      string `json:"ruleId,omitempty"`
		ClientCountryName           string `json:"clientCountryName,omitempty"`
		ClientRequestHTTPMethodName string `json:"clientRequestHTTPMethodName,omitempty"`
		ClientRequestPath           string `json:"clientRequestPath,omitempty"`
	}
	filter := filterObj{
		DatetimeGeq:                 input.DateFrom,
		DatetimeLt:                  input.DateTo,
		Source:                      input.Source,
		Action:                      input.Action,
		ClientIP:                    input.ClientIP,
		ClientRequestHTTPHost:       input.Host,
		RuleID:                      input.RuleID,
		ClientCountryName:           input.Country,
		ClientRequestHTTPMethodName: input.HTTPMethod,
		ClientRequestPath:           input.RequestPath,
	}

	type variables struct {
		ZoneTag string    `json:"zoneTag"`
		Filter  filterObj `json:"filter"`
		Limit   int       `json:"limit"`
	}

	const queryWithBotScore = `
query SecurityEvents($zoneTag: String!, $filter: FirewallEventsAdaptiveFilter_InputObject, $limit: Int) {
  viewer {
    zones(filter: { zoneTag: $zoneTag }) {
      firewallEventsAdaptive(
        filter: $filter
        limit: $limit
        orderBy: [datetime_DESC]
      ) {
        action
        botScore
        botScoreSrcName
        clientAsn
        clientCountryName
        clientIP
        clientRequestHTTPHost
        clientRequestHTTPMethodName
        clientRequestPath
        clientRequestQuery
        datetime
        description
        edgeResponseStatus
        originResponseStatus
        ruleId
        source
        userAgent
      }
    }
  }
}`

	const queryWithoutBotScore = `
query SecurityEvents($zoneTag: String!, $filter: FirewallEventsAdaptiveFilter_InputObject, $limit: Int) {
  viewer {
    zones(filter: { zoneTag: $zoneTag }) {
      firewallEventsAdaptive(
        filter: $filter
        limit: $limit
        orderBy: [datetime_DESC]
      ) {
        action
        clientAsn
        clientCountryName
        clientIP
        clientRequestHTTPHost
        clientRequestHTTPMethodName
        clientRequestPath
        clientRequestQuery
        datetime
        description
        edgeResponseStatus
        originResponseStatus
        ruleId
        source
        userAgent
      }
    }
  }
}`

	vars := variables{
		ZoneTag: input.ZoneID,
		Filter:  filter,
		Limit:   limit,
	}

	// Try with bot score fields first; fall back without them for zones without Bot Management.
	gqlResp, err := executeGraphQL(ctx, apiToken, queryWithBotScore, vars)
	if err != nil {
		return nil, nil, err
	}

	if hasBotScoreError(gqlResp) {
		gqlResp, err = executeGraphQL(ctx, apiToken, queryWithoutBotScore, vars)
		if err != nil {
			return nil, nil, err
		}
	}

	if len(gqlResp.Errors) > 0 {
		var msgs []string
		for _, e := range gqlResp.Errors {
			msgs = append(msgs, e.Message)
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "GraphQL error: " + strings.Join(msgs, "; ")}},
			IsError: true,
		}, nil, nil
	}

	formatted, err := json.MarshalIndent(gqlResp.Data, "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("formatting result: %w", err)
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: string(formatted)}},
	}, nil, nil
}

func executeGraphQL(ctx context.Context, apiToken string, query string, vars any) (*graphqlResponse, error) {
	reqBody := struct {
		Query     string `json:"query"`
		Variables any    `json:"variables"`
	}{
		Query:     query,
		Variables: vars,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshaling GraphQL request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, cfapi.APIBase+"/graphql", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	httpReq.Header.Set("Authorization", "Bearer "+apiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("calling Cloudflare GraphQL API: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	var gqlResp graphqlResponse
	if err := json.Unmarshal(respBody, &gqlResp); err != nil {
		return nil, fmt.Errorf("parsing GraphQL response: %w", err)
	}
	return &gqlResp, nil
}

func hasBotScoreError(resp *graphqlResponse) bool {
	for _, e := range resp.Errors {
		if strings.Contains(e.Message, "botscoresrcname") || strings.Contains(e.Message, "botscore") {
			return true
		}
	}
	return false
}

// RegisterTools registers security and firewall tools with the MCP server.
func RegisterTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_waf_managed_rulesets",
		Description: "Get the WAF managed rulesets entrypoint for a Cloudflare zone. Returns which managed rulesets (e.g. Cloudflare Managed Ruleset, OWASP) are enabled and their configuration.",
	}, listWAFManagedRulesets)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_ip_access_rules",
		Description: "List IP access rules for a Cloudflare zone. Returns rules that block, challenge, or allow specific IPs, CIDRs, ASNs, or countries.",
	}, listIPAccessRules)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "query_security_events",
		Description: "Query security events for a Cloudflare zone using the GraphQL Analytics API. Returns event details including client info, action taken, rule details, bot score, and response status for security event triage.",
	}, querySecurityEvents)
}
