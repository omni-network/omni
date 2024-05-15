package relayer

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	bufferLen = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "worker",
		Name:      "buffer_length",
		Help:      "The length of the async send worker activeBuffer per destination chain. Alert if too high",
	}, []string{"dst_chain"})

	mempoolLen = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "worker",
		Name:      "mempool_length",
		Help:      "The length of the mempool per destination chain. Alert if too high",
	}, []string{"dst_chain"})

	workerResets = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "relayer",
		Subsystem: "worker",
		Name:      "reset_total",
		Help:      "The total number of times the worker has reset by destination chain. Alert if too high",
	}, []string{"dst_chain"})

	submissionTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "relayer",
		Subsystem: "worker",
		Name:      "submission_total",
		Help:      "The total number of submissions to destination chain from a specific source chain",
	}, []string{"src_chain", "dst_chain"})

	msgTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "relayer",
		Subsystem: "worker",
		Name:      "msg_total",
		Help:      "The total number of messages submitted to a destination chain from a specific source chain",
	}, []string{"src_chain", "dst_chain"})

	revertedSubmissionTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "relayer",
		Subsystem: "worker",
		Name:      "reverted_submission_total",
		Help:      "The total number of reverted (unsuccessful) submissions to destination chain from a specific source chain",
	}, []string{"src_chain", "dst_chain"})

	gasEstimated = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "relayer",
		Subsystem: "worker",
		Name:      "estimated_gas",
		Help:      "Estimated max gas usage by submissions by destination chain",
		Buckets:   prometheus.ExponentialBucketsRange(21_000, 10_000_000, 8),
	}, []string{"dst_chain"})

	emitMsgOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "monitor",
		Name:      "emit_stream_offset",
		Help:      "The latest emitted xmsg offset on a source chain for a specific destination chain",
	}, []string{"stream"})

	submitMsgOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "monitor",
		Name:      "submit_stream_offset",
		Help:      "The latest submitted xmsg stream offset on a destination chain for a specific source chain",
	}, []string{"stream"})

	submitBlockOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "monitor",
		Name:      "submit_block_offset",
		Help:      "The latest submitted xblock offset on a destination chain for a specific source chain",
	}, []string{"src_chain", "dst_chain"})

	accountBalance = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "monitor",
		Name:      "account_balance_ether",
		Help:      "The balance of the relayer account on a specific chain in ether",
	}, []string{"chain"})

	accountNonce = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "monitor",
		Name:      "account_nonce",
		Help:      "The nonce of the relayer account on a specific chain",
	}, []string{"chain"})

	headHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "monitor",
		Name:      "head_height",
		Help:      "The latest height of different types of head blocks on a specific chain",
	}, []string{"chain", "type"})

	attestedHeight = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "monitor",
		Name:      "halo_attested_height",
		Help:      "The latest halo attested height of a specific chain",
	}, []string{"chain"})

	attestedBlockOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "monitor",
		Name:      "halo_attested_block_offset",
		Help:      "The latest halo attested block offset of a specific chain",
	}, []string{"chain"})

	attestedMsgOffset = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "monitor",
		Name:      "halo_attested_stream_offset",
		Help:      "The latest halo attested msg offset of a specific stream",
	}, []string{"stream"})
)
