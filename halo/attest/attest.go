package attest

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/crypto"
	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
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

	return xchain.Attestation{
		BlockHeader: block.BlockHeader,
		BlockRoot:   root,
		Signature: xchain.SigTuple{
			ValidatorPubKey: [33]byte(pubkey),
			Signature:       sig,
		},
	}, nil
}

// VerifyAttestation verifies the attestation signature.
func VerifyAttestation(att xchain.Attestation) error {
	pk := k1.PubKey(att.Signature.ValidatorPubKey[:])
	ok, err := k1util.Verify(pk, att.BlockRoot, att.Signature.Signature)
	if err != nil {
		return err
	} else if !ok {
		return ErrInvalidAttestation
	}

	return nil
}
