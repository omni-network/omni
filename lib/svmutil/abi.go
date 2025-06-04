package svmutil

import (
	"crypto/rand"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/solvernet/fillhash"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/common"

	bin "github.com/gagliardetto/binary"
)

// FillHash returns the fill hash for the given parameters.
func FillHash(
	orderID [32]byte,
	srcChainID uint64,
	destChainID uint64,
	fillDeadline uint32,
	callTarget common.Address,
	callSelector [4]byte,
	callValue *big.Int,
	callParams []byte,
	expenseSpender common.Address,
	expenseToken common.Address,
	expenseAmount *big.Int,
) (common.Hash, error) {
	// Construct fill origin data
	fillData := bindings.SolverNetFillOriginData{
		SrcChainId:   srcChainID,
		DestChainId:  destChainID,
		FillDeadline: fillDeadline,
		Calls: []bindings.SolverNetCall{
			{
				Target:   callTarget,
				Selector: callSelector,
				Value:    callValue,
				Params:   callParams,
			},
		},
		Expenses: []bindings.SolverNetTokenExpense{
			{
				Spender: expenseSpender,
				Token:   expenseToken,
				Amount:  expenseAmount,
			},
		},
	}

	resp, err := fillhash.FillHash(orderID, fillData)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "encode fill data")
	}

	return resp, nil
}

// U128 returns the big int as a solana binary uint128.
func U128(i *big.Int) (bin.Uint128, error) {
	if bi.GT(i, umath.MaxUint128) {
		return bin.Uint128{}, errors.New("value too large")
	}

	// Explicitly don't use bin.NewUint128LittleEndian,
	// rather rely on defaultOrder (little endian), since
	// values then match unmarshalled solana client values.
	var resp bin.Uint128
	if err := resp.UnmarshalJSON([]byte(`"` + i.String() + `"`)); err != nil {
		return bin.Uint128{}, errors.Wrap(err, "unmarshal JSON [BUG]")
	}

	return resp, nil
}

// RandomU96 generates a random 96-bit unsigned integer of type bin.Uint128.
func RandomU96() bin.Uint128 {
	var b [12]byte
	_, _ = rand.Read(b[:])

	resp, err := U128(new(big.Int).SetBytes(b[:]))
	if err != nil {
		panic(err)
	}

	return resp
}
