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

func TestQuerySecurityEvents_missingToken(t *testing.T) {
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	result, _, err := querySecurityEvents(context.Background(), &mcp.CallToolRequest{}, QuerySecurityEventsInput{
		ZoneID:   "abc123",
		DateFrom: "2026-03-23T08:19:58Z",
		DateTo:   "2026-03-23T08:20:58Z",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Fatal("expected IsError to be true when token is missing")
	}
}

func TestQuerySecurityEventsInput_defaults(t *testing.T) {
	input := QuerySecurityEventsInput{}
	if input.ZoneID != "" || input.DateFrom != "" || input.DateTo != "" || input.Source != "" || input.Action != "" || input.ClientIP != "" || input.Host != "" || input.RuleID != "" || input.Country != "" || input.HTTPMethod != "" || input.RequestPath != "" || input.Limit != 0 {
		t.Fatal("expected zero values for default QuerySecurityEventsInput")
	}
}

func TestListIPAccessRules_missingToken(t *testing.T) {
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	result, _, err := listIPAccessRules(context.Background(), &mcp.CallToolRequest{}, ListIPAccessRulesInput{ZoneID: "abc123"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Fatal("expected IsError to be true when token is missing")
	}
}

func TestListIPAccessRulesInput_defaults(t *testing.T) {
	input := ListIPAccessRulesInput{}
	if input.ZoneID != "" || input.IP != "" || input.Mode != "" || input.Page != 0 || input.PerPage != 0 {
		t.Fatal("expected zero values for default ListIPAccessRulesInput")
	}
}

func TestListWAFManagedRulesets_missingToken(t *testing.T) {
	t.Setenv("CLOUDFLARE_API_TOKEN", "")

	result, _, err := listWAFManagedRulesets(context.Background(), &mcp.CallToolRequest{}, ListWAFManagedRulesetsInput{ZoneID: "abc123"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Fatal("expected IsError to be true when token is missing")
	}
}

func TestListWAFManagedRulesetsInput_required(t *testing.T) {
	input := ListWAFManagedRulesetsInput{}
	if input.ZoneID != "" {
		t.Fatal("expected zero value for default ListWAFManagedRulesetsInput.ZoneID")
	}
}
