package contracts

import (
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/create3"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// AVS returns the AVS contract address for the given network.
func AVS(network netconf.ID) common.Address {
	if network == netconf.Mainnet {
		return common.HexToAddress("0xed2f4d90b073128ae6769a9A8D51547B1Df766C8")
	} else if network == netconf.Omega {
		return common.HexToAddress("0xa7b2e7830C51728832D33421670DbBE30299fD92")
	}

	return create3.Address(Create3Factory(network), AVSSalt(network), eoa.MustAddress(network, eoa.RoleDeployer))
}

// Create3Factory returns the Create3 factory address for the given network.
func Create3Factory(network netconf.ID) common.Address {
	return crypto.CreateAddress(eoa.MustAddress(network, eoa.RoleCreate3Deployer), 0)
}

// Portal returns the Portal contract address for the given network.
func Portal(network netconf.ID) common.Address {
	return create3.Address(Create3Factory(network), PortalSalt(network), eoa.MustAddress(network, eoa.RoleDeployer))
}

// L1Bridge returns the L1Bridge contract address for the given network.
//
// We use create3 deployments so we can have predictable addresses in ephemeral networks.
func L1Bridge(network netconf.ID) common.Address {
	return create3.Address(Create3Factory(network), L1BridgeSalt(network), eoa.MustAddress(network, eoa.RoleDeployer))
}

// Token returns the Token contract address for the given network.
func Token(network netconf.ID) common.Address {
	if network == netconf.Mainnet {
		return common.HexToAddress("0x36e66fbbce51e4cd5bd3c62b637eb411b18949d4")
	}

	return create3.Address(Create3Factory(network), TokenSalt(network), eoa.MustAddress(network, eoa.RoleDeployer))
}

//
// Salts.
//

func PortalSalt(network netconf.ID) string {
	return salt(network, "portal-"+network.Version())
}

func AVSSalt(network netconf.ID) string {
	// AVS not versioned, as requiring re-registration per each version is too cumbersome.
	return salt(network, "avs")
}

func L1BridgeSalt(network netconf.ID) string {
	return salt(network, "l1-bridge-"+network.Version())
}

func TokenSalt(network netconf.ID) string {
	return salt(network, "token-"+network.Version())
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
