package attester

import (
	"bytes"
	"context"
	"sort"

	"github.com/omni-network/omni/halo2/attest/types"
	evmenginetypes "github.com/omni-network/omni/halo2/evmengine/types"
	"github.com/omni-network/omni/lib/errors"

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/gogoproto/proto"
)

var _ evmenginetypes.CPayloadProvider = CPayloadProvider{}

// CPayloadProvider implements the CPayloadProvider interface for the attester module.
// It extracts the aggregate attestations from the vote extensions of the last local commit.
type CPayloadProvider struct {
}

func (CPayloadProvider) PreparePayload(_ context.Context, _ uint64, commit abci.ExtendedCommitInfo) ([]sdk.Msg, error) {
	aggAtts, ok, err := aggregatesFromLastCommit(commit)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, nil
	}

	return []sdk.Msg{aggAtts}, nil
}

// aggregatesFromLastCommit returns the aggregateAtts attestations contained in vote extensions
// of the last local commit.
func aggregatesFromLastCommit(info abci.ExtendedCommitInfo) (*types.MsgAggAttestations, bool, error) {
	var allAtts []*types.Attestation
	for _, vote := range info.Votes {
		atts, ok, err := attestationsFromVoteExt(vote.VoteExtension)
		if err != nil {
			return nil, false, err
		} else if !ok {
			continue
		}

		allAtts = append(allAtts, atts.Attestations...)
	}

	return &types.MsgAggAttestations{
		Authority:  authtypes.NewModuleAddress(types.ModuleName).String(),
		Aggregates: aggregateAtts(allAtts),
	}, len(allAtts) > 0, nil
}

// aggregateAtts aggregates the provided attestations by block header.
func aggregateAtts(atts []*types.Attestation) []*types.AggAttestation {
	aggsByHash := make(map[string]*types.AggAttestation) // map[BlockHash]AggAttestation
	for _, att := range atts {
		hashStr := string(att.BlockHeader.Hash)
		agg, ok := aggsByHash[hashStr]
		if !ok {
			agg = &types.AggAttestation{
				BlockHeader: att.BlockHeader,
				BlockRoot:   att.BlockRoot,
			}
		}

		agg.Signatures = append(agg.Signatures, att.Signature)
		aggsByHash[hashStr] = agg
	}

	return flattenAggsByHeader(aggsByHash)
}

// flattenAggsByHeader returns the provided map of aggregates by header as a slice in a deterministic order.
func flattenAggsByHeader(aggsByHeader map[string]*types.AggAttestation) []*types.AggAttestation {
	aggs := make([]*types.AggAttestation, 0, len(aggsByHeader))
	for _, agg := range aggsByHeader {
		aggs = append(aggs, agg)
	}

	return sortAggregates(aggs)
}

// sortAggregates returns the provided aggregates in a deterministic order.
// Note the provided slice is also sorted in-place.
func sortAggregates(aggs []*types.AggAttestation) []*types.AggAttestation {
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

// attestationsFromVoteExt returns the attestations contained in the vote extension, or false if none or an error.
func attestationsFromVoteExt(voteExtension []byte) (*types.Attestations, bool, error) {
	if len(voteExtension) == 0 {
		return nil, false, nil
	}

	resp := new(types.Attestations)
	if err := proto.Unmarshal(voteExtension, resp); err != nil {
		return nil, false, errors.Wrap(err, "decode vote extension")
	}

	return resp, true, nil
}
