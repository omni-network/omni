package comet

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//nolint:gochecknoglobals // Promauto metrics are global.
var (
	pendingHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "consensus",
		Name:      "pending_aggregate_height",
		Help: "Latest pending aggregate attestation height per source chain. Alert if not growing. " +
			"Or if out-growing approved height.",
	}, []string{"chain"})

	approvedHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "consensus",
		Name:      "approved_aggregate_height",
		Help:      "Latest approved aggregate attestation height per source chain. Alert if not growing.",
	}, []string{"chain"})
)
