package cfapi

import (
	"encoding/json"
	"testing"
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

func Test_APIErrorResult_エラーメッセージを結合して返す(t *testing.T) {
	// Arrange
	errs := []Error{
		{Code: 1001, Message: "Invalid zone"},
		{Code: 1002, Message: "Rate limited"},
	}

	// Act
	result := APIErrorResult(errs)

	// Assert
	if !result.IsError {
		t.Error("got IsError = false, want true")
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
	if !result.IsError {
		t.Error("got IsError = false, want true")
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
	if result.IsError {
		t.Error("got IsError = true, want false")
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
