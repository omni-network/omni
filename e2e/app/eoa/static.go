package eoa

import (
	"math/big"

	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/crypto"
)

//nolint:gochecknoglobals // Static mappings.
var statics = map[netconf.ID][]Account{
	netconf.Devnet: {
		{
			Type:          TypeRemote,
			Role:          RoleCreate3Deployer,
			Address:       devnetCreate3Deployer,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(10),
			TargetBalance: big.NewInt(100)},

		{
			Type:          TypeRemote,
			Role:          RoleDeployer,
			Address:       devnetDeployer,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(10),
			TargetBalance: big.NewInt(100),
		},
		{
			Type:          TypeRemote,
			Role:          RoleAdmin,
			Address:       devnetAdmin,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(10),
			TargetBalance: big.NewInt(100),
		},
		{
			Type:          TypeWellKnown,
			Role:          RoleRelayer,
			Address:       crypto.PubkeyToAddress((anvil.DevPrivateKey5()).PublicKey),
			PrivateKey:    anvil.DevPrivateKey5(),
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(10),
			TargetBalance: big.NewInt(100),
		},
		{
			Type:          TypeWellKnown,
			Role:          RoleMonitor,
			Address:       crypto.PubkeyToAddress((anvil.DevPrivateKey6()).PublicKey),
			PrivateKey:    anvil.DevPrivateKey6(),
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(10),
			TargetBalance: big.NewInt(100),
		},
		{
			Type:          TypeRemote,
			Role:          RoleFbDev,
			Address:       addr(fbDev),
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(10),
			TargetBalance: big.NewInt(100),
		},
	},
	netconf.Staging: {
		{
			Type:          TypeRemote,
			Role:          RoleCreate3Deployer,
			Address:       stagingCreate3Deployer,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(2),
			TargetBalance: big.NewInt(10)},

		{
			Type:          TypeRemote,
			Role:          RoleDeployer,
			Address:       stagingDeployer,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(2),
			TargetBalance: big.NewInt(10),
		},
		{
			Type:          TypeRemote,
			Role:          RoleAdmin,
			Address:       stagingAdmin,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeSecret,
			Role:          RoleRelayer,
			Address:       stagingRelayer,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeSecret,
			Role:          RoleMonitor,
			Address:       stagingMonitor,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeRemote,
			Role:          RoleFbDev,
			Address:       addr(fbDev),
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(5),
			TargetBalance: big.NewInt(10),
		},
	},
	netconf.Testnet: {
		{
			Type:          TypeRemote,
			Role:          RoleCreate3Deployer,
			Address:       testnetCreate3Deployer,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(2),
			TargetBalance: big.NewInt(10),
		},
		{
			Type:          TypeRemote,
			Role:          RoleDeployer,
			Address:       testnetDeployer,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(2),
			TargetBalance: big.NewInt(10),
		},
		{
			Type:          TypeRemote,
			Role:          RoleAdmin,
			Address:       testnetAdmin,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeSecret,
			Role:          RoleRelayer,
			Address:       testnetRelayer,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeSecret,
			Role:          RoleMonitor,
			Address:       testnetMonitor,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
	},
	netconf.Mainnet: {
		{
			Type:          TypeSecret,
			Role:          RoleRelayer,
			Address:       mainnetRelayer,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeSecret,
			Role:          RoleMonitor,
			Address:       mainnetMonitor,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
	},
}
