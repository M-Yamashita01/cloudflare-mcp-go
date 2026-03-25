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

func TestGetZone_missingToken(t *testing.T) {
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	result, _, err := getZone(context.Background(), &mcp.CallToolRequest{}, GetZoneInput{ZoneID: "abc123"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Fatal("expected IsError to be true when token is missing")
	}
}

func TestGetZoneInput_required(t *testing.T) {
	input := GetZoneInput{}
	if input.ZoneID != "" {
		t.Fatal("expected zero value for default GetZoneInput.ZoneID")
	}
}

func TestListDNSRecords_missingToken(t *testing.T) {
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	result, _, err := listDNSRecords(context.Background(), &mcp.CallToolRequest{}, ListDNSRecordsInput{ZoneID: "abc123"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Fatal("expected IsError to be true when token is missing")
	}
}

func TestListDNSRecordsInput_defaults(t *testing.T) {
	input := ListDNSRecordsInput{}
	if input.ZoneID != "" || input.Type != "" || input.Name != "" || input.Content != "" || input.Page != 0 || input.PerPage != 0 {
		t.Fatal("expected zero values for default ListDNSRecordsInput")
	}
}

func TestListR2Buckets_missingToken(t *testing.T) {
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	result, _, err := listR2Buckets(context.Background(), &mcp.CallToolRequest{}, ListR2BucketsInput{AccountID: "acc123"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Fatal("expected IsError to be true when token is missing")
	}
}

func TestListR2BucketsInput_defaults(t *testing.T) {
	input := ListR2BucketsInput{}
	if input.AccountID != "" || input.NameContains != "" || input.StartAfter != "" || input.PerPage != 0 || input.Direction != "" || input.Jurisdiction != "" {
		t.Fatal("expected zero values for default ListR2BucketsInput")
	}
}

func TestListKVNamespaces_missingToken(t *testing.T) {
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	result, _, err := listKVNamespaces(context.Background(), &mcp.CallToolRequest{}, ListKVNamespacesInput{AccountID: "acc123"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Fatal("expected IsError to be true when token is missing")
	}
}

func TestListKVNamespacesInput_defaults(t *testing.T) {
	input := ListKVNamespacesInput{}
	if input.AccountID != "" || input.Page != 0 || input.PerPage != 0 || input.Order != "" || input.Direction != "" {
		t.Fatal("expected zero values for default ListKVNamespacesInput")
	}
}

func TestListAccounts_missingToken(t *testing.T) {
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	result, _, err := listAccounts(context.Background(), &mcp.CallToolRequest{}, ListAccountsInput{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Fatal("expected IsError to be true when token is missing")
	}
}

func TestListAccountsInput_defaults(t *testing.T) {
	input := ListAccountsInput{}
	if input.Name != "" || input.Page != 0 || input.PerPage != 0 {
		t.Fatal("expected zero values for default ListAccountsInput")
	}
}

func TestDeleteDNSRecord_missingToken(t *testing.T) {
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	result, _, err := deleteDNSRecord(context.Background(), &mcp.CallToolRequest{}, DeleteDNSRecordInput{
		ZoneID:      "abc123",
		DNSRecordID: "rec456",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Fatal("expected IsError to be true when token is missing")
	}
}

func TestUpdateDNSRecord_missingToken(t *testing.T) {
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	result, _, err := updateDNSRecord(context.Background(), &mcp.CallToolRequest{}, UpdateDNSRecordInput{
		ZoneID:      "abc123",
		DNSRecordID: "rec456",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Fatal("expected IsError to be true when token is missing")
	}
}

func TestCreateDNSRecord_missingToken(t *testing.T) {
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	result, _, err := createDNSRecord(context.Background(), &mcp.CallToolRequest{}, CreateDNSRecordInput{
		ZoneID:  "abc123",
		Type:    "A",
		Name:    "example.com",
		Content: "1.2.3.4",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Fatal("expected IsError to be true when token is missing")
	}
}
