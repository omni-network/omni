package attester

import (
	"github.com/omni-network/omni/halo/attest/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/crypto"
)

// CreateAttestation creates an attestation for the given block.
func CreateAttestation(privKey crypto.PrivKey, block xchain.Block) (*types.Attestation, error) {
	pubkey := privKey.PubKey().Bytes()
	if len(pubkey) != 33 {
		return nil, errors.New("invalid pubkey length", "length", len(pubkey))
	}

	tree, err := xchain.NewBlockTree(block)
	if err != nil {
		return nil, err
	}
	root := tree.Root()

	sig, err := k1util.Sign(privKey, root)
	if err != nil {
		return nil, errors.Wrap(err, "sign attestation")
	}

	address, err := k1util.PubKeyToAddress(privKey.PubKey())
	if err != nil {
		return nil, err
	}

	return &types.Attestation{
		BlockHeader: &types.BlockHeader{
			ChainId: block.SourceChainID,
			Height:  block.BlockHeight,
			Hash:    block.BlockHash[:],
		},
		BlockRoot: root[:],
		Signature: &types.SigTuple{
			ValidatorAddress: address[:],
			Signature:        sig[:],
		},
	}, nil
}
