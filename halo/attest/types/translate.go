package types

import (
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

func (h *BlockHeader) XChainVersion() xchain.ChainVersion {
	return xchain.ChainVersion{
		ID:        h.ChainId,
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
		sig, err := sigFromProto(sigpb)
		if err != nil {
			return xchain.Attestation{}, err
		}
		sigs = append(sigs, sig)
	}

	return xchain.Attestation{
		BlockHeader:     blockHeaderFromProto(att.GetBlockHeader()),
		ValidatorSetID:  att.ValidatorSetId,
		AttestationRoot: common.BytesToHash(att.GetAttestationRoot()),
		Signatures:      sigs,
	}, nil
}

// sigFromProto converts a protobuf SigTuple to a xchain.SigTuple.
func sigFromProto(sig *SigTuple) (xchain.SigTuple, error) {
	if err := sig.Verify(); err != nil {
		return xchain.SigTuple{}, err
	}

	return xchain.SigTuple{
		ValidatorAddress: common.Address(sig.GetValidatorAddress()),
		Signature:        xchain.Signature65(sig.GetSignature()),
	}, nil
}

// blockHeaderFromProto converts a protobuf BlockHeader to a xchain.BlockHeader.
func blockHeaderFromProto(header *BlockHeader) xchain.BlockHeader {
	var offsets []xchain.BlockStreamOffset
	for _, offset := range header.GetStreamOffsets() {
		off, err := streamOffsetFromProto(offset)
		if err != nil {
			return xchain.BlockHeader{}
		}
		offsets = append(offsets, off)
	}

	return xchain.BlockHeader{
		SourceChainID: header.GetChainId(),
		ConfLevel:     xchain.ConfLevel(byte(header.ConfLevel)),
		BlockOffset:   header.GetOffset(),
		BlockHeight:   header.GetHeight(),
		BlockHash:     common.Hash(header.GetHash()),
		StreamOffsets: offsets,
	}
}

func BlockHeaderToProto(header xchain.BlockHeader) *BlockHeader {
	offsets := make([]*BlockStreamOffset, 0, len(header.StreamOffsets))
	for _, offset := range header.StreamOffsets {
		offsets = append(offsets, streamOffsetToProto(offset))
	}

	return &BlockHeader{
		ChainId:       header.SourceChainID,
		ConfLevel:     uint32(header.ConfLevel),
		Offset:        header.BlockOffset,
		Height:        header.BlockHeight,
		Hash:          header.BlockHash[:],
		StreamOffsets: offsets,
	}
}

func streamOffsetFromProto(offset *BlockStreamOffset) (xchain.BlockStreamOffset, error) {
	if err := offset.Verify(); err != nil {
		return xchain.BlockStreamOffset{}, err
	}

	return xchain.BlockStreamOffset{
		DestChainID: offset.GetDestChainId(),
		ShardID:     xchain.ShardID(offset.GetShardId()),
		MsgOffset:   offset.GetMsgOffset(),
	}, nil
}

func streamOffsetToProto(offset xchain.BlockStreamOffset) *BlockStreamOffset {
	return &BlockStreamOffset{
		DestChainId: offset.DestChainID,
		ShardId:     uint64(offset.ShardID),
		MsgOffset:   offset.MsgOffset,
	}
}
