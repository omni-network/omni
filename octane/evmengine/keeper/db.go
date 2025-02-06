package keeper

import (
	"bytes"
	"context"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (h *ExecutionHead) Hash() (common.Hash, error) {
	return cast.EthHash(h.GetBlockHash())
}

// executionHeadID is the ID of the singleton execution head row in the database.
const executionHeadID = 1

// InsertGenesisHead inserts the genesis execution head into the database.
func (k *Keeper) InsertGenesisHead(ctx context.Context, executionBlockHash []byte) error {
	if len(executionBlockHash) != common.HashLength {
		return errors.New("invalid execution block hash length", "length", len(executionBlockHash))
	} else if bytes.Equal(executionBlockHash, common.Hash{}.Bytes()) {
		return errors.New("invalid zero execution block hash")
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

// insertWithdrawal inserts a withdrawal request.
func (k *Keeper) InsertWithdrawal(ctx context.Context, withdrawalAddr sdk.AccAddress, amountGwei uint64) error {
	err := k.withdrawalTable.Insert(ctx, &Withdrawal{
		Address:       withdrawalAddr[:],
		CreatedHeight: uint64(sdk.UnwrapSDKContext(ctx).BlockHeight()),
		AmountGwei:    amountGwei,
	})
	if err != nil {
		return errors.Wrap(err, "insert withdrawal")
	}

	withdrawals.Inc()

	return nil
}

// GetWithdrawalsByAddress returns all stored withdrawals.
func (k *Keeper) getWithdrawalsByAddress(ctx context.Context, withdrawalAddr sdk.AccAddress) ([]*Withdrawal, error) {
	iter, err := k.withdrawalTable.List(ctx, WithdrawalAddressIndexKey{}.WithAddress(withdrawalAddr[:]))
	if err != nil {
		return nil, errors.Wrap(err, "list withdrawals")
	}
	defer iter.Close()

	var withdrawals []*Withdrawal

	for iter.Next() {
		val, err := iter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "get withdrawal")
		}
		withdrawals = append(withdrawals, val)
	}

	return withdrawals, nil
}
