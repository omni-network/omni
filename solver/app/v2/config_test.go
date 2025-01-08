package appv2_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tutil"
	solver "github.com/omni-network/omni/solver/app"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestDefaultConfigReference(t *testing.T) {
	t.Parallel()
	tempDir := t.TempDir()

	cfg := solver.DefaultConfig()
	cfg.LoadGenPrivKey = "loadgen.key"

	path := filepath.Join(tempDir, "solver.toml")

	require.NoError(t, os.MkdirAll(tempDir, 0o755))
	require.NoError(t, solver.WriteConfigTOML(cfg, log.DefaultConfig(), path))

	b, err := os.ReadFile(path)
	require.NoError(t, err)

	tutil.RequireGoldenBytes(t, b, tutil.WithFilename("default_solver.toml"))
}
