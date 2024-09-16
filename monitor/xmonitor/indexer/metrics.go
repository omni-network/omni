package indexer

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	latencyHist = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "monitor",
		Subsystem: "indexer",
		Name:      "latency_seconds",
		Help:      "Cross chain latency in seconds per stream per xdapp (submit-emit timestamp)",
		Buckets:   prometheus.ExponentialBucketsRange(time.Second.Seconds(), time.Hour.Seconds(), 10),
	}, []string{"stream", "xdapp"})

	successCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "monitor",
		Subsystem: "indexer",
		Name:      "success_total",
		Help:      "Total number of successful cross chain transactions per stream per xdapp",
	}, []string{"stream", "xdapp"})

	revertCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "monitor",
		Subsystem: "indexer",
		Name:      "revert_total",
		Help:      "Total number of reverted cross chain transactions per stream per xdapp",
	}, []string{"stream", "xdapp"})

	fuzzyOverrideCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "monitor",
		Subsystem: "indexer",
		Name:      "fuzzy_override_total",
		Help:      "Total number of fuzzy override cross chain transactions per stream per xdapp",
	}, []string{"stream", "xdapp"})

	excessGasHist = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "monitor",
		Subsystem: "indexer",
		Name:      "excess_gas",
		Help:      "Excess gas per stream per xdapp (msg.GasLimit - receipt.GasUsed)",
		Buckets:   prometheus.ExponentialBucketsRange(1, 1e6, 10),
	}, []string{"stream", "xdapp"})
)

type sample struct {
	Stream        string
	XDApp         string
	Latency       time.Duration
	ExcessGas     uint64
	Success       bool
	FuzzyOverride bool
}

func instrumentSample(s sample) {
	// Initialize success/revert counters so both exist
	revertCounter.WithLabelValues(s.Stream, s.XDApp).Add(0)
	successCounter.WithLabelValues(s.Stream, s.XDApp).Add(0)
	if s.Success {
		successCounter.WithLabelValues(s.Stream, s.XDApp).Inc()
	} else {
		revertCounter.WithLabelValues(s.Stream, s.XDApp).Inc()
	}

	if s.FuzzyOverride {
		fuzzyOverrideCounter.WithLabelValues(s.Stream, s.XDApp).Inc()
	}
	latencyHist.WithLabelValues(s.Stream, s.XDApp).Observe(s.Latency.Seconds())
	excessGasHist.WithLabelValues(s.Stream, s.XDApp).Observe(float64(s.ExcessGas))
}
