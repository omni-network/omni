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
	"github.com/omni-network/omni/lib/forkjoin"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/monitor/flowgen/types"
	solver "github.com/omni-network/omni/solver/app"
	sclient "github.com/omni-network/omni/solver/client"
	stypes "github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"
)

// parallel is the number of parallel orders to open.
const parallelDev = 2
const parallel = 64

// Jobs returns two jobs bridging native ETH from one chain to another one and back.
func Jobs(
	networkID netconf.ID,
	backends ethbackend.Backends,
	scl sclient.Client,
) []types.Job {
	confs, ok := config[networkID]
	if !ok {
		return nil
	}

	var jobs []types.Job
	for _, conf := range confs {
		jobs = append(jobs,
			newJob(networkID, backends, conf, scl),
			newJob(networkID, backends, conf.Flip(), scl), // Second job using flipped config
		)
	}

	return jobs
}

// NewJob returns a job that bridges native tokens.
func newJob(
	networkID netconf.ID,
	backends ethbackend.Backends,
	conf flowConfig,
	scl sclient.Client,
) types.Job {
	cadence := 25 * time.Minute
	if networkID == netconf.Devnet {
		cadence = time.Second * 20
	}

	namer := netconf.ChainNamer(networkID)

	return types.Job{
		Name:       fmt.Sprintf("Bridging (%v->%v)", namer(conf.srcChain), namer(conf.dstChain)),
		Cadence:    cadence,
		SrcChainID: conf.srcChain,
		OpenOrdersFunc: func(ctx context.Context) ([]types.Result, error) {
			return openOrders(ctx, backends, networkID, conf, scl)
		},
	}
}

func openOrders(
	ctx context.Context,
	backends ethbackend.Backends,
	network netconf.ID,
	conf flowConfig,
	scl sclient.Client,
) ([]types.Result, error) {
	flowgenAddr := eoa.MustAddress(network, eoa.RoleFlowgen)

	srcToken, ok := tokens.Native(conf.srcChain)
	if !ok {
		return nil, errors.New("src token not found")
	}

	dstToken, ok := tokens.Native(conf.dstChain)
	if !ok {
		return nil, errors.New("dst token not found")
	}

	srcBackend, err := backends.Backend(conf.srcChain)
	if err != nil {
		return nil, err
	}
	destBackend, err := backends.Backend(dstToken.ChainID)
	if err != nil {
		return nil, err
	}

	price, err := swapPrice(ctx, scl, srcToken, dstToken)
	if err != nil {
		return nil, errors.Wrap(err, "price")
	}

	// Use all available flowgen balance (without dropping below minimums)
	totalAmount, err := availableBalance(ctx, network, srcBackend, eoa.RoleFlowgen, srcToken)
	if err != nil {
		return nil, errors.Wrap(err, "get source available")
	}

	// Limit to available solver balance on destination.
	if solverDstAvail, err := availableBalance(ctx, network, destBackend, eoa.RoleSolver, dstToken); err != nil {
		return nil, errors.Wrap(err, "get dest available")
	} else if solverSrcAvail := price.ToDeposit(solverDstAvail); bi.GT(totalAmount, solverSrcAvail) {
		totalAmount = solverSrcAvail
	}

	expenseBound, ok := solver.GetSpendBounds(dstToken)
	if !ok {
		return nil, errors.New("no expense bounds found", "token", dstToken)
	}
	depositBounds := expenseBound.DepositBounds(price)

	p := parallel
	if network == netconf.Devnet {
		p = parallelDev
	}

	var orderDatas []bindings.SolverNetOrderData
	for _, depositAmount := range splitOrderAmounts(depositBounds, totalAmount, p) {
		orderData, err := nativeOrderData(ctx, scl, flowgenAddr, srcToken, dstToken, depositAmount)
		if err != nil {
			return nil, err
		}

		orderDatas = append(orderDatas, orderData)
	}

	work := func(ctx context.Context, orderData bindings.SolverNetOrderData) (output, error) {
		return openOrder(ctx, backends, network, flowgenAddr, conf.srcChain, scl, orderData)
	}

	outputs, cancel := forkjoin.NewWithInputs(ctx, work, orderDatas, forkjoin.WithWorkers(16))
	defer cancel()

	all, err := outputs.Flatten()
	if err != nil {
		return nil, errors.Wrap(err, "open orders")
	}

	// Filter out skipped orders
	var results []types.Result
	for _, result := range all {
		if result.Skip {
			continue
		}

		results = append(results, result.Result)
	}

	return results, nil
}

func splitOrderAmounts(bounds solver.SpendBounds, total *big.Int, split int) []*big.Int {
	avg := bi.DivRaw(total, split)
	remaining := bi.Clone(total)

	var resp []*big.Int
	for len(resp) < split {
		next, ok := nextOrderAmount(bounds, remaining, avg)
		if !ok {
			break
		}

		remaining = bi.Sub(remaining, next)

		resp = append(resp, next)
	}

	return resp
}

