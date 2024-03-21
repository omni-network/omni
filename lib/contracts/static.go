package contracts

import (
	"github.com/omni-network/omni/lib/anvil"

	"github.com/ethereum/go-ethereum/common"
)

//nolint:gochecknoglobals // Static addresses
var (
	DevnetCreate3Factory  = common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	MainnetCreate3Factory = common.HexToAddress("0x0") // TODO
	TestnetCreate3Factory = common.HexToAddress("0x0") // TODO

	DevnetProxyAdmin  = common.HexToAddress("0x733AA9e7E4025E9F69DBEd9e05155e081D720565")
	MainnetProxyAdmin = common.HexToAddress("0x0") // TODO
	TestnetProxyAdmin = common.HexToAddress("0x0") // TODO

	DevnetProxyAdminOwner  = anvil.Account2
	MainnetProxyAdminOwner = common.HexToAddress("0x0") // TODO
	TestnetProxyAdminOwner = common.HexToAddress("0x0") // TODO

	// addrress that can deploy the create3 factory.
	DevnetCreate3Deployer  = anvil.Account0
	MainnetCreate3Deployer = common.HexToAddress("0x0") // TODO
	TestnetCreate3Deployer = common.HexToAddress("0x0") // TODO

	// address that can deploy protocol contracts via Create3 factory.
	DevnetDeployer  = anvil.Account1
	MainnetDeployer = common.HexToAddress("0x0") // TODO
	TestnetDeployer = common.HexToAddress("0x0") // TODO
)
