package attest

import (
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/cometbft/cometbft/crypto"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
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

	sig, err := secp256k1.Sign(root[:], privKey.Bytes())
	if err != nil {
		return xchain.Attestation{}, errors.Wrap(err, "sign attestation")
	} else if len(sig) != 65 {
		return xchain.Attestation{}, errors.New("invalid signature length", "length", len(sig))
	}

	return xchain.Attestation{
		BlockHeader: block.BlockHeader,
		BlockRoot:   root,
		Signature: xchain.SigTuple{
			ValidatorPubKey: [33]byte(pubkey),
			Signature:       [65]byte(sig),
		},
	}, nil
}

// VerifyAttestation verifies the attestation signature.
func VerifyAttestation(att xchain.Attestation) error {
	// Trim recovery ID
	trimmedSig := att.Signature.Signature[:64]

	ok := secp256k1.VerifySignature(att.Signature.ValidatorPubKey[:], att.BlockRoot[:], trimmedSig)
	if !ok {
		return ErrInvalidAttestation
	}

	return nil
}
