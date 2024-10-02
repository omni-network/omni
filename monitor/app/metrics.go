package monitor

import (
	"time"

	"github.com/omni-network/omni/lib/promutil"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	syncDiff = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "rpc",
		Name:      "sync_diff",
		Help:      "Maximum sync difference (concurrent latest heights) per chain",
	}, []string{"chain"})

	histBaseline = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "monitor",
		Subsystem: "cprovider",
		Name:      "historical_baseline_seconds",
		Help:      "Baseline time (in seconds) to stream historical approved attestation",
		Buckets:   prometheus.ExponentialBucketsRange(time.Second.Seconds(), time.Hour.Seconds(), 8),
	}, []string{"chain", "size"})

	plannedUpgradeGauge = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Name:      "planned_upgrade",
		Help:      "Constant gauge set to 1 with labels matching the upgrade name and height",
	}, []string{"upgrade", "height"})
)
