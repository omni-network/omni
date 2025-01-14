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

// FillInstruction returns the order fill instruction. SolverNet v1 ERC7683 orders only have one fill instruction.
func (o Order) FillInstruction() bindings.IERC7683FillInstruction {
	return o.Resolved.FillInstructions[0]
}

// DestinationSettler returns the destination settler address.
func (o Order) DestinationSettler() (common.Address, error) {
	settler := o.FillInstruction().DestinationSettler
	addr, err := cast.EthAddress(settler[:])
	if err != nil {
		return common.Address{}, errors.Wrap(err, "cast to eth addr")
	}

	return addr, nil
}

// LogAttrs returns a map of order attributes for logging. It tries to extract fil origin data with known format.
func (o Order) LogAttrs(ctx context.Context) []any {
	attrs := []any{
		"id", o.ID.String(),
		"status", o.Status,
		"accepted_by", o.AcceptedBy.Hex(),
		"src_chain_id", o.SourceChainID(),
		"dst_chain_id", o.DestinationChainID(),
	}

	fill, err := parseFillOriginData(o.FillInstruction().OriginData)
	if err != nil {
		log.Warn(ctx, "Failed to parse fill origin data", err, attrs...)

		return append(attrs,
			"call_target", unknown,
			"call_value", unknown,
			"call_data", unknown,
			"call_expenses", unknown,
		)
	}

	return append(attrs,
		"call_target", hexutil.Encode(fill.Call.Target[:]),
		"call_value", fill.Call.Value.Uint64(),
		"call_data", hexutil.Encode(fill.Call.Data),
		"call_expenses", len(fill.Call.Expenses),
	)
}

// SourceChainID returns the order's source chain ID.
func (o Order) SourceChainID() uint64 {
	return o.Resolved.OriginChainId.Uint64()
}

// DestinationChainID returns the order's destination chain ID.
func (o Order) DestinationChainID() uint64 {
	return o.FillInstruction().DestinationChainId
}

// FillOriginData returns the order's fill origin data, as packed bytes.
func (o Order) FillOriginData() []byte {
	return o.FillInstruction().OriginData
}
