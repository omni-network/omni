package ethp2p_test

import (
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/omni-network/omni/lib/ethp2p"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/enode"

	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "Run integration tests")

func TestP2PClient(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("Skipping integration test")
	}

	peer := executionSeedENR(t, "../netconf/mainnet/execution-seeds.txt")

	key, err := crypto.GenerateKey()
	require.NoError(t, err)

	ctx := t.Context()
	cl, status, err := ethp2p.Dial(ctx, key, peer)
	require.NoError(t, err)

	headers, err := cl.HeadersDownFrom(ctx, status.Head, 100)
	require.NoError(t, err)
	require.NotEmpty(t, headers)

	n, err := cl.SnapshotRange(ctx, status.Head, 256)
	require.NoError(t, err)
	t.Logf("Snapshot range: %d", n)

	err = cl.Disconnect()
	require.NoError(t, err)
}

// executionSeedENR returns an enode by parsing the first ENR from the provided path.
func executionSeedENR(t *testing.T, path string) *enode.Node {
	t.Helper()
	bz, err := os.ReadFile(path)
	require.NoError(t, err)

	line := strings.TrimSpace(strings.Split(string(bz), "\n")[0])
	n, err := enode.ParseV4(line)
	require.NoError(t, err)

	n, err = ethp2p.DNSResolveHostname(t.Context(), n)
	require.NoError(t, err)

	return n
}
