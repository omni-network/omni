package keeper

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	approvedHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "halo",
		Subsystem: "attest",
		Name:      "approved_height",
		Help:      "The height of latest approved attestation per source chain",
	}, []string{"chain"})
)
