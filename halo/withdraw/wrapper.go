// Package withdraw provides tools to automatically create EVM withdrawals
// for any transfers into user accounts.
package withdraw

import (
	"context"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

// BankWrapper wraps x/bank.Keeper by overriding methods with
// creation of a new withdrawal requests.
type BankWrapper struct {
	bankkeeper.Keeper

	EVMEngineKeeper EVMEngineKeeper
}

func NewBankWrapper(k bankkeeper.Keeper) *BankWrapper {
	return &BankWrapper{Keeper: k}
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
// TODO(corver): This is unexpected in magellan upgrade, must be implemented in drake.
func (w *BankWrapper) UndelegateCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	log.Error(ctx, "Unexpected call to DelegateCoinsFromAccountToModule [BUG]", nil, "sender", senderModule, "recipient", recipientAddr, "amt", amt)

	if err := w.Keeper.UndelegateCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt); err != nil {
		return errors.Wrap(err, "undelegate coins from module to account")
	}

	return nil
}

// SendCoinsFromModuleToAccount intercepts all "normal" bank transfers from modules to users and
// creates EVM withdrawal to the user account and burns the funds from the module.
func (w *BankWrapper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, coins sdk.Coins) error {
	if w.EVMEngineKeeper == nil {
		return errors.New("nil EVMEngineKeeper [BUG]")
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

	gwei, dust, err := toGwei(coins[0].Amount)
	if err != nil {
		return errors.Wrap(err, "to gwei conversion")
	}
	dustCounter.Add(float64(dust))

	if gwei == 0 {
		log.Debug(ctx, "Not creating all-dust withdrawal", "addr", addr, "amount_wei", coins[0].Amount)
	} else if err := w.EVMEngineKeeper.InsertWithdrawal(ctx, addr, gwei); err != nil {
		return err
	}

	if err := w.BurnCoins(ctx, senderModule, coins); err != nil {
		return errors.Wrap(err, "burn coins")
	}

	return nil
}

// toGwei converts a math.Int to Gwei and the wei remainder.
func toGwei(amount math.Int) (gwei uint64, wei uint64, err error) { //nolint:nonamedreturns // Disambiguate identical return types.
	gweiInt := amount.QuoRaw(params.GWei)
	weiInt := amount.Sub(gweiInt.MulRaw(params.GWei))

	// This should work up to 18G ETH
	if !gweiInt.IsUint64() {
		return 0, 0, errors.New("invalid amount [BUG]")
	}

	return gweiInt.Uint64(), weiInt.Uint64(), nil
}
