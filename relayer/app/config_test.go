package relayer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tutil"
	relayer "github.com/omni-network/omni/relayer/app"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestDefaultConfigReference(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()

	cfg := relayer.DefaultConfig()

	path := filepath.Join(tempDir, "relayer.toml")

	require.NoError(t, os.MkdirAll(tempDir, 0o755))
	require.NoError(t, relayer.WriteConfigTOML(cfg, log.DefaultConfig(), path))

	b, err := os.ReadFile(path)
	require.NoError(t, err)

	tutil.RequireGoldenBytes(t, b, tutil.WithFilename("default_relayer.toml"))
}
