package solve

import (
	"bytes"
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"
	solver "github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
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

func Test(ctx context.Context, network netconf.Network, endpoints xchain.RPCEndpoints) error {
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

	if err = testCheckAPI(ctx, backends, orders); err != nil {
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
	timeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Minute+30*time.Second)
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
			proc := func(ctx context.Context, _ uint64, logs []types.Log) error {
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
		case <-timeoutCtx.Done():
			return errors.Wrap(timeoutCtx.Err(), "context timeout or canceled")
		case <-ticker.C:
			done, err := tracker.Done()
			if err != nil {
				return errors.Wrap(err, "solver tracker failed")
			}

			if done {
				log.Info(timeoutCtx, "Solver test success")
				return nil
			}
		}
	}
}

func makeOrders() []TestOrder {
	users := anvil.DevAccounts()
	var orders []TestOrder

	// erc20 OMNI -> native OMNI orders
	for i, user := range users {
		requestAmt := math.NewInt(10).MulRaw(params.Ether).BigInt()

		// make some insufficient (should reject)
		insufficientDeposit := i%2 == 0
		depositAmt := new(big.Int).Set(requestAmt)
		if insufficientDeposit {
			depositAmt = depositAmt.Div(depositAmt, big.NewInt(2))
		}

		shouldReject := insufficientDeposit
		rejectReason := ""
		if insufficientDeposit {
			rejectReason = solver.RejectInsufficientDeposit.String()
		}

		orders = append(orders, TestOrder{
			Owner:         user,
			FillDeadline:  time.Now().Add(1 * time.Hour),
			SourceChainID: evmchain.IDMockL1,
			DestChainID:   evmchain.IDOmniDevnet,
			Expenses:      nativeExpense(requestAmt),
			Calls:         nativeTransferCall(requestAmt, user),
			Deposit:       erc20Deposit(depositAmt, addrs.Token),
			ShouldReject:  shouldReject,
			RejectReason:  rejectReason,
		})

		orders = append(orders, TestOrder{
			Owner:         user,
			FillDeadline:  time.Now().Add(1 * time.Hour),
			SourceChainID: evmchain.IDMockL1,
			DestChainID:   evmchain.IDOmniDevnet,
			Expenses:      nativeExpense(requestAmt),
			Calls:         nativeTransferCall(requestAmt, user),
			Deposit:       nativeDeposit(depositAmt),
			ShouldReject:  true,
			RejectReason:  solver.RejectInvalidDeposit.String(),
		})

		orders = append(orders, TestOrder{
			Owner:         user,
			FillDeadline:  time.Now().Add(1 * time.Hour),
			SourceChainID: evmchain.IDOmniDevnet,
			DestChainID:   evmchain.IDMockL1,
			Expenses:      nativeExpense(requestAmt),
			Calls:         nativeTransferCall(requestAmt, user),
			Deposit:       unsupportedERC20Deposit(depositAmt),
			ShouldReject:  true,
			RejectReason:  solver.RejectUnsupportedDeposit.String(),
		})
	}

	// native ETH transfers
	for i, user := range users {
		amt := math.NewInt(1).MulRaw(params.Ether).BigInt()

		// make some under min or over max expense
		overMax := i < 3
		underMin := i > 6

		if overMax {
			// max is 1 ETH
			amt = math.NewInt(2).MulRaw(params.Ether).BigInt()
		}

		if underMin {
			// min is 0.001 ETH
			amt = big.NewInt(1)
		}

		shouldReject := underMin || overMax
		rejectReason := ""
		if underMin {
			rejectReason = solver.RejectExpenseUnderMin.String()
		}
		if overMax {
			rejectReason = solver.RejectExpenseOverMax.String()
		}

		order := TestOrder{
			Owner:         user,
			FillDeadline:  time.Now().Add(1 * time.Hour),
			SourceChainID: evmchain.IDMockL1,
			DestChainID:   evmchain.IDMockL2,
			Expenses:      nativeExpense(amt),
			Calls:         nativeTransferCall(amt, user),
			Deposit:       nativeDeposit(new(big.Int).Add(amt, big.NewInt(1e17))), // add enough to cover fee
			ShouldReject:  shouldReject,
			RejectReason:  rejectReason,
		}

		orders = append(orders, order)
	}

	// bad src chain
	orders = append(orders, TestOrder{
		Owner:         users[0],
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: invalidChainID,
		DestChainID:   evmchain.IDMockL1,
		Expenses:      nativeExpense(big.NewInt(1)),
		Calls:         nativeTransferCall(big.NewInt(1), users[0]),
		Deposit:       erc20Deposit(big.NewInt(1), zeroAddr),
		ShouldReject:  true,
		RejectReason:  solver.RejectUnsupportedSrcChain.String(),
	})

	// bad dest chain
	orders = append(orders, TestOrder{
		Owner:         users[0],
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL1,
		DestChainID:   invalidChainID,
		Expenses:      nativeExpense(big.NewInt(1)),
		Calls:         nativeTransferCall(big.NewInt(1), users[0]),
		Deposit:       erc20Deposit(big.NewInt(1), addrs.Token),
		ShouldReject:  true,
		RejectReason:  solver.RejectUnsupportedDestChain.String(),
	})

	// same chain for src and dest
	orders = append(orders, TestOrder{
		Owner:         users[0],
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL1,
		DestChainID:   evmchain.IDMockL1,
		Expenses:      nativeExpense(big.NewInt(1)),
		Calls:         nativeTransferCall(big.NewInt(1), users[0]),
		Deposit:       erc20Deposit(big.NewInt(1), addrs.Token),
		ShouldReject:  true,
		RejectReason:  solver.RejectSameChain.String(),
	})

	// unsupported expense token
	orders = append(orders, TestOrder{
		Owner:         users[0],
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL1,
		DestChainID:   evmchain.IDMockL2,
		Expenses:      unsupportedExpense(big.NewInt(1)),
		Calls:         nativeTransferCall(big.NewInt(1), users[0]),
		Deposit:       erc20Deposit(big.NewInt(1), addrs.Token),
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
		Calls:         nativeTransferCall(big.NewInt(1), users[0]),
		Deposit:       erc20Deposit(big.NewInt(1), addrs.Token),
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

func testCheckAPI(ctx context.Context, backends ethbackend.Backends, orders []TestOrder) error {
	const url = "http://localhost:26661/api/v1/check"

	for i, order := range orders {
		// If this order requires balance draining, do it before test logic.
		if isInsufficientInventory(order) {
			// Drain solver native balance.
			if err := setSolverAccountNativeBalance(ctx, order.DestChainID, backends, big.NewInt(0)); err != nil {
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

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
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
			eth1m := math.NewInt(1_000_000).MulRaw(params.Ether).BigInt()
			if err := setSolverAccountNativeBalance(ctx, order.DestChainID, backends, eth1m); err != nil {
				return errors.Wrap(err, "refund solver account failed")
			}
		}
	}

	log.Info(ctx, "Test /check success")

	return nil
}
