// Package intel provides MCP tools for Cloudflare threat intelligence.
package intel

import (
	"context"
	"fmt"
	"net/http"
	"os"

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

// RegisterTools registers threat intelligence tools with the MCP server.
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
}
