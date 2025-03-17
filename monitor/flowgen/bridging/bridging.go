package bridging

import (
	"fmt"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/monitor/flowgen/types"
	"github.com/omni-network/omni/solver/app"

	"github.com/ethereum/go-ethereum/common"
)

// NewJob returns a job that bridges native tokens.
func newJob(
	networkID netconf.ID,
	srcChain, dstChain uint64,
	owner common.Address,
	amount *big.Int,
) (types.Job, error) {
	data, err := nativeOrderData(owner, srcChain, dstChain, amount)
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

	expense := app.TokenAmt{Token: dstToken, Amount: amount}

	depositWithFee, err := app.QuoteDeposit(srcToken, expense)
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

// Jobs bridges native ETH from one chain to another one.
func Jobs(networkID netconf.ID, owner common.Address) ([]types.Job, error) {
	type balanced struct {
		From uint64
		To   uint64
	}

	b, ok := map[netconf.ID]balanced{
		netconf.Devnet:  {evmchain.IDMockL1, evmchain.IDMockL2},
		netconf.Staging: {evmchain.IDBaseSepolia, evmchain.IDOpSepolia},
		netconf.Omega:   {evmchain.IDOpSepolia, evmchain.IDArbSepolia},
		netconf.Mainnet: {evmchain.IDOptimism, evmchain.IDArbitrumOne},
	}[networkID]
	if !ok {
		return nil, nil
	}

	// Bridging of native ETH
	amount := bi.Ether(0.02) // 0.02 ETH

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
