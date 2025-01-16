package appv2

import (
	"context"
	"encoding/binary"
	"strconv"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

	fillInstr := resolved.FillInstructions[0]
	o := Order{
		FillInstruction:    fillInstr,
		SourceChainID:      resolved.OriginChainId.Uint64(),
		DestinationChainID: fillInstr.DestinationChainId,
		Status:             state.Status,
		AcceptedBy:         state.AcceptedBy,
		History:            history,
	}

	settlerBz := fillInstr.DestinationSettler
	settler, err := cast.EthAddress(settlerBz[:])
	if err != nil {
		return Order{}, errors.Wrap(err, "cast destination settler")
	}
	o.DestinationSettler = settler

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

// LogAttrs returns a map of order attributes for logging. It tries to extract fil origin data with known format.
func (o Order) LogAttrs(ctx context.Context) []any {
	attrs := []any{
		"id", o.ID.String(),
		"status", o.Status,
		"src_chain_id", o.SourceChainID,
		"dst_chain_id", o.DestinationChainID,
	}

	fill, err := parseFillOriginData(o.FillOriginData)
	if err != nil {
		log.Warn(ctx, "Failed to parse fill origin data", err, attrs...)

		return append(attrs,
			"call_target", unknown,
			"call_value", unknown,
			"call_data", unknown,
		)
	}

	return append(attrs,
		"call_target", hexutil.Encode(fill.Call.Target[:]),
		"call_value", fill.Call.Value.Uint64(),
		"call_data", hexutil.Encode(fill.Call.Data),
	)
}

func (o Order) ParsedFillOriginData() (FillOriginData, error) {
	return parseFillOriginData(o.FillOriginData)
}
