package keeper

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	"cosmossdk.io/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// getOrCreateEpoch return the unattested epoch created at the current block height or creates a new one
// using the latest valset and network as base.
func (k *Keeper) getOrCreateEpoch(ctx context.Context) (*Epoch, error) {
	createHeight := uint64(sdk.UnwrapSDKContext(ctx).BlockHeight())

	epoch, err := k.epochTable.GetByAttestedCreatedHeight(ctx, false, createHeight)
	if ormerrors.IsNotFound(err) {
		// Create a new epoch using the latest valset and network as base
		networkID, err := k.networkTable.LastInsertedSequence(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get last network ID")
		}

		valsetID, err := k.valsetTable.LastInsertedSequence(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "get last valset ID")
		}

		epoch = &Epoch{
			CreatedHeight: createHeight,
			Attested:      false,
			NetworkId:     networkID,
			ValsetId:      valsetID,
		}
		id, err := k.epochTable.InsertReturningId(ctx, epoch)
		if err != nil {
			return nil, errors.Wrap(err, "insert epoch")
		}

		epoch.Id = id
	} else if err != nil {
		return nil, errors.Wrap(err, "get epoch")
	}

	return epoch, nil
}
