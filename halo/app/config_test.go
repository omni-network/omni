package app_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/halo/app"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/tutil"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestDefaultConfigReference(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()

	cfg := app.DefaultHaloConfig()
	cfg.HomeDir = tempDir

	require.NoError(t, os.MkdirAll(filepath.Join(tempDir, "config"), 0o755))
	require.NoError(t, app.WriteConfigTOML(cfg, log.DefaultConfig()))

	b, err := os.ReadFile(filepath.Join(tempDir, "config", "halo.toml"))
	require.NoError(t, err)

	tutil.RequireGoldenBytes(t, b, tutil.WithFilename("default_halo.toml"))
}
