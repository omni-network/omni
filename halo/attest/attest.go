package attest

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/crypto"
)

var ErrInvalidAttestation = errors.New("invalid attestation signature")

// CreateAttestation creates an attestation for the given block.
func CreateAttestation(privKey crypto.PrivKey, block xchain.Block) (xchain.Attestation, error) {
	pubkey := privKey.PubKey().Bytes()
	if len(pubkey) != 33 {
		return xchain.Attestation{}, errors.New("invalid pubkey length", "length", len(pubkey))
	}

	tree, err := xchain.NewBlockTree(block)
	if err != nil {
		return xchain.Attestation{}, err
	}
	root := tree.Root()

	sig, err := k1util.Sign(privKey, root)
	if err != nil {
		return xchain.Attestation{}, errors.Wrap(err, "sign attestation")
	}

	address, err := k1util.PubKeyToAddress(privKey.PubKey())
	if err != nil {
		return xchain.Attestation{}, err
	}

	return xchain.Attestation{
		BlockHeader: block.BlockHeader,
		BlockRoot:   root,
		Signature: xchain.SigTuple{
			ValidatorAddress: address,
			Signature:        sig,
		},
	}, nil
}

// VerifyAttestation verifies the attestation signature.
func VerifyAttestation(att xchain.Attestation) error {
	ok, err := k1util.Verify(att.Signature.ValidatorAddress, att.BlockRoot, att.Signature.Signature)
	if err != nil {
		return err
	} else if !ok {
		return ErrInvalidAttestation
	}

	return nil
}
