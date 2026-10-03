package dns

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_create_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := create(context.Background(), &mcp.CallToolRequest{}, CreateInput{
		ZoneID:  "abc123",
		Type:    "A",
		Name:    "example.com",
		Content: "203.0.113.10",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_update_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := update(context.Background(), &mcp.CallToolRequest{}, UpdateInput{
		ZoneID:   "abc123",
		RecordID: "rec123",
		Content:  "203.0.113.20",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_deleteRecord_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := deleteRecord(context.Background(), &mcp.CallToolRequest{}, DeleteInput{
		ZoneID:   "abc123",
		RecordID: "rec123",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_CreateInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := CreateInput{}

	// Assert
	if input != (CreateInput{}) {
		t.Error("got non-zero defaults, want zero values for CreateInput")
	}
}

func Test_UpdateInput_proxied_defaults_to_nil(t *testing.T) {
	// Arrange & Act
	input := UpdateInput{}

	// Assert
	if input.Proxied != nil {
		t.Errorf("got Proxied = %v, want nil", input.Proxied)
	}
}

func Test_overwrite_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := overwrite(context.Background(), &mcp.CallToolRequest{}, OverwriteInput{
		ZoneID:   "abc123",
		RecordID: "rec123",
		Type:     "A",
		Name:     "example.com",
		Content:  "203.0.113.10",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_scanTrigger_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := scanTrigger(context.Background(), &mcp.CallToolRequest{}, ScanTriggerInput{ZoneID: "abc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_scan_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := scan(context.Background(), &mcp.CallToolRequest{}, ScanInput{ZoneID: "abc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_importRecords_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := importRecords(context.Background(), &mcp.CallToolRequest{}, ImportInput{ZoneID: "abc123", File: "example.com. 300 IN A 203.0.113.10"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_batch_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := batch(context.Background(), &mcp.CallToolRequest{}, BatchInput{ZoneID: "abc123", Operations: `{"deletes":[]}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_batch_returns_error_when_operations_is_invalid_json(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "token")

	// Act
	result, _, err := batch(context.Background(), &mcp.CallToolRequest{}, BatchInput{ZoneID: "abc123", Operations: "not-json"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_DeleteInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := DeleteInput{}

	// Assert
	if input != (DeleteInput{}) {
		t.Error("got non-zero defaults, want zero values for DeleteInput")
	}
}
