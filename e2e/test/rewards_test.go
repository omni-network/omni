package e2e_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/cchain/queryutil"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/rpc/client/http"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"
)

func TestInflation(t *testing.T) {
	t.Parallel()
	testNetwork(t, func(ctx context.Context, t *testing.T, network netconf.Network, endpoints xchain.RPCEndpoints) {
		t.Helper()

		cl, err := http.New(network.ID.Static().ConsensusRPC(), "/websocket")
		require.NoError(t, err)
		cprov := provider.NewABCI(cl, network.ID)

		inf, changed, err := queryutil.AvgInflationRate(ctx, cprov, 3)
		if changed {
			t.Log("staking state changed") // Avoids test flapping given delegation race
			return
		}
		require.NoError(t, err)

		target := math.LegacyNewDecWithPrec(11, 2) // 11%
		delta := math.LegacyNewDecWithPrec(1, 2)   // Allow +-1% error
		minInf, maxInf := target.Sub(delta), target.Add(delta)
		if inf.LT(minInf) || inf.GT(maxInf) {
			require.Fail(t, "inflation average not within bounds", "rate: %v, min: %v, max: %v", inf, minInf, maxInf)
		}
	})
}
