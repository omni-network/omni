package buildinfo

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	commitGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "buildinfo",
		Name:      "git_commit",
		Help:      "Constant gauge with label 'commit' set to the build info git commit hash.",
	}, []string{"commit"})

	timestampGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "buildinfo",
		Name:      "git_timestamp_unix",
		Help:      "Build info git commit timestamp in unix seconds.",
	})

	versionGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "buildinfo",
		Name:      "version",
		Help:      "Constant gauge with label 'version' set to the build info omni version",
	}, []string{"version"})
)
