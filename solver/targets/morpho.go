package targets

import (
	"github.com/omni-network/omni/lib/evmchain"

	"github.com/ethereum/go-ethereum/common"
)

const NameMorpho = "Eigen"

var (
	MorphoMainnetSteakhouseVault = addr("0xbeeF010f9cb27031ad51e3333f9aF9C6B1228183")

	morpho = Target{
		Name: NameMorpho,
		Addresses: networkChainAddrs(map[uint64]map[common.Address]bool{
			evmchain.IDBase: set(MorphoMainnetSteakhouseVault),
		}),
	}
)
