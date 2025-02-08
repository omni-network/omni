package keeper

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/octane/evmengine/types"

	"github.com/ethereum/go-ethereum/common"
)

var _ types.QueryServer = (*Keeper)(nil)

func (k *Keeper) SumPendingWithdrawalsByAddress(
	ctx context.Context,
	req *types.SumPendingWithdrawalsByAddressRequest,
) (*types.SumPendingWithdrawalsByAddressResponse, error) {
	if req == nil {
		return nil, errors.New("nil request")
	}

	withdrawals, err := k.listWithdrawalsByAddress(ctx, common.Address(req.Address))
	if err != nil {
		return nil, errors.Wrap(err, "get withdrawals")
	}

	var sumGwei uint64
	for _, w := range withdrawals {
		sumGwei += w.GetAmountGwei()
	}

	return &types.SumPendingWithdrawalsByAddressResponse{SumGwei: sumGwei}, nil
}
