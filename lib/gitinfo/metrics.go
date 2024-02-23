package gitinfo

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	commitGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "git",
		Name:      "commit",
		Help:      "Constant gauge with label 'commit' set to the build info git commit hash.",
	}, []string{"commit"})

	timestampGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "git",
		Name:      "timestamp_unix",
		Help:      "Build info git commit timestamp in unix seconds.",
	})
)
