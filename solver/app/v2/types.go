package appv2

import (
	"encoding/binary"
	"strconv"

	"github.com/omni-network/omni/contracts/bindings"

	"github.com/ethereum/go-ethereum/common"
)

type OrderResolved = bindings.IERC7683ResolvedCrossChainOrder
type OrderState = bindings.ISolverNetInboxOrderState
type OrderHistory = []bindings.ISolverNetInboxStatusUpdate

// TODO: Replace solver/types::ReqID with this.
type OrderID [32]byte

type Order struct {
	ID         OrderID
	Resolved   OrderResolved
	Status     uint8
	AcceptedBy common.Address
	History    OrderHistory
}

// Uint64 returns the order ID as a BigEndian uint64 (monotonically incrementing number).
func (id OrderID) Uint64() uint64 {
	return binary.BigEndian.Uint64(id[32-8:])
}

// String returns the Uint64 representation of the order ID as a string.
func (id OrderID) String() string {
	return strconv.FormatUint(id.Uint64(), 10)
}
