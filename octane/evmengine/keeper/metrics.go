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
)
