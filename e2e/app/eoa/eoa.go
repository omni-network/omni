// Package eoa defines well-known (non-fireblocks) eoa private keys used in an omni network.
package eoa

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/omni-network/omni/e2e/app/key"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Role string

const (
	RoleRelayer         Role = "relayer"
	RoleMonitor         Role = "monitor"
	RoleCreate3Deployer Role = "create3-deployer"
	RoleDeployer        Role = "deployer"
	RoleAdmin           Role = "admin"
	RoleFbDev           Role = "fb-dev"
)

type Type string

const (
	TypeRemote    Type = "remote"     // stored in (fireblocks) accessible via API to sign
	TypeSecret    Type = "secret"     // stored in GCP can be downloaded to disk
	TypeWellKnown Type = "well-known" // well-known eoa private keys in the repo
)

type ChainSelector func(netconf.Network) []netconf.Chain

var ChainSelectorAll = func(network netconf.Network) []netconf.Chain { return network.Chains }
var ChainSelectorL1 = func(network netconf.Network) []netconf.Chain { return network.EVMChains() /* todo: impl l1chains*/ }

// Account is a well-known EOA account.
type Account struct {
	Type    Type
	Role    Role
	Address common.Address

	Chains        ChainSelector
	MinBalance    *big.Int
	TargetBalance *big.Int
}

func (r Role) Verify() error {
	if r != RoleRelayer && r != RoleMonitor {
		return errors.New("invalid role", "role", r)
	}

	return nil
}

var (
	devnetKeys = map[Role]*ecdsa.PrivateKey{
		RoleRelayer: anvil.DevPrivateKey5(),
		RoleMonitor: anvil.DevPrivateKey6(),
	}
)

var secureAddrs = map[netconf.ID]map[Role]common.Address{
	netconf.Staging: {
		RoleRelayer: common.HexToAddress("0xfE921e06Ed0a22c035b4aCFF0A5D3a434A330c96"),
		RoleMonitor: common.HexToAddress("0x0De553555Fa19d787Af4273B18bDB77282D618c4"),
	},
	netconf.Testnet: {
		RoleRelayer: common.HexToAddress("0x01654f55E4F5E2f2ff8080702676F1984CBf7d8a"),
		RoleMonitor: common.HexToAddress("0x12Dc870b3F5b7f810c3d1e489e32a64d4E25AaCA"),
	},
	netconf.Mainnet: {
		TypeMonitor: common.HexToAddress("0x07082fcbFA5F5AC9FBc03A48B7f6391441DB8332"),
		TypeRelayer: common.HexToAddress("0x07804D7B8be635c0C68Cdf3E946114221B12f4F7"),
	},
}

// MustAddress returns the address for the EOA identified by the network and role.
func MustAddress(network netconf.ID, role Role) common.Address {
	resp, _ := Address(network, role)
	return resp
}

// Address returns the address for the EOA identified by the network and role.
func Address(network netconf.ID, role Role) (common.Address, bool) {
	if network == netconf.Devnet {
		return crypto.PubkeyToAddress(devnetKeys[role].PublicKey), true
	}

	resp, ok := secureAddrs[network][role]

	return resp, ok
}

// PrivateKey returns the private key for the EOA identified by the network and role.
func PrivateKey(ctx context.Context, network netconf.ID, role Role) (*ecdsa.PrivateKey, error) {
	if network == netconf.Devnet {
		return devnetKeys[role], nil
	}

	addr, ok := secureAddrs[network][role]
	if !ok {
		return nil, errors.New("eoa key not defined", "network", network, "role", role)
	}

	k, err := key.Download(ctx, network, string(role), key.EOA, addr.Hex())
	if err != nil {
		return nil, errors.Wrap(err, "download key")
	}

	return k.ECDSA()
}
