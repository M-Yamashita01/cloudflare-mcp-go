package security

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_listIPAccessRules_returns_error_when_token_is_not_set(t *testing.T) {
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

func Test_ListIPAccessRulesInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := ListIPAccessRulesInput{}

	// Assert
	if input.ZoneID != "" || input.IP != "" || input.Mode != "" || input.Page != 0 || input.PerPage != 0 {
		t.Error("got non-zero defaults, want zero values for ListIPAccessRulesInput")
	}
}

func Test_listWAFManagedRulesets_returns_error_when_token_is_not_set(t *testing.T) {
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

func Test_ListWAFManagedRulesetsInput_ZoneID_defaults_to_empty_string(t *testing.T) {
	// Arrange & Act
	input := ListWAFManagedRulesetsInput{}

	// Assert
	if input.ZoneID != "" {
		t.Errorf("got ZoneID = %q, want empty string", input.ZoneID)
	}
}

func Test_querySecurityEvents_returns_error_when_token_is_not_set(t *testing.T) {
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

func Test_hasBotScoreError_returns_true_when_error_contains_botscoresrcname(t *testing.T) {
	// Arrange
	resp := &graphqlResponse{
		Errors: []graphqlError{
			{Message: "zone 'abc' does not have access to the field 'botscoresrcname' from the path"},
		},
	}

	// Act
	got := hasBotScoreError(resp)

	// Assert
	if !got {
		t.Error("got false, want true")
	}
}

func Test_hasBotScoreError_returns_true_when_error_contains_botscore(t *testing.T) {
	// Arrange
	resp := &graphqlResponse{
		Errors: []graphqlError{
			{Message: "zone 'abc' does not have access to the field 'botscore'"},
		},
	}

	// Act
	got := hasBotScoreError(resp)

	// Assert
	if !got {
		t.Error("got false, want true")
	}
}

func Test_hasBotScoreError_returns_false_when_no_bot_score_error(t *testing.T) {
	// Arrange
	resp := &graphqlResponse{
		Errors: []graphqlError{
			{Message: "some other error"},
		},
	}

	// Act
	got := hasBotScoreError(resp)

	// Assert
	if got {
		t.Error("got true, want false")
	}
}

func Test_hasBotScoreError_returns_false_when_no_errors(t *testing.T) {
	// Arrange
	resp := &graphqlResponse{}

	// Act
	got := hasBotScoreError(resp)

	// Assert
	if got {
		t.Error("got true, want false")
	}
}

func Test_QuerySecurityEventsInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := QuerySecurityEventsInput{}

	// Assert
	if input.ZoneID != "" || input.DateFrom != "" || input.DateTo != "" || input.Source != "" || input.Action != "" || input.ClientIP != "" || input.Host != "" || input.RuleID != "" || input.Country != "" || input.HTTPMethod != "" || input.RequestPath != "" || input.Limit != 0 {
		t.Error("got non-zero defaults, want zero values for QuerySecurityEventsInput")
	}
}
