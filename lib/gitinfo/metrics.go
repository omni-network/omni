package gitinfo

import "github.com/prometheus/client_golang/prometheus"

var (
	commitGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "git",
		Name:      "commit",
		Help:      "Constant gauge with label 'commit' set to the current git commit hash.",
	}, []string{"commit"})

	timestampGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "lib",
		Subsystem: "git",
		Name:      "timestamp",
		Help:      "Constant gauge with label 'timestamp' set to the current git commit timestamp.",
	}, []string{"timestamp"})
)
