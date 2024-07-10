package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	syncDiff = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "rpc",
		Name:      "sync_diff",
		Help:      "Maximum sync difference (concurrent latest heights) per chain",
	}, []string{"chain"})
)
