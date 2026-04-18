package zone

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_list_トークンが未設定の場合エラーを返す(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := list(context.Background(), &mcp.CallToolRequest{}, ListInput{})

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
	if input.Name != "" || input.Page != 0 || input.PerPage != 0 {
		t.Error("got non-zero defaults, want zero values for ListInput")
	}
}

func Test_get_トークンが未設定の場合エラーを返す(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := get(context.Background(), &mcp.CallToolRequest{}, GetInput{ZoneID: "abc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_GetInput_ZoneIDのデフォルト値が空文字列である(t *testing.T) {
	// Arrange & Act
	input := GetInput{}

	// Assert
	if input.ZoneID != "" {
		t.Errorf("got ZoneID = %q, want empty string", input.ZoneID)
	}
}
