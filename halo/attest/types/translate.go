package types

import (
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/xchain"
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

	header, err := BlockHeaderFromProto(att.GetBlockHeader())
	if err != nil {
		return xchain.Attestation{}, err
	}

	msgRoot, err := cast.Array32(att.GetMsgRoot())
	if err != nil {
		return xchain.Attestation{}, err
	}

	return xchain.Attestation{
		AttestHeader:   AttestHeaderFromProto(att.GetAttestHeader()),
		BlockHeader:    header,
		ValidatorSetID: att.GetValidatorSetId(),
		MsgRoot:        msgRoot,
		Signatures:     sigs,
	}, nil
}

// SigFromProto converts a protobuf SigTuple to a xchain.SigTuple.
func SigFromProto(sig *SigTuple) (xchain.SigTuple, error) {
	if err := sig.Verify(); err != nil {
		return xchain.SigTuple{}, err
	}

	return sig.ToXChain()
}

// BlockHeaderFromProto converts a protobuf BlockHeader to a xchain.BlockHeader.
func BlockHeaderFromProto(header *BlockHeader) (xchain.BlockHeader, error) {
	addr, err := cast.Array32(header.GetBlockHash())
	if err != nil {
		return xchain.BlockHeader{}, err
	}

	return xchain.BlockHeader{
		ChainID:     header.GetChainId(),
		BlockHeight: header.GetBlockHeight(),
		BlockHash:   addr,
	}, nil
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
