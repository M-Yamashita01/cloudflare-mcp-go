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

// RegisterWriteTools registers intel write (mutation) tools with the MCP server.
//
// It is called only when write mode is enabled via CLOUDFLARE_MCP_ENABLE_WRITE.
func RegisterWriteTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "dismiss_security_center_insight",
		Description: "Dismiss (archive) or un-dismiss a Security Center attack-surface insight by issue ID. Set dismissed to true to archive.",
	}, dismissInsight)
}
