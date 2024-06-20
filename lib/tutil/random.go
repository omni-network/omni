package tutil

import (
	"crypto/rand"

	"github.com/ethereum/go-ethereum/common"
)

// RandomBytes returns a random byte slice of length l.
func RandomBytes(l int) []byte {
	resp := make([]byte, l)
	_, _ = rand.Read(resp)

	return resp
}

// RandomHash returns a random 32-byte 256-bit hash.
func RandomHash() common.Hash {
	var resp common.Hash
	_, _ = rand.Read(resp[:])

	return resp
}

// RandomAddress returns a random 20-byte ethereum address.
func RandomAddress() common.Address {
	var resp common.Address
	_, _ = rand.Read(resp[:])

	return resp
}
