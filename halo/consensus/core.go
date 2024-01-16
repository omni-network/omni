package consensus

import (
	"github.com/omni-network/omni/halo/attest"
	"github.com/omni-network/omni/lib/engine"
	"github.com/omni-network/omni/lib/xchain"

	abci "github.com/cometbft/cometbft/abci/types"
)

// Core of the Halo consensus client with the following responsibilities:
// - Implements the server side of the ABCI++ interface, see abci.go.
// - Maintains the consensus chain state.
type Core struct {
	ethCl     engine.Client
	state     *State
	attestSvc attest.Service
}

// NewCore returns a new Core instance.
func NewCore(ethCl engine.Client, attestSvc attest.Service, state *State) *Core {
	return &Core{
		ethCl:     ethCl,
		attestSvc: attestSvc,
		state:     state,
	}
}

// ApprovedAggregates returns a copy of the latest state's approved aggregates.
// For testing purposes only.
func (c *Core) ApprovedAggregates() []xchain.AggAttestation {
	return c.state.ApprovedAggregates()
}

var _ abci.Application = (*Core)(nil)
