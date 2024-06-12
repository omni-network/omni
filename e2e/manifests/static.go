package manifests

import (
	_ "embed"
)

var (
	//go:embed devnet0.toml
	devnet0 []byte

	//go:embed omega.toml
	omega []byte
)

// Devnet0 returns the devnet0.toml manifest bytes.
func Devnet0() []byte {
	return devnet0
}

// Omega returns the omega.toml manifest bytes.
func Omega() []byte {
	return omega
}
