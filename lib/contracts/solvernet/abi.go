package solvernet

import (
	"crypto/sha256"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const (
	typUint256    = "uint256"
	typUint96     = "uint96"
	typUint64     = "uint64"
	typUint32     = "uint32"
	typBytes32    = "bytes32"
	typBytes4     = "bytes4"
	typBytes      = "bytes"
	typAddress    = "address"
	typTuple      = "tuple"
	typTupleArray = "tuple[]"
)

// FillHash returns the fill hash for the given order and fill data.
func FillHash(
	orderID OrderID,
	fillData bindings.SolverNetFillOriginData,
) (common.Hash, error) {
	encoded, err := encodeFillData(orderID, fillData)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "encode fill data")
	}

	return sha256.Sum256(encoded), nil
}

// This is equivalent to: abi.encode(orderId, fillOriginData);.
func encodeFillData(orderID OrderID, fillData bindings.SolverNetFillOriginData) ([]byte, error) {
	for _, expense := range fillData.Expenses {
		if bi.GT(expense.Amount, umath.MaxUint96) {
			return nil, errors.New("expense amount too large")
		}
	}

	encoded, err := fillHashArgs.Pack(orderID, fillData)
	if err != nil {
		return nil, errors.Wrap(err, "pack fill data")
	}

	return encoded, nil
}

var fillHashArgs = abi.Arguments{
	mustArg(typBytes32, nil),
	mustArg(typTuple, []abi.ArgumentMarshaling{
		{Name: "SrcChainId", Type: typUint64},
		{Name: "DestChainId", Type: typUint64},
		{Name: "FillDeadline", Type: typUint32},
		{Name: "Calls", Type: typTupleArray,
			Components: []abi.ArgumentMarshaling{
				{Name: "Target", Type: typAddress},
				{Name: "Selector", Type: typBytes4},
				{Name: "Value", Type: typUint256},
				{Name: "Params", Type: typBytes},
			},
		},
		{Name: "Expenses", Type: typTupleArray,
			Components: []abi.ArgumentMarshaling{
				{Name: "Spender", Type: typAddress},
				{Name: "Token", Type: typAddress},
				{Name: "Amount", Type: typUint96},
			},
		},
	}),
}

// mustArg returns an ABI argument type for a simple (non-nested) type.
func mustArg(typ string, components []abi.ArgumentMarshaling) abi.Argument {
	t, err := abi.NewType(typ, "", components)
	if err != nil {
		panic(err)
	}

	return abi.Argument{Type: t}
}
