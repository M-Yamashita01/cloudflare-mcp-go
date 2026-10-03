package kv

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_listNamespaces_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listNamespaces(context.Background(), &mcp.CallToolRequest{}, ListNamespacesInput{AccountID: "acc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_ListNamespacesInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := ListNamespacesInput{}

	// Assert
	if input.AccountID != "" || input.Page != 0 || input.PerPage != 0 || input.Order != "" || input.Direction != "" {
		t.Error("got non-zero defaults, want zero values for ListNamespacesInput")
	}
}

func Test_listKeys_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listKeys(context.Background(), &mcp.CallToolRequest{}, ListKeysInput{AccountID: "acc123", NamespaceID: "ns123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_getValue_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getValue(context.Background(), &mcp.CallToolRequest{}, GetValueInput{AccountID: "acc123", NamespaceID: "ns123", Key: "k"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_getMetadata_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getMetadata(context.Background(), &mcp.CallToolRequest{}, GetMetadataInput{AccountID: "acc123", NamespaceID: "ns123", Key: "k"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_ListKeysInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := ListKeysInput{}

	// Assert
	if input.AccountID != "" || input.NamespaceID != "" || input.Prefix != "" || input.Limit != 0 || input.Cursor != "" {
		t.Error("got non-zero defaults, want zero values for ListKeysInput")
	}
}

func Test_GetValueInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := GetValueInput{}

	// Assert
	if input != (GetValueInput{}) {
		t.Error("got non-zero defaults, want zero values for GetValueInput")
	}
}

func Test_getNamespace_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getNamespace(context.Background(), &mcp.CallToolRequest{}, GetNamespaceInput{AccountID: "acc123", NamespaceID: "ns123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}
