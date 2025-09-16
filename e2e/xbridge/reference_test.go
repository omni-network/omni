package xbridge_test

import (
	"testing"

	"github.com/omni-network/omni/e2e/xbridge"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

//go:generate go test . -golden -clean

func TestReference(t *testing.T) {
	t.Parallel()

	golden := make(map[string]map[netconf.ID]map[string]any)
	ctx := t.Context()

	seen := make(map[common.Address]bool)

	for _, token := range xbridge.Tokens() {
		golden[token.Symbol()] = make(map[netconf.ID]map[string]any)

		for _, network := range netconf.All() {
			if network == netconf.Simnet {
				continue // Skip simnet since it doesn't have eoas.
			}

			if network != netconf.Mainnet {
				continue // Skip staging because salt version is dynamic, only do mainnet, since that is only L1 chain.
			}

			bridge, err := xbridge.BridgeAddr(ctx, network, token)
			require.NoError(t, err)

			lockbox, err := xbridge.LockboxAddr(ctx, network, token)
			require.NoError(t, err)

			addr, err := token.Address(ctx, network)
			require.NoError(t, err)

			wraps := token.Wraps()

			canon, err := token.Canonical(ctx, network)
			require.NoError(t, err)

			see := func(addr common.Address) {
				require.False(t, seen[addr], "duplicate address: %s", addr.Hex())
				seen[addr] = true
			}

			see(addr)
			see(bridge)
			see(lockbox)
			see(canon.Address)

			json := map[string]any{
				"addr":    addr.Hex(),
				"lockbox": lockbox.Hex(),
				"bridge":  bridge.Hex(),
				"wraps": map[string]any{
					"name":   wraps.Name,
					"symbol": wraps.Symbol,
				},
				"canonical": map[string]any{
					"address":  canon.Address.Hex(),
					"chain_id": canon.ChainID,
				},
			}

			golden[token.Symbol()][network] = json
		}
	}

	tutil.RequireGoldenJSON(t, golden)
}
