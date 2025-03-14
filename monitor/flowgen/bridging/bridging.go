package bridging

import (
	"fmt"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/monitor/flowgen/types"
	"github.com/omni-network/omni/monitor/flowgen/util"
	"github.com/omni-network/omni/solver/app"

	"github.com/ethereum/go-ethereum/common"
)

// NewJob instantiates the job that bridges native ETH.
func newJob(
	networkID netconf.ID,
	srcChain,
	dstChain uint64,
	owner common.Address,
	amount *big.Int,
) (types.Job, error) {
	data, err := orderData(owner, srcChain, dstChain, amount)
	if err != nil {
		return types.Job{}, errors.Wrap(err, "new job")
	}

	cadence := 30 * time.Minute
	if networkID == netconf.Devnet {
		cadence = time.Second * 10
	}

	namer := netconf.ChainNamer(networkID)

	return types.Job{
		Name:      fmt.Sprintf("Bridging (%v->%v)", namer(srcChain), namer(dstChain)),
		Cadence:   cadence,
		NetworkID: networkID,

		SrcChain: srcChain,
		DstChain: dstChain,

		Owner: owner,

		OrderData: data,
	}, nil
}

// orderData returns the order data required to do the job.
func orderData(
	owner common.Address,
	srcChain, dstChain uint64,
	amount *big.Int,
) (bindings.SolverNetOrderData, error) {
	token, ok := app.AllTokens().Find(srcChain, common.Address{})
	if !ok {
		return bindings.SolverNetOrderData{}, errors.New("token not found")
	}

	expense := app.TokenAmt{Token: token, Amount: amount}

	depositWithFee, err := app.QuoteDeposit(expense.Token, expense)
	if err != nil {
		return bindings.SolverNetOrderData{}, errors.Wrap(err, "quote deposit")
	}

	orderData := bindings.SolverNetOrderData{
		Owner:       owner,
		DestChainId: dstChain,
		Deposit: solvernet.Deposit{
			Token:  depositWithFee.Token.Address,
			Amount: depositWithFee.Amount,
		},
		Expenses: []solvernet.Expense{},
		Calls: []bindings.SolverNetCall{
			{
				Target: owner,
				Value:  expense.Amount,
			},
		},
	}

	return orderData, nil
}

func Jobs(networkID netconf.ID, owner common.Address) ([]types.Job, error) {
	type balanced struct {
		From uint64
		To   uint64
	}

	b, ok := map[netconf.ID]balanced{
		netconf.Devnet:  {evmchain.IDMockL1, evmchain.IDMockL2},
		netconf.Staging: {evmchain.IDBaseSepolia, evmchain.IDOpSepolia},
		netconf.Omega:   {evmchain.IDOpSepolia, evmchain.IDArbSepolia},
	}[networkID]
	if !ok {
		return nil, nil
	}

	// Bridging of native ETH
	amount := big.NewInt(0).Mul(util.MilliEther, big.NewInt(20)) // 0.02 ETH

	job1, err := newJob(networkID, b.From, b.To, owner, amount)
	if err != nil {
		return nil, err
	}

	job2, err := newJob(networkID, b.To, b.From, owner, amount)
	if err != nil {
		return nil, err
	}

	return []types.Job{job1, job2}, nil
}
