package keeper

import (
	"context"

	"github.com/omni-network/omni/halo2/evmengine/types"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/beacon/engine"
)

type msgServer struct {
	Keeper
	types.UnimplementedMsgServiceServer
}

func (s msgServer) FinalisePayload(ctx context.Context, msg *types.MsgExecutionPayload,
) (*types.FinalisePayloadResponse, error) {
	var payload engine.ExecutableData
	if err := decode(msg.Data, &payload); err != nil {
		return nil, err
	}

	fcs := engine.ForkchoiceStateV1{
		HeadBlockHash:      payload.BlockHash,
		SafeBlockHash:      payload.BlockHash,
		FinalizedBlockHash: payload.BlockHash,
	}

	forchainResp, err := s.ethCl.ForkchoiceUpdatedV2(ctx, fcs, nil)
	if err != nil {
		return nil, err
	} else if forchainResp.PayloadStatus.Status != engine.VALID {
		return nil, errors.New("status not valid")
	}

	return &types.FinalisePayloadResponse{}, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServiceServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServiceServer = msgServer{}
