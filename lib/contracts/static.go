package contracts

import (
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//nolint:gochecknoglobals // Static addresses
var (
	// Create3Factory.
	DevnetCreate3Factory  = crypto.CreateAddress(DevnetCreate3Deployer, 0)
	MainnetCreate3Factory = crypto.CreateAddress(MainnetCreate3Deployer, 0)
	TestnetCreate3Factory = crypto.CreateAddress(TestnetCreate3Deployer, 0)

	// ProxyAdmin.
	DevnetProxyAdmin  = create3.Address(DevnetCreate3Factory, ProxyAdminSalt(netconf.Devnet), DevnetDeployer)
	MainnetProxyAdmin = create3.Address(MainnetCreate3Factory, ProxyAdminSalt(netconf.Mainnet), MainnetDeployer)
	TestnetProxyAdmin = create3.Address(TestnetCreate3Factory, ProxyAdminSalt(netconf.Testnet), TestnetDeployer)

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
	DevnetAVSAdmin  = anvil.Account2
	MainnetAVSAdmin = addr("0x0")
	TestnetAVSAdmin = addr("0x0")

	// Omni Portal.
	DevnetPortal  = create3.Address(DevnetCreate3Factory, PortalSalt(netconf.Devnet), DevnetDeployer)
	MainnetPortal = create3.Address(MainnetCreate3Factory, PortalSalt(netconf.Mainnet), MainnetDeployer)
	TestnetPortal = create3.Address(TestnetCreate3Factory, PortalSalt(netconf.Testnet), TestnetDeployer)

	// Fee Oracle V1.
	DevnetFeeOracleV1  = addr("0x1234") // TODO: stubbed for now, so portal tests don't fail
	MainnetFeeOracleV1 = addr("0x0")
	TestnetFeeOracleV1 = addr("0x0")

	// AVS.
	DevnetAVS  = create3.Address(DevnetCreate3Factory, AVSSalt(netconf.Devnet), DevnetDeployer)
	MainnetAVS = create3.Address(MainnetCreate3Factory, AVSSalt(netconf.Mainnet), MainnetDeployer)
	TestnetAVS = create3.Address(TestnetCreate3Factory, AVSSalt(netconf.Testnet), TestnetDeployer)
)

func ProxyAdminSalt(network string) string {
	return salt(network, "proxy-admin")
}

func PortalSalt(network string) string {
	return salt(network, "portal")
}

func AVSSalt(network string) string {
	return salt(network, "avs")
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
