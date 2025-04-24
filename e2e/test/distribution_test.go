package e2e_test

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/cchain/provider"

	"github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	_ "embed"
)

// TestDistribution tests:
// - that after withdrawals, the queried rewards are smaller.
func TestDistribution(t *testing.T) {
	t.Parallel()

	skipFunc := func(manifest types.Manifest) bool {
		return !manifest.AllE2ETests
	}

	maybeTestNetwork(t, skipFunc, func(ctx context.Context, t *testing.T, deps NetworkDeps) {
		t.Helper()

		network := deps.Network
		omniBackend, err := deps.OmniBackend()
		require.NoError(t, err)

		cl, err := http.New(network.ID.Static().ConsensusRPC(), "/websocket")
		require.NoError(t, err)
		cprov := provider.NewABCI(cl, network.ID)

		const valChangeWait = 30 * time.Second

		validators, err := cprov.SDKValidators(ctx)
		require.NoError(t, err)
		var validator cchain.SDKValidator
		for _, v := range validators {
			if !v.Jailed {
				validator = v
				break
			}
		}
		validatorAddr, err := validator.OperatorEthAddr()
		require.NoError(t, err)

		t.Run("delegate and withdraw rewards", func(t *testing.T) {
			t.Parallel()

			delegation := bi.Ether(50)

			_, delegatorEthAddr := GenFundedEOA(ctx, t, omniBackend)
			delegatorCosmosAddr := sdk.AccAddress(delegatorEthAddr.Bytes())

			val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
			require.True(t, ok)
			require.NoError(t, err)

			txOpts, err := omniBackend.BindOpts(ctx, delegatorEthAddr)
			require.NoError(t, err)
			txOpts.Value = delegation

			stakingContractAddr := common.HexToAddress(predeploys.Staking)
			contract, err := bindings.NewStaking(stakingContractAddr, omniBackend)
			require.NoError(t, err)
			tx, err := contract.Delegate(txOpts, validatorAddr)
			require.NoError(t, err)

			_, err = omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			// make sure the delegation can be found
			require.Eventuallyf(t, func() bool {
				delegatedAmt := delegatedAmount(t, ctx, cprov, val.OperatorAddress, delegatorCosmosAddr.String()).Amount.BigInt()
				return bi.EQ(delegatedAmt, delegation)
			}, valChangeWait, 500*time.Millisecond, "failed to delegate")

			waitForBlocks(ctx, t, cprov, 10)
			rewardsBeforeWithdrawal, ok := queryDelegationRewards(t, ctx, cprov, delegatorCosmosAddr, val.OperatorAddress)
			require.True(t, ok)

			// withdraw rewards
			distrContractAddr := common.HexToAddress(predeploys.Distribution)
			dContract, err := bindings.NewDistribution(distrContractAddr, omniBackend)
			require.NoError(t, err)
			_, err = dContract.Withdraw(txOpts, validatorAddr)
			require.NoError(t, err)

			// Make sure the rewards balance decreases eventually
			require.Eventuallyf(t, func() bool {
				rewardsAfterWithdrawal, ok := queryDelegationRewards(t, ctx, cprov, delegatorCosmosAddr, val.OperatorAddress)

				return ok && rewardsBeforeWithdrawal.GT(rewardsAfterWithdrawal)
			}, valChangeWait, 500*time.Millisecond, "failed to withdraw")
		})
	})
}
