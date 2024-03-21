package types

import (
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

// AttestationsToProto converts a slice of xchain.Attestations to a slice of protobuf Attestations.
func AttestationsToProto(atts []xchain.Attestation) []*Attestation {
	resp := make([]*Attestation, 0, len(atts))
	for _, att := range atts {
		resp = append(resp, AttestationToProto(att))
	}

	return resp
}

// AttestationsFromProto converts a slice of protobuf Attestations to a slice of xchain.Attestations.
func AttestationsFromProto(atts []*Attestation) ([]xchain.Attestation, error) {
	resp := make([]xchain.Attestation, 0, len(atts))
	for _, attpb := range atts {
		att, err := AttestationFromProto(attpb)
		if err != nil {
			return nil, err
		}
		resp = append(resp, att)
	}

	return resp, nil
}

// AttestationToProto converts a xchain.Attestation to a protobuf Attestation.
func AttestationToProto(att xchain.Attestation) *Attestation {
	sigs := make([]*SigTuple, 0, len(att.Signatures))
	for _, sig := range att.Signatures {
		sigs = append(sigs, SigToProto(sig))
	}

	return &Attestation{
		BlockHeader:     BlockHeaderToProto(att.BlockHeader),
		ValidatorSetId:  att.ValidatorSetID,
		AttestationRoot: att.AttestationRoot[:],
		Signatures:      sigs,
	}
}

// AttestationFromProto converts a protobuf Attestation to a xchain.Attestation.
func AttestationFromProto(att *Attestation) (xchain.Attestation, error) {
	if err := att.Verify(); err != nil {
		return xchain.Attestation{}, err
	}

	header, err := BlockHeaderFromProto(att.GetBlockHeader())
	if err != nil {
		return xchain.Attestation{}, err
	}

	sigs := make([]xchain.SigTuple, 0, len(att.GetSignatures()))
	for _, sigpb := range att.GetSignatures() {
		sig, err := SigFromProto(sigpb)
		if err != nil {
			return xchain.Attestation{}, err
		}
		sigs = append(sigs, sig)
	}

	return xchain.Attestation{
		BlockHeader:     header,
		ValidatorSetID:  att.ValidatorSetId,
		AttestationRoot: common.BytesToHash(att.GetAttestationRoot()),
		Signatures:      sigs,
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
