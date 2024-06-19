package keeper

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (h *ExecutionHead) Hash() common.Hash {
	return common.BytesToHash(h.GetBlockHash())
}

// executionHeadID is the ID of the singleton execution head row in the database.
const executionHeadID = 1

// InsertGenesisHead inserts the genesis execution head into the database.
func (k *Keeper) InsertGenesisHead(ctx context.Context, executionBlockHash []byte) error {
	if len(executionBlockHash) != common.HashLength {
		return errors.New("invalid execution block hash length", "length", len(executionBlockHash))
	}

	const genesisHeight = 0 // Genesis block height is 0.

	id, err := k.headTable.InsertReturningId(ctx, &ExecutionHead{
		CreatedHeight: genesisHeight,
		BlockHeight:   genesisHeight,
		BlockHash:     executionBlockHash,
		BlockTime:     0, // Timestamp isn't critical, skip it in genesis.
	})
	if err != nil {
		return errors.Wrap(err, "insert genesis head")
	} else if id != executionHeadID {
		return errors.New("unexpected genesis head id", "id", id)
	}

	return nil
}

// getExecutionHead returns the current execution head.
func (k *Keeper) getExecutionHead(ctx context.Context) (*ExecutionHead, error) {
	head, err := k.headTable.Get(ctx, executionHeadID)
	if err != nil {
		return nil, errors.Wrap(err, "update execution head")
	}

	return head, nil
}

// updateExecutionHead updates the execution head with the given payload.
func (k *Keeper) updateExecutionHead(ctx context.Context, payload engine.ExecutableData) error {
	head := &ExecutionHead{
		Id:            executionHeadID,
		CreatedHeight: uint64(sdk.UnwrapSDKContext(ctx).BlockHeight()),
		BlockHeight:   payload.Number,
		BlockHash:     payload.BlockHash.Bytes(),
		BlockTime:     payload.Timestamp,
	}

	err := k.headTable.Update(ctx, head)
	if err != nil {
		return errors.Wrap(err, "update execution head")
	}

	return nil
}
