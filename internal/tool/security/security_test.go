package security

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_listIPAccessRules_トークンが未設定の場合エラーを返す(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listIPAccessRules(context.Background(), &mcp.CallToolRequest{}, ListIPAccessRulesInput{ZoneID: "abc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_ListIPAccessRulesInput_デフォルト値がゼロ値である(t *testing.T) {
	// Arrange & Act
	input := ListIPAccessRulesInput{}

	// Assert
	if input.ZoneID != "" || input.IP != "" || input.Mode != "" || input.Page != 0 || input.PerPage != 0 {
		t.Error("got non-zero defaults, want zero values for ListIPAccessRulesInput")
	}
}

func Test_listWAFManagedRulesets_トークンが未設定の場合エラーを返す(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listWAFManagedRulesets(context.Background(), &mcp.CallToolRequest{}, ListWAFManagedRulesetsInput{ZoneID: "abc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_ListWAFManagedRulesetsInput_ZoneIDのデフォルト値が空文字列である(t *testing.T) {
	// Arrange & Act
	input := ListWAFManagedRulesetsInput{}

	// Assert
	if input.ZoneID != "" {
		t.Errorf("got ZoneID = %q, want empty string", input.ZoneID)
	}
}

func Test_querySecurityEvents_トークンが未設定の場合エラーを返す(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := querySecurityEvents(context.Background(), &mcp.CallToolRequest{}, QuerySecurityEventsInput{
		ZoneID:   "abc123",
		DateFrom: "2026-03-23T08:19:58Z",
		DateTo:   "2026-03-23T08:20:58Z",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_QuerySecurityEventsInput_デフォルト値がゼロ値である(t *testing.T) {
	// Arrange & Act
	input := QuerySecurityEventsInput{}

	// Assert
	if input.ZoneID != "" || input.DateFrom != "" || input.DateTo != "" || input.Source != "" || input.Action != "" || input.ClientIP != "" || input.Host != "" || input.RuleID != "" || input.Country != "" || input.HTTPMethod != "" || input.RequestPath != "" || input.Limit != 0 {
		t.Error("got non-zero defaults, want zero values for QuerySecurityEventsInput")
	}
}
