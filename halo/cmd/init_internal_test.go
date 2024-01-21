package cmd

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/omni-network/omni/test/tutil"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -run TestInit -update

func TestInitFiles(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	err := initFiles(context.Background(), dir, false)
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

	err = initFiles(context.Background(), dir, false)
	require.ErrorContains(t, err, "unexpected file")

	err = initFiles(context.Background(), dir, true)
	require.NoError(t, err)
}
