package allocs_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/omni-network/omni/contracts/allocs"
	"github.com/omni-network/omni/halo/genutil/evm"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"

	"github.com/stretchr/testify/require"
)

// TestIsLocked ensures that if a persistent network has it's genesis
// set in lib/netconf/static.go, then the allocs are locked in contracts/allocs.
func TestIsLocked(t *testing.T) {
	t.Parallel()

	for _, network := range netconf.All() {
		if network.IsEphemeral() {
			continue
		}

		if network.Static().ExecutionGenesisJSON == nil {
			continue
		}

		require.True(t, allocs.IsLocked(network), "allocs should be locked for %s", network)

		locked := allocs.MustAlloc(network)

		var genesis core.Genesis
		err := json.Unmarshal(network.Static().ExecutionGenesisJSON, &genesis)
		require.NoError(t, err)

		precompile := evm.PrecompilesAlloc()
		prefund, err := evm.PrefundAlloc(network)
		require.NoError(t, err)

		for addr, acc := range genesis.Alloc {
			if _, ok := precompile[addr]; ok {
				continue
			}

			if _, ok := prefund[addr]; ok {
				continue
			}

			lockedAcc, ok := locked[addr]
			require.True(t, ok, "missing locked alloc for %s", addr)
			require.Equal(t, hexBytes(acc.Code), hexBytes(lockedAcc.Code), "code mismatch for %s", addr)
			require.Equal(t, hexBig(acc.Balance), hexBig(lockedAcc.Balance), "balance mismatch for %s", addr)
			require.Equal(t, acc.Nonce, lockedAcc.Nonce, "nonce mismatch for %s", addr)

			for key, val := range acc.Storage {
				lockedVal, ok := lockedAcc.Storage[key]
				require.True(t, ok, "missing storage key %s for %s", key, addr)
				require.Equal(t, hexBytes(val[:]), hexBytes(lockedVal[:]), "storage mismatch for %s[%s]", addr, key)
			}
		}
	}
}

func hexBytes(b []byte) string { return hexutil.Encode(b) }
func hexBig(b *big.Int) string { return hexutil.EncodeBig(b) }
