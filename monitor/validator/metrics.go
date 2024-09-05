package validator

import (
	"github.com/omni-network/omni/lib/promutil"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	powerGauge = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "validator",
		Name:      "power",
		Help:      "Current power by validator",
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
)
