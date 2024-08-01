package emitcache

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	hitCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "monitor",
		Subsystem: "emitcache",
		Name:      "hit_total",
		Help:      "Total number of emitcache hits per stream",
	}, []string{"stream"})

	missCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "monitor",
		Subsystem: "emitcache",
		Name:      "miss_total",
		Help:      "Total number of emitcache misses per stream",
	}, []string{"stream"})
)
