package contract

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	spendTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "monitor",
		Subsystem: "xfeemngr",
		Name:      "spend_gwei_total",
		Help:      "Total amount of tokens spent by the xfeemngr on a chain (in gwei)",
	}, []string{"chain", "token", "method"})
)
