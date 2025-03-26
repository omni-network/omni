package targets

import (
	"github.com/omni-network/omni/lib/evmchain"

	"github.com/ethereum/go-ethereum/common"
)

const NameEigen = "Eigen"

var (
	// Eigen testnet.
	EigenHoleskyStrategyManager = addr("0xdfB5f6CE42aAA7830E94ECFCcAd411beF4d4D5b6")

	// Eigen mainnet.
	EigenMainnetStrategyManager = addr("0x858646372CC42E1A627fcE94aa7A7033e7CF075A")

	eigen = Target{
		Name: NameEigen,
		Addresses: networkChainAddrs(map[uint64]map[common.Address]bool{
			evmchain.IDHolesky:  set(EigenHoleskyStrategyManager),
			evmchain.IDEthereum: set(EigenMainnetStrategyManager),
		}),
	}
)
