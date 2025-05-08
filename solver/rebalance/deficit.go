package rebalance

import (
	"context"
	"math/big"
	"sort"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cctp"
	cctpdb "github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/netconf"
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
// USD amount is rebased to usdDecimals (6) to match USDC.
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

	return bi.Rebase(bi.MulF64(amount, price), token.Decimals, usdDecimals), nil
}

// GetUSDChainDeficit returns the total USD deficit for a given chain.
// Total Deficit = (Sum token deficits) - (Sum token surplus) - (Inflight USDC to the chain).
func GetUSDChainDeficit(
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
		sTkn, err := GetSurplus(ctx, client, token, solver)
		if err != nil {
			return nil, errors.Wrap(err, "get surplus")
		}

		if bi.LT(sTkn, GetFundThreshold(token).MinSwap()) {
			// If surplus < min swap, don't deduct from deficit.
			// We cannot use it to fill deficit.
			continue
		}

		// Convert surplus to USD
		sUSD, err := AmtToUSD(ctx, pricer, token, sTkn)
		if err != nil {
			return nil, errors.Wrap(err, "get surplus in usd")
		}

		deficit = bi.Sub(deficit, sUSD)
	}

	// If no DB, we can't get inflight USDC, so return deficit as is.
	if db == nil {
		return deficit, nil
	}

	// Subtract inflight USDC to the chain
	inflight, err := cctp.GetInflightUSDC(ctx, db, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "get inflight")
	}
	deficit = bi.Sub(deficit, inflight)

	return deficit, nil
}

type ChainAmount struct {
	ChainID uint64
	Amount  *big.Int
}

// GetUSDChainDeficits returns the total USD deficit by chain, sorted by amount descending.
func GetUSDChainDeficits(
	ctx context.Context,
	db *cctpdb.DB,
	network netconf.Network,
	clients map[uint64]ethclient.Client,
	pricer tokenpricer.Pricer,
	solver common.Address,
) ([]ChainAmount, error) {
	var deficits []ChainAmount

	for _, chain := range network.EVMChains() {
		client, ok := clients[chain.ID]
		if !ok {
			return nil, errors.New("no client", "chain", chain.ID)
		}

		deficit, err := GetUSDChainDeficit(ctx, db, client, pricer, chain.ID, solver)
		if err != nil {
			return nil, errors.Wrap(err, "get chain deficit")
		}

		deficits = append(deficits, ChainAmount{ChainID: chain.ID, Amount: deficit})
	}

	// Sort by amount descending
	sort.Slice(deficits, func(i, j int) bool { return bi.GT(deficits[i].Amount, deficits[j].Amount) })

	return deficits, nil
}
