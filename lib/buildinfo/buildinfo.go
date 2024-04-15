package buildinfo

import (
	"context"
	"runtime/debug"
	"strings"
	"time"

	"github.com/omni-network/omni/lib/log"

	"github.com/spf13/cobra"
)

// version of the whole omni-monorepo and all binaries built from this git commit.
// This value is set by goreleaser at build-time and should be the git tag for official releases.
var version = "v0.1.3"

// Version returns the version of the whole omni-monorepo and all binaries built from this git commit.
func Version() string {
	return version
}

// Instrument logs the version, git commit hash, and timestamp from the runtime build info.
// It also sets metrics.
func Instrument(ctx context.Context) {
	commit, timestamp := get()

	log.Info(ctx, "Version info",
		"version", version,
		"git_commit", commit,
		"git_timestamp", timestamp,
	)

	versionGauge.WithLabelValues(version).Set(1)
	commitGauge.WithLabelValues(commit).Set(1)

	ts, _ := time.Parse(time.RFC3339, timestamp)
	timestampGauge.Set(float64(ts.Unix()))
}

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version information of this binary",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			commit, timestamp := get()

			var sb strings.Builder
			_, _ = sb.WriteString("Version       " + version)
			_, _ = sb.WriteString("\n")
			_, _ = sb.WriteString("Git Commit    " + commit)
			_, _ = sb.WriteString("\n")
			_, _ = sb.WriteString("Git Timestamp " + timestamp)
			_, _ = sb.WriteString("\n")

			cmd.Printf(sb.String())
		},
	}
}

// get returns the git commit hash and timestamp from the runtime build info.
func get() (hash string, timestamp string) { //nolint:nonamedreturns // Disambiguate identical return types.
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
