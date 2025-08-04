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
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/common"

	"github.com/cosmos/cosmos-sdk/client/grpc/node"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	_ "embed"
)

const (
	StakingMethodDelegate   = 0
	StakingMethodUndelegate = 1
)

// burnFee is the fee burned on every undelegation.
var burnFee = bi.Ether(0.1)

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

		stakingContractAddr := common.HexToAddress(predeploys.Staking)
		contract, err := bindings.NewStaking(stakingContractAddr, omniBackend)
		require.NoError(t, err)

		// We deploy a proxy smart contract and let it stake and unstake in one batch
		t.Run("batched delegations and undelegations", func(t *testing.T) {
			t.Parallel()
			_, eoa := GenFundedEOA(ctx, t, omniBackend)

			proxyAddr, err := deploy(ctx, omniBackend, stakingContractAddr, eoa)
			proxyCosmosAddr := sdk.AccAddress(proxyAddr.Bytes())
			require.NoError(t, err)
			log.Debug(ctx, "Staking proxy deployed", "address", proxyAddr)

			delegation := bi.Ether(6 * factor)
			undelegation1 := bi.Ether(3 * factor)
			undelegation2 := bi.Ether(2 * factor)

			val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
			require.True(t, ok)

			require.NoError(t, err)

			// These calls will be mined in one block and batched by octane
			calls := []bindings.StakingProxyCall{
				{
					Method:    StakingMethodDelegate,
					Validator: validatorAddr,
					Value:     delegation,
					Amount:    bi.N(0),
				},
				{
					Method:    StakingMethodUndelegate,
					Validator: validatorAddr,
					Value:     burnFee,
					Amount:    undelegation1,
				},
				{
					Method:    StakingMethodUndelegate,
					Validator: validatorAddr,
					Value:     burnFee,
					Amount:    undelegation2,
				},
			}

			err = proxyCall(ctx, omniBackend, proxyAddr, eoa, calls)
			require.NoError(t, err)

			// make sure our delegated amount is delegation-undelegation1-undelegation2
			require.Eventuallyf(t, func() bool {
				stake := delegatedStake(t, ctx, cprov, val.OperatorAddress, proxyCosmosAddr.String())
				log.Debug(ctx, "Delegated amount", "stake", stake)

				expected := bi.Sub(delegation, undelegation1, undelegation2)
				actual, err := evmredenom.ToEVMAmount(stake) // actual=stake*factor
				require.NoError(t, err)

				return bi.EQ(expected, actual)
			}, valChangeWait, 500*time.Millisecond, "failed to execute batched staking calls")
		})

		const delegation = uint64(5 * factor)

		var anyBlock *big.Int

		t.Run("undelegate from a wrong validator", func(t *testing.T) {
			t.Parallel()
			_, delegatorEthAddr := GenFundedEOA(ctx, t, omniBackend)

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

			// Since we trigger an invalid undelegation, wait for a few blocks to make sure
			// the chain does not stall and our balance didn't change
			waitForBlocks(ctx, t, cprov, 2)

			require.Eventuallyf(t, func() bool {
				newBalance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
				require.NoError(t, err)

				// Amounts are roughly equal (we need to account for tx fee expected to be below burnFee)
				maxExpectedAmount := bi.Add(balance, burnFee)

				return bi.GT(maxExpectedAmount, newBalance)
			}, valChangeWait, 500*time.Millisecond, "failed to undeleate")
		})

		t.Run("undelegate from a non-existent validator", func(t *testing.T) {
			t.Parallel()
			_, delegatorEthAddr := GenFundedEOA(ctx, t, omniBackend)

			balance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
			require.NoError(t, err)

			txOpts, err := omniBackend.BindOpts(ctx, delegatorEthAddr)
			require.NoError(t, err)
			txOpts.Value = burnFee

			tx, err := contract.Undelegate(txOpts, tutil.RandomAddress(), bi.N(delegation))
			require.NoError(t, err)

			_, err = omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			// Since we trigger an invalid undelegation, wait for a few blocks to make sure
			// the chain does not stall and our balance didn't change
			waitForBlocks(ctx, t, cprov, 2)

			require.Eventuallyf(t, func() bool {
				newBalance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
				require.NoError(t, err)

				// Amounts are roughly equal (we need to account for tx fee expected to be below burnFee)
				maxExpectedAmount := bi.Add(balance, burnFee)

				return bi.GT(maxExpectedAmount, newBalance)
			}, valChangeWait, 500*time.Millisecond, "failed to undeleate")
		})

		t.Run("undelegation too big", func(t *testing.T) {
			t.Parallel()
			_, delegatorEthAddr := GenFundedEOA(ctx, t, omniBackend)

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

			// Since we trigger an invalid undelegation, wait for a few blocks to make sure
			// the chain does not stall and our balance didn't change
			waitForBlocks(ctx, t, cprov, 2)

			require.Eventuallyf(t, func() bool {
				newBalance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
				require.NoError(t, err)

				// Amounts are roughly equal (we need to account for tx fee expected to be below burnFee)
				maxExpectedAmount := bi.Add(balance, burnFee)

				return bi.GT(maxExpectedAmount, newBalance)
			}, valChangeWait, 500*time.Millisecond, "failed to undeleate")
		})

		t.Run("delegate, then undelegate (partially and completely)", func(t *testing.T) {
			t.Parallel()
			_, delegatorEthAddr := GenFundedEOA(ctx, t, omniBackend)
			delegatorCosmosAddr := sdk.AccAddress(delegatorEthAddr.Bytes())

			val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
			require.True(t, ok)
			require.NoError(t, err)

			txOpts, err := omniBackend.BindOpts(ctx, delegatorEthAddr)
			require.NoError(t, err)
			txOpts.Value = bi.Ether(delegation)

			tx, err := contract.Delegate(txOpts, validatorAddr)
			require.NoError(t, err)

			_, err = omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			// make sure the delegation can be found
			require.Eventuallyf(t, func() bool {
				stake := delegatedStake(t, ctx, cprov, val.OperatorAddress, delegatorCosmosAddr.String())
				actual, err := evmredenom.ToEVMAmount(stake) // actual=stake*factor
				require.NoError(t, err)

				return bi.EQ(actual, bi.Ether(delegation))
			}, valChangeWait, 500*time.Millisecond, "failed to delegate")

			require.NoError(t, err)
			balance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
			require.NoError(t, err)

			txOpts, err = omniBackend.BindOpts(ctx, delegatorEthAddr)
			require.NoError(t, err)
			txOpts.Value = burnFee

			// undelegate half of stake first
			undelegatedAmount := bi.Ether(delegation / 2)
			tx, err = contract.Undelegate(txOpts, validatorAddr, undelegatedAmount)
			require.NoError(t, err)

			_, err = omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			require.Eventuallyf(t, func() bool {
				newBalance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
				require.NoError(t, err)

				// we subtract the burn fee twice to account for the tx fees (which are expected to be below the burn fee)
				minBalanceAfterDelegation := bi.Sub(balance, burnFee, burnFee)
				minExpectedBalance := bi.Add(minBalanceAfterDelegation, undelegatedAmount)

				return bi.GTE(newBalance, minExpectedBalance)
			}, valChangeWait, 500*time.Millisecond, "failed to undeleate")

			// ensure rewards are still accruing
			require.Eventuallyf(t, func() bool {
				_, ok := queryDelegationRewards(t, ctx, cprov, delegatorCosmosAddr, val.OperatorAddress)

				return ok
			}, valChangeWait, 500*time.Millisecond, "no rewards increase")

			balance, err = omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
			require.NoError(t, err)

			txOpts, err = omniBackend.BindOpts(ctx, delegatorEthAddr)
			require.NoError(t, err)
			txOpts.Value = burnFee

			remainingStake := delegatedStake(t, ctx, cprov, validator.OperatorAddress, delegatorCosmosAddr.String())
			remainingAmount, err := evmredenom.ToEVMAmount(remainingStake) // remainingAmount=remainingStake*factor
			require.NoError(t, err)

			// undelegate remaining half
			tx, err = contract.Undelegate(txOpts, validatorAddr, remainingAmount)
			require.NoError(t, err)

			_, err = omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			require.Eventuallyf(t, func() bool {
				newBalance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
				require.NoError(t, err)

				// we subtract the burn fee twice to account for the tx fees (which are expected to be below the burn fee)
				minBalanceAfterDelegation := bi.Sub(balance, burnFee, burnFee)
				minExpectedBalance := bi.Add(minBalanceAfterDelegation, remainingAmount)

				return bi.GTE(newBalance, minExpectedBalance)
			}, valChangeWait, 500*time.Millisecond, "failed to undeleate")

			// ensure nothing is staked anymore
			remainingStake = delegatedStake(t, ctx, cprov, validator.OperatorAddress, delegatorCosmosAddr.String())
			remainingAmt, err := evmredenom.ToEVMAmount(remainingStake) // remainingAmt=remainingStake*factor
			require.NoError(t, err)

			log.Debug(ctx, "Remaining delegation", "stake", remainingStake, "amount", remainingAmt)
			require.Equal(t, int64(0), remainingAmt.Int64())

			// ensure no rewards accrue anymore because we undelegated everything
			waitForBlocks(ctx, t, cprov, 1)
			_, ok = queryDelegationRewards(t, ctx, cprov, delegatorCosmosAddr, validator.OperatorAddress)
			require.False(t, ok)
		})
	})
}

