package keeper

import (
	"github.com/omni-network/omni/halo2/attest/types"
)

func aggAttestationToORM(att *types.MsgAggAttestation) *AggAttestation {
	return &AggAttestation{
		ChainId:        att.BlockHeader.ChainId,
		Height:         att.BlockHeader.Height,
		Hash:           att.BlockHeader.Hash,
		BlockRoot:      att.BlockRoot,
		Status:         int32(AggStatus_Pending),
		ValidatorSetId: att.ValidatorSetId,
	}
}
