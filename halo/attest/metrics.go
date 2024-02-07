package attest

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//nolint:gochecknoglobals // Promauto metrics are global.
var (
	createLag = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "attester",
		Name:      "create_lag_seconds",
		Help: "Latest lag between attestation creation and xblock timestamp (in seconds) per source chain. " +
			"Alert if too high.",
	}, []string{"chain"})

	createHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "attester",
		Name:      "create_height",
		Help:      "Latest created attestation height per source chain. Alert if not growing.",
	}, []string{"chain"})

	commitHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "attester",
		Name:      "commit_height",
		Help:      "Latest committed attestation height per source chain. Alert if not growing.",
	}, []string{"chain"})

	availableCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "attester",
		Name:      "available_attestations",
		Help:      "Current number of available attestations per source chain. Alert if growing.",
	}, []string{"chain"})

	proposedCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "attester",
		Name:      "proposed_attestations",
		Help:      "Current number of proposed attestations per source chain. Alert if growing.",
	}, []string{"chain"})
)
