package rebalance

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cctp"
	cctpdb "github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"

	"github.com/ethereum/go-ethereum/common"
)

// GetDeficit returns deficit balance of `token` for `solver`.
func GetDeficit(
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

	if bi.GTE(balance, thresh.Target()) {
		// Balance > target, no deficit.
		return bi.Zero(), nil
	}

	return bi.Sub(thresh.Target(), balance), nil
}

// GetUSDDeficit returns deficit balance of `token` for `solver` in USD (rebaces to 6 decimals).
func GetUSDDeficit(
	ctx context.Context,
	client ethclient.Client,
	pricer tokenpricer.Pricer,
	token tokens.Token,
	solver common.Address,
) (*big.Int, error) {
	deficit, err := GetDeficit(ctx, client, token, solver)
	if err != nil {
		return nil, errors.Wrap(err, "get deficit")
	}

	return AmtToUSD(ctx, pricer, token, deficit)
}

// AmtToUSD converts a token amount to USD using the given price.
// USD amount is rebased to 6 decimals.
func AmtToUSD(
	ctx context.Context,
	pricer tokenpricer.Pricer,
	token tokens.Token,
	amount *big.Int,
) (*big.Int, error) {
	price, err := pricer.USDPrice(ctx, token.Asset)
	if err != nil {
		return nil, errors.Wrap(err, "get price")
	}

	return bi.Rebase(bi.MulF64(amount, price), token.Decimals, 6), nil
}

// GetChainUSDDeficit returns the total USD deficit for a given chain.
// Total Deficit = (Sum token deficits) - (Sum token surplus) - (Inflight USDC to the chain).
func GetChainUSDDeficit(
	ctx context.Context,
	db *cctpdb.DB,
	client ethclient.Client,
	pricer tokenpricer.Pricer,
	chainID uint64,
	solver common.Address,
) (*big.Int, error) {
	deficit := bi.Zero()

	// Add up all the token deficits
	for _, token := range TokensByChain(chainID) {
		d, err := GetUSDDeficit(ctx, client, pricer, token, solver)
		if err != nil {
			return nil, errors.Wrap(err, "get deficit")
		}
		deficit = bi.Add(deficit, d)
	}

	// Subtract all the token surpluses
	for _, token := range TokensByChain(chainID) {
		s, err := GetUSDSurplus(ctx, client, pricer, token, solver)
		if err != nil {
			return nil, errors.Wrap(err, "get surplus")
		}

		deficit = bi.Sub(deficit, s)
	}

	// Subtract inflight USDC to the chain
	inflight, err := cctp.GetInflightUSDC(ctx, db, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "get inflight")
	}
	deficit = bi.Sub(deficit, inflight)

	return deficit, nil
}
