package keeper

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	approvedHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "attest",
		Name:      "approved_height",
		Help:      "The height of latest approved attestation per source chain",
	}, []string{"chain_version"})

	approvedOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "attest",
		Name:      "approved_offset",
		Help:      "The offset of latest approved attestation per source chain",
	}, []string{"chain_version"})

	votesProposed = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "halo",
		Subsystem: "attest",
		Name:      "proposed_votes",
		Help:      "The number of votes proposed per block per source chain",
		Buckets:   []float64{1, 2, 5, 10, 25, 50, 100, 250, 500, 1000},
	}, []string{"chain_version"})

	votesExtended = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "halo",
		Subsystem: "attest",
		Name:      "extended_votes",
		Help:      "The number of votes included by a validator per block per source chain",
		Buckets:   []float64{1, 2, 5, 10, 25, 50, 100, 250, 500, 1000},
	}, []string{"chain_version"})

	dbLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "halo",
		Subsystem: "attest",
		Name:      "db_latency_seconds",
		Help:      "Latency (in seconds) for each db function call (both internal and external)",
		Buckets:   []float64{.001, .002, .005, .01, .025, .05, .1},
	}, []string{"method"})

	doubleSignCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "halo",
		Subsystem: "attest",
		Name:      "double_sign_total",
		Help:      "Total number of double sign votes detected per validator",
	}, []string{"validator"})
)

func latency(method string) func() {
	t0 := time.Now()
	return func() {
		dbLatency.WithLabelValues(method).Observe(time.Since(t0).Seconds())
	}
}
