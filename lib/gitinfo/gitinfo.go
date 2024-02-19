package gitinfo

import (
	"context"
	"runtime/debug"

	"github.com/omni-network/omni/lib/log"
)

// Get returns the git commit hash and timestamp from the runtime build info.
func Get() (hash string, timestamp string) { //nolint:nonamedreturns // Disambiguate identical return types.
	hash, timestamp = "unknown", "unknown"
	hashLen := 7

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return hash, timestamp
	}

	for _, s := range info.Settings {
		if s.Key == "vcs.revision" {
			if len(s.Value) < hashLen {
				hashLen = len(s.Value)
			}
			hash = s.Value[:hashLen]
		} else if s.Key == "vcs.time" {
			timestamp = s.Value
		}
	}

	return hash, timestamp
}

// Instrument logs the git commit hash and timestamp from the runtime build info.
// It also sets the commit and timestamp metrics.
func Instrument(ctx context.Context) {
	commit, timestamp := Get()

	log.Info(ctx, "Version info", "git_commit", commit, "git_timestamp", timestamp)

	commitGauge.WithLabelValues(commit).Set(1)
	timestampGauge.WithLabelValues(timestamp).Set(1)
}
