package e2e_test

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/omni-network/omni/cli/cmd"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/halo/evmredenom"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/txmgr"

	"github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	dtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
)

// Y $Native == Y * factor $STAKE(Power).
const factor = evmredenom.EVMToBondMultiplier

// execCLI will execute provided command with the arguments and return an error in case
// execution fails. It always returns stdOut and stdErr as well.
func execCLI(ctx context.Context, args ...string) (string, string, error) {
	outBuf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)

	root := cmd.New()
	root.SetOut(outBuf)
	root.SetErr(errBuf)

	root.SetArgs(args)
	if err := root.ExecuteContext(ctx); err != nil {
		return outBuf.String(), errBuf.String(), errors.Wrap(err, "executing CLI", "args", args)
	}

	return outBuf.String(), errBuf.String(), nil
}

// TestCLIOperator test the omni operator cli subcommands.
// The test runs the following commands:
// - operator create-validator creates a new validator and makes sure the validator is added to the consensus chain
// - operator delegate increases the newly created validator stake and makes sure its power is increased
// - delegator delegates stake
// - delegator makes sure rewards are accruing
// - delegator delegates more stake and make sure a withdrawals request is persisted
//
// Since they rely first on validator being created it must be run as a unit.
//
//nolint:paralleltest // We have to run self-delegation and delegation tests sequentially
func TestCLIOperator(t *testing.T) {
	t.Parallel()

	skipFunc := func(manifest types.Manifest) bool {
		return !manifest.AllE2ETests
	}
	maybeTestNetwork(t, skipFunc, func(ctx context.Context, t *testing.T, deps NetworkDeps) {
		t.Helper()

		network := deps.Network
		netID := network.ID
		omniBackend, err := deps.OmniBackend()
		require.NoError(t, err)

		// use an existing test anvil account for new validator and write it's pkey to temp file
		validatorPriv, validatorAddr := GenFundedEOA(ctx, t, omniBackend)
		validatorPubBz := ethcrypto.CompressPubkey(&validatorPriv.PublicKey)
		tmpDir := t.TempDir()
		privKeyFile := filepath.Join(tmpDir, "privkey")
		require.NoError(t, ethcrypto.SaveECDSA(privKeyFile, validatorPriv))

		cl, err := http.New(network.ID.Static().ConsensusRPC(), "/websocket")
		require.NoError(t, err)
		cprov := provider.NewABCI(cl, network.ID)

		const valChangeWait = 1 * time.Minute

		// operator's initial and self delegations
		const opInitDelegation = uint64(10 * factor)
		const opSelfDelegation = uint64(1 * factor)

		// create a new valdiator and self-delegate
		t.Run("create validator and self-delegate", func(t *testing.T) {
			// operator create-validator test
			stdOut, _, err := execCLI(
				ctx, "operator", "create-validator",
				"--network", netID.String(),
				"--private-key-file", privKeyFile,
				"--consensus-pubkey-hex", hex.EncodeToString(validatorPubBz),
				// we use minimum stake so the new validator doesn't affect the network too much
				"--self-delegation", fmt.Sprintf("%d", opInitDelegation),
				"--execution-rpc", omniBackend.Address(),
			)
			require.NoError(t, err)
			require.Empty(t, stdOut)

			require.Eventuallyf(t, func() bool {
				_, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
				return ok
			}, valChangeWait, 500*time.Millisecond, "failed to create validator")

			// make sure the validator now exists and has correct power
			val, ok, err := cprov.SDKValidator(ctx, validatorAddr)
			require.NoError(t, err)
			require.True(t, ok)
			power, err := val.Power()
			require.NoError(t, err)
			require.Equal(t, opInitDelegation, power*factor)

			// delegate more stake for the validator, since we are using an anvil account
			// it is already sufficiently funded
			stdOut, _, err = execCLI(
				ctx, "operator", "delegate",
				"--network", netID.String(),
				"--private-key-file", privKeyFile,
				"--amount", fmt.Sprintf("%d", opSelfDelegation),
				"--execution-rpc", omniBackend.Address(),
				"--self",
			)
			require.NoError(t, err)
			require.Empty(t, stdOut)

			// make sure the validator power is actually increased
			require.Eventuallyf(t, func() bool {
				val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
				require.True(t, ok)
				newPower, err := val.Power()
				require.NoError(t, err)

				return newPower*factor == opInitDelegation+opSelfDelegation
			}, valChangeWait, 500*time.Millisecond, "failed to self-delegate")
		})

		// delegator's keys
		delegatorPrivKey, delegatorEthAddr := GenFundedEOA(ctx, t, omniBackend)
		delegatorCosmosAddr := sdk.AccAddress(delegatorEthAddr.Bytes())
		delegatorPrivKeyFile := filepath.Join(tmpDir, "delegator_privkey")
		err = ethcrypto.SaveECDSA(delegatorPrivKeyFile, delegatorPrivKey)
		require.NoError(t, err)

		const delegatorDelegation = uint64(7 * factor)
		// delegate from a new account
		t.Run("delegation", func(t *testing.T) {
			// delegator delegate test
			stdOut, _, err := execCLI(
				ctx, "operator", "delegate",
				"--network", netID.String(),
				"--validator-address", validatorAddr.Hex(),
				"--private-key-file", delegatorPrivKeyFile,
				"--amount", fmt.Sprintf("%d", delegatorDelegation),
				"--execution-rpc", omniBackend.Address(),
			)
			require.NoError(t, err)
			require.Empty(t, stdOut)

			// make sure the validator power is increased and the delegation can be found
			require.Eventuallyf(t, func() bool {
				val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
				require.True(t, ok)
				power, err := val.Power()
				require.NoError(t, err)

				expect := opInitDelegation + opSelfDelegation + delegatorDelegation
				if power*factor != expect {
					return false
				}

				if delegatedStake(t, ctx, cprov, val.OperatorAddress, delegatorCosmosAddr.String()).IsZero() {
					return false
				}

				return true
			}, valChangeWait, 500*time.Millisecond, "failed to delegate")
		})

		// edit validator data
		t.Run("edit validator", func(t *testing.T) {
			val, ok, err := cprov.SDKValidator(ctx, validatorAddr)
			require.NoError(t, err)
			require.True(t, ok)

			// Edit validator moniker
			newMoniker := val.Description.Moniker + "*"
			// Add 1 Omni to current minSelf, then convert from wei to Omni.
			newMinSelfEther := val.MinSelfDelegation.AddRaw(params.Ether).QuoRaw(params.Ether)
			newMinSelfWei := newMinSelfEther.MulRaw(params.Ether)
			stdOut, stdErr, err := execCLI(
				ctx, "operator", "edit-validator",
				"--network", netID.String(),
				"--private-key-file", privKeyFile,
				"--execution-rpc", omniBackend.Address(),
				"--moniker", newMoniker,
				"--min-self-delegation", newMinSelfEther.String(),
			)
			require.NoError(t, err)
			require.Empty(t, stdOut)
			t.Log(stdErr)

			// make sure the validator moniker and min-self-delegation is actually increased
			require.Eventuallyf(t, func() bool {
				val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
				require.True(t, ok)

				return val.GetMoniker() == newMoniker && val.MinSelfDelegation.Equal(newMinSelfWei)
			}, valChangeWait, 500*time.Millisecond, "failed to edit validator")
		})

		// test rewards distribution
		t.Run("distribution", func(t *testing.T) {
			val, ok, err := cprov.SDKValidator(ctx, validatorAddr)
			require.NoError(t, err)
			require.True(t, ok)

			var originalRewards math.LegacyDec

			// fetch rewards and make sure they are present
			require.Eventuallyf(t, func() bool {
				var ok bool
				originalRewards, ok = queryDelegationRewards(t, ctx, cprov, delegatorCosmosAddr, val.OperatorAddress)

				return ok
			}, valChangeWait, 500*time.Millisecond, "no rewards")

			// fetch again and make sure they increased
			require.Eventuallyf(t, func() bool {
				latestRewards, ok := queryDelegationRewards(t, ctx, cprov, delegatorCosmosAddr, val.OperatorAddress)
				if !ok {
					return false
				}

				return latestRewards.GT(originalRewards)
			}, valChangeWait, 500*time.Millisecond, "no rewards increase")
		})

		t.Run("undelegation", func(t *testing.T) {
			var block *big.Int
			require.NoError(t, err)
			balance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, block)
			require.NoError(t, err)

			contract, err := bindings.NewStaking(common.HexToAddress(predeploys.Staking), omniBackend)
			require.NoError(t, err)

			burnFee := bi.Ether(0.1)

			txOpts, err := omniBackend.BindOpts(ctx, delegatorEthAddr)
			require.NoError(t, err)
			txOpts.Value = burnFee

			// undelegate everything
			undelegatedAmount := big.NewInt(int64(2 * delegatorDelegation))
			tx, err := contract.Undelegate(txOpts, validatorAddr, undelegatedAmount)
			require.NoError(t, err)

			_, err = omniBackend.WaitMined(ctx, tx)
			require.NoError(t, err)

			require.Eventuallyf(t, func() bool {
				newBalance, err := omniBackend.BalanceAt(ctx, delegatorEthAddr, block)
				require.NoError(t, err)

				// we subtract the burn fee twice to account for the tx fees (which are expected to be below the burn fee)
				return bi.GTE(newBalance, bi.Add(bi.Sub(balance, burnFee, burnFee), undelegatedAmount))
			}, valChangeWait, 500*time.Millisecond, "failed to undeleate")
		})
	})
}

