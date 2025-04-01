package solve

import (
	"bytes"
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"net/url"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"
	solver "github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"golang.org/x/sync/errgroup"
)

type TestOrder struct {
	Owner         common.Address
	FillDeadline  time.Time
	DestChainID   uint64
	SourceChainID uint64
	Expenses      []solvernet.Expense
	Calls         []solvernet.Call
	Deposit       solvernet.Deposit
	ShouldReject  bool
	RejectReason  string
}

func isDepositTokenInvalid(o TestOrder) bool {
	return o.Deposit.Token == invalidTokenAddress
}

func isDepositTokenEmpty(o TestOrder) bool {
	return o.Deposit.Token == zeroAddr
}

func srcAndDestChainAreSame(o TestOrder) bool {
	return o.SourceChainID == o.DestChainID
}

func isSrcChainInvalid(o TestOrder) bool {
	return o.SourceChainID == invalidChainID
}

func isInsufficientInventory(o TestOrder) bool {
	return o.RejectReason == solver.RejectInsufficientInventory.String()
}

func isInvalidExpense(o TestOrder) bool {
	return o.RejectReason == solver.RejectInvalidExpense.String()
}

func Test(ctx context.Context, network netconf.Network, endpoints xchain.RPCEndpoints, solverAddr string) error {
	if network.ID != netconf.Devnet {
		return errors.New("only devnet")
	}

	log.Info(ctx, "Running solver tests")

	// use anvil.DevAccounts instead of eoa.DevAccounts, because eoa.DevAccounts
	// are used frequently elsewhere in e2e / e2e tests, and nonce issues get annoying
	backends, err := ethbackend.BackendsFromNetwork(network, endpoints, anvil.DevPrivateKeys()...)
	if err != nil {
		return err
	}

	orders := makeOrders()

	if err = mintAndApproveAll(ctx, backends, orders); err != nil {
		return errors.Wrap(err, "mint omni")
	}

	if err = testCheckAPI(ctx, backends, orders, solverAddr); err != nil {
		return errors.Wrap(err, "test check api")
	}

	tracker, err := openAll(ctx, backends, orders)
	if err != nil {
		return errors.Wrap(err, "open all")
	}

	xprov := xprovider.New(network, backends.Clients(), nil)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// start event streams
	errChan := make(chan error, 1)
	// Set a 90 sec timeout which is a bit below the 120 sec go test timeout in order to avoid being canceled by the go runtime and leave time for cleanup.
	ctx, cancel = context.WithTimeout(ctx, 90*time.Second)
	defer cancel()

	for _, chain := range network.EVMChains() {
		go func() {
			req := xchain.EventLogsReq{
				ChainID:       chain.ID,
				ConfLevel:     xchain.ConfLatest,
				Height:        1,
				FilterAddress: addrs.SolverNetInbox,
				FilterTopics:  solvernet.AllEventTopics(),
			}

			// Stream all inbox event logs and update order status in tracker
			proc := func(ctx context.Context, _ *types.Header, logs []types.Log) error {
				for _, l := range logs {
					orderID, status, err := solvernet.ParseEvent(l)
					if err != nil {
						return err
					}

					log.Debug(ctx, "Order status updated", "status", status, "order_id", orderID)

					tracker.UpdateStatus(orderID, chain.ID, status)
				}

				return nil
			}

			err := xprov.StreamEventLogs(ctx, req, proc)
			if err != nil {
				errChan <- err
			}
		}()
	}

	// wait Done
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case err := <-errChan:
			return errors.Wrap(err, "stream event logs")
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "context timeout or canceled")
		case <-ticker.C:
			done, err := tracker.Done()
			if err != nil {
				return errors.Wrap(err, "solver tracker failed")
			}

			if done {
				// if done, wait for solver to rebalance OMNI (bridge L1 back to native)
				if err := waitRebalance(ctx, backends); err != nil {
					return errors.Wrap(err, "wait rebalance")
				}

				log.Info(ctx, "Solver test success")

				return nil
			}
		}
	}
}

