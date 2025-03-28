package relayer

import (
	"fmt"
	"testing"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"

	"github.com/stretchr/testify/require"
)

func TestQuorumSigs(t *testing.T) {
	t.Parallel()

	const total = 10
	var vals []cchain.PortalValidator
	for i := range total {
		val := cchain.PortalValidator{
			Address: common.BytesToAddress([]byte{byte(i)}),
			Power:   int64(i), // Power from 0 and 9
		}
		vals = append(vals, val)
	}

	// totalPower := 0 + 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 // 45
	// quorum := totalPower * 2 / 3                        // +30

	tests := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{},
			output: []int{},
		},
		{
			input:  []int{0},
			output: []int{},
		},
		{
			input:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, // 45
			output: []int{9, 8, 7, 6, 5},                // 35
		},
		{
			input:  []int{9, 8, 7, 6, 1, 0}, // 31
			output: []int{9, 8, 7, 6, 1},    // 31
		},
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8}, // 36
			output: []int{8, 7, 6, 5, 4, 3},       // 33
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.input), func(t *testing.T) {
			t.Parallel()

			var sigs []xchain.SigTuple
			for _, i := range test.input {
				sigs = append(sigs, xchain.SigTuple{ValidatorAddress: vals[i].Address})
			}

			actual, err := quorumSigs(vals, sigs)
			if len(test.output) == 0 {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			var got []int
			for _, sig := range actual {
				got = append(got, int(sig.ValidatorAddress[19]))
			}
			require.Equal(t, test.output, got)
		})
	}
}

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
