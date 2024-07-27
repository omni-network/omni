package voter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	createLag = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "voter",
		Name:      "create_lag_seconds",
		Help: "Latest lag between vote creation and xblock timestamp (in seconds) per source chain version. " +
			"Alert if too high.",
	}, []string{"chain_version"})

	createHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "voter",
		Name:      "create_height",
		Help:      "Latest created vote height per source chain version. Alert if not growing.",
	}, []string{"chain_version"})

	createAttestOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "voter",
		Name:      "create_attest_offset",
		Help:      "Latest created vote attest offset per source chain version. Alert if not growing.",
	}, []string{"chain_version"})

	createMsgOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "voter",
		Name:      "create_stream_offset",
		Help:      "Latest created vote xmsg offset per stream",
	}, []string{"stream"})

	commitHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "voter",
		Name:      "commit_height",
		Help:      "Latest committed vote height per source chain version. Alert if not growing.",
	}, []string{"chain_version"})

	availableCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "voter",
		Name:      "available_votes",
		Help:      "Current number of available votes per source chain version. Alert if growing.",
	}, []string{"chain_version"})

	proposedCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "voter",
		Name:      "proposed_votes",
		Help:      "Current number of proposed votes per source chain version. Alert if growing.",
	}, []string{"chain_version"})

	proposedPerBlock = promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: "halo",
		Subsystem: "voter",
		Name:      "proposed_per_block",
		Help:      "Number of proposed votes per block.",
		Buckets:   []float64{1, 2, 5, 10, 25, 50, 100, 250, 500, 1000},
	})

	committedPerBlock = promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: "halo",
		Subsystem: "voter",
		Name:      "committed_per_block",
		Help:      "Number of committed votes per block.",
		Buckets:   []float64{1, 2, 5, 10, 25, 50, 100, 250, 500, 1000},
	})

	trimTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "halo",
		Subsystem: "voter",
		Name:      "trim_total",
		Help:      "Total number of votes trimmed per source chain version.",
	}, []string{"chain_version"})

	reorgTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "halo",
		Subsystem: "voter",
		Name:      "reorg_total",
		Help:      "Total number of reorgs detected per source chain version.",
	}, []string{"chain_version"})
)
