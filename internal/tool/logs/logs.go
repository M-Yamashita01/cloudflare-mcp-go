// Package logs provides MCP tools for Cloudflare HTTP log investigation.
package logs

import (
	"context"
	"fmt"
	"io"
	"net/http"
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

// RegisterTools registers log investigation tools with the MCP server.
func RegisterTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_log_by_rayid",
		Description: "Look up an HTTP request log entry by its Cloudflare Ray ID. Returns request details including client IP, path, user agent, status code, and security actions. Useful for investigating why a specific request was blocked or challenged.",
	}, getByRayID)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_received_logs",
		Description: "Retrieve HTTP request logs for a Cloudflare zone within a time range. Returns NDJSON log entries. Time range is limited to 1 hour and data must be at least 5 minutes old. Useful for investigating traffic patterns and anomalies.",
	}, listReceived)
}
