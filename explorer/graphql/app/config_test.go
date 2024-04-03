package app_test

import (
	"os"
	"path/filepath"
	"testing"

	graphql "github.com/omni-network/omni/explorer/graphql/app"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestDefaultConfigReference(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()

	cfg := graphql.DefaultConfig()

	path := filepath.Join(tempDir, "explorer_graphql.toml")

	require.NoError(t, os.MkdirAll(tempDir, 0o755))
	require.NoError(t, graphql.WriteConfigTOML(cfg, log.DefaultConfig(), path))

	b, err := os.ReadFile(path)
	require.NoError(t, err)

	tutil.RequireGoldenBytes(t, b, tutil.WithFilename("default_explorer_graphql.toml"))
}
