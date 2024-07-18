package portal_test

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/create3"
	"github.com/omni-network/omni/lib/contracts/portal"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

const (
	chainName   = "test"
	chainID     = 111
	blockPeriod = time.Second
)

func TestDeployDevnet(t *testing.T) {
	t.Parallel()
	network := netconf.Devnet
	ctx := context.Background()

	client, _, stop, err := anvil.Start(ctx, tutil.TempDir(t), chainID)
	require.NoError(t, err)
	t.Cleanup(stop)

	backend, err := ethbackend.NewAnvilBackend(chainName, chainID, blockPeriod, client)
	require.NoError(t, err)

	// devnet create3 factory is required
	addr, _, err := create3.Deploy(ctx, network, backend)
	require.NoError(t, err)
	require.Equal(t, contracts.Create3Factory(network), addr)

	valSetID := uint64(1)
	vals := []bindings.Validator{
		{Addr: common.HexToAddress("0x1111"), Power: 100},
		{Addr: common.HexToAddress("0x2222"), Power: 100},
		{Addr: common.HexToAddress("0x3333"), Power: 100},
	}

	feeOracle := common.HexToAddress("0xfffff")
	addr, _, err = portal.Deploy(ctx, network, backend, feeOracle, valSetID, vals)
	require.NoError(t, err)
	require.Equal(t, contracts.Portal(network), addr)

	portal, err := bindings.NewOmniPortal(addr, backend)
	require.NoError(t, err)

	owner, err := portal.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, eoa.MustAddress(network, eoa.RoleAdmin), owner)

	// check validators
	totalPower, err := portal.ValSetTotalPower(nil, 1)
	require.NoError(t, err)
	require.Equal(t, uint64(300), totalPower)

	val1Power, err := portal.ValSet(nil, valSetID, vals[0].Addr)
	require.NoError(t, err)
	require.Equal(t, uint64(100), val1Power)

	val2Power, err := portal.ValSet(nil, valSetID, vals[1].Addr)
	require.NoError(t, err)
	require.Equal(t, uint64(100), val2Power)

	val3Power, err := portal.ValSet(nil, valSetID, vals[2].Addr)
	require.NoError(t, err)
	require.Equal(t, uint64(100), val3Power)
}
