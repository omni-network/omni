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
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/cosmos/cosmos-sdk/client/grpc/node"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	_ "embed"
)

const (
	DELEGATE   = 0
	UNDELEGATE = 1
)

// burnFee is the fee burned on every undelegation.
var burnFee = bi.Ether(0.1)

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

		const valChangeWait = 30 * time.Second

		validators, err := cprov.SDKValidators(ctx)
		require.NoError(t, err)
		// Skip if not enough validators
		if len(validators) < 2 {
			t.Skip()
		}
		validator := validators[0]
		altValidator := validators[1]
		validatorAddr, err := validator.OperatorEthAddr()
		require.NoError(t, err)

		delegatorPrivKey, delegatorEthAddr := GenFundedEOA(ctx, t, omniBackend)
		delegatorCosmosAddr := sdk.AccAddress(delegatorEthAddr.Bytes())
		delegatorPrivKeyFile := filepath.Join(tmpDir, "delegator_privkey")
		err = ethcrypto.SaveECDSA(delegatorPrivKeyFile, delegatorPrivKey)
		require.NoError(t, err)

		stakingContractAddr := common.HexToAddress(predeploys.Staking)

		contract, err := bindings.NewStaking(stakingContractAddr, omniBackend)
		require.NoError(t, err)

		// We deploy a proxy smart contract and let it stake and unstake in one batch
		t.Run("batched delegations and undelegations", func(t *testing.T) {
			proxyAddr, err := deploy(ctx, omniBackend, stakingContractAddr, delegatorEthAddr)
			proxyCosmosAddr := sdk.AccAddress(proxyAddr.Bytes())
			require.NoError(t, err)
			log.Debug(ctx, "Staking proxy deployed", "address", proxyAddr)

			delegation := bi.Ether(50)
			undelegation1 := bi.Ether(20)
			undelegation2 := bi.Ether(5)

			val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
			require.True(t, ok)

			require.NoError(t, err)

			// These calls will be mined in one block and batched by octane
			calls := []bindings.StakingProxyCall{
				{
					Method:    DELEGATE,
					Validator: validatorAddr,
					Value:     delegation,
					Amount:    bi.N(0),
				},
				{
					Method:    UNDELEGATE,
					Validator: validatorAddr,
					Value:     burnFee,
					Amount:    undelegation1,
				},
				{
					Method:    UNDELEGATE,
					Validator: validatorAddr,
					Value:     burnFee,
					Amount:    undelegation2,
				},
			}

			err = proxyCall(ctx, omniBackend, proxyAddr, delegatorEthAddr, calls)
			require.NoError(t, err)

			// make sure our delegated amount is delegation-undelegation1-undelegation2
			require.Eventuallyf(t, func() bool {
				amount := delegatedAmount(t, ctx, cprov, val.OperatorAddress, proxyCosmosAddr.String())
				log.Debug(ctx, "Delegated amount", "amount", amount)

				return bi.EQ(bi.Sub(delegation, undelegation1, undelegation2), amount.Amount.BigInt())
			}, valChangeWait, 500*time.Millisecond, "failed to execute batched staking calls")
		})

		const delegation = uint64(76)

		t.Run("delegation", func(t *testing.T) {
			val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
			require.True(t, ok)

			valPower, err := val.Power()
			require.NoError(t, err)

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

				if delegatedAmount(t, ctx, cprov, val.OperatorAddress, delegatorCosmosAddr.String()).IsZero() {
					return false
				}

				return newPower >= valPower+delegation
			}, valChangeWait, 500*time.Millisecond, "failed to delegate")
		})

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

			// Since we trigger an invalid undelegation, wait for a few blocks to make sure
			// the chain does not stall and our balance didn't change
			waitForBlocks(ctx, t, cprov, 5)

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

			// Since we trigger an invalid undelegation, wait for a few blocks to make sure
			// the chain does not stall and our balance didn't change
			waitForBlocks(ctx, t, cprov, 5)

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

			// Since we trigger an invalid undelegation, wait for a few blocks to make sure
			// the chain does not stall and our balance didn't change
			waitForBlocks(ctx, t, cprov, 5)

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

			// ensure rewards are still accruing after some time
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

			delegatedAmt := delegatedAmount(t, ctx, cprov, validator.OperatorAddress, delegatorCosmosAddr.String()).Amount.BigInt()

			// undelegate remaining half
			tx, err := contract.Undelegate(txOpts, validatorAddr, delegatedAmt)
			require.NoError(t, err)

			_, err = omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			require.Eventuallyf(t, func() bool {
				newBalance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, anyBlock)
				require.NoError(t, err)

				// we subtract the burn fee twice to account for the tx fees (which are expected to be below the burn fee)
				return bi.GTE(newBalance, bi.Add(bi.Sub(balance, burnFee, burnFee), delegatedAmt))
			}, valChangeWait, 500*time.Millisecond, "failed to undeleate")

			// ensure nothing is staked anymore
			remainingAmt := delegatedAmount(t, ctx, cprov, validator.OperatorAddress, delegatorCosmosAddr.String())
			log.Debug(ctx, "Remaining delegation", "amount", remainingAmt)
			require.Equal(t, int64(0), remainingAmt.Amount.Int64())

			// ensure no rewards accrue anymore because we undelegated everything
			waitForBlocks(ctx, t, cprov, 10)
			_, ok := queryDelegationRewards(t, ctx, cprov, delegatorCosmosAddr, validator.OperatorAddress)
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
