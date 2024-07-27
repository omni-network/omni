package xmonitor

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	emitMsgOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "xchain",
		Name:      "emit_stream_offset",
		Help:      "The latest emitted xmsg offset on a source chain for a specific destination chain",
	}, []string{"stream"})

	submitMsgOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "xchain",
		Name:      "submit_stream_offset",
		Help:      "The latest submitted xmsg stream offset on a destination chain for a specific source chain",
	}, []string{"stream"})

	submitBlockOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "xchain",
		Name:      "submit_block_offset",
		Help:      "The latest submitted xblock offset on a destination chain for a specific source chain",
	}, []string{"stream"})

	headHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "xchain",
		Name:      "head_height",
		Help:      "The latest height of different types of head blocks on a specific chain",
	}, []string{"chain", "type"})

	attestedHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "xchain",
		Name:      "halo_attested_height",
		Help:      "The latest halo attested height of a specific chain",
	}, []string{"chain"})

	attestedOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "xchain",
		Name:      "halo_attest_offset",
		Help:      "The latest halo attest offset of a specific chain",
	}, []string{"chain"})

	attestedMsgOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "monitor",
		Subsystem: "xchain",
		Name:      "halo_attested_stream_offset",
		Help:      "The latest halo attested msg offset of a specific stream",
	}, []string{"stream"})
)
