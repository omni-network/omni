package main

import (
	"math/big"
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/omnitoken"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokenmeta"

	"github.com/stretchr/testify/require"
)

// TestBridgeBalance the native bridge balance is as expected.
// For all non-mainnet network, balance should be omnitoken.TotalSupply.
// For mainnet, we should decrement evm prefund and genesis validator allocations.
func TestBridgeBalance(t *testing.T) {
	t.Parallel()

	// mainnet prefunds
	mp := bi.Zero()
	for _, role := range eoa.AllRoles() {
		th, ok := eoa.GetFundThresholds(tokenmeta.OMNI, netconf.Mainnet, role)
		if !ok {
			continue
		}
		mp = bi.Add(mp, th.TargetBalance())
	}

	// Note that there were actually only 2 100 OMNI mainnet genesis validators. These calcs are wrong.
	mp = bi.Add(mp,
		bi.Ether(1000), // 1000 OMNI: genesis validator 1
		bi.Ether(1000), // 1000 OMNI: genesis validator 2
		bi.Ether(1000), // 1000 OMNI: genesis validator 3
		bi.Ether(1000), // 1000 OMNI: genesis validator 4
	)

	tests := []struct {
		name     string
		network  netconf.ID
		expected *big.Int
	}{
		{
			name:     "devnet",
			network:  netconf.Devnet,
			expected: omnitoken.TotalSupply(),
		},
		{
			name:     "staging",
			network:  netconf.Staging,
			expected: omnitoken.TotalSupply(),
		},
		{
			name:     "omega",
			network:  netconf.Omega,
			expected: omnitoken.TotalSupply(),
		},
		{
			name:     "mainnet",
			network:  netconf.Mainnet,
			expected: bi.Sub(omnitoken.TotalSupply(), mp),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			balance, err := getNativeBridgeBalance(tt.network)
			require.NoError(t, err)
			require.Equal(t, tt.expected, balance)
		})
	}
}
