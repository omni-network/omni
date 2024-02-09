package txmgr

import (
	"context"
	"math/big"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type priceBumpTest struct {
	prevGasTip  int64
	prevBaseFee int64
	newGasTip   int64
	newBaseFee  int64
	expectedTip int64
	expectedFC  int64
}

func (tc *priceBumpTest) run(t *testing.T) {
	t.Helper()
	prevFC := calcGasFeeCap(big.NewInt(tc.prevBaseFee), big.NewInt(tc.prevGasTip))

	tip, fc := updateFees(context.Background(), big.NewInt(tc.prevGasTip), prevFC, big.NewInt(tc.newGasTip),
		big.NewInt(tc.newBaseFee))

	require.Equal(t, tc.expectedTip, tip.Int64(), "tip must be as expected")
	require.Equal(t, tc.expectedFC, fc.Int64(), "fee cap must be as expected")
}

func TestUpdateFees(t *testing.T) {
	t.Parallel()
	require.Equal(t, int64(10), PriceBump, "test must be updated if priceBump is adjusted")
	tests := []priceBumpTest{
		{
			prevGasTip: 100, prevBaseFee: 1000,
			newGasTip: 90, newBaseFee: 900,
			expectedTip: 110, expectedFC: 2310,
		},

		{
			prevGasTip: 100, prevBaseFee: 1000,
			newGasTip: 101, newBaseFee: 1000,
			expectedTip: 110, expectedFC: 2310,
		},

		{
			prevGasTip: 100, prevBaseFee: 1000,
			newGasTip: 100, newBaseFee: 1001,
			expectedTip: 110, expectedFC: 2310,
		},

		{
			prevGasTip: 100, prevBaseFee: 1000,
			newGasTip: 101, newBaseFee: 900,
			expectedTip: 110, expectedFC: 2310,
		},

		{
			prevGasTip: 100, prevBaseFee: 1000,
			newGasTip: 90, newBaseFee: 1010,
			expectedTip: 110, expectedFC: 2310,
		},

		{
			prevGasTip: 100, prevBaseFee: 1000,
			newGasTip: 101, newBaseFee: 2000,
			expectedTip: 110, expectedFC: 4110,
		},

		{
			prevGasTip: 100, prevBaseFee: 1000,
			newGasTip: 120, newBaseFee: 900,
			expectedTip: 120, expectedFC: 2310,
		},

		{
			prevGasTip: 100, prevBaseFee: 1000,
			newGasTip: 120, newBaseFee: 1100,
			expectedTip: 120, expectedFC: 2320,
		},

		{
			prevGasTip: 100, prevBaseFee: 1000,
			newGasTip: 120, newBaseFee: 1140,
			expectedTip: 120, expectedFC: 2400,
		},

		{
			prevGasTip: 100, prevBaseFee: 1000,
			newGasTip: 120, newBaseFee: 1200,
			expectedTip: 120, expectedFC: 2520,
		},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			test.run(t)
		})
	}
}
