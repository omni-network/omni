package job_test

import (
	"context"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/solver/job"

	"github.com/ethereum/go-ethereum/core/types"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
)

func TestDB(t *testing.T) {
	t.Parallel()
	db, err := job.New(dbm.NewMemDB())
	require.NoError(t, err)

	fuzzer := ethclient.NewFuzzer(0)
	fuzzLog := func() types.Log {
		var resp types.Log
		fuzzer.Fuzz(&resp)

		return resp
	}

	ctx := context.Background()
	elogs := []types.Log{fuzzLog(), fuzzLog(), fuzzLog()}

	for i, elog := range elogs {
		t0 := time.Now()
		id := uint64(i + 1)
		j, err := db.Insert(ctx, id, elog)
		require.NoError(t, err)
		require.Equal(t, id, j.GetId())
		require.Equal(t, id, j.GetChainId())
		require.WithinRange(t, j.GetCreatedAt().AsTime(), t0, time.Now())

		ok, err := db.Exists(ctx, id)
		require.NoError(t, err)
		require.True(t, ok)
		ok, err = db.Exists(ctx, id+1)
		require.NoError(t, err)
		require.False(t, ok)

		// Idempotent
		j, err = db.Insert(ctx, id, elog)
		require.NoError(t, err)
		require.Equal(t, id, j.GetId())

		el, err := j.EventLog()
		require.NoError(t, err)
		require.Equal(t, elog, el)

		all, err := db.All(ctx)
		require.NoError(t, err)
		require.Len(t, all, i+1)
		last, err := all[len(all)-1].EventLog()
		require.NoError(t, err)
		require.Equal(t, elog, last)

		// Inserting duplicate with different hash fails
		elog2 := elog
		elog2.BlockHash[0] = ^elog2.BlockHash[0]
		_, err = db.Insert(ctx, id, elog2)
		require.ErrorContains(t, err, "duplicate")
	}

	for i := range len(elogs) {
		id := uint64(i + 1)
		err := db.Delete(ctx, id)
		require.NoError(t, err)

		all, err := db.All(ctx)
		require.NoError(t, err)
		require.Len(t, all, len(elogs)-i-1)

		for _, j := range all {
			require.NotEqual(t, id, j.GetId())
		}
	}
}
