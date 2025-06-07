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
	"github.com/omni-network/omni/solver/app"
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

// TestGaslessOrder represents a gasless order for testing.
// Note: The sponsor/relayer who submits the transaction is not tracked here as they just pay gas.
type TestGaslessOrder struct {
	User              common.Address // The depositor (pays via Permit2 signature) - maps to order.user
	Owner             common.Address // The order owner (gets refunds if canceled/rejected) - maps to orderData.owner. If address(0), defaults to User
	Nonce             uint64
	OpenDeadline      time.Time
	FillDeadline      time.Time
	DestChainID       uint64
	SourceChainID     uint64
	Expenses          []solvernet.Expense
	Calls             []solvernet.Call
	Deposit           solvernet.Deposit
	ShouldReject      bool
	RejectReason      string
	RequiresSignature bool
}

func isDepositTokenInvalid(o TestOrder) bool {
	return o.Deposit.Token == invalidTokenAddress
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

	gaslessOrders := makeGaslessOrders() //nolint:contextcheck // Not critical code

	// Test check API with regular orders
	if err := testCheckAPI(ctx, backends, testOrdersToOrderLike(orders), solverAddr); err != nil {
		return errors.Wrap(err, "test check api onchain")
	}

	// Test check API with gasless orders using the new CheckRequest conversion
	if err := testCheckAPIGasless(ctx, backends, gaslessOrders, solverAddr); err != nil {
		return errors.Wrap(err, "test check api gasless")
	}

	if err := testRelayAPI(ctx, solverAddr); err != nil {
		return errors.Wrap(err, "test relay api")
	}

	// Create unified tracker for all orders
	tracker := newOrderTracker()

	// Mint and approve tokens for all regular orders
	if err := mintAndApproveAll(ctx, backends, testOrdersToOrderLike(orders)); err != nil {
		return errors.Wrap(err, "mint all onchain")
	}

	// Open regular orders and add to tracker
	if err := openAllToTracker(ctx, backends, orders, tracker); err != nil {
		return errors.Wrap(err, "open all regular orders")
	}

	// Mint and approve tokens for all gasless orders
	if err := mintAndApproveAll(ctx, backends, testGaslessOrdersToOrderLike(gaslessOrders)); err != nil {
		return errors.Wrap(err, "mint all gasless")
	}

	// Open gasless orders via relay API and add to same tracker
	if err := openAllGaslessToTracker(ctx, backends, gaslessOrders, solverAddr, tracker); err != nil {
		return errors.Wrap(err, "open all gasless orders")
	}

	// Mark all orders as tracked
	tracker.AllTracked()

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

			// if done, wait for solver to rebalance OMNI (bridge L1 back to native)
			if err := waitRebalance(ctx, backends); err != nil {
				return errors.Wrap(err, "wait rebalance")
			}

			log.Info(ctx, "Solver test success")

			return nil
		}
	}
}

// openAllToTracker opens regular orders and adds them directly to the provided tracker.
func openAllToTracker(ctx context.Context, backends ethbackend.Backends, orders []TestOrder, tracker *orderTracker) error {
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
		return errors.Wrap(err, "open onchain orders")
	}

	return nil
}

