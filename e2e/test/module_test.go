package e2e_test

import (
	"context"
	"testing"

	magellan2 "github.com/omni-network/omni/halo/app/upgrades/magellan"
	uluwatu1 "github.com/omni-network/omni/halo/app/upgrades/uluwatu"
	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/rpc/client/http"

	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/stretchr/testify/require"
)

func TestMint(t *testing.T) {
	t.Parallel()
	testNetwork(t, func(ctx context.Context, t *testing.T, network netconf.Network, endpoints xchain.RPCEndpoints) {
		t.Helper()

		cl, err := http.New(network.ID.Static().ConsensusRPC(), "/websocket")
		require.NoError(t, err)
		cprov := provider.NewABCI(cl, network.ID)

		paramResponse, err := cprov.QueryClients().Mint.Params(ctx, &minttypes.QueryParamsRequest{})
		require.NoError(t, err)
		require.Equal(t, magellan2.MintParams.String(), paramResponse.Params.String())

		inflResponse, err := cprov.QueryClients().Mint.Inflation(ctx, &minttypes.QueryInflationRequest{})
		require.NoError(t, err)
		require.Equal(t, magellan2.MintParams.InflationMin.String(), inflResponse.Inflation.String())
	})
}

func TestSlashing(t *testing.T) {
	t.Parallel()
	testNetwork(t, func(ctx context.Context, t *testing.T, network netconf.Network, endpoints xchain.RPCEndpoints) {
		t.Helper()

		cl, err := http.New(network.ID.Static().ConsensusRPC(), "/websocket")
		require.NoError(t, err)
		cprov := provider.NewABCI(cl, network.ID)

		paramResponse, err := cprov.QueryClients().Slashing.Params(ctx, &slashingtypes.QueryParamsRequest{})
		require.NoError(t, err)
		require.Equal(t, uluwatu1.SlashingParams.String(), paramResponse.Params.String())
	})
}
