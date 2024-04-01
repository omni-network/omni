package app

import (
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enode"

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

func TestWriteGethConfigTOML(t *testing.T) {
	t.Parallel()

	testKey, _ := crypto.HexToECDSA("45a915e4d060149eb4365960e6a7a45f334393093061116b197e3240065ff2d8")
	node1 := enode.NewV4(&testKey.PublicKey, net.IP{127, 0, 0, 1}, 1, 1)
	node2 := enode.NewV4(&testKey.PublicKey, net.IP{127, 0, 0, 2}, 2, 2)

	cfg := GethConfig{
		peers:     []*enode.Node{node1, node2},
		ChainID:   15651,
		IsArchive: true,
	}

	tempFile := filepath.Join(t.TempDir(), "geth.toml")

	err := WriteGethConfigTOML(cfg, tempFile)
	require.NoError(t, err)

	bz, err := os.ReadFile(tempFile)
	require.NoError(t, err)

	tutil.RequireGoldenBytes(t, bz)
}
