// Package logs provides MCP tools for Cloudflare HTTP log investigation.
package logs

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// doLogpullRequest executes a request against the Cloudflare Logpull API.
// Unlike the standard REST API, Logpull returns NDJSON (one JSON object per line)
// rather than the standard Cloudflare response envelope.
func doLogpullRequest(ctx context.Context, url, apiToken string) (string, error) {
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("creating request: %w", err)
	}
	httpReq.Header.Set("Authorization", "Bearer "+apiToken)

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("calling Cloudflare Logpull API: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("logpull API returned status %d: %s", resp.StatusCode, string(body))
	}

	return string(body), nil
}

// GetByRayIDInput holds parameters for retrieving a log entry by Ray ID.
type GetByRayIDInput struct {
	ZoneID     string `json:"zone_id" jsonschema:"required,The ID of the zone"`
	RayID      string `json:"ray_id" jsonschema:"required,The Ray ID of the request to look up"`
	Fields     string `json:"fields,omitempty" jsonschema:"Comma-separated list of log fields to return"`
	Timestamps string `json:"timestamps,omitempty" jsonschema:"Timestamp format: unixnano (default), unix, or rfc3339"`
}

func getByRayID(ctx context.Context, _ *mcp.CallToolRequest, input GetByRayIDInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/logs/rayids/" + input.RayID
	var params []string
	if input.Fields != "" {
		params = append(params, fmt.Sprintf("fields=%s", input.Fields))
	}
	if input.Timestamps != "" {
		params = append(params, fmt.Sprintf("timestamps=%s", input.Timestamps))
	}
	if len(params) > 0 {
		url += "?" + strings.Join(params, "&")
	}

	body, err := doLogpullRequest(ctx, url, apiToken)
	if err != nil {
		return nil, nil, err
	}

	if body == "" {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "No log entry found for the given Ray ID"}},
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: body}},
	}, nil, nil
}

// ListReceivedInput holds parameters for retrieving HTTP request logs by time range.
type ListReceivedInput struct {
	ZoneID     string  `json:"zone_id" jsonschema:"required,The ID of the zone"`
	Start      string  `json:"start" jsonschema:"required,Start timestamp (inclusive) in RFC3339 or UNIX format"`
	End        string  `json:"end" jsonschema:"required,End timestamp (exclusive) in RFC3339 or UNIX format"`
	Fields     string  `json:"fields,omitempty" jsonschema:"Comma-separated list of log fields to return"`
	Count      int     `json:"count,omitempty" jsonschema:"Maximum number of records to return"`
	Sample     float64 `json:"sample,omitempty" jsonschema:"Sampling rate between 0.0 and 1.0"`
	Timestamps string  `json:"timestamps,omitempty" jsonschema:"Timestamp format: unixnano (default), unix, or rfc3339"`
}

func listReceived(ctx context.Context, _ *mcp.CallToolRequest, input ListReceivedInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/logs/received"
	var params []string
	params = append(params, fmt.Sprintf("start=%s", input.Start))
	params = append(params, fmt.Sprintf("end=%s", input.End))
	if input.Fields != "" {
		params = append(params, fmt.Sprintf("fields=%s", input.Fields))
	}
	if input.Count > 0 {
		params = append(params, fmt.Sprintf("count=%d", input.Count))
	}
	if input.Sample > 0 {
		params = append(params, fmt.Sprintf("sample=%f", input.Sample))
	}
	if input.Timestamps != "" {
		params = append(params, fmt.Sprintf("timestamps=%s", input.Timestamps))
	}
	url += "?" + strings.Join(params, "&")

	body, err := doLogpullRequest(ctx, url, apiToken)
	if err != nil {
		return nil, nil, err
	}

	if body == "" {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "No log entries found for the given time range"}},
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: body}},
	}, nil, nil
}

// ListFieldsInput holds parameters for listing available log fields.
type ListFieldsInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
}

