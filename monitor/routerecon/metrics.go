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

	reconCompletedOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "routerecon",
		Name:      "completed_offset",
		Help:      "Latest completed offset per stream. Only measured periodically",
	}, []string{"stream"})

	reconCompletedLag = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "routerecon",
		Name:      "completed_lag",
		Help:      "Routescan completed lag per stream. Difference between latest completed offset and submit cursor. Only measured periodically",
	}, []string{"stream"})
)
