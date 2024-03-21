package create3

import (
	"github.com/ethereum/go-ethereum/crypto"
)

// HashSalt returns the [32]byte hash of the salt string.
func HashSalt(s string) [32]byte {
	var h [32]byte
	copy(h[:], crypto.Keccak256([]byte(s)))

	return h
}
