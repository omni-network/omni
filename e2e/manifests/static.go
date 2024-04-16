package manifests

import (
	_ "embed"
)

//go:embed devnet1.toml
var devnet1 []byte

func Devnet1() []byte {
	return devnet1
}
