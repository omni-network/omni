package bridging

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/monitor/flowgen/types"
	"github.com/omni-network/omni/solver/app"

	"github.com/ethereum/go-ethereum/common"
)

// NewJob returns a job that bridges native tokens.
func newJob(networkID netconf.ID, backends ethbackend.Backends, conf flowConfig, owner common.Address) (types.Job, error) {
	cadence := 30 * time.Minute
	if networkID == netconf.Devnet {
		cadence = time.Second * 10
	}

	backend, err := backends.Backend(conf.srcChain)
	if err != nil {
		return types.Job{}, errors.Wrap(err, "src chain backend")
	}

	namer := netconf.ChainNamer(networkID)

	return types.Job{
		Name:      fmt.Sprintf("Bridging (%v->%v)", namer(conf.srcChain), namer(conf.dstChain)),
		Cadence:   cadence,
		NetworkID: networkID,

		SrcChainBackend: backend,

		OpenOrderFunc: func(ctx context.Context) (solvernet.OrderID, bool, error) {
			return openOrder(ctx, backends, backend, networkID, owner, conf)
		},
	}, nil
}

// openOrder returns the order id if an order was opened successfully,
// it returns false if no order was opened or an error in case of an error.
func openOrder(
	ctx context.Context,
	backends ethbackend.Backends,
	backend *ethbackend.Backend,
	networkID netconf.ID,
	owner common.Address,
	conf flowConfig,
) (solvernet.OrderID, bool, error) {
	srcToken, ok := app.AllTokens().Find(conf.srcChain, app.NativeAddr)
	if !ok {
		return solvernet.OrderID{}, false, errors.New("src token not found")
	}
	dstToken, ok := app.AllTokens().Find(conf.dstChain, app.NativeAddr)
	if !ok {
		return solvernet.OrderID{}, false, errors.New("dst token not found")
	}

	orderSize, ok, err := estimateOrderSize(ctx, networkID, backend.Client, owner, conf, srcToken.Token)
	if err != nil {
		return solvernet.OrderID{}, false, errors.Wrap(err, "estimate order size")
	}
	if !ok {
		return solvernet.OrderID{}, false, nil
	}

	expense := app.TokenAmt{Token: dstToken, Amount: orderSize}

	depositWithFee, err := app.QuoteDeposit(srcToken, expense)
	if err != nil {
		return solvernet.OrderID{}, false, errors.Wrap(err, "quote deposit")
	}

	orderData := bindings.SolverNetOrderData{
		Owner:       owner,
		DestChainId: conf.dstChain,
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

	orderID, err := solvernet.OpenOrder(ctx, networkID, conf.srcChain, backends, owner, orderData)
	if err != nil {
		return solvernet.OrderID{}, false, errors.Wrap(err, "open order")
	}

	return orderID, true, nil
}

// Jobs returns two jobs bridging native ETH from one chain to another one and back.
func Jobs(networkID netconf.ID, backends ethbackend.Backends, owner common.Address) ([]types.Job, error) {
	conf, ok := config[networkID]
	if !ok {
		return nil, nil
	}

	job1, err := newJob(networkID, backends, conf, owner)
	if err != nil {
		return nil, err
	}

	conf2 := conf
	conf2.srcChain, conf2.dstChain = conf2.dstChain, conf2.srcChain
	job2, err := newJob(networkID, backends, conf2, owner)
	if err != nil {
		return nil, err
	}

	return []types.Job{job1, job2}, nil
}

// estimateOrderSize checks the current balance of the flowgen EOA and returns
// the maximal possible order size or false if the minimal balance is reached.
func estimateOrderSize(
	ctx context.Context,
	networkID netconf.ID,
	client ethclient.Client,
	owner common.Address,
	conf flowConfig,
	srcToken tokens.Token,
) (*big.Int, bool, error) {
	balance, err := client.BalanceAt(ctx, owner, nil)
	if err != nil {
		return nil, false, errors.Wrap(err, "balance at")
	}

	thresholds, ok := eoa.GetFundThresholds(srcToken, networkID, eoa.RoleFlowgen)
	if !ok {
		// Skip accounts without thresholds
		return nil, false, errors.New("no thresholds found", "role", eoa.RoleFlowgen)
	}

	orderSize := bi.Sub(balance, thresholds.MinBalance())

	// if order size is too small, do nothing
	if bi.LT(orderSize, conf.minOrderSize) {
		return nil, false, nil
	}

	// cap the order if necessary
	if bi.GT(orderSize, conf.maxOrderSize) {
		orderSize = conf.maxOrderSize
	}

	log.Debug(ctx, "Flowgen: order size estimated", "balance", balance, "min_threshold", thresholds.MinBalance(), "order", orderSize)

	return orderSize, true, nil
}
