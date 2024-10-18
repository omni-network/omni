package main

import (
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/contracts/omnitoken"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/params"

	"github.com/stretchr/testify/require"
)

// TestBridgeBalance the native bridge balance is as expected.
// For all non-mainnet network, balance should be omnitoken.TotalSupply.
// For mainnet, we should decrement evm prefund and genesis validator allocations.
func TestBridgeBalance(t *testing.T) {
	t.Parallel()

	// mainnet prefunds
	mp := big.NewInt(0)
	mp = add(mp, div(ether(1), 10)) // 0.1  OMNI: create3-deployer
	mp = add(mp, div(ether(1), 10)) // 0.1  OMNI: deployer
	mp = add(mp, ether(10))         // 10   OMNI: manager
	mp = add(mp, ether(10))         // 10   OMNI: upgrader
	mp = add(mp, ether(100))        // 100  OMNI: relayer
	mp = add(mp, ether(100))        // 100  OMNI: monitor
	mp = add(mp, ether(500))        // 500  OMNI: funder
	mp = add(mp, ether(1000))       // 1000 OMNI: genesis validator 1
	mp = add(mp, ether(1000))       // 1000 OMNI: genesis validator 2

	tests := []struct {
		name     string
		network  netconf.ID
		expected *big.Int
	}{
		{
			name:     "devnet",
			network:  netconf.Devnet,
			expected: omnitoken.TotalSupply,
		},
		{
			name:     "staging",
			network:  netconf.Staging,
			expected: omnitoken.TotalSupply,
		},
		{
			name:     "omega",
			network:  netconf.Omega,
			expected: omnitoken.TotalSupply,
		},
		{
			name:     "mainnet",
			network:  netconf.Mainnet,
			expected: new(big.Int).Sub(omnitoken.TotalSupply, mp),
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

func ether(n int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(n), big.NewInt(params.Ether))
}

func div(n *big.Int, d int64) *big.Int {
	return new(big.Int).Div(n, big.NewInt(d))
}

func add(x, y *big.Int) *big.Int {
	return new(big.Int).Add(x, y)
}
