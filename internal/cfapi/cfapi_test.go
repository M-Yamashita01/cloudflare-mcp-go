package cfapi

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Test_CheckToken_トークンが空の場合エラーを返す(t *testing.T) {
	// Arrange
	apiToken := ""

	// Act
	result := CheckToken(apiToken)

	// Assert
	if result == nil {
		t.Fatal("got nil, want error result")
	}
	if !result.IsError {
		t.Error("got IsError = false, want true")
	}
}

func Test_CheckToken_トークンが設定されている場合nilを返す(t *testing.T) {
	// Arrange
	apiToken := "valid-token"

	// Act
	result := CheckToken(apiToken)

	// Assert
	if result != nil {
		t.Errorf("got %v, want nil", result)
	}
}

func Test_APIErrorResult_複数エラーをセミコロンで結合して返す(t *testing.T) {
	// Arrange
	errs := []Error{
		{Code: 1001, Message: "Invalid zone"},
		{Code: 1002, Message: "Rate limited"},
	}

	// Act
	result := APIErrorResult(errs)

	// Assert
	got := result.Content[0].(*mcp.TextContent).Text
	if !strings.Contains(got, "[1001] Invalid zone") {
		t.Errorf("got %q, want it to contain [1001] Invalid zone", got)
	}
	if !strings.Contains(got, "[1002] Rate limited") {
		t.Errorf("got %q, want it to contain [1002] Rate limited", got)
	}
	if !strings.Contains(got, "; ") {
		t.Errorf("got %q, want errors joined with \"; \"", got)
	}
}

func Test_APIErrorResult_単一エラーの場合コードとメッセージを含む(t *testing.T) {
	// Arrange
	errs := []Error{
		{Code: 9999, Message: "test error"},
	}

	// Act
	result := APIErrorResult(errs)

	// Assert
	got := result.Content[0].(*mcp.TextContent).Text
	if !strings.Contains(got, "[9999] test error") {
		t.Errorf("got %q, want it to contain [9999] test error", got)
	}
}

func Test_FormatResult_JSONを整形して返す(t *testing.T) {
	// Arrange
	cfResp := &Response{
		Success: true,
		Result:  json.RawMessage(`{"id":"abc123"}`),
	}

	// Act
	result, err := FormatResult(cfResp)

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := result.Content[0].(*mcp.TextContent).Text
	want := "{\n  \"id\": \"abc123\"\n}"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func Test_FormatResult_不正なJSONの場合エラーを返す(t *testing.T) {
	// Arrange
	cfResp := &Response{
		Success: true,
		Result:  json.RawMessage(`invalid`),
	}

	// Act
	_, err := FormatResult(cfResp)

	// Assert
	if err == nil {
		t.Fatal("got nil error, want error for invalid JSON")
	}
}
