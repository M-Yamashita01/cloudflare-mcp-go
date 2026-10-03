package kv

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/M-Yamashita01/cloudflare-mcp-go/internal/cfapi"
)

// WritePairInput holds parameters for writing a key-value pair.
type WritePairInput struct {
	AccountID     string `json:"account_id"              jsonschema:"required,The ID of the Cloudflare account"`
	NamespaceID   string `json:"namespace_id"            jsonschema:"required,The ID of the KV namespace"`
	Key           string `json:"key"                     jsonschema:"required,The key to write"`
	Value         string `json:"value"                   jsonschema:"required,The value to store"`
	ExpirationTTL int    `json:"expiration_ttl,omitempty" jsonschema:"Seconds until the key expires (minimum 60)"`
}

func writePair(ctx context.Context, _ *mcp.CallToolRequest, input WritePairInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	reqURL := cfapi.APIBase + "/accounts/" + input.AccountID + "/storage/kv/namespaces/" + input.NamespaceID + "/values/" + url.PathEscape(input.Key)
	if input.ExpirationTTL > 0 {
		reqURL += fmt.Sprintf("?expiration_ttl=%d", input.ExpirationTTL)
	}

	cfResp, err := cfapi.DoRequest(ctx, http.MethodPut, reqURL, apiToken, strings.NewReader(input.Value))
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

// DeletePairInput holds parameters for deleting a key-value pair.
type DeletePairInput struct {
	AccountID   string `json:"account_id"   jsonschema:"required,The ID of the Cloudflare account"`
	NamespaceID string `json:"namespace_id" jsonschema:"required,The ID of the KV namespace"`
	Key         string `json:"key"          jsonschema:"required,The key to delete"`
}

func deletePair(ctx context.Context, _ *mcp.CallToolRequest, input DeletePairInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	reqURL := cfapi.APIBase + "/accounts/" + input.AccountID + "/storage/kv/namespaces/" + input.NamespaceID + "/values/" + url.PathEscape(input.Key)
	cfResp, err := cfapi.DoRequest(ctx, http.MethodDelete, reqURL, apiToken, nil)
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

// CreateNamespaceInput holds parameters for creating a KV namespace.
type CreateNamespaceInput struct {
	AccountID string `json:"account_id" jsonschema:"required,The ID of the Cloudflare account"`
	Title     string `json:"title"      jsonschema:"required,The title of the new namespace"`
}

func createNamespace(ctx context.Context, _ *mcp.CallToolRequest, input CreateNamespaceInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	payload, err := json.Marshal(map[string]any{"title": input.Title})
	if err != nil {
		return nil, nil, fmt.Errorf("marshaling request body: %w", err)
	}

	reqURL := cfapi.APIBase + "/accounts/" + input.AccountID + "/storage/kv/namespaces"
	cfResp, err := cfapi.DoRequest(ctx, http.MethodPost, reqURL, apiToken, bytes.NewReader(payload))
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

// DeleteNamespaceInput holds parameters for deleting a KV namespace.
type DeleteNamespaceInput struct {
	AccountID   string `json:"account_id"   jsonschema:"required,The ID of the Cloudflare account"`
	NamespaceID string `json:"namespace_id" jsonschema:"required,The ID of the KV namespace to delete"`
}

func deleteNamespace(ctx context.Context, _ *mcp.CallToolRequest, input DeleteNamespaceInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	reqURL := cfapi.APIBase + "/accounts/" + input.AccountID + "/storage/kv/namespaces/" + input.NamespaceID
	cfResp, err := cfapi.DoRequest(ctx, http.MethodDelete, reqURL, apiToken, nil)
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

// RenameNamespaceInput holds parameters for renaming a KV namespace.
type RenameNamespaceInput struct {
	AccountID   string `json:"account_id"   jsonschema:"required,The ID of the Cloudflare account"`
	NamespaceID string `json:"namespace_id" jsonschema:"required,The ID of the KV namespace to rename"`
	Title       string `json:"title"        jsonschema:"required,The new title for the namespace"`
}

func renameNamespace(ctx context.Context, _ *mcp.CallToolRequest, input RenameNamespaceInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	payload, err := json.Marshal(map[string]any{"title": input.Title})
	if err != nil {
		return nil, nil, fmt.Errorf("marshaling request body: %w", err)
	}

	reqURL := cfapi.APIBase + "/accounts/" + input.AccountID + "/storage/kv/namespaces/" + input.NamespaceID
	cfResp, err := cfapi.DoRequest(ctx, http.MethodPut, reqURL, apiToken, bytes.NewReader(payload))
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

// bulkKVRequest validates jsonBody and performs a KV bulk write request.
func bulkKVRequest(ctx context.Context, method, reqURL, apiToken, jsonBody string) (*mcp.CallToolResult, error) {
	if !json.Valid([]byte(jsonBody)) {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "Error: body must be valid JSON"}},
			IsError: true,
		}, nil
	}
	cfResp, err := cfapi.DoRequest(ctx, method, reqURL, apiToken, bytes.NewReader([]byte(jsonBody)))
	if err != nil {
		return nil, err
	}
	if !cfResp.Success {
		return cfapi.APIErrorResult(cfResp.Errors), nil
	}
	return cfapi.FormatResult(cfResp)
}

// BulkWriteInput holds parameters for writing multiple KV pairs.
type BulkWriteInput struct {
	AccountID   string `json:"account_id"   jsonschema:"required,The ID of the Cloudflare account"`
	NamespaceID string `json:"namespace_id" jsonschema:"required,The ID of the KV namespace"`
	Pairs       string `json:"pairs"        jsonschema:"required,JSON array of key-value pair objects (key, value, optional expiration, metadata, base64)"`
}

func bulkWrite(ctx context.Context, _ *mcp.CallToolRequest, input BulkWriteInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	reqURL := cfapi.APIBase + "/accounts/" + input.AccountID + "/storage/kv/namespaces/" + input.NamespaceID + "/bulk"
	result, err := bulkKVRequest(ctx, http.MethodPut, reqURL, apiToken, input.Pairs)
	return result, nil, err
}

// BulkDeleteInput holds parameters for deleting multiple KV pairs (DELETE bulk).
type BulkDeleteInput struct {
	AccountID   string `json:"account_id"   jsonschema:"required,The ID of the Cloudflare account"`
	NamespaceID string `json:"namespace_id" jsonschema:"required,The ID of the KV namespace"`
	Keys        string `json:"keys"         jsonschema:"required,JSON array of key names to delete"`
}

func bulkDelete(ctx context.Context, _ *mcp.CallToolRequest, input BulkDeleteInput) (*mcp.CallToolResult, any, error) {
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")
	if result := cfapi.CheckToken(apiToken); result != nil {
		return result, nil, nil
	}

	reqURL := cfapi.APIBase + "/accounts/" + input.AccountID + "/storage/kv/namespaces/" + input.NamespaceID + "/bulk"
	result, err := bulkKVRequest(ctx, http.MethodDelete, reqURL, apiToken, input.Keys)
	return result, nil, err
}

// RegisterWriteTools registers KV write (mutation) tools with the MCP server.
//
// It is called only when write mode is enabled via CLOUDFLARE_MCP_ENABLE_WRITE.
func RegisterWriteTools(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "write_kv_pair",
		Description: "Write a key-value pair to a Workers KV namespace. Optionally set expiration_ttl (seconds, minimum 60).",
	}, writePair)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_kv_pair",
		Description: "Delete a key-value pair from a Workers KV namespace by key.",
	}, deletePair)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_kv_namespace",
		Description: "Create a Workers KV namespace in a Cloudflare account with the given title.",
	}, createNamespace)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_kv_namespace",
		Description: "Delete a Workers KV namespace from a Cloudflare account by namespace ID.",
	}, deleteNamespace)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "rename_kv_namespace",
		Description: "Rename a Workers KV namespace (change its title) by namespace ID.",
	}, renameNamespace)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "write_kv_pairs_bulk",
		Description: "Write multiple key-value pairs to a Workers KV namespace at once (PUT bulk). The pairs argument is a JSON array of pair objects (key, value, optional expiration/metadata/base64).",
	}, bulkWrite)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_kv_pairs_bulk",
		Description: "Delete multiple key-value pairs from a Workers KV namespace at once (DELETE bulk). The keys argument is a JSON array of key names.",
	}, bulkDelete)
}
