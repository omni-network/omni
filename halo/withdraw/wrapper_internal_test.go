package withdraw

import (
	"context"
	gomath "math"
	"testing"

	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/stretchr/testify/require"
)

func TestWrapper(t *testing.T) {
	t.Parallel()
	tests := []struct {
		arg  math.Int
		gwei uint64
	}{
		{math.NewInt(1), 0},
		{math.NewInt(params.GWei), 1},
		{math.NewInt(params.GWei).AddRaw(1), 1},
	}
	for _, tt := range tests {
		t.Run(tt.arg.String(), func(t *testing.T) {
			t.Parallel()
			module := tt.arg.String()
			address := tutil.RandomAddress()

			var burnt, withdrawn bool
			keeper := testBankKeeper{
				BurnFunc: func(ctx context.Context, moduleName string, amt sdk.Coins) error {
					require.False(t, burnt)
					require.Equal(t, module, moduleName)
					require.Equal(t, tt.arg.String(), amt[0].Amount.String())
					burnt = true

					return nil
				},
			}
			engKeeper := testEVMEngKeeper(func(ctx context.Context, withdrawalAddr common.Address, amountGwei uint64) error {
				require.False(t, withdrawn)
				require.Equal(t, address, withdrawalAddr)
				require.Equal(t, tt.gwei, amountGwei)
				withdrawn = true

				return nil
			})

			w := NewBankWrapper(keeper)
			w.SetEVMEngineKeeper(engKeeper)
			err := w.SendCoinsFromModuleToAccount(t.Context(), module, address.Bytes(), sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, tt.arg)))
			require.NoError(t, err)
			require.True(t, burnt)
			if tt.gwei > 0 {
				require.True(t, withdrawn)
			} else {
				require.False(t, withdrawn)
			}
		})
	}
}

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
			gwei, wei, err := toGwei(tt.arg)
			require.NoError(t, err)
			require.Equal(t, tt.gwei, gwei)
			require.Equal(t, tt.wei, wei)
		})
	}

	const maxUint64 = gomath.MaxUint64
	_, _, err := toGwei(math.NewIntFromUint64(maxUint64).MulRaw(2).MulRaw(params.GWei))
	require.ErrorContains(t, err, "invalid amount [BUG]")
}

type testEVMEngKeeper func(ctx context.Context, withdrawalAddr common.Address, amountGwei uint64) error

func (t testEVMEngKeeper) InsertWithdrawal(ctx context.Context, withdrawalAddr common.Address, amountGwei uint64) error {
	return t(ctx, withdrawalAddr, amountGwei)
}

type testBankKeeper struct {
	bankkeeper.Keeper
	BurnFunc func(ctx context.Context, moduleName string, amt sdk.Coins) error
}

func (k testBankKeeper) BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error {
	return k.BurnFunc(ctx, moduleName, amt)
}