// openAllGaslessToTracker opens gasless orders via relay API and adds them directly to the provided tracker.
func openAllGaslessToTracker(ctx context.Context, backends ethbackend.Backends, orders []TestGaslessOrder, solverAddr string, tracker *orderTracker) error {
	var eg errgroup.Group
	for _, order := range orders {
		eg.Go(func() error {
			// Skip orders that should be rejected due to validation or backend issues
			if isInvalidExpenseGasless(order) || isSrcChainInvalidGasless(order) || isDepositTokenInvalidGasless(order) ||
				srcAndDestChainAreSameGasless(order) || isInsufficientInventoryGasless(order) || isRelayerErrorGasless(order) || isInvalidOpenDeadlineGasless(order) || isInvalidFillDeadlineGasless(order) {
				return nil
			}

			id, err := openGaslessOrder(ctx, backends, order, solverAddr)
			if err != nil {
				return err
			}

			tracker.TrackOrder(id, order)

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "open gasless orders")
	}

	return nil
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
			swapOrder(user, evmchain.IDMockL1, tokens.ETH, evmchain.IDOmniDevnet, tokens.OMNI),
			swapOrder(user, evmchain.IDMockL1, tokens.OMNI, evmchain.IDMockL2, tokens.ETH),
			swapOrder(user, evmchain.IDMockL1, tokens.USDC, evmchain.IDOmniDevnet, tokens.OMNI),
			swapOrder(user, evmchain.IDMockL1, tokens.OMNI, evmchain.IDMockL1, tokens.ETH), // same chain swap
		)
	}
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
			omniBridgeOrder(user, expense, erc20Deposit(expense, addrs.Token), ""),
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
		Deposit:       erc20Deposit(bi.Wei(1), addrs.Token),
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
		Deposit:       erc20Deposit(bi.Wei(1), addrs.Token),
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
		Deposit:       erc20Deposit(bi.Wei(1), addrs.Token),
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

