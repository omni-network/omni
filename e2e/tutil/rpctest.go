package tutil

import (
	"os"
	"testing"

	"github.com/cometbft/cometbft/config"
	rpctest "github.com/cometbft/cometbft/rpc/test"

	"github.com/stretchr/testify/require"

	_ "embed"
)

var (
	//go:embed testdata/genesis.json
	genesisJSON []byte

	//go:embed testdata/priv-validator-key.json
	privValKeyJSON []byte

	//go:embed testdata/priv-validator-state.json
	privValStateJSON []byte
)

// PrepRPCTestConfig creates the require cometbft config on disk for rpctest package to work with halo app.
func PrepRPCTestConfig(t *testing.T) *config.Config {
	t.Helper()

	// Write genesis and priv validator files to temp dir.
	conf := rpctest.GetConfig(true)

	err := os.WriteFile(conf.GenesisFile(), genesisJSON, 0o644)
	require.NoError(t, err)

	err = os.WriteFile(conf.PrivValidatorKeyFile(), privValKeyJSON, 0o644)
	require.NoError(t, err)

	err = os.WriteFile(conf.PrivValidatorStateFile(), privValStateJSON, 0o644)
	require.NoError(t, err)

	return conf
}
