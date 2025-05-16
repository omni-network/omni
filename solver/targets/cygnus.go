package targets

import (
	"github.com/omni-network/omni/lib/evmchain"

	"github.com/ethereum/go-ethereum/common"
)

const NameCygnus = "Cygnus"

var (
	BaseCygnusGlobalUSD = addr("0xCa72827a3D211CfD8F6b00Ac98824872b72CAb49")

	cygnus = Target{
		Name: NameCygnus,
		Addresses: func(chainID uint64) map[common.Address]bool {
			if chainID == evmchain.IDBase {
				return set(BaseCygnusGlobalUSD)
			}

			return map[common.Address]bool{}
		},
	}
)
