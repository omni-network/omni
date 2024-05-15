package e2e_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/omni-network/omni/e2e/app/geth"
	"github.com/omni-network/omni/lib/ethclient"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

// TestGethConfig ensure that the geth config is setup correctly.
func TestGethConfig(t *testing.T) {
	t.Parallel()
	testOmniEVM(t, func(t *testing.T, client ethclient.Client) {
		t.Helper()
		ctx := context.Background()

		cfg := geth.MakeGethConfig(geth.Config{})

		block, err := client.BlockByNumber(ctx, big.NewInt(1))
		require.NoError(t, err)

		require.EqualValues(t, int(cfg.Eth.Miner.GasCeil), int(block.GasLimit()))
		require.Equal(t, big.NewInt(0), block.Difficulty())

		require.NotNil(t, block.BeaconRoot())
		require.NotEqual(t, common.Hash{}, *block.BeaconRoot())
	})
}
