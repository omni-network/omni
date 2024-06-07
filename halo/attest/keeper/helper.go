package keeper

import (
	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/xchain"
)

func (a *Attestation) XChainVersion() xchain.ChainVersion {
	return xchain.ChainVersion{
		ID:        a.GetChainId(),
		ConfLevel: xchain.ConfLevel(a.GetConfLevel()),
	}
}

func AttestationFromDB(att *Attestation, sigs []*Signature) *types.Attestation {
	return &types.Attestation{
		BlockHeader: &types.BlockHeader{
			ChainId:       att.GetChainId(),
			ConfLevel:     att.GetConfLevel(),
			Offset:        att.GetOffset(),
			Height:        att.GetHeight(),
			Hash:          att.GetBlockHash(),
			StreamOffsets: offsetsFromDB(att.GetStreamOffsets()),
		},
		ValidatorSetId:  att.GetValidatorSetId(),
		AttestationRoot: att.GetAttestationRoot(),
		Signatures:      sigsFromDB(sigs),
	}
}

func sigsFromDB(sigs []*Signature) []*types.SigTuple {
	resp := make([]*types.SigTuple, 0, len(sigs))
	for _, sig := range sigs {
		resp = append(resp, &types.SigTuple{
			ValidatorAddress: sig.GetValidatorAddress(),
			Signature:        sig.GetSignature(),
		})
	}

	return resp
}

func offsetsFromDB(offsets []*BlockStreamOffset) []*types.BlockStreamOffset {
	resp := make([]*types.BlockStreamOffset, 0, len(offsets))
	for _, offset := range offsets {
		resp = append(resp, &types.BlockStreamOffset{
			DestChainId: offset.GetDestChainId(),
			ShardId:     offset.GetShardId(),
			MsgOffset:   offset.GetMsgOffset(),
		})
	}

	return resp
}

func offsetsToDB(offsets []*types.BlockStreamOffset) []*BlockStreamOffset {
	resp := make([]*BlockStreamOffset, 0, len(offsets))
	for _, offset := range offsets {
		resp = append(resp, &BlockStreamOffset{
			DestChainId: offset.GetDestChainId(),
			ShardId:     offset.GetShardId(),
			MsgOffset:   offset.GetMsgOffset(),
		})
	}

	return resp
}
