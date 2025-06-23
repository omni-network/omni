package uni

import (
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gagliardetto/solana-go"
)

// HexToAddress converts a hex string to an EVM address.
func HexToAddress(hex string) (Address, error) {
	if !common.IsHexAddress(hex) {
		return Address{}, errors.New("invalid hex address", "address", hex)
	}

	return EVMAddress(common.HexToAddress(hex)), nil
}

// Base58ToAddress converts a base58 string to a SVM address.
func Base58ToAddress(base58 string) (Address, error) {
	svm, err := solana.PublicKeyFromBase58(base58)
	if err != nil {
		return Address{}, errors.Wrap(err, "invalid base58 address")
	}

	return SVMAddress(svm), nil
}

func MustHexToAddress(hex string) Address {
	addr, err := HexToAddress(hex)
	if err != nil {
		panic(err)
	}

	return addr
}

func MustBase58ToAddress(base58 string) Address {
	addr, err := Base58ToAddress(base58)
	if err != nil {
		panic(err)
	}

	return addr
}

func ParseAddress(address string) (Address, error) {
	if a, err := HexToAddress(address); err == nil {
		return a, nil
	} else if a, err := Base58ToAddress(address); err == nil {
		return a, nil
	}

	return Address{}, errors.New("invalid address", "address", address)
}
