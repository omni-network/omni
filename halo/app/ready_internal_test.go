package app

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Checking that ExecutionP2PPeers does not appear in the JSON output if it's
// set to nil, but appears if it's set to 0 or any other number.
func TestReadinessStatusMarshaling(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  *uint64
		output string
	}{
		{
			name:  "ExecutionP2PPeers is 22",
			input: uint64Ptr(22),
		},
		{
			name:  "ExecutionP2PPeers is 0",
			input: uint64Ptr(0),
		},
		{
			name:  "ExecutionP2PPeers is nil",
			input: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			status := readinessStatus{
				ExecutionP2PPeers: tt.input,
			}
			data, err := json.Marshal(&status)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tt.input != nil {
				require.True(t,
					strings.Contains(string(data), fmt.Sprintf("\"execution_p2p_peers\":%v", *tt.input)),
					tt.name)
			} else {
				require.False(t, strings.Contains(string(data), "execution_p2p_peers"), tt.name)
			}
		})
	}
}

func uint64Ptr(i uint64) *uint64 {
	return &i
}
