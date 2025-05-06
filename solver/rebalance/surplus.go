package rebalance

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"

	"github.com/ethereum/go-ethereum/common"
)

// GetSurplus returns surplus balance of `token` for `solver`.
func GetSurplus(
	ctx context.Context,
	client ethclient.Client,
	token tokens.Token,
	solver common.Address,
) (*big.Int, error) {
	balance, err := tokenutil.BalanceOf(ctx, client, token, solver)
	if err != nil {
		return nil, errors.Wrap(err, "get balance")
	}

	thresh := GetFundThreshold(token)

	if bi.LTE(balance, thresh.Surplus()) {
		return bi.Zero(), nil
	}

	return bi.Sub(balance, thresh.Surplus()), nil
}

// GetUSDSurplus returns surplus balance of `token` for `solver` in USD.
func GetUSDSurplus(
	ctx context.Context,
	client ethclient.Client,
	pricer tokenpricer.Pricer,
	token tokens.Token,
	solver common.Address,
) (*big.Int, error) {
	surplus, err := GetSurplus(ctx, client, token, solver)
	if err != nil {
		return nil, errors.Wrap(err, "get surplus")
	}

	return AmtToUSD(ctx, pricer, token, surplus)
}
