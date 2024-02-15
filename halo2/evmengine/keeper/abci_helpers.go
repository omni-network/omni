package keeper

import (
	"bytes"
	"encoding/json"
	"sort"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	abci "github.com/cometbft/cometbft/abci/types"
)

// aggregatesFromProposal returns the aggregateAtts attestations contained in the proposal's last local
// commit's vote extensions.
func aggregatesFromProposal(info abci.ExtendedCommitInfo) ([]xchain.AggAttestation, error) {
	var attestations []xchain.Attestation
	for _, vote := range info.Votes {
		voteAtts, err := attestationsFromVoteExt(vote.VoteExtension)
		if err != nil {
			return nil, err
		}
		attestations = append(attestations, voteAtts...)
	}

	return aggregateAtts(attestations), nil
}

// aggregateAtts aggregates the provided attestations by block header.
func aggregateAtts(attestations []xchain.Attestation) []xchain.AggAttestation {
	aggsByHeader := make(map[xchain.BlockHeader]xchain.AggAttestation)
	for _, att := range attestations {
		agg, ok := aggsByHeader[att.BlockHeader]
		if !ok {
			agg = xchain.AggAttestation{
				BlockHeader: xchain.BlockHeader{
					SourceChainID: att.SourceChainID,
					BlockHeight:   att.BlockHeight,
					BlockHash:     att.BlockHash,
				},
				ValidatorSetID: 0, // TODO(corver): Figoure out how to map attestation to valsetid.
				BlockRoot:      att.BlockRoot,
			}
		}

		agg.Signatures = append(agg.Signatures, att.Signature)
		aggsByHeader[att.BlockHeader] = agg
	}

	return flattenAggsByHeader(aggsByHeader)
}

// flattenAggsByHeader returns the provided map of aggregates by header as a slice in a deterministic order.
func flattenAggsByHeader(aggsByHeader map[xchain.BlockHeader]xchain.AggAttestation) []xchain.AggAttestation {
	aggs := make([]xchain.AggAttestation, 0, len(aggsByHeader))
	for _, agg := range aggsByHeader {
		aggs = append(aggs, agg)
	}

	return sortAggregates(aggs)
}

// sortAggregates returns the provided aggregates in a deterministic order.
// Note the provided slice is also sorted in-place.
func sortAggregates(aggs []xchain.AggAttestation) []xchain.AggAttestation {
	sort.Slice(aggs, func(i, j int) bool {
		if aggs[i].BlockHeight != aggs[j].BlockHeight {
			return aggs[i].BlockHeight < aggs[j].BlockHeight
		}
		if aggs[i].SourceChainID != aggs[j].SourceChainID {
			return aggs[i].SourceChainID < aggs[j].SourceChainID
		}

		return bytes.Compare(aggs[i].BlockHash[:], aggs[j].BlockHash[:]) < 0
	})

	return aggs
}

// attestationsFromVoteExt returns the attestations contained in the vote extension.
func attestationsFromVoteExt(voteExtension []byte) ([]xchain.Attestation, error) {
	if len(voteExtension) == 0 {
		return nil, nil
	}

	var att []xchain.Attestation
	if err := decode(voteExtension, &att); err != nil {
		return nil, errors.Wrap(err, "decode vote extension")
	}

	return att, nil
}

// encode serializes the provided value.
// TODO(corver): We should use an optimized serialization, not json, probably SSZ or protobuf.
func encode(atts any) ([]byte, error) {
	buf, err := json.Marshal(atts)
	if err != nil {
		return nil, errors.Wrap(err, "marshal json")
	}

	return buf, nil
}

// decode de-serializes the provided data in the pointer.
// TODO(corver): We should use an optimized serialization, not json, probably SSZ or protobuf.
func decode(data []byte, ptr any) error {
	if err := json.Unmarshal(data, ptr); err != nil {
		return errors.Wrap(err, "unmarshal json")
	}

	return nil
}
