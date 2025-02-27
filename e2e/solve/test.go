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
	Expenses      solvernet.Expenses
	Calls         solvernet.Calls
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

	if err = testCheckAPI(ctx, orders); err != nil {
		return errors.Wrap(err, "test check api")
	}

	tracker, err := openAll(ctx, backends, orders)
	if err != nil {
		return errors.Wrap(err, "open all")
	}

	xprov := xprovider.New(network, backends.Clients(), nil)

	proc := func(ctx context.Context, _ uint64, logs []types.Log) error {
		for _, l := range logs {
			event, ok := solvernet.EventByTopic(l.Topics[0])
			if !ok {
				return errors.New("unknown event", "topic", l.Topics[0])
			}

			id := solvernet.OrderID(l.Topics[1])

			log.Info(ctx, "Order updated", "status", event.Status, "order", id)

			tracker.setStatus(id, event.Status)
		}

		return nil
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// start event streams
	errChan := make(chan error, 1)
	for _, chain := range network.EVMChains() {
		go func() {
			req := xchain.EventLogsReq{
				ChainID:       chain.ID,
				ConfLevel:     xchain.ConfLatest,
				Height:        1,
				FilterAddress: addrs.SolverNetInbox,
				FilterTopics:  solvernet.AllEventTopics(),
			}

			err := xprov.StreamEventLogs(ctx, req, proc)
			if err != nil {
				errChan <- err
			}
		}()
	}

	// wait done
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case err := <-errChan:
			return errors.Wrap(err, "stream event logs")
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "context done")
		case <-ticker.C:
			done, err := tracker.done()
			if err != nil {
				log.Error(ctx, "Solver test error", err)
				return err
			}

			if done {
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
			Deposit:       unsupportedDeposit(depositAmt),
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
		Deposit:       erc20Deposit(big.NewInt(1), common.Address{}),
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
		Expenses:      invalidExpense(),
		Calls:         nativeTransferCall(big.NewInt(1), users[0]),
		Deposit:       erc20Deposit(big.NewInt(1), addrs.Token),
		ShouldReject:  true,
		RejectReason:  solver.RejectInvalidExpense.String(),
	})

	// TODO: more tests orders (different rejection cases, valid orders, etc)

	return orders
}

func openAll(ctx context.Context, backends ethbackend.Backends, orders []TestOrder) (*orderTracker, error) {
	tracker := newOrderTracker()

	var eg errgroup.Group
	for _, order := range orders {
		eg.Go(func() error {
			if isSrcChainInvalid(order) || isDepositTokenInvalid(order) || srcAndDestChainAreSame(order) {
				return nil
			}

			id, err := openOrder(ctx, backends, order)
			if err != nil {
				return err
			}

			tracker.add(id, order)

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, errors.Wrap(err, "wait group")
	}

	return tracker, nil
}

func openOrder(ctx context.Context, backends ethbackend.Backends, order TestOrder) ([32]byte, error) {
	return solvernet.OpenOrder(ctx, netconf.Devnet, order.SourceChainID, backends, order.Owner, bindings.SolverNetOrderData{
		DestChainId: order.DestChainID,
		Expenses:    order.Expenses.NoNative(),
		Calls:       order.Calls.ToBindings(),
		Deposit:     order.Deposit,
	}, solvernet.WithFillDeadline(order.FillDeadline))
}

func testCheckAPI(ctx context.Context, orders []TestOrder) error {
	const url = "http://localhost:26661/api/v1/check"

	var eg errgroup.Group
	for i, order := range orders {
		eg.Go(func() error {
			checkReq := solver.CheckRequest{
				FillDeadline:       uint32(order.FillDeadline.Unix()), //nolint:gosec // this is fine for tests
				SourceChainID:      order.SourceChainID,
				DestinationChainID: order.DestChainID,
				Expenses:           solver.ToJSONExpenses(order.Expenses),
				Calls:              solver.ToJSONCalls(order.Calls),
				Deposit:            solver.ToJSONDeposit(order.Deposit),
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

			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				return errors.New("bad status", "status", resp.StatusCode)
			}

			var checkResp solver.CheckResponse
			if err = json.NewDecoder(resp.Body).Decode(&checkResp); err != nil {
				return errors.Wrap(err, "decode response")
			}

			if checkResp.Rejected != order.ShouldReject {
				return errors.New("unexpected rejection",
					"expected", order.ShouldReject,
					"actual", checkResp.Rejected,
					"reason", checkResp.RejectReason,
					"order_idx", i,
				)
			}

			if checkResp.RejectReason != order.RejectReason {
				return errors.New("unexpected reject reason", "expected", order.RejectReason, "actual", checkResp.RejectReason)
			}

			if checkResp.Accepted && order.ShouldReject {
				return errors.New("accepted but should reject")
			}

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		log.Error(ctx, "Test /check error", err)
		return errors.Wrap(err, "wait checks")
	}

	log.Info(ctx, "Test /check success")

	return nil
}
