package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// --- list_zones ---

type ListZonesInput struct {
	Name    string `json:"name,omitempty"     jsonschema:"A domain name to filter zones by"`
	Page    int    `json:"page,omitempty"     jsonschema:"Page number of paginated results (default: 1)"`
	PerPage int    `json:"per_page,omitempty" jsonschema:"Number of zones per page (default: 20, max: 50)"`
}

func listZones(ctx context.Context, _ *mcp.CallToolRequest, input ListZonesInput) (*mcp.CallToolResult, any, error) {
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

// --- get_zone ---

type GetZoneInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone to retrieve"`
}

func getZone(ctx context.Context, _ *mcp.CallToolRequest, input GetZoneInput) (*mcp.CallToolResult, any, error) {
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

// --- list_dns_records ---

type ListDNSRecordsInput struct {
	ZoneID  string `json:"zone_id"            jsonschema:"required,The ID of the zone"`
	Type    string `json:"type,omitempty"     jsonschema:"DNS record type to filter by (A, AAAA, CNAME, TXT, MX, etc.)"`
	Name    string `json:"name,omitempty"     jsonschema:"DNS record name to filter by"`
	Content string `json:"content,omitempty"  jsonschema:"DNS record content to filter by"`
	Page    int    `json:"page,omitempty"     jsonschema:"Page number of paginated results (default: 1)"`
	PerPage int    `json:"per_page,omitempty" jsonschema:"Number of records per page (default: 100, max: 5000)"`
}

func listDNSRecords(ctx context.Context, _ *mcp.CallToolRequest, input ListDNSRecordsInput) (*mcp.CallToolResult, any, error) {
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

// --- list_accounts ---

type ListAccountsInput struct {
	Name    string `json:"name,omitempty"     jsonschema:"Account name to filter by"`
	Page    int    `json:"page,omitempty"     jsonschema:"Page number of paginated results (default: 1)"`
	PerPage int    `json:"per_page,omitempty" jsonschema:"Number of accounts per page (default: 20, max: 50)"`
}

func listAccounts(ctx context.Context, _ *mcp.CallToolRequest, input ListAccountsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/accounts"
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

// --- list_ip_access_rules ---

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

// --- list_waf_managed_rulesets ---

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

// --- query_security_events ---

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

// graphqlResponse is the top-level GraphQL response envelope.
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

	// Build the filter object conditionally.
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

	const query = `
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
        ruleId
        source
        userAgent
      }
    }
  }
}`

	reqBody := struct {
		Query     string    `json:"query"`
		Variables variables `json:"variables"`
	}{
		Query: query,
		Variables: variables{
			ZoneTag: input.ZoneID,
			Filter:  filter,
			Limit:   limit,
		},
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, nil, fmt.Errorf("marshaling GraphQL request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, cfapi.APIBase+"/graphql", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, nil, fmt.Errorf("creating request: %w", err)
	}
	httpReq.Header.Set("Authorization", "Bearer "+apiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, nil, fmt.Errorf("calling Cloudflare GraphQL API: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("reading response body: %w", err)
	}

	var gqlResp graphqlResponse
	if err := json.Unmarshal(respBody, &gqlResp); err != nil {
		return nil, nil, fmt.Errorf("parsing GraphQL response: %w", err)
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

// --- list_kv_namespaces ---

type ListKVNamespacesInput struct {
	AccountID string `json:"account_id"          jsonschema:"required,The ID of the Cloudflare account"`
	Page      int    `json:"page,omitempty"      jsonschema:"Page number of paginated results (default: 1)"`
	PerPage   int    `json:"per_page,omitempty"  jsonschema:"Number of namespaces per page (default: 20, max: 100)"`
	Order     string `json:"order,omitempty"     jsonschema:"Order results by field: id or title"`
	Direction string `json:"direction,omitempty" jsonschema:"Sort direction: asc or desc"`
}

func listKVNamespaces(ctx context.Context, _ *mcp.CallToolRequest, input ListKVNamespacesInput) (*mcp.CallToolResult, any, error) {
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

func main() {
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "cloudflare-mcp-go",
			Version: "0.1.0",
		},
		nil,
	)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_zones",
		Description: "List zones in your Cloudflare account. Returns zone details such as ID, name, status, and plan.",
	}, listZones)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_zone",
		Description: "Get details of a specific Cloudflare zone. Returns zone details such as ID, name, status, and plan.",
	}, getZone)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_accounts",
		Description: "List Cloudflare accounts accessible with the current API token. Returns account details such as ID, name, and settings.",
	}, listAccounts)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_kv_namespaces",
		Description: "List Workers KV namespaces in a Cloudflare account. Returns namespace details such as ID and title.",
	}, listKVNamespaces)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_dns_records",
		Description: "List DNS records for a Cloudflare zone. Returns record details such as ID, type, name, content, TTL, and proxy status.",
	}, listDNSRecords)

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
		Description: "Query WAF and firewall security events for a Cloudflare zone using the GraphQL Analytics API. Returns event details such as client IP, action, rule ID, request path, and source service.",
	}, querySecurityEvents)

	log.Println("Starting Cloudflare MCP server (stdio)...")
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}
