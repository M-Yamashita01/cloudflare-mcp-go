package logs

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_getByRayID_returns_error_when_token_is_missing(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getByRayID(context.Background(), &mcp.CallToolRequest{}, GetByRayIDInput{
		ZoneID: "abc123",
		RayID:  "ray123",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_GetByRayIDInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := GetByRayIDInput{}

	// Assert
	if input.ZoneID != "" || input.RayID != "" || input.Fields != "" || input.Timestamps != "" {
		t.Error("got non-zero defaults, want zero values for GetByRayIDInput")
	}
}

func Test_listReceived_returns_error_when_token_is_missing(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listReceived(context.Background(), &mcp.CallToolRequest{}, ListReceivedInput{
		ZoneID: "abc123",
		Start:  "2026-05-18T00:00:00Z",
		End:    "2026-05-18T01:00:00Z",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_ListReceivedInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := ListReceivedInput{}

	// Assert
	if input.ZoneID != "" || input.Start != "" || input.End != "" || input.Fields != "" || input.Count != 0 || input.Sample != 0 || input.Timestamps != "" {
		t.Error("got non-zero defaults, want zero values for ListReceivedInput")
	}
}

func Test_listFields_returns_error_when_token_is_missing(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listFields(context.Background(), &mcp.CallToolRequest{}, ListFieldsInput{
		ZoneID: "abc123",
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_ListFieldsInput_has_zero_value_defaults(t *testing.T) {
	// Arrange & Act
	input := ListFieldsInput{}

	// Assert
	if input.ZoneID != "" {
		t.Errorf("got ZoneID = %q, want empty string", input.ZoneID)
	}
}

func Test_listLogpushJobs_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listLogpushJobs(context.Background(), &mcp.CallToolRequest{}, ListLogpushJobsInput{ZoneID: "abc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_getLogpushJob_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getLogpushJob(context.Background(), &mcp.CallToolRequest{}, GetLogpushJobInput{ZoneID: "abc123", JobID: "1"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_listLogpushDatasetJobs_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listLogpushDatasetJobs(context.Background(), &mcp.CallToolRequest{}, ListLogpushDatasetJobsInput{ZoneID: "abc123", DatasetID: "http_requests"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_listLogpushDatasetFields_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listLogpushDatasetFields(context.Background(), &mcp.CallToolRequest{}, ListLogpushDatasetFieldsInput{ZoneID: "abc123", DatasetID: "http_requests"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_listInstantLogsJobs_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listInstantLogsJobs(context.Background(), &mcp.CallToolRequest{}, ListInstantLogsJobsInput{ZoneID: "abc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_getRetentionFlag_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getRetentionFlag(context.Background(), &mcp.CallToolRequest{}, GetRetentionFlagInput{ZoneID: "abc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_listLogDatasets_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listLogDatasets(context.Background(), &mcp.CallToolRequest{}, ListLogDatasetsInput{ZoneID: "abc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_listAvailableLogDatasets_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listAvailableLogDatasets(context.Background(), &mcp.CallToolRequest{}, ListAvailableLogDatasetsInput{ZoneID: "abc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_getLogDataset_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getLogDataset(context.Background(), &mcp.CallToolRequest{}, GetLogDatasetInput{ZoneID: "abc123", DatasetID: "ds1"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_queryLogsSQL_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := queryLogsSQL(context.Background(), &mcp.CallToolRequest{}, QueryLogsSQLInput{ZoneID: "abc123", Query: "SELECT 1"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_listAccountLogpushJobs_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listAccountLogpushJobs(context.Background(), &mcp.CallToolRequest{}, ListAccountLogpushJobsInput{AccountID: "acc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_getAccountLogpushJob_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getAccountLogpushJob(context.Background(), &mcp.CallToolRequest{}, GetAccountLogpushJobInput{AccountID: "acc123", JobID: "1"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_listAccountLogpushDatasetJobs_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listAccountLogpushDatasetJobs(context.Background(), &mcp.CallToolRequest{}, ListAccountLogpushDatasetJobsInput{AccountID: "acc123", DatasetID: "http_requests"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_listAccountLogpushDatasetFields_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listAccountLogpushDatasetFields(context.Background(), &mcp.CallToolRequest{}, ListAccountLogpushDatasetFieldsInput{AccountID: "acc123", DatasetID: "http_requests"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_listAccountLogpushTransformers_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listAccountLogpushTransformers(context.Background(), &mcp.CallToolRequest{}, ListAccountLogpushTransformersInput{AccountID: "acc123"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_getAccountLogpushTransformer_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getAccountLogpushTransformer(context.Background(), &mcp.CallToolRequest{}, GetAccountLogpushTransformerInput{AccountID: "acc123", TransformerID: "t1"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_getAccountLogpushTransformerContent_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := getAccountLogpushTransformerContent(context.Background(), &mcp.CallToolRequest{}, GetAccountLogpushTransformerContentInput{AccountID: "acc123", TransformerID: "t1"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_listAccountLogpushTransformerVersions_returns_error_when_token_is_not_set(t *testing.T) {
	// Arrange
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	// Act
	result, _, err := listAccountLogpushTransformerVersions(context.Background(), &mcp.CallToolRequest{}, ListAccountLogpushTransformerVersionsInput{AccountID: "acc123", TransformerID: "t1"})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}
