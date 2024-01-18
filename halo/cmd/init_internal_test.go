package cmd

import (
	"context"
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
	err := initFiles(context.Background(), dir)
	require.NoError(t, err)

	files, err := filepath.Glob(dir + "/**/*")
	require.NoError(t, err)

	var resp string
	for _, file := range files {
		resp += strings.TrimPrefix(file, dir) + "\n"
	}

	tutil.RequireGoldenBytes(t, []byte(resp))
}
