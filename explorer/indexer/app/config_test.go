package app_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/e2e/tutil"
	"github.com/omni-network/omni/explorer/indexer/app"
	"github.com/omni-network/omni/lib/log"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestWriteConfigTOML(t *testing.T) {
	t.Parallel()

	cfg := app.DefaultConfig()
	logCfg := log.DefaultConfig()

	dir := t.TempDir()

	err := app.WriteConfigTOML(cfg, logCfg, dir)
	require.NoError(t, err)

	bz, err := os.ReadFile(filepath.Join(dir, "indexer.toml"))
	require.NoError(t, err)

	tutil.RequireGoldenBytes(t, bz)
}
