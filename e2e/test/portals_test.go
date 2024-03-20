package e2e_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestPortalOffsets ensures that cross chain messages are sent from all source chains to all destination chains
// and that at least half of the messages are received by the destination chains.
func TestPortalOffsets(t *testing.T) {
	t.Parallel()
	testPortal(t, func(t *testing.T, source Portal, dests []Portal) {
		t.Helper()
		for _, dest := range dests {
			if source.Chain.ID == dest.Chain.ID {
				continue
			}

			sourceOffset, err := source.Contract.OutXStreamOffset(nil, dest.Chain.ID)
			require.NoError(t, err)

			destOffset, err := dest.Contract.InXStreamOffset(nil, source.Chain.ID)
			require.NoError(t, err)

			// require at least some xmsgs were sent
			require.Greater(t, sourceOffset, uint64(0),
				"no xmsgs sent from source chain %v to dest chain %v",
				source.Chain.ID, dest.Chain.ID)

			// require at least half were received
			require.GreaterOrEqual(t, destOffset, sourceOffset/2,
				"dest chain %v offset=%d, source chain %v offset=%d",
				dest.Chain.ID, destOffset, source.Chain.ID, sourceOffset)
		}
	})
}
