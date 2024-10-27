package e2e_test

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/omni-network/omni/cli/cmd"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/rpc/client/http"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

// execCLI will execute provided command with the arguments and return an error in case
// execution fails, or command output as string in case of success.
func execCLI(ctx context.Context, args ...string) (string, error) {
	outBuf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)

	root := cmd.New()
	root.SetOut(outBuf)
	root.SetErr(errBuf)

	root.SetArgs(args)
	if err := root.ExecuteContext(ctx); err != nil {
		return "", errors.Wrap(err, "executing CLI", "args", args)
	}

	return outBuf.String(), nil
}

// TestCLIOperator test the omni operator cli subcommands.
// The test runs the following commands:
// - operator create-validator creates a new validator and makes sure the validator is added to the consensus chain
// - operator delegate increases the newly created validator stake and makes sure its power is increased
//
// Since they rely first on validator being created it must be run as a unit.
func TestCLIOperator(t *testing.T) {
	t.Parallel()

	testNetwork(t, func(t *testing.T, network netconf.Network, endpoints xchain.RPCEndpoints) {
		t.Helper()
		testnet, _, _, _ := loadEnv(t)

		ctx := context.Background()
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

		// operator create-validator test
		const selfDelegation = uint64(100)
		res, err := execCLI(
			ctx, "operator", "create-validator",
			"--network", "devnet",
			"--private-key-file", privKeyFile,
			"--consensus-pubkey-hex", hex.EncodeToString(validatorPubBz),
			// we use minimum stake so the new validator doesn't affect the network too much
			"--self-delegation", fmt.Sprintf("%d", selfDelegation),
			"--execution-rpc", executionRPC,
		)
		require.NoError(t, err)
		require.Empty(t, res)

		cl, err := http.New(testnet.Network.Static().ConsensusRPC(), "/websocket")
		require.NoError(t, err)

		cprov := provider.NewABCI(cl, network.ID)

		// wait for validator to be created
		const valChangeWait = 15 * time.Second
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
		require.Equal(t, selfDelegation, power)

		// operator delegate test

		// delegate more stake for the validator, since we are using an anvil account
		// it is already sufficiently funded
		const delegation = uint64(1)
		res, err = execCLI(
			ctx, "operator", "delegate",
			"--network", "devnet",
			"--private-key-file", privKeyFile,
			"--amount", fmt.Sprintf("%d", delegation),
			"--execution-rpc", executionRPC,
			"--self",
		)
		require.NoError(t, err)
		require.Empty(t, res)

		// make sure the validator power is actually increased
		require.Eventuallyf(t, func() bool {
			val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
			require.True(t, ok)
			newPower, err := val.Power()
			require.NoError(t, err)

			return newPower > power
		}, valChangeWait, 500*time.Millisecond, "failed to create validator")
	})
}
