package eoa

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

var (
	gwei100 = new(big.Int).Mul(big.NewInt(100), big.NewInt(params.GWei))
	gwei500 = new(big.Int).Mul(big.NewInt(500), big.NewInt(params.GWei))

	ether1   = new(big.Int).Mul(big.NewInt(1), big.NewInt(params.Ether))
	ether5   = new(big.Int).Mul(big.NewInt(5), big.NewInt(params.Ether))
	ether10  = new(big.Int).Mul(big.NewInt(10), big.NewInt(params.Ether))
	ether100 = new(big.Int).Mul(big.NewInt(100), big.NewInt(params.Ether))

	minBalanceSmall    = gwei100
	targetBalanceSmall = gwei500

	minBalanceMedium    = ether1
	targetBalanceMedium = ether5

	minBalanceLarge    = ether10
	targetBalanceLarge = ether100
)

//nolint:gochecknoglobals // Static mappings.
var statics = map[netconf.ID][]Account{
	netconf.Devnet: {
		{
			Type:          TypeRemote,
			Role:          RoleCreate3Deployer,
			Address:       devnetCreate3Deployer,
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(10),
			TargetBalance: big.NewInt(100)},

		{
			Type:          TypeRemote,
			Role:          RoleDeployer,
			Address:       devnetDeployer,
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(10),
			TargetBalance: big.NewInt(100),
		},
		{
			Type:          TypeRemote,
			Role:          RoleAdmin,
			Address:       devnetAdmin,
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(10),
			TargetBalance: big.NewInt(100),
		},
		{
			Type:          TypeWellKnown,
			Role:          RoleRelayer,
			Address:       anvil.DevAccount5(),
			privateKey:    anvil.DevPrivateKey5(),
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(10),
			TargetBalance: big.NewInt(100),
		},
		{
			Type:          TypeWellKnown,
			Role:          RoleMonitor,
			Address:       anvil.DevAccount6(),
			privateKey:    anvil.DevPrivateKey6(),
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(10),
			TargetBalance: big.NewInt(100),
		},
		{
			Type:          TypeRemote,
			Role:          RoleFbDev,
			Address:       common.HexToAddress(fbDev),
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(10),
			TargetBalance: big.NewInt(100),
		},
	},
	netconf.Staging: {
		{
			Type:          TypeRemote,
			Role:          RoleCreate3Deployer,
			Address:       stagingCreate3Deployer,
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(2),
			TargetBalance: big.NewInt(10)},

		{
			Type:          TypeRemote,
			Role:          RoleDeployer,
			Address:       stagingDeployer,
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(2),
			TargetBalance: big.NewInt(10),
		},
		{
			Type:          TypeRemote,
			Role:          RoleAdmin,
			Address:       stagingAdmin,
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeSecret,
			Role:          RoleRelayer,
			Address:       stagingRelayer,
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeSecret,
			Role:          RoleMonitor,
			Address:       stagingMonitor,
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeRemote,
			Role:          RoleFbDev,
			Address:       common.HexToAddress(fbDev),
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(5),
			TargetBalance: big.NewInt(10),
		},
	},
	netconf.Testnet: {
		{
			Type:          TypeRemote,
			Role:          RoleCreate3Deployer,
			Address:       testnetCreate3Deployer,
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(2),
			TargetBalance: big.NewInt(10),
		},
		{
			Type:          TypeRemote,
			Role:          RoleDeployer,
			Address:       testnetDeployer,
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(2),
			TargetBalance: big.NewInt(10),
		},
		{
			Type:          TypeRemote,
			Role:          RoleAdmin,
			Address:       testnetAdmin,
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeSecret,
			Role:          RoleRelayer,
			Address:       testnetRelayer,
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeSecret,
			Role:          RoleMonitor,
			Address:       testnetMonitor,
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
	},
	netconf.Mainnet: {
		{
			Type:          TypeSecret,
			Role:          RoleRelayer,
			Address:       mainnetRelayer,
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeSecret,
			Role:          RoleMonitor,
			Address:       mainnetMonitor,
			Chains:        chainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
	},
}