// deploy deploys a proxy smart contract that simply batches arbitrary calls to Staking.sol.
func deploy(
	ctx context.Context,
	backend *ethbackend.Backend,
	stakingContractAddr,
	deployer common.Address,
) (common.Address, error) {
	txOpts, err := backend.BindOpts(ctx, deployer)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "binding opts")
	}

	address, tx, _, err := bindings.DeployStakingProxy(txOpts, backend, stakingContractAddr)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "deployment")
	}
	if _, err = backend.WaitMined(ctx, tx); err != nil {
		return common.Address{}, errors.Wrap(err, "mining")
	}

	return address, nil
}

// proxyCall executes batches delegations and undelegations to Staking.sol.
func proxyCall(
	ctx context.Context,
	backend *ethbackend.Backend,
	contractAddress, caller common.Address,
	calls []bindings.StakingProxyCall,
) error {
	totalValue := big.NewInt(0)
	for _, call := range calls {
		totalValue = bi.Add(totalValue, call.Value)
	}

	opts, err := backend.BindOpts(ctx, caller)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	opts.Value = totalValue

	contract, err := bindings.NewStakingProxy(contractAddress, backend)
	if err != nil {
		return errors.Wrap(err, "instantiate bindings")
	}

	tx, err := contract.Proxy(opts, calls)
	if err != nil {
		return errors.Wrap(err, "proxy")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "mined")
	}

	return nil
}

// waitForBlocks wait for `n` blocks.
func waitForBlocks(ctx context.Context, t *testing.T, cprov provider.Provider, n uint64) {
	t.Helper()

	log.Debug(ctx, "Awaiting", "blocks", n)

	s, err := cprov.QueryClients().Node.Status(ctx, &node.StatusRequest{})
	require.NoError(t, err)

	target := s.Height + n

	require.Eventually(t, func() bool {
		s, err := cprov.QueryClients().Node.Status(ctx, &node.StatusRequest{})
		require.NoError(t, err)

		return int64(s.Height) >= int64(target)
	}, time.Second*time.Duration(target*2), time.Millisecond*100)
}
