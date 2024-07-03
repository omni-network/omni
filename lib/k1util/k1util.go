// Package k1util provides functions to sign and verify Ethereum RSV style signatures.
package k1util

import (
	stdecdsa "crypto/ecdsa"

	"github.com/omni-network/omni/lib/errors"

	"github.com/cometbft/cometbft/crypto"
	k1 "github.com/cometbft/cometbft/crypto/secp256k1"
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

// pubkeyCompressedLen is the length of a secp256k1 compressed public key.
const pubkeyCompressedLen = 33

// pubkeyUncompressedLen is the length of a secp256k1 uncompressed public key.
const pubkeyUncompressedLen = 65

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
	if len(pubkeyBytes) != pubkeyCompressedLen {
		return common.Address{}, errors.New("invalid pubkey length", "length", len(pubkeyBytes))
	}

	ethPubKey, err := ethcrypto.DecompressPubkey(pubkeyBytes)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "decompress pubkey")
	}

	return ethcrypto.PubkeyToAddress(*ethPubKey), nil
}

func StdPrivKeyToComet(privkey *stdecdsa.PrivateKey) (crypto.PrivKey, error) {
	bz := ethcrypto.FromECDSA(privkey)
	if len(bz) != privKeyLen {
		return nil, errors.New("invalid private key length")
	}

	return k1.PrivKey(bz), nil
}

func StdPrivKeyFromComet(privkey crypto.PrivKey) (*stdecdsa.PrivateKey, error) {
	bz := privkey.Bytes()
	if len(bz) != privKeyLen {
		return nil, errors.New("invalid private key length")
	}

	resp, err := ethcrypto.ToECDSA(bz)
	if err != nil {
		return nil, errors.Wrap(err, "convert to ECDSA")
	}

	return resp, nil
}

func StdPubKeyToCosmos(pubkey *stdecdsa.PublicKey) (cosmoscrypto.PubKey, error) {
	return PubKeyBytesToCosmos(ethcrypto.CompressPubkey(pubkey))
}

func PubKeyToCosmos(pubkey crypto.PubKey) (cosmoscrypto.PubKey, error) {
	return PubKeyBytesToCosmos(pubkey.Bytes())
}

func PubKeyBytesToCosmos(pubkey []byte) (cosmoscrypto.PubKey, error) {
	if len(pubkey) != pubkeyCompressedLen {
		return nil, errors.New("invalid pubkey length", "length", len(pubkey))
	}

	return &cosmosk1.PubKey{
		Key: pubkey,
	}, nil
}

func PBPubKeyFromBytes(pubkey []byte) (cryptopb.PublicKey, error) {
	if len(pubkey) != pubkeyCompressedLen {
		return cryptopb.PublicKey{}, errors.New("invalid pubkey length", "length", len(pubkey))
	}

	return cryptopb.PublicKey{Sum: &cryptopb.PublicKey_Secp256K1{Secp256K1: pubkey}}, nil
}

// PubKeyPBToAddress returns the Ethereum address for the given k1 public key.
func PubKeyPBToAddress(pubkey cryptopb.PublicKey) (common.Address, error) {
	pubkeyBytes := pubkey.GetSecp256K1()
	if len(pubkeyBytes) != pubkeyCompressedLen {
		return common.Address{}, errors.New("invalid pubkey length", "length", len(pubkeyBytes))
	}

	ethPubKey, err := ethcrypto.DecompressPubkey(pubkeyBytes)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "decompress pubkey")
	}

	return ethcrypto.PubkeyToAddress(*ethPubKey), nil
}

// PubKeyToBytes64 returns the 64 byte uncompressed version of the public key, by removing the prefix (0x04 for uncompressed keys).
func PubKeyToBytes64(pubkey *stdecdsa.PublicKey) ([]byte, error) {
	bz := ethcrypto.FromECDSAPub(pubkey)
	if len(bz) != pubkeyUncompressedLen {
		return nil, errors.New("invalid pubkey")
	}

	return bz[1:], nil
}

// PubKeyFromBytes64 returns the public key from the 64 byte uncompressed version.
// It adds the prefix (0x04 for uncompressed keys) to the input bytes.
func PubKeyFromBytes64(pubkey []byte) (*stdecdsa.PublicKey, error) {
	if len(pubkey) != pubkeyUncompressedLen-1 {
		return nil, errors.New("invalid pubkey length", "length", len(pubkey))
	}

	const prefix = 0x04

	resp, err := ethcrypto.UnmarshalPubkey(append([]byte{prefix}, pubkey...))
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal pubkey")
	}

	return resp, nil
}
