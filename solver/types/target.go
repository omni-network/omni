package types

import (
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
)

// Target is the interface for a target contract the solver can interact with.
type Target interface {
	// ChainID returns the chain ID of the target contract.
	ChainID() uint64

	// Address returns the address of the target contract.
	Address() common.Address

	// IsAllowedCall returns true if the call is allowed.
	IsAllowedCall(call bindings.SolveCall) bool

	// TokenPrereqs returns the token prerequisites required for the call.
	TokenPrereqs(call bindings.SolveCall) ([]bindings.SolveTokenPrereq, error)

	// Verify returns an error if the call should not be fulfilled.
	Verify(srcChainID uint64, call bindings.SolveCall, deposits []bindings.SolveDeposit) error
}

type Targets []Target

func (t Targets) ForCall(call bindings.SolveCall) (Target, error) {
	var match Target
	var matched bool

	for _, target := range t {
		if target.IsAllowedCall(call) {
			if matched {
				return nil, errors.New("multiple targets found")
			}

			match = target
			matched = true
		}
	}

	if !matched {
		return nil, errors.New("no target found")
	}

	return match, nil
}