func delegatedStake(t *testing.T, ctx context.Context, cprov provider.Provider, valAddr string, delegatorAddr string) sdk.Coin {
	t.Helper()
	response, err := cprov.QueryClients().Staking.ValidatorDelegations(ctx, &stypes.QueryValidatorDelegationsRequest{
		ValidatorAddr: valAddr,
		Pagination:    nil,
	})
	require.NoError(t, err)
	require.NotNil(t, response)

	balance := sdk.Coin{
		Denom:  "stake",
		Amount: math.NewInt(0),
	}
	for _, response := range response.DelegationResponses {
		if response.Delegation.DelegatorAddress == delegatorAddr {
			balance = balance.Add(response.Balance)
		}
	}

	return balance
}

func queryDelegationRewards(t *testing.T, ctx context.Context, cprov provider.Provider, delegatorAddr sdk.AccAddress, validatorAddr string) (math.LegacyDec, bool) {
	t.Helper()

	resp, err := cprov.QueryClients().Distribution.DelegationRewards(ctx, &dtypes.QueryDelegationRewardsRequest{
		DelegatorAddress: delegatorAddr.String(),
		ValidatorAddress: validatorAddr,
	})
	if err != nil && strings.Contains(err.Error(), "no delegation for") {
		// No delegation found
		return math.LegacyDec{}, false
	}

	if len(resp.Rewards) == 0 {
		return math.LegacyDec{}, false
	}

	require.Len(t, resp.Rewards, 1)
	require.Equal(t, sdk.DefaultBondDenom, resp.Rewards[0].Denom)

	return resp.Rewards[0].Amount, true
}

func GenFundedEOA(ctx context.Context, t *testing.T, backend *ethbackend.Backend) (*ecdsa.PrivateKey, common.Address) {
	t.Helper()

	amount1k := bi.Ether(1_000)
	funderAddr := anvil.DevAccount9()

	newKey, err := ethcrypto.GenerateKey()
	require.NoError(t, err)
	newAddr := ethcrypto.PubkeyToAddress(newKey.PublicKey)

	_, rec, err := backend.Send(ctx, funderAddr, txmgr.TxCandidate{
		To:    &newAddr,
		Value: amount1k,
	})
	require.NoError(t, err)
	require.Equal(t, ethtypes.ReceiptStatusSuccessful, rec.Status)

	bal, err := backend.BalanceAt(ctx, newAddr, nil)
	require.NoError(t, err)
	tutil.RequireEQ(t, amount1k, bal)

	_, err = backend.AddAccount(newKey)
	require.NoError(t, err)

	log.Debug(ctx, "Funded new EOA", "addr", newAddr.Hex(), "amount", bi.ToEtherF64(amount1k), "chain", backend.Name())

	return newKey, newAddr
}
