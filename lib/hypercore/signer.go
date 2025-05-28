package hypercore

import (
	"context"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

type pkSigner struct {
	pk *ecdsa.PrivateKey
}

// NewPrivateKeySigner returns a new Signer that uses the provided ECDSA private key.
func NewPrivateKeySigner(privateKey *ecdsa.PrivateKey) Signer {
	return pkSigner{pk: privateKey}
}

func (s pkSigner) Sign(_ context.Context, digest []byte) ([65]byte, error) {
	sig, err := crypto.Sign(digest, s.pk)
	if err != nil {
		return [65]byte{}, err
	}

	// Convert to fixed-size array
	var out [65]byte
	copy(out[:], sig)

	return out, nil
}
