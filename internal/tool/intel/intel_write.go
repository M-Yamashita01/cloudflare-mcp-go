package intel

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// sendWrite performs a write request against the standard Cloudflare REST API
// and formats the response. It is shared by the intel write tools.
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

// DismissInsightInput holds parameters for dismissing a Security Center insight.
type DismissInsightInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	IssueID   string `json:"issue_id"   jsonschema:"required,The ID of the attack-surface issue/insight"`
	Dismissed bool   `json:"dismissed,omitempty" jsonschema:"Whether to dismiss (true) or un-dismiss (false) the insight"`
}

func dismissInsight(ctx context.Context, _ *mcp.CallToolRequest, input DismissInsightInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	payload, err := json.Marshal(map[string]any{"dismiss": input.Dismissed})
	if err != nil {
		return nil, nil, fmt.Errorf("marshaling request body: %w", err)
	}

	url := cfapi.APIBase + "/accounts/" + input.AccountID + "/intel/attack-surface-report/issues/" + input.IssueID + "/dismiss"
	result, err := sendWrite(ctx, http.MethodPut, url, apiToken, bytes.NewReader(payload))
	return result, nil, err
}

// CreateIndicatorFeedInput holds parameters for creating an indicator feed.
type CreateIndicatorFeedInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Config    string `json:"config"     jsonschema:"required,JSON object describing the feed (name, description, is_attributable, etc.)"`
}

func createIndicatorFeed(ctx context.Context, _ *mcp.CallToolRequest, input CreateIndicatorFeedInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPost, cfapi.APIBase+"/accounts/"+input.AccountID+"/intel/indicator-feeds", apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// GrantFeedPermissionInput holds parameters for granting indicator feed permission.
type GrantFeedPermissionInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Config    string `json:"config"     jsonschema:"required,JSON object with the feed_id to grant permission for"`
}

func grantFeedPermission(ctx context.Context, _ *mcp.CallToolRequest, input GrantFeedPermissionInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPut, cfapi.APIBase+"/accounts/"+input.AccountID+"/intel/indicator-feeds/permissions/add", apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// CreateFeedProviderInput holds parameters for creating an indicator feed provider.
type CreateFeedProviderInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Config    string `json:"config"     jsonschema:"required,JSON object describing the provider"`
}

func createFeedProvider(ctx context.Context, _ *mcp.CallToolRequest, input CreateFeedProviderInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPut, cfapi.APIBase+"/accounts/"+input.AccountID+"/intel/indicator-feeds/permissions/createProvider", apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// RevokeFeedPermissionInput holds parameters for revoking indicator feed permission.
type RevokeFeedPermissionInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Config    string `json:"config"     jsonschema:"required,JSON object with the feed_id to revoke permission for"`
}

func revokeFeedPermission(ctx context.Context, _ *mcp.CallToolRequest, input RevokeFeedPermissionInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPut, cfapi.APIBase+"/accounts/"+input.AccountID+"/intel/indicator-feeds/permissions/remove", apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// UpdateIndicatorFeedInput holds parameters for updating indicator feed metadata.
type UpdateIndicatorFeedInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	FeedID    string `json:"feed_id"    jsonschema:"required,The ID of the indicator feed"`
	Config    string `json:"config"     jsonschema:"required,JSON object of feed fields to update"`
}

func updateIndicatorFeed(ctx context.Context, _ *mcp.CallToolRequest, input UpdateIndicatorFeedInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}
	if result := invalidJSON("config", input.Config); result != nil {
		return result, nil, nil
	}

	result, err := sendWrite(ctx, http.MethodPut, cfapi.APIBase+"/accounts/"+input.AccountID+"/intel/indicator-feeds/"+input.FeedID, apiToken, bytes.NewReader([]byte(input.Config)))
	return result, nil, err
}

// RegisterWriteTools registers intel write (mutation) tools with the MCP server.
//
// It is called only when write mode is enabled via CLOUDFLARE_MCP_ENABLE_WRITE.
func RegisterWriteTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "dismiss_security_center_insight",
		Description: "Dismiss (archive) or un-dismiss a Security Center attack-surface insight by issue ID. Set dismissed to true to archive.",
	}, dismissInsight)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_indicator_feed",
		Description: "Create a new threat-intelligence indicator feed for a Cloudflare account. The config argument is a JSON object (name, description, etc.).",
	}, createIndicatorFeed)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "grant_indicator_feed_permission",
		Description: "Grant an account permission to view a threat-intelligence indicator feed. The config argument is a JSON object with the feed_id.",
	}, grantFeedPermission)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_indicator_feed_provider",
		Description: "Create an indicator feed provider for a Cloudflare account. The config argument is a JSON object describing the provider.",
	}, createFeedProvider)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "revoke_indicator_feed_permission",
		Description: "Revoke an account's permission to view a threat-intelligence indicator feed. The config argument is a JSON object with the feed_id.",
	}, revokeFeedPermission)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_indicator_feed",
		Description: "Update the metadata of a threat-intelligence indicator feed by feed ID. The config argument is a JSON object of fields to change.",
	}, updateIndicatorFeed)
}
