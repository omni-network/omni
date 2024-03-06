package keeper

import (
	"context"

	"github.com/omni-network/omni/lib/errors"

	"golang.org/x/exp/maps"
)

// processApprovedReceipts processes all receipts for the given approved attestation.
// It finds matching source chain msg offsets (and their attestation IDs) and marks them as submitted.
// It returns updated attestation IDs.
func (k *Keeper) processApprovedReceipts(ctx context.Context, attID uint64) ([]uint64, error) {
	// Query all receipts for the given attestation
	attIDIdx := ReceiptOffsetAttIdIndexKey{}.WithAttId(attID)
	receiptIter, err := k.receiptOffsetTable.List(ctx, attIDIdx)
	if err != nil {
		return nil, errors.Wrap(err, "list receipts")
	}
	defer receiptIter.Close()

	// For each receipt, update all equaled or lower msg offsets as submitted.
	// Also collect all updated attestations.
	uniqAttIDs := make(map[uint64]struct{})
	for receiptIter.Next() {
		receipt, err := receiptIter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "receipt value")
		}

		attIDs, err := k.updateMsgsSubmitted(ctx, receipt)
		if err != nil {
			return nil, errors.Wrap(err, "mark msgs submitted")
		}

		for _, attID := range attIDs {
			uniqAttIDs[attID] = struct{}{}
		}
	}

	return maps.Keys(uniqAttIDs), nil
}

// updateMsgsSubmitted updates all equaled-or-lower msg offsets as submitted for the given approved receipt.
// It returns the associated attestation IDs for the updated msg offsets.
func (k *Keeper) updateMsgsSubmitted(ctx context.Context, receipt *ReceiptOffset) ([]uint64, error) {
	msgIdxFrom, msgIdxTo := msgIndexesForReceipt(receipt)
	msgIter, err := k.msgOffsetTable.ListRange(ctx, msgIdxFrom, msgIdxTo)
	if err != nil {
		return nil, errors.Wrap(err, "list msgs")
	}
	defer msgIter.Close()

	var attIDs []uint64
	for msgIter.Next() {
		msg, err := msgIter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "msg value")
		}

		// Mark the message as submitted.
		msg.Submitted = true

		if err := k.msgOffsetTable.Update(ctx, msg); err != nil {
			return nil, errors.Wrap(err, "update msg")
		}

		attIDs = append(attIDs, msg.GetAttId())
	}

	return attIDs, nil
}

func (k *Keeper) maybeUpdateAttsSubmitted(ctx context.Context, attIDs []uint64) error {
	for _, attID := range attIDs {
		// Check if all msg offsets are submitted
		ok, err := k.IsAllMsgsSubmitted(ctx, attID)
		if err != nil {
			return errors.Wrap(err, "check all msgs submitted")
		} else if !ok {
			continue
		}

		// Mark the attestation as submitted
		att, err := k.attTable.Get(ctx, attID)
		if err != nil {
			return errors.Wrap(err, "get att")
		}

		att.Status = int32(Status_Submitted)

		if err := k.attTable.Update(ctx, att); err != nil {
			return errors.Wrap(err, "update att")
		}
	}

	return nil
}

// IsAllMsgsSubmitted returns true if all messages for the given attestation ID are submitted.
func (k *Keeper) IsAllMsgsSubmitted(ctx context.Context, attID uint64) (bool, error) {
	msgIdx := MsgOffsetAttIdIndexKey{}.WithAttId(attID)
	msgIter, err := k.msgOffsetTable.List(ctx, msgIdx)
	if err != nil {
		return false, errors.Wrap(err, "list msgs")
	}
	defer msgIter.Close()

	for msgIter.Next() {
		msg, err := msgIter.Value()
		if err != nil {
			return false, errors.Wrap(err, "msg value")
		}

		if !msg.GetSubmitted() {
			return false, nil
		}
	}

	return true, nil
}

func msgIndexesForReceipt(receipt *ReceiptOffset) (MsgOffsetSubmittedSourceChainIdDestChainIdStreamOffsetIndexKey, MsgOffsetSubmittedSourceChainIdDestChainIdStreamOffsetIndexKey) {
	return MsgOffsetSubmittedSourceChainIdDestChainIdStreamOffsetIndexKey{}.
			WithSubmittedSourceChainIdDestChainIdStreamOffset(
				false,                      // Not submitted yet.
				receipt.GetSourceChainId(), // For the given source chain.
				receipt.GetDestChainId(),   // For the given dest chain.
				0,                          // Query from 0
			),
		MsgOffsetSubmittedSourceChainIdDestChainIdStreamOffsetIndexKey{}.
			WithSubmittedSourceChainIdDestChainIdStreamOffset(
				false,                      // Not submitted yet.
				receipt.GetSourceChainId(), // For the given source chain.
				receipt.GetDestChainId(),   // For the given dest chain.
				receipt.GetStreamOffset(),  // Query up to the stream offset of the receipt.
			)
}
