package app

import (
	"context"
	"sync"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/solver/job"

	"github.com/ethereum/go-ethereum/core/types"
)

// asyncWorkFunc abstracts the async processing of the job.
// It only returns an error if the job is invalid.
type asyncWorkFunc func(context.Context, *job.Job) error

func newAsyncWorkerFunc(
	jobDB job.DB,
	procs map[uint64]eventProcFunc,
	namer func(uint64) string,
) asyncWorkFunc {
	var active sync.Map // Stores active job IDs.
	return func(ctx context.Context, j *job.Job) error {
		chainName := namer(j.GetChainId())
		elog, err := j.EventLog()
		if err != nil {
			return errors.Wrap(err, "parse event log [BUG]")
		}

		proc, ok := procs[j.GetChainId()]
		if !ok {
			return errors.New("unknown chain [BUG]")
		}

		if _, ok := active.LoadOrStore(j.GetId(), struct{}{}); ok {
			// Already processing this job.
			return nil
		}

		go func() {
			workActive.WithLabelValues(chainName).Inc()
			defer workActive.WithLabelValues(chainName).Dec()

			ctx := log.WithCtx(ctx, "height", elog.BlockNumber)

			backoff := expbackoff.New(ctx)
			for {
				err := processJobOnce(ctx, jobDB, proc, j.GetId(), elog)
				if ctx.Err() != nil {
					return // Shutdown
				} else if err == nil {
					return // Done
				}

				log.Warn(ctx, "Failed to process job (will retry)", err, "job_id", j.GetId())
				backoff()
			}
		}()

		return nil
	}
}

func processJobOnce(ctx context.Context, jobDB job.DB, proc eventProcFunc, jobID uint64, elog types.Log) error {
	if ok, err := jobDB.Exists(ctx, jobID); err != nil {
		return err
	} else if !ok {
		log.Warn(ctx, "Job unexpectedly deleted [BUG]", nil)
		return nil
	}

	err := proc(ctx, elog)
	if err != nil {
		return err
	}

	return jobDB.Delete(ctx, jobID)
}
