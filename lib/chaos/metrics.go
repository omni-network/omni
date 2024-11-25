package chaos

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	chaosErrorCount = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "lib",
		Subsystem: "chaos",
		Name:      "errors_total",
		Help:      "Total number of chaos errors injected",
	})
)
