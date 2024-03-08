package keeper

import (
	"context"

	"github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/errors"

	"github.com/cosmos/gogoproto/proto"
)

type proposalServer struct {
	*Keeper
	types.UnimplementedMsgServiceServer
}

// ExecutionPayload handles a new execution payload proposed in a block.
func (s proposalServer) ExecutionPayload(ctx context.Context, msg *types.MsgExecutionPayload,
) (*types.ExecutionPayloadResponse, error) {
	// Push the payload to the EVM, ensuring it is valid.
	payload, err := pushPayload(ctx, s.engineCl, msg)
	if err != nil {
		return nil, err
	}

	// Collect local view of the evm logs from the previous payload.
	evmLogs, err := s.evmLogs(ctx, payload.ParentHash)
	if err != nil {
		return nil, errors.Wrap(err, "prepare evm logs")
	}

	// Ensure the proposed evm logs are equal to the local view.
	if err := evmLogsEqual(evmLogs, msg.PrevPayloadLogs); err != nil {
		return nil, errors.Wrap(err, "verify prev payload logs")
	}

	return &types.ExecutionPayloadResponse{}, nil
}

// NewProposalServer returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewProposalServer(keeper *Keeper) types.MsgServiceServer {
	return &proposalServer{Keeper: keeper}
}

var _ types.MsgServiceServer = proposalServer{}

func evmLogsEqual(a, b []*types.EVMLog) error {
	if len(a) != len(b) {
		return errors.New("count mismatch")
	}

	for i := range a {
		if !proto.Equal(a[i], b[i]) {
			return errors.New("log mismatch", "index", i)
		}
	}

	return nil
}
