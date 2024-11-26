package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type AuthKeeper interface {
	HasAccount(_ context.Context, _ sdk.AccAddress) bool
	NewAccountWithAddress(_ context.Context, _ sdk.AccAddress) sdk.AccountI
	SetAccount(_ context.Context, _ sdk.AccountI)
}

type BankKeeper interface {
	MintCoins(_ context.Context, _ string, _ sdk.Coins) error
	SendCoinsFromModuleToAccount(_ context.Context, _ string, _ sdk.AccAddress, _ sdk.Coins) error
}

type StakingKeeper interface {
	GetValidator(_ context.Context, _ sdk.ValAddress) (stypes.Validator, error)
}

type StakingMsgServer interface {
	CreateValidator(_ context.Context, _ *stypes.MsgCreateValidator) (*stypes.MsgCreateValidatorResponse, error)
	Delegate(_ context.Context, _ *stypes.MsgDelegate) (*stypes.MsgDelegateResponse, error)
}
