package flowgen

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/tokens"
)

type Spend map[tokens.Token]*big.Int

type RunFunc func(context.Context, Spend) error

type Job struct {
	// Name is the friendly name of the job
	Name string

	// Run is the function to run
	Run RunFunc

	// Cadence is intrerval at which to run the job
	Cadence time.Duration

	// Spend is spend this job uses on one run (exluding gas)
	Spend Spend
}
