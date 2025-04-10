package keeper

import (
	"bytes"
	"context"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	etypes "github.com/ethereum/go-ethereum/core/types"

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
	height, err := umath.ToUint64(sdk.UnwrapSDKContext(ctx).BlockHeight())
	if err != nil {
		return err
	}

	head := &ExecutionHead{
		Id:            executionHeadID,
		CreatedHeight: height,
		BlockHeight:   payload.Number,
		BlockHash:     payload.BlockHash.Bytes(),
		BlockTime:     payload.Timestamp,
	}

	err = k.headTable.Update(ctx, head)
	if err != nil {
		return errors.Wrap(err, "update execution head")
	}

	return nil
}

// InsertWithdrawal inserts a new withdrawal.
func (k *Keeper) InsertWithdrawal(ctx context.Context, withdrawalAddr common.Address, amountGwei uint64) error {
	if amountGwei == 0 {
		return errors.New("zero withdrawal amount")
	}

	height, err := umath.ToUint64(sdk.UnwrapSDKContext(ctx).BlockHeight())
	if err != nil {
		return err
	}

	err = k.withdrawalTable.Insert(ctx, &Withdrawal{
		Address:       withdrawalAddr.Bytes(),
		CreatedHeight: height,
		AmountGwei:    amountGwei,
	})
	if err != nil {
		return errors.Wrap(err, "insert withdrawal")
	}

	insertedWithdrawals.Inc()

	return nil
}

// listWithdrawalsByAddress returns all withdrawals with provided address.
func (k *Keeper) listWithdrawalsByAddress(ctx context.Context, withdrawalAddr common.Address) ([]*Withdrawal, error) {
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

// eligibleWithdrawals returns all withdrawals created below the specified height, sorted by the
// id in ascending order, limited by the configured count.
// Note we exclude the provided height, because this function is called during the proposal
// verification and execution, but in the case of the latter we also execute BeginBlockers, which
// can trigger creation of new withdrawals, that were not present during the proposal creation.
func (k *Keeper) eligibleWithdrawals(ctx context.Context, height uint64) ([]*etypes.Withdrawal, error) {
	// Note: items are ordered by the id in ascending order (oldest to newest).
	iter, err := k.withdrawalTable.List(ctx, WithdrawalPrimaryKey{})
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

		if val.GetCreatedHeight() >= height {
			// Withdrawals created in this block are not eligible
			break
		}

		withdrawals = append(withdrawals, val)

		if umath.Len(withdrawals) == k.maxWithdrawalsPerBlock {
			// Reached the max number of withdrawals
			break
		}
	}

	// This can't be nil, because the engine API would reject it otherwise.
	evmWithdrawals := []*etypes.Withdrawal{}
	for _, w := range withdrawals {
		addr, err := cast.EthAddress(w.GetAddress())
		if err != nil {
			return nil, errors.Wrap(err, "address conversion")
		}
		evmWithdrawals = append(evmWithdrawals, &etypes.Withdrawal{
			Index:   w.GetId(),
			Address: addr,
			Amount:  w.GetAmountGwei(),
			// The validator index is not used for withdrawals.
			Validator: 0,
		})
	}

	return evmWithdrawals, nil
}

// deleteWithdrawals removes all passed withdrawals by the id.
func (k *Keeper) deleteWithdrawals(ctx context.Context, withdrawals []*etypes.Withdrawal) error {
	for _, w := range withdrawals {
		err := k.withdrawalTable.Delete(ctx, &Withdrawal{Id: w.Index})
		if err != nil {
			return errors.Wrap(err, "removing withdrawal", "id", w.Index)
		}
	}

	return nil
}
