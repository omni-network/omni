package keeper

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

type proposalServer struct {
	Keeper
	types.UnimplementedMsgServiceServer
}

// AddAggAttestations handles aggregate attestations proposed in a block.
func (s proposalServer) AddAggAttestations(ctx context.Context, msg *types.MsgAggAttestations,
) (*types.AddAggAttestationsResponse, error) {
	localHeaders := headersByAddress(msg.Aggregates, s.attester.LocalAddress())
	if err := s.attester.SetProposed(localHeaders); err != nil {
		return nil, errors.Wrap(err, "set committed")
	}

	if len(msg.Aggregates) == 0 {
		return &types.AddAggAttestationsResponse{}, nil
	}

	// Make nice logs
	heights := make(map[uint64][]uint64)
	for _, header := range localHeaders {
		heights[header.ChainId] = append(heights[header.ChainId], header.Height)
	}
	attrs := []any{
		slog.Int("attestations", len(localHeaders)),
		log.Hex7("validator", s.attester.LocalAddress().Bytes()),
	}
	for cid, hs := range heights {
		attrs = append(attrs, slog.String(
			strconv.FormatUint(cid, 10),
			fmt.Sprint(hs),
		))
	}

	log.Debug(ctx, "Marked local attestations as proposed", attrs...)

	return &types.AddAggAttestationsResponse{}, nil
}

// NewProposalServer returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewProposalServer(keeper Keeper) types.MsgServiceServer {
	return &proposalServer{Keeper: keeper}
}

var _ types.MsgServiceServer = proposalServer{}
