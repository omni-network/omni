package e2e_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/xchain"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	"github.com/stretchr/testify/require"
)

// TestApprovedAttestations tests that all halo instances contain approved aggregate attestation
// for at least half of all the source chain blocks.
func TestApprovedAttestations(t *testing.T) {
	t.Parallel()
	test(t, func(t *testing.T, node e2e.Node, portals []Portal) {
		t.Helper()
		client, err := node.Client()
		require.NoError(t, err)
		cprov := provider.NewABCIProvider(client, nil)

		ctx := context.Background()
		for _, portal := range portals {
			height, err := portal.Client.BlockNumber(ctx)
			require.NoError(t, err)

			aggs, err := fetchAllAggs(ctx, cprov, portal.Chain.ID, portal.Chain.DeployHeight)
			require.NoError(t, err)

			totalBlocks := height - portal.Chain.DeployHeight
			require.GreaterOrEqual(t, len(aggs), int(totalBlocks/2)) // Assert that at least half of the blocks are approved
		}
	}, nil)
}

func fetchAllAggs(ctx context.Context, cprov cchain.Provider, chainID, from uint64) ([]xchain.AggAttestation, error) {
	var resp []xchain.AggAttestation
	for {
		aggs, err := cprov.ApprovedFrom(ctx, chainID, from)
		if err != nil {
			return nil, err
		}
		if len(aggs) == 0 { // No more attestation to fetch
			break
		}
		resp = append(resp, aggs...)

		// Update the from height to fetch the next batch of attestation
		from = aggs[len(aggs)-1].BlockHeight + 1
	}

	return resp, nil
}
