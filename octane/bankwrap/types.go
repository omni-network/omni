package bankwrap

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	"cosmossdk.io/math"
)

type WithdrawalKeeper interface {
	// InsertWithdrawal creates a new withdrawal request into the local DB.
	InsertWithdrawal(ctx context.Context, withdrawalAddr common.Address, amountGwei math.Int) error
}
