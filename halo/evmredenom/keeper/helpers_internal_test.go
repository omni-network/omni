package keeper

import (
	"fmt"
	"testing"

	"github.com/omni-network/omni/halo/evmredenom"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum/common"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdktestutil "github.com/cosmos/cosmos-sdk/testutil"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	t.Parallel()

	key := storetypes.NewKVStoreKey(types.ModuleName)
	storeSvc := runtime.NewKVStoreService(key)
	k, err := New(storeSvc)
	require.NoError(t, err)
	ctx := sdktestutil.DefaultContext(key, storetypes.NewTransientStoreKey("test_key"))

	err = k.InitStatus(ctx, tutil.RandomHash())
	require.NoError(t, err)
	err = k.InitStatus(ctx, tutil.RandomHash())
	require.Error(t, err)
}

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
			result, err := calcMint(uint256.NewInt(tt.amount), evmredenom.EVMToBondMultiplier)
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
