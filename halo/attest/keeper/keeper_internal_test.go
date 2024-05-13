package keeper

import (
	"fmt"
	"testing"
)

// AttestTable returns the attestations ORM table.
func (k *Keeper) AttestTable() AttestationTable {
	return k.attTable
}

// SignatureTable returns the attestations ORM table.
func (k *Keeper) SignatureTable() SignatureTable {
	return k.sigTable
}

func TestWindowCompose(t *testing.T) {
	t.Parallel()
	const window = 64
	const mid = 32
	const bigMid = 256

	const in = 0
	const above = 1
	const below = -1

	tests := []struct {
		Mid      uint64
		Target   uint64
		Expected int
	}{
		{mid, mid, in},
		{mid, mid + 1, in},
		{mid, mid - 1, in},
		{mid, mid + window, in},
		{mid, 0, in},
		{mid, mid + window + 1, above},
		{mid, mid + window + window, above},
		{bigMid, bigMid - window - 1, below},
		{bigMid, bigMid - window - window, below},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("mid_%d_target_%d", tt.Mid, tt.Target), func(t *testing.T) {
			t.Parallel()
			got := windowCompare(window, tt.Mid, tt.Target)
			if got != tt.Expected {
				t.Errorf("Test %d: Expected %d, got %d", i, tt.Expected, got)
			}
		})
	}
}
