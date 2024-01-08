package consensus

import (
	"bytes"
	"encoding/json"
	"sort"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	abci "github.com/cometbft/cometbft/api/cometbft/abci/v1"
	crypto "github.com/cometbft/cometbft/api/cometbft/crypto/v1"
)

// cpayload is the value that we are coming to consensus on.
// It is the single tx contained in the cometBFT consensus block.
type cpayload struct {
	// TODO(corver): Add execution cpayload.
	Aggregates []xchain.AggAttestation `json:"attestations"`
}

// payloadFromTXs returns the consensus cpayload contained in the list of raw txs.
func payloadFromTXs(txs [][]byte) (cpayload, error) {
	if len(txs) != 1 {
		return cpayload{}, errors.New("invalid number of consensus transactions, only 1 ever expected")
	}

	var resp cpayload
	if err := decode(txs[0], &resp); err != nil {
		return cpayload{}, err
	}

	return resp, nil
}

// headersByPubkey returns the attestations for the provided key.
func headersByPubkey(aggregates []xchain.AggAttestation, key crypto.PublicKey) []xchain.BlockHeader {
	var filtered []xchain.BlockHeader
	for _, agg := range aggregates {
		for _, sig := range agg.Signatures {
			if sig.ValidatorPubKey == [33]byte(key.GetSecp256K1()) {
				filtered = append(filtered, agg.BlockHeader)
				break
			}
		}
	}

	return filtered
}

// aggregatesFromProposal returns the aggregate attestations contained in the proposal's last local
// commit's vote extensions.
func aggregatesFromProposal(req *abci.PrepareProposalRequest) ([]xchain.AggAttestation, error) {
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
	var att []xchain.Attestation
	if err := decode(voteExtension, &att); err != nil {
		return nil, err
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
func isApproved(agg xchain.AggAttestation, validators []abci.ValidatorUpdate) bool {
	quorum := 2*len(validators)/3 + 1 //nolint:gomnd // Formula for 2/3+1 quorum.
	return len(agg.Signatures) >= quorum
}
