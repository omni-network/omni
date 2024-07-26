package voter

import (
	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/crypto"
)

// CreateVote creates a vote for the given block.
func CreateVote(privKey crypto.PrivKey, attHeader xchain.AttestHeader, block xchain.Block) (*types.Vote, error) {
	var msgRoot [32]byte
	if len(block.Msgs) > 0 {
		tree, err := xchain.NewMsgTree(block.Msgs)
		if err != nil {
			return nil, err
		}

		msgRoot = tree.MsgRoot()
	} // else use zero value msgRoot

	attRoot, err := xchain.AttestationRoot(attHeader, block.BlockHeader, msgRoot)
	if err != nil {
		return nil, err
	}

	sig, err := k1util.Sign(privKey, attRoot)
	if err != nil {
		return nil, errors.Wrap(err, "sign attestation")
	}

	address, err := k1util.PubKeyToAddress(privKey.PubKey())
	if err != nil {
		return nil, err
	}

	return &types.Vote{
		AttestHeader: &types.AttestHeader{
			ConsensusChainId: attHeader.ConsensusChainID,
			SourceChainId:    attHeader.ChainVersion.ID,
			ConfLevel:        uint32(attHeader.ChainVersion.ConfLevel),
			AttestOffset:     attHeader.AttestOffset,
		},
		BlockHeader: &types.BlockHeader{
			ChainId:     block.ChainID,
			BlockHeight: block.BlockHeight,
			BlockHash:   block.BlockHash.Bytes(),
		},
		MsgRoot: msgRoot[:],
		Signature: &types.SigTuple{
			ValidatorAddress: address[:],
			Signature:        sig[:],
		},
	}, nil
}
