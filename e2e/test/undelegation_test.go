package e2e_test

import (
	"context"
	"math/big"
	"path/filepath"
	"testing"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/cosmos/cosmos-sdk/client/grpc/node"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

//nolint:paralleltest // We have to run tests sequentially
func TestUndelegations(t *testing.T) {
	t.Parallel()

	skipFunc := func(manifest types.Manifest) bool {
		return !manifest.AllE2ETests
	}
	maybeTestNetwork(t, skipFunc, func(ctx context.Context, t *testing.T, deps NetworkDeps) {
		t.Helper()

		network := deps.Network
		omniBackend, err := deps.OmniBackend()
		require.NoError(t, err)

		tmpDir := t.TempDir()

		cl, err := http.New(network.ID.Static().ConsensusRPC(), "/websocket")
		require.NoError(t, err)
		cprov := provider.NewABCI(cl, network.ID)

		const valChangeWait = 15 * time.Second

		validators, err := cprov.SDKValidators(ctx)
		require.NoError(t, err)
		require.Len(t, validators, 2)
		validator := validators[0]
		altValidator := validators[1]
		validatorAddr, err := validator.OperatorEthAddr()
		require.NoError(t, err)
		valPower, err := validator.Power()
		require.NoError(t, err)

		// delegator's keys
		delegatorPrivKey, delegatorEthAddr := GenFundedEOA(ctx, t, omniBackend)
		delegatorCosmosAddr := sdk.AccAddress(delegatorEthAddr.Bytes())
		delegatorPrivKeyFile := filepath.Join(tmpDir, "delegator_privkey")
		err = ethcrypto.SaveECDSA(delegatorPrivKeyFile, delegatorPrivKey)
		require.NoError(t, err)

		contract, err := bindings.NewStaking(common.HexToAddress(predeploys.Staking), omniBackend)
		require.NoError(t, err)

		const delegation = uint64(76)

		t.Run("delegation", func(t *testing.T) {
			txOpts, err := omniBackend.BindOpts(ctx, delegatorEthAddr)
			require.NoError(t, err)
			txOpts.Value = bi.Ether(delegation)

			tx, err := contract.Delegate(txOpts, validatorAddr)
			require.NoError(t, err)

			_, err = omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			// make sure the validator power is increased and the delegation can be found
			require.Eventuallyf(t, func() bool {
				val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
				require.True(t, ok)
				newPower, err := val.Power()
				require.NoError(t, err)

				if degelatedAmount(t, ctx, cprov, val.OperatorAddress, delegatorCosmosAddr.String()).IsZero() {
					return false
				}

				return newPower >= valPower+delegation
			}, valChangeWait, 500*time.Millisecond, "failed to delegate")
		})

		burnFee := bi.Ether(0.1)
		var anyBlock *big.Int

		t.Run("undelegate from a wrong validator", func(t *testing.T) {
			txOpts, err := omniBackend.BindOpts(ctx, delegatorEthAddr)
			require.NoError(t, err)
			txOpts.Value = burnFee

			balance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
			require.NoError(t, err)

			addr, err := altValidator.OperatorEthAddr()
			require.NoError(t, err)

			tx, err := contract.Undelegate(txOpts, addr, bi.N(delegation))
			require.NoError(t, err)

			_, err = omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			waitForBlocks(ctx, t, cprov, 10)

			require.Eventuallyf(t, func() bool {
				newBalance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
				require.NoError(t, err)

				// Amounts are roughly equal (we need to account for tx fee expected to be below burnFee)
				return bi.GT(bi.Add(balance, burnFee), newBalance)
			}, valChangeWait, 500*time.Millisecond, "failed to undeleate")
		})

		t.Run("undelegate from a non-existent validator", func(t *testing.T) {
			balance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
			require.NoError(t, err)

			txOpts, err := omniBackend.BindOpts(ctx, delegatorEthAddr)
			require.NoError(t, err)
			txOpts.Value = burnFee

			tx, err := contract.Undelegate(txOpts, tutil.RandomAddress(), bi.N(delegation))
			require.NoError(t, err)

			_, err = omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			require.Eventuallyf(t, func() bool {
				newBalance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
				require.NoError(t, err)

				// Amounts are roughly equal (we need to account for tx fee expected to be below burnFee)
				return bi.GT(bi.Add(balance, burnFee), newBalance)
			}, valChangeWait, 500*time.Millisecond, "failed to undeleate")
		})

		t.Run("undelegation too big", func(t *testing.T) {
			require.NoError(t, err)
			balance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
			require.NoError(t, err)

			txOpts, err := omniBackend.BindOpts(ctx, delegatorEthAddr)
			require.NoError(t, err)
			txOpts.Value = burnFee

			// undelegate more than delegated
			tooLargeAmt := bi.N(2 * delegation)
			tx, err := contract.Undelegate(txOpts, validatorAddr, tooLargeAmt)
			require.NoError(t, err)

			_, err = omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			waitForBlocks(ctx, t, cprov, 10)

			require.Eventuallyf(t, func() bool {
				newBalance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
				require.NoError(t, err)

				// Amounts are roughly equal (we need to account for tx fee expected to be below burnFee)
				return bi.GT(bi.Add(balance, burnFee), newBalance)
			}, valChangeWait, 500*time.Millisecond, "failed to undeleate")
		})

		t.Run("partial undelegation", func(t *testing.T) {
			require.NoError(t, err)
			balance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
			require.NoError(t, err)

			txOpts, err := omniBackend.BindOpts(ctx, delegatorEthAddr)
			require.NoError(t, err)
			txOpts.Value = burnFee

			// undelegate half
			undelegatedAmount := bi.N(delegation / 2)
			tx, err := contract.Undelegate(txOpts, validatorAddr, undelegatedAmount)
			require.NoError(t, err)

			_, err = omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			require.Eventuallyf(t, func() bool {
				newBalance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
				require.NoError(t, err)

				// we subtract the burn fee twice to account for the tx fees (which are expected to be below the burn fee)
				return bi.GTE(newBalance, bi.Add(bi.Sub(balance, burnFee, burnFee), undelegatedAmount))
			}, valChangeWait, 500*time.Millisecond, "failed to undeleate")

			// ensure rewards are still accruing
			waitForBlocks(ctx, t, cprov, 10)
			_, ok := queryDelegationRewards(t, ctx, cprov, delegatorCosmosAddr, validator.OperatorAddress)
			require.True(t, ok)
		})

		t.Run("complete undelegation", func(t *testing.T) {
			require.NoError(t, err)
			balance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
			require.NoError(t, err)

			txOpts, err := omniBackend.BindOpts(ctx, delegatorEthAddr)
			require.NoError(t, err)
			txOpts.Value = burnFee

			delegatedAmount := degelatedAmount(t, ctx, cprov, validator.OperatorAddress, delegatorCosmosAddr.String()).Amount.BigInt()

			// undelegate remaining half
			tx, err := contract.Undelegate(txOpts, validatorAddr, delegatedAmount)
			require.NoError(t, err)

			_, err = omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			require.Eventuallyf(t, func() bool {
				newBalance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
				require.NoError(t, err)

				// we subtract the burn fee twice to account for the tx fees (which are expected to be below the burn fee)
				return bi.GTE(newBalance, bi.Add(bi.Sub(balance, burnFee, burnFee), delegatedAmount))
			}, valChangeWait, 500*time.Millisecond, "failed to undeleate")

			// ensure nothing is staked anymore
			remainingAmt := degelatedAmount(t, ctx, cprov, validator.OperatorAddress, delegatorCosmosAddr.String())
			log.Debug(ctx, "remaining delegation", "amount", remainingAmt)
			require.Equal(t, int64(0), remainingAmt.Amount.Int64())

			// ensure no rewards accrue anymore
			waitForBlocks(ctx, t, cprov, 10)
			_, ok := queryDelegationRewards(t, ctx, cprov, delegatorCosmosAddr, validator.OperatorAddress)
			require.False(t, ok)
		})

		t.Run("Delegate and deploy a multiplicator contract", func(t *testing.T) {
			// Delegate again
			txOpts, err := omniBackend.BindOpts(ctx, delegatorEthAddr)
			require.NoError(t, err)
			txOpts.Value = bi.Ether(delegation)

			tx, err := contract.Delegate(txOpts, validatorAddr)
			require.NoError(t, err)

			_, err = omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			// make sure the validator power is increased and the delegation can be found
			require.Eventuallyf(t, func() bool {
				val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
				require.True(t, ok)
				newPower, err := val.Power()
				require.NoError(t, err)

				if degelatedAmount(t, ctx, cprov, val.OperatorAddress, delegatorCosmosAddr.String()).IsZero() {
					return false
				}

				return newPower >= valPower+delegation
			}, valChangeWait, 500*time.Millisecond, "failed to delegate")
		})
	})
}

// waitForBlocks wait for `n` blocks.
//
//nolint:unparam // we might want to change it from call to call
func waitForBlocks(ctx context.Context, t *testing.T, cprov provider.Provider, n uint64) {
	t.Helper()

	log.Debug(ctx, "awaiting", "blocks", n)

	s, err := cprov.QueryClients().Node.Status(ctx, &node.StatusRequest{})
	require.NoError(t, err)

	target := s.Height + n

	require.Eventually(t, func() bool {
		s, err := cprov.QueryClients().Node.Status(ctx, &node.StatusRequest{})
		require.NoError(t, err)

		return int64(s.Height) >= int64(target)
	}, time.Second*time.Duration(target*2), time.Millisecond*100)
}
