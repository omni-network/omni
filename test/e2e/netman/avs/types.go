package avs

import (
	"encoding/json"
	"math/big"
	"os"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
)

type StrategyParams struct {
	Strategy   common.Address `json:"strategy"`
	Multiplier *big.Int       `json:"multiplier"`
}

type AVSConfig struct {
	MinimumOperatorStake *big.Int         `json:"minimumOperatorStake"`
	MaximumOperatorCount uint32           `json:"maximumOperatorCount"`
	StrategyParams       []StrategyParams `json:"strategyParams"`
}

type EigenDeployments struct {
	// core deployments
	AVSDirectory      common.Address `json:"AVSDirectory"`
	DelegationManager common.Address `json:"DelegationManager"`
	StrategyManager   common.Address `json:"StrategyManager"`
	EigenPodManager   common.Address `json:"EigenPodManager"`

	// test token & strategies
	EigenStrategy common.Address `json:"EigenStrategy"`
	EigenToken    common.Address `json:"EigenToken"`
	WETH          common.Address `json:"WETH"`
	WETHStrategy  common.Address `json:"WETHStrategy"`
}

func LoadDeployments(file string) (EigenDeployments, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return EigenDeployments{}, errors.Wrap(err, "read eigen layer resp", "path", file)
	}

	var resp EigenDeployments
	if err := json.Unmarshal(data, &resp); err != nil {
		return EigenDeployments{}, errors.Wrap(err, "unmarshal eigen layer resp")
	}

	return resp, nil
}

func DefaultTestAVSConfig(eigen EigenDeployments) AVSConfig {
	return AVSConfig{
		MinimumOperatorStake: big.NewInt(1e18),
		MaximumOperatorCount: 10,
		StrategyParams: []StrategyParams{
			{
				Strategy:   eigen.EigenStrategy,
				Multiplier: big.NewInt(1e18), // OmniAVS.WEIGHTING_DIVISOR
			},
			{
				Strategy:   eigen.WETHStrategy,
				Multiplier: big.NewInt(1e18), // OmniAVS.WEIGHTING_DIVISOR
			},
		},
	}
}
