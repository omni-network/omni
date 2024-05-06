package e2e_test

import (
	"context"
	"encoding/hex"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/netconf"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	"github.com/stretchr/testify/require"
)

// Tests that the app hash (as reported by the app) matches the last
// block and the node sync status.
func TestApp_Hash(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	testNode(t, func(t *testing.T, _ netconf.Network, node *e2e.Node, _ []Portal) {
		t.Helper()
		client, err := node.Client()
		require.NoError(t, err)

		info, err := client.ABCIInfo(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, info.Response.LastBlockAppHash, "expected app to return app hash")

		// In next-block execution, the app hash is stored in the next block
		requestedHeight := info.Response.LastBlockHeight + 1

		require.Eventually(t, func() bool {
			status, err := client.Status(ctx)
			require.NoError(t, err)
			require.NotZero(t, status.SyncInfo.LatestBlockHeight)

			return status.SyncInfo.LatestBlockHeight >= requestedHeight
		}, 5*time.Second, 500*time.Millisecond)

		block, err := client.Block(ctx, &requestedHeight)
		require.NoError(t, err)
		require.Equal(t,
			hex.EncodeToString(info.Response.LastBlockAppHash),
			hex.EncodeToString(block.Block.AppHash.Bytes()),
			"app hash does not match last block's app hash")
	})
}
