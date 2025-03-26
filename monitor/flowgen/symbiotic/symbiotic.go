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
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/monitor/flowgen/types"
	"github.com/omni-network/omni/monitor/flowgen/util"
	solver "github.com/omni-network/omni/solver/app"
	stokens "github.com/omni-network/omni/solver/tokens"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	_ "embed"
)

//go:embed symbiotic_abi.json
var abiJSON []byte

// newJob returns a symbiotic deposit job if a config for the given network exists, or it returns false.
func newJob(
	ctx context.Context,
	backends ethbackend.Backends,
	networkID netconf.ID,
	owner common.Address,
) (types.Job, bool, error) {
	addrs, err := contracts.GetAddresses(ctx, networkID)
	if err != nil {
		return types.Job{}, false, errors.New("contract addresses")
	}

	conf, ok := config[networkID]
	if !ok {
		return types.Job{}, false, nil
	}

	backend, err := backends.Backend(conf.srcChain)
	if err != nil {
		return types.Job{}, false, errors.Wrap(err, "src chain backend")
	}

	token := tokens.WSTETH

	srcChainTkn, ok := stokens.BySymbol(conf.srcChain, token.Symbol)
	if !ok {
		return types.Job{}, false, errors.Wrap(err, "src token not found")
	}

	dstChainTkn, ok := stokens.BySymbol(conf.dstChain, token.Symbol)
	if !ok {
		return types.Job{}, false, errors.Wrap(err, "dst token not found")
	}

	if err := util.ApproveToken(ctx, backend, srcChainTkn.Address, owner, addrs.SolverNetInbox); err != nil {
		return types.Job{}, false, errors.Wrap(err, "token approval")
	}

	namer := netconf.ChainNamer(networkID)
	cadence := 30 * time.Minute
	if networkID == netconf.Devnet {
		cadence = time.Second * 10
	}

	return types.Job{
		Name:      fmt.Sprintf("Symbiotic deposit (%v->%v)", namer(conf.srcChain), namer(conf.dstChain)),
		Cadence:   cadence,
		NetworkID: networkID,

		SrcChainBackend: backend,

		OpenOrderFunc: func(ctx context.Context) (types.Result, bool, error) {
			return openOrder(ctx, backends, networkID, owner, srcChainTkn, dstChainTkn, conf)
		},
	}, true, nil
}

// openOrder returns the order id if an order was opened successfully,
// it returns false if no order was opened or an error in case of an error.
func openOrder(
	ctx context.Context,
	backends ethbackend.Backends,
	networkID netconf.ID,
	owner common.Address,
	srcToken, dstToken stokens.Token,
	conf flowConfig,
) (types.Result, bool, error) {
	expense := solver.TokenAmt{Token: dstToken, Amount: conf.orderSize}

	depositWithFee, err := solver.QuoteDeposit(srcToken, solver.TokenAmt{Token: srcToken, Amount: conf.orderSize})
	if err != nil {
		return types.Result{}, false, errors.Wrap(err, "quote deposit")
	}

	abi, err := metaData.GetAbi()
	if err != nil {
		return types.Result{}, false, errors.Wrap(err, "get abi")
	}

	data, err := abi.Pack("deposit", owner, expense.Amount)
	if err != nil {
		return types.Result{}, false, errors.Wrap(err, "packing")
	}

	orderData := bindings.SolverNetOrderData{
		Owner:       owner,
		DestChainId: conf.dstChain,
		Deposit: solvernet.Deposit{
			Token:  depositWithFee.Token.Address,
			Amount: depositWithFee.Amount,
		},
		Expenses: []solvernet.Expense{{
			Token:   expense.Token.Address,
			Amount:  expense.Amount,
			Spender: conf.vaultAddr,
		}},
		Calls: []bindings.SolverNetCall{
			solvernet.Call{
				Target: conf.vaultAddr,
				Data:   data,
				Value:  new(big.Int),
			}.ToBinding(),
		},
	}

	orderID, err := solvernet.OpenOrder(ctx, networkID, conf.srcChain, backends, owner, orderData)
	if err != nil {
		return types.Result{}, false, errors.Wrap(err, "open order")
	}

	return types.Result{OrderID: orderID, Expense: expense}, true, nil
}

var metaData = &bind.MetaData{
	ABI: string(abiJSON),
}

// Jobs creates the following jobs:
// - deposit wstETH from the source to the destination chain.
func Jobs(ctx context.Context, backends ethbackend.Backends, networkID netconf.ID, owner common.Address) ([]types.Job, error) {
	var jobs []types.Job
	job, ok, err := newJob(
		ctx,
		backends,
		networkID,
		owner,
	)
	if err != nil {
		return jobs, errors.Wrap(err, "symbiotic job")
	}
	if ok {
		jobs = append(jobs, job)
	}

	return jobs, nil
}
