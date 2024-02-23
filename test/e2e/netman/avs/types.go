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
	AVSDirectory      common.Address `json:"avsDirectory"`
	DelegationManager common.Address `json:"delegationManager"`
	StrategyManager   common.Address `json:"strategyManager"`
	EigenPodManager   common.Address `json:"eigenPodManager"`

	// maps token symbol to strategy address
	Strategies map[string]common.Address `json:"strategies"`
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
	strategyParams := make([]StrategyParams, 0, len(eigen.Strategies))
	for _, strategy := range eigen.Strategies {
		strategyParams = append(strategyParams, StrategyParams{
			Strategy:   strategy,
			Multiplier: big.NewInt(1e18), // OmniAVS.WEIGHTING_DIVISOR
		})
	}

	return AVSConfig{
		MinimumOperatorStake: big.NewInt(1e18),
		MaximumOperatorCount: 10,
		StrategyParams:       strategyParams,
	}
}
