package dns

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// CreateInput holds parameters for creating a DNS record.
type CreateInput struct {
	ZoneID   string `json:"zone_id"            jsonschema:"required,The ID of the zone"`
	Type     string `json:"type"               jsonschema:"required,DNS record type (A, AAAA, CNAME, TXT, MX, NS, SRV, etc.)"`
	Name     string `json:"name"               jsonschema:"required,DNS record name (e.g. example.com or www.example.com)"`
	Content  string `json:"content"            jsonschema:"required,DNS record content (e.g. an IP address for A records)"`
	TTL      int    `json:"ttl,omitempty"      jsonschema:"Time to live in seconds; 1 means automatic (default: 1)"`
	Proxied  bool   `json:"proxied,omitempty"  jsonschema:"Whether the record is proxied through Cloudflare"`
	Priority int    `json:"priority,omitempty" jsonschema:"Record priority, required for MX and SRV records"`
	Comment  string `json:"comment,omitempty"  jsonschema:"Optional comment describing the record"`
}

func create(ctx context.Context, _ *mcp.CallToolRequest, input CreateInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	body := map[string]any{
		"type":    input.Type,
		"name":    input.Name,
		"content": input.Content,
	}
	if input.TTL > 0 {
		body["ttl"] = input.TTL
	}
	if input.Proxied {
		body["proxied"] = input.Proxied
	}
	if input.Priority > 0 {
		body["priority"] = input.Priority
	}
	if input.Comment != "" {
		body["comment"] = input.Comment
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return nil, nil, fmt.Errorf("marshaling request body: %w", err)
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_records"
	cfResp, err := cfapi.DoRequest(ctx, http.MethodPost, url, apiToken, bytes.NewReader(payload))
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

// UpdateInput holds parameters for updating an existing DNS record.
//
// Only the fields that are set are sent to the API (partial update via PATCH).
type UpdateInput struct {
	ZoneID   string `json:"zone_id"            jsonschema:"required,The ID of the zone"`
	RecordID string `json:"record_id"          jsonschema:"required,The ID of the DNS record to update"`
	Type     string `json:"type,omitempty"     jsonschema:"DNS record type (A, AAAA, CNAME, TXT, MX, NS, SRV, etc.)"`
	Name     string `json:"name,omitempty"     jsonschema:"DNS record name (e.g. example.com or www.example.com)"`
	Content  string `json:"content,omitempty"  jsonschema:"DNS record content (e.g. an IP address for A records)"`
	TTL      int    `json:"ttl,omitempty"      jsonschema:"Time to live in seconds; 1 means automatic"`
	Proxied  *bool  `json:"proxied,omitempty"  jsonschema:"Whether the record is proxied through Cloudflare"`
	Priority int    `json:"priority,omitempty" jsonschema:"Record priority, used for MX and SRV records"`
	Comment  string `json:"comment,omitempty"  jsonschema:"Optional comment describing the record"`
}

func update(ctx context.Context, _ *mcp.CallToolRequest, input UpdateInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	body := map[string]any{}
	if input.Type != "" {
		body["type"] = input.Type
	}
	if input.Name != "" {
		body["name"] = input.Name
	}
	if input.Content != "" {
		body["content"] = input.Content
	}
	if input.TTL > 0 {
		body["ttl"] = input.TTL
	}
	if input.Proxied != nil {
		body["proxied"] = *input.Proxied
	}
	if input.Priority > 0 {
		body["priority"] = input.Priority
	}
	if input.Comment != "" {
		body["comment"] = input.Comment
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return nil, nil, fmt.Errorf("marshaling request body: %w", err)
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_records/" + input.RecordID
	cfResp, err := cfapi.DoRequest(ctx, http.MethodPatch, url, apiToken, bytes.NewReader(payload))
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

// DeleteInput holds parameters for deleting a DNS record.
type DeleteInput struct {
	ZoneID   string `json:"zone_id"   jsonschema:"required,The ID of the zone"`
	RecordID string `json:"record_id" jsonschema:"required,The ID of the DNS record to delete"`
}

func deleteRecord(ctx context.Context, _ *mcp.CallToolRequest, input DeleteInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_records/" + input.RecordID
	cfResp, err := cfapi.DoRequest(ctx, http.MethodDelete, url, apiToken, nil)
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

// OverwriteInput holds parameters for overwriting (replacing) a DNS record.
//
// The Cloudflare overwrite endpoint uses PUT and replaces the whole record, so
// type, name, and content are required.
type OverwriteInput struct {
	ZoneID   string `json:"zone_id"            jsonschema:"required,The ID of the zone"`
	RecordID string `json:"record_id"          jsonschema:"required,The ID of the DNS record to overwrite"`
	Type     string `json:"type"               jsonschema:"required,DNS record type (A, AAAA, CNAME, TXT, MX, NS, SRV, etc.)"`
	Name     string `json:"name"               jsonschema:"required,DNS record name (e.g. example.com or www.example.com)"`
	Content  string `json:"content"            jsonschema:"required,DNS record content (e.g. an IP address for A records)"`
	TTL      int    `json:"ttl,omitempty"      jsonschema:"Time to live in seconds; 1 means automatic (default: 1)"`
	Proxied  bool   `json:"proxied,omitempty"  jsonschema:"Whether the record is proxied through Cloudflare"`
	Priority int    `json:"priority,omitempty" jsonschema:"Record priority, required for MX and SRV records"`
	Comment  string `json:"comment,omitempty"  jsonschema:"Optional comment describing the record"`
}

func overwrite(ctx context.Context, _ *mcp.CallToolRequest, input OverwriteInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	body := map[string]any{
		"type":    input.Type,
		"name":    input.Name,
		"content": input.Content,
	}
	if input.TTL > 0 {
		body["ttl"] = input.TTL
	}
	if input.Proxied {
		body["proxied"] = input.Proxied
	}
	if input.Priority > 0 {
		body["priority"] = input.Priority
	}
	if input.Comment != "" {
		body["comment"] = input.Comment
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return nil, nil, fmt.Errorf("marshaling request body: %w", err)
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_records/" + input.RecordID
	cfResp, err := cfapi.DoRequest(ctx, http.MethodPut, url, apiToken, bytes.NewReader(payload))
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

// BatchInput holds parameters for a batch DNS record operation.
//
// The Cloudflare batch endpoint accepts an object with optional "deletes",
// "patches", "posts", and "puts" arrays applied atomically. The caller provides
// that object as a JSON string, which is validated before being sent.
type BatchInput struct {
	ZoneID     string `json:"zone_id"    jsonschema:"required,The ID of the zone"`
	Operations string `json:"operations" jsonschema:"required,JSON object with optional deletes/patches/posts/puts arrays (applied atomically)"`
}

func batch(ctx context.Context, _ *mcp.CallToolRequest, input BatchInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	var obj map[string]json.RawMessage
	if err := json.Unmarshal([]byte(input.Operations), &obj); err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "Error: operations must be a JSON object: " + err.Error()}},
			IsError: true,
		}, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_records/batch"
	cfResp, err := cfapi.DoRequest(ctx, http.MethodPost, url, apiToken, bytes.NewReader([]byte(input.Operations)))
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

// ImportInput holds parameters for importing DNS records from a BIND file.
type ImportInput struct {
	ZoneID  string `json:"zone_id"           jsonschema:"required,The ID of the zone"`
	File    string `json:"file"              jsonschema:"required,BIND zone file contents to import"`
	Proxied string `json:"proxied,omitempty" jsonschema:"Optional proxied override expression (e.g. true to proxy all imported records)"`
}

func importRecords(ctx context.Context, _ *mcp.CallToolRequest, input ImportInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	fw, err := w.CreateFormFile("file", "import.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("building multipart form: %w", err)
	}
	if _, err := io.WriteString(fw, input.File); err != nil {
		return nil, nil, fmt.Errorf("writing import file: %w", err)
	}
	if input.Proxied != "" {
		if err := w.WriteField("proxied", input.Proxied); err != nil {
			return nil, nil, fmt.Errorf("writing proxied field: %w", err)
		}
	}
	if err := w.Close(); err != nil {
		return nil, nil, fmt.Errorf("closing multipart form: %w", err)
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_records/import"
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
	if err != nil {
		return nil, nil, fmt.Errorf("creating request: %w", err)
	}
	httpReq.Header.Set("Authorization", "Bearer "+apiToken)
	httpReq.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, nil, fmt.Errorf("calling Cloudflare API: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("reading response body: %w", err)
	}

	var cfResp cfapi.Response
	if err := json.Unmarshal(respBody, &cfResp); err != nil {
		return nil, nil, fmt.Errorf("parsing response: %w", err)
	}
	if !cfResp.Success {
		return cfapi.APIErrorResult(cfResp.Errors), nil, nil
	}

	result, err := cfapi.FormatResult(&cfResp)
	if err != nil {
		return nil, nil, err
	}
	return result, nil, nil
}

// ScanInput holds parameters for scanning a zone's DNS records.
type ScanInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
}

func scan(ctx context.Context, _ *mcp.CallToolRequest, input ScanInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_records/scan"
	cfResp, err := cfapi.DoRequest(ctx, http.MethodPost, url, apiToken, nil)
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

// ScanTriggerInput holds parameters for triggering a DNS record scan.
type ScanTriggerInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
}

func scanTrigger(ctx context.Context, _ *mcp.CallToolRequest, input ScanTriggerInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_records/scan/trigger"
	cfResp, err := cfapi.DoRequest(ctx, http.MethodPost, url, apiToken, nil)
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

// ReviewScannedInput holds parameters for reviewing scanned DNS records.
//
// The review body (which scanned records to accept/reject) is a rich structure,
// so it is passed as a JSON string and validated before being sent.
type ReviewScannedInput struct {
	ZoneID   string `json:"zone_id"  jsonschema:"required,The ID of the zone"`
	Decision string `json:"decision" jsonschema:"required,JSON body describing which scanned records to accept or reject"`
}

func reviewScanned(ctx context.Context, _ *mcp.CallToolRequest, input ReviewScannedInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	if !json.Valid([]byte(input.Decision)) {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "Error: decision must be valid JSON"}},
			IsError: true,
		}, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_records/scan/review"
	cfResp, err := cfapi.DoRequest(ctx, http.MethodPost, url, apiToken, bytes.NewReader([]byte(input.Decision)))
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

// UpdateSettingsInput holds parameters for updating a zone's DNS settings.
//
// DNS settings cover many optional fields, so they are passed as a JSON object
// string and validated before being sent.
type UpdateSettingsInput struct {
	ZoneID   string `json:"zone_id"  jsonschema:"required,The ID of the zone"`
	Settings string `json:"settings" jsonschema:"required,JSON object of DNS settings to update (e.g. foundation_dns, multi_provider, nameservers)"`
}

func updateSettings(ctx context.Context, _ *mcp.CallToolRequest, input UpdateSettingsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	var obj map[string]json.RawMessage
	if err := json.Unmarshal([]byte(input.Settings), &obj); err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "Error: settings must be a JSON object: " + err.Error()}},
			IsError: true,
		}, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dns_settings"
	cfResp, err := cfapi.DoRequest(ctx, http.MethodPatch, url, apiToken, bytes.NewReader([]byte(input.Settings)))
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

// EditDNSSECInput holds parameters for editing a zone's DNSSEC status.
type EditDNSSECInput struct {
	ZoneID      string `json:"zone_id"            jsonschema:"required,The ID of the zone"`
	Status      string `json:"status,omitempty"   jsonschema:"DNSSEC status: active or disabled"`
	MultiSigner *bool  `json:"dnssec_multi_signer,omitempty" jsonschema:"Enable multi-signer DNSSEC"`
	Presigned   *bool  `json:"dnssec_presigned,omitempty"    jsonschema:"Enable presigned DNSSEC"`
	UseNSEC3    *bool  `json:"dnssec_use_nsec3,omitempty"    jsonschema:"Use NSEC3 instead of NSEC"`
}

func editDNSSEC(ctx context.Context, _ *mcp.CallToolRequest, input EditDNSSECInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	body := map[string]any{}
	if input.Status != "" {
		body["status"] = input.Status
	}
	if input.MultiSigner != nil {
		body["dnssec_multi_signer"] = *input.MultiSigner
	}
	if input.Presigned != nil {
		body["dnssec_presigned"] = *input.Presigned
	}
	if input.UseNSEC3 != nil {
		body["dnssec_use_nsec3"] = *input.UseNSEC3
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return nil, nil, fmt.Errorf("marshaling request body: %w", err)
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/dnssec"
	cfResp, err := cfapi.DoRequest(ctx, http.MethodPatch, url, apiToken, bytes.NewReader(payload))
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

// RegisterWriteTools registers DNS write (mutation) tools with the MCP server.
//
// It is called only when write mode is enabled via CLOUDFLARE_MCP_ENABLE_WRITE,
// so these tools stay out of tools/list by default.
func RegisterWriteTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_dns_record",
		Description: "Create a DNS record in a Cloudflare zone. Requires type, name, and content. Returns the created record.",
	}, create)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_dns_record",
		Description: "Update an existing DNS record in a Cloudflare zone. Only the provided fields are changed (partial update). Returns the updated record.",
	}, update)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_dns_record",
		Description: "Delete a DNS record from a Cloudflare zone by record ID. Returns the ID of the deleted record.",
	}, deleteRecord)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "overwrite_dns_record",
		Description: "Overwrite (fully replace) an existing DNS record by ID via PUT. Requires type, name, and content; any field not provided reverts to its default.",
	}, overwrite)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "batch_dns_records",
		Description: "Apply a batch of DNS record changes atomically. The operations argument is a JSON object with optional deletes, patches, posts, and puts arrays.",
	}, batch)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "import_dns_records",
		Description: "Import DNS records into a Cloudflare zone from a BIND zone file. Provide the file contents; optionally set a proxied override expression.",
	}, importRecords)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "scan_dns_records",
		Description: "Scan a Cloudflare zone for DNS records at common names (for zones with a known provider). Returns the scan result.",
	}, scan)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "trigger_dns_record_scan",
		Description: "Trigger an asynchronous DNS record scan for a Cloudflare zone. Use list_scanned_dns_records to review results and review_scanned_dns_records to accept them.",
	}, scanTrigger)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "review_scanned_dns_records",
		Description: "Accept or reject DNS records found by a scan. The decision argument is a JSON body describing which scanned records to apply.",
	}, reviewScanned)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_dns_settings",
		Description: "Update DNS settings for a Cloudflare zone (e.g. Foundation DNS, multi-provider, nameservers). The settings argument is a JSON object of fields to change.",
	}, updateSettings)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "edit_dnssec",
		Description: "Edit DNSSEC status for a Cloudflare zone. Set status to active or disabled, and optionally toggle multi-signer, presigned, or NSEC3.",
	}, editDNSSEC)
}
