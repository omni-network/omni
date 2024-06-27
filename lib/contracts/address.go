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
	// This mainnet AVS was deployed outside of the e2e deployment flow, without Create3.
	return common.HexToAddress("0xed2f4d90b073128ae6769a9A8D51547B1Df766C8")
}

func OmegaAVS() common.Address {
	// This address was deployed outside of the e2e deployment flow, without Create3.
	return common.HexToAddress("0xa7b2e7830C51728832D33421670DbBE30299fD92")
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

func OmegaCreate3Factory() common.Address {
	return crypto.CreateAddress(eoa.MustAddress(netconf.Omega, eoa.RoleCreate3Deployer), 0)
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

func OmegaPortal() common.Address {
	return create3.Address(OmegaCreate3Factory(), PortalSalt(netconf.Omega), eoa.MustAddress(netconf.Omega, eoa.RoleDeployer))
}

func StagingPortal() common.Address {
	return create3.Address(StagingCreate3Factory(), PortalSalt(netconf.Staging), eoa.MustAddress(netconf.Staging, eoa.RoleDeployer))
}

func DevnetPortal() common.Address {
	return create3.Address(DevnetCreate3Factory(), PortalSalt(netconf.Devnet), eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer))
}

func Portal(network netconf.ID) (common.Address, bool) {
	switch network {
	case netconf.Mainnet:
		return MainnetPortal(), true
	case netconf.Omega:
		return OmegaPortal(), true
	case netconf.Staging:
		return StagingPortal(), true
	case netconf.Devnet:
		return DevnetPortal(), true
	default:
		return common.Address{}, false
	}
}

//
// ProxyAdmin.
//

func MainnetProxyAdmin() common.Address {
	return create3.Address(MainnetCreate3Factory(), ProxyAdminSalt(netconf.Mainnet), eoa.MustAddress(netconf.Mainnet, eoa.RoleDeployer))
}

func OmegaProxyAdmin() common.Address {
	return create3.Address(OmegaCreate3Factory(), ProxyAdminSalt(netconf.Omega), eoa.MustAddress(netconf.Omega, eoa.RoleDeployer))
}

func StagingProxyAdmin() common.Address {
	return create3.Address(StagingCreate3Factory(), ProxyAdminSalt(netconf.Staging), eoa.MustAddress(netconf.Staging, eoa.RoleDeployer))
}

func DevnetProxyAdmin() common.Address {
	return create3.Address(DevnetCreate3Factory(), ProxyAdminSalt(netconf.Devnet), eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer))
}

//
// L1Bridge.
//
// We use create3 deployments so we can have predictable addresses in ephemeral networks.

func StagingL1Bridge() common.Address {
	return create3.Address(StagingCreate3Factory(), L1BridgeSalt(netconf.Staging), eoa.MustAddress(netconf.Staging, eoa.RoleDeployer))
}

func DevnetL1Bridge() common.Address {
	return create3.Address(DevnetCreate3Factory(), L1BridgeSalt(netconf.Devnet), eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer))
}

func L1Bridge(network netconf.ID) (common.Address, bool) {
	switch network {
	case netconf.Staging:
		return StagingL1Bridge(), true
	case netconf.Devnet:
		return DevnetL1Bridge(), true
	default:
		return common.Address{}, false
	}
}

//
// Token.
//
// We use create3 deployments so we can have predictable addresses in ephemeral networks.

func StagingToken() common.Address {
	return create3.Address(StagingCreate3Factory(), TokenSalt(netconf.Staging), eoa.MustAddress(netconf.Staging, eoa.RoleDeployer))
}

func DevnetToken() common.Address {
	return create3.Address(DevnetCreate3Factory(), TokenSalt(netconf.Devnet), eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer))
}

func Token(network netconf.ID) (common.Address, bool) {
	switch network {
	case netconf.Staging:
		return StagingToken(), true
	case netconf.Devnet:
		return DevnetToken(), true
	default:
		return common.Address{}, false
	}
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

func L1BridgeSalt(network netconf.ID) string {
	return salt(network, "l1-bridge")
}

func TokenSalt(network netconf.ID) string {
	return salt(network, "token")
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
