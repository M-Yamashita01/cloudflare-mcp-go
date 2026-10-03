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

// UpdateAccountLogpushJobInput holds parameters for updating a Logpush job in an account.
type UpdateAccountLogpushJobInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	JobID     string `json:"job_id"     jsonschema:"required,The ID of the Logpush job"`
	Config    string `json:"config"     jsonschema:"required,JSON object of Logpush job fields to update"`
}

func updateAccountLogpushJob(ctx context.Context, _ *mcp.CallToolRequest, input UpdateAccountLogpushJobInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPut, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/jobs/"+input.JobID, apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// DeleteAccountLogpushJobInput holds parameters for deleting a Logpush job in an account.
type DeleteAccountLogpushJobInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	JobID     string `json:"job_id"     jsonschema:"required,The ID of the Logpush job to delete"`
}

func deleteAccountLogpushJob(ctx context.Context, _ *mcp.CallToolRequest, input DeleteAccountLogpushJobInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodDelete, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/jobs/"+input.JobID, apiToken, nil)
	return result, nil, err
}

// GetAccountLogpushOwnershipInput holds parameters for requesting a Logpush ownership challenge in an account.
type GetAccountLogpushOwnershipInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Config    string `json:"config"     jsonschema:"required,JSON object with the destination_conf to challenge"`
}

func getAccountLogpushOwnership(ctx context.Context, _ *mcp.CallToolRequest, input GetAccountLogpushOwnershipInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPost, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/ownership", apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// ValidateAccountLogpushOwnershipInput holds parameters for validating a Logpush ownership challenge in an account.
type ValidateAccountLogpushOwnershipInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Config    string `json:"config"     jsonschema:"required,JSON object with destination_conf and ownership_challenge"`
}

func validateAccountLogpushOwnership(ctx context.Context, _ *mcp.CallToolRequest, input ValidateAccountLogpushOwnershipInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPost, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/ownership/validate", apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// CreateAccountLogpushTransformerInput holds parameters for creating a Logpush transformer in an account.
type CreateAccountLogpushTransformerInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Config    string `json:"config"     jsonschema:"required,JSON object describing the transformer (name, content, etc.)"`
}

func createAccountLogpushTransformer(ctx context.Context, _ *mcp.CallToolRequest, input CreateAccountLogpushTransformerInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPost, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/transformers", apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// PreviewAccountLogpushTransformerInput holds parameters for previewing a Logpush transformer in an account.
type PreviewAccountLogpushTransformerInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Config    string `json:"config"     jsonschema:"required,JSON object with the transformer and sample input to preview"`
}

func previewAccountLogpushTransformer(ctx context.Context, _ *mcp.CallToolRequest, input PreviewAccountLogpushTransformerInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPost, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/transformers/preview", apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// UpdateAccountLogpushTransformerInput holds parameters for updating a Logpush transformer in an account.
type UpdateAccountLogpushTransformerInput struct {
	AccountID     string `json:"account_id"     jsonschema:"required,The ID of the Cloudflare account"`
	TransformerID string `json:"transformer_id" jsonschema:"required,The ID of the Logpush transformer"`
	Config        string `json:"config"         jsonschema:"required,JSON object of transformer fields to update"`
}

func updateAccountLogpushTransformer(ctx context.Context, _ *mcp.CallToolRequest, input UpdateAccountLogpushTransformerInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPut, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/transformers/"+input.TransformerID, apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// DeleteAccountLogpushTransformerInput holds parameters for deleting a Logpush transformer in an account.
type DeleteAccountLogpushTransformerInput struct {
	AccountID     string `json:"account_id"     jsonschema:"required,The ID of the Cloudflare account"`
	TransformerID string `json:"transformer_id" jsonschema:"required,The ID of the Logpush transformer to delete"`
}

func deleteAccountLogpushTransformer(ctx context.Context, _ *mcp.CallToolRequest, input DeleteAccountLogpushTransformerInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodDelete, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/transformers/"+input.TransformerID, apiToken, nil)
	return result, nil, err
}

// ValidateAccountLogpushDestinationInput holds parameters for validating a Logpush destination in an account.
type ValidateAccountLogpushDestinationInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Config    string `json:"config"     jsonschema:"required,JSON object with the destination_conf to validate"`
}

func validateAccountLogpushDestination(ctx context.Context, _ *mcp.CallToolRequest, input ValidateAccountLogpushDestinationInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPost, cfapi.APIBase+"/accounts/"+input.AccountID+"/logpush/validate/destination", apiToken, bytes.NewReader([]byte(input.Config)))
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

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_account_logpush_job",
		Description: "Update a Logpush job in a Cloudflare account by job ID. The config argument is a JSON object of fields to change.",
	}, updateAccountLogpushJob)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_account_logpush_job",
		Description: "Delete a Logpush job from a Cloudflare account by job ID.",
	}, deleteAccountLogpushJob)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_account_logpush_ownership_challenge",
		Description: "Request a Logpush ownership challenge for a destination in a Cloudflare account. The config argument is a JSON object with destination_conf.",
	}, getAccountLogpushOwnership)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "validate_account_logpush_ownership",
		Description: "Validate a Logpush ownership challenge for a Cloudflare account. The config argument is a JSON object with destination_conf and ownership_challenge.",
	}, validateAccountLogpushOwnership)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_account_logpush_transformer",
		Description: "Create a Logpush transformer for a Cloudflare account. The config argument is a JSON object (name, content, etc.).",
	}, createAccountLogpushTransformer)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "preview_account_logpush_transformer",
		Description: "Preview the output of a Logpush transformer for a Cloudflare account against sample input. The config argument is a JSON object.",
	}, previewAccountLogpushTransformer)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_account_logpush_transformer",
		Description: "Update a Logpush transformer in a Cloudflare account by transformer ID. The config argument is a JSON object of fields to change.",
	}, updateAccountLogpushTransformer)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_account_logpush_transformer",
		Description: "Delete a Logpush transformer from a Cloudflare account by transformer ID.",
	}, deleteAccountLogpushTransformer)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "validate_account_logpush_destination",
		Description: "Validate a Logpush destination configuration for a Cloudflare account. The config argument is a JSON object with destination_conf.",
	}, validateAccountLogpushDestination)
}
