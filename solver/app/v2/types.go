package appv2

import (
	"encoding/binary"
	"strconv"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
)

type OrderResolved = bindings.IERC7683ResolvedCrossChainOrder
type OrderState = bindings.ISolverNetInboxOrderState
type OrderHistory = []bindings.ISolverNetInboxStatusUpdate
type FillOriginData = bindings.ISolverNetFillOriginData

// TODO: Replace solver/types::ReqID with this.
type OrderID [32]byte

type Order struct {
	ID                 OrderID
	FillInstruction    bindings.IERC7683FillInstruction
	FillOriginData     []byte
	DestinationSettler common.Address
	DestinationChainID uint64
	SourceChainID      uint64
	MaxSpent           []bindings.IERC7683Output
	MinReceived        []bindings.IERC7683Output

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

func newOrder(resolved OrderResolved, state OrderState, history OrderHistory) (Order, error) {
	if err := validateResolved(resolved); err != nil {
		return Order{}, errors.Wrap(err, "validate resolved")
	}

	o := Order{
		ID:                 resolved.OrderId,
		Status:             state.Status,
		AcceptedBy:         state.AcceptedBy,
		History:            history,
		FillInstruction:    resolved.FillInstructions[0],
		FillOriginData:     resolved.FillInstructions[0].OriginData,
		DestinationChainID: resolved.FillInstructions[0].DestinationChainId,
		DestinationSettler: toEthAddr(resolved.FillInstructions[0].DestinationSettler),
		SourceChainID:      resolved.OriginChainId.Uint64(),
		MaxSpent:           resolved.MaxSpent,
		MinReceived:        resolved.MinReceived,
	}

	return o, nil
}

func validateResolved(o OrderResolved) error {
	if o.OrderId == [32]byte{} {
		return errors.New("order ID is empty")
	}

	if o.OriginChainId == nil {
		return errors.New("origin chain ID is nil")
	}

	if o.FillInstructions == nil {
		return errors.New("fill instructions are nil")
	}

	if o.MaxSpent == nil {
		return errors.New("max spent is nil")
	}

	if len(o.FillInstructions) != 1 {
		return errors.New("expected exactly one fill instruction")
	}

	return nil
}

func (o Order) ParsedFillOriginData() (FillOriginData, error) {
	return solvernet.ParseFillOriginData(o.FillOriginData)
}
