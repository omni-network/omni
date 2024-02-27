package keeper

import (
	"context"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"

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
	vals := s.validatorsByAddress(ctx, sdkContext.BlockHeight()-1)
	if err := verifyAggVotes(ctx, vals, s.windower, msg.Votes); err != nil {
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
