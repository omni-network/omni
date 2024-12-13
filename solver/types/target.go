package types

import (
	"context"
	"encoding/binary"
	"strconv"

	"github.com/omni-network/omni/contracts/bindings"

	"github.com/ethereum/go-ethereum/common"
)

// ReqID is a inbox request ID.
type ReqID [32]byte

// Uint64 returns the req ID as a BigEndian uint64 (monotonically incrementing number).
func (r ReqID) Uint64() uint64 {
	return binary.BigEndian.Uint64(r[32-8:])
}

// String returns the Uint64 representation of the req ID as a string.
func (r ReqID) String() string {
	return strconv.FormatUint(r.Uint64(), 10)
}

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

	// LogCall logs the call for debugging purposes.
	LogCall(ctx context.Context, call bindings.SolveCall) error

	// LogMetadata logs target metadata.
	LogMetadata(ctx context.Context)
}
