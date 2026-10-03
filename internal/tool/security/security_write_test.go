package security

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_createIPAccessRule_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := createIPAccessRule(context.Background(), &mcp.CallToolRequest{}, CreateIPAccessRuleInput{ZoneID: "abc123", Mode: "block", Target: "ip", Value: "203.0.113.10"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_updateIPAccessRule_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := updateIPAccessRule(context.Background(), &mcp.CallToolRequest{}, UpdateIPAccessRuleInput{ZoneID: "abc123", RuleID: "rule123", Mode: "challenge"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_deleteIPAccessRule_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := deleteIPAccessRule(context.Background(), &mcp.CallToolRequest{}, DeleteIPAccessRuleInput{ZoneID: "abc123", RuleID: "rule123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_CreateIPAccessRuleInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := CreateIPAccessRuleInput{}

	// Assert
	if input != (CreateIPAccessRuleInput{}) {
		t.Error("got non-zero defaults, want zero values for CreateIPAccessRuleInput")
	}
}
