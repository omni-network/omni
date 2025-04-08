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
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/monitor/flowgen/types"
	sclient "github.com/omni-network/omni/solver/client"
	stokens "github.com/omni-network/omni/solver/tokens"
	stypes "github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"
)

// parallel is the number of parallel orders to open.
const parallelDev = 4
const parallel = 64

// Jobs returns two jobs bridging native ETH from one chain to another one and back.
func Jobs(
	networkID netconf.ID,
	backends ethbackend.Backends,
	owner common.Address,
	scl sclient.Client,
) []types.Job {
	conf, ok := config[networkID]
	if !ok {
		return nil
	}

	return []types.Job{
		newJob(networkID, backends, conf, owner, scl),
		newJob(networkID, backends, conf.Flip(), owner, scl), // Second job using flipped config
	}
}

// NewJob returns a job that bridges native tokens.
func newJob(
	networkID netconf.ID,
	backends ethbackend.Backends,
	conf flowConfig,
	owner common.Address,
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
			return openOrders(ctx, backends, networkID, owner, conf, scl)
		},
	}
}

func openOrders(
	ctx context.Context,
	backends ethbackend.Backends,
	networkID netconf.ID,
	owner common.Address,
	conf flowConfig,
	scl sclient.Client,
) ([]types.Result, error) {
	srcToken, ok := stokens.Native(conf.srcChain)
	if !ok {
		return nil, errors.New("src token not found")
	}

	dstToken, ok := stokens.Native(conf.dstChain)
	if !ok {
		return nil, errors.New("dst token not found")
	}

	backend, err := backends.Backend(conf.srcChain)
	if err != nil {
		return nil, err
	}

	totalAmount, err := availableBalance(ctx, networkID, backend.Client, owner, srcToken.Token)
	if err != nil {
		return nil, errors.Wrap(err, "get order size")
	}

	p := parallel
	if networkID == netconf.Devnet {
		p = parallelDev
	}

	var orderDatas []bindings.SolverNetOrderData
	for _, amount := range splitOrderAmounts(dstToken, totalAmount, p) {
		orderData, err := nativeOrderData(ctx, scl, owner, srcToken, dstToken, amount)
		if err != nil {
			return nil, err
		}

		orderDatas = append(orderDatas, orderData)
	}

	work := func(ctx context.Context, orderData bindings.SolverNetOrderData) (types.Result, error) {
		return openOrder(ctx, backends, networkID, owner, conf.srcChain, scl, orderData)
	}

	results, cancel := forkjoin.NewWithInputs(ctx, work, orderDatas, forkjoin.WithWorkers(16))
	defer cancel()

	return results.Flatten()
}

func splitOrderAmounts(dstChain stokens.Token, total *big.Int, split int) []*big.Int {
	avg := bi.DivRaw(total, split)
	remaining := bi.Clone(total)

	var resp []*big.Int
	for len(resp) < split {
		next, ok := nextOrderAmount(dstChain, remaining, avg)
		if !ok {
			break
		}

		remaining = bi.Sub(remaining, next)

		resp = append(resp, next)
	}

	return resp
}

func nextOrderAmount(dstChain stokens.Token, remaining *big.Int, target *big.Int) (*big.Int, bool) {
	// If not enough remaining, return nothing
	if bi.LT(remaining, dstChain.MinSpend) {
		return nil, false
	}

	// If target amount is less than min spend, increase it
	if bi.LT(target, dstChain.MinSpend) {
		target = dstChain.MinSpend
	}

	// If target amount is greater than remaining, decrease it
	if bi.GT(target, remaining) {
		target = remaining
	}

	// If target amount is greater than max spend, decrease it
	if bi.GT(target, dstChain.MaxSpend) {
		target = dstChain.MaxSpend
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

func nativeOrderData(ctx context.Context, scl sclient.Client, owner common.Address, srcToken, dstToken stokens.Token, expenseAmt *big.Int) (bindings.SolverNetOrderData, error) {
	quoteReq := stypes.QuoteRequest{
		SourceChainID:      srcToken.ChainID,
		DestinationChainID: dstToken.ChainID,
		Expense: stypes.AddrAmt{
			Token:  dstToken.Address,
			Amount: expenseAmt,
		},
	}

	quoteResp, err := scl.Quote(ctx, quoteReq)
	if err != nil {
		return bindings.SolverNetOrderData{}, errors.Wrap(err, "quote deposit")
	} else if quoteResp.Rejected {
		return bindings.SolverNetOrderData{}, errors.New("quote rejected", "description", quoteResp.RejectDescription, "reason", quoteResp.RejectCode)
	}

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

func openOrder(
	ctx context.Context,
	backends ethbackend.Backends,
	networkID netconf.ID,
	owner common.Address,
	srcChainID uint64,
	scl sclient.Client,
	orderData bindings.SolverNetOrderData,
) (types.Result, error) {
	if resp, err := check(ctx, scl, srcChainID, orderData); err != nil {
		return types.Result{}, errors.Wrap(err, "check")
	} else if resp.Rejected {
		return types.Result{}, errors.New("order rejected", "description", resp.RejectDescription, "reason", resp.RejectCode)
	}

	orderID, err := solvernet.OpenOrder(ctx, networkID, srcChainID, backends, owner, orderData)
	if err != nil {
		return types.Result{}, errors.Wrap(err, "open order")
	}

	return types.Result{OrderID: orderID, Data: orderData}, nil
}

// availableBalance returns the available flowgen balance to spend on orders.
func availableBalance(
	ctx context.Context,
	networkID netconf.ID,
	client ethclient.Client,
	owner common.Address,
	srcToken tokens.Token,
) (*big.Int, error) {
	balance, err := client.BalanceAt(ctx, owner, nil)
	if err != nil {
		return nil, errors.Wrap(err, "balance at")
	}

	thresholds, ok := eoa.GetFundThresholds(srcToken, networkID, eoa.RoleFlowgen)
	if !ok {
		return nil, errors.New("no thresholds found", "role", eoa.RoleFlowgen)
	}

	reserved := bi.Ether(0.01) // overhead that should cover solver commission and tx fees

	return bi.Sub(balance, thresholds.MinBalance(), reserved), nil
}
