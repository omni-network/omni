package avs

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type StrategyParam struct {
	Strategy   common.Address
	Multiplier *big.Int
}

var (
	stdMultiplier = big.NewInt(1e18) // OmniAVS.STRATEGY_WEIGHTING_DIVISOR
)

func holeskyStrategyParams() []StrategyParam {
	return []StrategyParam{
		// sETH
		{
			Strategy:   common.HexToAddress("0x7D704507b76571a51d9caE8AdDAbBFd0ba0e63d3"),
			Multiplier: stdMultiplier,
		},
		// rETH
		{
			Strategy:   common.HexToAddress("0x3A8fBdf9e77DFc25d09741f51d3E181b25d0c4E0"),
			Multiplier: stdMultiplier,
		},
		// WETH
		{
			Strategy:   common.HexToAddress("0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9"),
			Multiplier: stdMultiplier,
		},
		// beacon eth
		{
			Strategy:   common.HexToAddress("0xbeaC0eeEeeeeEEeEeEEEEeeEEeEeeeEeeEEBEaC0"),
			Multiplier: stdMultiplier,
		},
	}
}

func devnetStrategyParams() []StrategyParam {
	return []StrategyParam{
		// devnet WETH
		{
			Strategy:   common.HexToAddress("0xdBD296711eC8eF9Aacb623ee3F1C0922dce0D7b2"),
			Multiplier: stdMultiplier,
		},
	}
}
