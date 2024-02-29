package avs

import (
	"encoding/json"
	"math/big"
	"os"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
)

type StrategyParam struct {
	Strategy   common.Address `json:"strategy"`
	Multiplier *big.Int       `json:"multiplier"`
}

type AVSConfig struct {
	StrategyParams []StrategyParam `json:"strategyParams"`
}

type EigenDeployments struct {
	// core deployments
	AVSDirectory      common.Address `json:"AVSDirectory"`
	DelegationManager common.Address `json:"DelegationManager"`
	StrategyManager   common.Address `json:"StrategyManager"`
	EigenPodManager   common.Address `json:"EigenPodManager"`

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
	strategyParams := make([]StrategyParam, 0, len(eigen.Strategies))
	for _, strategy := range eigen.Strategies {
		strategyParams = append(strategyParams, StrategyParam{
			Strategy:   strategy,
			Multiplier: big.NewInt(1e18), // OmniAVS.STRATEGY_WEIGHTING_DIVISOR
		})
	}

	return AVSConfig{
		StrategyParams: strategyParams,
	}
}
