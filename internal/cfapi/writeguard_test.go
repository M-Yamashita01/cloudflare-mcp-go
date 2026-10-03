package cfapi

import (
	"os"
	"strings"
	"testing"
)

type writeEnabledResult struct {
	enabled bool
	hasWarn bool
}

func Test_WriteEnabled_covers_env_value_partitions(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  writeEnabledResult
	}{
		{name: "exactly true enables write", value: "true", want: writeEnabledResult{enabled: true, hasWarn: false}},
		{name: "empty value stays disabled without warning", value: "", want: writeEnabledResult{enabled: false, hasWarn: false}},
		{name: "false is invalid", value: "false", want: writeEnabledResult{enabled: false, hasWarn: true}},
		{name: "uppercase TRUE is invalid", value: "TRUE", want: writeEnabledResult{enabled: false, hasWarn: true}},
		{name: "1 is invalid", value: "1", want: writeEnabledResult{enabled: false, hasWarn: true}},
		{name: "yes is invalid", value: "yes", want: writeEnabledResult{enabled: false, hasWarn: true}},
		{name: "true with surrounding spaces is invalid", value: " true ", want: writeEnabledResult{enabled: false, hasWarn: true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(EnableWriteEnv, tt.value)

			enabled, warn := WriteEnabled()

			got := writeEnabledResult{enabled: enabled, hasWarn: warn != ""}
			if got != tt.want {
				t.Errorf("WriteEnabled() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_WriteEnabled_returns_false_when_env_is_unset(t *testing.T) {
	t.Setenv(EnableWriteEnv, "")
	if err := os.Unsetenv(EnableWriteEnv); err != nil {
		t.Fatalf("unsetting %s: %v", EnableWriteEnv, err)
	}

	enabled, warn := WriteEnabled()

	got := writeEnabledResult{enabled: enabled, hasWarn: warn != ""}
	want := writeEnabledResult{enabled: false, hasWarn: false}
	if got != want {
		t.Errorf("WriteEnabled() = %+v, want %+v", got, want)
	}
}

func Test_WriteEnabled_warn_names_the_invalid_value(t *testing.T) {
	t.Setenv(EnableWriteEnv, "yes")

	_, warn := WriteEnabled()

	want := `Error: invalid CLOUDFLARE_MCP_ENABLE_WRITE value "yes" (expected "true"); write tools disabled`
	if warn != want {
		t.Errorf("warn = %q, want %q", warn, want)
	}
}

func Test_WriteEnabled_warn_is_empty_when_enabled(t *testing.T) {
	t.Setenv(EnableWriteEnv, "true")

	_, warn := WriteEnabled()

	if warn != "" {
		t.Errorf("warn = %q, want empty", warn)
	}
}

// Guard against accidental renames of the public env var constant that users
// depend on in their MCP configuration.
func Test_EnableWriteEnv_has_expected_name(t *testing.T) {
	if !strings.EqualFold(EnableWriteEnv, "CLOUDFLARE_MCP_ENABLE_WRITE") {
		t.Errorf("EnableWriteEnv = %q, want %q", EnableWriteEnv, "CLOUDFLARE_MCP_ENABLE_WRITE")
	}
}
