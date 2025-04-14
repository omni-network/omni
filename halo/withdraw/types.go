package withdraw

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type EVMEngineKeeper interface {
	// InsertWithdrawal creates a new withdrawal request.
	InsertWithdrawal(ctx context.Context, withdrawalAddr common.Address, amountGwei uint64) error
}

type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	GetModuleAccount(ctx context.Context, moduleName string) sdk.ModuleAccountI
}
