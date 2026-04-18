package dns

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_list_トークンが未設定の場合エラーを返す(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := list(context.Background(), &mcp.CallToolRequest{}, ListInput{ZoneID: "abc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_ListInput_デフォルト値がゼロ値である(t *testing.T) {
	// Arrange & Act
	input := ListInput{}

	// Assert
	if input.ZoneID != "" || input.Type != "" || input.Name != "" || input.Content != "" || input.Page != 0 || input.PerPage != 0 {
		t.Error("got non-zero defaults, want zero values for ListInput")
	}
}
