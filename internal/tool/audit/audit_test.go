package audit

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_listAuditLogs_returns_error_when_token_is_missing(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listAuditLogs(context.Background(), &mcp.CallToolRequest{}, ListAuditLogsInput{
		AccountID: "acc123",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_ListAuditLogsInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := ListAuditLogsInput{}

	// Assert
	if input.AccountID != "" || input.Since != "" || input.Before != "" || input.ActorIP != "" || input.ActorEmail != "" || input.Direction != "" || input.Page != 0 || input.PerPage != 0 {
		t.Error("got non-zero defaults, want zero values for ListAuditLogsInput")
	}
}
