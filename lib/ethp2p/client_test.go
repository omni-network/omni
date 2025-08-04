package ethp2p_test

import (
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/omni-network/omni/lib/ethp2p"

	etypes "github.com/ethereum/go-ethereum/core/types"
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

	headers, err := cl.HeadersDownFrom(ctx, status.Head, 1)
	require.NoError(t, err)
	require.NotEmpty(t, headers)

	n, err := cl.AllAccountRanges(ctx, status.Head, 64*1024) // 64kB
	require.NoError(t, err)
	t.Logf("Snapshot ranges:%d", len(n))

	var accounts int
	var all []int
	var mini, maxi, sum int
	for _, r := range n {
		var nonZeros int
		accounts += len(r.Accounts)
		for _, v := range r.Accounts {
			acc, err := etypes.FullAccount(v.Body)
			require.NoError(t, err)
			if acc.Balance.IsZero() {
				continue
			}
			nonZeros++
		}
		all = append(all, nonZeros)
		sum += nonZeros
		if nonZeros > maxi {
			maxi = nonZeros
		}
		if mini == 0 || nonZeros < mini {
			mini = nonZeros
		}
	}

	t.Logf("Accounts: total=%d, nonZeros=%d, min=%d, max=%d, avg=%.0f", accounts, sum, mini, maxi, float64(sum)/float64(len(all)))

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
