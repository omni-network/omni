package e2e_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/cchain/provider"

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
		cprov := provider.NewABCIProvider(client)

		ctx := context.Background()
		for _, portal := range portals {
			height, err := portal.Client.BlockNumber(ctx)
			require.NoError(t, err)

			totalBlocks := height - portal.Chain.DeployHeight

			aggs, err := cprov.ApprovedFrom(ctx, portal.Chain.ID, portal.Chain.DeployHeight)
			require.NoError(t, err)

			require.GreaterOrEqual(t, len(aggs), int(totalBlocks/2)) // Assert that at least half of the blocks are approved
		}
	}, nil)
}
