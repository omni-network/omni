package withdraw

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type EVMEngineKeeper interface {
	// InsertWithdrawal creates a new withdrawal request.
	// Note the amount is the native EVM token amount in wei.
	// Withdrawals are rounded to gwei, so small amounts result in noop.
	InsertWithdrawal(ctx context.Context, withdrawalAddr common.Address, weiAmount *big.Int) error
}

type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	GetModuleAccount(ctx context.Context, moduleName string) sdk.ModuleAccountI
}
