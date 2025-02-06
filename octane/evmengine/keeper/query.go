package keeper

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/octane/evmengine/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = (*Keeper)(nil)
var _ types.WithdrawalsProvider = (*Keeper)(nil)

func (k *Keeper) SumPendingWithdrawalsByAddress(
	ctx context.Context,
	in *types.SumPendingWithdrawalsByAddressRequest,
) (*types.SumPendingWithdrawalsByAddressResponse, error) {
	if in == nil {
		return nil, errors.New("no address")
	}
	withdrawals, err := k.getWithdrawalsByAddress(ctx, sdk.AccAddress(in.Address))
	if err != nil {
		return nil, errors.Wrap(err, "get withdrawals")
	}

	var amount uint64

	for _, w := range withdrawals {
		amount += w.GetAmountGwei()
	}

	return &types.SumPendingWithdrawalsByAddressResponse{Amount: amount}, nil
}