func makeOrders() []TestOrder {
	users := anvil.DevAccounts()
	var orders []TestOrder

	// erc20 OMNI -> native OMNI orders
	{
		omniBridgeOrder := func(addr common.Address, expense *big.Int, deposit solvernet.Deposit, rejectReason string) TestOrder {
			return TestOrder{
				Owner:         addr,
				FillDeadline:  time.Now().Add(1 * time.Hour),
				SourceChainID: evmchain.IDMockL1,
				DestChainID:   evmchain.IDOmniDevnet,
				Expenses:      nativeExpense(expense),
				Calls:         nativeTransferCall(expense, addr),
				Deposit:       deposit,
				ShouldReject:  rejectReason != "",
				RejectReason:  rejectReason,
			}
		}

		expense := bi.Ether(10) // Note that fee is not charged for OMNI, so deposit==expense
		orders = append(orders,
			omniBridgeOrder(users[0], expense, erc20Deposit(expense, addrs.Token), ""),
			omniBridgeOrder(users[1], expense, nativeDeposit(expense), solver.RejectInvalidDeposit.String()),
			omniBridgeOrder(users[2], expense, unsupportedERC20Deposit(expense), solver.RejectUnsupportedDeposit.String()),
		)
	}

	// native ETH transfers
	{
		ethTransferOrder := func(addr common.Address, expense *big.Int, rejectReason string) TestOrder {
			return TestOrder{
				Owner:         addr,
				FillDeadline:  time.Now().Add(1 * time.Hour),
				SourceChainID: evmchain.IDMockL1,
				DestChainID:   evmchain.IDMockL2,
				Expenses:      nativeExpense(expense),
				Calls:         nativeTransferCall(expense, addr),
				Deposit:       nativeDeposit(bi.Add(expense, bi.Ether(0.1))), // add enough to cover fee
				ShouldReject:  rejectReason != "",
				RejectReason:  rejectReason,
			}
		}

		validExpense := bi.Ether(1)
		underMinExpense := bi.Gwei(1)   // Min is 0.001 ETH
		overMaxExpense := bi.Ether(100) // Max is 1 ETH

		orders = append(orders,
			ethTransferOrder(users[3], validExpense, ""),
			ethTransferOrder(users[4], underMinExpense, solver.RejectExpenseUnderMin.String()),
			ethTransferOrder(users[5], overMaxExpense, solver.RejectExpenseOverMax.String()),
		)
	}

	// bad src chain
	orders = append(orders, TestOrder{
		Owner:         users[0],
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: invalidChainID,
		DestChainID:   evmchain.IDMockL1,
		Expenses:      nativeExpense(bi.Wei(1)),
		Calls:         nativeTransferCall(bi.Wei(1), users[0]),
		Deposit:       erc20Deposit(bi.Wei(1), zeroAddr),
		ShouldReject:  true,
		RejectReason:  solver.RejectUnsupportedSrcChain.String(),
	})

	// bad dest chain
	orders = append(orders, TestOrder{
		Owner:         users[0],
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL1,
		DestChainID:   invalidChainID,
		Expenses:      nativeExpense(bi.Wei(1)),
		Calls:         nativeTransferCall(bi.Wei(1), users[0]),
		Deposit:       erc20Deposit(bi.Wei(1), addrs.Token),
		ShouldReject:  true,
		RejectReason:  solver.RejectUnsupportedDestChain.String(),
	})

	// same chain for src and dest
	orders = append(orders, TestOrder{
		Owner:         users[0],
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL1,
		DestChainID:   evmchain.IDMockL1,
		Expenses:      nativeExpense(bi.Wei(1)),
		Calls:         nativeTransferCall(bi.Wei(1), users[0]),
		Deposit:       erc20Deposit(bi.Wei(1), addrs.Token),
		ShouldReject:  true,
		RejectReason:  solver.RejectSameChain.String(),
	})

	// unsupported expense token
	orders = append(orders, TestOrder{
		Owner:         users[0],
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL1,
		DestChainID:   evmchain.IDMockL2,
		Expenses:      unsupportedExpense(bi.Wei(1)),
		Calls:         nativeTransferCall(bi.Wei(1), users[0]),
		Deposit:       erc20Deposit(bi.Wei(1), addrs.Token),
		ShouldReject:  true,
		RejectReason:  solver.RejectUnsupportedExpense.String(),
	})

	// invalid expense
	orders = append(orders, TestOrder{
		Owner:         users[0],
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL1,
		DestChainID:   evmchain.IDMockL2,
		Expenses:      invalidExpenseOutOfBounds(),
		Calls:         nativeTransferCall(bi.Wei(1), users[0]),
		Deposit:       erc20Deposit(bi.Wei(1), addrs.Token),
		ShouldReject:  true,
		RejectReason:  solver.RejectInvalidExpense.String(),
	})

	// invalid expense with multiple native expenses
	orders = append(orders, TestOrder{
		Owner:         users[0],
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL1,
		DestChainID:   evmchain.IDMockL2,
		Expenses:      multipleNativeExpenses(validETHSpend),
		Calls:         nativeTransferCall(validETHSpend, users[0]),
		Deposit:       nativeDeposit(maxETHSpend),
		ShouldReject:  true,
		RejectReason:  solver.RejectInvalidExpense.String(),
	})

	// invalid expense with multiple ERC20 expenses
	orders = append(orders, TestOrder{
		Owner:         users[0],
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL2,
		DestChainID:   evmchain.IDMockL1,
		Expenses:      multipleERC20Expenses(validETHSpend),
		Deposit:       nativeDeposit(maxETHSpend),
		ShouldReject:  true,
		RejectReason:  solver.RejectInvalidExpense.String(),
	})

	// valid order with valid ETH spend
	orders = append(orders, TestOrder{
		Owner:         users[0],
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL1,
		DestChainID:   evmchain.IDMockL2,
		Expenses:      nativeExpense(validETHSpend),
		Calls:         nativeTransferCall(validETHSpend, users[0]),
		Deposit:       nativeDeposit(maxETHSpend),
		ShouldReject:  false,
	})

	// dest call reverts
	orders = append(orders, TestOrder{
		Owner:         users[0],
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL1,
		DestChainID:   evmchain.IDMockL2,
		Expenses:      nativeExpense(validETHSpend),
		Calls:         contractCallWithInvalidCallData(),
		Deposit:       nativeDeposit(maxETHSpend),
		ShouldReject:  true,
		RejectReason:  solver.RejectDestCallReverts.String(),
	})

	// insufficient inventory for native expense
	orders = append(orders, TestOrder{
		Owner:         users[0],
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL2,
		DestChainID:   evmchain.IDMockL1,
		Expenses:      nativeExpense(validETHSpend),
		Calls:         nativeTransferCall(validETHSpend, users[0]),
		Deposit:       nativeDeposit(maxETHSpend),
		ShouldReject:  true,
		RejectReason:  solver.RejectInsufficientInventory.String(),
	})

	// TODO: add test order for insufficient inventory for ERC20 expenses

	return orders
}

