package app

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestUpdateConfigStateSync(t *testing.T) {
	t.Parallel()
	config := `
# Database directory
db_dir = "data"

# Output level for logging, including package level options
log_level = "error"

# Output format: 'plain' (colored text) or 'json'
log_format = "plain"

# For Cosmos SDK-based chains, trust_period should usually be about 2/3 of the unbonding time (~2
# weeks) during which they can be financially punished (slashed) for misbehavior.
rpc_servers = ""
trust_height = 0
trust_hash = ""
trust_period = "168h0m0s"
`

	dir := os.TempDir()
	configFile := filepath.Join(dir, "config", "config.toml")
	require.NoError(t, os.MkdirAll(filepath.Dir(configFile), 0o755))
	err := os.WriteFile(configFile, []byte(config), 0o644)
	require.NoError(t, err)

	err = updateConfigStateSync(dir, 1, []byte("test"))
	require.NoError(t, err)

	bz, err := os.ReadFile(configFile)
	require.NoError(t, err)

	tutil.RequireGoldenBytes(t, bz)
}
