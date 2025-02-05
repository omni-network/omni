// Package evmupgrade monitors the Upgrade pre-deploy contract and converts
// its log events to cosmosSDK x/upgrade logic.
//
//nolint:wrapcheck // Wrapping not needed in this package.
package mybank

import (
	"context"

	"github.com/omni-network/omni/lib/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

const ModuleName = "mybank"

type Keeper struct {
	bankkeeper.Keeper
}

func (k Keeper) SendCoinsFromModuleToAccountForReal(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	log.Info(ctx, "Wrapped method called")

	return k.Keeper.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt)
}

func (Keeper) SendCoinsFromModuleToAccount(context.Context, string, sdk.AccAddress, sdk.Coins) error {
	panic("unreachable")
}
