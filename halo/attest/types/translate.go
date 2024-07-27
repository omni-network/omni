package types

import (
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

func (h *AttestHeader) XChainVersion() xchain.ChainVersion {
	return xchain.ChainVersion{
		ID:        h.SourceChainId,
		ConfLevel: xchain.ConfLevel(h.ConfLevel),
	}
}

func (h *WindowCompareRequest) XChainVersion() xchain.ChainVersion {
	return xchain.ChainVersion{
		ID:        h.ChainId,
		ConfLevel: xchain.ConfLevel(h.ConfLevel),
	}
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

// AttestationFromProto converts a protobuf Attestation to a xchain.Attestation.
func AttestationFromProto(att *Attestation) (xchain.Attestation, error) {
	if err := att.Verify(); err != nil {
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
		AttestHeader:   AttestHeaderFromProto(att.GetAttestHeader()),
		BlockHeader:    BlockHeaderFromProto(att.GetBlockHeader()),
		ValidatorSetID: att.GetValidatorSetId(),
		MsgRoot:        common.BytesToHash(att.GetMsgRoot()),
		Signatures:     sigs,
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

// BlockHeaderFromProto converts a protobuf BlockHeader to a xchain.BlockHeader.
func BlockHeaderFromProto(header *BlockHeader) xchain.BlockHeader {
	return xchain.BlockHeader{
		ChainID:     header.GetChainId(),
		BlockHeight: header.GetBlockHeight(),
		BlockHash:   common.Hash(header.GetBlockHash()),
	}
}

// AttestHeaderFromProto converts a protobuf AttestHeader to a xchain.AttestHeader.
func AttestHeaderFromProto(header *AttestHeader) xchain.AttestHeader {
	return xchain.AttestHeader{
		ConsensusChainID: header.GetConsensusChainId(),
		ChainVersion: xchain.ChainVersion{
			ID:        header.GetSourceChainId(),
			ConfLevel: xchain.ConfLevel(header.GetConfLevel()),
		},
		AttestOffset: header.GetAttestOffset(),
	}
}
