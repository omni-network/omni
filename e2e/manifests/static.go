package manifests

import (
	_ "embed"
)

//go:embed devnet0.toml
var devnet0 []byte

func Devnet0() []byte {
	return devnet0
}
