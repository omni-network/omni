package e2e_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/halo/evmredenom"
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
// - after a rewards withdrawal, the ETH balance increases,
// - no balance increase if withdrawals are triggered without any delegations,
// - no balance increase if delegations exist, but the validator address is wrong.
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
		// Skip if not enough validators
		if len(validators) < 2 {
			t.Skip()
		}
		var validator cchain.SDKValidator
		var valIndex int
		for i, v := range validators {
			if !v.Jailed {
				validator = v
				valIndex = i

				break
			}
		}
		validatorAddr, err := validator.OperatorEthAddr()
		require.NoError(t, err)
		// for some tests, we need a valid alternative validator address
		altValidator := validators[(valIndex+1)%len(validators)]

		var anyBlock *big.Int
		delegation := bi.Ether(5 * factor)
		distrContractAddr := common.HexToAddress(predeploys.Distribution)
		dContract, err := bindings.NewDistribution(distrContractAddr, omniBackend)
		require.NoError(t, err)

		stakingContractAddr := common.HexToAddress(predeploys.Staking)
		stContract, err := bindings.NewStaking(stakingContractAddr, omniBackend)
		require.NoError(t, err)

		t.Run("don't delegate and withdraw rewards", func(t *testing.T) {
			t.Parallel()

			_, delegatorEthAddr := GenFundedEOA(ctx, t, omniBackend)
			delegatorCosmosAddr := sdk.AccAddress(delegatorEthAddr.Bytes())

			val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
			require.True(t, ok)
			require.NoError(t, err)

			txOpts, err := omniBackend.BindOpts(ctx, delegatorEthAddr)
			require.NoError(t, err)
			txOpts.Value = delegation

			// make sure no rewards are pending
			_, ok = queryDelegationRewards(t, ctx, cprov, delegatorCosmosAddr, val.OperatorAddress)
			require.False(t, ok)

			// try to withdraw rewards
			tx, err := dContract.Withdraw(txOpts, validatorAddr)
			require.NoError(t, err)
			receipt, err := omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			// fetch balance at the block when withdrawal request was mined
			balanceBeforeWithdrawal, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, receipt.BlockNumber)
			require.NoError(t, err)

			// Wait and make sure the balance stays the same
			waitForBlocks(ctx, t, cprov, 3)

			balanceAfterWithdrawal, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
			require.NoError(t, err)

			require.Equal(t, balanceBeforeWithdrawal, balanceAfterWithdrawal)
		})

		t.Run("delegate and withdraw rewards from a wrong validator", func(t *testing.T) {
			t.Parallel()

			_, delegatorEthAddr := GenFundedEOA(ctx, t, omniBackend)
			delegatorCosmosAddr := sdk.AccAddress(delegatorEthAddr.Bytes())

			val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
			require.True(t, ok)
			require.NoError(t, err)

			txOpts, err := omniBackend.BindOpts(ctx, delegatorEthAddr)
			require.NoError(t, err)
			txOpts.Value = delegation

			tx, err := stContract.Delegate(txOpts, validatorAddr)
			require.NoError(t, err)

			_, err = omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			// make sure the delegation can be found
			require.Eventuallyf(t, func() bool {
				stake := delegatedStake(t, ctx, cprov, val.OperatorAddress, delegatorCosmosAddr.String())
				actual, err := evmredenom.ToEVMAmount(stake) // actual = stake * factor
				require.NoError(t, err)

				return bi.EQ(delegation, actual)
			}, valChangeWait, 500*time.Millisecond, "failed to delegate")

			// Wait for rewards to accrue
			waitForBlocks(ctx, t, cprov, 3)
			_, ok = queryDelegationRewards(t, ctx, cprov, delegatorCosmosAddr, val.OperatorAddress)
			require.True(t, ok)

			// withdraw rewards from the wrong validator
			altValAddr, err := altValidator.OperatorEthAddr()
			require.NoError(t, err)
			require.NotEqual(t, validatorAddr, altValAddr)
			tx, err = dContract.Withdraw(txOpts, altValAddr)
			require.NoError(t, err)
			receipt, err := omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			// fetch balance at the block when withdrawal request was mined
			balanceBeforeWithdrawal, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, receipt.BlockNumber)
			require.NoError(t, err)

			// Wait and make sure the balance stays the same
			waitForBlocks(ctx, t, cprov, 3)

			balanceAfterWithdrawal, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
			require.NoError(t, err)

			require.Equal(t, balanceBeforeWithdrawal, balanceAfterWithdrawal)
		})

		t.Run("delegate and withdraw rewards", func(t *testing.T) {
			t.Parallel()

			_, delegatorEthAddr := GenFundedEOA(ctx, t, omniBackend)
			delegatorCosmosAddr := sdk.AccAddress(delegatorEthAddr.Bytes())

			val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
			require.True(t, ok)
			require.NoError(t, err)

			txOpts, err := omniBackend.BindOpts(ctx, delegatorEthAddr)
			require.NoError(t, err)
			txOpts.Value = delegation

			tx, err := stContract.Delegate(txOpts, validatorAddr)
			require.NoError(t, err)

			_, err = omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			// make sure the delegation can be found
			require.Eventuallyf(t, func() bool {
				stake := delegatedStake(t, ctx, cprov, val.OperatorAddress, delegatorCosmosAddr.String())
				actual, err := evmredenom.ToEVMAmount(stake) // actual = stake * factor
				require.NoError(t, err)

				return bi.EQ(delegation, actual)
			}, valChangeWait, 500*time.Millisecond, "failed to delegate")

			// Wait for rewards to accrue
			waitForBlocks(ctx, t, cprov, 3)
			_, ok = queryDelegationRewards(t, ctx, cprov, delegatorCosmosAddr, val.OperatorAddress)
			require.True(t, ok)

			// withdraw rewards
			tx, err = dContract.Withdraw(txOpts, validatorAddr)
			require.NoError(t, err)
			receipt, err := omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			// fetch balance at the block when withdrawal request was mined
			balanceBeforeWithdrawal, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, receipt.BlockNumber)
			require.NoError(t, err)

			// Make sure the ETH balance increases eventually
			require.Eventuallyf(t, func() bool {
				balanceAfterWithdrawal, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
				require.NoError(t, err)

				return bi.GT(balanceAfterWithdrawal, balanceBeforeWithdrawal)
			}, valChangeWait, 500*time.Millisecond, "failed to withdraw")
		})
	})
}
