package zone

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// PurgeCacheInput holds parameters for purging a zone's cache.
//
// Exactly one purge mode should be provided. If PurgeEverything is true it
// takes precedence; otherwise the first non-empty list (Files, Tags, Hosts,
// Prefixes) is sent.
type PurgeCacheInput struct {
	ZoneID          string   `json:"zone_id"                   jsonschema:"required,The ID of the zone whose cache to purge"`
	PurgeEverything bool     `json:"purge_everything,omitempty" jsonschema:"Purge all cached files for the zone"`
	Files           []string `json:"files,omitempty"           jsonschema:"Purge specific cached URLs"`
	Tags            []string `json:"tags,omitempty"            jsonschema:"Purge by cache tag (Enterprise only)"`
	Hosts           []string `json:"hosts,omitempty"           jsonschema:"Purge by host (Enterprise only)"`
	Prefixes        []string `json:"prefixes,omitempty"        jsonschema:"Purge by URL prefix (Enterprise only)"`
}

func purgeCache(ctx context.Context, _ *mcp.CallToolRequest, input PurgeCacheInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	body := map[string]any{}
	switch {
	case input.PurgeEverything:
		body["purge_everything"] = true
	case len(input.Files) > 0:
		body["files"] = input.Files
	case len(input.Tags) > 0:
		body["tags"] = input.Tags
	case len(input.Hosts) > 0:
		body["hosts"] = input.Hosts
	case len(input.Prefixes) > 0:
		body["prefixes"] = input.Prefixes
	default:
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "Error: specify purge_everything, files, tags, hosts, or prefixes"}},
			IsError: true,
		}, nil, nil
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return nil, nil, fmt.Errorf("marshaling request body: %w", err)
	}

	url := cfapi.APIBase + "/zones/" + input.ZoneID + "/purge_cache"
	cfResp, err := cfapi.DoRequest(ctx, http.MethodPost, url, apiToken, bytes.NewReader(payload))
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

// RegisterWriteTools registers zone write (mutation) tools with the MCP server.
//
// It is called only when write mode is enabled via CLOUDFLARE_MCP_ENABLE_WRITE.
func RegisterWriteTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "purge_cache",
		Description: "Purge a Cloudflare zone's cache. Choose one mode: purge_everything, or a list of files (URLs), tags, hosts, or prefixes. Tags/hosts/prefixes require an Enterprise plan.",
	}, purgeCache)
}
