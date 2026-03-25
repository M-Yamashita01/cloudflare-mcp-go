package main

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestListZones_missingToken(t *testing.T) {
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	result, _, err := listZones(context.Background(), &mcp.CallToolRequest{}, ListZonesInput{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Fatal("expected IsError to be true when token is missing")
	}
}

func TestListZonesInput_defaults(t *testing.T) {
	input := ListZonesInput{}
	if input.Name != "" || input.Page != 0 || input.PerPage != 0 {
		t.Fatal("expected zero values for default ListZonesInput")
	}
}
