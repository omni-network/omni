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

	emitCursor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "monitor",
		Name:      "emit_cursor",
		Help:      "The latest emitted cursor on a source chain for a specific destination chain",
	}, []string{"src_chain", "dst_chain"})

	submitCursor = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "relayer",
		Subsystem: "monitor",
		Name:      "submit_cursor",
		Help:      "The latest submitted cursor on a destination chain for a specific source chain",
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
)
