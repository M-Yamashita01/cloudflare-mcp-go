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

func Test_listAccountInsights_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listAccountInsights(context.Background(), &mcp.CallToolRequest{}, ListAccountInsightsInput{AccountID: "acc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_getAccountInsightCounts_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getAccountInsightCounts(context.Background(), &mcp.CallToolRequest{}, GetAccountInsightCountsInput{AccountID: "acc123", Dimension: "class"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_getAccountInsightsAuditLog_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getAccountInsightsAuditLog(context.Background(), &mcp.CallToolRequest{}, GetAccountInsightsAuditLogInput{AccountID: "acc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_getAccountPartnerCount_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getAccountPartnerCount(context.Background(), &mcp.CallToolRequest{}, GetAccountPartnerCountInput{AccountID: "acc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_getAccountScans_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getAccountScans(context.Background(), &mcp.CallToolRequest{}, GetAccountScansInput{AccountID: "acc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_getAccountIssueAuditLog_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getAccountIssueAuditLog(context.Background(), &mcp.CallToolRequest{}, GetAccountIssueAuditLogInput{AccountID: "acc123", IssueID: "i1"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_getAccountInsightContext_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getAccountInsightContext(context.Background(), &mcp.CallToolRequest{}, GetAccountInsightContextInput{AccountID: "acc123", IssueID: "i1"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_getPartnerSettings_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getPartnerSettings(context.Background(), &mcp.CallToolRequest{}, GetPartnerSettingsInput{AccountID: "acc123", Partner: "p"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_listShadowZones_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listShadowZones(context.Background(), &mcp.CallToolRequest{}, ListShadowZonesInput{AccountID: "acc123", Partner: "p"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}
