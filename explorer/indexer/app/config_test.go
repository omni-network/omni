package app_test

import (
	"os"
	"path/filepath"
	"testing"

	indexer "github.com/omni-network/omni/explorer/indexer/app"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestDefaultConfigReference(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()

	cfg := indexer.DefaultConfig()

	path := filepath.Join(tempDir, "explorer_indexer.toml")

	require.NoError(t, os.MkdirAll(tempDir, 0o755))
	require.NoError(t, indexer.WriteConfigTOML(cfg, log.DefaultConfig(), path))

	b, err := os.ReadFile(path)
	require.NoError(t, err)

	tutil.RequireGoldenBytes(t, b, tutil.WithFilename("default_explorer_indexer.toml"))
}
