package kv

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_writePair_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := writePair(context.Background(), &mcp.CallToolRequest{}, WritePairInput{AccountID: "acc123", NamespaceID: "ns123", Key: "k", Value: "v"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_deletePair_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := deletePair(context.Background(), &mcp.CallToolRequest{}, DeletePairInput{AccountID: "acc123", NamespaceID: "ns123", Key: "k"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_createNamespace_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := createNamespace(context.Background(), &mcp.CallToolRequest{}, CreateNamespaceInput{AccountID: "acc123", Title: "my-ns"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_deleteNamespace_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := deleteNamespace(context.Background(), &mcp.CallToolRequest{}, DeleteNamespaceInput{AccountID: "acc123", NamespaceID: "ns123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_WritePairInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := WritePairInput{}

	// Assert
	if input != (WritePairInput{}) {
		t.Error("got non-zero defaults, want zero values for WritePairInput")
	}
}

func Test_CreateNamespaceInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := CreateNamespaceInput{}

	// Assert
	if input != (CreateNamespaceInput{}) {
		t.Error("got non-zero defaults, want zero values for CreateNamespaceInput")
	}
}

func Test_renameNamespace_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := renameNamespace(context.Background(), &mcp.CallToolRequest{}, RenameNamespaceInput{AccountID: "acc123", NamespaceID: "ns123", Title: "new"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_bulkWrite_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := bulkWrite(context.Background(), &mcp.CallToolRequest{}, BulkWriteInput{AccountID: "acc123", NamespaceID: "ns123", Pairs: `[{"key":"k","value":"v"}]`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_bulkWrite_returns_error_when_pairs_is_invalid_json(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "token")

	// Act
	result, _, err := bulkWrite(context.Background(), &mcp.CallToolRequest{}, BulkWriteInput{AccountID: "acc123", NamespaceID: "ns123", Pairs: "oops"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_bulkDelete_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := bulkDelete(context.Background(), &mcp.CallToolRequest{}, BulkDeleteInput{AccountID: "acc123", NamespaceID: "ns123", Keys: `["k1","k2"]`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_bulkDeletePost_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := bulkDeletePost(context.Background(), &mcp.CallToolRequest{}, BulkDeletePostInput{AccountID: "acc123", NamespaceID: "ns123", Keys: `{"keys":["k1"]}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}
