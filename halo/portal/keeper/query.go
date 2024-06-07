package keeper

import (
	"context"

	"github.com/omni-network/omni/halo/portal/types"
	"github.com/omni-network/omni/lib/errors"

	"cosmossdk.io/orm/types/ormerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = (*Keeper)(nil)

func (k Keeper) Block(ctx context.Context, req *types.BlockRequest) (*types.BlockResponse, error) {
	blockID := req.Id
	if req.Latest {
		var err error
		blockID, err = k.blockTable.LastInsertedSequence(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	block, msgs, err := k.getBlockAndMsgs(ctx, blockID)
	if errors.Is(err, ormerrors.NotFound) {
		return nil, status.Error(codes.NotFound, "no block found for id")
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var messages []*types.Msg
	for _, msg := range msgs {
		messages = append(messages, &types.Msg{
			Id:           msg.GetId(),
			Type:         msg.GetMsgType(),
			MsgTypeId:    msg.GetMsgTypeId(),
			DestChainId:  msg.GetDestChainId(),
			ShardId:      msg.GetShardId(),
			StreamOffset: msg.GetStreamOffset(),
		})
	}

	return &types.BlockResponse{
		Id:            block.GetId(),
		CreatedHeight: block.GetCreatedHeight(),
		Msgs:          messages,
	}, nil
}
