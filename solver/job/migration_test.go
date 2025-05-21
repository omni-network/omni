package job_test

import (
	"testing"

	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/solver/job"
	"github.com/omni-network/omni/solver/jobprev"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	db "github.com/cosmos/cosmos-db"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
)

func TestMigrate(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	fuzzer := fuzz.New().NilChance(0).NumElements(0, 0)

	dir := t.TempDir()
	const dbName = "test"
	dbm, err := db.NewGoLevelDB(dbName, dir, nil)
	require.NoError(t, err)

	// Create old DB
	prevdb, err := jobprev.New(dbm)
	require.NoError(t, err)

	// Insert N jobs
	const N = 10

	var oldJobs []*jobprev.Job
	for i := range N {
		var elog types.Log
		fuzzer.Fuzz(&elog)
		elog.Topics = []common.Hash{
			solvernet.TopicClosed,
			tutil.RandomHash(), // OrderID
		}

		j, err := prevdb.Insert(ctx, uint64(i), elog)
		require.NoError(t, err)
		oldJobs = append(oldJobs, j)
	}
	require.NoError(t, dbm.Close())

	// Reopen DB
	dbm, err = db.NewGoLevelDB(dbName, dir, nil)
	require.NoError(t, err)

	// Create new DB
	newdb, err := job.New(dbm)
	require.NoError(t, err)

	count, err := job.Migrate(ctx, newdb)
	require.NoError(t, err)
	require.Equal(t, N, count)

	// Migrating again should not do anything
	count, err = job.Migrate(ctx, newdb)
	require.NoError(t, err)
	require.Zero(t, count)

	newJobs, err := newdb.All(ctx)
	require.NoError(t, err)
	require.Len(t, newJobs, N)

	for i, newJob := range newJobs {
		require.Equal(t, oldJobs[i].GetId(), newJob.GetId())
		require.Equal(t, oldJobs[i].GetChainId(), newJob.GetChainId())
		require.Equal(t, oldJobs[i].GetBlockHeight(), newJob.GetHeight())
		require.Equal(t, oldJobs[i].GetEventIndex(), newJob.GetEventIndex())
		require.Equal(t, oldJobs[i].GetCreatedAt(), newJob.GetCreatedAt())

		require.Empty(t, newJob.GetDeprecatedEventJson())

		oldLog, err := oldJobs[i].EventLog()
		require.NoError(t, err)

		orderID, status, err := solvernet.ParseEvent(oldLog)
		require.NoError(t, err)

		require.Equal(t, orderID[:], newJob.GetOrderId())
		require.EqualValues(t, status, newJob.GetStatus())
		require.Equal(t, oldLog.TxHash.String(), newJob.GetTxString())
	}

	// Add more jobs to the new DB
	for i := N; i < 2*N; i++ {
		var elog types.Log
		fuzzer.Fuzz(&elog)
		elog.Topics = []common.Hash{
			solvernet.TopicClosed,
			tutil.RandomHash(), // OrderID
		}

		_, err := newdb.InsertLog(ctx, uint64(i), elog)
		require.NoError(t, err)
	}
}
