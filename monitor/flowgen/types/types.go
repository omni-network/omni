package types

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/ethclient/ethbackend"
)

type Job interface {
	// Name is the friendly name of the job
	Name() string

	// Cadence is intrerval at which to run the job
	Cadence() time.Duration

	// Run runs the job exactly once
	Run(context.Context, ethbackend.Backends) error
}
