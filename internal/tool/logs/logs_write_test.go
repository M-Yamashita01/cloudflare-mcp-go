package logs

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_createAccountLogpushJob_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := createAccountLogpushJob(context.Background(), &mcp.CallToolRequest{}, CreateAccountLogpushJobInput{AccountID: "acc123", Config: `{"dataset":"http_requests"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_createAccountLogpushJob_returns_error_when_config_is_invalid_json(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "token")

	// Act
	result, _, err := createAccountLogpushJob(context.Background(), &mcp.CallToolRequest{}, CreateAccountLogpushJobInput{AccountID: "acc123", Config: "oops"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_updateAccountLogpushJob_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := updateAccountLogpushJob(context.Background(), &mcp.CallToolRequest{}, UpdateAccountLogpushJobInput{AccountID: "acc123", JobID: "1", Config: `{"enabled":true}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_deleteAccountLogpushJob_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := deleteAccountLogpushJob(context.Background(), &mcp.CallToolRequest{}, DeleteAccountLogpushJobInput{AccountID: "acc123", JobID: "1"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_getAccountLogpushOwnership_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getAccountLogpushOwnership(context.Background(), &mcp.CallToolRequest{}, GetAccountLogpushOwnershipInput{AccountID: "acc123", Config: `{"destination_conf":"s3://b"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}
