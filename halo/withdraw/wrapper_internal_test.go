package withdraw

import (
	"context"
	gomath "math"
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/stretchr/testify/require"
)

func TestWrapper(t *testing.T) {
	t.Parallel()
	tests := []struct {
		arg math.Int
	}{
		{math.NewInt(1)},
		{math.NewInt(params.GWei)},
		{math.NewInt(params.GWei).AddRaw(1)},
	}

	for _, tt := range tests {
		module := tt.arg.String()
		address := tutil.RandomAddress()

		t.Run(tt.arg.String(), func(t *testing.T) {
			t.Parallel()

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
			engKeeper := testEVMEngKeeper(func(ctx context.Context, withdrawalAddr common.Address, amountWei *big.Int) error {
				require.False(t, withdrawn)
				require.Equal(t, address, withdrawalAddr)
				tutil.RequireEQ(t, tt.arg.BigInt(), amountWei)
				withdrawn = true

				return nil
			})

			w := NewBankWrapper(keeper, testAccountKeeper{})
			w.SetEVMEngineKeeper(engKeeper)
			err := w.SendCoinsFromModuleToAccount(t.Context(), module, address.Bytes(), sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, tt.arg)))
			require.NoError(t, err)
			require.True(t, burnt)
			require.True(t, withdrawn)
		})

		t.Run(tt.arg.String(), func(t *testing.T) {
			t.Parallel()

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
			engKeeper := testEVMEngKeeper(func(ctx context.Context, withdrawalAddr common.Address, amountWei *big.Int) error {
				require.False(t, withdrawn)
				require.Equal(t, address, withdrawalAddr)
				tutil.RequireEQ(t, tt.arg.BigInt(), amountWei)
				withdrawn = true

				return nil
			})
			ak := testAccountKeeper{}

			w := NewBankWrapper(keeper, ak)
			w.SetEVMEngineKeeper(engKeeper)
			err := w.UndelegateCoinsFromModuleToAccount(t.Context(), module, address.Bytes(), sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, tt.arg)))
			require.NoError(t, err)
			require.True(t, burnt)
			require.True(t, withdrawn)
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

type testEVMEngKeeper func(ctx context.Context, withdrawalAddr common.Address, amountWei *big.Int) error

func (t testEVMEngKeeper) InsertWithdrawal(ctx context.Context, withdrawalAddr common.Address, amountWei *big.Int) error {
	return t(ctx, withdrawalAddr, amountWei)
}

type testBankKeeper struct {
	bankkeeper.Keeper
	BurnFunc func(ctx context.Context, moduleName string, amt sdk.Coins) error
}

func (k testBankKeeper) BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error {
	return k.BurnFunc(ctx, moduleName, amt)
}

type testAccountKeeper struct{}

func (testAccountKeeper) GetAccount(context.Context, sdk.AccAddress) sdk.AccountI {
	return &authtypes.BaseAccount{}
}

func (testAccountKeeper) GetModuleAccount(context.Context, string) sdk.ModuleAccountI {
	return authtypes.NewModuleAccount(&authtypes.BaseAccount{}, "fake", authtypes.Staking, authtypes.Burner)
}
