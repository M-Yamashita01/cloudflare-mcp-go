package kv

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_listNamespaces_トークンが未設定の場合エラーを返す(t *testing.T) {
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

func Test_ListNamespacesInput_デフォルト値がゼロ値である(t *testing.T) {
	// Arrange & Act
	input := ListNamespacesInput{}

	// Assert
	if input.AccountID != "" || input.Page != 0 || input.PerPage != 0 || input.Order != "" || input.Direction != "" {
		t.Error("got non-zero defaults, want zero values for ListNamespacesInput")
	}
}
