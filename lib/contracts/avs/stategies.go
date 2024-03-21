package avs

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type StrategyParam struct {
	Strategy   common.Address
	Multiplier *big.Int
}

func holeskeyStrategyParams() []StrategyParam {
	return []StrategyParam{} // TODO: setup holesky strategy params
}

func devnetStrategyParams() []StrategyParam {
	// devnet weth strategy
	wethStrat := common.HexToAddress("0xdBD296711eC8eF9Aacb623ee3F1C0922dce0D7b2")

	return []StrategyParam{
		{
			Strategy:   wethStrat,
			Multiplier: big.NewInt(1e18), // OmniAVS.STRATEGY_WEIGHTING_DIVISOR
		},
	}
}
