// Package cast provides save casting functions for converting between types without panicking.
package cast

import (
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
)

// Array65 casts a slice to an array of length 65.
func Array65[A any](slice []A) ([65]A, error) {
	if len(slice) == 65 {
		return [65]A(slice), nil
	}

	return [65]A{}, errors.New("slice length not 65", "len", len(slice))
}

// Array48 casts a slice to an array of length 48.
func Array48[A any](slice []A) ([48]A, error) {
	if len(slice) == 48 {
		return [48]A(slice), nil
	}

	return [48]A{}, errors.New("slice length not 48", "len", len(slice))
}

// Must32 casts a slice to an array of length 32.
func Must32[A any](slice []A) [32]A {
	arr, err := Array32(slice)
	if err != nil {
		panic(err)
	}

	return arr
}

// Must20 casts a slice to an array of length 20.
func Must20[A any](slice []A) [20]A {
	arr, err := Array20(slice)
	if err != nil {
		panic(err)
	}

	return arr
}

// EthHash casts a byte slice to an Ethereum hash.
func EthHash(b []byte) (common.Hash, error) {
	resp, err := Array32(b)
	if err != nil {
		return common.Hash{}, errors.New("invalid hash length", "len", len(b))
	}

	return resp, nil
}

// Array64 casts a slice to an array of length 64.
func Array64[A any](slice []A) ([64]A, error) {
	if len(slice) == 64 {
		return [64]A(slice), nil
	}

	return [64]A{}, errors.New("slice length not 64", "len", len(slice))
}

// Array32 casts a slice to an array of length 32.
func Array32[A any](slice []A) ([32]A, error) {
	if len(slice) == 32 {
		return [32]A(slice), nil
	}

	return [32]A{}, errors.New("slice length not 32", "len", len(slice))
}

func EthAddress32(addr common.Address) [32]byte {
	var resp [32]byte
	copy(resp[12:], addr[:])

	return resp
}

// EthAddress casts a byte slice to an Ethereum address.
func EthAddress(b []byte) (common.Address, error) {
	resp, err := Array20(b)
	if err != nil {
		return common.Address{}, errors.New("invalid address length", "len", len(b))
	}

	return resp, nil
}

// MustEthAddress casts a byte slice to an Ethereum address.
func MustEthAddress(b []byte) common.Address {
	addr, err := EthAddress(b)
	if err != nil {
		panic(err)
	}

	return addr
}

// Array20 casts a slice to an array of length 32.
func Array20[A any](slice []A) ([20]A, error) {
	if len(slice) == 20 {
		return [20]A(slice), nil
	}

	return [20]A{}, errors.New("slice length not 20", "len", len(slice))
}

// Array8 casts a slice to an array of length 8.
func Array8[A any](slice []A) ([8]A, error) {
	if len(slice) == 8 {
		return [8]A(slice), nil
	}

	return [8]A{}, errors.New("slice length not 8", "len", len(slice))
}

// Array4 casts a slice to an array of length 4.
func Array4[A any](slice []A) ([4]A, error) {
	if len(slice) == 4 {
		return [4]A(slice), nil
	}

	return [4]A{}, errors.New("slice length not 4", "len", len(slice))
}

// BytesToAddress converts a byte slice to an Ethereum address.
// Handles both 20-byte addresses and 32-byte storage slots (takes last 20 bytes).
func BytesToAddress(b []byte) (common.Address, error) {
	switch len(b) {
	case 20:
		return EthAddress(b)
	case 32:
		// Take the last 20 bytes (addresses are right-aligned in 32-byte slots)
		return EthAddress(b[12:])
	default:
		return common.Address{}, errors.New("invalid bytes length for address", "len", len(b))
	}
}

// MustBytesToAddress converts a byte slice to an Ethereum address, panicking on error.
func MustBytesToAddress(b []byte) common.Address {
	addr, err := BytesToAddress(b)
	if err != nil {
		panic(err)
	}

	return addr
}
