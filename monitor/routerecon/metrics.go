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
		Help:      "Latest completed attest offset per stream",
	}, []string{"stream"})
)
