package app

import (
	"context"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/solver/job"
)

// asyncWorkFunc abstracts the async processing of the job.
// It only returns an error if the job is invalid.
type asyncWorkFunc func(context.Context, *job.Job) error

func newAsyncWorkerFunc(
	jobDB *job.DB,
	procs map[uint64]eventProcFunc,
	namer func(uint64) string,
) asyncWorkFunc {
	var active sync.Map // Stores active job IDs.
	return func(ctx context.Context, j *job.Job) error {
		chainName := namer(j.GetChainId())
		event, err := jobToEvent(j)
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

			ctx, span := startTrace(ctx, chainName, event.OrderID, event.Status)
			defer span.End()

			ctx = log.WithCtx(ctx, "height", event.Height, "src_chain", chainName)

			backoff := expbackoff.New(ctx)
			for {
				err := processJobOnce(ctx, jobDB, proc, j.GetId(), event)
				if ctx.Err() != nil {
					return // Shutdown
				} else if err == nil {
					// Done
					duration := time.Since(j.GetCreatedAt().AsTime()).Seconds()
					workDuration.WithLabelValues(chainName, event.Status.String()).Observe(duration)

					return
				}

				log.Warn(ctx, "Failed to process job (will retry)", err, "job_id", j.GetId(), "order_id", event.OrderID, "status", event.Status)
				workErrors.WithLabelValues(chainName, event.Status.String()).Inc()
				backoff()
			}
		}()

		return nil
	}
}

func processJobOnce(ctx context.Context, jobDB *job.DB, proc eventProcFunc, jobID uint64, e Event) error {
	if ok, err := jobDB.Exists(ctx, jobID); err != nil {
		return err
	} else if !ok {
		log.Warn(ctx, "Job unexpectedly deleted [BUG]", nil)
		return nil
	}

	err := proc(ctx, e)
	if err != nil {
		return err
	}

	return jobDB.Delete(ctx, jobID)
}

func jobToEvent(job *job.Job) (Event, error) {
	orderID, err := cast.Array32(job.GetOrderId())
	if err != nil {
		return Event{}, errors.Wrap(err, "cast order id")
	}

	status, err := umath.ToUint8(job.GetStatus())
	if err != nil {
		return Event{}, errors.Wrap(err, "cast status")
	}

	return Event{
		OrderID: orderID,
		Status:  solvernet.OrderStatus(status),
		Height:  job.GetHeight(),
		Tx:      job.GetTxString(),
	}, nil
}
