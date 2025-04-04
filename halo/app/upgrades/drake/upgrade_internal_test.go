package drake

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestStakingParams(t *testing.T) {
	t.Parallel()

	p := StakingParams
	require.Equal(t, time.Duration(0), p.UnbondingTime)
	require.Equal(t, uint32(30), p.MaxValidators)
	require.Equal(t, uint32(7), p.MaxEntries)
	require.Equal(t, uint32(10_000), p.HistoricalEntries)
	require.Equal(t, "stake", p.BondDenom)
	require.True(t, p.MinCommissionRate.IsZero())
}
