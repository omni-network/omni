package e2e_test

import (
	"context"
	"testing"

	magellan2 "github.com/omni-network/omni/halo/app/upgrades/magellan"
	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/rpc/client/http"

	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"
)

func TestInflation(t *testing.T) {
	t.Parallel()
	testNetwork(t, func(ctx context.Context, t *testing.T, network netconf.Network, endpoints xchain.RPCEndpoints) {
		t.Helper()

		cl, err := http.New(network.ID.Static().ConsensusRPC(), "/websocket")
		require.NoError(t, err)
		cprov := provider.NewABCI(cl, network.ID)

		paramResponse, err := cprov.QueryClients().Mint.Params(ctx, &minttypes.QueryParamsRequest{})
		require.NoError(t, err)
		proto.Equal(&magellan2.MintParams, &paramResponse.Params)
	})
}
