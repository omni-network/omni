package types_test

import (
	"testing"

	"github.com/omni-network/omni/e2e/types"

	"github.com/stretchr/testify/require"
)

func TestCanaryRegex(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Service string
		Canary  bool
	}{
		{"validator01", true},
		{"validator01_evm", true},
		{"validator02", false},
		{"validator02_evm", false},
		{"fullnode01", true},
		{"fullnode01_evm", true},
		{"fullnode02", false},
		{"fullnode02_evm", false},
		{"archive01", true},
		{"archive01_evm", true},
		{"archive02", false},
		{"archive02_evm", false},
		{"relayer", true},
		{"monitor", true},
		{"solver", true},
	}
	for _, test := range tests {
		t.Run(test.Service, func(t *testing.T) {
			t.Parallel()
			ok := types.ServiceConfig{Regexp: "canary"}.MatchService(test.Service)
			require.Equal(t, test.Canary, ok)

			ok = types.ServiceConfig{Regexp: "non-canary"}.MatchService(test.Service)
			require.Equal(t, !test.Canary, ok)
		})
	}
}
