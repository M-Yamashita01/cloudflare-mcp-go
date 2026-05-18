package intel

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_getIPIntel_returns_error_when_token_is_missing(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getIPIntel(context.Background(), &mcp.CallToolRequest{}, GetIPIntelInput{
		AccountID: "acc123",
		IP:        "203.0.113.50",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_GetIPIntelInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := GetIPIntelInput{}

	// Assert
	if input.AccountID != "" || input.IP != "" {
		t.Error("got non-zero defaults, want zero values for GetIPIntelInput")
	}
}

func Test_getDomainIntel_returns_error_when_token_is_missing(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getDomainIntel(context.Background(), &mcp.CallToolRequest{}, GetDomainIntelInput{
		AccountID: "acc123",
		Domain:    "example.com",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_GetDomainIntelInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := GetDomainIntelInput{}

	// Assert
	if input.AccountID != "" || input.Domain != "" {
		t.Error("got non-zero defaults, want zero values for GetDomainIntelInput")
	}
}

func Test_getDomainHistory_returns_error_when_token_is_missing(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getDomainHistory(context.Background(), &mcp.CallToolRequest{}, GetDomainHistoryInput{
		AccountID: "acc123",
		Domain:    "example.com",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_listPassiveDNS_returns_error_when_token_is_missing(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listPassiveDNS(context.Background(), &mcp.CallToolRequest{}, ListPassiveDNSInput{
		AccountID: "acc123",
		IP:        "203.0.113.50",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}
