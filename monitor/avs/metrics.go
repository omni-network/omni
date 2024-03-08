package avs

import (
	"github.com/omni-network/omni/lib/promutil"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	numOperatorsGuage = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "avs",
		Name:      "num_operators",
		Help:      "The number of operators registered with the AVS",
	})

	totalDelegationsGuage = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "avs",
		Name:      "total_delegations",
		Help:      "The total amount of delegations made all operators registered with the AVS",
	})

	operatorStakeGuage = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "avs",
		Name:      "operator_stake",
		Help:      "The total amount staked (self-delegations) by operators registered with the AVS",
	}, []string{"operator"})

	operatorDelegationsGuage = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "avs",
		Name:      "operator_delegations",
		Help:      "The total amount of delegations (non self-delegations) made to operators registered with the AVS",
	}, []string{"operator"})

	ownerGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "avs",
		Name:      "owner",
		Help:      "Constant gauge with label 'owner' set to the owner of the AVS",
	}, []string{"owner"})

	pausedGuage = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "avs",
		Name:      "paused",
		Help:      "Set to 1 if the AVS is paused, 0 otherwise",
	})

	allowlistEnabledGuage = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "avs",
		Name:      "allowlist_enabled",
		Help:      "Set to 1 if the AVS is using an allowlist, 0 otherwise",
	})

	minStakeGuage = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "avs",
		Name:      "min_stake",
		Help:      "The minimum amount of stake required to be an operator",
	})

	maxOperatorsGuage = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "avs",
		Name:      "max_operators",
		Help:      "The maximum number of operators allowed",
	})

	strategyParamsGuage = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "avs",
		Name:      "strategy_params",
		Help:      "The AVS strategy parameters, label 'strategy' is the strategy  address, value is the multiplier.",
	}, []string{"strategy"})
)
