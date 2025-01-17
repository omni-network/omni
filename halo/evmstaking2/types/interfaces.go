//go:generate mockgen -source ./interfaces.go -package testutil -destination ../testutil/mock_interfaces.go
package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type AuthKeeper interface {
	HasAccount(ctx context.Context, addr sdk.AccAddress) bool
	NewAccountWithAddress(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	SetAccount(ctx context.Context, acc sdk.AccountI)
}

type BankKeeper interface {
	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromModuleToAccountForReal(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

type StakingKeeper interface {
	GetValidator(ctx context.Context, addr sdk.ValAddress) (stypes.Validator, error)
}

type StakingMsgServer interface {
	CreateValidator(ctx context.Context, msg *stypes.MsgCreateValidator) (*stypes.MsgCreateValidatorResponse, error)
	Delegate(ctx context.Context, msg *stypes.MsgDelegate) (*stypes.MsgDelegateResponse, error)
}
