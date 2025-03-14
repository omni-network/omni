package bridging

import (
	"fmt"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/monitor/flowgen/types"
	"github.com/omni-network/omni/solver/app"

	"github.com/ethereum/go-ethereum/common"
)

// NewJob returns a job that bridges native tokens.
func NewJob(
	network netconf.ID,
	srcChain,
	dstChain uint64,
	role eoa.Role,
	amount *big.Int,
) (types.Job, error) {
	owner := eoa.MustAddress(network, role)
	data, err := nativeOrderData(owner, srcChain, dstChain, amount)
	if err != nil {
		return types.Job{}, errors.Wrap(err, "new job")
	}

	cadence := 30 * time.Minute
	if network == netconf.Devnet {
		cadence = time.Second * 10
	}

	namer := netconf.ChainNamer(network)

	return types.Job{
		Name:    fmt.Sprintf("Bridging (%v->%v)", namer(srcChain), namer(dstChain)),
		Cadence: cadence,
		Network: network,

		SrcChain: srcChain,
		DstChain: dstChain,

		Owner: owner,

		OrderData: data,
	}, nil
}

// nativeOrderData returns the order data required bridge native token amount from source to destination chain.
func nativeOrderData(
	owner common.Address,
	srcChain, dstChain uint64,
	amount *big.Int,
) (bindings.SolverNetOrderData, error) {
	srcToken, ok := app.AllTokens().Find(srcChain, app.NativeAddr)
	if !ok {
		return bindings.SolverNetOrderData{}, errors.New("src token not found")
	}
	dstToken, ok := app.AllTokens().Find(dstChain, app.NativeAddr)
	if !ok {
		return bindings.SolverNetOrderData{}, errors.New("dst token not found")
	}

	// Tokens that will be deposited to the user on the destination chain.
	expense := app.TokenAmt{Token: dstToken, Amount: amount}

	depositWithFee, err := app.QuoteDeposit(srcToken, expense)
	if err != nil {
		return bindings.SolverNetOrderData{}, errors.Wrap(err, "quote expense")
	}

	orderData := bindings.SolverNetOrderData{
		Owner:       owner,
		DestChainId: dstChain,
		Deposit: solvernet.Deposit{
			Token:  depositWithFee.Token.Address,
			Amount: depositWithFee.Amount,
		},
		Expenses: []solvernet.Expense{}, // Explicit expense not required for native transfer calls.
		Calls: []bindings.SolverNetCall{
			{
				Target: owner,
				Value:  expense.Amount,
			},
		},
	}

	return orderData, nil
}
