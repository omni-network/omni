package eoa

import (
	"crypto/ecdsa"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

// Admin returns the address of the admin for the given network.
// NOTE: this relies on the fact that we use the same admin account for all "admin" roles.
func Admin(network netconf.ID) (common.Address, error) {
	for _, account := range statics[network] {
		if account.Role == RoleAdmin {
			return account.Address, nil
		}
	}

	return common.Address{}, errors.New("no admin account for network")
}

// Funder returns the address of the funder account.
func Funder() common.Address {
	return common.HexToAddress(fbFunder)
}

func fbDevAcc() []Account {
	return []Account{
		{
			Type:    TypeRemote,
			Role:    RoleFbDev,
			Address: common.HexToAddress(fbDev),
		},
	}
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
