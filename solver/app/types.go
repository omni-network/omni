package app

import (
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"
)

type (
	OrderID        = solvernet.OrderID
	OrderResolved  = solvernet.OrderResolved
	OrderState     = solvernet.OrderState
	FillOriginData = solvernet.FillOriginData
)

type Order struct {
	ID                 OrderID
	FillInstruction    bindings.IERC7683FillInstruction
	FillOriginData     []byte
	DestinationSettler common.Address
	DestinationChainID uint64
	SourceChainID      uint64
	MaxSpent           []bindings.IERC7683Output
	MinReceived        []bindings.IERC7683Output

	Status    solvernet.OrderStatus
	UpdatedBy common.Address
}

func newOrder(resolved OrderResolved, state OrderState) (Order, error) {
	if err := validateResolved(resolved); err != nil {
		return Order{}, errors.Wrap(err, "validate resolved")
	}

	o := Order{
		ID:                 resolved.OrderId,
		Status:             solvernet.OrderStatus(state.Status),
		UpdatedBy:          state.UpdatedBy,
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
	resp, err := solvernet.ParseFillOriginData(o.FillOriginData)
	if err != nil {
		return FillOriginData{}, errors.Wrap(err, "parse fill origin data")
	} else if len(resp.Calls) == 0 {
		return FillOriginData{}, errors.New("no calls in fill origin data")
	}

	return resp, nil
}
