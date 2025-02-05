// Package bank wraps the x/bank by overriding `SendCoinsFromModuleToAccount` with
// creation of a new withdrawal request.
//
//nolint:wrapcheck // Wrapping not needed in this package.
package bank

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

const ModuleName = "bank"

type Keeper struct {
	bankkeeper.Keeper
}

func (k Keeper) SendCoinsFromModuleToAccountNoWithdrawal(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	return k.Keeper.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt)
}

func (Keeper) SendCoinsFromModuleToAccount(context.Context, string, sdk.AccAddress, sdk.Coins) error {
	panic("unreachable")
}
