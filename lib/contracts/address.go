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
	return create3.Address(MainnetCreate3Factory(), AVSSalt(netconf.Mainnet), eoa.MustAddress(netconf.Mainnet, eoa.RoleDeployer))
}

func TestnetAVS() common.Address {
	return create3.Address(TestnetCreate3Factory(), AVSSalt(netconf.Testnet), eoa.MustAddress(netconf.Testnet, eoa.RoleDeployer))
}

func StagingAVS() common.Address {
	return create3.Address(StagingCreate3Factory(), AVSSalt(netconf.Staging), eoa.MustAddress(netconf.Staging, eoa.RoleDeployer))
}

func DevnetAVS() common.Address {
	return create3.Address(DevnetCreate3Factory(), AVSSalt(netconf.Devnet), eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer))
}

//
// Create3Factory.
//

func MainnetCreate3Factory() common.Address {
	return crypto.CreateAddress(eoa.MustAddress(netconf.Mainnet, eoa.RoleCreate3Deployer), 0)
}

func TestnetCreate3Factory() common.Address {
	return crypto.CreateAddress(eoa.MustAddress(netconf.Testnet, eoa.RoleCreate3Deployer), 0)
}

func StagingCreate3Factory() common.Address {
	return crypto.CreateAddress(eoa.MustAddress(netconf.Staging, eoa.RoleCreate3Deployer), 0)
}

func DevnetCreate3Factory() common.Address {
	return crypto.CreateAddress(eoa.MustAddress(netconf.Devnet, eoa.RoleCreate3Deployer), 0)
}

//
// Portal.
//

func MainnetPortal() common.Address {
	return create3.Address(MainnetCreate3Factory(), PortalSalt(netconf.Mainnet), eoa.MustAddress(netconf.Mainnet, eoa.RoleDeployer))
}

func TestnetPortal() common.Address {
	return create3.Address(TestnetCreate3Factory(), PortalSalt(netconf.Testnet), eoa.MustAddress(netconf.Testnet, eoa.RoleDeployer))
}

func StagingPortal() common.Address {
	return create3.Address(StagingCreate3Factory(), PortalSalt(netconf.Staging), eoa.MustAddress(netconf.Staging, eoa.RoleDeployer))
}

func DevnetPortal() common.Address {
	return create3.Address(DevnetCreate3Factory(), PortalSalt(netconf.Devnet), eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer))
}

//
// ProxyAdmin.
//

func MainnetProxyAdmin() common.Address {
	return create3.Address(MainnetCreate3Factory(), ProxyAdminSalt(netconf.Mainnet), eoa.MustAddress(netconf.Mainnet, eoa.RoleDeployer))
}

func TestnetProxyAdmin() common.Address {
	return create3.Address(TestnetCreate3Factory(), ProxyAdminSalt(netconf.Testnet), eoa.MustAddress(netconf.Testnet, eoa.RoleDeployer))
}

func StagingProxyAdmin() common.Address {
	return create3.Address(StagingCreate3Factory(), ProxyAdminSalt(netconf.Staging), eoa.MustAddress(netconf.Staging, eoa.RoleDeployer))
}

func DevnetProxyAdmin() common.Address {
	return create3.Address(DevnetCreate3Factory(), ProxyAdminSalt(netconf.Devnet), eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer))
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
