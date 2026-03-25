package main

import (
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

type ListZonesInput struct {
	Name    string `json:"name,omitempty" jsonschema:"A domain name to filter zones by"`
	Page    int    `json:"page,omitempty" jsonschema:"Page number of paginated results (default: 1)"`
	PerPage int    `json:"per_page,omitempty" jsonschema:"Number of zones per page (default: 20, max: 50)"`
}

type CloudflareResponse struct {
	Success  bool            `json:"success"`
	Errors   []CloudflareErr `json:"errors"`
	Messages []any           `json:"messages"`
	Result   json.RawMessage `json:"result"`
}

type CloudflareErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func listZones(ctx context.Context, req *mcp.CallToolRequest, input ListZonesInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if apiToken == "" {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "Error: CLOUDFLARE_API_TOKEN environment variable is not set"}},
			IsError: true,
		}, nil, nil
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

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("creating request: %w", err)
	}
	httpReq.Header.Set("Authorization", "Bearer "+apiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, nil, fmt.Errorf("calling Cloudflare API: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("reading response body: %w", err)
	}

	var cfResp CloudflareResponse
	if err := json.Unmarshal(body, &cfResp); err != nil {
		return nil, nil, fmt.Errorf("parsing response: %w", err)
	}

	if !cfResp.Success {
		var errMsgs []string
		for _, e := range cfResp.Errors {
			errMsgs = append(errMsgs, fmt.Sprintf("[%d] %s", e.Code, e.Message))
		}
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "Cloudflare API error: " + strings.Join(errMsgs, "; ")}},
			IsError: true,
		}, nil, nil
	}

	formatted, err := json.MarshalIndent(json.RawMessage(cfResp.Result), "", "  ")
	if err != nil {
		return nil, nil, fmt.Errorf("formatting result: %w", err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: string(formatted)}},
	}, nil, nil
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

	log.Println("Starting Cloudflare MCP server (stdio)...")
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}
