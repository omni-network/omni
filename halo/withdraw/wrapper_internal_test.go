package withdraw

import (
	"context"
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
