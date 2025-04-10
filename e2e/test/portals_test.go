package e2e_test

import (
	"testing"
	"time"

	"github.com/omni-network/omni/lib/netconf"

	"github.com/stretchr/testify/require"
)

// TestPortalOffsets ensures that cross chain messages are sent from all source chains to all destination chains
// and that at least half of the messages are received by the destination chains.
func TestPortalOffsets(t *testing.T) {
	t.Parallel()
	testPortal(t, func(t *testing.T, network netconf.Network, source Portal, dests []Portal) {
		t.Helper()
		for _, dest := range dests {
			for _, stream := range network.StreamsBetween(source.Chain.ID, dest.Chain.ID) {
				// Require some messages were sent
				require.Eventuallyf(t, func() bool {
					sourceOffset, err := source.Contract.OutXMsgOffset(nil, dest.Chain.ID, uint64(stream.ShardID))
					return err == nil && sourceOffset > 0
				}, time.Second*30, time.Second, "no xmsgs sent from source chain %v to dest chain %v", source.Chain.ID, dest.Chain.ID)

				// Require some messages were received
				require.Eventuallyf(t, func() bool {
					destOffset, err := dest.Contract.InXMsgOffset(nil, source.Chain.ID, uint64(stream.ShardID))
					return err == nil && destOffset > 0
				}, time.Second*30, time.Second, "no xmsgs received by dest chain %v from source chain %v", dest.Chain.ID, source.Chain.ID)
			}
		}
	})
}

// TestSupportedChains ensures that all portals have been relayed supported chains from the PortalRegistry, via the XRegistry.
func TestSupportedChains(t *testing.T) {
	// TODO: enable when cchain setNetwork xmsgs are enabled
	t.Skip()

	t.Parallel()
	testPortal(t, func(t *testing.T, network netconf.Network, source Portal, dests []Portal) {
		t.Helper()
		for _, dest := range dests {
			supported, err := source.Contract.IsSupportedDest(nil, dest.Chain.ID)
			require.NoError(t, err)

			if source.Chain.ID == dest.Chain.ID {
				require.False(t, supported,
					"source chain %v supports itself", source.Chain.ID)
			} else {
				require.True(t, supported,
					"source chain %v does not support dest chain %v",
					source.Chain.ID, dest.Chain.ID)
			}
		}
	})
}
