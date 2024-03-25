package monitor_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tutil"
	monitor "github.com/omni-network/omni/monitor/app"
	"github.com/omni-network/omni/monitor/loadgen"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestDefaultConfigReference(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()

	cfg := monitor.DefaultConfig()
	cfg.LoadGen = loadgen.Config{
		ValidatorKeysGlob: "path/*/1",
	}

	path := filepath.Join(tempDir, "monitor.toml")

	require.NoError(t, os.MkdirAll(tempDir, 0o755))
	require.NoError(t, monitor.WriteConfigTOML(cfg, log.DefaultConfig(), path))

	b, err := os.ReadFile(path)
	require.NoError(t, err)

	tutil.RequireGoldenBytes(t, b, tutil.WithFilename("default_monitor.toml"))
}