func nextOrderAmount(bounds solver.SpendBounds, remaining *big.Int, target *big.Int) (*big.Int, bool) {
	// If not enough remaining, return nothing
	if bi.LT(remaining, bounds.MinSpend) {
		return nil, false
	}

	// If target amount is less than min spend, increase it
	if bi.LT(target, bounds.MinSpend) {
		target = bounds.MinSpend
	}

	// If target amount is greater than remaining, decrease it
	if bi.GT(target, remaining) {
		target = remaining
	}

	// If target amount is greater than max spend, decrease it
	if bi.GT(target, bounds.MaxSpend) {
		target = bounds.MaxSpend
	}

	return target, true
}

func check(ctx context.Context, scl sclient.Client, srcChainID uint64, orderData bindings.SolverNetOrderData) (stypes.CheckResponse, error) {
	checkReq, err := stypes.CheckRequestFromOrderData(srcChainID, orderData)
	if err != nil {
		return stypes.CheckResponse{}, err
	}

	return scl.Check(ctx, checkReq)
}

// swapPrice returns the price of the source token in destination tokens.
func swapPrice(ctx context.Context, scl sclient.Client, srcToken, dstToken tokens.Token) (stypes.Price, error) {
	priceReq := stypes.PriceRequest{
		SourceChainID:      srcToken.ChainID,
		DestinationChainID: dstToken.ChainID,
		DepositToken:       srcToken.Address,
		ExpenseToken:       dstToken.Address,
	}

	return scl.Price(ctx, priceReq)
}

func nativeOrderData(ctx context.Context, scl sclient.Client, owner common.Address, srcToken, dstToken tokens.Token, depositAmt *big.Int) (bindings.SolverNetOrderData, error) {
	quoteReq := stypes.QuoteRequest{
		SourceChainID:      srcToken.ChainID,
		DestinationChainID: dstToken.ChainID,
		Deposit: stypes.AddrAmt{
			Token:  srcToken.Address,
			Amount: depositAmt,
		},
	}

	quoteResp, err := scl.Quote(ctx, quoteReq)
	if err != nil {
		return bindings.SolverNetOrderData{}, errors.Wrap(err, "quote deposit")
	} else if quoteResp.Rejected {
		return bindings.SolverNetOrderData{}, errors.New("quote rejected", "description", quoteResp.RejectDescription, "reason", quoteResp.RejectCode)
	} else if !bi.EQ(quoteResp.Deposit.Amount, depositAmt) {
		return bindings.SolverNetOrderData{}, errors.New("quote deposit amount mismatch", "expected", depositAmt, "actual", quoteResp.Deposit.Amount)
	}

	expenseAmt := quoteResp.Expense.Amount
	call := solvernet.Call{
		Target: owner,
		Value:  expenseAmt,
	}

	return bindings.SolverNetOrderData{
		Owner:       owner,
		DestChainId: dstToken.ChainID,
		Deposit:     solvernet.Deposit(quoteResp.Deposit),
		Expenses:    []solvernet.Expense{}, // Explicit expense not required for native transfer calls.
		Calls:       []bindings.SolverNetCall{call.ToBinding()},
	}, nil
}

type output struct {
	types.Result
	Skip bool
}

func openOrder(
	ctx context.Context,
	backends ethbackend.Backends,
	networkID netconf.ID,
	owner common.Address,
	srcChainID uint64,
	scl sclient.Client,
	orderData bindings.SolverNetOrderData,
) (output, error) {
	if resp, err := check(ctx, scl, srcChainID, orderData); err != nil {
		return output{}, errors.Wrap(err, "check")
	} else if resp.RejectCode == stypes.RejectInsufficientInventory {
		log.Debug(ctx, "Skipping bridge order due to insufficient inventory", "description", resp.RejectDescription)
		return output{Skip: true}, nil
	} else if resp.Rejected {
		return output{}, errors.New("order rejected", "description", resp.RejectDescription, "reason", resp.RejectCode)
	}

	orderID, err := solvernet.OpenOrder(ctx, networkID, srcChainID, backends, owner, orderData)
	if err != nil {
		return output{}, errors.Wrap(err, "open order")
	}

	return output{Result: types.Result{OrderID: orderID, Data: orderData}}, nil
}

// availableBalance returns the available balance to spend on orders.
func availableBalance(
	ctx context.Context,
	network netconf.ID,
	client ethclient.Client,
	role eoa.Role,
	token tokens.Token,
) (*big.Int, error) {
	if !token.IsNative() {
		return nil, errors.New("only native tokens supported", "token", token)
	}

	addr, ok := eoa.Address(network, role)
	if !ok {
		return nil, errors.New("invalid role", "role", role)
	}

	balance, err := client.BalanceAt(ctx, addr, nil)
	if err != nil {
		return nil, errors.Wrap(err, "balance at")
	}

	thresholds, ok := eoa.GetFundThresholds(token.Asset, network, role)
	if !ok {
		return nil, errors.New("no role thresholds found", "asset", token.Asset, "role", role)
	}

	reserved := bi.Ether(0.01) // overhead that should cover solver commission and tx fees

	return bi.Sub(balance, thresholds.MinBalance(), reserved), nil
}
