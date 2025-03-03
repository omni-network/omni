package flowgen

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	jobsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "monitor",
		Subsystem: "flowgen",
		Name:      "jobs_total",
		Help:      "Total number of jobs executed",
	})

	jobsFailed = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "monitor",
		Subsystem: "flowgen",
		Name:      "failures_total",
		Help:      "Total number of failed jobs",
	})
)
