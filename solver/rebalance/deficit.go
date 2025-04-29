package rebalance

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
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
