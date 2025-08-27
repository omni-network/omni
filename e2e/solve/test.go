package solve

import (
	"context"
	"math/big"
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
	"github.com/omni-network/omni/lib/tokenpricer"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/xchain"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"
	sclient "github.com/omni-network/omni/solver/client"
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

func Test(ctx context.Context, network netconf.Network, backends ethbackend.Backends, solverAddr string) error {
	if network.ID != netconf.Devnet {
		return errors.New("only devnet")
	}

	log.Info(ctx, "Running solver tests")

	orders := makeOrders() //nolint:contextcheck // Not critical code

	if err := mintAndApproveAll(ctx, backends, orders); err != nil {
		return errors.Wrap(err, "mint all")
	}

	if err := testCheckAPI(ctx, backends, orders, solverAddr); err != nil {
		return errors.Wrap(err, "test check api")
	}

	if err := testRelayAPI(ctx, solverAddr); err != nil {
		return errors.Wrap(err, "test relay api")
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
				ChainID:         chain.ID,
				ConfLevel:       xchain.ConfLatest,
				Height:          1,
				FilterAddresses: []common.Address{addrs.SolverNetInbox},
				FilterTopics:    solvernet.AllEventTopics(),
			}

			// Stream all inbox event logs and update order status in tracker
			proc := func(_ context.Context, _ *types.Header, logs []types.Log) error {
				for _, l := range logs {
					orderID, status, err := solvernet.ParseEvent(l)
					if err != nil {
						return err
					}

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
			remaining, err := tracker.Remaining()
			if err != nil {
				return errors.Wrap(err, "solver tracker failed")
			} else if remaining > 0 {
				log.Debug(ctx, "Some orders still pending", "remaining", remaining, "total", tracker.Len())
				tracker.DebugFirstPending(ctx)

				continue
			}

			// if done, wait for solver to rebalance NOM (bridge L1 back to native)
			if err := waitRebalance(ctx, backends); err != nil {
				return errors.Wrap(err, "wait rebalance")
			}

			log.Info(ctx, "Solver test success")

			return nil
		}
	}
}

func makeOrders() []TestOrder {
	user := anvil.DevAccount8()
	var orders []TestOrder

	// swaps
	{
		//nolint:unparam // test code
		swapOrder := func(addr common.Address, srcChain uint64, srcAsset tokens.Asset, dstChain uint64, dstAsset tokens.Asset) TestOrder {
			amount := bi.Ether(2)
			srcToken := mustTokenByAsset(srcChain, srcAsset)
			dstToken := mustTokenByAsset(dstChain, dstAsset)
			expense, call := expenseAndCall(amount, dstToken, addr)
			depositAmount := mustDepositAmount(amount, srcToken, dstToken)

			return TestOrder{
				Owner:         addr,
				FillDeadline:  time.Now().Add(1 * time.Hour),
				DestChainID:   dstChain,
				SourceChainID: srcChain,
				Expenses:      expense,
				Calls:         call,
				Deposit: solvernet.Deposit{
					Token:  srcToken.Address,
					Amount: depositAmount,
				},
			}
		}

		orders = append(orders,
			swapOrder(user, evmchain.IDMockL1, tokens.ETH, evmchain.IDOmniDevnet, tokens.NOM),
			swapOrder(user, evmchain.IDMockL1, tokens.NOM, evmchain.IDMockL2, tokens.ETH),
			swapOrder(user, evmchain.IDMockL1, tokens.USDC, evmchain.IDOmniDevnet, tokens.NOM),
			swapOrder(user, evmchain.IDMockL1, tokens.NOM, evmchain.IDMockL1, tokens.ETH), // same chain swap
		)
	}
	// erc20 NOM -> native NOM orders
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

		expense := bi.Ether(10) // Note that fee is not charged for NOM, so deposit==expense
		orders = append(orders,
			omniBridgeOrder(user, expense, erc20Deposit(expense, addrs.NomToken), ""),
			omniBridgeOrder(user, expense, unsupportedERC20Deposit(expense), solver.RejectUnsupportedDeposit.String()),
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
			ethTransferOrder(user, validExpense, ""),
			ethTransferOrder(user, underMinExpense, solver.RejectExpenseUnderMin.String()),
			ethTransferOrder(user, overMaxExpense, solver.RejectExpenseOverMax.String()),
		)
	}

	// bad src chain
	orders = append(orders, TestOrder{
		Owner:         user,
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: invalidChainID,
		DestChainID:   evmchain.IDMockL1,
		Expenses:      nativeExpense(bi.Wei(1)),
		Calls:         nativeTransferCall(bi.Wei(1), user),
		Deposit:       erc20Deposit(bi.Wei(1), zeroAddr),
		ShouldReject:  true,
		RejectReason:  solver.RejectUnsupportedSrcChain.String(),
	})

	// bad dest chain
	orders = append(orders, TestOrder{
		Owner:         user,
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL1,
		DestChainID:   invalidChainID,
		Expenses:      nativeExpense(bi.Wei(1)),
		Calls:         nativeTransferCall(bi.Wei(1), user),
		Deposit:       erc20Deposit(bi.Wei(1), addrs.NomToken),
		ShouldReject:  true,
		RejectReason:  solver.RejectUnsupportedDestChain.String(),
	})

	// unsupported expense token
	orders = append(orders, TestOrder{
		Owner:         user,
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL1,
		DestChainID:   evmchain.IDMockL2,
		Expenses:      unsupportedExpense(bi.Wei(1)),
		Calls:         nativeTransferCall(bi.Wei(1), user),
		Deposit:       erc20Deposit(bi.Wei(1), addrs.NomToken),
		ShouldReject:  true,
		RejectReason:  solver.RejectUnsupportedExpense.String(),
	})

	// invalid expense
	orders = append(orders, TestOrder{
		Owner:         user,
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL1,
		DestChainID:   evmchain.IDMockL2,
		Expenses:      invalidExpenseOutOfBounds(),
		Calls:         nativeTransferCall(bi.Wei(1), user),
		Deposit:       erc20Deposit(bi.Wei(1), addrs.NomToken),
		ShouldReject:  true,
		RejectReason:  solver.RejectInvalidExpense.String(),
	})

	// invalid expense with multiple native expenses
	orders = append(orders, TestOrder{
		Owner:         user,
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL1,
		DestChainID:   evmchain.IDMockL2,
		Expenses:      multipleNativeExpenses(validETHSpend),
		Calls:         nativeTransferCall(validETHSpend, user),
		Deposit:       nativeDeposit(maxETHSpend),
		ShouldReject:  true,
		RejectReason:  solver.RejectInvalidExpense.String(),
	})

	// invalid expense with multiple ERC20 expenses
	orders = append(orders, TestOrder{
		Owner:         user,
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
		Owner:         user,
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL1,
		DestChainID:   evmchain.IDMockL2,
		Expenses:      nativeExpense(validETHSpend),
		Calls:         nativeTransferCall(validETHSpend, user),
		Deposit:       nativeDeposit(maxETHSpend),
		ShouldReject:  false,
	})

	// dest call reverts
	orders = append(orders, TestOrder{
		Owner:         user,
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
		Owner:         user,
		FillDeadline:  time.Now().Add(1 * time.Hour),
		SourceChainID: evmchain.IDMockL2,
		DestChainID:   evmchain.IDMockL1,
		Expenses:      nativeExpense(validETHSpend),
		Calls:         nativeTransferCall(validETHSpend, user),
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
	scl := sclient.New(solverAddr)

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

		checkResp, err := scl.Check(ctx, checkReq)
		if err != nil {
			return errors.Wrap(err, "check request")
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

func testRelayAPI(ctx context.Context, solverAddr string) error {
	scl := sclient.New(solverAddr)

	// Create a minimal gasless order that should be rejected for invalid values
	// This tests that the endpoint exists and handles requests properly
	dummyOrder := bindings.IERC7683GaslessCrossChainOrder{
		OriginSettler: common.Address{}, // Invalid address
		User:          common.Address{}, // Invalid address
		Nonce:         big.NewInt(0),    // Invalid nonce
		OriginChainId: big.NewInt(0),    // Invalid chain ID
		OpenDeadline:  0,                // Invalid deadline
		FillDeadline:  0,                // Invalid deadline
		OrderDataType: common.Hash{},    // Invalid type hash
		OrderData:     []byte{},         // Empty order data
	}

	relayReq := solver.RelayRequest{
		Order:            dummyOrder,
		Signature:        []byte{}, // Empty signature should be rejected
		OriginFillerData: []byte{}, // Empty filler data
	}

	relayResp, err := scl.Relay(ctx, relayReq)
	if err != nil {
		return errors.Wrap(err, "relay request failed unexpectedly")
	}

	// The request should be handled but not successful due to validation failures
	if relayResp.Success {
		return errors.New("relay request unexpectedly succeeded with invalid data")
	}

	// Should have an error indicating why the request failed
	if relayResp.Error == nil {
		return errors.New("relay request failed but no error details provided")
	}

	// Check that we got a meaningful error code
	if relayResp.Error.Code == "" {
		return errors.New("relay error missing error code")
	}

	log.Info(ctx, "Test /relay success, endpoint exists and handles validation")

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

			// solver will have claimed much more than 1 NOM
			// if balance is < 1, rebalancing is working
			if bi.LTE(balance, bi.Ether(1)) {
				return nil
			}
		}
	}
}

func mustTokenByAsset(chainID uint64, asset tokens.Asset) tokens.Token {
	t, ok := tokens.ByAsset(chainID, asset)
	if !ok {
		panic("token not found by asset")
	}

	return t
}

func mustDepositAmount(expenseAmount *big.Int, srcToken tokens.Token, dstToken tokens.Token) *big.Int {
	pricer := tokenpricer.NewDevnetMock()
	price, err := pricer.Price(context.TODO(), srcToken.Asset, dstToken.Asset)
	if err != nil {
		panic(err)
	}

	sprice := solver.Price{
		Price:   price,
		Deposit: srcToken.Asset,
		Expense: dstToken.Asset,
	}

	// Add fee bips
	sprice = sprice.WithFeeBips(30)

	return sprice.ToDeposit(expenseAmount)
}
