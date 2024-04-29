package key

import (
	"crypto/ecdsa"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/cometbft/cometbft/crypto"
	ed "github.com/cometbft/cometbft/crypto/ed25519"
	k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

// Type represents the type of cryptographic key.
type Type string

const (
	Validator    Type = "validator"
	P2PConsensus Type = "p2p_consensus"
	P2PExecution Type = "p2p_execution"
	EOA          Type = "eoa"
)

func (t Type) Verify() error {
	if t == "" {
		return errors.New("empty key type")
	}
	switch t {
	case Validator, P2PConsensus, P2PExecution, EOA:
		return nil
	default:
		return errors.New("invalid key type")
	}
}

func (t Type) String() string {
	return string(t)
}

// Key wraps a cometBFT private key adding custom a address function.
type Key struct {
	crypto.PrivKey
}

// Addr returns the address of the key.
// K1 keys have ethereum 0x addresses,
// ED keys have the default SHA256-20 of the raw pubkey bytes.
func (k Key) Addr() (string, error) {
	switch t := k.PrivKey.(type) {
	case k1.PrivKey:
		addr, err := k1util.PubKeyToAddress(t.PubKey())
		if err != nil {
			return "", err
		}

		return addr.Hex(), nil
	case ed.PrivKey:
		return t.PubKey().Address().String(), nil
	default:
		return "", errors.New("unknown key type")
	}
}

// ECDSA returns the ECDSA representation of the k1 key.
// This returns an error for ed25519 keys.
func (k Key) ECDSA() (*ecdsa.PrivateKey, error) {
	switch t := k.PrivKey.(type) {
	case k1.PrivKey:
		resp, err := ethcrypto.ToECDSA(t.Bytes())
		if err != nil {
			return nil, errors.Wrap(err, "converting to ECDSA")
		}

		return resp, nil
	default:
		return nil, errors.New("ed25519 keys do not have ECDSA representation")
	}
}

// Generate generates a cryptographic key of the specified type.
// It panics since it assumes that the type is valid.
func Generate(typ Type) Key {
	switch typ {
	case Validator, P2PExecution, EOA:
		return Key{
			PrivKey: k1.GenPrivKey(),
		}

	case P2PConsensus:
		return Key{
			PrivKey: ed.GenPrivKey(),
		}
	default:
		panic("invalid key type:" + typ)
	}
}

// GenerateInsecureDeterministic generates an insecure deterministic key of the specified type
// from the provided seed.
// NOTE THIS MUST ONLY BE USED FOR TESTING: It panics if network is not ephemeral.
func GenerateInsecureDeterministic(network netconf.ID, typ Type, seed string) Key {
	if !network.IsEphemeral() {
		panic("only ephemeral keys are supported")
	}

	secret := []byte(string(typ) + "|" + seed) // Deterministic secret from typ+seed.

	switch typ {
	case Validator, P2PExecution, EOA:
		return Key{
			PrivKey: k1.GenPrivKeySecp256k1(secret),
		}

	case P2PConsensus:
		return Key{
			PrivKey: ed.GenPrivKeyFromSecret(secret),
		}
	default:
		panic("invalid key type:" + typ)
	}
}

// FromBytes parses the given bytes into th eprovided key type.
func FromBytes(typ Type, b []byte) (Key, error) {
	switch typ {
	case Validator, P2PExecution, EOA:
		if len(b) != k1.PrivKeySize {
			return Key{}, errors.New("invalid key size")
		}

		return Key{PrivKey: k1.PrivKey(b)}, nil
	case P2PConsensus:

		if len(b) != ed.PrivateKeySize {
			return Key{}, errors.New("invalid key size")
		}

		return Key{PrivKey: ed.PrivKey(b)}, nil
	default:
		return Key{}, errors.New("invalid key type")
	}
}
