package logs

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_getByRayID_returns_error_when_token_is_missing(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getByRayID(context.Background(), &mcp.CallToolRequest{}, GetByRayIDInput{
		ZoneID: "abc123",
		RayID:  "ray123",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_GetByRayIDInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := GetByRayIDInput{}

	// Assert
	if input.ZoneID != "" || input.RayID != "" || input.Fields != "" || input.Timestamps != "" {
		t.Error("got non-zero defaults, want zero values for GetByRayIDInput")
	}
}