// makeGaslessOrders creates a collection of gasless test orders similar to makeOrders.
// These orders are designed to test the relayer functionality and gasless order flow.
//
//nolint:maintidx // Just a test function that generates many test cases
func makeGaslessOrders() []TestGaslessOrder {
	user := anvil.DevAccount8()
	var orders []TestGaslessOrder

	// Helper function to generate unique nonces
	nextNonce := uint64(1000) // Start at 1000 to avoid conflicts with any existing tests
	getNonce := func() uint64 {
		nextNonce++
		return nextNonce
	}

	// Craft valid order params
	defaultExpenseAmount := bi.Ether(1)
	defaultSrcToken := mustTokenByAsset(evmchain.IDMockL1, tokens.OMNI) // Source token for deposit
	defaultDstToken := mustTokenByAsset(evmchain.IDMockL2, tokens.ETH)  // Native ETH destination
	defaultDepositAmount := mustDepositAmount(defaultExpenseAmount, defaultSrcToken, defaultDstToken)

	// ERC20 swaps (gasless orders only support ERC20 deposits)
	{
		gaslessSwapOrder := func(addr common.Address, srcChain uint64, srcAsset tokens.Asset, dstChain uint64, dstAsset tokens.Asset) TestGaslessOrder {
			// Use appropriate amount based on destination token decimals
			var amount *big.Int
			if dstAsset == tokens.USDC || dstAsset == tokens.USDT || dstAsset == tokens.USDT0 {
				amount = bi.Dec6(2) // 2 tokens with 6 decimals (2 USDC/USDT)
			} else {
				amount = bi.Ether(2) // 2 tokens with 18 decimals (2 ETH/OMNI)
			}
			srcToken := mustTokenByAsset(srcChain, srcAsset)
			dstToken := mustTokenByAsset(dstChain, dstAsset)
			expense, call := expenseAndCall(amount, dstToken, addr)
			depositAmount := mustDepositAmount(amount, srcToken, dstToken)

			return TestGaslessOrder{
				User:          addr,
				Owner:         addr, // Most common case: user is also the order owner
				Nonce:         getNonce(),
				OpenDeadline:  time.Now().Add(30 * time.Minute), // 30 min to open
				FillDeadline:  time.Now().Add(1 * time.Hour),    // 1 hour to fill
				DestChainID:   dstChain,
				SourceChainID: srcChain,
				Expenses:      expense,
				Calls:         call,
				Deposit: solvernet.Deposit{
					Token:  srcToken.Address,
					Amount: depositAmount,
				},
				RequiresSignature: true,
			}
		}

		orders = append(orders,
			gaslessSwapOrder(user, evmchain.IDMockL1, tokens.OMNI, evmchain.IDOmniDevnet, tokens.OMNI), // OMNI -> OMNI (no fee)
			gaslessSwapOrder(user, evmchain.IDMockL1, tokens.USDC, evmchain.IDMockL2, tokens.ETH),      // USDC -> ETH
			gaslessSwapOrder(user, evmchain.IDMockL1, tokens.OMNI, evmchain.IDMockL1, tokens.USDC),     // same chain swap
		)
	}

	// ERC20 OMNI -> native OMNI orders (bridging)
	{
		gaslessOmniBridgeOrder := func(addr common.Address, expense *big.Int, deposit solvernet.Deposit, rejectReason string) TestGaslessOrder {
			return TestGaslessOrder{
				User:              addr,
				Owner:             addr, // User is also the order owner
				Nonce:             getNonce(),
				OpenDeadline:      time.Now().Add(30 * time.Minute),
				FillDeadline:      time.Now().Add(1 * time.Hour),
				SourceChainID:     evmchain.IDMockL1,
				DestChainID:       evmchain.IDOmniDevnet,
				Expenses:          []solvernet.Expense{},             // NO native expenses - contract rejects expenses with token==address(0)
				Calls:             nativeTransferCall(expense, addr), // Native value goes in call.value
				Deposit:           deposit,
				ShouldReject:      rejectReason != "",
				RejectReason:      rejectReason,
				RequiresSignature: true,
			}
		}

		expense := bi.Ether(5) // Smaller amount for gasless test
		orders = append(orders,
			gaslessOmniBridgeOrder(user, expense, erc20Deposit(expense, addrs.Token), ""),
			gaslessOmniBridgeOrder(user, expense, unsupportedERC20Deposit(expense), solver.RejectUnsupportedDeposit.String()),
		)
	}

	// Invalid gasless order cases
	{
		// Bad source chain
		orders = append(orders, TestGaslessOrder{
			User:              user,
			Owner:             user,
			Nonce:             getNonce(),
			OpenDeadline:      time.Now().Add(30 * time.Minute),
			FillDeadline:      time.Now().Add(1 * time.Hour),
			SourceChainID:     invalidChainID,
			DestChainID:       evmchain.IDMockL1,
			Expenses:          []solvernet.Expense{}, // NO native expenses
			Calls:             nativeTransferCall(bi.Wei(1), user),
			Deposit:           erc20Deposit(bi.Wei(1), zeroAddr),
			ShouldReject:      true,
			RejectReason:      solver.RejectUnsupportedSrcChain.String(),
			RequiresSignature: true,
		})

		// Bad destination chain
		orders = append(orders, TestGaslessOrder{
			User:              user,
			Owner:             user,
			Nonce:             getNonce(),
			OpenDeadline:      time.Now().Add(30 * time.Minute),
			FillDeadline:      time.Now().Add(1 * time.Hour),
			SourceChainID:     evmchain.IDMockL1,
			DestChainID:       invalidChainID,
			Expenses:          []solvernet.Expense{}, // NO native expenses
			Calls:             nativeTransferCall(bi.Wei(1), user),
			Deposit:           erc20Deposit(bi.Wei(1), zeroAddr),
			ShouldReject:      true,
			RejectReason:      solver.RejectUnsupportedDestChain.String(),
			RequiresSignature: true,
		})

		// Invalid expense (out of bounds) - this will be caught by check endpoint
		orders = append(orders, TestGaslessOrder{
			User:              user,
			Owner:             user,
			Nonce:             getNonce(),
			OpenDeadline:      time.Now().Add(30 * time.Minute),
			FillDeadline:      time.Now().Add(1 * time.Hour),
			SourceChainID:     evmchain.IDMockL1,
			DestChainID:       evmchain.IDMockL2,
			Expenses:          []solvernet.Expense{},                   // NO native expenses
			Calls:             nativeTransferCall(bi.Ether(100), user), // Way over max in call value
			Deposit:           erc20Deposit(defaultDepositAmount, addrs.Token),
			ShouldReject:      true,
			RejectReason:      solver.RejectExpenseOverMax.String(), // Check endpoint validation
			RequiresSignature: true,
		})

		// Invalid open deadline (already expired)
		orders = append(orders, TestGaslessOrder{
			User:              user,
			Owner:             user,
			Nonce:             getNonce(),
			OpenDeadline:      time.Now().Add(-1 * time.Hour), // Already expired
			FillDeadline:      time.Now().Add(1 * time.Hour),
			SourceChainID:     evmchain.IDMockL1,
			DestChainID:       evmchain.IDMockL2,
			Expenses:          []solvernet.Expense{}, // NO native expenses
			Calls:             nativeTransferCall(defaultExpenseAmount, user),
			Deposit:           erc20Deposit(defaultDepositAmount, addrs.Token),
			ShouldReject:      true,
			RejectReason:      solver.RejectInvalidDeposit.String(), // Use check API rejection reason
			RequiresSignature: true,
		})

		// Invalid fill deadline (before open deadline)
		orders = append(orders, TestGaslessOrder{
			User:              user,
			Owner:             user,
			Nonce:             getNonce(),
			OpenDeadline:      time.Now().Add(1 * time.Hour),
			FillDeadline:      time.Now().Add(30 * time.Minute), // Before open deadline
			SourceChainID:     evmchain.IDMockL1,
			DestChainID:       evmchain.IDMockL2,
			Expenses:          []solvernet.Expense{}, // NO native expenses
			Calls:             nativeTransferCall(defaultExpenseAmount, user),
			Deposit:           erc20Deposit(defaultDepositAmount, addrs.Token),
			ShouldReject:      true,
			RejectReason:      solver.RejectInvalidDeposit.String(), // Use check API rejection reason
			RequiresSignature: true,
		})

		// Unsupported expense token (matching onchain test)
		orders = append(orders, TestGaslessOrder{
			User:              user,
			Owner:             user,
			Nonce:             getNonce(),
			OpenDeadline:      time.Now().Add(30 * time.Minute),
			FillDeadline:      time.Now().Add(1 * time.Hour),
			SourceChainID:     evmchain.IDMockL1,
			DestChainID:       evmchain.IDMockL2,
			Expenses:          unsupportedExpense(bi.Wei(1)),
			Calls:             nativeTransferCall(bi.Wei(1), user),
			Deposit:           erc20Deposit(defaultDepositAmount, addrs.Token),
			ShouldReject:      true,
			RejectReason:      solver.RejectUnsupportedExpense.String(),
			RequiresSignature: true,
		})

		// Invalid native expense (should be rejected by contract - expenses cannot have token==address(0))
		orders = append(orders, TestGaslessOrder{
			User:              user,
			Owner:             user,
			Nonce:             getNonce(),
			OpenDeadline:      time.Now().Add(30 * time.Minute),
			FillDeadline:      time.Now().Add(1 * time.Hour),
			SourceChainID:     evmchain.IDMockL1,
			DestChainID:       evmchain.IDMockL2,
			Expenses:          nativeExpense(validETHSpend), // Contract rejects native expenses
			Calls:             nativeTransferCall(validETHSpend, user),
			Deposit:           erc20Deposit(defaultDepositAmount, addrs.Token),
			ShouldReject:      true,
			RejectReason:      solver.RejectInvalidExpense.String(),
			RequiresSignature: true,
		})

		// Multiple native expenses (should be rejected - same reason as above)
		orders = append(orders, TestGaslessOrder{
			User:              user,
			Owner:             user,
			Nonce:             getNonce(),
			OpenDeadline:      time.Now().Add(30 * time.Minute),
			FillDeadline:      time.Now().Add(1 * time.Hour),
			SourceChainID:     evmchain.IDMockL1,
			DestChainID:       evmchain.IDMockL2,
			Expenses:          multipleNativeExpenses(validETHSpend), // Contract rejects native expenses
			Calls:             nativeTransferCall(validETHSpend, user),
			Deposit:           erc20Deposit(defaultDepositAmount, addrs.Token),
			ShouldReject:      true,
			RejectReason:      solver.RejectInvalidExpense.String(),
			RequiresSignature: true,
		})

		// Multiple ERC20 expenses (valid in contract, but solver may reject)
		dstToken := mustTokenByAsset(evmchain.IDMockL2, tokens.USDC)   // Valid token on destination chain
		validAmount := bi.Dec6(2)                                      // 2 USDC (appropriate amount for 6-decimal token)
		validExpense, _ := expenseAndCall(validAmount, dstToken, user) // Get valid single expense
		orders = append(orders, TestGaslessOrder{
			User:              user,
			Owner:             user,
			Nonce:             getNonce(),
			OpenDeadline:      time.Now().Add(30 * time.Minute),
			FillDeadline:      time.Now().Add(1 * time.Hour),
			SourceChainID:     evmchain.IDMockL1,
			DestChainID:       evmchain.IDMockL2,
			Expenses:          append(validExpense, validExpense[0]), // Duplicate the valid expense to create multiple
			Calls:             nativeTransferCall(validETHSpend, user),
			Deposit:           erc20Deposit(defaultDepositAmount, addrs.Token),
			ShouldReject:      true,
			RejectReason:      solver.RejectInvalidExpense.String(),
			RequiresSignature: true,
		})

		// Dest call reverts (matching onchain test)
		orders = append(orders, TestGaslessOrder{
			User:              user,
			Owner:             user,
			Nonce:             getNonce(),
			OpenDeadline:      time.Now().Add(30 * time.Minute),
			FillDeadline:      time.Now().Add(1 * time.Hour),
			SourceChainID:     evmchain.IDMockL1,
			DestChainID:       evmchain.IDMockL2,
			Expenses:          []solvernet.Expense{},
			Calls:             contractCallWithInvalidCallData(),
			Deposit:           erc20Deposit(defaultDepositAmount, addrs.Token),
			ShouldReject:      true,
			RejectReason:      solver.RejectDestCallReverts.String(),
			RequiresSignature: true,
		})

		// Insufficient inventory for native expense
		srcTokenL2 := mustTokenByAsset(evmchain.IDMockL2, tokens.USDC) // Valid token on source chain (MockL2)
		largeAmount := bi.Ether(2.5)                                   // Amount larger than solver inventory but below max limit (3 ETH) to trigger insufficient inventory
		largeDepositAmount := mustDepositAmount(largeAmount, srcTokenL2, mustTokenByAsset(evmchain.IDMockL1, tokens.ETH))
		orders = append(orders, TestGaslessOrder{
			User:              user,
			Owner:             user,
			Nonce:             getNonce(),
			OpenDeadline:      time.Now().Add(30 * time.Minute),
			FillDeadline:      time.Now().Add(1 * time.Hour),
			SourceChainID:     evmchain.IDMockL2,
			DestChainID:       evmchain.IDMockL1,
			Expenses:          []solvernet.Expense{},
			Calls:             nativeTransferCall(largeAmount, user),                // Large native value in call to exceed inventory
			Deposit:           erc20Deposit(largeDepositAmount, srcTokenL2.Address), // Use token that exists on source chain
			ShouldReject:      true,
			RejectReason:      solver.RejectInsufficientInventory.String(),
			RequiresSignature: true,
		})
	}

	// Valid ERC20 transfer orders
	{
		erc20TransferOrder := func(addr common.Address, srcToken, dstToken tokens.Token, expense *big.Int, rejectReason string) TestGaslessOrder {
			depositAmount := mustDepositAmount(expense, srcToken, dstToken)
			return TestGaslessOrder{
				User:          addr,
				Owner:         addr,
				Nonce:         getNonce(),
				OpenDeadline:  time.Now().Add(30 * time.Minute),
				FillDeadline:  time.Now().Add(1 * time.Hour),
				SourceChainID: evmchain.IDMockL1,
				DestChainID:   evmchain.IDMockL2,
				Expenses:      []solvernet.Expense{{Token: dstToken.Address, Amount: expense}},
				Calls: func() []solvernet.Call {
					call, err := erc20Call(dstToken, expense, addr)
					if err != nil {
						panic(err)
					}

					return call
				}(),
				Deposit: solvernet.Deposit{
					Token:  srcToken.Address,
					Amount: depositAmount,
				},
				ShouldReject:      rejectReason != "",
				RejectReason:      rejectReason,
				RequiresSignature: true,
			}
		}

		srcToken := mustTokenByAsset(evmchain.IDMockL1, tokens.USDC)
		dstToken := mustTokenByAsset(evmchain.IDMockL2, tokens.USDC)
		validExpense := bi.Dec6(1) // 1 USDC (6 decimals)

		orders = append(orders,
			erc20TransferOrder(user, srcToken, dstToken, validExpense, ""),
		)
	}

	// Sponsored transaction example (User != Owner)
	{
		sponsor := anvil.DevAccount9() // Different account as the sponsor/owner

		// Sponsored order where user pays deposit but sponsor owns the order
		orders = append(orders, TestGaslessOrder{
			User:              user,    // User pays the deposit via Permit2
			Owner:             sponsor, // Sponsor owns the order and gets refunds
			Nonce:             getNonce(),
			OpenDeadline:      time.Now().Add(30 * time.Minute),
			FillDeadline:      time.Now().Add(1 * time.Hour),
			SourceChainID:     evmchain.IDMockL1,
			DestChainID:       evmchain.IDMockL2,
			Expenses:          nativeExpense(defaultExpenseAmount),
			Calls:             nativeTransferCall(defaultExpenseAmount, sponsor), // Transfer to sponsor
			Deposit:           erc20Deposit(defaultDepositAmount, addrs.Token),
			ShouldReject:      false,
			RejectReason:      "",
			RequiresSignature: true,
		})
	}

	// Test address(0) owner fallback behavior
	{
		// When Owner is address(0), the contract should default it to User
		// This tests the fallback logic in _validateOrderData
		orders = append(orders, TestGaslessOrder{
			User:              user,
			Owner:             zeroAddr, // address(0) - should default to User in contract
			Nonce:             getNonce(),
			OpenDeadline:      time.Now().Add(30 * time.Minute),
			FillDeadline:      time.Now().Add(1 * time.Hour),
			SourceChainID:     evmchain.IDMockL1,
			DestChainID:       evmchain.IDMockL2,
			Expenses:          nativeExpense(defaultExpenseAmount),
			Calls:             nativeTransferCall(defaultExpenseAmount, user), // Transfer to user since they'll become owner
			Deposit:           erc20Deposit(defaultDepositAmount, addrs.Token),
			ShouldReject:      false,
			RejectReason:      "",
			RequiresSignature: true,
		})
	}

	return orders
}

