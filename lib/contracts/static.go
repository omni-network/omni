package contracts

import (
	"github.com/omni-network/omni/lib/anvil"

	"github.com/ethereum/go-ethereum/common"
)

//nolint:gochecknoglobals // Static addresses
var (
	// Create3Factory.
	DevnetCreate3Factory  = addr("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	MainnetCreate3Factory = addr("0x0")
	TestnetCreate3Factory = addr("0x0")

	// ProxyAdmin.
	DevnetProxyAdmin  = addr("0x733AA9e7E4025E9F69DBEd9e05155e081D720565")
	MainnetProxyAdmin = addr("0x0")
	TestnetProxyAdmin = addr("0x0")

	// ProxyAdminOwner.
	DevnetProxyAdminOwner  = anvil.Account2
	MainnetProxyAdminOwner = addr("0x0")
	TestnetProxyAdminOwner = addr("0x0")

	// Create3 Deployer - addrress that can deploy the create3 factory.
	DevnetCreate3Deployer  = anvil.Account0
	MainnetCreate3Deployer = addr("0x0")
	TestnetCreate3Deployer = addr("0x0")

	// Deployer - address that can deploy protocol contracts via Create3 factory.
	DevnetDeployer  = anvil.Account1
	MainnetDeployer = addr("0x0")
	TestnetDeployer = addr("0x0")

	// Portal Admin.
	DevnetPortalAdmin  = anvil.Account2
	MainnetPortalAdmin = addr("0x0")
	TestnetPortalAdmin = addr("0x0")

	// AVS Admin.
	DevnetAVSAAdmin  = anvil.Account2
	MainnetAVSAAdmin = addr("0x0")
	TestnetAVSAAdmin = addr("0x0")

	// Omni Portal.
	DevnetPortal  = addr("0x1Fa76B04A827b7BBF34646815358E2ADE0dFCB77")
	MainnetPortal = addr("0x0")
	TestnetPortal = addr("0x0")

	// Fee Oracle V1.
	DevnetFeeOracleV1  = addr("0x1234") // TODO: stubbed for now, so portal tests don't fail
	MainnetFeeOracleV1 = addr("0x0")
	TestnetFeeOracleV1 = addr("0x0")
)

func addr(hex string) common.Address {
	return common.HexToAddress(hex)
}
