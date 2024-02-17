package types

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

type EigenLayerDeployments struct {
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

func LoadEigenLayerDeployments(file string) (EigenLayerDeployments, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return EigenLayerDeployments{}, errors.Wrap(err, "read eigen layer deployments", "path", file)
	}

	var deployments EigenLayerDeployments
	if err := json.Unmarshal(data, &deployments); err != nil {
		return EigenLayerDeployments{}, errors.Wrap(err, "unmarshal eigen layer deployments")
	}

	return deployments, nil
}

func DefaultTestAVSConfig(eigen EigenLayerDeployments) AVSConfig {
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
