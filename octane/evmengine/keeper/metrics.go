package keeper

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	insertedWithdrawals = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "octane",
		Subsystem: "evmengine",
		Name:      "withdrawals_inserted_total",
		Help:      "Total number of inserted withdrawals",
	})

	completedWithdrawals = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "octane",
		Subsystem: "evmengine",
		Name:      "withdrawals_completed_total",
		Help:      "Total number of completed withdrawals",
	})

	dustCounter = promauto.NewCounter(prometheus.CounterOpts{
		Namespace:   "octane",
		Subsystem:   "evmengine",
		Name:        "withdrawals_dust_total",
		Help:        "Total withdrawal creation dust in wei (dust is amounts less than 1 gwei)",
		ConstLabels: nil,
	})
)
