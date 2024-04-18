package contracts

import (
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

//
// AVS.
//

func MainnetAVS() common.Address {
	return create3.Address(MainnetCreate3Factory(), AVSSalt(netconf.Mainnet), eoa.MainnetDeployer())
}

func TestnetAVS() common.Address {
	return create3.Address(TestnetCreate3Factory(), AVSSalt(netconf.Testnet), eoa.TestnetDeployer())
}

func StagingAVS() common.Address {
	return create3.Address(StagingCreate3Factory(), AVSSalt(netconf.Staging), eoa.StagingDeployer())
}

func DevnetAVS() common.Address {
	return create3.Address(DevnetCreate3Factory(), AVSSalt(netconf.Devnet), eoa.DevnetDeployer())
}

//
// Create3Factory.
//

func MainnetCreate3Factory() common.Address {
	return crypto.CreateAddress(eoa.MainnetCreate3Deployer(), 0)
}

func TestnetCreate3Factory() common.Address {
	return crypto.CreateAddress(eoa.TestnetCreate3Deployer(), 0)
}

func StagingCreate3Factory() common.Address {
	return crypto.CreateAddress(eoa.StagingCreate3Deployer(), 0)
}

func DevnetCreate3Factory() common.Address {
	return crypto.CreateAddress(eoa.DevnetCreate3Deployer(), 0)
}

//
// Portal.
//

func MainnetPortal() common.Address {
	return create3.Address(MainnetCreate3Factory(), PortalSalt(netconf.Mainnet), eoa.MainnetDeployer())
}

func TestnetPortal() common.Address {
	return create3.Address(TestnetCreate3Factory(), PortalSalt(netconf.Testnet), eoa.TestnetDeployer())
}

func StagingPortal() common.Address {
	return create3.Address(StagingCreate3Factory(), PortalSalt(netconf.Staging), eoa.StagingDeployer())
}

func DevnetPortal() common.Address {
	return create3.Address(DevnetCreate3Factory(), PortalSalt(netconf.Devnet), eoa.DevnetDeployer())
}

//
// ProxyAdmin.
//

func MainnetProxyAdmin() common.Address {
	return create3.Address(MainnetCreate3Factory(), ProxyAdminSalt(netconf.Mainnet), eoa.MainnetDeployer())
}

func TestnetProxyAdmin() common.Address {
	return create3.Address(TestnetCreate3Factory(), ProxyAdminSalt(netconf.Testnet), eoa.TestnetDeployer())
}

func StagingProxyAdmin() common.Address {
	return create3.Address(StagingCreate3Factory(), ProxyAdminSalt(netconf.Staging), eoa.StagingDeployer())
}

func DevnetProxyAdmin() common.Address {
	return create3.Address(DevnetCreate3Factory(), ProxyAdminSalt(netconf.Devnet), eoa.DevnetDeployer())
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
