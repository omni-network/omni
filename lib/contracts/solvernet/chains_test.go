package solvernet_test

import (
	"testing"

	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/stretchr/testify/require"
)

func TestHLChains(t *testing.T) {
	t.Parallel()

	require.NotPanics(t, func() {
		for _, network := range netconf.All() {
			_ = solvernet.HLChains(network)
		}
	})
}
