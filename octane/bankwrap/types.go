package bankwrap

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type WithdrawalKeeper interface {
	// InsertWithdrawal creates a new withdrawal request into the local DB.
	InsertWithdrawal(ctx context.Context, withdrawalAddr sdk.AccAddress, amountGwei uint64) error
}
