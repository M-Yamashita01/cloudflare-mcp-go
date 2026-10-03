// Package intel provides MCP tools for Cloudflare threat intelligence.
package intel

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// GetIPIntelInput holds parameters for retrieving IP threat intelligence.
type GetIPIntelInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	IP        string `json:"ip" jsonschema:"required,The IPv4 or IPv6 address to look up"`
}

func getIPIntel(ctx context.Context, _ *mcp.CallToolRequest, input GetIPIntelInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := fmt.Sprintf("%s/accounts/%s/intel/ip?ipv4=%s", cfapi.APIBase, input.AccountID, input.IP)

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

// GetDomainIntelInput holds parameters for retrieving domain threat intelligence.
type GetDomainIntelInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Domain    string `json:"domain" jsonschema:"required,The domain name to look up"`
}

func getDomainIntel(ctx context.Context, _ *mcp.CallToolRequest, input GetDomainIntelInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := fmt.Sprintf("%s/accounts/%s/intel/domain?domain=%s", cfapi.APIBase, input.AccountID, input.Domain)

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

// GetDomainHistoryInput holds parameters for retrieving domain threat history.
type GetDomainHistoryInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Domain    string `json:"domain" jsonschema:"required,The domain name to look up"`
}

func getDomainHistory(ctx context.Context, _ *mcp.CallToolRequest, input GetDomainHistoryInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := fmt.Sprintf("%s/accounts/%s/intel/domain-history?domain=%s", cfapi.APIBase, input.AccountID, input.Domain)

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

// ListPassiveDNSInput holds parameters for listing domains resolved to an IP.
type ListPassiveDNSInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	IP        string `json:"ip" jsonschema:"required,The IPv4 address to look up"`
}

func listPassiveDNS(ctx context.Context, _ *mcp.CallToolRequest, input ListPassiveDNSInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := fmt.Sprintf("%s/accounts/%s/intel/dns?ipv4=%s", cfapi.APIBase, input.AccountID, input.IP)

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

// GetWhoisInput holds parameters for retrieving WHOIS data for a domain.
type GetWhoisInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Domain    string `json:"domain" jsonschema:"required,The domain name to look up"`
}

func getWhois(ctx context.Context, _ *mcp.CallToolRequest, input GetWhoisInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := fmt.Sprintf("%s/accounts/%s/intel/whois?domain=%s", cfapi.APIBase, input.AccountID, input.Domain)

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

// GetASNIntelInput holds parameters for retrieving ASN intelligence.
type GetASNIntelInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	ASN       int    `json:"asn" jsonschema:"required,The Autonomous System Number to look up"`
}

func getASNIntel(ctx context.Context, _ *mcp.CallToolRequest, input GetASNIntelInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := fmt.Sprintf("%s/accounts/%s/intel/asn/%d", cfapi.APIBase, input.AccountID, input.ASN)

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

// GetDomainIntelBulkInput holds parameters for bulk domain intelligence lookup.
type GetDomainIntelBulkInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Domains   string `json:"domains" jsonschema:"required,Comma-separated list of domain names to look up"`
}

