package avs

import (
	"github.com/omni-network/omni/lib/promutil"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	numOperators = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "avs",
		Name:      "num_operators",
		Help:      "The number of operators registered with the AVS",
	})

	totalDelegations = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "avs",
		Name:      "total_delegations",
		Help:      "The total amount of delegations made all operators registered with the AVS",
	})

	operatorStake = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "avs",
		Name:      "operator_stake",
		Help:      "The total amount staked (self-delegations) by operators registered with the AVS",
	}, []string{"operator"})

	operatorDelegations = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "avs",
		Name:      "operator_delegations",
		Help:      "The total amount of delegations (non self-delegations) made to operators registered with the AVS",
	}, []string{"operator"})
)
