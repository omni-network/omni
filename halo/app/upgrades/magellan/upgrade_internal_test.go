package magellan

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/stretchr/testify/require"
)

func TestSlashingParams(t *testing.T) {
	t.Parallel()

	p := SlashingParams()
	require.True(t, p.SlashFractionDowntime.IsZero())
	require.True(t, p.SlashFractionDoubleSign.IsZero())
	require.Equal(t, time.Minute, p.DowntimeJailDuration)
	require.Equal(t, "0.050000000000000000", p.MinSignedPerWindow.String())
	require.EqualValues(t, 2000, p.SignedBlocksWindow)
}

func TestTargetInflation(t *testing.T) {
	t.Parallel()

	require.Equal(t, "0.110000000000000000", targetInflation.String())
	require.EqualValues(t, 21024000, MintParams.BlocksPerYear)

	totalStaked := math.NewInt(1000).MulRaw(params.Ether) // 1000 omni staked
	bondedRatio := math.LegacyNewDec(1)                   // 100% bonded

	minter := minttypes.InitialMinter(targetInflation)

	for range 5 {
		minter.Inflation = minter.NextInflationRate(MintParams, bondedRatio)
		minter.AnnualProvisions = minter.NextAnnualProvisions(MintParams, totalStaked)

		// Inflation doesn't change
		require.Equal(t, targetInflation.String(), minter.Inflation.String())
		// Anaual rewards is always totalStoked*11%
		require.Equal(t, totalStaked.ToLegacyDec().Mul(targetInflation).String(), minter.AnnualProvisions.String())

		blockRewards := minter.BlockProvision(MintParams)
		annualRewards := blockRewards.Amount.MulRaw(int64(MintParams.BlocksPerYear))
		delta := annualRewards.ToLegacyDec().Sub(minter.AnnualProvisions).Abs()
		gwei := math.LegacyNewDec(params.GWei)
		require.True(t, delta.LT(gwei))

		totalStaked = totalStaked.AddRaw(100) // 100 more omni staked
	}
}
