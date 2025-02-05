package monitor_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/xchain"
	monitor "github.com/omni-network/omni/monitor/app"
	"github.com/omni-network/omni/monitor/loadgen"
	"github.com/omni-network/omni/monitor/xfeemngr"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestDefaultConfigReference(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()

	cfg := monitor.DefaultConfig()
	cfg.HaloGRPCURL = "localhost:1111"
	cfg.LoadGen = loadgen.Config{
		ValidatorKeysGlob: "path/*/1",
		XCallerKey:        "path/xcaller_privatekey",
	}
	cfg.XFeeMngr = xfeemngr.Config{
		RPCEndpoints:    xchain.RPCEndpoints{"test_chain": "http://localhost:8545"},
		CoinGeckoAPIKey: "secret",
	}

	path := filepath.Join(tempDir, "monitor.toml")

	require.NoError(t, os.MkdirAll(tempDir, 0o755))
	require.NoError(t, monitor.WriteConfigTOML(cfg, log.DefaultConfig(), path))

	b, err := os.ReadFile(path)
	require.NoError(t, err)

	tutil.RequireGoldenBytes(t, b, tutil.WithFilename("default_monitor.toml"))
}
