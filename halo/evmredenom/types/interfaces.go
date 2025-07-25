package types

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type EVMEngineKeeper interface {
	InsertWithdrawal(ctx context.Context, withdrawalAddr common.Address, amountWei *big.Int) error
}
