package relayer

import (
	"testing"

	"github.com/omni-network/omni/lib/xchain"

	"github.com/stretchr/testify/require"
)

func TestGroupingMsgsByCost(t *testing.T) {
	t.Parallel()

	// Constants defined in contracts/src/protocol/OmniPortalConstants.sol
	const msgGasLimitMin uint64 = 21_000
	const msgGasLimitMax uint64 = 5_000_000

	tests := []struct {
		name     string
		msgGas   []uint64
		expected []uint64
	}{
		{
			name:     "empty",
			msgGas:   nil,
			expected: uints(subGasBase),
		},
		{
			name:     "one min",
			msgGas:   uints(msgGasLimitMin),
			expected: uints(subGasBase + subGasXmsgOverhead + msgGasLimitMin),
		},
		{
			name:     "one max",
			msgGas:   uints(msgGasLimitMax),
			expected: uints(subGasBase + subGasXmsgOverhead + msgGasLimitMax),
		},
		{
			name:   "two min",
			msgGas: uints(msgGasLimitMin, msgGasLimitMin),
			expected: uints(
				subGasBase + ((subGasXmsgOverhead + msgGasLimitMin) * 2),
			),
		},
		{
			name:   "two max",
			msgGas: uints(msgGasLimitMax, msgGasLimitMax),
			expected: uints(
				subGasBase+subGasXmsgOverhead+msgGasLimitMax,
				subGasBase+subGasXmsgOverhead+msgGasLimitMax,
			),
		},
		{
			name:   "many",
			msgGas: uints(msgGasLimitMin, msgGasLimitMax, msgGasLimitMin, msgGasLimitMax, msgGasLimitMin, msgGasLimitMax),
			expected: uints(
				subGasBase+(subGasXmsgOverhead+msgGasLimitMax)+(subGasXmsgOverhead+msgGasLimitMin)*2,
				subGasBase+(subGasXmsgOverhead+msgGasLimitMax)+(subGasXmsgOverhead+msgGasLimitMin)*1,
				subGasBase+(subGasXmsgOverhead+msgGasLimitMax)+(subGasXmsgOverhead+msgGasLimitMin)*0,
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var xmsgs []xchain.Msg
			for _, gas := range test.msgGas {
				xmsgs = append(xmsgs, xchain.Msg{DestGasLimit: gas})
			}

			groups := groupMsgsByCost(xmsgs)

			require.Len(t, groups, len(test.expected))
			for i, group := range groups {
				require.Equal(t, int(test.expected[i]), int(naiveSubmissionGas(group)))

				// Ensure we never cross the max
				require.LessOrEqual(t, int(naiveSubmissionGas(group)), int(subGasMax))
			}
		})
	}
}

func uints(ii ...uint64) []uint64 {
	return ii
}
