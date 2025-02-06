package e2e_test

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/omni-network/omni/cli/cmd"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/feature"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/rpc/client/http"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	dtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stretchr/testify/require"
)

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

	testNetwork(t, func(ctx context.Context, t *testing.T, network netconf.Network, endpoints xchain.RPCEndpoints) {
		t.Helper()

		e, ok := network.OmniEVMChain()
		require.True(t, ok)
		executionRPC, err := endpoints.ByNameOrID(e.Name, e.ID)
		require.NoError(t, err)

		// use an existing test anvil account for new validator and write it's pkey to temp file
		validatorPriv := anvil.DevPrivateKey6()
		validatorPubBz := ethcrypto.CompressPubkey(&validatorPriv.PublicKey)
		validatorAddr := ethcrypto.PubkeyToAddress(validatorPriv.PublicKey)
		tmpDir := t.TempDir()
		privKeyFile := filepath.Join(tmpDir, "privkey")
		require.NoError(
			t,
			ethcrypto.SaveECDSA(privKeyFile, validatorPriv),
			"failed to save new validator private key to temp file",
		)

		cl, err := http.New(network.ID.Static().ConsensusRPC(), "/websocket")
		require.NoError(t, err)
		cprov := provider.NewABCI(cl, network.ID)

		const valChangeWait = 15 * time.Second

		// operator's initial and self delegations
		const opInitDelegation = uint64(100)
		const opSelfDelegation = uint64(1)

		// create a new valdiator and self-delegate
		t.Run("self delegation", func(t *testing.T) {
			// operator create-validator test
			stdOut, _, err := execCLI(
				ctx, "operator", "create-validator",
				"--network", "devnet",
				"--private-key-file", privKeyFile,
				"--consensus-pubkey-hex", hex.EncodeToString(validatorPubBz),
				// we use minimum stake so the new validator doesn't affect the network too much
				"--self-delegation", fmt.Sprintf("%d", opInitDelegation),
				"--execution-rpc", executionRPC,
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
			require.Equal(t, opInitDelegation, power)

			// delegate more stake for the validator, since we are using an anvil account
			// it is already sufficiently funded
			stdOut, _, err = execCLI(
				ctx, "operator", "delegate",
				"--network", "devnet",
				"--private-key-file", privKeyFile,
				"--amount", fmt.Sprintf("%d", opSelfDelegation),
				"--execution-rpc", executionRPC,
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

				return newPower == opInitDelegation+opSelfDelegation
			}, valChangeWait, 500*time.Millisecond, "failed to self-delegate")
		})

		// delegator's keys
		privKey, pubKey := anvil.DevPrivateKey5(), anvil.DevAccount5()
		delegatorCosmosAddr := sdk.AccAddress(pubKey.Bytes()).String()
		delegatorPrivKeyFile := filepath.Join(tmpDir, "delegator_privkey")
		require.NoError(
			t,
			ethcrypto.SaveECDSA(delegatorPrivKeyFile, privKey),
			"failed to save new validator private key to temp file",
		)

		// delegate from a new account
		t.Run("delegation", func(t *testing.T) {
			if !feature.FlagEVMStakingModule.Enabled(ctx) {
				t.Skip("Skipping delegation tests")
			}

			// delegator delegate test
			const delegatorDelegation = uint64(700)
			stdOut, _, err := execCLI(
				ctx, "operator", "delegate",
				"--network", "devnet",
				"--validator-address", validatorAddr.Hex(),
				"--private-key-file", delegatorPrivKeyFile,
				"--amount", fmt.Sprintf("%d", delegatorDelegation),
				"--execution-rpc", executionRPC,
			)
			require.NoError(t, err)
			require.Empty(t, stdOut)

			// make sure the validator power is increased and the delegation can be found
			require.Eventuallyf(t, func() bool {
				val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
				require.True(t, ok)
				newPower, err := val.Power()
				require.NoError(t, err)

				if newPower != opInitDelegation+opSelfDelegation+delegatorDelegation {
					return false
				}

				if !delegationFound(t, ctx, cprov, val.OperatorAddress, delegatorCosmosAddr) {
					return false
				}

				return true
			}, valChangeWait, 500*time.Millisecond, "failed to delegate")
		})

		// edit validator data
		t.Run("edit validator", func(t *testing.T) {
			if !feature.FlagEVMStakingModule.Enabled(ctx) {
				t.Skip("Skipping evmstaking2 tests")
			}

			// Edit validator moniker
			const moniker = "new-moniker"
			const minSelf = 2 // TODO(corver): Also here
			stdOut, stdErr, err := execCLI(
				ctx, "operator", "edit-validator",
				"--network", netconf.Devnet.String(),
				"--private-key-file", privKeyFile,
				"--execution-rpc", executionRPC,
				"--moniker", moniker,
				"--min-self-delegation", strconv.FormatInt(minSelf, 10),
			)
			require.NoError(t, err)
			require.Empty(t, stdOut)
			t.Log(stdErr)

			minSelfWei := math.NewInt(minSelf).MulRaw(params.Ether)

			// make sure the validator moniker and min-self-delegation is actually increased
			require.Eventuallyf(t, func() bool {
				val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
				require.True(t, ok)

				return val.GetMoniker() == moniker && val.MinSelfDelegation.Equal(minSelfWei)
			}, valChangeWait, 500*time.Millisecond, "failed to edit validator")
		})

		// test rewards distribution
		t.Run("distribution", func(t *testing.T) {
			if !feature.FlagEVMStakingModule.Enabled(ctx) {
				t.Skip("Skipping evmstaking2 tests")
			}

			val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
			require.True(t, ok)

			// fetch rewards and make sure they are positive
			resp, err := cprov.QueryClients().Distribution.DelegationRewards(ctx, &dtypes.QueryDelegationRewardsRequest{
				DelegatorAddress: delegatorCosmosAddr,
				ValidatorAddress: val.OperatorAddress,
			})
			require.NoError(t, err)
			require.NotEmpty(t, resp.Rewards)

			for _, coin := range resp.Rewards {
				require.Equal(t, "stake", coin.Denom)
				require.True(t, coin.Amount.IsPositive())
			}

			// fetch again and make sure they increased
			wait := time.Second * 2
			require.Eventuallyf(t, func() bool {
				resp2, err := cprov.QueryClients().Distribution.DelegationRewards(ctx, &dtypes.QueryDelegationRewardsRequest{
					DelegatorAddress: delegatorCosmosAddr,
					ValidatorAddress: val.OperatorAddress,
				})
				require.NoError(t, err)
				require.NotEmpty(t, resp2.Rewards)

				// all fetched values should be strictly larger
				for i, coin2 := range resp2.Rewards {
					coin := resp.Rewards[i]
					if !coin2.Amount.GT(coin.Amount) {
						return false
					}
				}

				return true
			}, wait, 500*time.Millisecond, "no rewards increase")
		})
	})
}

func delegationFound(t *testing.T, ctx context.Context, cprov provider.Provider, valAddr string, delegatorAddr string) bool {
	t.Helper()
	response, err := cprov.QueryClients().Staking.ValidatorDelegations(ctx, &stypes.QueryValidatorDelegationsRequest{
		ValidatorAddr: valAddr,
		Pagination:    nil,
	})
	require.NoError(t, err)
	require.NotNil(t, response)
	log.Info(ctx, "delegations found", "number", len(response.DelegationResponses))
	for _, response := range response.DelegationResponses {
		if response.Delegation.DelegatorAddress == delegatorAddr {
			return true
		}
	}

	return false
}
