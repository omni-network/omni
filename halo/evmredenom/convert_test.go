package evmredenom_test

import (
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/omni-network/omni/halo/evmredenom"
	"github.com/omni-network/omni/lib/bi"

	cosmosmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestToBondCoin(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		amount   *big.Int
		expected sdk.Coin
	}{
		{
			name:     "zero amount",
			amount:   bi.N(0),
			expected: sdk.NewCoin(sdk.DefaultBondDenom, toSDKMath(bi.Ether(0))),
		},
		{
			name:     "one wei",
			amount:   bi.N(1),
			expected: sdk.NewCoin(sdk.DefaultBondDenom, toSDKMath(bi.N(1))),
		},
		{
			name:     "one ether in wei",
			amount:   bi.Ether(1),
			expected: sdk.NewCoin(sdk.DefaultBondDenom, toSDKMath(bi.Ether(1))),
		},
		{
			name:     "large amount",
			amount:   bi.N(math.MaxInt64),
			expected: sdk.NewCoin(sdk.DefaultBondDenom, toSDKMath(bi.N(math.MaxInt64))),
		},
		{
			name:     "random amount",
			amount:   bi.N(123456789),
			expected: sdk.NewCoin(sdk.DefaultBondDenom, toSDKMath(bi.N(123456789))),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := evmredenom.ToBondCoin(tt.amount)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestToEVMAmount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		coin        sdk.Coin
		expected    *big.Int
		expectError bool
	}{
		{
			name:        "zero bond coin",
			coin:        sdk.NewCoin(sdk.DefaultBondDenom, toSDKMath(bi.Ether(0))),
			expected:    bi.N(0),
			expectError: false,
		},
		{
			name:        "one bond coin",
			coin:        sdk.NewCoin(sdk.DefaultBondDenom, toSDKMath(bi.Ether(1))),
			expected:    bi.Ether(1),
			expectError: false,
		},
		{
			name:        "large bond coin amount",
			coin:        sdk.NewCoin(sdk.DefaultBondDenom, toSDKMath(bi.Ether(1))),
			expected:    bi.Ether(1),
			expectError: false,
		},
		{
			name:        "random bond coin amount",
			coin:        sdk.NewCoin(sdk.DefaultBondDenom, toSDKMath(bi.Ether(987654321))),
			expected:    bi.Ether(987654321),
			expectError: false,
		},
		{
			name:        "wrong denomination",
			coin:        sdk.NewCoin("wrongdenom", toSDKMath(bi.Ether(100))),
			expected:    nil,
			expectError: true,
		},
		{
			name:        "custom denomination",
			coin:        sdk.NewCoin("custom", toSDKMath(bi.Ether(100))),
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result, err := evmredenom.ToEVMAmount(tt.coin)

			if tt.expectError {
				require.Error(t, err)
				require.Nil(t, result)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestRoundTripConversion(t *testing.T) {
	t.Parallel()

	testAmounts := []*big.Int{
		bi.N(0),
		bi.N(1),
		bi.N(1000),
		bi.Ether(1), // 1 ETH in wei
		bi.N(123456789),
		bi.N(1311768467463790320),
	}

	for i, amount := range testAmounts {
		t.Run(fmt.Sprintf("round_trip_%d", i), func(t *testing.T) {
			t.Parallel()

			// Convert to bond coin
			bondCoin := evmredenom.ToBondCoin(amount)

			// Convert back to EVM amount
			evmAmount, err := evmredenom.ToEVMAmount(bondCoin)
			require.NoError(t, err)

			// Should be equal to original
			require.Equal(t, amount, evmAmount)
		})
	}
}

// toSDKMath converts a *big.Int to cosmossdk.io/math.Int.
func toSDKMath(amount *big.Int) cosmosmath.Int {
	return cosmosmath.NewIntFromBigInt(amount)
}
