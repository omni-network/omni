package types

import (
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

// AggregatesToProto converts a slice of xchain.AggAttestations to a slice of protobuf AggAttestations.
func AggregatesToProto(aggs []xchain.AggAttestation) []*Attestation {
	resp := make([]*Attestation, 0, len(aggs))
	for _, agg := range aggs {
		resp = append(resp, AggregateToProto(agg))
	}

	return resp
}

// AggregatesFromProto converts a slice of protobuf AggAttestations to a slice of xchain.AggAttestations.
func AggregatesFromProto(aggs []*Attestation) ([]xchain.AggAttestation, error) {
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
func AggregateToProto(agg xchain.AggAttestation) *Attestation {
	sigs := make([]*SigTuple, 0, len(agg.Signatures))
	for _, sig := range agg.Signatures {
		sigs = append(sigs, SigToProto(sig))
	}

	return &Attestation{
		BlockHeader:    BlockHeaderToProto(agg.BlockHeader),
		ValidatorsHash: agg.ValidatorSetHash[:],
		BlockRoot:      agg.BlockRoot[:],
		Signatures:     sigs,
	}
}

// AggregateFromProto converts a protobuf AggAttestation to a xchain.AggAttestation.
func AggregateFromProto(agg *Attestation) (xchain.AggAttestation, error) {
	if err := agg.Verify(); err != nil {
		return xchain.AggAttestation{}, err
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

	return xchain.AggAttestation{
		BlockHeader:      header,
		ValidatorSetHash: common.BytesToHash(agg.GetValidatorsHash()),
		BlockRoot:        common.BytesToHash(agg.GetBlockRoot()),
		Signatures:       sigs,
	}, nil
}

// SigFromProto converts a protobuf SigTuple to a xchain.SigTuple.
func SigFromProto(sig *SigTuple) (xchain.SigTuple, error) {
	if err := sig.Verify(); err != nil {
		return xchain.SigTuple{}, err
	}

	return xchain.SigTuple{
		ValidatorAddress: common.Address(sig.GetValidatorAddress()),
		Signature:        xchain.Signature65(sig.GetSignature()),
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
	if err := header.Verify(); err != nil {
		return xchain.BlockHeader{}, err
	}

	return xchain.BlockHeader{
		SourceChainID: header.GetChainId(),
		BlockHeight:   header.GetHeight(),
		BlockHash:     common.Hash(header.GetHash()),
	}, nil
}
