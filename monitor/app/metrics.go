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
		Help:      "Height of current planned (non-processed) upgrade by name",
	}, []string{"upgrade"})

	appliedUpgradeGauge = promutil.NewResetGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Name:      "applied_upgrade",
		Help:      "Height of last applied (processed) upgrade by name",
	}, []string{"upgrade"})

	publicRPCSyncDiff = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "public_rpc",
		Name:      "sync_diff",
		Help:      "Sync difference (highest blocks) between public RPC and omni node",
	})

	gasTipCap = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "omni_evm",
		Name:      "gas_tip_cap_gwei",
		Help:      "Suggested OmniEVM gas tip cap in gwei",
	})
)
