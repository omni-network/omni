package e2e_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/ethclient"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/stretchr/testify/require"
)

// TestPredeploys esnures that the predeploys are correctly set up.
func TestPredeploys(t *testing.T) {
	t.Parallel()
	testOmniEVM(t, func(t *testing.T, client ethclient.Client) {
		t.Helper()
		ctx := context.Background()
		code, err := client.CodeAt(ctx, common.HexToAddress(predeploys.OmniStake), nil)
		require.NoError(t, err)
		require.Equal(t, bindings.OmniStakeDeployedBytecode, hexutil.Encode(code), "OmniStake predeploy code mismatch")
	})
}
