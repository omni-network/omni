package keeper

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum/common"

	"cosmossdk.io/math"
)

var _ types.QueryServer = (*Keeper)(nil)

func (k *Keeper) SumPendingWithdrawalsByAddress(
	ctx context.Context,
	req *types.SumPendingWithdrawalsByAddressRequest,
) (*types.SumPendingWithdrawalsByAddressResponse, error) {
	if req == nil {
		return nil, errors.New("nil request")
	}
	withdrawals, err := k.getWithdrawalsByAddress(ctx, common.BytesToAddress(req.Address)) //nolint:forbidigo // should be padded
	if err != nil {
		return nil, errors.Wrap(err, "get withdrawals")
	}

	amount := math.NewInt(0)

	for _, w := range withdrawals {
		amount = amount.Add(math.NewIntFromUint64(w.GetAmountGwei()))
	}

	return &types.SumPendingWithdrawalsByAddressResponse{Amount: amount.Uint64()}, nil
}
