package eoa

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

// Funder returns the address of the funder account.
func Funder() common.Address {
	return common.HexToAddress(fbFunder)
}

func dummy(roles ...Role) []Account {
	var resp []Account
	for _, role := range roles {
		resp = append(resp, Account{
			Type:    TypeWellKnown,
			Role:    role,
			Address: common.HexToAddress(ZeroXDead),
		})
	}

	return resp
}

func remote(hex string, roles ...Role) []Account {
	var resp []Account
	for _, role := range roles {
		resp = append(resp, Account{
			Type:    TypeRemote,
			Role:    role,
			Address: common.HexToAddress(hex),
		})
	}

	return resp
}

func wellKnown(pk *ecdsa.PrivateKey, roles ...Role) []Account {
	var resp []Account
	for _, role := range roles {
		resp = append(resp, Account{
			Type:       TypeWellKnown,
			Role:       role,
			Address:    ethcrypto.PubkeyToAddress(pk.PublicKey),
			privateKey: pk,
		})
	}

	return resp
}

func secret(hex string, roles ...Role) []Account {
	var resp []Account
	for _, role := range roles {
		resp = append(resp, Account{
			Type:    TypeSecret,
			Role:    role,
			Address: common.HexToAddress(hex),
		})
	}

	return resp
}

func flatten[T any](slices ...[]T) []T {
	var resp []T
	for _, slice := range slices {
		resp = append(resp, slice...)
	}

	return resp
}
