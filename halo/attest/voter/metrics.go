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
		Help: "Latest lag between vote creation and xblock timestamp (in seconds) per source chain. " +
			"Alert if too high.",
	}, []string{"chain"})

	createHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "voter",
		Name:      "create_height",
		Help:      "Latest created vote height per source chain. Alert if not growing.",
	}, []string{"chain"})

	commitHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "voter",
		Name:      "commit_height",
		Help:      "Latest committed vote height per source chain. Alert if not growing.",
	}, []string{"chain"})

	availableCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "voter",
		Name:      "available_votes",
		Help:      "Current number of available votes per source chain. Alert if growing.",
	}, []string{"chain"})

	proposedCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "voter",
		Name:      "proposed_votes",
		Help:      "Current number of proposed votes per source chain. Alert if growing.",
	}, []string{"chain"})

	trimTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "halo",
		Subsystem: "voter",
		Name:      "trim_total",
		Help:      "Total number of votes trimmed per source chain.",
	}, []string{"chain"})
)