func getDomainIntelBulk(ctx context.Context, _ *mcp.CallToolRequest, input GetDomainIntelBulkInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := fmt.Sprintf("%s/accounts/%s/intel/domain/bulk?domain=%s", cfapi.APIBase, input.AccountID, input.Domains)

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

// RegisterTools registers threat intelligence tools with the MCP server.
// doGet performs a GET against the standard Cloudflare REST API and formats the
// response. It is shared by the newer intel read tools.
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

// ListASRIssueTypesInput holds parameters for listing attack-surface issue types.
type ListASRIssueTypesInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func listASRIssueTypes(ctx context.Context, _ *mcp.CallToolRequest, input ListASRIssueTypesInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/intel/attack-surface-report/issue-types", apiToken)
	return result, nil, err
}

// ListASRIssuesInput holds parameters for listing attack-surface issues.
type ListASRIssuesInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func listASRIssues(ctx context.Context, _ *mcp.CallToolRequest, input ListASRIssuesInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/intel/attack-surface-report/issues", apiToken)
	return result, nil, err
}

// GetASRClassCountsInput holds parameters for attack-surface issue counts by class.
type GetASRClassCountsInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func getASRClassCounts(ctx context.Context, _ *mcp.CallToolRequest, input GetASRClassCountsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/intel/attack-surface-report/issues/class", apiToken)
	return result, nil, err
}

// GetASRSeverityCountsInput holds parameters for attack-surface issue counts by severity.
type GetASRSeverityCountsInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func getASRSeverityCounts(ctx context.Context, _ *mcp.CallToolRequest, input GetASRSeverityCountsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/intel/attack-surface-report/issues/severity", apiToken)
	return result, nil, err
}

// GetASRTypeCountsInput holds parameters for attack-surface issue counts by type.
type GetASRTypeCountsInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func getASRTypeCounts(ctx context.Context, _ *mcp.CallToolRequest, input GetASRTypeCountsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/intel/attack-surface-report/issues/type", apiToken)
	return result, nil, err
}

// GetIPListsInput holds parameters for listing available IP lists.
type GetIPListsInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func getIPLists(ctx context.Context, _ *mcp.CallToolRequest, input GetIPListsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/intel/ip-lists", apiToken)
	return result, nil, err
}

// GetBulkDomainDetailsInput holds parameters for getting multiple domain details.
type GetBulkDomainDetailsInput struct {
	AccountID string   `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Domains   []string `json:"domains"    jsonschema:"required,List of domain names to look up"`
}

func getBulkDomainDetails(ctx context.Context, _ *mcp.CallToolRequest, input GetBulkDomainDetailsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	var params []string
	for _, d := range input.Domains {
		params = append(params, "domain="+url.QueryEscape(d))
	}
	reqURL := cfapi.APIBase + "/accounts/" + input.AccountID + "/intel/domain/bulk"
	if len(params) > 0 {
		reqURL += "?" + strings.Join(params, "&")
	}

	result, err := doGet(ctx, reqURL, apiToken)
	return result, nil, err
}

// GetURLIntelInput holds parameters for getting URL threat intelligence.
type GetURLIntelInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	URL       string `json:"url"        jsonschema:"required,The URL to look up"`
}

func getURLIntel(ctx context.Context, _ *mcp.CallToolRequest, input GetURLIntelInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	reqURL := cfapi.APIBase + "/accounts/" + input.AccountID + "/intel/url?url=" + url.QueryEscape(input.URL)
	result, err := doGet(ctx, reqURL, apiToken)
	return result, nil, err
}

// ListIndicatorFeedsInput holds parameters for listing indicator feeds.
type ListIndicatorFeedsInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func listIndicatorFeeds(ctx context.Context, _ *mcp.CallToolRequest, input ListIndicatorFeedsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/intel/indicator-feeds", apiToken)
	return result, nil, err
}

func RegisterTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_ip_intel",
		Description: "Get threat intelligence for an IP address. Returns geolocation, ASN, infrastructure type, and security threat categories. Useful for investigating suspicious IPs found in security events or access logs.",
	}, getIPIntel)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_domain_intel",
		Description: "Get security intelligence for a domain. Returns risk scores, content categories, and DNS information. Useful for investigating suspicious domains found in referrer headers or access logs.",
	}, getDomainIntel)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_domain_history",
		Description: "Get historical threat data for a domain. Returns past and current security threat categories and content classifications. Useful for checking if a domain has a pattern of malicious behavior over time.",
	}, getDomainHistory)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_passive_dns",
		Description: "List domains that have resolved to a specific IP address (passive DNS). Useful for identifying shared hosting or malicious infrastructure by revealing which domains point to a given IP.",
	}, listPassiveDNS)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_whois",
		Description: "Get WHOIS registration data for a domain. Returns registrant information, nameservers, and registration/expiration dates. Useful for investigating domain ownership and detecting newly registered suspicious domains.",
	}, getWhois)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_asn_intel",
		Description: "Get an overview of an Autonomous System Number (ASN) and its subnet allocations. Useful for understanding the network behind a suspicious IP and assessing whether an entire ASN is involved in attacks.",
	}, getASNIntel)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_domain_intel_bulk",
		Description: "Get threat intelligence for multiple domains at once. Returns risk scores and content categories for each domain. Useful for batch assessment of suspicious domains found in logs.",
	}, getDomainIntelBulk)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_attack_surface_issue_types",
		Description: "List the Security Center attack-surface-report issue types for a Cloudflare account.",
	}, listASRIssueTypes)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_attack_surface_issues",
		Description: "List Security Center attack-surface-report issues for a Cloudflare account (misconfigurations and exposures).",
	}, listASRIssues)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_attack_surface_issue_counts_by_class",
		Description: "Get Security Center attack-surface issue counts grouped by class for a Cloudflare account.",
	}, getASRClassCounts)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_attack_surface_issue_counts_by_severity",
		Description: "Get Security Center attack-surface issue counts grouped by severity for a Cloudflare account.",
	}, getASRSeverityCounts)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_attack_surface_issue_counts_by_type",
		Description: "Get Security Center attack-surface issue counts grouped by type for a Cloudflare account.",
	}, getASRTypeCounts)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_ip_lists",
		Description: "Get the available Cloudflare threat-intelligence IP lists for a Cloudflare account.",
	}, getIPLists)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_bulk_domain_details",
		Description: "Get threat intelligence details for multiple domains at once (GET bulk). Returns per-domain risk and category data.",
	}, getBulkDomainDetails)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_url_intel",
		Description: "Get threat intelligence for a URL. Returns risk and category information for the URL.",
	}, getURLIntel)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_indicator_feeds",
		Description: "List the threat-intelligence indicator feeds owned by a Cloudflare account.",
	}, listIndicatorFeeds)
}
