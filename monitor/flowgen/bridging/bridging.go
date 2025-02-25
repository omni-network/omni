package bridging

import (
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

// NewJob instantiates the flow-specific job configuration.
func NewJob(
	network netconf.ID,
	srcChain,
	dstChain uint64,
	role eoa.Role,
	token common.Address,
	amount *big.Int,
) (*types.Job, error) {
	owner := eoa.MustAddress(network, role)
	data, err := orderData(owner, srcChain, dstChain, token, amount)
	if err != nil {
		return nil, errors.Wrap(err, "new job")
	}

	job := types.Job{
		Name:    "Bridging",
		Cadence: 1 * time.Minute,
		Network: network,

		SrcChain: srcChain,
		DstChain: dstChain,

		Owner: owner,

		OrderData: data,
	}

	return &job, nil
}

// OrderData returns the flow-specific order data.
func orderData(
	owner common.Address,
	srcChain, dstChain uint64,
	tokenAddr common.Address,
	amount *big.Int,
) (bindings.SolverNetOrderData, error) {
	token, ok := app.AllTokens().Find(srcChain, tokenAddr)
	if !ok {
		return bindings.SolverNetOrderData{}, errors.New("token not found")
	}
	deposit := app.Payment{Token: token, Amount: amount}

	expense, err := app.QuoteExpense(deposit.Token, deposit)
	if err != nil {
		return bindings.SolverNetOrderData{}, errors.Wrap(err, "quote expense")
	}

	orderData := bindings.SolverNetOrderData{
		Owner:       owner,
		DestChainId: dstChain,
		Deposit: solvernet.Deposit{
			Token:  deposit.Token.Address,
			Amount: deposit.Amount,
		},
		Expenses: solvernet.Expenses{},
		Calls: solvernet.Calls{
			{
				Target: owner,
				Value:  expense.Amount,
				Data:   nil,
			},
		}.ToBindings(),
	}

	return orderData, nil
}
