package flowgen

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tokens"

	"github.com/robfig/cron/v3"
)

func Start(ctx context.Context) error {
	c := cron.New(
		cron.WithLocation(time.UTC),
		cron.WithLogger(newCronLogger(ctx)),
	)

	jobs := []Job{
		{
			Name:  "Symbiotic",
			Run:   newSymbioticRunner(),
			Spec:  "0 0 * * *",
			Spend: spend(tokens.WSTETH, 0.1),
		},
	}

	for _, j := range jobs {
		if _, err := c.AddJob(j.Spec, cronner(ctx, j)); err != nil {
			return errors.Wrap(err, "add job", "job", j.Name)
		}
	}

	c.Start()

	return nil
}