func listFields(ctx context.Context, _ *mcp.CallToolRequest, input ListFieldsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/logs/received/fields"

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

// RegisterTools registers log investigation tools with the MCP server.
// doGet performs a GET against the standard Cloudflare REST API and formats the
// response. It is shared by the logs read tools.
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

// ListLogpushJobsInput holds parameters for listing Logpush jobs in a zone.
type ListLogpushJobsInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
}

func listLogpushJobs(ctx context.Context, _ *mcp.CallToolRequest, input ListLogpushJobsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/zones/"+input.ZoneID+"/logpush/jobs", apiToken)
	return result, nil, err
}

// GetLogpushJobInput holds parameters for getting a Logpush job in a zone.
type GetLogpushJobInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
	JobID  string `json:"job_id"  jsonschema:"required,The ID of the Logpush job"`
}

func getLogpushJob(ctx context.Context, _ *mcp.CallToolRequest, input GetLogpushJobInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/zones/"+input.ZoneID+"/logpush/jobs/"+input.JobID, apiToken)
	return result, nil, err
}

// ListLogpushDatasetJobsInput holds parameters for listing a dataset's Logpush jobs in a zone.
type ListLogpushDatasetJobsInput struct {
	ZoneID    string `json:"zone_id"    jsonschema:"required,The ID of the zone"`
	DatasetID string `json:"dataset_id" jsonschema:"required,The Logpush dataset ID (e.g. http_requests, firewall_events)"`
}

func listLogpushDatasetJobs(ctx context.Context, _ *mcp.CallToolRequest, input ListLogpushDatasetJobsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/zones/"+input.ZoneID+"/logpush/datasets/"+input.DatasetID+"/jobs", apiToken)
	return result, nil, err
}

// ListLogpushDatasetFieldsInput holds parameters for listing a dataset's Logpush fields in a zone.
type ListLogpushDatasetFieldsInput struct {
	ZoneID    string `json:"zone_id"    jsonschema:"required,The ID of the zone"`
	DatasetID string `json:"dataset_id" jsonschema:"required,The Logpush dataset ID (e.g. http_requests, firewall_events)"`
}

func listLogpushDatasetFields(ctx context.Context, _ *mcp.CallToolRequest, input ListLogpushDatasetFieldsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/zones/"+input.ZoneID+"/logpush/datasets/"+input.DatasetID+"/fields", apiToken)
	return result, nil, err
}

// ListInstantLogsJobsInput holds parameters for listing Instant Logs jobs in a zone.
type ListInstantLogsJobsInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
}

func listInstantLogsJobs(ctx context.Context, _ *mcp.CallToolRequest, input ListInstantLogsJobsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/zones/"+input.ZoneID+"/logpush/edge/jobs", apiToken)
	return result, nil, err
}

// GetRetentionFlagInput holds parameters for getting a zone's log retention flag.
type GetRetentionFlagInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
}

func getRetentionFlag(ctx context.Context, _ *mcp.CallToolRequest, input GetRetentionFlagInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/zones/"+input.ZoneID+"/logs/control/retention/flag", apiToken)
	return result, nil, err
}

// ListLogDatasetsInput holds parameters for listing Logs Explorer datasets in a zone.
type ListLogDatasetsInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
}

func listLogDatasets(ctx context.Context, _ *mcp.CallToolRequest, input ListLogDatasetsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/zones/"+input.ZoneID+"/logs/explorer/datasets", apiToken)
	return result, nil, err
}

// ListAvailableLogDatasetsInput holds parameters for listing available Logs Explorer datasets in a zone.
type ListAvailableLogDatasetsInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
}

func listAvailableLogDatasets(ctx context.Context, _ *mcp.CallToolRequest, input ListAvailableLogDatasetsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/zones/"+input.ZoneID+"/logs/explorer/datasets/available", apiToken)
	return result, nil, err
}

// GetLogDatasetInput holds parameters for getting a Logs Explorer dataset in a zone.
type GetLogDatasetInput struct {
	ZoneID    string `json:"zone_id"    jsonschema:"required,The ID of the zone"`
	DatasetID string `json:"dataset_id" jsonschema:"required,The ID of the Logs Explorer dataset"`
}

