package routerecon

import (
	"context"
	"testing"

	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/xchain"
	xconnect "github.com/omni-network/omni/lib/xchain/connect"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -integration -v -run=TestReconOnce

func TestReconOnce(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := context.Background()
	endpoints := xchain.RPCEndpoints{
		"omni_evm":     "https://omega.omni.network",
		"op_sepolia":   types.PublicRPCByName("op_sepolia"),
		"arb_sepolia":  types.PublicRPCByName("arb_sepolia"),
		"base_sepolia": types.PublicRPCByName("base_sepolia"),
		"holesky":      types.PublicRPCByName("holesky"),
	}
	conn, err := xconnect.New(ctx, netconf.Omega, endpoints)
	require.NoError(t, err)

	crossTx, err := paginateLatestCrossTx(ctx)
	require.NoError(t, err)

	err = reconOnce(ctx, conn.Network, conn.XProvider, conn.EthClients, crossTx)
	tutil.RequireNoError(t, err)
}
