package keeper

import (
	"bytes"
	"context"
	"sort"

	"github.com/omni-network/omni/halo/attest/types"
	evmenginetypes "github.com/omni-network/omni/halo/evmengine/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/gogoproto/proto"
)

var _ evmenginetypes.VoteExtensionProvider = (*Keeper)(nil)

func (k *Keeper) PrepareVotes(ctx context.Context, commit abci.ExtendedCommitInfo) ([]sdk.Msg, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if err := baseapp.ValidateVoteExtensions(sdkCtx, k.skeeper, sdkCtx.BlockHeight(), sdkCtx.ChainID(), commit); err != nil {
		log.Error(ctx, "Cannot include invalid vote extensions in payload", err, "height", sdkCtx.BlockHeight())
		return nil, nil
	}

	msg, err := votesFromLastCommit(ctx, k.windowCompare, commit)
	if err != nil {
		return nil, err
	}

	return []sdk.Msg{msg}, nil
}

type windowCompareFunc func(context.Context, uint64, uint64) (int, error)

// votesFromLastCommit returns the aggregated votes contained in vote extensions
// of the last local commit.
func votesFromLastCommit(ctx context.Context, windowCompare windowCompareFunc, info abci.ExtendedCommitInfo) (*types.MsgAddVotes, error) {
	var allVotes []*types.Vote
	for _, vote := range info.Votes {
		if vote.BlockIdFlag != cmtproto.BlockIDFlagCommit {
			continue // Skip non block votes
		}
		votes, ok, err := votesFromExtension(vote.VoteExtension)
		if err != nil {
			return nil, err
		} else if !ok {
			continue
		}

		var selected []*types.Vote
		for _, v := range votes.Votes {
			cmp, err := windowCompare(ctx, v.BlockHeader.ChainId, v.BlockHeader.Height)
			if err != nil {
				return nil, err
			} else if cmp != 0 {
				// Skip votes that are not in the current window anymore.
				continue
			}

			selected = append(selected, v)
		}

		allVotes = append(allVotes, selected...)
	}

	return &types.MsgAddVotes{
		Authority: authtypes.NewModuleAddress(types.ModuleName).String(),
		Votes:     aggregateVotes(allVotes),
	}, nil
}

// aggregateVotes aggregates the provided attestations by block header.
func aggregateVotes(votes []*types.Vote) []*types.AggVote {
	uniqueAggs := make(map[types.UniqueKey]*types.AggVote)
	for _, vote := range votes {
		key := vote.UniqueKey()
		agg, ok := uniqueAggs[key]
		if !ok {
			agg = &types.AggVote{
				BlockHeader:     vote.BlockHeader,
				AttestationRoot: vote.AttestationRoot,
			}
		}

		agg.Signatures = append(agg.Signatures, vote.Signature)
		uniqueAggs[key] = agg
	}

	return sortAggregates(flattenAggs(uniqueAggs))
}

// flattenAggs returns the values of the provided map.
func flattenAggs(aggsByHeader map[types.UniqueKey]*types.AggVote) []*types.AggVote {
	aggs := make([]*types.AggVote, 0, len(aggsByHeader))
	for _, agg := range aggsByHeader {
		aggs = append(aggs, agg)
	}

	return aggs
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
