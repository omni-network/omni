package keeper

import (
	"context"
	"encoding/json"

	"github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	*Keeper
	types.UnimplementedMsgServiceServer
}

// ExecutionPayload handles a new execution payload included in the current finalized block.
func (s msgServer) ExecutionPayload(ctx context.Context, msg *types.MsgExecutionPayload,
) (*types.ExecutionPayloadResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if sdkCtx.ExecMode() != sdk.ExecModeFinalize {
		return nil, errors.New("only allowed in finalize mode")
	}

	var payload engine.ExecutableData
	err := retryForever(ctx, func(ctx context.Context) (bool, error) {
		pload, status, err := pushPayload(ctx, s.engineCl, msg)
		if err != nil || isUnknown(status) {
			// We need to retry forever on networking errors, but can't easily identify them, so retry all errors.
			log.Warn(ctx, "Processing finalized payload failed: push new payload to evm (will retry)", err,
				"status", status.Status)

			return false, nil // Retry
		} else if invalid, err := isInvalid(status); invalid {
			// This should never happen. This node will stall now.
			log.Error(ctx, "Processing finalized payload failed; payload invalid [BUG]", err)

			return false, err // Don't retry, error out.
		} else if isSyncing(status) {
			log.Warn(ctx, "Processing finalized payload; evm syncing", nil)
		}

		payload = pload

		return true, nil // We are done, don't retry
	})
	if err != nil {
		return nil, err
	}

	// CometBFT has instant finality, so head/safe/finalized is latest height.
	fcs := engine.ForkchoiceStateV1{
		HeadBlockHash:      payload.BlockHash,
		SafeBlockHash:      payload.BlockHash,
		FinalizedBlockHash: payload.BlockHash,
	}

	err = retryForever(ctx, func(ctx context.Context) (bool, error) {
		fcr, err := s.engineCl.ForkchoiceUpdatedV3(ctx, fcs, nil)
		if err != nil || isUnknown(fcr.PayloadStatus) {
			// We need to retry forever on networking errors, but can't easily identify them, so retry all errors.
			log.Warn(ctx, "Processing finalized payload failed: evm fork choice update (will retry)", err,
				"status", fcr.PayloadStatus.Status)

			return false, nil // Retry
		} else if isSyncing(fcr.PayloadStatus) {
			log.Warn(ctx, "Processing finalized payload halted while evm syncing (will retry)", nil, "payload_height", payload.Number)

			return false, nil // Retry
		} else if invalid, err := isInvalid(fcr.PayloadStatus); invalid {
			// This should never happen. This node will stall now.
			log.Error(ctx, "Processing finalized payload failed; forkchoice update invalid [BUG]", err,
				"payload_height", payload.Number)

			return false, err // Don't retry
		}

		return true, nil
	})
	if err != nil {
		return nil, err
	}

	// Deliver all the previous payload log events
	if err := s.deliverEvents(ctx, payload.Number-1, payload.ParentHash, msg.PrevPayloadEvents); err != nil {
		return nil, errors.Wrap(err, "deliver event logs")
	}

	return &types.ExecutionPayloadResponse{}, nil
}

// deliverEvents delivers the given logs to the registered log providers.
func (s msgServer) deliverEvents(ctx context.Context, height uint64, blockHash common.Hash, logs []*types.EVMEvent) error {
	procs := make(map[common.Address]types.EvmEventProcessor)
	for _, proc := range s.eventProcs {
		for _, addr := range proc.Addresses() {
			procs[addr] = proc
		}
	}

	for _, evmLog := range logs {
		if err := evmLog.Verify(); err != nil {
			return errors.Wrap(err, "verify log [BUG]") // This shouldn't happen
		}

		proc, ok := procs[common.BytesToAddress(evmLog.Address)]
		if !ok {
			return errors.New("unknown log address [BUG]", log.Hex7("address", evmLog.Address))
		}

		if err := proc.Deliver(ctx, blockHash, evmLog); err != nil {
			return errors.Wrap(err, "deliver log")
		}
	}

	log.Debug(ctx, "Delivered evm logs", "height", height, "count", len(logs))

	return nil
}

// pushPayload creates a new payload from the given message and pushes it to the execution client.
// It returns the new forkchoice state and engine payload status or an error.
func pushPayload(ctx context.Context, engineCl ethclient.EngineClient, msg *types.MsgExecutionPayload,
) (engine.ExecutableData, engine.PayloadStatusV1, error) {
	var payload engine.ExecutableData
	if err := json.Unmarshal(msg.ExecutionPayload, &payload); err != nil {
		return engine.ExecutableData{}, engine.PayloadStatusV1{}, errors.Wrap(err, "unmarshal payload")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	appHash := common.BytesToHash(sdkCtx.BlockHeader().AppHash)
	if appHash == (common.Hash{}) {
		return engine.ExecutableData{}, engine.PayloadStatusV1{}, errors.New("app hash is empty")
	}

	emptyVersionHashes := make([]common.Hash, 0) // Cannot use nil.

	// Push it back to the execution client (mark it as possible new head).
	status, err := engineCl.NewPayloadV3(ctx, payload, emptyVersionHashes, &appHash)
	if err != nil {
		return engine.ExecutableData{}, engine.PayloadStatusV1{}, errors.Wrap(err, "new payload")
	}

	return payload, status, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServiceServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServiceServer = msgServer{}

func isUnknown(status engine.PayloadStatusV1) bool {
	if status.Status == engine.VALID ||
		status.Status == engine.INVALID ||
		status.Status == engine.SYNCING ||
		status.Status == engine.ACCEPTED {
		return false
	}

	return true
}

func isSyncing(status engine.PayloadStatusV1) bool {
	return status.Status == engine.SYNCING || status.Status == engine.ACCEPTED
}

func isInvalid(status engine.PayloadStatusV1) (bool, error) {
	if status.Status != engine.INVALID {
		return false, nil
	}

	valErr := "nil"
	if status.ValidationError != nil {
		valErr = *status.ValidationError
	}

	hash := "nil"
	if status.LatestValidHash != nil {
		hash = status.LatestValidHash.Hex()
	}

	return true, errors.New("payload invalid", "validation_err", valErr, "last_valid_hash", hash)
}
