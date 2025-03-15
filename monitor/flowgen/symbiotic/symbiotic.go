package symbiotic

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/monitor/flowgen/types"
	"github.com/omni-network/omni/monitor/flowgen/util"
	solver "github.com/omni-network/omni/solver/app"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	_ "embed"
)

//go:embed symbiotic_abi.json
var abiJSON []byte

func newJob(
	ctx context.Context,
	backends ethbackend.Backends,
	networkID netconf.ID,
	owner common.Address,
	amount *big.Int,
) (types.Job, error) {
	addrs, err := contracts.GetAddresses(ctx, networkID)
	if err != nil {
		return types.Job{}, errors.New("contract addresses")
	}

	conf, ok := config[networkID]
	if !ok {
		return types.Job{}, errors.New("flow config missing")
	}

	backend, err := backends.Backend(conf.srcChain)
	if err != nil {
		return types.Job{}, errors.Wrap(err, "get backend")
	}

	srcChainTkn, ok := solver.AllTokens().FindBySymbol(conf.srcChain, conf.depositToken.Symbol)
	if !ok {
		return types.Job{}, errors.Wrap(err, "src token not found")
	}

	dstChainTkn, ok := solver.AllTokens().FindBySymbol(conf.dstChain, conf.expenseToken.Symbol)
	if !ok {
		return types.Job{}, errors.Wrap(err, "dst token not found")
	}

	if err := util.ApproveToken(ctx, backend, srcChainTkn.Address, owner, addrs.SolverNetInbox); err != nil {
		return types.Job{}, errors.Wrap(err, "token approval")
	}

	data, err := orderData(owner, conf.dstChain, srcChainTkn, dstChainTkn, conf.vaultAddr, amount)
	if err != nil {
		return types.Job{}, errors.Wrap(err, "new job")
	}

	namer := netconf.ChainNamer(networkID)

	return types.Job{
		Name:      fmt.Sprintf("Symbiotic deposit (%v->%v)", namer(conf.srcChain), namer(conf.dstChain)),
		Cadence:   30 * time.Minute,
		NetworkID: networkID,

		SrcChain: conf.srcChain,
		DstChain: conf.dstChain,

		Owner: owner,

		OrderData: data,
	}, nil
}

// orderData returns the order data required to do the job.
func orderData(
	owner common.Address,
	dstChain uint64,
	srcToken, dstToken solver.Token,
	vaultAddr common.Address,
	amount *big.Int,
) (bindings.SolverNetOrderData, error) {
	expense := solver.TokenAmt{Token: dstToken, Amount: amount}

	depositWithFee, err := solver.QuoteDeposit(srcToken, solver.TokenAmt{Token: srcToken, Amount: amount})
	if err != nil {
		return bindings.SolverNetOrderData{}, errors.Wrap(err, "quote deposit")
	}

	abi, err := metaData.GetAbi()
	if err != nil {
		return bindings.SolverNetOrderData{}, errors.Wrap(err, "get abi")
	}

	data, err := abi.Pack("deposit", owner, expense.Amount)
	if err != nil {
		return bindings.SolverNetOrderData{}, errors.Wrap(err, "packing")
	}

	orderData := bindings.SolverNetOrderData{
		Owner:       owner,
		DestChainId: dstChain,
		Deposit: solvernet.Deposit{
			Token:  depositWithFee.Token.Address,
			Amount: depositWithFee.Amount,
		},
		Expenses: []solvernet.Expense{{
			Token:   expense.Token.Address,
			Amount:  expense.Amount,
			Spender: vaultAddr,
		}},
		Calls: []bindings.SolverNetCall{
			solvernet.Call{
				Target: vaultAddr,
				Data:   data,
				Value:  new(big.Int),
			}.ToBinding(),
		},
	}

	return orderData, nil
}

var metaData = &bind.MetaData{
	ABI: string(abiJSON),
}

func Jobs(ctx context.Context, backends ethbackend.Backends, networkID netconf.ID, owner common.Address) ([]types.Job, error) {
	var jobs []types.Job
	deposit := big.NewInt(0).Mul(util.MilliEther, big.NewInt(20)) // 0.02 ETH
	job, err := newJob(
		ctx,
		backends,
		networkID,
		owner,
		deposit,
	)
	if err != nil {
		return jobs, errors.Wrap(err, "symbiotic job")
	}
	jobs = append(jobs, job)

	return jobs, nil
}
