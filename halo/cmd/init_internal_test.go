package cmd

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -run TestInit -golden

func TestInitFiles(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	cfg := InitConfig{
		HomeDir: dir,
		Network: netconf.Simnet,
	}
	err := InitFiles(context.Background(), cfg)
	require.NoError(t, err)

	files, err := filepath.Glob(dir + "/**/*")
	require.NoError(t, err)

	var resp string
	for _, file := range files {
		resp += strings.TrimPrefix(file, dir) + "\n"
	}

	tutil.RequireGoldenBytes(t, []byte(resp))
}

func TestInitForce(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()

	// Create a dummy file
	err := os.WriteFile(filepath.Join(dir, "dummy"), nil, 0o644)
	require.NoError(t, err)

	cfg := InitConfig{
		HomeDir: dir,
		Network: netconf.Simnet,
	}

	err = InitFiles(context.Background(), cfg)
	require.ErrorContains(t, err, "unexpected file")

	cfg.Force = true
	err = InitFiles(context.Background(), cfg)
	require.NoError(t, err)
}
