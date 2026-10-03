package account

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
// and formats the response. It is shared by the account write tools.
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

// CreateAccountInput holds parameters for creating an account.
type CreateAccountInput struct {
	Config string `json:"config" jsonschema:"required,JSON object describing the account (name, type, unit)"`
}

func createAccount(ctx context.Context, _ *mcp.CallToolRequest, input CreateAccountInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPost, cfapi.APIBase+"/accounts", apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// UpdateAccountInput holds parameters for updating an account.
type UpdateAccountInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Config    string `json:"config"     jsonschema:"required,JSON object of account fields to update (name, settings)"`
}

func updateAccount(ctx context.Context, _ *mcp.CallToolRequest, input UpdateAccountInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPut, cfapi.APIBase+"/accounts/"+input.AccountID, apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// RegisterWriteTools registers account write (mutation) tools with the MCP server.
//
// It is called only when write mode is enabled via CLOUDFLARE_MCP_ENABLE_WRITE.
func RegisterWriteTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_account",
		Description: "Create a new Cloudflare account. The config argument is a JSON object (name, type, and unit for tenant/reseller accounts).",
	}, createAccount)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_account",
		Description: "Update a Cloudflare account by ID. The config argument is a JSON object of fields to change (name, settings).",
	}, updateAccount)
}
