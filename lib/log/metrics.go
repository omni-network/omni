package log

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	logTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "lib",
		Subsystem: "log",
		Name:      "total",
		Help:      "Total number of log messages per level.",
	}, []string{"level"})
)

// zeroLogMetrics zeros the log metrics so they display nicely in grafana.
func zeroLogMetrics() {
	for _, level := range levels {
		logTotal.WithLabelValues(level).Add(0)
	}
}
