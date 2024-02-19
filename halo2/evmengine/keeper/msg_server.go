package keeper

import (
	"context"
	"encoding/json"

	"github.com/omni-network/omni/halo2/evmengine/types"
	engineapi "github.com/omni-network/omni/lib/engine"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/beacon/engine"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
	types.UnimplementedMsgServiceServer
}

func (s msgServer) ExecutionPayload(ctx context.Context, msg *types.MsgExecutionPayload,
) (*types.ExecutionPayloadResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if sdkCtx.ExecMode() != sdk.ExecModeFinalize {
		return nil, errors.New("only allowed in finalize mode")
	}

	fcs, err := newPayload(ctx, s.ethCl, msg)
	if err != nil {
		return nil, err
	}

	forchainResp, err := s.ethCl.ForkchoiceUpdatedV2(ctx, fcs, nil)
	if err != nil {
		return nil, err
	} else if forchainResp.PayloadStatus.Status != engine.VALID {
		return nil, errors.New("status not valid")
	}

	return &types.ExecutionPayloadResponse{}, nil
}

// newPayload creates a new payload from the given message and pushes it to the execution client.
// It returns the new forkchoice state.
func newPayload(ctx context.Context, ethCl engineapi.API, msg *types.MsgExecutionPayload,
) (engine.ForkchoiceStateV1, error) {
	var payload engine.ExecutableData
	if err := json.Unmarshal(msg.Data, &payload); err != nil {
		return engine.ForkchoiceStateV1{}, errors.Wrap(err, "unmarshal payload")
	}

	// Push it back to the execution client (mark it as possible new head).
	status, err := ethCl.NewPayloadV2(ctx, payload)
	if err != nil {
		return engine.ForkchoiceStateV1{}, errors.Wrap(err, "new payload")
	} else if status.Status != engine.VALID {
		msg := "unknown"
		if status.ValidationError != nil {
			msg = *status.ValidationError
		}

		return engine.ForkchoiceStateV1{}, errors.New("new payload invalid", "msg", msg)
	}

	return engine.ForkchoiceStateV1{
		HeadBlockHash:      payload.BlockHash,
		SafeBlockHash:      payload.BlockHash,
		FinalizedBlockHash: payload.BlockHash,
	}, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServiceServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServiceServer = msgServer{}
