package keeper

import (
	"context"

	"github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/errors"

	"cosmossdk.io/orm/types/ormerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = (*Keeper)(nil)

//nolint:wrapcheck // This is a grpc server method, so has to return grpc errors.
func (k Keeper) ValidatorSet(ctx context.Context, req *types.ValidatorSetRequest) (*types.ValidatorSetResponse, error) {
	vatset, err := k.valsetTable.Get(ctx, req.Id)
	if errors.Is(err, ormerrors.NotFound) {
		return nil, status.Error(codes.NotFound, "no validator set found for id")
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	valIter, err := k.valTable.List(ctx, ValidatorValsetIdAddressIndexKey{}.WithValsetId(vatset.GetId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer valIter.Close()

	var vals []*types.Validator
	for valIter.Next() {
		val, err := valIter.Value()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		vals = append(vals, &types.Validator{
			Address: val.GetAddress(),
			Power:   val.GetPower(),
		})
	}

	return &types.ValidatorSetResponse{
		Id:            vatset.GetId(),
		CreatedHeight: vatset.GetCreatedHeight(),
		Validators:    vals,
	}, nil
}
