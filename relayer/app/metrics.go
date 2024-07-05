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
)