func openAll(ctx context.Context, backends ethbackend.Backends, orders []TestOrder) (*orderTracker, error) {
	tracker := newOrderTracker()

	var eg errgroup.Group
	for _, order := range orders {
		eg.Go(func() error {
			if isInvalidExpense(order) || isSrcChainInvalid(order) || isDepositTokenInvalid(order) || srcAndDestChainAreSame(order) || isInsufficientInventory(order) {
				return nil
			}

			id, err := openOrder(ctx, backends, order)
			if err != nil {
				return err
			}

			tracker.TrackOrder(id, order)

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, errors.Wrap(err, "wait group")
	}

	// Mark all orders as tracked
	tracker.AllTracked()

	return tracker, nil
}

func openOrder(ctx context.Context, backends ethbackend.Backends, order TestOrder) ([32]byte, error) {
	return solvernet.OpenOrder(ctx, netconf.Devnet, order.SourceChainID, backends, order.Owner, bindings.SolverNetOrderData{
		DestChainId: order.DestChainID,
		Expenses:    solvernet.FilterNativeExpenses(order.Expenses),
		Calls:       solvernet.CallsToBindings(order.Calls),
		Deposit:     order.Deposit,
	}, solvernet.WithFillDeadline(order.FillDeadline))
}

func testCheckAPI(ctx context.Context, backends ethbackend.Backends, orders []TestOrder, solverAddr string) error {
	uri, err := url.JoinPath(solverAddr, "/api/v1/check")
	if err != nil {
		return errors.Wrap(err, "get api uri")
	}

	for i, order := range orders {
		// If this order requires balance draining, do it before test logic.
		if isInsufficientInventory(order) {
			// Drain solver native balance.
			if err := setSolverAccountNativeBalance(ctx, order.DestChainID, backends, bi.Zero()); err != nil {
				return errors.Wrap(err, "drain solver account failed")
			}
		}

		checkReq := solver.CheckRequest{
			FillDeadline:       uint32(order.FillDeadline.Unix()), //nolint:gosec // this is fine for tests
			SourceChainID:      order.SourceChainID,
			DestinationChainID: order.DestChainID,
			Expenses:           expensesFromBindings(order.Expenses),
			Calls:              callsFromBindings(order.Calls),
			Deposit:            addrAmtFromDeposit(order.Deposit),
		}

		body, err := json.Marshal(checkReq)
		if err != nil {
			return errors.Wrap(err, "marshal request")
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, bytes.NewBuffer(body))
		if err != nil {
			return errors.Wrap(err, "new request")
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return errors.Wrap(err, "do request")
		}
		if resp.StatusCode != http.StatusOK {
			return errors.New("bad status", "status", resp.StatusCode)
		}

		var checkResp solver.CheckResponse
		if err = json.NewDecoder(resp.Body).Decode(&checkResp); err != nil {
			return errors.Wrap(err, "decode response")
		}

		if err = resp.Body.Close(); err != nil {
			return errors.Wrap(err, "close response body")
		}

		if checkResp.Rejected != order.ShouldReject {
			return errors.New("unexpected rejection",
				"expected", order.ShouldReject,
				"actual", checkResp.Rejected,
				"reason", checkResp.RejectReason,
				"description", checkResp.RejectDescription,
				"order_idx", i,
			)
		}

		if checkResp.RejectReason != order.RejectReason {
			return errors.New("unexpected reject reason", "expected", order.RejectReason, "actual", checkResp.RejectReason)
		}

		if checkResp.Accepted && order.ShouldReject {
			return errors.New("accepted but should reject")
		}

		// Refund solver native balance after test logic.
		if isInsufficientInventory(order) {
			eth1m := bi.Ether(1_000_000)
			if err := setSolverAccountNativeBalance(ctx, order.DestChainID, backends, eth1m); err != nil {
				return errors.Wrap(err, "refund solver account failed")
			}
		}
	}

	log.Info(ctx, "Test /check success")

	return nil
}

func waitRebalance(ctx context.Context, backends ethbackend.Backends) error {
	backend, err := backends.Backend(evmchain.IDMockL1)
	if err != nil {
		return errors.Wrap(err, "l1 backend")
	}

	solver := eoa.MustAddress(netconf.Devnet, eoa.RoleSolver)

	token, err := bindings.NewIERC20(contracts.TokenAddr(netconf.Devnet), backend)
	if err != nil {
		return errors.Wrap(err, "new erc20")
	}

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "timeout")
		case <-ticker.C:
			balance, err := token.BalanceOf(&bind.CallOpts{Context: ctx}, solver)
			if err != nil {
				return errors.Wrap(err, "balance of")
			}

			// solver will have claimed much more than 1 OMNI
			// if balance is < 1, rebalancing is working
			if bi.LTE(balance, bi.Ether(1)) {
				return nil
			}
		}
	}
}
