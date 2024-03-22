package contracts

import (
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

//nolint:gochecknoglobals // Static addresses
var (
	// Create3Factory.
	DevnetCreate3Factory  = addr("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	MainnetCreate3Factory = addr("0x0")
	TestnetCreate3Factory = addr("0x0")

	// ProxyAdmin.
	DevnetProxyAdmin  = Create3Address(DevnetCreate3Factory, ProxyAdminSalt(netconf.Devnet), DevnetDeployer)
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
	DevnetPortal  = Create3Address(DevnetCreate3Factory, PortalSalt(netconf.Devnet), DevnetDeployer)
	MainnetPortal = addr("0x0")
	TestnetPortal = addr("0x0")

	// Fee Oracle V1.
	DevnetFeeOracleV1  = addr("0x1234") // TODO: stubbed for now, so portal tests don't fail
	MainnetFeeOracleV1 = addr("0x0")
	TestnetFeeOracleV1 = addr("0x0")
)

func ProxyAdminSalt(network string) string {
	return salt(network, "proxy-admin")
}

func PortalSalt(network string) string {
	return salt(network, "portal")
}

// salt generates a salt for a contract deployment, adding git build info for staging.
func salt(network string, contract string) string {
	salt := network + "-" + contract

	if network == netconf.Staging {
		return salt + "-" + buildinfo.Version()
	}

	return salt
}

func addr(hex string) common.Address {
	return common.HexToAddress(hex)
}
