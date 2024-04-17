package eoa

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/netconf"
)

const (
	// address of the "staging-create3-deployer" fireblocks account.
	fbStagingCreate3Deployer = "0xC8103859Ac7CB547d70307EdeF1A2319FC305fdC"
	// address of the "staging-deployer" fireblocks account.
	fbStagingDeployer = "0x274c4B3e5d27A65196d63964532366872F81D261"
	// address of the "staging-owner" fireblocks account.
	fbStagingAdmin = "0x4891925c4f13A34FC26453FD168Db80aF3273014"
	// address of the "testnet-create3-deployer" fireblocks account.
	fbTestnetCreate3Deployer = "0xeC5134556da0797A5C5cD51DD622b689Cac97Fe9"
	// address of the "testnet-deployer" fireblocks account.
	fbTestnetDeployer = "0x0CdCc644158b7D03f40197f55454dc7a11Bd92c1"
	// address of the "testnet-owner" fireblocks account.
	fbTestnetAdmin = "0xEAD625eB2011394cdD739E91Bf9D51A7169C22F5"

	fbDev = "0x7a6cF389082dc698285474976d7C75CAdE08ab7e"
)

//nolint:gochecknoglobals // Static addresses
var (
	// Admin - used as contract owner.

	MainnetAdmin = addr("0x0")
	TestnetAdmin = addr(fbTestnetAdmin)
	StagingAdmin = addr(fbStagingAdmin)
	DevnetAdmin  = anvil.DevAccount2()

	// Create3 Deployer - addrress that can deploy the create3 factory.

	MainnetCreate3Deployer = addr("0x0")
	TestnetCreate3Deployer = addr(fbTestnetCreate3Deployer)
	StagingCreate3Deployer = addr(fbStagingCreate3Deployer)
	DevnetCreate3Deployer  = anvil.DevAccount0()

	// Deployer - address that can deploy protocol contracts via Create3 factory.

	MainnetDeployer = addr("0x0")
	TestnetDeployer = addr(fbTestnetDeployer)
	StagingDeployer = addr(fbStagingDeployer)
	DevnetDeployer  = anvil.DevAccount1()
)

//nolint:gochecknoglobals // Static mappings.
var statics = map[netconf.ID][]Account{
	netconf.Staging: {
		{
			Type:          TypeRemote,
			Role:          RoleCreate3Deployer,
			Address:       StagingCreate3Deployer,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5)},

		{
			Type:          TypeRemote,
			Role:          RoleDeployer,
			Address:       StagingDeployer,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeRemote,
			Role:          RoleAdmin,
			Address:       StagingAdmin,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeSecret,
			Role:          RoleRelayer,
			Address:       addr("0xfE921e06Ed0a22c035b4aCFF0A5D3a434A330c96"),
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeSecret,
			Role:          RoleMonitor,
			Address:       addr("0x0De553555Fa19d787Af4273B18bDB77282D618c4"),
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeRemote,
			Role:          RoleFbDev,
			Address:       addr(fbDev),
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
	},
	netconf.Testnet: {
		{
			Type:          TypeRemote,
			Role:          RoleCreate3Deployer,
			Address:       TestnetCreate3Deployer,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeRemote,
			Role:          RoleDeployer,
			Address:       TestnetDeployer,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeRemote,
			Role:          RoleAdmin,
			Address:       TestnetAdmin,
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeSecret,
			Role:          RoleRelayer,
			Address:       addr("0x01654f55E4F5E2f2ff8080702676F1984CBf7d8a"),
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeSecret,
			Role:          RoleMonitor,
			Address:       addr("0x12Dc870b3F5b7f810c3d1e489e32a64d4E25AaCA"),
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
		{
			Type:          TypeRemote,
			Role:          RoleFbDev,
			Address:       addr(fbDev),
			Chains:        ChainSelectorAll,
			MinBalance:    big.NewInt(1),
			TargetBalance: big.NewInt(5),
		},
	},
}

func MustAddresses(network netconf.Network) []common.Address {
	accounts := statics[network.ID]
	var addresses []common.Address
	for _, account := range accounts {
		addresses = append(addresses, account.Address)

	}
	return addresses
}

func addr(hex string) common.Address {
	return common.HexToAddress(hex)
}
