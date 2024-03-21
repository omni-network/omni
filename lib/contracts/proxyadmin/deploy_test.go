package proxyadmin_test

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/create3"
	"github.com/omni-network/omni/lib/contracts/proxyadmin"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/stretchr/testify/require"
)

const (
	chainName   = "test"
	chainID     = 111
	blockPeriod = time.Second
)

func TestDeployDevnet(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	client, stop, err := anvil.Start(ctx, tutil.TempDir(t), chainID)
	require.NoError(t, err)
	t.Cleanup(stop)

	backend, err := anvil.NewBackend(chainName, chainID, blockPeriod, client)
	require.NoError(t, err)

	// devnet create3 factory is required
	addr, _, err := create3.DeployDevnet(ctx, backend)
	require.NoError(t, err)
	require.Equal(t, contracts.DevnetCreate3Factory, addr)

	addr, _, err = proxyadmin.DeployDevnet(ctx, backend)
	require.NoError(t, err)
	require.Equal(t, contracts.DevnetProxyAdmin, addr)

	proxyAdmin, err := bindings.NewProxyAdmin(addr, backend)
	require.NoError(t, err)

	owner, err := proxyAdmin.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, contracts.DevnetProxyAdminOwner, owner)
}
