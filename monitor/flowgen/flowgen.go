package flowgen

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/tokens"
)

func Start(ctx context.Context) error {
	jobs := []Job{
		{
			Name:    "Symbiotic",
			Run:     newSymbioticRunner(),
			Cadence: 1 * time.Hour,
			Spend:   spend(tokens.WSTETH, ether1),
		},
	}

	_ = jobs
	_ = ctx

	// TODO

	return nil
}
