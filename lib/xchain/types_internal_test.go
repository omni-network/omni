package xchain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfLevelFuzzy(t *testing.T) {
	t.Parallel()
	var fuzzies []ConfLevel
	for conf := ConfUnknown; conf < confSentinel; conf++ {
		if conf.IsFuzzy() {
			fuzzies = append(fuzzies, conf)
		}
	}

	require.EqualValues(t, fuzzies, FuzzyConfLevels())
}
