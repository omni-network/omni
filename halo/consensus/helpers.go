package consensus

import (
	"bytes"
	"encoding/json"
	"sort"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/ethereum/go-ethereum/beacon/engine"
)

// cPayload is the value that we are coming to consensus on.
// It is the single tx contained in the cometBFT consensus block.
type cPayload struct {
	EPayload   engine.ExecutableData   `json:"executable_payload"`
	Aggregates []xchain.AggAttestation `json:"attestations"`
}

// payloadFromTXs returns the consensus cPayload contained in the list of raw txs.
func payloadFromTXs(txs [][]byte) (cPayload, error) {
	if len(txs) != 1 {
		return cPayload{}, errors.New("invalid number of consensus transactions, only 1 ever expected")
	}

	var resp cPayload
	if err := decode(txs[0], &resp); err != nil {
		return cPayload{}, errors.Wrap(err, "decode cpayload")
	}

	return resp, nil
}

// headersByPubKey returns the attestations for the provided key.
func headersByPubKey(aggregates []xchain.AggAttestation, pubkey [33]byte) []xchain.BlockHeader {
	var filtered []xchain.BlockHeader
	for _, agg := range aggregates {
		for _, sig := range agg.Signatures {
			if sig.ValidatorPubKey == pubkey {
				filtered = append(filtered, agg.BlockHeader)
				break
			}
		}
	}

	return filtered
}

// aggregatesFromProposal returns the aggregate attestations contained in the proposal's last local
// commit's vote extensions.
func aggregatesFromProposal(req *abci.RequestPrepareProposal) ([]xchain.AggAttestation, error) {
	var attestations []xchain.Attestation
	for _, vote := range req.LocalLastCommit.Votes {
		voteAtts, err := attestationsFromVoteExt(vote.VoteExtension)
		if err != nil {
			return nil, err
		}
		attestations = append(attestations, voteAtts...)
	}

	return aggregate(attestations), nil
}

// aggregate aggregates the provided attestations by block header.
func aggregate(attestations []xchain.Attestation) []xchain.AggAttestation {
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
				Signatures:     []xchain.SigTuple{att.Signature},
			}
		}

		agg.Signatures = append(agg.Signatures, att.Signature)
		aggsByHeader[att.BlockHeader] = agg
	}

	aggs := make([]xchain.AggAttestation, 0, len(aggsByHeader))
	for _, agg := range aggsByHeader {
		aggs = append(aggs, agg)
	}

	// Sort deterministically.
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

// isApproved returns true if the provided aggregate attestation is approved by the provided validator set.
func isApproved(agg xchain.AggAttestation, validators []validator) bool {
	quorum := 2*len(validators)/3 + 1 //nolint:gomnd // Formula for 2/3+1 quorum.
	return len(agg.Signatures) >= quorum
}