func openOrder(ctx context.Context, backends ethbackend.Backends, order TestOrder) ([32]byte, error) {
	return solvernet.OpenOrder(ctx, netconf.Devnet, order.SourceChainID, backends, order.Owner, bindings.SolverNetOrderData{
		DestChainId: order.DestChainID,
		Expenses:    solvernet.FilterNativeExpenses(order.Expenses),
		Calls:       solvernet.CallsToBindings(order.Calls),
		Deposit:     order.Deposit,
	}, solvernet.WithFillDeadline(order.FillDeadline))
}

func testCheckAPI(ctx context.Context, backends ethbackend.Backends, orders []OrderLike, solverAddr string) error {
	scl := sclient.New(solverAddr)

	for i, order := range orders {
		// If this order requires balance draining, do it before test logic.
		if isInsufficientInventoryForOrderLike(order) {
			// Drain solver native balance.
			if err := setSolverAccountNativeBalance(ctx, order.GetDestChainID(), backends, bi.Zero()); err != nil {
				return errors.Wrap(err, "drain solver account failed")
			}
		}

		checkReq := orderLikeToCheckRequest(order)

		checkResp, err := scl.Check(ctx, checkReq)
		if err != nil {
			return errors.Wrap(err, "check request")
		}

		if checkResp.Rejected != order.GetShouldReject() {
			return errors.New("unexpected rejection",
				"expected", order.GetShouldReject(),
				"actual", checkResp.Rejected,
				"reason", checkResp.RejectReason,
				"description", checkResp.RejectDescription,
				"order_idx", i,
			)
		}

		if checkResp.RejectReason != order.GetRejectReason() {
			return errors.New("unexpected reject reason", "expected", order.GetRejectReason(), "actual", checkResp.RejectReason)
		}

		if checkResp.Accepted && order.GetShouldReject() {
			return errors.New("accepted but should reject")
		}

		// Refund solver native balance after test logic.
		if isInsufficientInventoryForOrderLike(order) {
			eth1m := bi.Ether(1_000_000)
			if err := setSolverAccountNativeBalance(ctx, order.GetDestChainID(), backends, eth1m); err != nil {
				return errors.Wrap(err, "refund solver account failed")
			}
		}
	}

	log.Info(ctx, "Test /check success")

	return nil
}

