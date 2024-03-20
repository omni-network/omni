package keeper

import (
	"context"

	"github.com/omni-network/omni/halo/valsync/types"
	"github.com/omni-network/omni/lib/errors"

	"cosmossdk.io/orm/model/ormlist"
	"cosmossdk.io/orm/types/ormerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = (*Keeper)(nil)

// ValidatorsAtHeight returns the validators at the given height.
// Note: This MUST only be used for querying last few sets, it is inefficient otherwise.
func (k Keeper) ValidatorsAtHeight(ctx context.Context, height uint64) ([]*Validator, error) {
	setIter, err := k.valsetTable.List(ctx, ValidatorSetPrimaryKey{}, ormlist.Reverse())
	if err != nil {
		return nil, errors.Wrap(err, "failed to list validators")
	}
	defer setIter.Close()

	// Find the latest set less-than-or-equal to the given height.
	var valsetID uint64
	for setIter.Next() {
		set, err := setIter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "failed to get validator")
		}
		if set.GetCreatedHeight() <= height {
			valsetID = set.GetId()
			break
		}
	}
	if valsetID == 0 {
		return nil, errors.New("no validator set found for height")
	}

	valIter, err := k.valTable.List(ctx, ValidatorValsetIdIndexKey{}.WithValsetId(valsetID))
	if err != nil {
		return nil, errors.Wrap(err, "failed to list validators")
	}
	defer valIter.Close()

	var vals []*Validator
	for valIter.Next() {
		val, err := valIter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "failed to get validator")
		}

		vals = append(vals, val)
	}

	return vals, nil
}

//nolint:wrapcheck // This is a grpc server method, so has to return grpc errors.
func (k Keeper) ValidatorSet(ctx context.Context, req *types.ValidatorSetRequest) (*types.ValidatorSetResponse, error) {
	vatset, err := k.valsetTable.Get(ctx, req.Id)
	if errors.Is(err, ormerrors.NotFound) {
		return nil, status.Error(codes.NotFound, "no validator set found for id")
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	valIter, err := k.valTable.List(ctx, ValidatorValsetIdIndexKey{}.WithValsetId(vatset.GetId()))
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

		addr, err := val.Address()
		if err != nil {
			return nil, err
		}

		vals = append(vals, &types.Validator{
			Address: addr.Bytes(),
			Power:   val.GetPower(),
		})
	}

	return &types.ValidatorSetResponse{
		Id:            vatset.GetId(),
		CreatedHeight: vatset.GetCreatedHeight(),
		Validators:    vals,
	}, nil
}
