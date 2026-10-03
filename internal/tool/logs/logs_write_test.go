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

func Test_validateAccountLogpushOwnership_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := validateAccountLogpushOwnership(context.Background(), &mcp.CallToolRequest{}, ValidateAccountLogpushOwnershipInput{AccountID: "acc123", Config: `{"destination_conf":"s3://b"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_createAccountLogpushTransformer_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := createAccountLogpushTransformer(context.Background(), &mcp.CallToolRequest{}, CreateAccountLogpushTransformerInput{AccountID: "acc123", Config: `{"name":"t"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_previewAccountLogpushTransformer_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := previewAccountLogpushTransformer(context.Background(), &mcp.CallToolRequest{}, PreviewAccountLogpushTransformerInput{AccountID: "acc123", Config: `{"content":"x"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_updateAccountLogpushTransformer_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := updateAccountLogpushTransformer(context.Background(), &mcp.CallToolRequest{}, UpdateAccountLogpushTransformerInput{AccountID: "acc123", TransformerID: "t1", Config: `{"name":"x"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_deleteAccountLogpushTransformer_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := deleteAccountLogpushTransformer(context.Background(), &mcp.CallToolRequest{}, DeleteAccountLogpushTransformerInput{AccountID: "acc123", TransformerID: "t1"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_validateAccountLogpushDestination_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := validateAccountLogpushDestination(context.Background(), &mcp.CallToolRequest{}, ValidateAccountLogpushDestinationInput{AccountID: "acc123", Config: `{"destination_conf":"s3://b"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_checkAccountLogpushDestinationExists_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := checkAccountLogpushDestinationExists(context.Background(), &mcp.CallToolRequest{}, CheckAccountLogpushDestinationExistsInput{AccountID: "acc123", Config: `{"destination_conf":"s3://b"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_validateAccountLogpushOrigin_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := validateAccountLogpushOrigin(context.Background(), &mcp.CallToolRequest{}, ValidateAccountLogpushOriginInput{AccountID: "acc123", Config: `{"dataset":"http_requests"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_updateCMBConfig_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := updateCMBConfig(context.Background(), &mcp.CallToolRequest{}, UpdateCMBConfigInput{AccountID: "acc123", Config: `{"regions":"eu"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_deleteCMBConfig_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := deleteCMBConfig(context.Background(), &mcp.CallToolRequest{}, DeleteCMBConfigInput{AccountID: "acc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_createAccountLogDataset_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := createAccountLogDataset(context.Background(), &mcp.CallToolRequest{}, CreateAccountLogDatasetInput{AccountID: "acc123", Config: `{"name":"ds"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_updateAccountLogDataset_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := updateAccountLogDataset(context.Background(), &mcp.CallToolRequest{}, UpdateAccountLogDatasetInput{AccountID: "acc123", DatasetID: "ds1", Config: `{"name":"x"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_deleteAccountLogDataset_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := deleteAccountLogDataset(context.Background(), &mcp.CallToolRequest{}, DeleteAccountLogDatasetInput{AccountID: "acc123", DatasetID: "ds1"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_runAccountLogsSQL_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := runAccountLogsSQL(context.Background(), &mcp.CallToolRequest{}, RunAccountLogsSQLInput{AccountID: "acc123", Config: `{"query":"SELECT 1"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_createLogpushJob_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := createLogpushJob(context.Background(), &mcp.CallToolRequest{}, CreateLogpushJobInput{ZoneID: "abc123", Config: `{"dataset":"http_requests"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_updateLogpushJob_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := updateLogpushJob(context.Background(), &mcp.CallToolRequest{}, UpdateLogpushJobInput{ZoneID: "abc123", JobID: "1", Config: `{"enabled":true}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_deleteLogpushJob_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := deleteLogpushJob(context.Background(), &mcp.CallToolRequest{}, DeleteLogpushJobInput{ZoneID: "abc123", JobID: "1"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_createInstantLogsJob_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := createInstantLogsJob(context.Background(), &mcp.CallToolRequest{}, CreateInstantLogsJobInput{ZoneID: "abc123", Config: `{"fields":"RayID"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_getLogpushOwnership_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getLogpushOwnership(context.Background(), &mcp.CallToolRequest{}, GetLogpushOwnershipInput{ZoneID: "abc123", Config: `{"destination_conf":"s3://b"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_validateLogpushOwnership_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := validateLogpushOwnership(context.Background(), &mcp.CallToolRequest{}, ValidateLogpushOwnershipInput{ZoneID: "abc123", Config: `{"destination_conf":"s3://b"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_validateLogpushDestination_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := validateLogpushDestination(context.Background(), &mcp.CallToolRequest{}, ValidateLogpushDestinationInput{ZoneID: "abc123", Config: `{"destination_conf":"s3://b"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_checkLogpushDestinationExists_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := checkLogpushDestinationExists(context.Background(), &mcp.CallToolRequest{}, CheckLogpushDestinationExistsInput{ZoneID: "abc123", Config: `{"destination_conf":"s3://b"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_validateLogpushOrigin_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := validateLogpushOrigin(context.Background(), &mcp.CallToolRequest{}, ValidateLogpushOriginInput{ZoneID: "abc123", Config: `{"dataset":"http_requests"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_updateRetentionFlag_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := updateRetentionFlag(context.Background(), &mcp.CallToolRequest{}, UpdateRetentionFlagInput{ZoneID: "abc123", Flag: true})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_createLogDataset_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := createLogDataset(context.Background(), &mcp.CallToolRequest{}, CreateLogDatasetInput{ZoneID: "abc123", Config: `{"name":"ds"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_updateLogDataset_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := updateLogDataset(context.Background(), &mcp.CallToolRequest{}, UpdateLogDatasetInput{ZoneID: "abc123", DatasetID: "ds1", Config: `{"name":"x"}`})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_deleteLogDataset_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := deleteLogDataset(context.Background(), &mcp.CallToolRequest{}, DeleteLogDatasetInput{ZoneID: "abc123", DatasetID: "ds1"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}
