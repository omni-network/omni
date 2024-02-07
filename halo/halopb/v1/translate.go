package v1

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"
)

// AggregatesToProto converts a slice of xchain.AggAttestations to a slice of protobuf AggAttestations.
func AggregatesToProto(aggs []xchain.AggAttestation) []*AggAttestation {
	resp := make([]*AggAttestation, 0, len(aggs))
	for _, agg := range aggs {
		resp = append(resp, AggregateToProto(agg))
	}

	return resp
}

// AggregatesFromProto converts a slice of protobuf AggAttestations to a slice of xchain.AggAttestations.
func AggregatesFromProto(aggs []*AggAttestation) ([]xchain.AggAttestation, error) {
	resp := make([]xchain.AggAttestation, 0, len(aggs))
	for _, aggpb := range aggs {
		agg, err := AggregateFromProto(aggpb)
		if err != nil {
			return nil, err
		}
		resp = append(resp, agg)
	}

	return resp, nil
}

// AggregateToProto converts a xchain.AggAttestation to a protobuf AggAttestation.
func AggregateToProto(agg xchain.AggAttestation) *AggAttestation {
	sigs := make([]*SigTuple, 0, len(agg.Signatures))
	for _, sig := range agg.Signatures {
		sigs = append(sigs, SigToProto(sig))
	}

	return &AggAttestation{
		BlockHeader:    BlockHeaderToProto(agg.BlockHeader),
		ValidatorSetId: agg.ValidatorSetID,
		BlockRoot:      agg.BlockRoot[:],
		Signatures:     sigs,
	}
}

// AggregateFromProto converts a protobuf AggAttestation to a xchain.AggAttestation.
func AggregateFromProto(agg *AggAttestation) (xchain.AggAttestation, error) {
	if agg == nil {
		return xchain.AggAttestation{}, errors.New("nil aggregate attestation")
	}

	header, err := BlockHeaderFromProto(agg.GetBlockHeader())
	if err != nil {
		return xchain.AggAttestation{}, err
	}

	sigs := make([]xchain.SigTuple, 0, len(agg.GetSignatures()))
	for _, sigpb := range agg.GetSignatures() {
		sig, err := SigFromProto(sigpb)
		if err != nil {
			return xchain.AggAttestation{}, err
		}
		sigs = append(sigs, sig)
	}

	var zero xchain.AggAttestation
	if len(agg.GetBlockRoot()) != len(zero.BlockRoot) {
		return xchain.AggAttestation{}, errors.New("invalid block root length")
	}

	return xchain.AggAttestation{
		BlockHeader:    header,
		ValidatorSetID: agg.GetValidatorSetId(),
		BlockRoot:      [32]byte(agg.GetBlockRoot()),
		Signatures:     sigs,
	}, nil
}

// SigFromProto converts a protobuf SigTuple to a xchain.SigTuple.
func SigFromProto(sig *SigTuple) (xchain.SigTuple, error) {
	var zero xchain.SigTuple

	if sig == nil {
		return xchain.SigTuple{}, errors.New("nil sig tuple")
	} else if len(sig.GetValidatorAddress()) != len(zero.ValidatorAddress) {
		return xchain.SigTuple{}, errors.New("invalid validator address length")
	} else if len(sig.GetSignature()) != len(zero.Signature) {
		return xchain.SigTuple{}, errors.New("invalid signature length")
	}

	return xchain.SigTuple{
		ValidatorAddress: [20]byte(sig.GetValidatorAddress()),
		Signature:        [65]byte(sig.GetSignature()),
	}, nil
}

// SigToProto converts a xchain.SigTuple to a protobuf SigTuple.
func SigToProto(sig xchain.SigTuple) *SigTuple {
	return &SigTuple{
		ValidatorAddress: sig.ValidatorAddress[:],
		Signature:        sig.Signature[:],
	}
}

// BlockHeaderToProto converts a xchain.BlockHeader to a protobuf BlockHeader.
func BlockHeaderToProto(header xchain.BlockHeader) *BlockHeader {
	return &BlockHeader{
		ChainId: header.SourceChainID,
		Height:  header.BlockHeight,
		Hash:    header.BlockHash[:],
	}
}

// BlockHeaderFromProto converts a protobuf BlockHeader to a xchain.BlockHeader.
func BlockHeaderFromProto(header *BlockHeader) (xchain.BlockHeader, error) {
	if header == nil {
		return xchain.BlockHeader{}, errors.New("nil block header")
	}

	var zero xchain.BlockHeader
	if len(header.GetHash()) != len(zero.BlockHash) {
		return xchain.BlockHeader{}, errors.New("invalid block hash length")
	}

	return xchain.BlockHeader{
		SourceChainID: header.GetChainId(),
		BlockHeight:   header.GetHeight(),
		BlockHash:     [32]byte(header.GetHash()),
	}, nil
}
