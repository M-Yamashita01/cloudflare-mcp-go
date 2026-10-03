package securitycenter

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_updateAccountClassification_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := updateAccountClassification(context.Background(), &mcp.CallToolRequest{}, UpdateAccountClassificationInput{AccountID: "acc123", IssueID: "i1", Config: `{"severity":"low"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_updateAccountClassification_returns_error_when_config_is_invalid_json(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "token")

	// Act
	result, _, err := updateAccountClassification(context.Background(), &mcp.CallToolRequest{}, UpdateAccountClassificationInput{AccountID: "acc123", IssueID: "i1", Config: "oops"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_dismissAccountInsight_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := dismissAccountInsight(context.Background(), &mcp.CallToolRequest{}, DismissAccountInsightInput{AccountID: "acc123", IssueID: "i1", Config: `{"dismiss":true}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_startAccountScan_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := startAccountScan(context.Background(), &mcp.CallToolRequest{}, StartAccountScanInput{AccountID: "acc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_updatePartnerSettings_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := updatePartnerSettings(context.Background(), &mcp.CallToolRequest{}, UpdatePartnerSettingsInput{AccountID: "acc123", Partner: "p1", Config: "{}"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_updatePartnerSettings_returns_error_when_config_is_invalid_json(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "token")

	// Act
	result, _, err := updatePartnerSettings(context.Background(), &mcp.CallToolRequest{}, UpdatePartnerSettingsInput{AccountID: "acc123", Partner: "p1", Config: "{invalid"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_updateAccountState_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := updateAccountState(context.Background(), &mcp.CallToolRequest{}, UpdateAccountStateInput{AccountID: "acc123", Config: "{}"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_updateAccountState_returns_error_when_config_is_invalid_json(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "token")

	// Act
	result, _, err := updateAccountState(context.Background(), &mcp.CallToolRequest{}, UpdateAccountStateInput{AccountID: "acc123", Config: "{invalid"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}
