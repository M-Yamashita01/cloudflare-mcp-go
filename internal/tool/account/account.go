// Package account provides MCP tools for Cloudflare account management.
package account

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// ListInput holds query parameters for listing accounts.
type ListInput struct {
	Name    string `json:"name,omitempty"     jsonschema:"Account name to filter by"`
	Page    int    `json:"page,omitempty"     jsonschema:"Page number of paginated results (default: 1)"`
	PerPage int    `json:"per_page,omitempty" jsonschema:"Number of accounts per page (default: 20, max: 50)"`
}

func list(ctx context.Context, _ *mcp.CallToolRequest, input ListInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/accounts"
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

// doGet performs a GET against the standard Cloudflare REST API and formats the
// response. It is shared by the account read tools.
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

// GetDetailsInput holds parameters for getting account details.
type GetDetailsInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func getDetails(ctx context.Context, _ *mcp.CallToolRequest, input GetDetailsInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID, apiToken)
	return result, nil, err
}

// ListMembersInput holds parameters for listing account members.
type ListMembersInput struct {
	AccountID string `json:"account_id"         jsonschema:"required,The ID of the Cloudflare account"`
	Page      int    `json:"page,omitempty"     jsonschema:"Page number of paginated results (default: 1)"`
	PerPage   int    `json:"per_page,omitempty" jsonschema:"Number of members per page (default: 20, max: 100)"`
}

func listMembers(ctx context.Context, _ *mcp.CallToolRequest, input ListMembersInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/accounts/" + input.AccountID + "/members"
	var params []string
	if input.Page > 0 {
		params = append(params, fmt.Sprintf("page=%d", input.Page))
	}
	if input.PerPage > 0 {
		params = append(params, fmt.Sprintf("per_page=%d", input.PerPage))
	}
	if len(params) > 0 {
		url += "?" + strings.Join(params, "&")
	}

	result, err := doGet(ctx, url, apiToken)
	return result, nil, err
}

// GetMemberInput holds parameters for getting an account member.
type GetMemberInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	MemberID  string `json:"member_id"  jsonschema:"required,The ID of the account member"`
}

func getMember(ctx context.Context, _ *mcp.CallToolRequest, input GetMemberInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/members/"+input.MemberID, apiToken)
	return result, nil, err
}

// ListRolesInput holds parameters for listing account roles.
type ListRolesInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
}

func listRoles(ctx context.Context, _ *mcp.CallToolRequest, input ListRolesInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	result, err := doGet(ctx, cfapi.APIBase+"/accounts/"+input.AccountID+"/roles", apiToken)
	return result, nil, err
}

// RegisterTools registers account management tools with the MCP server.
func RegisterTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_accounts",
		Description: "List Cloudflare accounts accessible with the current API token. Returns account details such as ID, name, and settings.",
	}, list)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_account_details",
		Description: "Get details of a specific Cloudflare account by ID (name, settings, created date).",
	}, getDetails)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_account_members",
		Description: "List members of a Cloudflare account. Returns member details including user, roles, and status.",
	}, listMembers)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_account_member",
		Description: "Get details of a specific Cloudflare account member by member ID (user, roles, policies, status).",
	}, getMember)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_account_roles",
		Description: "List the roles available in a Cloudflare account. Returns role IDs, names, descriptions, and permissions.",
	}, listRoles)
}
