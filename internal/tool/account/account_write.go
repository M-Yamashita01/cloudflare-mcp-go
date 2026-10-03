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

// DeleteAccountInput holds parameters for deleting an account.
type DeleteAccountInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account to delete"`
}

func deleteAccount(ctx context.Context, _ *mcp.CallToolRequest, input DeleteAccountInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodDelete, cfapi.APIBase+"/accounts/"+input.AccountID, apiToken, nil)
	return result, nil, err
}

// AddMemberInput holds parameters for adding an account member.
type AddMemberInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Config    string `json:"config"     jsonschema:"required,JSON object describing the member (email, roles or policies, status)"`
}

func addMember(ctx context.Context, _ *mcp.CallToolRequest, input AddMemberInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPost, cfapi.APIBase+"/accounts/"+input.AccountID+"/members", apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// UpdateMemberInput holds parameters for updating an account member.
type UpdateMemberInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	MemberID  string `json:"member_id"  jsonschema:"required,The ID of the account member"`
	Config    string `json:"config"     jsonschema:"required,JSON object of member fields to update (roles or policies, status)"`
}

func updateMember(ctx context.Context, _ *mcp.CallToolRequest, input UpdateMemberInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPut, cfapi.APIBase+"/accounts/"+input.AccountID+"/members/"+input.MemberID, apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// RemoveMemberInput holds parameters for removing an account member.
type RemoveMemberInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	MemberID  string `json:"member_id"  jsonschema:"required,The ID of the account member to remove"`
}

func removeMember(ctx context.Context, _ *mcp.CallToolRequest, input RemoveMemberInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodDelete, cfapi.APIBase+"/accounts/"+input.AccountID+"/members/"+input.MemberID, apiToken, nil)
	return result, nil, err
}

// CreateSubscriptionInput holds parameters for creating an account subscription.
type CreateSubscriptionInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Config    string `json:"config"     jsonschema:"required,JSON object describing the subscription (rate_plan, etc.)"`
}

func createSubscription(ctx context.Context, _ *mcp.CallToolRequest, input CreateSubscriptionInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPost, cfapi.APIBase+"/accounts/"+input.AccountID+"/subscriptions", apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// CancelDowngradeInput holds parameters for canceling a delayed subscription downgrade.
type CancelDowngradeInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func cancelDowngrade(ctx context.Context, _ *mcp.CallToolRequest, input CancelDowngradeInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPost, cfapi.APIBase+"/accounts/"+input.AccountID+"/subscriptions/cancel-downgrade", apiToken, nil)
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

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_account",
		Description: "Delete a Cloudflare account by ID. This is destructive and only works for tenant/reseller sub-accounts.",
	}, deleteAccount)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "add_account_member",
		Description: "Add a member to a Cloudflare account. The config argument is a JSON object (email, roles or policies, status).",
	}, addMember)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_account_member",
		Description: "Update a Cloudflare account member by member ID. The config argument is a JSON object of fields to change (roles or policies, status).",
	}, updateMember)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "remove_account_member",
		Description: "Remove a member from a Cloudflare account by member ID.",
	}, removeMember)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_account_subscription",
		Description: "Create a subscription for a Cloudflare account. The config argument is a JSON object (rate_plan, etc.).",
	}, createSubscription)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "cancel_delayed_downgrade",
		Description: "Cancel a pending (delayed) subscription downgrade for a Cloudflare account.",
	}, cancelDowngrade)
}
