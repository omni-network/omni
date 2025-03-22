package xchain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfLevelFuzzy(t *testing.T) {
	t.Parallel()
	var fuzzies []ConfLevel
	var minXs []ConfLevel
	for conf := ConfUnknown; conf < confSentinel; conf++ {
		if !conf.Valid() {
			continue
		}
		if conf.IsFuzzy() {
			fuzzies = append(fuzzies, conf)
		}
		if conf.MinX() > 0 {
			minXs = append(minXs, conf)
		}
	}

	require.EqualValues(t, fuzzies, FuzzyConfLevels())
	require.EqualValues(t, []ConfLevel{ConfMin1, ConfMin2}, minXs)
}
