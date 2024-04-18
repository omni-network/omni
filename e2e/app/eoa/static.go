package eoa

import (
	"math/big"

	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

var (
	gwei10000 = new(big.Int).Mul(big.NewInt(10000), big.NewInt(params.GWei))
	gwei50000 = new(big.Int).Mul(big.NewInt(50000), big.NewInt(params.GWei))

	ether1   = new(big.Int).Mul(big.NewInt(1), big.NewInt(params.Ether))
	ether5   = new(big.Int).Mul(big.NewInt(5), big.NewInt(params.Ether))
	ether10  = new(big.Int).Mul(big.NewInt(10), big.NewInt(params.Ether))
	ether100 = new(big.Int).Mul(big.NewInt(100), big.NewInt(params.Ether))

	minBalanceSmall    = gwei10000
	targetBalanceSmall = gwei50000

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
			MinBalance:    minBalanceLarge,
			TargetBalance: targetBalanceLarge,
		},
		{
			Type:          TypeRemote,
			Role:          RoleDeployer,
			Address:       devnetDeployer,
			Chains:        chainSelectorAll,
			MinBalance:    minBalanceLarge,
			TargetBalance: targetBalanceLarge,
		},
		{
			Type:          TypeRemote,
			Role:          RoleAdmin,
			Address:       devnetAdmin,
			Chains:        chainSelectorAll,
			MinBalance:    minBalanceLarge,
			TargetBalance: targetBalanceLarge,
		},
		{
			Type:          TypeWellKnown,
			Role:          RoleRelayer,
			Address:       anvil.DevAccount5(),
			privateKey:    anvil.DevPrivateKey5(),
			Chains:        chainSelectorAll,
			MinBalance:    minBalanceLarge,
			TargetBalance: targetBalanceLarge,
		},
		{
			Type:          TypeWellKnown,
			Role:          RoleMonitor,
			Address:       anvil.DevAccount6(),
			privateKey:    anvil.DevPrivateKey6(),
			Chains:        chainSelectorAll,
			MinBalance:    minBalanceLarge,
			TargetBalance: targetBalanceLarge,
		},
		{
			Type:          TypeRemote,
			Role:          RoleFbDev,
			Address:       common.HexToAddress(fbDev),
			Chains:        chainSelectorAll,
			MinBalance:    minBalanceLarge,
			TargetBalance: targetBalanceLarge,
		},
	},
	netconf.Staging: {
		{
			Type:          TypeRemote,
			Role:          RoleCreate3Deployer,
			Address:       stagingCreate3Deployer,
			Chains:        chainSelectorAll,
			MinBalance:    minBalanceMedium,
			TargetBalance: targetBalanceMedium,
		},

		{
			Type:          TypeRemote,
			Role:          RoleDeployer,
			Address:       stagingDeployer,
			Chains:        chainSelectorAll,
			MinBalance:    minBalanceMedium,
			TargetBalance: targetBalanceMedium,
		},
		{
			Type:          TypeRemote,
			Role:          RoleAdmin,
			Address:       stagingAdmin,
			Chains:        chainSelectorAll,
			MinBalance:    minBalanceMedium,
			TargetBalance: targetBalanceMedium,
		},
		{
			Type:          TypeSecret,
			Role:          RoleRelayer,
			Address:       stagingRelayer,
			Chains:        chainSelectorAll,
			MinBalance:    minBalanceMedium,
			TargetBalance: targetBalanceMedium,
		},
		{
			Type:          TypeSecret,
			Role:          RoleMonitor,
			Address:       stagingMonitor,
			Chains:        chainSelectorAll,
			MinBalance:    minBalanceSmall,
			TargetBalance: targetBalanceSmall,
		},
		{
			Type:          TypeRemote,
			Role:          RoleFbDev,
			Address:       common.HexToAddress(fbDev),
			Chains:        chainSelectorAll,
			MinBalance:    targetBalanceMedium,
			TargetBalance: targetBalanceMedium,
		},
	},
	netconf.Testnet: {
		{
			Type:          TypeRemote,
			Role:          RoleCreate3Deployer,
			Address:       testnetCreate3Deployer,
			Chains:        chainSelectorAll,
			MinBalance:    minBalanceMedium,
			TargetBalance: targetBalanceMedium,
		},
		{
			Type:          TypeRemote,
			Role:          RoleDeployer,
			Address:       testnetDeployer,
			Chains:        chainSelectorAll,
			MinBalance:    minBalanceMedium,
			TargetBalance: targetBalanceMedium,
		},
		{
			Type:          TypeRemote,
			Role:          RoleAdmin,
			Address:       testnetAdmin,
			Chains:        chainSelectorAll,
			MinBalance:    minBalanceMedium,
			TargetBalance: targetBalanceMedium,
		},
		{
			Type:          TypeSecret,
			Role:          RoleRelayer,
			Address:       testnetRelayer,
			Chains:        chainSelectorAll,
			MinBalance:    minBalanceMedium,
			TargetBalance: targetBalanceMedium,
		},
		{
			Type:          TypeSecret,
			Role:          RoleMonitor,
			Address:       testnetMonitor,
			Chains:        chainSelectorAll,
			MinBalance:    minBalanceSmall,
			TargetBalance: targetBalanceSmall,
		},
	},
	netconf.Mainnet: {
		{
			Type:          TypeSecret,
			Role:          RoleRelayer,
			Address:       mainnetRelayer,
			Chains:        chainSelectorAll,
			MinBalance:    minBalanceMedium,
			TargetBalance: targetBalanceMedium,
		},
		{
			Type:          TypeSecret,
			Role:          RoleMonitor,
			Address:       mainnetMonitor,
			Chains:        chainSelectorAll,
			MinBalance:    minBalanceSmall,
			TargetBalance: targetBalanceSmall,
		},
	},
}
