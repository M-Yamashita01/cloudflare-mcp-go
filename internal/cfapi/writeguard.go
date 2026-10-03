package cfapi

import (
	"fmt"
	"os"
)

// EnableWriteEnv is the environment variable that gates write (mutation) tools.
const EnableWriteEnv = "CLOUDFLARE_MCP_ENABLE_WRITE"

// enableWriteValue is the only value that enables write tools.
const enableWriteValue = "true"

// WriteEnabled reports whether write (mutation) tools should be enabled.
//
// Write tools are disabled by default. They are enabled only when the
// CLOUDFLARE_MCP_ENABLE_WRITE environment variable is set to exactly "true".
// An unset or empty value keeps write tools disabled without a warning. Any
// other value keeps write tools disabled and returns a non-empty warn message
// describing the invalid value; callers should log warn and continue serving
// read-only tools.
func WriteEnabled() (enabled bool, warn string) {
	v, ok := os.LookupEnv(EnableWriteEnv)
	if !ok || v == "" {
		return false, ""
	}
	if v == enableWriteValue {
		return true, ""
	}
	return false, fmt.Sprintf("Error: invalid %s value %q (expected %q); write tools disabled", EnableWriteEnv, v, enableWriteValue)
}
