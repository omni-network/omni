package types

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
)

type Job struct {
	// Name is the friendly name of the job
	Name string

	// Cadence is intrerval at which to run the job
	Cadence time.Duration

	NetworkID netconf.ID

	SrcChainBackend *ethbackend.Backend

	OpenOrderFunc func(ctx context.Context) (solvernet.OrderID, bool, error)
}
