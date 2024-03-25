package contracts

import (
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/buildinfo"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	// fireblocksDev is the address of the dev fireblocks account.
	// Currently used for all staging deployments / operations.
	fireblocksDev = "0x7a6cF389082dc698285474976d7C75CAdE08ab7e"
)

//nolint:gochecknoglobals // Static addresses
var (
	// ProxyAdminOwner.
	mainnetProxyAdminOwner = addr("0x0")
	testnetProxyAdminOwner = addr("0x0")
	stagingProxyAdminOwner = addr(fireblocksDev)
	devnetProxyAdminOwner  = anvil.DevAccount2()

	// Create3 Deployer - addrress that can deploy the create3 factory.
	mainnetCreate3Deployer = addr("0x0")
	testnetCreate3Deployer = addr("0x0")
	stagingCreate3Deployer = addr(fireblocksDev)
	devnetCreate3Deployer  = anvil.DevAccount0()

	// Deployer - address that can deploy protocol contracts via Create3 factory.
	mainnetDeployer = addr("0x0")
	testnetDeployer = addr("0x0")
	stagingDeployer = addr("0x0")
	devnetDeployer  = anvil.DevAccount1()

	// Portal Admin.
	mainnetPortalAdmin = addr("0x0")
	testnetPortalAdmin = addr("0x0")
	stagingPortalAdmin = addr(fireblocksDev)
	devnetPortalAdmin  = anvil.DevAccount2()

	// AVS Admin.
	mainnetAVSAdmin = addr("0x0")
	testnetAVSAdmin = addr("0x0")
	stagingAVSAdmin = addr(fireblocksDev)
	devnetAVSAdmin  = anvil.DevAccount2()
)

//
// ProxyAdminOwner.
//

func MainnetProxyAdminOwner() common.Address {
	return mainnetProxyAdminOwner
}

func TestnetProxyAdminOwner() common.Address {
	return testnetProxyAdminOwner
}

func StagingProxyAdminOwner() common.Address {
	return stagingProxyAdminOwner
}

func DevnetProxyAdminOwner() common.Address {
	return devnetProxyAdminOwner
}

//
// Create3Deployer.
//

func MainnetCreate3Deployer() common.Address {
	return mainnetCreate3Deployer
}

func TestnetCreate3Deployer() common.Address {
	return testnetCreate3Deployer
}

func StagingCreate3Deployer() common.Address {
	return stagingCreate3Deployer
}

func DevnetCreate3Deployer() common.Address {
	return devnetCreate3Deployer
}

//
// Deployer.
//

func MainnetDeployer() common.Address {
	return mainnetDeployer
}

func TestnetDeployer() common.Address {
	return testnetDeployer
}

func StagingDeployer() common.Address {
	return stagingDeployer
}

func DevnetDeployer() common.Address {
	return devnetDeployer
}

//
// PortalAdmin.
//

func MainnetPortalAdmin() common.Address {
	return mainnetPortalAdmin
}

func TestnetPortalAdmin() common.Address {
	return testnetPortalAdmin
}

func StagingPortalAdmin() common.Address {
	return stagingPortalAdmin
}

func DevnetPortalAdmin() common.Address {
	return devnetPortalAdmin
}

//
// AVSAdmin.
//

func MainnetAVSAdmin() common.Address {
	return mainnetAVSAdmin
}

func TestnetAVSAdmin() common.Address {
	return testnetAVSAdmin
}

func StagingAVSAdmin() common.Address {
	return stagingAVSAdmin
}

func DevnetAVSAdmin() common.Address {
	return devnetAVSAdmin
}

//
// Create3Factory.
//

func MainnetCreate3Factory() common.Address {
	return crypto.CreateAddress(mainnetCreate3Deployer, 0)
}

func TestnetCreate3Factory() common.Address {
	return crypto.CreateAddress(testnetCreate3Deployer, 0)
}

func StagingCreate3Factory() common.Address {
	return crypto.CreateAddress(stagingCreate3Deployer, 0)
}

func DevnetCreate3Factory() common.Address {
	return crypto.CreateAddress(devnetCreate3Deployer, 0)
}

//
// ProxyAdmin.
//

func MainnetProxyAdmin() common.Address {
	return create3.Address(MainnetCreate3Factory(), ProxyAdminSalt(netconf.Mainnet), mainnetDeployer)
}

func TestnetProxyAdmin() common.Address {
	return create3.Address(TestnetCreate3Factory(), ProxyAdminSalt(netconf.Testnet), testnetDeployer)
}

func StagingProxyAdmin() common.Address {
	return create3.Address(StagingCreate3Factory(), ProxyAdminSalt(netconf.Staging), stagingDeployer)
}

func DevnetProxyAdmin() common.Address {
	return create3.Address(DevnetCreate3Factory(), ProxyAdminSalt(netconf.Devnet), devnetDeployer)
}

//
// Portal.
//

func MainnetPortal() common.Address {
	return create3.Address(MainnetCreate3Factory(), PortalSalt(netconf.Mainnet), mainnetDeployer)
}

func TestnetPortal() common.Address {
	return create3.Address(TestnetCreate3Factory(), PortalSalt(netconf.Testnet), testnetDeployer)
}

func StagingPortal() common.Address {
	return create3.Address(StagingCreate3Factory(), PortalSalt(netconf.Staging), stagingDeployer)
}

func DevnetPortal() common.Address {
	return create3.Address(DevnetCreate3Factory(), PortalSalt(netconf.Devnet), devnetDeployer)
}

//
// AVS.
//

func MainnetAVS() common.Address {
	return create3.Address(MainnetCreate3Factory(), AVSSalt(netconf.Mainnet), mainnetDeployer)
}

func TestnetAVS() common.Address {
	return create3.Address(TestnetCreate3Factory(), AVSSalt(netconf.Testnet), testnetDeployer)
}

func StagingAVS() common.Address {
	return create3.Address(StagingCreate3Factory(), AVSSalt(netconf.Staging), stagingDeployer)
}

func DevnetAVS() common.Address {
	return create3.Address(DevnetCreate3Factory(), AVSSalt(netconf.Devnet), devnetDeployer)
}

//
// Salts.
//

func ProxyAdminSalt(network string) string {
	return salt(network, "proxy-admin")
}

func PortalSalt(network string) string {
	return salt(network, "portal")
}

func AVSSalt(network string) string {
	return salt(network, "avs")
}

//
// Utils.
//

// salt generates a salt for a contract deployment, adding git build info.
func salt(network string, contract string) string {
	return network + "-" + contract + "-" + buildinfo.Version()
}

func addr(hex string) common.Address {
	return common.HexToAddress(hex)
}
