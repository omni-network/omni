package rlusd

import (
	"github.com/omni-network/omni/e2e/xbridge/types"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

var (
	xtoken = types.TokenDescriptors{
		Symbol: "RLUSD.e",
		Name:   "Bridged RLUSD (Omni)",
	}

	wraps = types.TokenDescriptors{
		Symbol: "RLUSD",
		Name:   "RLUSD",
	}

	canonicals = map[netconf.ID]types.TokenDeployment{
		netconf.Mainnet: {
			Name:    wraps.Name,
			Symbol:  wraps.Symbol,
			Address: common.HexToAddress("0x8292bb45bf1ee4d140127049757c2e0ff06317ed"),
			ChainID: evmchain.IDEthereum,
		},
	}
)
