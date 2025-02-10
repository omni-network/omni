package withdraw

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
)

type EVMEngineKeeper interface {
	// InsertWithdrawal creates a new withdrawal request.
	InsertWithdrawal(ctx context.Context, withdrawalAddr common.Address, amountGwei uint64) error
}
