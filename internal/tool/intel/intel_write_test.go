package intel

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_dismissInsight_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := dismissInsight(context.Background(), &mcp.CallToolRequest{}, DismissInsightInput{AccountID: "acc123", IssueID: "i1", Dismissed: true})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_createIndicatorFeed_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := createIndicatorFeed(context.Background(), &mcp.CallToolRequest{}, CreateIndicatorFeedInput{AccountID: "acc123", Config: `{"name":"f"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_createIndicatorFeed_returns_error_when_config_is_invalid_json(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "token")

	// Act
	result, _, err := createIndicatorFeed(context.Background(), &mcp.CallToolRequest{}, CreateIndicatorFeedInput{AccountID: "acc123", Config: "oops"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_grantFeedPermission_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := grantFeedPermission(context.Background(), &mcp.CallToolRequest{}, GrantFeedPermissionInput{AccountID: "acc123", Config: `{"feed_id":1}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_createFeedProvider_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := createFeedProvider(context.Background(), &mcp.CallToolRequest{}, CreateFeedProviderInput{AccountID: "acc123", Config: `{"name":"p"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_revokeFeedPermission_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := revokeFeedPermission(context.Background(), &mcp.CallToolRequest{}, RevokeFeedPermissionInput{AccountID: "acc123", Config: `{"feed_id":1}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_updateIndicatorFeed_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := updateIndicatorFeed(context.Background(), &mcp.CallToolRequest{}, UpdateIndicatorFeedInput{AccountID: "acc123", FeedID: "1", Config: `{"description":"x"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}
