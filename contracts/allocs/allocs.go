package allocs

import (
	"encoding/json"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/core/types"

	_ "embed"
)

var (
	//go:embed devnet.json
	devnetJSON []byte

	//go:embed staging.json
	stagingJSON []byte

	//go:embed omega.json
	omegaJSON []byte

	//go:embed mainnet.json
	mainnetJSON []byte

	devnetAlloc  = mustUnmarshalAlloc(devnetJSON)
	stagingAlloc = mustUnmarshalAlloc(stagingJSON)
	omegaAlloc   = mustUnmarshalAlloc(omegaJSON)
	mainnetAlloc = mustUnmarshalAlloc(mainnetJSON)
)

// locked maps network ID to whether the network's allocations are locked.
// Once a persistent network has been deployed, its allocs should be locked,
// so that allocs here are consistent with the live network's genesis.
var locked = map[netconf.ID]bool{
	netconf.Omega:   true,
	netconf.Mainnet: true,
}

func IsLocked(network netconf.ID) bool {
	return locked[network]
}

func Alloc(network netconf.ID) (types.GenesisAlloc, error) {
	switch network {
	case netconf.Devnet:
		return devnetAlloc, nil
	case netconf.Staging:
		return stagingAlloc, nil
	case netconf.Omega:
		return omegaAlloc, nil
	case netconf.Mainnet:
		return mainnetAlloc, nil
	default:
		return nil, errors.New("unknown network")
	}
}

func MustAlloc(network netconf.ID) types.GenesisAlloc {
	alloc, err := Alloc(network)
	if err != nil {
		panic(err)
	}

	return alloc
}

func mustUnmarshalAlloc(data []byte) types.GenesisAlloc {
	var alloc types.GenesisAlloc
	if err := json.Unmarshal(data, &alloc); err != nil {
		panic(err)
	}

	return alloc
}
