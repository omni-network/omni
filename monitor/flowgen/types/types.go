package types

import (
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

type Job struct {
	// Name is the friendly name of the job
	Name string

	// Cadence is intrerval at which to run the job
	Cadence time.Duration

	Network netconf.ID

	SrcChain uint64
	DstChain uint64

	Owner common.Address

	OrderData bindings.SolverNetOrderData
}
