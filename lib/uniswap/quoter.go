package uniswap

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/lib/bi"
	// "github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common/hexutil"
)

// CallQuoteExactInput calls the quoteExactInput method on the QuoterV2 contract.
// This is needed because quoteExactInput is a mutator, not a view function,
// so we need to use contract.Call instead of the generated QuoteExactInput method.
func (q *UniQuoterV2) CallQuoteExactInput(ctx context.Context, path []byte, amountIn *big.Int) (*big.Int, error) {
	var result []any
	err := q.UniQuoterV2Caller.contract.Call(&bind.CallOpts{Context: ctx}, &result, "quoteExactInput", path, amountIn)
	if err != nil {
		return nil, errors.Wrap(err, "quote exact input")
	}

	if len(result) == 0 {
		return nil, errors.New("empty result")
	}

	amountOut, ok := result[0].(*big.Int)
	if !ok {
		return nil, errors.New("invalid type")
	}

	if bi.IsZero(amountOut) {
		return nil, errors.New("zero out")
	}

	return amountOut, nil
}

// CallQuoteExactOutput calls the quoteExactOutput method on the QuoterV2 contract.
func (q *UniQuoterV2) CallQuoteExactOutput(ctx context.Context, path []byte, amountOut *big.Int) (*big.Int, error) {
	var result []any
	err := q.UniQuoterV2Caller.contract.Call(&bind.CallOpts{Context: ctx}, &result, "quoteExactOutput", path, amountOut)
	if err != nil {
		return nil, errors.Wrap(err, "quote exact output")
	}

	if len(result) == 0 {
		return nil, errors.New("empty result")
	}

	amountIn, ok := result[0].(*big.Int)
	if !ok {
		return nil, errors.New("invalid type")
	}

	if bi.IsZero(amountIn) {
		return nil, errors.New("zero in")
	}

	return amountIn, nil
}
