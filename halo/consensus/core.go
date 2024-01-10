package consensus

import (
	"github.com/omni-network/omni/halo/attest"
	"github.com/omni-network/omni/lib/engine"

	abci "github.com/cometbft/cometbft/abci/types"
)

// Core of the Halo consensus client with the following responsibilities:
// - Implements the server side of the ABCI++ interface, see abci.go.
// - Maintains the consensus chain state.
type Core struct {
	ethCl     engine.Client
	state     state
	attestSvc attest.Service
}

// NewCore returns a new Core instance.
func NewCore(ethCl engine.Client, attestSvc attest.Service) *Core {
	return &Core{
		ethCl:     ethCl,
		attestSvc: attestSvc,
	}
}

var _ abci.Application = (*Core)(nil)
