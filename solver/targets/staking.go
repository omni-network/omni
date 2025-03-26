package targets

import (
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"

	"github.com/ethereum/go-ethereum/common"
)

const NameOmniStaking = "OmniStaking"

var staking = Target{
	Name: NameOmniStaking,
	Addresses: func(uint64) map[common.Address]bool {
		return map[common.Address]bool{
			common.HexToAddress(predeploys.Staking): true,
		}
	},
}
