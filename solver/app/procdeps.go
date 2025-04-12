package app

import (
	"context"
	"log/slog"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	stypes "github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// procDeps abstracts dependencies for the event processor allowed simplified testing.
type procDeps struct {
	ParseID      func(log types.Log) (OrderID, error)
	GetOrder     func(ctx context.Context, chainID uint64, id OrderID) (Order, bool, error)
	SetCursor    func(ctx context.Context, chainID uint64, height uint64) error
	ShouldReject func(ctx context.Context, order Order) (stypes.RejectReason, bool, error)
	DidFill      func(ctx context.Context, order Order) (bool, error)

	Accept func(ctx context.Context, order Order) error
	Reject func(ctx context.Context, order Order, reason stypes.RejectReason) error
	Fill   func(ctx context.Context, order Order) error
	Claim  func(ctx context.Context, order Order) error

	// Monitoring helpers
	ProcessorName string
	TargetName    func(PendingData) string
	ChainName     func(chainID uint64) string
	InstrumentAge func(ctx context.Context, chainID uint64, height uint64, order Order) slog.Attr
}

func newClaimer(
	networkID netconf.ID,
	inboxContracts map[uint64]*bindings.SolverNetInbox,
	backends ethbackend.Backends,
	solverAddr common.Address,
	pnl updatePnLFunc,
) func(ctx context.Context, order Order) error {
	return func(ctx context.Context, order Order) error {
		inbox, ok := inboxContracts[order.SourceChainID]
		if !ok {
			return errors.New("unknown chain")
		}

		backend, err := backends.Backend(order.SourceChainID)
		if err != nil {
			return err
		}

		txOpts, err := backend.BindOpts(ctx, solverAddr)
		if err != nil {
			return err
		}

		claimant, ok, err := getClaimant(networkID, order)
		if err != nil {
			return errors.Wrap(err, "get claimant")
		} else if !ok {
			claimant = solverAddr
		}

		// Claim to solver address for now
		// TODO: consider claiming to hot / cold funding wallet
		tx, err := inbox.Claim(txOpts, order.ID, claimant)
		if err != nil {
			return errors.Wrap(err, "claim order")
		}
		rec, err := backend.WaitConfirmed(ctx, tx)
		if ok, err2 := maybeDebugRevert(ctx, backend, solverAddr, tx, rec); ok {
			return errors.Wrap(err2, "claim reverted") // Best effort improving of revert errors
		} else if err != nil {
			return errors.Wrap(err, "wait confirmed")
		}

		return pnl(ctx, order, rec, "Inbox:Claim")
	}
}

func newFiller(
	outboxContracts map[uint64]*bindings.SolverNetOutbox,
	backends ethbackend.Backends,
	solverAddr, outboxAddr common.Address,
	pnl filledPnLFunc,
) func(ctx context.Context, order Order) error {
	return func(ctx context.Context, order Order) error {
		pendingData, err := order.PendingData()
		if err != nil {
			return err
		}

		if pendingData.DestinationSettler != outboxAddr {
			return errors.New("destination settler mismatch [BUG] ", "got", pendingData.DestinationSettler.Hex(), "expected", outboxAddr.Hex())
		}

		destChainID := pendingData.DestinationChainID
		outbox, ok := outboxContracts[destChainID]
		if !ok {
			return errors.New("unknown chain")
		}

		backend, err := backends.Backend(destChainID)
		if err != nil {
			return err
		}

		callOpts := &bind.CallOpts{Context: ctx}
		txOpts, err := backend.BindOpts(ctx, solverAddr)
		if err != nil {
			return err
		}

		if ok, err := outbox.DidFill(callOpts, order.ID, pendingData.FillOriginData); err != nil {
			return errors.Wrap(err, "did fill")
		} else if ok {
			log.Info(ctx, "Skipping already filled order", "order_id", order.ID)
			return nil
		}

		nativeValue := bi.Zero()
		for _, output := range pendingData.MaxSpent {
			if output.ChainId.Uint64() != destChainID {
				// We error on this case for now, as our contracts only allow single dest chain orders
				// ERC7683 allows for orders with multiple destination chains, so continue-ing here
				// would also be appropriate.
				return errors.New("destination chain mismatch [BUG] ")
			}

			// zero token address means native token
			if output.Token == [32]byte{} {
				nativeValue.Add(nativeValue, output.Amount)
				continue
			}

			tknAddr, err := toEthAddr(output.Token)
			if err != nil {
				return errors.Wrap(err, "output token address")
			}

			tkn, ok := tokens.ByAddress(destChainID, tknAddr)
			if !ok || !IsSupportedToken(tkn) {
				return errors.New("unsupported token, should have been rejected [BUG]", "addr", tknAddr.Hex(), "chain_id", destChainID)
			}

			if ok, err = isAppproved(ctx, tknAddr, backend, solverAddr, outboxAddr, output.Amount); err != nil {
				return errors.Wrap(err, "is approved")
			} else if !ok {
				return errors.New("outbox not approved to spend token",
					"token", tkn.Symbol,
					"chain_id", destChainID,
					"addr", tknAddr.Hex(),
					"amount", output.Amount,
				)
			}
		}

		// xcall fee
		fee, err := outbox.FillFee(callOpts, pendingData.FillOriginData)
		if err != nil {
			return errors.Wrap(err, "get fulfill fee")
		}

		txOpts.Value = bi.Add(nativeValue, fee)
		fillerData := []byte{} // fillerData is optional ERC7683 custom filler specific data, unused in our contracts
		tx, err := outbox.Fill(txOpts, order.ID, pendingData.FillOriginData, fillerData)
		if err != nil {
			return errors.Wrap(err, "fill order", "custom", solvernet.DetectCustomError(err))
		}
		rec, err := backend.WaitConfirmed(ctx, tx)
		if ok, err2 := maybeDebugRevert(ctx, backend, solverAddr, tx, rec); ok {
			return errors.Wrap(err2, "fill reverted") // Best effort improving of revert errors
		} else if err != nil {
			return errors.Wrap(err, "wait confirmed")
		}

		if ok, err := outbox.DidFill(callOpts, order.ID, pendingData.FillOriginData); err != nil {
			return errors.Wrap(err, "did fill")
		} else if !ok {
			return errors.New("fill failed [BUG]")
		}

		return pnl(ctx, order, rec)
	}
}

func newRejector(
	inboxContracts map[uint64]*bindings.SolverNetInbox,
	backends ethbackend.Backends,
	solverAddr common.Address,
	pnl updatePnLFunc,
) func(ctx context.Context, order Order, reason stypes.RejectReason) error {
	return func(ctx context.Context, order Order, reason stypes.RejectReason) error {
		inbox, ok := inboxContracts[order.SourceChainID]
		if !ok {
			return errors.New("unknown chain")
		}

		backend, err := backends.Backend(order.SourceChainID)
		if err != nil {
			return err
		}

		// Ensure latest on-chain order is still pending
		if latest, err := inbox.GetOrder(&bind.CallOpts{Context: ctx}, order.ID); err != nil {
			return errors.Wrap(err, "get order")
		} else if latest.State.Status != solvernet.StatusPending.Uint8() {
			return errors.New("order status not pending anymore")
		}

		txOpts, err := backend.BindOpts(ctx, solverAddr)
		if err != nil {
			return err
		}

		tx, err := inbox.Reject(txOpts, order.ID, uint8(reason))
		if err != nil {
			return errors.Wrap(err, "reject order", "custom", solvernet.DetectCustomError(err))
		}
		rec, err := backend.WaitConfirmed(ctx, tx)
		if ok, err2 := maybeDebugRevert(ctx, backend, solverAddr, tx, rec); ok {
			return errors.Wrap(err2, "reject reverted") // Best effort improving of revert errors
		} else if err != nil {
			return errors.Wrap(err, "wait confirmed")
		}

		return pnl(ctx, order, rec, "Inbox:Reject")
	}
}

func newDidFiller(outboxContracts map[uint64]*bindings.SolverNetOutbox) func(ctx context.Context, order Order) (bool, error) {
	return func(ctx context.Context, order Order) (bool, error) {
		pendingData, err := order.PendingData()
		if err != nil {
			return false, err
		}

		outbox, ok := outboxContracts[pendingData.DestinationChainID]
		if !ok {
			return false, errors.New("unknown chain")
		}

		filled, err := outbox.DidFill(&bind.CallOpts{Context: ctx}, order.ID, pendingData.FillOriginData)
		if err != nil {
			return false, errors.Wrap(err, "did fill")
		}

		return filled, nil
	}
}

func newIDParser() func(log types.Log) (OrderID, error) {
	return func(log types.Log) (OrderID, error) {
		event, ok := solvernet.EventByTopic(log.Topics[0])
		if !ok {
			return OrderID{}, errors.New("unknown event")
		}

		return event.ParseID(log)
	}
}

func newOrderGetter(inboxContracts map[uint64]*bindings.SolverNetInbox) func(ctx context.Context, chainID uint64, id OrderID) (Order, bool, error) {
	return func(ctx context.Context, chainID uint64, id OrderID) (Order, bool, error) {
		inbox, ok := inboxContracts[chainID]
		if !ok {
			return Order{}, false, errors.New("unknown chain")
		}

		o, err := inbox.GetOrder(&bind.CallOpts{Context: ctx}, id)
		if err != nil {
			return Order{}, false, errors.Wrap(err, "get order")
		}

		// not found
		if o.Resolved.OrderId == [32]byte{} {
			return Order{}, false, nil
		}

		if o.Resolved.OrderId != id {
			return Order{}, false, errors.New("[BUG] order ID mismatch")
		}

		order, err := newOrder(o.Resolved, o.State, o.Offset)
		if err != nil {
			return Order{}, false, errors.Wrap(err, "new order")
		}

		return order, true, nil
	}
}
