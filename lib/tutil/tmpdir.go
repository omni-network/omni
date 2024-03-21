package tutil

import (
	"crypto/rand"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/stretchr/testify/require"
)

// TempDir creates a temporary directory with a random name, as opposed to t.TempDir which
// uses consecutive numbers per test (001, 002 ...).
func TempDir(t *testing.T) string {
	t.Helper()

	dir, err := os.MkdirTemp(t.TempDir(), randStr(t))
	require.NoError(t, err)

	return dir
}

func randStr(t *testing.T) string {
	t.Helper()

	n, err := rand.Int(rand.Reader, big.NewInt(1e18))
	require.NoError(t, err)

	return hexutil.EncodeBig(n)
}
