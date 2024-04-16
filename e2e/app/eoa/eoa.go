// Package eoa defines well-known (non-fireblocks) eoa private keys used in an omni network.
package eoa

import (
	"context"
	"crypto/ecdsa"

	"github.com/omni-network/omni/e2e/app/key"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Type string

const (
	TypeRelayer Type = "relayer"
	TypeMonitor Type = "monitor"
)

func (t Type) Verify() error {
	if t != TypeRelayer && t != TypeMonitor {
		return errors.New("invalid type", "type", t)
	}

	return nil
}

var (
	devnetKeys = map[Type]*ecdsa.PrivateKey{
		TypeRelayer: anvil.DevPrivateKey5(),
		TypeMonitor: anvil.DevPrivateKey6(),
	}
)

var secureAddrs = map[netconf.ID]map[Type]common.Address{
	netconf.Staging: {
		TypeRelayer: common.HexToAddress("0xfE921e06Ed0a22c035b4aCFF0A5D3a434A330c96"),
		TypeMonitor: common.HexToAddress("0x0De553555Fa19d787Af4273B18bDB77282D618c4"),
	},
	netconf.Testnet: {
		TypeRelayer: common.HexToAddress("0x01654f55E4F5E2f2ff8080702676F1984CBf7d8a"),
		TypeMonitor: common.HexToAddress("0x12Dc870b3F5b7f810c3d1e489e32a64d4E25AaCA"),
	},
	netconf.Mainnet: {
		TypeMonitor: common.HexToAddress("0x07082fcbFA5F5AC9FBc03A48B7f6391441DB8332"),
		TypeRelayer: common.HexToAddress("0x07804D7B8be635c0C68Cdf3E946114221B12f4F7"),
	},
}

// MustAddress returns the address for the EOA identified by the network and type.
func MustAddress(network netconf.ID, typ Type) common.Address {
	resp, _ := Address(network, typ)
	return resp
}

// Address returns the address for the EOA identified by the network and type.
func Address(network netconf.ID, typ Type) (common.Address, bool) {
	if network == netconf.Devnet {
		return crypto.PubkeyToAddress(devnetKeys[typ].PublicKey), true
	}

	resp, ok := secureAddrs[network][typ]

	return resp, ok
}

// PrivateKey returns the private key for the EOA identified by the network and type.
func PrivateKey(ctx context.Context, network netconf.ID, typ Type) (*ecdsa.PrivateKey, error) {
	if network == netconf.Devnet {
		return devnetKeys[typ], nil
	}

	addr, ok := secureAddrs[network][typ]
	if !ok {
		return nil, errors.New("eoa key not defined", "network", network, "type", typ)
	}

	k, err := key.Download(ctx, network, string(typ), key.EOA, addr.Hex())
	if err != nil {
		return nil, errors.Wrap(err, "download key")
	}

	return k.ECDSA()
}
