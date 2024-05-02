package contract

import (
	"context"
	"math/big"
)

type FeeOracleV1 interface {
	SetGasPriceOn(ctx context.Context, destChainID uint64, gasPrice *big.Int) error
	GasPriceOn(ctx context.Context, destChainID uint64) (*big.Int, error)

	SetToNativeRate(ctx context.Context, destChainID uint64, rate *big.Int) error
	ToNativeRate(ctx context.Context, destChainID uint64) (*big.Int, error)
}
