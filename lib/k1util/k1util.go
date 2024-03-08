// Package k1util provides functions to sign and verify Ethereum RSV style signatures.
package k1util

import (
	"github.com/omni-network/omni/lib/errors"

	"github.com/cometbft/cometbft/crypto"
	cryptopb "github.com/cometbft/cometbft/proto/tendermint/crypto"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	cosmosk1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmoscrypto "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
)

// privKeyLen is the length of a secp256k1 private key.
const privKeyLen = 32

// pubkeyLen is the length of a secp256k1 compressed public key.
const pubkeyLen = 33

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
// and Ethereum address.
//
// Note the signature MUST be 65 bytes in the Ethereum [R || S || V] format.
func Verify(address common.Address, hash [32]byte, sig [65]byte) (bool, error) {
	// Adjust V from Ethereum 27/28 to secp256k1 0/1
	const vIdx = 64
	if v := sig[vIdx]; v != 27 && v != 28 {
		return false, errors.New("invalid recovery id (V) format, must be 27 or 28")
	}
	sig[vIdx] -= 27

	pubkey, err := ethcrypto.SigToPub(hash[:], sig[:])
	if err != nil {
		return false, errors.Wrap(err, "recover public key")
	}

	actual := ethcrypto.PubkeyToAddress(*pubkey)

	return actual == address, nil
}

// PubKeyToAddress returns the Ethereum address for the given k1 public key.
func PubKeyToAddress(pubkey crypto.PubKey) (common.Address, error) {
	pubkeyBytes := pubkey.Bytes()
	if len(pubkeyBytes) != pubkeyLen {
		return common.Address{}, errors.New("invalid pubkey length", "length", len(pubkeyBytes))
	}

	ethPubKey, err := ethcrypto.DecompressPubkey(pubkeyBytes)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "decompress pubkey")
	}

	return ethcrypto.PubkeyToAddress(*ethPubKey), nil
}

func PubKeyToCosmos(pubkey crypto.PubKey) (cosmoscrypto.PubKey, error) {
	pubkeyBytes := pubkey.Bytes()
	if len(pubkeyBytes) != pubkeyLen {
		return nil, errors.New("invalid pubkey length", "length", len(pubkeyBytes))
	}

	return &cosmosk1.PubKey{
		Key: pubkey.Bytes(),
	}, nil
}

// PubKeyPBToAddress returns the Ethereum address for the given k1 public key.
func PubKeyPBToAddress(pubkey cryptopb.PublicKey) (common.Address, error) {
	pubkeyBytes := pubkey.GetSecp256K1()
	if len(pubkeyBytes) != pubkeyLen {
		return common.Address{}, errors.New("invalid pubkey length", "length", len(pubkeyBytes))
	}

	ethPubKey, err := ethcrypto.DecompressPubkey(pubkeyBytes)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "decompress pubkey")
	}

	return ethcrypto.PubkeyToAddress(*ethPubKey), nil
}
