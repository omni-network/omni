package hypercore

import (
	"context"
	"crypto/ecdsa"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/fireblocks"

	"github.com/ethereum/go-ethereum/common"
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

type fbSigner struct {
	client fireblocks.Client
	addr   common.Address
}

// NewFireblocksSigner returns a new Signer that uses Fireblocks to sign messages.
func NewFireblocksSigner(client fireblocks.Client, addr common.Address) Signer {
	return fbSigner{
		client: client,
		addr:   addr,
	}
}

func (s fbSigner) Sign(ctx context.Context, digest []byte) ([65]byte, error) {
	hash, err := cast.EthHash(digest)
	if err != nil {
		return [65]byte{}, errors.Wrap(err, "cast eth hash")
	}

	return s.client.Sign(ctx, hash, s.addr)
}
