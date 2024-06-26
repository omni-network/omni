package predeploys

import (
	contracts "github.com/omni-network/omni/contracts/allocs"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/core/types"
)

const (
	// ProxyAdmin for all namespaces.
	ProxyAdmin = "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	// Omni Predeploys.
	PortalRegistry   = "0x121E240000000000000000000000000000000001"
	OmniBridgeNative = "0x121E240000000000000000000000000000000002"
	WOmni            = "0x121E240000000000000000000000000000000003"

	// Octane Predeploys.
	Staking  = "0xcccccc0000000000000000000000000000000001"
	Slashing = "0xcccccc0000000000000000000000000000000002"
)

// Alloc returns the genesis allocs for the predeployed contracts.
func Alloc(network netconf.ID) (types.GenesisAlloc, error) {
	return contracts.Alloc(network)
}
