package flowgen

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/tokens"

	"github.com/robfig/cron/v3"
)

type Spend map[tokens.Token]*big.Int

type RunFunc func(context.Context, Spend) error

type Job struct {
	// Name is the friendly name of the job
	Name string

	// Run is the function to run
	Run RunFunc

	// Spec is a cron job schedule spec string
	Spec string

	// Spend is spend this job uses on one run (exluding gas)
	Spend Spend
}

// cronner returns a cron.Job for the given Job.
func cronner(ctx context.Context, j Job) cron.Job {
	return cron.FuncJob(func() {
		if err := j.Run(ctx, j.Spend); err != nil {
			log.Error(ctx, "Job failed", err, "name", j.Name)
		}

		log.Info(ctx, "Job finished", "name", j.Name)
	})
}
