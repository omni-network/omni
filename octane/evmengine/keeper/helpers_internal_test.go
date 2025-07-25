package keeper

import (
	gomath "math"
	"testing"

	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"
)

func TestToGwei(t *testing.T) {
	t.Parallel()
	tests := []struct {
		arg  math.Int
		gwei uint64
		wei  uint64
	}{
		{math.NewInt(0), 0, 0},
		{math.NewInt(1), 0, 1},
		{math.NewInt(params.GWei), 1, 0},
		{math.NewInt(params.GWei).AddRaw(1), 1, 1},
		{math.NewInt(params.Ether), params.GWei, 0},
		{math.NewInt(params.Ether).AddRaw(1), params.GWei, 1},
	}
	for _, tt := range tests {
		t.Run(tt.arg.String(), func(t *testing.T) {
			t.Parallel()
			gwei, wei, err := toGwei(tt.arg.BigInt())
			require.NoError(t, err)
			require.Equal(t, tt.gwei, gwei)
			require.Equal(t, tt.wei, wei)
		})
	}

	const maxUint64 = gomath.MaxUint64
	_, _, err := toGwei(math.NewIntFromUint64(maxUint64).MulRaw(2).MulRaw(params.GWei).BigInt())
	require.ErrorContains(t, err, "invalid amount [BUG]")
}
