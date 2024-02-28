package keeper

import (
	"context"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type proposalServer struct {
	*Keeper
	types.UnimplementedMsgServiceServer
}

// AddVotes verifies all aggregated votes included in a proposed block.
func (s proposalServer) AddVotes(ctx context.Context, msg *types.MsgAddVotes,
) (*types.AddVotesResponse, error) {
	// Verify proposed msg
	sdkContext := sdk.UnwrapSDKContext(ctx)
	vals, ok, err := s.validatorsByAddress(ctx, sdkContext.BlockHeight()-1)
	if err != nil {
		return nil, errors.Wrap(err, "fetch validators")
	} else if !ok {
		log.Warn(ctx, "Skipping vote verification since prev block validators not available. "+
			"Assume this is first block after snapshot restore", nil)
		// This is ok, since this should only occur on single validators
	} else if err := s.verifyAggVotes(ctx, vals, msg.Votes); err != nil {
		return nil, errors.Wrap(err, "verify votes")
	}

	localHeaders := headersByAddress(msg.Votes, s.voter.LocalAddress())
	logLocalVotes(ctx, localHeaders, "proposed")
	if err := s.voter.SetProposed(localHeaders); err != nil {
		return nil, errors.Wrap(err, "set committed")
	}

	return &types.AddVotesResponse{}, nil
}

// NewProposalServer returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewProposalServer(keeper *Keeper) types.MsgServiceServer {
	return &proposalServer{Keeper: keeper}
}

var _ types.MsgServiceServer = proposalServer{}
