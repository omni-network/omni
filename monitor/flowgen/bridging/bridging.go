package bridging

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/monitor/flowgen/types"
	solver "github.com/omni-network/omni/solver/app"
	stokens "github.com/omni-network/omni/solver/tokens"
	stypes "github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"
)

// Don't go too low, because we need to generate volume.
var minOrderSize = bi.Ether(0.1)

// NewJob returns a job that bridges native tokens.
func newJob(ctx context.Context, networkID netconf.ID, backends ethbackend.Backends, conf flowConfig, owner common.Address) (types.Job, error) {
	cadence := 25 * time.Minute
	if networkID == netconf.Devnet {
		cadence = time.Second * 10
	}

	backend, err := backends.Backend(conf.srcChain)
	if err != nil {
		return types.Job{}, errors.Wrap(err, "src chain backend")
	}

	addrs, err := contracts.GetAddresses(ctx, networkID)
	if err != nil {
		return types.Job{}, errors.New("contract addresses")
	}

	solverAddr := eoa.MustAddress(networkID, eoa.RoleSolver)
	checker := solver.NewChecker(
		backends,
		func(uint64, common.Address, []byte) bool { return true },
		solverAddr,
		addrs.SolverNetOutbox,
	)

	namer := netconf.ChainNamer(networkID)

	return types.Job{
		Name:      fmt.Sprintf("Bridging (%v->%v)", namer(conf.srcChain), namer(conf.dstChain)),
		Cadence:   cadence,
		NetworkID: networkID,

		SrcChainBackend: backend,

		OpenOrderFunc: func(ctx context.Context) (types.Result, bool, error) {
			return openOrder(ctx, backends, backend, networkID, owner, conf, checker)
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
	checker solver.CheckFunc,
) (types.Result, bool, error) {
	srcToken, ok := stokens.Native(conf.srcChain)
	if !ok {
		return types.Result{}, false, errors.New("src token not found")
	}

	dstToken, ok := stokens.Native(conf.dstChain)
	if !ok {
		return types.Result{}, false, errors.New("dst token not found")
	}

	orderSize, ok, err := getOrderSize(ctx, networkID, backend.Client, owner, conf, srcToken.Token)
	if err != nil {
		return types.Result{}, false, errors.Wrap(err, "get order size")
	}

	if !ok {
		return types.Result{}, false, nil
	}

	expense := solver.TokenAmt{Token: dstToken, Amount: orderSize}

	depositWithFee, err := solver.QuoteDeposit(srcToken, expense)
	if err != nil {
		return types.Result{}, false, errors.Wrap(err, "quote deposit")
	}

	call := solvernet.Call{
		Target: owner,
		Value:  expense.Amount,
	}

	orderData := bindings.SolverNetOrderData{
		Owner:       owner,
		DestChainId: conf.dstChain,
		Deposit: solvernet.Deposit{
			Token:  depositWithFee.Token.Address,
			Amount: depositWithFee.Amount,
		},
		Expenses: []solvernet.Expense{}, // Explicit expense not required for native transfer calls.
		Calls:    []bindings.SolverNetCall{call.ToBinding()},
	}

	orderID, err := solvernet.OpenOrder(ctx, networkID, conf.srcChain, backends, owner, orderData)
	if err != nil {
		return types.Result{}, false, errors.Wrap(err, "open order")
	}

	if err = checker(ctx, stypes.CheckRequest{
		SourceChainID:      conf.srcChain,
		DestinationChainID: conf.dstChain,
		FillDeadline:       uint32(time.Now().Add(time.Hour).Unix()), //nolint:gosec // unix timestamp won't overflow
		Deposit:            stypes.AddrAmt(orderData.Deposit),
		Expenses: []stypes.Expense{{
			Amount: call.Value,
		}},
		Calls: []stypes.Call{stypes.Call(call)},
	}); err != nil {
		var rejErr *solver.RejectionError
		if errors.As(err, &rejErr) && rejErr.Reason == stypes.RejectInsufficientInventory {
			return types.Result{}, false, nil
		}

		return types.Result{}, false, err
	}

	return types.Result{OrderID: orderID, Expense: expense}, true, nil
}

// Jobs returns two jobs bridging native ETH from one chain to another one and back.
func Jobs(ctx context.Context, networkID netconf.ID, backends ethbackend.Backends, owner common.Address) ([]types.Job, error) {
	conf, ok := config[networkID]
	if !ok {
		return nil, nil
	}

	job1, err := newJob(ctx, networkID, backends, conf, owner)
	if err != nil {
		return nil, err
	}

	// Clone the job and flip the chains
	conf2 := conf
	conf2.srcChain, conf2.dstChain = conf2.dstChain, conf2.srcChain
	job2, err := newJob(ctx, networkID, backends, conf2, owner)
	if err != nil {
		return nil, err
	}

	return []types.Job{job1, job2}, nil
}

// getOrderSize checks the current balance of the flowgen EOA and returns
// the maximal possible order size or false if the minimal balance is reached.
func getOrderSize(
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

	// use solver's spend bounds on the destination chain
	nativeEthTkn, ok := stokens.Native(conf.dstChain)
	if !ok {
		return nil, false, nil
	}

	reserved := bi.Ether(0.01) // overhead that should cover solver commission and tx fees
	orderSize := bi.Sub(balance, thresholds.MinBalance(), reserved)

	// if order size is too small, do nothing
	if bi.LT(orderSize, minOrderSize) {
		return nil, false, nil
	}

	// cap the order if necessary
	if bi.GT(orderSize, nativeEthTkn.MaxSpend) {
		orderSize = nativeEthTkn.MaxSpend
	}

	return orderSize, true, nil
}
