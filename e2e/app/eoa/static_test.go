package eoa_test

import (
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

func TestStatic(t *testing.T) {
	t.Parallel()
	for _, network := range []netconf.ID{netconf.Devnet, netconf.Staging, netconf.Testnet, netconf.Mainnet} {
		for _, role := range eoa.AllRoles() {
			acc, ok := eoa.AccountForRole(network, role)
			require.True(t, ok, "account not found: %s %s", network, role)
			require.NotZero(t, acc.Address)
			require.True(t, common.IsHexAddress(acc.Address.Hex()))
		}
	}
}
