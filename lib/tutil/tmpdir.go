package tutil

import (
	"crypto/rand"
	"os"
	"testing"

	"github.com/omni-network/omni/lib/bi"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/stretchr/testify/require"
)

// TempDir creates a temporary directory with a random name, as opposed to t.TempDir which
// uses consecutive numbers per test (001, 002 ...).
//
//nolint:usetesting // Explicitly not using t.TempDir() to avoid consecutive numbers.
func TempDir(t *testing.T) string {
	t.Helper()

	dir, err := os.MkdirTemp(t.TempDir(), randStr(t))
	require.NoError(t, err)

	return dir
}

func randStr(t *testing.T) string {
	t.Helper()

	n, err := rand.Int(rand.Reader, bi.Ether(1))
	require.NoError(t, err)

	return hexutil.EncodeBig(n)
}