func getLogDataset(ctx context.Context, _ *mcp.CallToolRequest, input GetLogDatasetInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/zones/"+input.ZoneID+"/logs/explorer/datasets/"+input.DatasetID, apiToken)
	return result, nil, err
}

// QueryLogsSQLInput holds parameters for running a Logs Explorer SQL query in a zone.
type QueryLogsSQLInput struct {
	ZoneID string `json:"zone_id" jsonschema:"required,The ID of the zone"`
	Query  string `json:"query"   jsonschema:"required,The SQL query to run against Logs Explorer"`
}

func queryLogsSQL(ctx context.Context, _ *mcp.CallToolRequest, input QueryLogsSQLInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	reqURL := cfapi.APIBase + "/zones/" + input.ZoneID + "/logs/explorer/query/sql?query=" + url.QueryEscape(input.Query)
	result, err := doGet(ctx, reqURL, apiToken)
	return result, nil, err
}

// ListAccountLogpushJobsInput holds parameters for listing Logpush jobs in an account.
type ListAccountLogpushJobsInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func listAccountLogpushJobs(ctx context.Context, _ *mcp.CallToolRequest, input ListAccountLogpushJobsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/jobs", apiToken)
	return result, nil, err
}

// GetAccountLogpushJobInput holds parameters for getting a Logpush job in an account.
type GetAccountLogpushJobInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	JobID     string `json:"job_id"     jsonschema:"required,The ID of the Logpush job"`
}

func getAccountLogpushJob(ctx context.Context, _ *mcp.CallToolRequest, input GetAccountLogpushJobInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/jobs/"+input.JobID, apiToken)
	return result, nil, err
}

// ListAccountLogpushDatasetJobsInput holds parameters for listing a dataset's Logpush jobs in an account.
type ListAccountLogpushDatasetJobsInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	DatasetID string `json:"dataset_id" jsonschema:"required,The Logpush dataset ID (e.g. http_requests, audit_logs)"`
}

func listAccountLogpushDatasetJobs(ctx context.Context, _ *mcp.CallToolRequest, input ListAccountLogpushDatasetJobsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/datasets/"+input.DatasetID+"/jobs", apiToken)
	return result, nil, err
}

// ListAccountLogpushDatasetFieldsInput holds parameters for listing a dataset's Logpush fields in an account.
type ListAccountLogpushDatasetFieldsInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	DatasetID string `json:"dataset_id" jsonschema:"required,The Logpush dataset ID (e.g. http_requests, audit_logs)"`
}

func listAccountLogpushDatasetFields(ctx context.Context, _ *mcp.CallToolRequest, input ListAccountLogpushDatasetFieldsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/datasets/"+input.DatasetID+"/fields", apiToken)
	return result, nil, err
}

// ListAccountLogpushTransformersInput holds parameters for listing Logpush transformers in an account.
type ListAccountLogpushTransformersInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func listAccountLogpushTransformers(ctx context.Context, _ *mcp.CallToolRequest, input ListAccountLogpushTransformersInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/transformers", apiToken)
	return result, nil, err
}

// GetAccountLogpushTransformerInput holds parameters for getting a Logpush transformer in an account.
type GetAccountLogpushTransformerInput struct {
	AccountID     string `json:"account_id"     jsonschema:"required,The ID of the Cloudflare account"`
	TransformerID string `json:"transformer_id" jsonschema:"required,The ID of the Logpush transformer"`
}

func getAccountLogpushTransformer(ctx context.Context, _ *mcp.CallToolRequest, input GetAccountLogpushTransformerInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/transformers/"+input.TransformerID, apiToken)
	return result, nil, err
}

// GetAccountLogpushTransformerContentInput holds parameters for getting a Logpush transformer's content in an account.
type GetAccountLogpushTransformerContentInput struct {
	AccountID     string `json:"account_id"     jsonschema:"required,The ID of the Cloudflare account"`
	TransformerID string `json:"transformer_id" jsonschema:"required,The ID of the Logpush transformer"`
}

