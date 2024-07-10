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
var _ types.ValidatorProvider = (*Keeper)(nil)

// ActiveSetByHeight returns the active cometBFT validator set at the given height.
func (k *Keeper) ActiveSetByHeight(ctx context.Context, height uint64) (*types.ValidatorSetResponse, error) {
	valset, vals, err := k.activeSetByHeight(ctx, height)
	if err != nil {
		return nil, err
	}

	return valSetResponse(valset, vals), nil
}

// ActiveSetByHeight returns the active cometBFT validator set at the given height. Zero power validators are skipped.
// Note: This MUST only be used for querying last few sets, it is inefficient otherwise.
// Note2: We could add an index, but that would be a waste of space.
func (k *Keeper) activeSetByHeight(ctx context.Context, height uint64) (*ValidatorSet, []*Validator, error) {
	setIter, err := k.valsetTable.List(ctx, ValidatorSetPrimaryKey{}, ormlist.Reverse())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to list validators")
	}
	defer setIter.Close()

	// Find the latest activated set less-than-or-equal to the given height.
	var valset *ValidatorSet
	for setIter.Next() {
		set, err := setIter.Value()
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to get validator")
		}
		if !set.GetAttested() {
			continue // Skip unattested sets.
		}
		if set.GetActivatedHeight() <= height {
			valset = set
			break
		}
	}
	if valset == nil {
		return nil, nil, errors.New("no validator set found for height")
	}

	valIter, err := k.valTable.List(ctx, ValidatorValsetIdIndexKey{}.WithValsetId(valset.GetId()))
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to list validators")
	}
	defer valIter.Close()

	var vals []*Validator
	for valIter.Next() {
		val, err := valIter.Value()
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to get validator")
		}

		if val.GetPower() == 0 {
			continue // Skip zero power validators.
		}

		vals = append(vals, val)
	}

	return valset, vals, nil
}

func (k *Keeper) ValidatorSet(ctx context.Context, req *types.ValidatorSetRequest) (*types.ValidatorSetResponse, error) {
	valsetID := req.Id
	if req.Latest {
		var err error
		valsetID, err = k.valsetTable.LastInsertedSequence(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	vatset, err := k.valsetTable.Get(ctx, valsetID)
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

		if val.GetPower() == 0 {
			continue // Skip zero power validators.
		}

		vals = append(vals, &types.Validator{
			ConsensusPubkey: val.GetPubKey(),
			Power:           val.GetPower(),
		})
	}

	return &types.ValidatorSetResponse{
		Id:              vatset.GetId(),
		CreatedHeight:   vatset.GetCreatedHeight(),
		ActivatedHeight: vatset.GetActivatedHeight(),
		Validators:      vals,
	}, nil
}

func valSetResponse(set *ValidatorSet, vals []*Validator) *types.ValidatorSetResponse {
	var validators []*types.Validator
	for _, val := range vals {
		validators = append(validators, &types.Validator{
			ConsensusPubkey: val.GetPubKey(),
			Power:           val.GetPower(),
		})
	}

	return &types.ValidatorSetResponse{
		Id:              set.GetId(),
		CreatedHeight:   set.GetCreatedHeight(),
		ActivatedHeight: set.GetActivatedHeight(),
		Validators:      validators,
	}
}
