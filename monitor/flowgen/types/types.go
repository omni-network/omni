package types

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
	solver "github.com/omni-network/omni/solver/app"
)

type Receipt struct {
	OrderID solvernet.OrderID
	Expense solver.TokenAmt
	Success bool
}

type Job struct {
	// Name is the friendly name of the job
	Name string

	// Cadence is intrerval at which to run the job
	Cadence time.Duration

	NetworkID netconf.ID

	SrcChainBackend *ethbackend.Backend

	OpenOrderFunc func(ctx context.Context) (Receipt, error)
}
