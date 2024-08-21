package routerecon

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	reconSuccess = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "monitor",
		Subsystem: "routerecon",
		Name:      "success_total",
		Help:      "Total count of successful route recons",
	})

	reconFailure = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "monitor",
		Subsystem: "routerecon",
		Name:      "failure_total",
		Help:      "Total count of failed route recons",
	})
)
