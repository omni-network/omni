package provider_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/halo/attest"
	"github.com/omni-network/omni/halo/comet"
	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/engine"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/test/tutil"

	rpclocal "github.com/cometbft/cometbft/rpc/client/local"
	rpctest "github.com/cometbft/cometbft/rpc/test"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestQueryApprovedFrom(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	f := fuzz.New().NilChance(0).NumElements(2, 8)
	tutil.PrepRPCTestConfig(t) // Write genesis and priv validator files to temp dir.

	state, err := comet.LoadOrGenState(t.TempDir(), 0, map[uint64]string{})
	require.NoError(t, err)

	ethCl, err := engine.NewMock()
	require.NoError(t, err)

	app := comet.NewApp(ethCl, attest.Service(nil), state, nil, 0)
	node := rpctest.NewTendermint(app, new(rpctest.Options))

	fetcher := provider.NewABCIFetcher(rpclocal.New(node))

	assert := func(t *testing.T, chainID uint64, from uint64, total int) {
		t.Helper()

		actual, err := fetcher.ApprovedFrom(ctx, chainID, from)
		require.NoError(t, err)

		require.Len(t, actual, total)
		for i, agg := range actual {
			require.Equal(t, chainID, agg.SourceChainID)
			require.Equal(t, from+uint64(i), agg.BlockHeight)
		}
	}

	// Insert some red herrings
	state.AddAttestations(newAgg(f, 98, 1))
	state.AddAttestations(newAgg(f, 99, 2))
	state.AddAttestations(newAgg(f, 99, 3))

	const chain = 1
	assert(t, chain, 0, 0) // No aggregates yet.

	// Insert 3 aggregates
	state.AddAttestations(newAgg(f, chain, 1))
	state.AddAttestations(newAgg(f, chain, 2))
	state.AddAttestations(newAgg(f, chain, 3))

	assert(t, chain, 0, 0) // From 0 == 0 aggregates, since 0 doesn't exist
	assert(t, chain, 1, 3) // From 1 == 3 aggregates.
	assert(t, chain, 2, 2) // From 2 == 2 aggregates.
	assert(t, chain, 3, 1) // From 3 == 1 aggregates.
	assert(t, chain, 4, 0) // From 4 == 0 aggregates, since 4 doesn't exist

	// Insert 5 and 6 (skipping 4)
	state.AddAttestations(newAgg(f, chain, 5))
	state.AddAttestations(newAgg(f, chain, 6))

	assert(t, chain, 1, 3) // From 1 == 3 aggregates, since 4 doesn't exist
	assert(t, chain, 4, 0) // From 4 == 0 aggregates, since 4 doesn't exist
	assert(t, chain, 5, 2) // From 5 == 2 aggregates
	assert(t, chain, 6, 1) // From 6 == 1 aggregates
}

func newAgg(f *fuzz.Fuzzer, chainID uint64, height uint64) []xchain.AggAttestation {
	var resp xchain.AggAttestation
	f.Fuzz(&resp)
	resp.BlockHeight = height
	resp.SourceChainID = chainID

	return []xchain.AggAttestation{resp}
}
