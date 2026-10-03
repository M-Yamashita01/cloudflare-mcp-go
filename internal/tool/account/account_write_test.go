package account

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_createAccount_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := createAccount(context.Background(), &mcp.CallToolRequest{}, CreateAccountInput{Config: `{"name":"acme"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_createAccount_returns_error_when_config_is_invalid_json(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "token")

	// Act
	result, _, err := createAccount(context.Background(), &mcp.CallToolRequest{}, CreateAccountInput{Config: "oops"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_updateAccount_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := updateAccount(context.Background(), &mcp.CallToolRequest{}, UpdateAccountInput{AccountID: "acc123", Config: `{"name":"acme"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}
