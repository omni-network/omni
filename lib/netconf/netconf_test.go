package netconf_test

import (
	"testing"

	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/stretchr/testify/require"
)

// TestStrats ensures the netconf.StratX matches ethclient.HeadX.
// Netconf shouldn't import ethclient, so using this test to keep in-sync.
func TestStrats(t *testing.T) {
	t.Parallel()

	require.EqualValues(t, ethclient.HeadLatest, netconf.StratLatest)
	require.EqualValues(t, ethclient.HeadFinalized, netconf.StratFinalized)
}