// testCheckAPIGasless tests the check API with gasless orders by converting them to CheckRequest format.
func testCheckAPIGasless(ctx context.Context, backends ethbackend.Backends, gaslessOrders []TestGaslessOrder, solverAddr string) error {
	scl := sclient.New(solverAddr)

	for i, order := range gaslessOrders {
		log.Info(ctx, "Checking gasless order via check API", "order_idx", i)

		// If this order requires balance draining, do it before test logic.
		if isInsufficientInventoryGasless(order) {
			// Drain solver native balance.
			if err := setSolverAccountNativeBalance(ctx, order.DestChainID, backends, bi.Zero()); err != nil {
				return errors.Wrap(err, "drain solver account failed")
			}
		}

		// Create order data in the same format as used in openGaslessOrder
		orderData := bindings.SolverNetOrderData{
			Owner:       order.Owner,
			DestChainId: order.DestChainID,
			Deposit:     order.Deposit,
			Calls:       solvernet.CallsToBindings(order.Calls),
			Expenses:    solvernet.FilterNativeExpenses(order.Expenses), // Filter native expenses for gasless orders
		}

		// Special case: for testing native expense validation, don't filter
		if order.RejectReason == solver.RejectInvalidExpense.String() {
			// Check if this is testing native expense validation specifically
			hasNativeExpenses := false
			for _, exp := range order.Expenses {
				if exp.Token == (common.Address{}) {
					hasNativeExpenses = true
					break
				}
			}
			if hasNativeExpenses {
				orderData.Expenses = order.Expenses // Don't filter so solver can validate
			}
		}

		// Create gasless order structure (simulating what OpenGaslessOrder would create)
		gaslessOrder := bindings.IERC7683GaslessCrossChainOrder{
			User:          order.User,
			Nonce:         new(big.Int).SetUint64(order.Nonce),
			OpenDeadline:  uint32(order.OpenDeadline.Unix()),
			FillDeadline:  uint32(order.FillDeadline.Unix()),
			OriginChainId: new(big.Int).SetUint64(order.SourceChainID),
			// We don't need the other fields for validation
		}

		// Convert to check request using the new helper function
		req, err := solver.CheckRequestFromGaslessOrder(gaslessOrder, orderData)
		if err != nil {
			return errors.Wrap(err, "gasless order to check request", "order_idx", i)
		}

		// Add debug flag
		req.Debug = true

		// Make the check request
		checkResp, err := scl.Check(ctx, req)
		if err != nil {
			return errors.Wrap(err, "check api request", "order_idx", i)
		}

		// Validate response matches expected behavior
		if order.ShouldReject {
			if checkResp.Accepted {
				return errors.New("unexpected acceptance", "expected", false, "actual", true, "order_idx", i)
			}
			if checkResp.RejectReason != order.GetRejectReason() {
				return errors.New("unexpected reject reason", "expected", order.GetRejectReason(), "actual", checkResp.RejectReason, "order_idx", i)
			}
		} else {
			if checkResp.Rejected {
				return errors.New("unexpected rejection", "expected", false, "actual", true, "reason", checkResp.RejectReason, "description", checkResp.RejectDescription, "order_idx", i)
			}
		}

		// Refund solver native balance after test logic.
		if isInsufficientInventoryGasless(order) {
			eth1m := bi.Ether(1_000_000)
			if err := setSolverAccountNativeBalance(ctx, order.DestChainID, backends, eth1m); err != nil {
				return errors.Wrap(err, "refund solver account failed")
			}
		}
	}

	log.Info(ctx, "Test /check gasless success")

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

			// solver will have claimed much more than 1 OMNI
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

// Helper functions for gasless order skip logic.
func isInvalidExpenseGasless(o TestGaslessOrder) bool {
	// Skip gasless orders with native expenses
	for _, expense := range o.Expenses {
		if expense.Token == zeroAddr {
			return true
		}
	}

	return false
}

func isSrcChainInvalidGasless(o TestGaslessOrder) bool {
	return o.SourceChainID == invalidChainID
}

func isDepositTokenInvalidGasless(o TestGaslessOrder) bool {
	return o.Deposit.Token == invalidTokenAddress || o.Deposit.Token == zeroAddr // Also skip native deposits for gasless orders
}

func srcAndDestChainAreSameGasless(o TestGaslessOrder) bool {
	return o.SourceChainID == o.DestChainID
}

func isInsufficientInventoryGasless(o TestGaslessOrder) bool {
	return o.RejectReason == solver.RejectInsufficientInventory.String()
}

func isRelayerErrorGasless(o TestGaslessOrder) bool {
	return o.RejectReason == app.RelayErrorMissingSignature ||
		o.RejectReason == app.RelayErrorInvalidOrder ||
		o.RejectReason == app.RelayErrorUnsupportedChain ||
		o.RejectReason == app.RelayErrorInvalidOriginSettler
}

func isInvalidOpenDeadlineGasless(o TestGaslessOrder) bool {
	// Skip gasless orders with open deadlines in the past
	return o.OpenDeadline.Before(time.Now())
}

func isInvalidFillDeadlineGasless(o TestGaslessOrder) bool {
	// Skip gasless orders where fill deadline is not after open deadline
	return !o.FillDeadline.After(o.OpenDeadline)
}

// openGaslessOrder opens a single gasless order via the relay API.
func openGaslessOrder(ctx context.Context, backends ethbackend.Backends, order TestGaslessOrder, solverAddr string) ([32]byte, error) {
	// Get the user's private key
	userKey, ok := anvil.PrivateKey(order.User)
	if !ok {
		return [32]byte{}, errors.New("user private key not found", "user", order.User.Hex())
	}

	// Create order data
	orderData := bindings.SolverNetOrderData{
		Owner:       order.Owner,
		DestChainId: order.DestChainID,
		Deposit:     order.Deposit,
		Calls:       solvernet.CallsToBindings(order.Calls),
		Expenses:    solvernet.FilterNativeExpenses(order.Expenses), // Filter native expenses for gasless orders
	}

	// Get the actual solver Ethereum address (not the URL)
	solverEthAddr := eoa.MustAddress(netconf.Devnet, eoa.RoleSolver)

	// Prepare gasless order and signature
	gaslessOrder, signature, err := solvernet.OpenGaslessOrder(
		ctx,
		netconf.Devnet,
		order.SourceChainID,
		backends,
		userKey,
		orderData,
		solvernet.WithGaslessOpenDeadline(order.OpenDeadline),
		solvernet.WithGaslessFillDeadline(order.FillDeadline),
		solvernet.WithGaslessNonce(new(big.Int).SetUint64(order.Nonce)),
		solvernet.WithSolverAddr(solverEthAddr),
	)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "prepare gasless order")
	}

	// Skip signature if not required (for testing error cases)
	if !order.RequiresSignature {
		signature = []byte{} // Empty signature should be rejected by relayer
	}

	// Create relay request
	relayReq := solver.RelayRequest{
		Order:            gaslessOrder,
		Signature:        signature,
		OriginFillerData: []byte{}, // Empty for now
	}

	// Submit via relay API (solverAddr is the URL for the API)
	client := sclient.New(solverAddr)
	relayResp, err := client.Relay(ctx, relayReq)
	if err != nil {
		return [32]byte{}, errors.Wrap(err, "relay request failed")
	}

	// Check if the relay was successful
	if !relayResp.Success {
		if relayResp.Error != nil {
			return [32]byte{}, errors.New("relay request rejected",
				"code", relayResp.Error.Code,
				"message", relayResp.Error.Message,
				"description", relayResp.Error.Description)
		}

		return [32]byte{}, errors.New("relay request failed with no error details")
	}

	// Extract order ID from response
	return relayResp.OrderID, nil
}
