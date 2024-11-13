package cursor

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	confirmedOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "cursors",
		Name:      "confirmed_offset",
		Help:      "Confirmed cursor offset",
	}, []string{"src_chain_version", "dst_chain"})

	latestOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "cursors",
		Name:      "latest_offset",
		Help:      "Latest cursor offset",
	}, []string{"src_chain_version", "dst_chain"})
)
