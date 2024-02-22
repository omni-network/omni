package keeper

import (
	"bytes"
	"context"
	"sort"

	"github.com/omni-network/omni/halo/attest/types"
	evmenginetypes "github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/gogoproto/proto"
)

var _ evmenginetypes.CPayloadProvider = (*Keeper)(nil)

func (k *Keeper) PreparePayload(ctx context.Context, height uint64, commit abci.ExtendedCommitInfo) ([]sdk.Msg, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if err := baseapp.ValidateVoteExtensions(sdkCtx, k.skeeper, int64(height), sdkCtx.ChainID(), commit); err != nil {
		log.Error(ctx, "Cannot include invalid vote extensions in payload", err, "height", height)
		return nil, nil
	}

	msg, ok, err := votesFromLastCommit(commit)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, nil
	}

	return []sdk.Msg{msg}, nil
}

// votesFromLastCommit returns the aggregated votes contained in vote extensions
// of the last local commit.
func votesFromLastCommit(info abci.ExtendedCommitInfo) (*types.MsgAddVotes, bool, error) {
	var allVotes []*types.Vote
	for _, vote := range info.Votes {
		if vote.BlockIdFlag != cmtproto.BlockIDFlagCommit {
			continue // Skip non block votes
		}
		votes, ok, err := votesFromExtension(vote.VoteExtension)
		if err != nil {
			return nil, false, err
		} else if !ok {
			continue
		}

		allVotes = append(allVotes, votes.Votes...)
	}

	return &types.MsgAddVotes{
		Authority: authtypes.NewModuleAddress(types.ModuleName).String(),
		Votes:     aggregateVotes(allVotes),
	}, len(allVotes) > 0, nil
}

// aggregateVotes aggregates the provided attestations by block header.
func aggregateVotes(votes []*types.Vote) []*types.AggVote {
	aggsByHeader := make(map[xchain.BlockHeader]*types.AggVote) // map[BlockHash]AggAttestation
	for _, vote := range votes {
		header := vote.BlockHeader.ToXChain()
		agg, ok := aggsByHeader[header]
		if !ok {
			agg = &types.AggVote{
				BlockHeader: vote.BlockHeader,
				BlockRoot:   vote.BlockRoot,
			}
		}

		agg.Signatures = append(agg.Signatures, vote.Signature)
		aggsByHeader[header] = agg
	}

	return flattenAggsByHeader(aggsByHeader)
}

// flattenAggsByHeader returns the provided map of aggregates by header as a slice in a deterministic order.
func flattenAggsByHeader(aggsByHeader map[xchain.BlockHeader]*types.AggVote) []*types.AggVote {
	aggs := make([]*types.AggVote, 0, len(aggsByHeader))
	for _, agg := range aggsByHeader {
		aggs = append(aggs, agg)
	}

	return sortAggregates(aggs)
}

// sortAggregates returns the provided aggregates in a deterministic order.
// Note the provided slice is also sorted in-place.
func sortAggregates(aggs []*types.AggVote) []*types.AggVote {
	sort.Slice(aggs, func(i, j int) bool {
		if aggs[i].BlockHeader.Height != aggs[j].BlockHeader.Height {
			return aggs[i].BlockHeader.Height < aggs[j].BlockHeader.Height
		}
		if aggs[i].BlockHeader.ChainId != aggs[j].BlockHeader.ChainId {
			return aggs[i].BlockHeader.ChainId < aggs[j].BlockHeader.ChainId
		}

		return bytes.Compare(aggs[i].BlockHeader.Hash, aggs[j].BlockHeader.Hash) < 0
	})

	return aggs
}

// votesFromExtension returns the attestations contained in the vote extension, or false if none or an error.
func votesFromExtension(voteExtension []byte) (*types.Votes, bool, error) {
	if len(voteExtension) == 0 {
		return nil, false, nil
	}

	resp := new(types.Votes)
	if err := proto.Unmarshal(voteExtension, resp); err != nil {
		return nil, false, errors.Wrap(err, "decode vote extension")
	}

	return resp, true, nil
}
