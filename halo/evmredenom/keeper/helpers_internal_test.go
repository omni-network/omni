package keeper

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

func TestCalcMint(t *testing.T) {
	t.Parallel()
	const gwei uint64 = 1e9
	tests := []struct {
		amount uint64
		expect uint64
	}{
		{amount: 0, expect: 0},
		{amount: 1, expect: 74},
		{amount: gwei, expect: gwei * 74},
		{amount: gwei + 1, expect: (gwei + 1) * 74},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.amount), func(t *testing.T) {
			t.Parallel()
			result, err := calcMint(uint256.NewInt(tt.amount), evmToBondMultiplier)
			require.NoError(t, err)
			require.Equal(t, tt.expect, result.Uint64())
		})
	}
}

func TestIncHash(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input common.Hash
		want  common.Hash
	}{
		{"zero hash", common.Hash{}, common.HexToHash("0x01")},
		{"max hash", common.MaxHash, common.Hash{}},
		{"incremented hash", common.HexToHash("0x1234567890abcdee"), common.HexToHash("0x1234567890abcdef")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := incHash(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
