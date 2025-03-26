package types

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
	solver "github.com/omni-network/omni/solver/app"
)

type Result struct {
	OrderID solvernet.OrderID
	Expense solver.TokenAmt
}

type Job struct {
	// Name is the friendly name of the job
	Name string

	// Cadence is intrerval at which to run the job
	Cadence time.Duration

	NetworkID netconf.ID

	SrcChainBackend *ethbackend.Backend

	// OpenOrderFunc opens an order and returns the result, or false if the order wasn't opened, or an error.
	OpenOrderFunc func(ctx context.Context) (Result, bool, error)
}
