package e2e_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/feature"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/rpc/client/http"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

// TestDelegation creates a delegation from a non-validator and ensures it succeeds.
func TestDelegation(t *testing.T) {
	t.Parallel()
	testNetwork(t, func(t *testing.T, network netconf.Network, endpoints xchain.RPCEndpoints) {
		t.Helper()
		ctx := context.Background()
		testnet, _, _, _ := loadEnv(t)

		if feature.FlagDelegations.Enabled(ctx) {
			t.Log("Testing delegations")
			runTest(t, ctx, network, endpoints, testnet)
		}
	})
}

func runTest(t *testing.T, ctx context.Context, network netconf.Network, endpoints xchain.RPCEndpoints, testnet types.Testnet) {
	t.Helper()

	if network.ID != netconf.Devnet {
		t.Fatal("only devnet")
	}

	log.Info(ctx, "Running delegations test")

	backends, err := ethbackend.BackendsFromNetwork(network, endpoints, anvil.DevPrivateKeys()...)
	require.NoError(t, err)

	backend, err := backends.Backend(evmchain.IDOmniDevnet)
	require.NoError(t, err)

	cl, err := http.New(testnet.Network.Static().ConsensusRPC(), "/websocket")
	require.NoError(t, err)

	cprov := provider.NewABCI(cl, network.ID)
	vals, err := cprov.SDKValidators(ctx)
	require.NoError(t, err)

	if len(vals) == 0 {
		t.Fatal("no validators")
	}

	validator := vals[0]

	power, err := validator.Power()
	require.NoError(t, err)
	shares, err := validator.DelegatorShares.Float64()
	require.NoError(t, err)

	validatorAddr, err := validator.ConsensusEthAddr()
	require.NoError(t, err)

	err = delegate(ctx, backend, validatorAddr)
	require.NoError(t, err)

	const valChangeWait = 15 * time.Second
	require.Eventuallyf(t, func() bool {
		val, ok, _ := cprov.SDKValidator(ctx, validatorAddr)
		require.True(t, ok)

		newPower, err := val.Power()
		require.NoError(t, err)

		newShares, err := validator.DelegatorShares.Float64()
		require.NoError(t, err)

		log.Info(ctx, "new data", "power", power, "new_power", newPower, "shares", shares, "new_shares", newShares)

		return newPower > power && newShares > shares
	}, valChangeWait, 500*time.Millisecond, "failed to create validator")

	log.Info(ctx, "Delegation test success")
}

func delegate(ctx context.Context, backend *ethbackend.Backend, validatorAddr common.Address) error {
	contractAddress := common.HexToAddress(predeploys.Staking)
	contract, err := bindings.NewStaking(contractAddress, backend)
	if err != nil {
		return errors.Wrap(err, "new staking")
	}

	address := anvil.DevAccount8()

	txOpts, err := backend.BindOpts(ctx, address)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}
	txOpts.Value = big.NewInt(7000000000000000000)

	tx, err := contract.DelegateFor(txOpts, address, validatorAddr)
	if err != nil {
		return errors.Wrap(err, "delegateFor")
	}
	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}
	log.Info(ctx, "Delegation tx receipt", "tx_hash", receipt.TxHash)

	return nil
}
