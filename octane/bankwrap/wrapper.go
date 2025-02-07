// Package bankwrap wraps the x/bank by overriding `SendCoinsFromModuleToAccount` with
// creation of a new withdrawal request.
//
//nolint:wrapcheck // Wrapping not needed in this package.
package bankwrap

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

type Wrapper struct {
	bankkeeper.Keeper

	WithdrawalKeeper WithdrawalKeeper
}

func (k Wrapper) SendCoinsFromModuleToAccountNoWithdrawal(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	return k.Keeper.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt)
}

func (k Wrapper) SendCoinsFromModuleToAccount(ctx context.Context, _ string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	totalAmount := math.NewInt(0)
	for _, coin := range amt {
		totalAmount = totalAmount.Add(coin.Amount)
	}

	return k.WithdrawalKeeper.InsertWithdrawal(ctx, common.BytesToAddress(recipientAddr), totalAmount) //nolint:forbidigo // should be padded
}
