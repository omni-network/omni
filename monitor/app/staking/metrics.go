package staking

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	rewardsAvg = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "staking",
		Name:      "avg_rewards",
		Help:      "Average staking rewards percentage [0,1]",
	})

	delegatorsCount = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "staking",
		Name:      "delegators",
		Help:      "Number of unique delegators",
	})

	stakeAvg = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "staking",
		Name:      "avg_stake",
		Help:      "Average stake size in gwei",
	})

	stakeMedian = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "staking",
		Name:      "median_stake",
		Help:      "Median stake size in gwei",
	})
)
