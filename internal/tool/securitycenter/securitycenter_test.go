package securitycenter

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_listInsights_returns_error_when_token_is_missing(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listInsights(context.Background(), &mcp.CallToolRequest{}, ListInsightsInput{
		ZoneID: "abc123",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_ListInsightsInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := ListInsightsInput{}

	// Assert
	if input.ZoneID != "" || input.Severity != "" || input.IssueType != "" || input.IssueClass != "" {
		t.Error("got non-zero defaults, want zero values for ListInsightsInput")
	}
}

func Test_getInsightCounts_returns_error_when_token_is_missing(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getInsightCounts(context.Background(), &mcp.CallToolRequest{}, GetInsightCountsInput{
		ZoneID:    "abc123",
		Dimension: "severity",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}
