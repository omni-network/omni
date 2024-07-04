package eoa_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"github.com/stretchr/testify/require"
)

func TestStatic(t *testing.T) {
	t.Parallel()
	for _, network := range []netconf.ID{netconf.Devnet, netconf.Staging, netconf.Omega, netconf.Mainnet} {
		for _, role := range eoa.AllRoles() {
			acc, ok := eoa.AccountForRole(network, role)
			require.True(t, ok, "account not found: %s %s", network, role)
			require.NotZero(t, acc.Address)
			require.True(t, common.IsHexAddress(acc.Address.Hex()))

			thresholds, ok := eoa.GetFundThresholds(network, acc.Role)
			require.True(t, ok, "thresholds not found")

			require.NotPanics(t, func() {
				mini := thresholds.MinBalance()
				target := thresholds.TargetBalance()
				t.Logf("Thresholds: network=%s, role=%s, min=%s, target=%s",
					network, acc.Role, etherStr(mini), etherStr(target))
			})
		}
	}
}

func etherStr(amount *big.Int) string {
	b, _ := amount.Float64()
	b /= params.Ether

	return fmt.Sprintf("%.4f", b)
}
