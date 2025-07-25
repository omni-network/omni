// Package withdraw provides tools to automatically create EVM withdrawals
// for any transfers into user accounts.
package withdraw

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/halo/evmredenom"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// BankWrapper wraps x/bank.Keeper by overriding methods with
// creation of a new withdrawal requests.
type BankWrapper struct {
	bankkeeper.Keeper

	ak              AccountKeeper
	EVMEngineKeeper EVMEngineKeeper
}

func NewBankWrapper(k bankkeeper.Keeper, ak AccountKeeper) *BankWrapper {
	return &BankWrapper{Keeper: k, ak: ak}
}

func (w *BankWrapper) SetEVMEngineKeeper(keeper EVMEngineKeeper) {
	w.EVMEngineKeeper = keeper
}

// SendCoinsFromModuleToAccountNoWithdrawal bypasses the EVM withdrawal creation.
// This is required when "depositing" funds from the EVM.
func (w *BankWrapper) SendCoinsFromModuleToAccountNoWithdrawal(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, coins sdk.Coins) error {
	err := w.Keeper.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, coins)
	if err != nil {
		return errors.Wrap(err, "send coins from module to account")
	}

	return nil
}

// UndelegateCoinsFromModuleToAccount intercepts all principal undelegations and
// creates EVM withdrawal to the user account.
func (w *BankWrapper) UndelegateCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, coins sdk.Coins) error {
	log.Debug(ctx, "Undelegating coins from module to account", "sender", senderModule, "recipient", recipientAddr, "coins", coins)

	if acc := w.ak.GetModuleAccount(ctx, senderModule); acc == nil {
		return errors.New("module account does not exist [BUG]", "module_name", senderModule)
	} else if !acc.HasPermission(authtypes.Staking) {
		return errors.New("module account does not have permissions to undelegate coins", "module_name", senderModule)
	}

	return w.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, coins)
}

// SendCoinsFromModuleToAccount intercepts all "normal" bank transfers from modules to users and
// creates EVM withdrawal to the user account and burns the funds from the module.
func (w *BankWrapper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, coins sdk.Coins) error {
	if acc := w.ak.GetModuleAccount(ctx, senderModule); acc == nil {
		return errors.New("module account does not exist [BUG]", "module_name", senderModule)
	} else if !acc.HasPermission(authtypes.Burner) {
		return errors.New("module account does not have permissions to burn coins", "module_name", senderModule)
	}

	acc := w.ak.GetAccount(ctx, recipientAddr)
	if acc == nil {
		return errors.New("recipient account does not exist [BUG]", "recipient_addr", recipientAddr)
	}

	_, ok := acc.(banktypes.VestingAccount)
	if ok {
		return errors.New("vesting accounts are not supported [BUG]")
	}

	if w.EVMEngineKeeper == nil {
		return errors.New("nil EVMEngineKeeper [BUG]")
	} else if anyAmountNil(coins) {
		return errors.New("invalid nil amount [BUG]")
	} else if !coins.IsValid() { // This ensures amounts are positive
		return errors.New("invalid coins [BUG]")
	} else if len(coins) != 1 {
		return errors.New("invalid number of coins, only 1 supported [BUG]")
	} else if coins[0].Denom != sdk.DefaultBondDenom {
		return errors.New("invalid coin denom, only bond denom supported [BUG]")
	}

	addr, err := cast.EthAddress(recipientAddr)
	if err != nil {
		return errors.Wrap(err, "convert to eth address [BUG]")
	}

	weiAmount, err := evmredenom.ToEVMAmount(coins[0])
	if err != nil {
		return err
	}

	if err := w.EVMEngineKeeper.InsertWithdrawal(ctx, addr, weiAmount); err != nil {
		return err
	}

	if err := w.BurnCoins(ctx, senderModule, coins); err != nil {
		return errors.Wrap(err, "burn coins")
	}

	return nil
}

// toGwei converts a wei amount to a gwei amount and the wei remainder.
func toGwei(weiAmount *big.Int) (gweiU64 uint64, weiRemU64 uint64, err error) { //nolint:nonamedreturns // Disambiguate identical return types.
	const giga uint64 = 1e9
	gweiAmount := bi.DivRaw(weiAmount, giga)
	weiRem := bi.Sub(weiAmount, bi.MulRaw(gweiAmount, giga))

	// This should work up to 18G ETH
	if !gweiAmount.IsUint64() {
		return 0, 0, errors.New("invalid amount [BUG]")
	}

	return gweiAmount.Uint64(), weiRem.Uint64(), nil
}

// anyAmountNil returns true if any coin has a nil amount.
// This was raised during SigmaPrime audit that found that
// coins.Valid will panic if any coin has nil amount.
func anyAmountNil(coins sdk.Coins) bool {
	for _, coin := range coins {
		if coin.Amount.IsNil() {
			return true
		}
	}

	return false
}
