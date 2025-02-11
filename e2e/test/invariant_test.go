package e2e_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/feature"
	"github.com/omni-network/omni/lib/netconf"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
)

// TestInvariantNoUserBalance asserts the invariant that only the staking module has balance.
func TestInvariantNoUserBalance(t *testing.T) {
	t.Parallel()
	testNode(t, func(t *testing.T, network netconf.Network, node *e2e.Node, portals []Portal) {
		t.Helper()

		ctx := context.Background()
		if !feature.FlagEVMStakingModule.Enabled(ctx) {
			t.Skip("EVM staking module not enabled")
		}

		client, err := node.Client()
		require.NoError(t, err)
		cprov := provider.NewABCI(client, network.ID)

		resp, err := cprov.QueryClients().Bank.DenomOwners(ctx, &banktypes.QueryDenomOwnersRequest{
			Denom: sdk.DefaultBondDenom,
		})
		require.NoError(t, err)

		// Only allow staking module to have balance
		expected := map[common.Address]bool{
			common.BytesToAddress(address.Module(stakingtypes.BondedPoolName)):    true,
			common.BytesToAddress(address.Module(stakingtypes.NotBondedPoolName)): true,
			common.BytesToAddress(address.Module(distrtypes.ModuleName)):          true,
		}

		for _, owner := range resp.DenomOwners {
			sdkAddr, err := sdk.AccAddressFromBech32(owner.Address)
			require.NoError(t, err)
			ethAddr := common.BytesToAddress(sdkAddr.Bytes())

			if expected[ethAddr] {
				continue
			}

			require.Fail(t, "unexpected balance", "address: %v, balance: %v", ethAddr, owner.Balance)
		}
	})
}
