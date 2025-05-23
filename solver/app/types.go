package app

import (
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/tokens"

	"github.com/ethereum/go-ethereum/common"
)

type (
	OrderID        = solvernet.OrderID
	OrderResolved  = solvernet.OrderResolved
	OrderState     = solvernet.OrderState
	FillOriginData = solvernet.FillOriginData
)

// TokenAmt represents a token and an amount.
// It differs from types.AddrAmt in that it contains fully resolved token type, not just the token address.
type TokenAmt struct {
	Token  tokens.Token
	Amount *big.Int
}

// String returns a string representation of the token amount
// in primary units with the token symbol appended.
func (a TokenAmt) String() string {
	return a.Token.FormatAmt(a.Amount)
}

type Order struct {
	ID            OrderID
	Offset        uint64
	SourceChainID uint64
	Status        solvernet.OrderStatus
	UpdatedBy     common.Address

	pendingData PendingData
	filledData  FilledData
}

func (o Order) PendingData() (PendingData, error) {
	if o.Status != solvernet.StatusPending {
		return PendingData{}, errors.New("order is not pending")
	}

	return o.pendingData, nil
}

func (o Order) MinReceived() ([]bindings.IERC7683Output, error) {
	if o.Status == solvernet.StatusPending {
		return o.pendingData.MinReceived, nil
	} else if o.Status == solvernet.StatusFilled {
		return o.filledData.MinReceived, nil
	}

	return nil, errors.New("order is not pending or filled")
}

// PendingData contains order data that is only available for pending orders.
type PendingData struct {
	MinReceived        []bindings.IERC7683Output
	DestinationSettler common.Address
	DestinationChainID uint64
	FillOriginData     []byte
	MaxSpent           []bindings.IERC7683Output
}

// FilledData contains order data that is only available for filled orders.
type FilledData struct {
	MinReceived []bindings.IERC7683Output
}

func newOrder(resolved OrderResolved, state OrderState, offset *big.Int) (Order, error) {
	if err := validateResolved(resolved); err != nil {
		return Order{}, errors.Wrap(err, "validate resolved")
	}

	settler, err := toUniAddr(
		resolved.FillInstructions[0].DestinationChainId,
		resolved.FillInstructions[0].DestinationSettler,
	)
	if err != nil {
		return Order{}, errors.Wrap(err, "settler")
	} else if settler.IsEVM() {
		return Order{}, errors.New("only evm settler supported", "settler", settler)
	}

	return Order{
		ID:            resolved.OrderId,
		Offset:        offset.Uint64(),
		Status:        solvernet.OrderStatus(state.Status),
		UpdatedBy:     state.UpdatedBy,
		SourceChainID: resolved.OriginChainId.Uint64(),
		filledData: FilledData{
			MinReceived: resolved.MinReceived,
		},
		pendingData: PendingData{
			MinReceived:        resolved.MinReceived,
			FillOriginData:     resolved.FillInstructions[0].OriginData,
			DestinationChainID: resolved.FillInstructions[0].DestinationChainId,
			DestinationSettler: settler.EVM(),
			MaxSpent:           resolved.MaxSpent,
		},
	}, nil
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

func (d PendingData) ParsedFillOriginData() (FillOriginData, error) {
	resp, err := solvernet.ParseFillOriginData(d.FillOriginData)
	if err != nil {
		return FillOriginData{}, errors.Wrap(err, "parse fill origin data")
	} else if len(resp.Calls) == 0 {
		return FillOriginData{}, errors.New("no calls in fill origin data")
	}

	return resp, nil
}
