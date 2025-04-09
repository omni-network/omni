package bridging

import (
	"testing"

	"github.com/omni-network/omni/lib/bi"
	solver "github.com/omni-network/omni/solver/app"

	"github.com/stretchr/testify/require"
)

func TestSplitOrderAmounts(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Name     string
		MinSpend float64
		MaxSpend float64
		Total    float64
		Parallel int
		Expected []float64
	}{
		{
			Name:     "empty",
			MinSpend: 1,
			MaxSpend: 10,
			Total:    0,
			Parallel: 1,
			Expected: nil,
		},
		{
			Name:     "single",
			MinSpend: 1,
			MaxSpend: 10,
			Total:    5,
			Parallel: 1,
			Expected: []float64{5},
		},
		{
			Name:     "single too small",
			MinSpend: 10,
			MaxSpend: 100,
			Total:    5,
			Parallel: 1,
			Expected: nil,
		},
		{
			Name:     "single too big",
			MinSpend: 1,
			MaxSpend: 10,
			Total:    15,
			Parallel: 1,
			Expected: []float64{10},
		},
		{
			Name:     "multiple 2",
			MinSpend: 1,
			MaxSpend: 10,
			Total:    22.2,
			Parallel: 2,
			Expected: []float64{10, 10},
		},
		{
			Name:     "multiple 3",
			MinSpend: 1,
			MaxSpend: 10,
			Total:    22.2,
			Parallel: 3,
			Expected: []float64{7.4, 7.4, 7.4},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			bounds := solver.SpendBounds{
				MinSpend: bi.Ether(test.MinSpend),
				MaxSpend: bi.Ether(test.MaxSpend),
			}
			actual := splitOrderAmounts(bounds, bi.Ether(test.Total), test.Parallel)

			var actualFloats []float64
			for _, amt := range actual {
				actualFloats = append(actualFloats, bi.ToEtherF64(amt))
			}

			require.Equal(t, test.Expected, actualFloats)
		})
	}
}
