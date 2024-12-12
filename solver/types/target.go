package types

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"

	"github.com/ethereum/go-ethereum/common"
)

// Target is the interface for a target contract the solver can interact with.
type Target interface {
	// Name of the target
	Name() string

	// ChainID returns the chain ID of the target contract.
	ChainID() uint64

	// Address returns the address of the target contract.
	Address() common.Address

	// TokenPrereqs returns the token prerequisites required for the call.
	TokenPrereqs(call bindings.SolveCall) ([]bindings.SolveTokenPrereq, error)

	// Verify returns an error if the call should not be fulfilled.
	// TODO(corver): Return reject reason.
	Verify(srcChainID uint64, call bindings.SolveCall, deposits []bindings.SolveDeposit) error

	// DebugCall logs the call for debugging purposes.
	DebugCall(ctx context.Context, call bindings.SolveCall) error
}