func getAccountLogpushTransformerContent(ctx context.Context, _ *mcp.CallToolRequest, input GetAccountLogpushTransformerContentInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/transformers/"+input.TransformerID+"/content", apiToken)
	return result, nil, err
}

// ListAccountLogpushTransformerVersionsInput holds parameters for listing a transformer's versions in an account.
type ListAccountLogpushTransformerVersionsInput struct {
	AccountID     string `json:"account_id"     jsonschema:"required,The ID of the Cloudflare account"`
	TransformerID string `json:"transformer_id" jsonschema:"required,The ID of the Logpush transformer"`
}

func listAccountLogpushTransformerVersions(ctx context.Context, _ *mcp.CallToolRequest, input ListAccountLogpushTransformerVersionsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/transformers/"+input.TransformerID+"/versions", apiToken)
	return result, nil, err
}

// GetAccountAuditLogsInput holds parameters for the account audit logs (v2) endpoint.
type GetAccountAuditLogsInput struct {
	AccountID string `json:"account_id"     jsonschema:"required,The ID of the Cloudflare account"`
	Since     string `json:"since,omitempty" jsonschema:"Start time in RFC3339 format"`
	Before    string `json:"before,omitempty" jsonschema:"End time in RFC3339 format"`
	Limit     int    `json:"limit,omitempty" jsonschema:"Maximum number of entries to return"`
}

func getAccountAuditLogs(ctx context.Context, _ *mcp.CallToolRequest, input GetAccountAuditLogsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	reqURL := cfapi.APIBase + "/accounts/" + input.AccountID + "/logs/audit"
	var params []string
	if input.Since != "" {
		params = append(params, "since="+url.QueryEscape(input.Since))
	}
	if input.Before != "" {
		params = append(params, "before="+url.QueryEscape(input.Before))
	}
	if input.Limit > 0 {
		params = append(params, fmt.Sprintf("limit=%d", input.Limit))
	}
	if len(params) > 0 {
		reqURL += "?" + strings.Join(params, "&")
	}

	result, err := doGet(ctx, reqURL, apiToken)
	return result, nil, err
}

// ListAuditLogProductCategoriesInput holds parameters for listing audit log product categories in an account.
type ListAuditLogProductCategoriesInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func listAuditLogProductCategories(ctx context.Context, _ *mcp.CallToolRequest, input ListAuditLogProductCategoriesInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/logs/audit/product_categories", apiToken)
	return result, nil, err
}

// GetAuditLogHistoryInput holds parameters for getting an audit log entry's change history in an account.
type GetAuditLogHistoryInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	ID        string `json:"id"         jsonschema:"required,The ID of the audit log entry"`
}

func getAuditLogHistory(ctx context.Context, _ *mcp.CallToolRequest, input GetAuditLogHistoryInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/logs/audit/"+input.ID+"/history", apiToken)
	return result, nil, err
}

// GetCMBConfigInput holds parameters for getting the Customer Metadata Boundary config in an account.
type GetCMBConfigInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func getCMBConfig(ctx context.Context, _ *mcp.CallToolRequest, input GetCMBConfigInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/logs/control/cmb/config", apiToken)
	return result, nil, err
}

// ListAccountLogDatasetsInput holds parameters for listing Logs Explorer datasets in an account.
type ListAccountLogDatasetsInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func listAccountLogDatasets(ctx context.Context, _ *mcp.CallToolRequest, input ListAccountLogDatasetsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/logs/explorer/datasets", apiToken)
	return result, nil, err
}

func RegisterTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_log_by_rayid",
		Description: "Look up an HTTP request log entry by its Cloudflare Ray ID. Returns request details including client IP, path, user agent, status code, and security actions. Useful for investigating why a specific request was blocked or challenged.",
	}, getByRayID)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_received_logs",
		Description: "Retrieve HTTP request logs for a Cloudflare zone within a time range. Returns NDJSON log entries. Time range is limited to 1 hour and data must be at least 5 minutes old. Useful for investigating traffic patterns and anomalies.",
	}, listReceived)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_log_fields",
		Description: "List all available HTTP request log fields for a Cloudflare zone. Returns field names and descriptions. Use this to discover which fields can be specified when calling get_log_by_rayid or list_received_logs.",
	}, listFields)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_logpush_jobs",
		Description: "List Logpush jobs for a Cloudflare zone. Returns job details including dataset, destination, and enabled status.",
	}, listLogpushJobs)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_logpush_job",
		Description: "Get details of a specific Logpush job in a Cloudflare zone by job ID.",
	}, getLogpushJob)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_logpush_dataset_jobs",
		Description: "List Logpush jobs for a specific dataset in a Cloudflare zone (e.g. http_requests, firewall_events).",
	}, listLogpushDatasetJobs)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_logpush_dataset_fields",
		Description: "List the available Logpush fields for a specific dataset in a Cloudflare zone.",
	}, listLogpushDatasetFields)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_instant_logs_jobs",
		Description: "List Instant Logs (edge) jobs for a Cloudflare zone.",
	}, listInstantLogsJobs)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_log_retention_flag",
		Description: "Get the log retention flag for a Cloudflare zone (whether Logpull log retention is enabled).",
	}, getRetentionFlag)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_log_datasets",
		Description: "List Logs Explorer datasets configured for a Cloudflare zone.",
	}, listLogDatasets)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_available_log_datasets",
		Description: "List the Logs Explorer datasets available to be configured for a Cloudflare zone.",
	}, listAvailableLogDatasets)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_log_dataset",
		Description: "Get a specific Logs Explorer dataset for a Cloudflare zone by dataset ID.",
	}, getLogDataset)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "query_logs_sql",
		Description: "Run a Logs Explorer SQL query for a Cloudflare zone (GET). Returns the query results.",
	}, queryLogsSQL)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_account_logpush_jobs",
		Description: "List Logpush jobs for a Cloudflare account. Returns job details including dataset, destination, and enabled status.",
	}, listAccountLogpushJobs)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_account_logpush_job",
		Description: "Get details of a specific Logpush job in a Cloudflare account by job ID.",
	}, getAccountLogpushJob)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_account_logpush_dataset_jobs",
		Description: "List Logpush jobs for a specific dataset in a Cloudflare account.",
	}, listAccountLogpushDatasetJobs)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_account_logpush_dataset_fields",
		Description: "List the available Logpush fields for a specific dataset in a Cloudflare account.",
	}, listAccountLogpushDatasetFields)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_account_logpush_transformers",
		Description: "List Logpush transformers for a Cloudflare account.",
	}, listAccountLogpushTransformers)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_account_logpush_transformer",
		Description: "Get a specific Logpush transformer in a Cloudflare account by transformer ID.",
	}, getAccountLogpushTransformer)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_account_logpush_transformer_content",
		Description: "Get the content (transform definition) of a specific Logpush transformer in a Cloudflare account.",
	}, getAccountLogpushTransformerContent)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_account_logpush_transformer_versions",
		Description: "List the versions of a specific Logpush transformer in a Cloudflare account.",
	}, listAccountLogpushTransformerVersions)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_account_audit_logs_v2",
		Description: "Get account audit logs (Version 2) for a Cloudflare account. Supports since/before time filters and a limit. Records who changed what and when.",
	}, getAccountAuditLogs)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_account_audit_log_product_categories",
		Description: "List the product categories available for filtering account audit logs (v2).",
	}, listAuditLogProductCategories)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_account_audit_log_history",
		Description: "Get the resource change history for a specific account audit log entry (v2) by entry ID.",
	}, getAuditLogHistory)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_cmb_config",
		Description: "Get the Customer Metadata Boundary (CMB) config for a Cloudflare account (data localization region for logs).",
	}, getCMBConfig)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_account_log_datasets",
		Description: "List Logs Explorer datasets configured for a Cloudflare account.",
	}, listAccountLogDatasets)
}
