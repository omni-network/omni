package consensus

import (
	"github.com/omni-network/omni/halo/attest"

	abci "github.com/cometbft/cometbft/api/cometbft/abci/v1"
)

// Core of the Halo consensus client with the following responsibilities:
// - Implements the server side of the ABCI++ interface, see abci.go.
// - Maintains the consensus chain state.
type Core struct {
	state     state
	attestSvc attest.Service
}

var _ abci.ABCIServiceServer = (*Core)(nil)
