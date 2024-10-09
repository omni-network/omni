package main

import (
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/contracts/omnitoken"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"
)

var (
	eth1   = math.NewInt(1).MulRaw(params.Ether).BigInt()
	eth10  = math.NewInt(10).MulRaw(params.Ether).BigInt()
	eth20  = math.NewInt(20).MulRaw(params.Ether).BigInt()
	eth100 = math.NewInt(100).MulRaw(params.Ether).BigInt()

	addr1 = common.HexToAddress("0x1")
	addr2 = common.HexToAddress("0x2")
)

// TestBridgeBalance the mainnet native bridge balance is total supply minus prefunds.
func TestBridgeBalance(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		prefunds  types.GenesisAlloc
		expected  *big.Int
		shouldErr bool
	}{
		{
			name: "1 eth",
			prefunds: types.GenesisAlloc{
				addr1: {Balance: eth1},
			},
			expected:  new(big.Int).Sub(omnitoken.TotalSupply, eth1),
			shouldErr: false,
		},
		{
			name: "20 eth",
			prefunds: types.GenesisAlloc{
				addr1: {Balance: eth10},
				addr2: {Balance: eth10},
			},
			expected:  new(big.Int).Sub(omnitoken.TotalSupply, eth20),
			shouldErr: false,
		},
		{
			name: "200 eth - more than sane",
			prefunds: types.GenesisAlloc{
				addr1: {Balance: eth100},
				addr2: {Balance: eth100},
			},
			shouldErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cfg, err := allocConfig(netconf.Mainnet, tt.prefunds)

			if tt.shouldErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expected, cfg.NativeBridgeBalance)
		})
	}
}
