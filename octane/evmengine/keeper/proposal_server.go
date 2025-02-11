package keeper

import (
	"context"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/octane/evmengine/types"

	"github.com/cosmos/gogoproto/proto"
)

type proposalServer struct {
	*Keeper
	types.UnimplementedMsgServiceServer
}

// ExecutionPayload handles a new execution payload proposed in a block.
func (s proposalServer) ExecutionPayload(ctx context.Context, msg *types.MsgExecutionPayload,
) (*types.ExecutionPayloadResponse, error) {
	payload, err := s.parseAndVerifyProposedPayload(ctx, msg)
	if err != nil {
		return nil, err
	}

	blockHashes, err := blobHashes(msg.BlobCommitments)
	if err != nil {
		return nil, errors.Wrap(err, "blob commitments")
	}

	// Block of magellan network upgrade height is built by uluwatu logic,
	// just ensure events are valid, but log that we are ignoring/dropping them.
	if len(msg.PrevPayloadEvents) > 0 {
		// This is only supported along with legacy payloads
		if len(msg.ExecutionPayload) == 0 {
			return nil, errors.New("prev payload events only supported with legacy payloads")
		}

		// Collect local view of the evm logs from the previous payload.
		evmEvents, err := s.legacyEVMEvents(ctx, payload.ParentHash)
		if err != nil {
			return nil, errors.Wrap(err, "prepare evm event logs")
		}

		// Ensure the proposed evm event logs are equal to the local view.
		if err := evmEventsEqual(evmEvents, msg.PrevPayloadEvents); err != nil {
			return nil, errors.Wrap(err, "verify prev payload events")
		}

		log.Warn(ctx, "Ignoring legacy proposed prev payload events", nil, "count", len(msg.PrevPayloadEvents))
	}

	// Push the payload to the EVM.
	err = retryForever(ctx, func(ctx context.Context) (bool, error) {
		status, err := pushPayload(ctx, s.engineCl, payload, blockHashes)
		if err != nil {
			// We need to retry forever on networking errors, but can't easily identify them, so retry all errors.
			log.Warn(ctx, "Verifying proposal failed: push new payload to evm (will retry)", err)

			return false, nil // Retry
		} else if invalid, err := isInvalid(status); invalid {
			return false, errors.Wrap(err, "invalid payload, rejecting proposal") // Abort, don't retry
		} else if isSyncing(status) {
			// If this is initial sync, we need to continue and set a target head to sync to, so don't retry.
			log.Warn(ctx, "Can't properly verifying proposal: evm syncing", err,
				"payload_height", payload.Number)
		} /* else isValid(status) */

		return true, nil // Done
	})
	if err != nil {
		return nil, err
	}

	return &types.ExecutionPayloadResponse{}, nil
}

// NewProposalServer returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewProposalServer(keeper *Keeper) types.MsgServiceServer {
	return &proposalServer{Keeper: keeper}
}

var _ types.MsgServiceServer = proposalServer{}

func evmEventsEqual(a, b []types.EVMEvent) error {
	if len(a) != len(b) {
		return errors.New("count mismatch", "a", len(a), "b", len(b))
	}

	for i := range a {
		if !proto.Equal(&a[i], &b[i]) {
			return errors.New("log mismatch", "index", i)
		}
	}

	return nil
}
