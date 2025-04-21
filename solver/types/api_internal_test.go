package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/stretchr/testify/require"
)

func TestFlexBigUnmarshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    *big.Int
		wantErr bool
	}{
		{
			name:  "hex string",
			input: `"0x123"`,
			want:  bi.N(0x123),
		},
		{
			name:  "decimal string",
			input: `"123"`,
			want:  bi.N(123),
		},
		{
			name:  "1 ether",
			input: `"1000000000000000000"`,
			want:  bi.Ether(1),
		},
		{
			name:  "zero hex",
			input: `"0x0"`,
			want:  bi.Zero(),
		},
		{
			name:  "zero decimal",
			input: `"0"`,
			want:  bi.Zero(),
		},
		{
			name:  "empty string",
			input: `""`,
			want:  bi.Zero(),
		},
		{
			name:    "invalid decimal",
			input:   `"abc"`,
			wantErr: true,
		},
		{
			name:    "invalid hex",
			input:   `"0xzzz"`,
			wantErr: true,
		},
		{
			name:    "json number",
			input:   `123`,
			wantErr: true,
		},
		{
			name:    "null",
			input:   `null`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b := new(bigIntJSON)
			err := json.Unmarshal([]byte(tt.input), b)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			tutil.RequireEQ(t, tt.want, b.ToIntOrZero())
		})
	}
}

func TestFlexBigMarshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input *big.Int
	}{
		{
			name:  "positive number",
			input: bi.N(123),
		},
		{
			name:  "zero",
			input: bi.N(0),
		},
		{
			name:  "1 ether",
			input: bi.Ether(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			expect := fmt.Sprintf("%q", hexutil.EncodeBig(tt.input))

			got, err := json.Marshal((*bigIntJSON)(tt.input))
			require.NoError(t, err)
			require.Equal(t, expect, string(got), "expected %s, got %s", expect, string(got))
		})
	}
}
