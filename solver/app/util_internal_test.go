package app

import (
	"flag"
	"sync"
	"testing"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain/connect"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

var (
	integration = flag.Bool("integration", false, "enable integration tests")

	claimNetwork = netconf.Omega
	toClaim      = map[uint64][]OrderID{
		evmchain.IDBaseSepolia: {
			OrderID(common.HexToHash("0xf57063cbb9ae3149a61e2ab4b6680341b060a7a7ea6e089baab1cff6d54ea44c")),
			OrderID(common.HexToHash("0xafb56f273df8e28757cab87880c4b8d0c8f526c0c35621e4686e7dc8fb6a1e1e")),
			OrderID(common.HexToHash("0x3b88a49665cb86e305133ae71ab63b8f75ebffc187824f6e4055935041469548")),
			OrderID(common.HexToHash("0x9d886375111a35c85f01d680e647aae8584afeadd903c6d8c24ff5ca037e6308")),
		},
	}
)

// TestManualClaim manually claims specific orders on a network.
func TestManualClaim(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("integration tests not enabled")
	}

	ctx := t.Context()
	conn, err := connect.New(ctx, claimNetwork, connect.WithInfuraENV("INFURA_SECRET"))
	require.NoError(t, err)

	var solverAddr common.Address
	var once sync.Once

	for chainID, orders := range toClaim {
		t.Logf("Claiming orders for %s", evmchain.Name(chainID))

		inbox, err := conn.SolverNetInbox(ctx, chainID)
		require.NoError(t, err)

		orderGetter := newOrderGetter(map[uint64]*bindings.SolverNetInbox{chainID: inbox})

		for _, orderID := range orders {
			order, ok, err := orderGetter(ctx, chainID, orderID)
			require.NoError(t, err)
			require.True(t, ok, "order not found")

			if order.Status != solvernet.StatusFilled {
				t.Logf("Skipping order no in filled state: %s=%s", orderID, order.Status)
				continue
			}

			// When required, download the solver private key and add to backends.
			once.Do(func() {
				t.Log("Downloading solver private key")
				solverPrivKey, err := eoa.PrivateKey(ctx, claimNetwork, eoa.RoleSolver)
				require.NoError(t, err)
				solverAddr, err = conn.Backends.AddAccount(solverPrivKey)
				require.NoError(t, err)
			})

			txOpts, err := conn.BindOpts(ctx, chainID, solverAddr)
			require.NoError(t, err)

			tx, err := inbox.Claim(txOpts, order.ID, solverAddr)
			require.NoError(t, err)

			rec, err := conn.WaitMined(ctx, chainID, tx)
			require.NoError(t, err)

			t.Logf("Claimed order %s on %s: tx=%s (height=%s)", order.ID, evmchain.Name(chainID), tx.Hash(), rec.BlockNumber)
		}
	}
}
