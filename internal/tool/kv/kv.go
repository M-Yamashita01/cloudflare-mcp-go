// Package kv provides MCP tools for Cloudflare Workers KV management.
package kv

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// ListNamespacesInput holds query parameters for listing KV namespaces.
type ListNamespacesInput struct {
	AccountID string `json:"account_id"          jsonschema:"required,The ID of the Cloudflare account"`
	Page      int    `json:"page,omitempty"      jsonschema:"Page number of paginated results (default: 1)"`
	PerPage   int    `json:"per_page,omitempty"  jsonschema:"Number of namespaces per page (default: 20, max: 100)"`
	Order     string `json:"order,omitempty"     jsonschema:"Order results by field: id or title"`
	Direction string `json:"direction,omitempty" jsonschema:"Sort direction: asc or desc"`
}

func listNamespaces(ctx context.Context, _ *mcp.CallToolRequest, input ListNamespacesInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	url := cfapi.APIBase + "/accounts/" + input.AccountID + "/storage/kv/namespaces"
	var params []string
	if input.Page > 0 {
		params = append(params, fmt.Sprintf("page=%d", input.Page))
	}
	if input.PerPage > 0 {
		params = append(params, fmt.Sprintf("per_page=%d", input.PerPage))
	}
	if input.Order != "" {
		params = append(params, fmt.Sprintf("order=%s", input.Order))
	}
	if input.Direction != "" {
		params = append(params, fmt.Sprintf("direction=%s", input.Direction))
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

// ListKeysInput holds parameters for listing keys in a KV namespace.
type ListKeysInput struct {
	AccountID   string `json:"account_id"          jsonschema:"required,The ID of the Cloudflare account"`
	NamespaceID string `json:"namespace_id"        jsonschema:"required,The ID of the KV namespace"`
	Prefix      string `json:"prefix,omitempty"    jsonschema:"Filter keys by prefix"`
	Limit       int    `json:"limit,omitempty"     jsonschema:"Number of keys to return (default: 1000, max: 1000)"`
	Cursor      string `json:"cursor,omitempty"    jsonschema:"Pagination cursor from a previous response"`
}

func listKeys(ctx context.Context, _ *mcp.CallToolRequest, input ListKeysInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	reqURL := cfapi.APIBase + "/accounts/" + input.AccountID + "/storage/kv/namespaces/" + input.NamespaceID + "/keys"
	var params []string
	if input.Prefix != "" {
		params = append(params, "prefix="+url.QueryEscape(input.Prefix))
	}
	if input.Limit > 0 {
		params = append(params, fmt.Sprintf("limit=%d", input.Limit))
	}
	if input.Cursor != "" {
		params = append(params, "cursor="+url.QueryEscape(input.Cursor))
	}
	if len(params) > 0 {
		reqURL += "?" + strings.Join(params, "&")
	}

	cfResp, err := cfapi.DoRequest(ctx, http.MethodGet, reqURL, apiToken, nil)
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

// GetValueInput holds parameters for reading a KV pair's value.
type GetValueInput struct {
	AccountID   string `json:"account_id"   jsonschema:"required,The ID of the Cloudflare account"`
	NamespaceID string `json:"namespace_id" jsonschema:"required,The ID of the KV namespace"`
	Key         string `json:"key"          jsonschema:"required,The key whose value to read"`
}

func getValue(ctx context.Context, _ *mcp.CallToolRequest, input GetValueInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	reqURL := cfapi.APIBase + "/accounts/" + input.AccountID + "/storage/kv/namespaces/" + input.NamespaceID + "/values/" + url.PathEscape(input.Key)
	body, status, err := cfapi.DoRawRequest(ctx, http.MethodGet, reqURL, apiToken, nil)
	if err != nil {
		return nil, nil, err
	}
	if status != http.StatusOK {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Cloudflare API error: status %d: %s", status, string(body))}},
			IsError: true,
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: string(body)}},
	}, nil, nil
}

// GetMetadataInput holds parameters for reading a KV pair's metadata.
type GetMetadataInput struct {
	AccountID   string `json:"account_id"   jsonschema:"required,The ID of the Cloudflare account"`
	NamespaceID string `json:"namespace_id" jsonschema:"required,The ID of the KV namespace"`
	Key         string `json:"key"          jsonschema:"required,The key whose metadata to read"`
}

func getMetadata(ctx context.Context, _ *mcp.CallToolRequest, input GetMetadataInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	reqURL := cfapi.APIBase + "/accounts/" + input.AccountID + "/storage/kv/namespaces/" + input.NamespaceID + "/metadata/" + url.PathEscape(input.Key)
	cfResp, err := cfapi.DoRequest(ctx, http.MethodGet, reqURL, apiToken, nil)
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

// GetNamespaceInput holds parameters for getting a KV namespace.
type GetNamespaceInput struct {
	AccountID   string `json:"account_id"   jsonschema:"required,The ID of the Cloudflare account"`
	NamespaceID string `json:"namespace_id" jsonschema:"required,The ID of the KV namespace"`
}

func getNamespace(ctx context.Context, _ *mcp.CallToolRequest, input GetNamespaceInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	cfResp, err := cfapi.DoRequest(ctx, http.MethodGet, cfapi.APIBase+"/accounts/"+input.AccountID+"/storage/kv/namespaces/"+input.NamespaceID, apiToken, nil)
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

// BulkGetInput holds parameters for getting multiple KV pairs at once.
type BulkGetInput struct {
	AccountID   string `json:"account_id"   jsonschema:"required,The ID of the Cloudflare account"`
	NamespaceID string `json:"namespace_id" jsonschema:"required,The ID of the KV namespace"`
	Keys        string `json:"keys"         jsonschema:"required,JSON object with a keys array, e.g. {\"keys\":[\"k1\",\"k2\"]}"`
}

func bulkGet(ctx context.Context, _ *mcp.CallToolRequest, input BulkGetInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	reqURL := cfapi.APIBase + "/accounts/" + input.AccountID + "/storage/kv/namespaces/" + input.NamespaceID + "/bulk/get"
	result, err := bulkKVRequest(ctx, http.MethodPost, reqURL, apiToken, input.Keys)
	return result, nil, err
}

// RegisterTools registers KV management tools with the MCP server.
func RegisterTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_kv_namespaces",
		Description: "List Workers KV namespaces in a Cloudflare account. Returns namespace details such as ID and title.",
	}, listNamespaces)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_kv_keys",
		Description: "List keys in a Workers KV namespace. Supports prefix filtering and cursor pagination. Returns key names, expiration, and metadata.",
	}, listKeys)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_kv_value",
		Description: "Read the value stored for a key in a Workers KV namespace. Returns the raw stored value.",
	}, getValue)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_kv_metadata",
		Description: "Read the metadata associated with a key in a Workers KV namespace.",
	}, getMetadata)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_kv_namespace",
		Description: "Get a Workers KV namespace by ID. Returns the namespace details (ID, title, supports URL encoding).",
	}, getNamespace)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_kv_pairs_bulk",
		Description: "Get multiple key-value pairs from a Workers KV namespace at once. The keys argument is a JSON object with a keys array. Read-only (no mutation) despite using POST.",
	}, bulkGet)
}
