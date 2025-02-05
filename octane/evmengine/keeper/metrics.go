package keeper

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	withdrawals = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "octane",
		Subsystem: "evmengine",
		Name:      "withdrawals_queued",
		Help:      "The number of queued withdrawals",
	})
)
