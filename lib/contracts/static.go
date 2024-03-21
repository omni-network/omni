package contracts

import (
	"github.com/ethereum/go-ethereum/common"
)

//nolint:gochecknoglobals // Static addresses
var (
	MainnetCreate3Factory = common.HexToAddress("0x0") // TODO
	TestnetCreate3Factory = common.HexToAddress("0x0") // TODO

	MainnetProxyAdmin = common.HexToAddress("0x0") // TODO
	TestnetProxyAdmin = common.HexToAddress("0x0") // TODO

	MainnetProxyAdminOwner = common.HexToAddress("0x0") // TODO
	TestnetProxyAdminOwner = common.HexToAddress("0x0") // TODO

	// addrress that can deploy the create3 factory.
	MainnetCreate3Deployer = common.HexToAddress("0x0") // TODO
	TestnetCreate3Deployer = common.HexToAddress("0x0") // TODO

	// address that can deploy protocol contracts via Create3 factory.
	MainnetDeployer = common.HexToAddress("0x0") // TODO
	TestnetDeployer = common.HexToAddress("0x0") // TODO
)
