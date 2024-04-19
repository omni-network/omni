package eoa

import (
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
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
	// fbDev is the address of the fireblocks "dev" account.
	fbDev = "0x7a6cF389082dc698285474976d7C75CAdE08ab7e"
	// fbFunder is the address of the fireblocks "funder" account.
	fbFunder = "0xf63316AA39fEc9D2109AB0D9c7B1eE3a6F60AEA4"
)

//nolint:gochecknoglobals // Static addresses
var (
	// Admin used as contract owner.
	mainnetAdmin = common.HexToAddress("0x0")
	testnetAdmin = common.HexToAddress(fbTestnetAdmin)
	stagingAdmin = common.HexToAddress(fbStagingAdmin)
	devnetAdmin  = anvil.DevAccount2()

	// Create3 Deployer address that can deploy the create3 factory.
	mainnetCreate3Deployer = common.HexToAddress("0x0")
	testnetCreate3Deployer = common.HexToAddress(fbTestnetCreate3Deployer)
	stagingCreate3Deployer = common.HexToAddress(fbStagingCreate3Deployer)
	devnetCreate3Deployer  = anvil.DevAccount0()

	// Deployer address that can deploy protocol contracts via Create3 factory.
	mainnetDeployer = common.HexToAddress("0x0")
	testnetDeployer = common.HexToAddress(fbTestnetDeployer)
	stagingDeployer = common.HexToAddress(fbStagingDeployer)
	devnetDeployer  = anvil.DevAccount1()

	stagingRelayer = common.HexToAddress("0xfE921e06Ed0a22c035b4aCFF0A5D3a434A330c96")
	stagingMonitor = common.HexToAddress("0x0De553555Fa19d787Af4273B18bDB77282D618c4")

	testnetRelayer = common.HexToAddress("0x01654f55E4F5E2f2ff8080702676F1984CBf7d8a")
	testnetMonitor = common.HexToAddress("0x12Dc870b3F5b7f810c3d1e489e32a64d4E25AaCA")

	mainnetMonitor = common.HexToAddress("0x07082fcbFA5F5AC9FBc03A48B7f6391441DB8332")
	mainnetRelayer = common.HexToAddress("0x07804D7B8be635c0C68Cdf3E946114221B12f4F7")

	fbFunderAddr = common.HexToAddress(fbFunder)
)

// Admin returns the address of the admin for the given network.
func Admin(network netconf.ID) (common.Address, error) {
	switch network {
	case netconf.Mainnet:
		return mainnetAdmin, nil
	case netconf.Testnet:
		return testnetAdmin, nil
	case netconf.Staging:
		return stagingAdmin, nil
	case netconf.Devnet:
		return devnetAdmin, nil
	default:
		return common.Address{}, errors.New("unknown network", "network", network)
	}
}

// Funder returns the address of the funder account.
func Funder() common.Address {
	return fbFunderAddr
}
