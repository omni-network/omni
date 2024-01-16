package gitinfo

import "runtime/debug"

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
