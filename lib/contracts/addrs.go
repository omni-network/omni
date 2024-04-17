package contracts

import (
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Admin returns the address of the admin for the given network.
func Admin(network netconf.ID) (common.Address, error) {
	switch network {
	case netconf.Mainnet:
		return eoa.MainnetAdmin, nil
	case netconf.Testnet:
		return eoa.TestnetAdmin, nil
	case netconf.Staging:
		return eoa.StagingAdmin, nil
	case netconf.Devnet:
		return eoa.DevnetAdmin, nil
	default:
		return common.Address{}, errors.New("unknown network", "network", network)
	}
}

//
// ProxyAdminOwner.
//

func MainnetProxyAdminOwner() common.Address {
	return eoa.MainnetAdmin
}

func TestnetProxyAdminOwner() common.Address {
	return eoa.TestnetAdmin
}

func StagingProxyAdminOwner() common.Address {
	return eoa.StagingAdmin
}

func DevnetProxyAdminOwner() common.Address {
	return eoa.DevnetAdmin
}

//
// Create3Deployer.
//

func MainnetCreate3Deployer() common.Address {
	return eoa.MainnetCreate3Deployer
}

func TestnetCreate3Deployer() common.Address {
	return eoa.TestnetCreate3Deployer
}

func StagingCreate3Deployer() common.Address {
	return eoa.StagingCreate3Deployer
}

func DevnetCreate3Deployer() common.Address {
	return eoa.DevnetCreate3Deployer
}

//
// Deployer.
//

func MainnetDeployer() common.Address {
	return eoa.MainnetDeployer
}

func TestnetDeployer() common.Address {
	return eoa.TestnetDeployer
}

func StagingDeployer() common.Address {
	return eoa.StagingDeployer
}

func DevnetDeployer() common.Address {
	return eoa.DevnetDeployer
}

//
// PortalAdmin.
//

func MainnetPortalAdmin() common.Address {
	return eoa.MainnetAdmin
}

func TestnetPortalAdmin() common.Address {
	return eoa.TestnetAdmin
}

func StagingPortalAdmin() common.Address {
	return eoa.StagingAdmin
}

func DevnetPortalAdmin() common.Address {
	return eoa.DevnetAdmin
}

//
// AVSAdmin.
//

func MainnetAVSAdmin() common.Address {
	return eoa.MainnetAdmin
}

func TestnetAVSAdmin() common.Address {
	return eoa.TestnetAdmin
}

func StagingAVSAdmin() common.Address {
	return eoa.StagingAdmin
}

func DevnetAVSAdmin() common.Address {
	return eoa.DevnetAdmin
}

//
// Create3Factory.
//

func MainnetCreate3Factory() common.Address {
	return crypto.CreateAddress(eoa.MainnetCreate3Deployer, 0)
}

func TestnetCreate3Factory() common.Address {
	return crypto.CreateAddress(eoa.TestnetCreate3Deployer, 0)
}

func StagingCreate3Factory() common.Address {
	return crypto.CreateAddress(eoa.StagingCreate3Deployer, 0)
}

func DevnetCreate3Factory() common.Address {
	return crypto.CreateAddress(eoa.DevnetCreate3Deployer, 0)
}

//
// ProxyAdmin.
//

func MainnetProxyAdmin() common.Address {
	return create3.Address(MainnetCreate3Factory(), ProxyAdminSalt(netconf.Mainnet), eoa.MainnetDeployer)
}

func TestnetProxyAdmin() common.Address {
	return create3.Address(TestnetCreate3Factory(), ProxyAdminSalt(netconf.Testnet), eoa.TestnetDeployer)
}

func StagingProxyAdmin() common.Address {
	return create3.Address(StagingCreate3Factory(), ProxyAdminSalt(netconf.Staging), eoa.StagingDeployer)
}

func DevnetProxyAdmin() common.Address {
	return create3.Address(DevnetCreate3Factory(), ProxyAdminSalt(netconf.Devnet), eoa.DevnetDeployer)
}

//
// Portal.
//

func MainnetPortal() common.Address {
	return create3.Address(MainnetCreate3Factory(), PortalSalt(netconf.Mainnet), eoa.MainnetDeployer)
}

func TestnetPortal() common.Address {
	return create3.Address(TestnetCreate3Factory(), PortalSalt(netconf.Testnet), eoa.TestnetDeployer)
}

func StagingPortal() common.Address {
	return create3.Address(StagingCreate3Factory(), PortalSalt(netconf.Staging), eoa.StagingDeployer)
}

func DevnetPortal() common.Address {
	return create3.Address(DevnetCreate3Factory(), PortalSalt(netconf.Devnet), eoa.DevnetDeployer)
}

//
// AVS.
//

func MainnetAVS() common.Address {
	return create3.Address(MainnetCreate3Factory(), AVSSalt(netconf.Mainnet), eoa.MainnetDeployer)
}

func TestnetAVS() common.Address {
	return create3.Address(TestnetCreate3Factory(), AVSSalt(netconf.Testnet), eoa.TestnetDeployer)
}

func StagingAVS() common.Address {
	return create3.Address(StagingCreate3Factory(), AVSSalt(netconf.Staging), eoa.StagingDeployer)
}

func DevnetAVS() common.Address {
	return create3.Address(DevnetCreate3Factory(), AVSSalt(netconf.Devnet), eoa.DevnetDeployer)
}

//
// Salts.
//

func ProxyAdminSalt(network netconf.ID) string {
	return salt(network, "proxy-admin")
}

func PortalSalt(network netconf.ID) string {
	// only portal salts are versioned
	return salt(network, "portal-"+network.Version())
}

func AVSSalt(network netconf.ID) string {
	return salt(network, "avs")
}

//
// Utils.
//

// salt generates a salt for a contract deployment. For ephemeral networks,
// the salt includes a random per-run suffix. For persistent networks, the
// sale is static.
func salt(network netconf.ID, contract string) string {
	return string(network) + "-" + contract
}

func addr(hex string) common.Address {
	return common.HexToAddress(hex)
}
