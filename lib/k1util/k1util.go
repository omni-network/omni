// Package k1util provides functions to sign and verify Ethereum RSV style signatures.
// It was copied from https://github.com/ObolNetwork/charon/blob/main/app/k1util/k1util.go.
package k1util

import (
	"github.com/omni-network/omni/lib/errors"

	"github.com/cometbft/cometbft/crypto"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
)

// privKeyLen is the length of a secp256k1 private key.
const privKeyLen = 32

const scalarLen = 32

// Sign returns a signature from input data.
//
// The produced signature is 65 bytes in the [R || S || V] format where V is 27 or 28.
func Sign(key crypto.PrivKey, input [32]byte) ([65]byte, error) {
	bz := key.Bytes()
	if len(bz) != privKeyLen {
		return [65]byte{}, errors.New("invalid private key length")
	}

	sig := ecdsa.SignCompact(secp256k1.PrivKeyFromBytes(bz), input[:], false)

	// Convert signature from "compact" into "Ethereum R S V" format.
	return [65]byte(append(sig[1:], sig[0])), nil
}

// Verify returns whether the 65 byte signature is valid for the provided hash
// and secp256k1 public key.
//
// Note the signature MUST be 65 bytes in the Ethereum [R || S || V] format.
func Verify(pubkey crypto.PubKey, hash [32]byte, sig [65]byte) (bool, error) {
	r, err := to32Scalar(sig[:scalarLen])
	if err != nil {
		return false, errors.Wrap(err, "invalid signature R")
	}

	s, err := to32Scalar(sig[scalarLen : 2*scalarLen])
	if err != nil {
		return false, errors.Wrap(err, "invalid signature S")
	}

	pk, err := secp256k1.ParsePubKey(pubkey.Bytes())
	if err != nil {
		return false, errors.Wrap(err, "invalid public key")
	}

	return ecdsa.NewSignature(r, s).Verify(hash[:], pk), nil
}

// to32Scalar returns the 256-bit big-endian unsigned
// integer as a scalar.
func to32Scalar(b []byte) (*secp256k1.ModNScalar, error) {
	if len(b) != scalarLen {
		return nil, errors.New("invalid scalar length")
	}

	// Strip leading zeroes from S.
	for len(b) > 0 && b[0] == 0x00 {
		b = b[1:]
	}

	var s secp256k1.ModNScalar
	if overflow := s.SetByteSlice(b); overflow {
		return nil, errors.New("scalar overflow")
	} else if s.IsZero() {
		return nil, errors.New("zero overflow")
	}

	return &s, nil
}
