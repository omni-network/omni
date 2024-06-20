package keeper

import (
	"bytes"
	"context"
	"sort"

	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"
	evmenginetypes "github.com/omni-network/omni/octane/evmengine/types"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/gogoproto/proto"
)

var _ evmenginetypes.VoteExtensionProvider = (*Keeper)(nil)

// PrepareVotes returns the cosmosSDK transaction MsgAddVotes that will include all the validator votes included
// in the previous block's vote extensions into the attest module.
//
// Note that the commit is trusted to be valid and only contains valid VEs from the previous block as
// provided by a trusted cometBFT.
func (k *Keeper) PrepareVotes(ctx context.Context, commit abci.ExtendedCommitInfo) ([]sdk.Msg, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if err := baseapp.ValidateVoteExtensions(sdkCtx, k.skeeper, sdkCtx.BlockHeight(), sdkCtx.ChainID(), commit); err != nil {
		log.Error(ctx, "Cannot include invalid vote extensions in payload", err, "height", sdkCtx.BlockHeight())
		return nil, nil
	}

	msg, err := votesFromLastCommit(
		ctx,
		k.windowCompare,
		k.portalRegistry.SupportedChain,
		commit,
	)
	if err != nil {
		return nil, err
	}

	return []sdk.Msg{msg}, nil
}

type windowCompareFunc func(context.Context, xchain.ChainVersion, uint64) (int, error)
type supportedChainFunc func(context.Context, uint64) (bool, error)

// votesFromLastCommit returns the aggregated votes contained in vote extensions
// of the last local commit.
func votesFromLastCommit(
	ctx context.Context,
	windowCompare windowCompareFunc,
	supportedChain supportedChainFunc,
	info abci.ExtendedCommitInfo,

) (*types.MsgAddVotes, error) {
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
			if ok, err := supportedChain(ctx, v.BlockHeader.ChainId); err != nil {
				return nil, err
			} else if !ok {
				// Skip votes for unsupported chains.
				continue
			}

			cmp, err := windowCompare(ctx, v.BlockHeader.XChainVersion(), v.BlockHeader.Offset)
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

	votes, err := aggregateVotes(allVotes)
	if err != nil {
		return nil, err
	}

	return &types.MsgAddVotes{
		Authority: authtypes.NewModuleAddress(types.ModuleName).String(),
		Votes:     votes,
	}, nil
}

// aggregateVotes aggregates the provided attestations by block header.
func aggregateVotes(votes []*types.Vote) ([]*types.AggVote, error) {
	uniqueAggs := make(map[[32]byte]*types.AggVote)
	for _, vote := range votes {
		attRoot, err := vote.AttestationRoot()
		if err != nil {
			return nil, err
		}
		agg, ok := uniqueAggs[attRoot]
		if !ok {
			agg = &types.AggVote{
				BlockHeader: vote.BlockHeader,
				MsgRoot:     vote.MsgRoot,
			}
		}

		agg.Signatures = append(agg.Signatures, vote.Signature)
		uniqueAggs[attRoot] = agg
	}

	return sortAggregates(flattenAggs(uniqueAggs)), nil
}

// flattenAggs returns the values of the provided map.
func flattenAggs(aggsByHeader map[[32]byte]*types.AggVote) []*types.AggVote {
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
		if aggs[i].BlockHeader.Offset != aggs[j].BlockHeader.Offset {
			return aggs[i].BlockHeader.Offset < aggs[j].BlockHeader.Offset
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
