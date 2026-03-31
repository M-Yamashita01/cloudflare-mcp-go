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
)

const cloudflareAPIBase = "https://api.cloudflare.com/client/v4"

// CloudflareResponse is the standard Cloudflare API v4 response envelope.
type CloudflareResponse struct {
	Success  bool            `json:"success"`
	Errors   []CloudflareErr `json:"errors"`
	Messages []any           `json:"messages"`
	Result   json.RawMessage `json:"result"`
}

// CloudflareErr holds error details returned by the Cloudflare API.
type CloudflareErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// checkToken returns an error CallToolResult when apiToken is empty, or nil if it is set.
func checkToken(apiToken string) *mcp.CallToolResult {
	if apiToken == "" {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "Error: CLOUDFLARE_API_TOKEN environment variable is not set"}},
			IsError: true,
		}
	}
	return nil
}

// apiErrorResult converts Cloudflare API errors into an error CallToolResult.
func apiErrorResult(errs []CloudflareErr) *mcp.CallToolResult {
	var msgs []string
	for _, e := range errs {
		msgs = append(msgs, fmt.Sprintf("[%d] %s", e.Code, e.Message))
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: "Cloudflare API error: " + strings.Join(msgs, "; ")}},
		IsError: true,
	}
}

// doCloudflareRequest executes an HTTP request against the Cloudflare API and returns the parsed response.
func doCloudflareRequest(ctx context.Context, method, url, apiToken string, body io.Reader) (*CloudflareResponse, error) {
	httpReq, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	httpReq.Header.Set("Authorization", "Bearer "+apiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("calling Cloudflare API: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	var cfResp CloudflareResponse
	if err := json.Unmarshal(respBody, &cfResp); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}
	return &cfResp, nil
}

// formatResult pretty-prints the result field of a CloudflareResponse as a CallToolResult.
func formatResult(cfResp *CloudflareResponse) (*mcp.CallToolResult, error) {
	formatted, err := json.MarshalIndent(cfResp.Result, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("formatting result: %w", err)
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: string(formatted)}},
	}, nil
}

// --- list_zones ---

type ListZonesInput struct {
	Name    string `json:"name,omitempty"     jsonschema:"A domain name to filter zones by"`
	Page    int    `json:"page,omitempty"     jsonschema:"Page number of paginated results (default: 1)"`
	PerPage int    `json:"per_page,omitempty" jsonschema:"Number of zones per page (default: 20, max: 50)"`
}

func listZones(ctx context.Context, _ *mcp.CallToolRequest, input ListZonesInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := checkToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cloudflareAPIBase + "/zones"
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

	cfResp, err := doCloudflareRequest(ctx, http.MethodGet, url, apiToken, nil)
	if err != nil {
		return nil, nil, err
	}
	if !cfResp.Success {
		return apiErrorResult(cfResp.Errors), nil, nil
	}

	result, err := formatResult(cfResp)
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
	if result := checkToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cloudflareAPIBase + "/zones/" + input.ZoneID

	cfResp, err := doCloudflareRequest(ctx, http.MethodGet, url, apiToken, nil)
	if err != nil {
		return nil, nil, err
	}
	if !cfResp.Success {
		return apiErrorResult(cfResp.Errors), nil, nil
	}

	result, err := formatResult(cfResp)
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
	if result := checkToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cloudflareAPIBase + "/zones/" + input.ZoneID + "/dns_records"
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

	cfResp, err := doCloudflareRequest(ctx, http.MethodGet, url, apiToken, nil)
	if err != nil {
		return nil, nil, err
	}
	if !cfResp.Success {
		return apiErrorResult(cfResp.Errors), nil, nil
	}

	result, err := formatResult(cfResp)
	if err != nil {
		return nil, nil, err
	}
	return result, nil, nil
}

// --- update_dns_record ---

type UpdateDNSRecordInput struct {
	ZoneID      string `json:"zone_id"            jsonschema:"required,The ID of the zone"`
	DNSRecordID string `json:"dns_record_id"      jsonschema:"required,The ID of the DNS record to update"`
	Type        string `json:"type,omitempty"     jsonschema:"DNS record type (A, AAAA, CNAME, TXT, MX, NS, etc.)"`
	Name        string `json:"name,omitempty"     jsonschema:"DNS record name"`
	Content     string `json:"content,omitempty"  jsonschema:"DNS record content"`
	TTL         int    `json:"ttl,omitempty"      jsonschema:"Time to live in seconds (1=automatic, 60-86400)"`
	Proxied     *bool  `json:"proxied,omitempty"  jsonschema:"Whether the record is proxied through Cloudflare"`
	Comment     string `json:"comment,omitempty"  jsonschema:"Comment for the DNS record"`
}

type updateDNSRecordRequest struct {
	Type    string `json:"type,omitempty"`
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
	TTL     int    `json:"ttl,omitempty"`
	Proxied *bool  `json:"proxied,omitempty"`
	Comment string `json:"comment,omitempty"`
}

func updateDNSRecord(ctx context.Context, _ *mcp.CallToolRequest, input UpdateDNSRecordInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := checkToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cloudflareAPIBase + "/zones/" + input.ZoneID + "/dns_records/" + input.DNSRecordID

	reqBody := updateDNSRecordRequest{
		Type:    input.Type,
		Name:    input.Name,
		Content: input.Content,
		TTL:     input.TTL,
		Proxied: input.Proxied,
		Comment: input.Comment,
	}
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, nil, fmt.Errorf("marshaling request body: %w", err)
	}

	cfResp, err := doCloudflareRequest(ctx, http.MethodPatch, url, apiToken, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, nil, err
	}
	if !cfResp.Success {
		return apiErrorResult(cfResp.Errors), nil, nil
	}

	result, err := formatResult(cfResp)
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
		Name:        "update_dns_record",
		Description: "Update an existing DNS record in a Cloudflare zone. Returns the updated record details such as ID, type, name, content, TTL, and proxy status.",
	}, updateDNSRecord)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_dns_records",
		Description: "List DNS records for a Cloudflare zone. Returns record details such as ID, type, name, content, TTL, and proxy status.",
	}, listDNSRecords)

	log.Println("Starting Cloudflare MCP server (stdio)...")
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}
