package keeper

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	*Keeper
	types.UnimplementedMsgServiceServer
}

// AddVotes is called with all aggregated votes included in a new finalized block.
func (s msgServer) AddVotes(ctx context.Context, msg *types.MsgAddVotes,
) (*types.AddVotesResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if sdkCtx.ExecMode() != sdk.ExecModeFinalize {
		return nil, errors.New("only allowed in finalize mode")
	}

	// Not verifying votes here since this block is now finalized, so it is too late to reject votes.

	err := s.Keeper.Add(ctx, msg)
	if err != nil {
		return nil, errors.Wrap(err, "add votes")
	}

	// Update the voter state with the local headers.
	localHeaders := headersByAddress(msg.Votes, s.voter.LocalAddress())
	logLocalVotes(ctx, localHeaders, "committed")
	if err := s.voter.SetCommitted(localHeaders); err != nil {
		return nil, errors.Wrap(err, "set committed")
	}

	return &types.AddVotesResponse{}, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServiceServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServiceServer = msgServer{}

// headersByAddress returns the attestations for the provided address.
func headersByAddress(aggregates []*types.AggVote, address common.Address) []*types.BlockHeader {
	var filtered []*types.BlockHeader
	for _, agg := range aggregates {
		for _, sig := range agg.Signatures {
			if bytes.Equal(sig.ValidatorAddress, address[:]) {
				filtered = append(filtered, agg.BlockHeader)
				break // Continue to the next aggregate.
			}
		}
	}

	return filtered
}

func logLocalVotes(ctx context.Context, headers []*types.BlockHeader, typ string) {
	if len(headers) == 0 {
		return
	}

	const limit = 5
	heights := make(map[uint64][]string)
	for _, header := range headers {
		hs := heights[header.ChainId]
		if len(hs) == limit {
			hs = append(hs, "...")
		} else if len(hs) < limit {
			hs = append(hs, strconv.FormatUint(header.Height, 10))
		} else {
			continue
		}
		heights[header.ChainId] = hs
	}
	attrs := []any{
		slog.Int("votes", len(headers)),
	}
	for cid, hs := range heights {
		attrs = append(attrs, slog.String(
			strconv.FormatUint(cid, 10),
			fmt.Sprint(hs),
		))
	}

	log.Debug(ctx, "Marked local votes as "+typ, attrs...)
}
