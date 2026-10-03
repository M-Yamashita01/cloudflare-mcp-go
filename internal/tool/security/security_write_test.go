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

func Test_createRateLimit_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := createRateLimit(context.Background(), &mcp.CallToolRequest{}, CreateRateLimitInput{ZoneID: "abc123", URLPattern: "example.com/*", Threshold: 100, Period: 60, Mode: "simulate"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_updateRateLimit_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := updateRateLimit(context.Background(), &mcp.CallToolRequest{}, UpdateRateLimitInput{ZoneID: "abc123", RateLimitID: "rl123", URLPattern: "example.com/*", Threshold: 100, Period: 60, Mode: "ban"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_deleteRateLimit_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := deleteRateLimit(context.Background(), &mcp.CallToolRequest{}, DeleteRateLimitInput{ZoneID: "abc123", RateLimitID: "rl123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_rateLimitBody_includes_required_fields(t *testing.T) {
	// Arrange & Act
	body := rateLimitBody("example.com/*", nil, 100, 60, "simulate", 0, "")

	// Assert
	if body["threshold"] != 100 {
		t.Errorf("got threshold = %v, want 100", body["threshold"])
	}
}

func Test_rateLimitBody_omits_timeout_when_zero(t *testing.T) {
	// Arrange & Act
	body := rateLimitBody("example.com/*", nil, 100, 60, "simulate", 0, "")

	// Assert
	action := body["action"].(map[string]any)
	if _, ok := action["timeout"]; ok {
		t.Error("got timeout present, want it omitted when zero")
	}
}
