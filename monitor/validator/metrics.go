package validator

import (
	"github.com/omni-network/omni/lib/promutil"

	"github.com/cometbft/cometbft/crypto"

	"github.com/ethereum/go-ethereum/common"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	powerGauge = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "validator",
		Name:      "power",
		Help:      "Current potential power by validator (may not be bonded)",
	}, []string{"validator", "validator_address", "operator"}) // Main metric with additional label identifiers

	jailedGauge = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "validator",
		Name:      "jailed",
		Help:      "Constant gauge set to 1 if the validator is jailed otherwise 0",
	}, []string{"validator"})

	bondedGauge = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "validator",
		Name:      "bonded",
		Help:      "Constant gauge set to 1 if the validator is bonded otherwise 0",
	}, []string{"validator"})

	tombstonedGauge = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "validator",
		Name:      "tombstoned",
		Help:      "Constant gauge set to 1 if the validator is tombstoned otherwise 0",
	}, []string{"validator"})

	uptimeGauge = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "validator",
		Name:      "uptime",
		Help:      "Percentage of blocks signed in past <slashing_signing_window>",
	}, []string{"validator"})

	rewardsGauge = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "validator",
		Name:      "rewards",
		Help:      "Validator rewards",
	}, []string{"validator"})
)

type sample struct {
	ConsensusEthAddr common.Address
	ConsensusCmtAddr crypto.Address
	OperatorEthAddr  common.Address
	Power            int64
	Jailed           bool
	Bonded           bool
	Tombstoned       bool
	Uptime           float64
	Rewards          float64
}

func sampleValidator(s sample) {
	powerGauge.WithLabelValues(s.ConsensusEthAddr.String(), s.ConsensusCmtAddr.String(), s.OperatorEthAddr.String()).Set(float64(s.Power))
	jailedGauge.WithLabelValues(s.ConsensusEthAddr.String()).Set(boolToFloat(s.Jailed))
	bondedGauge.WithLabelValues(s.ConsensusEthAddr.String()).Set(boolToFloat(s.Bonded))
	tombstonedGauge.WithLabelValues(s.ConsensusEthAddr.String()).Set(boolToFloat(s.Tombstoned))
	uptimeGauge.WithLabelValues(s.ConsensusEthAddr.String()).Set(s.Uptime)
	rewardsGauge.WithLabelValues(s.ConsensusEthAddr.String()).Set(s.Rewards)
}
