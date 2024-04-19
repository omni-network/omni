// Package eoa defines well-known (non-fireblocks) eoa private keys used in an omni network.
package eoa

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/omni-network/omni/e2e/app/key"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

type Role string

const (
	RoleRelayer         Role = "relayer"
	RoleMonitor         Role = "monitor"
	RoleCreate3Deployer Role = "create3-deployer"
	RoleDeployer        Role = "deployer"
	RoleProxyAdminOwner Role = "proxy-admin-owner"
	RolePortalAdmin     Role = "portal-admin"
	RoleAVSAdmin        Role = "avs-admin"
	RoleFbDev           Role = "fb-dev"
)

func AllRoles() []Role {
	return []Role{
		RoleRelayer,
		RoleMonitor,
		RoleCreate3Deployer,
		RoleDeployer,
		RoleProxyAdminOwner,
		RolePortalAdmin,
		RoleAVSAdmin,
		RoleFbDev,
	}
}

func (r Role) Verify() error {
	for _, role := range AllRoles() {
		if r == role {
			return nil
		}
	}

	return errors.New("invalid role", "role", r)
}

type Type string

const (
	TypeRemote    Type = "remote"     // stored in (fireblocks) accessible via API to sign
	TypeSecret    Type = "secret"     // stored in GCP can be downloaded to disk
	TypeWellKnown Type = "well-known" // well-known eoa private keys in the repo
)

type chainSelector func(netconf.Network) []netconf.Chain

var chainSelectorAll = func(network netconf.Network) []netconf.Chain { return network.EVMChains() }

//nolint:unused // will be used.
var chainSelectorL1 = func(network netconf.Network) []netconf.Chain {
	c, ok := network.EthereumChain()
	if !ok {
		return nil
	}

	return []netconf.Chain{c}
}

// Account defines a EOA account used within the Omni network.
type Account struct {
	Type       Type
	Role       Role
	Address    common.Address
	privateKey *ecdsa.PrivateKey // only for devnet (well-known type)

	Chains        chainSelector // all vs specific chains
	MinBalance    *big.Int      // check if below this balance
	TargetBalance *big.Int      // fund to this balance
}

// privKey returns the private key for the account.
func (a Account) privKey() *ecdsa.PrivateKey {
	return a.privateKey
}

// MustAddress returns the address for the EOA identified by the network and role.
func MustAddress(network netconf.ID, role Role) common.Address {
	resp, ok := Address(network, role)
	if !ok {
		panic(errors.New("eoa address not defined", "network", network, "role", role))
	}

	return resp
}

// Address returns the address for the EOA identified by the network and role.
func Address(network netconf.ID, role Role) (common.Address, bool) {
	accounts, ok := statics[network]
	if !ok {
		return common.Address{}, false
	}

	for _, account := range accounts {
		if account.Role == role {
			return account.Address, true
		}
	}

	return common.Address{}, false
}

// PrivateKey returns the private key for the EOA identified by the network and role.
func PrivateKey(ctx context.Context, network netconf.ID, role Role) (*ecdsa.PrivateKey, error) {
	acc, ok := AccountForRole(network, role)
	if !ok {
		return nil, errors.New("eoa key not defined", "network", network, "role", role)
	}
	if acc.Type == TypeWellKnown {
		return acc.privKey(), nil
	} else if acc.Type == TypeRemote {
		return nil, errors.New("private key not available for remote keys", "network", network, "role", role)
	}

	k, err := key.Download(ctx, network, string(role), key.EOA, acc.Address.Hex())
	if err != nil {
		return nil, errors.Wrap(err, "download key")
	}

	return k.ECDSA()
}

// AccountForRole returns the account for the network and role.
func AccountForRole(network netconf.ID, role Role) (Account, bool) {
	accounts, ok := statics[network]
	if !ok {
		return Account{}, false
	}
	for _, account := range accounts {
		if account.Role == role {
			return account, true
		}
	}

	return Account{}, false
}

// MustAddresses returns the addresses for the network and roles.
func MustAddresses(network netconf.ID, roles ...Role) []common.Address {
	accounts := statics[network]
	var addresses []common.Address
	for _, role := range roles {
		for _, account := range accounts {
			if account.Role == role {
				addresses = append(addresses, account.Address)
			}
		}
	}

	return addresses
}

// AllAccounts returns all accounts for the network.
func AllAccounts(network netconf.ID) ([]Account, bool) {
	acc, ok := statics[network]
	return acc, ok
}
