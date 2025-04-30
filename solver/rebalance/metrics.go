package rebalance

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	balanceDecifit = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "solver",
		Subsystem: "rebalance",
		Name:      "balance_deficit",
		Help:      "The size of deficit for a token on a chain",
	}, []string{"chain", "token"})

	balanceSurplus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "solver",
		Subsystem: "rebalance",
		Name:      "balance_surplus",
		Help:      "The size of surplus for a token on a chain",
	}, []string{"chain", "token"})

	balanceCurrent = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "solver",
		Subsystem: "rebalance",
		Name:      "balance_current",
		Help:      "The current balance for a token on a chain",
	}, []string{"chain", "token"})

	// Threshold gauges will mostly be static, unless they are changed via upgrade.

	thresholdTarget = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "solver",
		Subsystem: "rebalance",
		Name:      "threshold_target",
		Help:      "The target balance for a token on a chain",
	}, []string{"chain", "token"})

	thresholdSurplus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "solver",
		Subsystem: "rebalance",
		Name:      "threshold_surplus",
		Help:      "The surplus threshold for a token on a chain",
	}, []string{"chain", "token"})

	thresholdMin = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "solver",
		Subsystem: "rebalance",
		Name:      "threshold_min",
		Help:      "The minimum threshold for a token on a chain",
	}, []string{"chain", "token"})
)
