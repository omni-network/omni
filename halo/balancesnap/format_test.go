package balancesnap_test

import (
	"math/big"
	"testing"

	"github.com/omni-network/omni/halo/balancesnap"

	"github.com/stretchr/testify/require"
)

func TestFormatBalance(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string // big.Int string representation
		expected string
	}{
		{
			name:     "zero",
			input:    "0",
			expected: "0.000000000000000000",
		},
		{
			name:     "one wei",
			input:    "1",
			expected: "0.000000000000000001",
		},
		{
			name:     "one token (10^18 wei)",
			input:    "1000000000000000000",
			expected: "1.000000000000000000",
		},
		{
			name:     "1234 tokens",
			input:    "1234000000000000000000",
			expected: "1_234.000000000000000000",
		},
		{
			name:     "1 million tokens",
			input:    "1000000000000000000000000",
			expected: "1_000_000.000000000000000000",
		},
		{
			name:     "1 billion tokens",
			input:    "1000000000000000000000000000",
			expected: "1_000_000_000.000000000000000000",
		},
		{
			name:     "7.5 billion tokens (Omni total supply)",
			input:    "7500000000000000000000000000",
			expected: "7_500_000_000.000000000000000000",
		},
		{
			name:     "fractional tokens",
			input:    "1234567890123456789",
			expected: "1.234567890123456789",
		},
		{
			name:     "large number with fractions",
			input:    "1234567890123456789012345678",
			expected: "1_234_567_890.123456789012345678",
		},
		{
			name:     "999 tokens",
			input:    "999000000000000000000",
			expected: "999.000000000000000000",
		},
		{
			name:     "1000 tokens (boundary)",
			input:    "1000000000000000000000",
			expected: "1_000.000000000000000000",
		},
		{
			name:     "small fraction",
			input:    "123",
			expected: "0.000000000000000123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			input := new(big.Int)
			input.SetString(tt.input, 10)

			result := balancesnap.FormatBalance(input)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatBalance_Nil(t *testing.T) {
	t.Parallel()

	result := balancesnap.FormatBalance(nil)
	require.Equal(t, "0.000000000000000000", result)
}
