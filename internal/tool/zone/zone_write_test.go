package zone

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_purgeCache_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := purgeCache(context.Background(), &mcp.CallToolRequest{}, PurgeCacheInput{ZoneID: "abc123", PurgeEverything: true})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_purgeCache_returns_error_when_no_mode_is_specified(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "token")

	// Act
	result, _, err := purgeCache(context.Background(), &mcp.CallToolRequest{}, PurgeCacheInput{ZoneID: "abc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_PurgeCacheInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := PurgeCacheInput{}

	// Assert
	if input.ZoneID != "" || input.PurgeEverything || input.Files != nil || input.Tags != nil || input.Hosts != nil || input.Prefixes != nil {
		t.Error("got non-zero defaults, want zero values for PurgeCacheInput")
	}
}
