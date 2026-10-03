package logs

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// sendWrite performs a write request against the standard Cloudflare REST API
// and formats the response. It is shared by the logs write tools.
func sendWrite(ctx context.Context, method, url, apiToken string, body io.Reader) (*mcp.CallToolResult, error) {
	cfResp, err := cfapi.DoRequest(ctx, method, url, apiToken, body)
	if err != nil {
		return nil, err
	}
	if !cfResp.Success {
		return cfapi.APIErrorResult(cfResp.Errors), nil
	}
	return cfapi.FormatResult(cfResp)
}

// invalidJSON returns an error result when s is not valid JSON, or nil otherwise.
func invalidJSON(field, s string) *mcp.CallToolResult {
	if json.Valid([]byte(s)) {
		return nil
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: "Error: " + field + " must be valid JSON"}},
		IsError: true,
	}
}

// CreateAccountLogpushJobInput holds parameters for creating a Logpush job in an account.
type CreateAccountLogpushJobInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Config    string `json:"config"     jsonschema:"required,JSON object describing the Logpush job (dataset, destination_conf, logpull_options, enabled, etc.)"`
}

func createAccountLogpushJob(ctx context.Context, _ *mcp.CallToolRequest, input CreateAccountLogpushJobInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPost, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/jobs", apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// RegisterWriteTools registers logs write (mutation) tools with the MCP server.
//
// It is called only when write mode is enabled via CLOUDFLARE_MCP_ENABLE_WRITE.
func RegisterWriteTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_account_logpush_job",
		Description: "Create a Logpush job for a Cloudflare account. The config argument is a JSON object (dataset, destination_conf, logpull_options, enabled, etc.).",
	}, createAccountLogpushJob)
}
