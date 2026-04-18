// Package cfapi provides shared types and helpers for the Cloudflare API v4.
package cfapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// APIBase is the base URL for the Cloudflare API v4.
const APIBase = "https://api.cloudflare.com/client/v4"

// Response is the standard Cloudflare API v4 response envelope.
type Response struct {
	Success  bool            `json:"success"`
	Errors   []Error         `json:"errors"`
	Messages []any           `json:"messages"`
	Result   json.RawMessage `json:"result"`
}

// Error holds error details returned by the Cloudflare API.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// CheckToken returns an error CallToolResult when apiToken is empty, or nil if it is set.
func CheckToken(apiToken string) *mcp.CallToolResult {
	if apiToken == "" {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "Error: CLOUDFLARE_API_TOKEN environment variable is not set"}},
			IsError: true,
		}
	}
	return nil
}

// APIErrorResult converts Cloudflare API errors into an error CallToolResult.
func APIErrorResult(errs []Error) *mcp.CallToolResult {
	var msgs []string
	for _, e := range errs {
		msgs = append(msgs, fmt.Sprintf("[%d] %s", e.Code, e.Message))
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: "Cloudflare API error: " + strings.Join(msgs, "; ")}},
		IsError: true,
	}
}

// DoRequest executes an HTTP request against the Cloudflare API and returns the parsed response.
func DoRequest(ctx context.Context, method, url, apiToken string, body io.Reader) (*Response, error) {
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

	var cfResp Response
	if err := json.Unmarshal(respBody, &cfResp); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}
	return &cfResp, nil
}

// FormatResult pretty-prints the result field of a Response as a CallToolResult.
func FormatResult(cfResp *Response) (*mcp.CallToolResult, error) {
	formatted, err := json.MarshalIndent(cfResp.Result, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("formatting result: %w", err)
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: string(formatted)}},
	}, nil
}
