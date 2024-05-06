package manifests

import (
	_ "embed"
)

var (
	//go:embed devnet0.toml
	devnet0 []byte

	//go:embed testnet.toml
	testnet []byte
)

// Devnet0 returns the devnet0.toml manifest bytes.
func Devnet0() []byte {
	return devnet0
}

// Testnet returns the testnet.toml manifest bytes.
func Testnet() []byte {
	return testnet
}
