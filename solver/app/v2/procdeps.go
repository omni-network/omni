//nolint:unused // This package is a work in progress.
package appv2

import (
	"context"
	"math/big"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// procDeps abstracts dependencies for the event processor allowed simplified testing.
type procDeps struct {
	ParseID      func(chainID uint64, log types.Log) (OrderID, error)
	GetOrder     func(ctx context.Context, chainID uint64, id OrderID) (Order, bool, error)
	ShouldReject func(ctx context.Context, chainID uint64, order Order) (rejectReason, bool, error)
	SetCursor    func(ctx context.Context, chainID uint64, height uint64) error

	Accept func(ctx context.Context, chainID uint64, order Order) error
	Reject func(ctx context.Context, chainID uint64, order Order, reason rejectReason) error
	Fill   func(ctx context.Context, chainID uint64, order Order) error
	Claim  func(ctx context.Context, chainID uint64, order Order) error

	// Monitoring helpers
	TargetName func(Order) string
	ChainName  func(chainID uint64) string
}

func newClaimer(
	inboxContracts map[uint64]*bindings.SolverNetInbox,
	backends ethbackend.Backends,
	solverAddr common.Address,
) func(ctx context.Context, chainID uint64, order Order) error {
	return func(ctx context.Context, chainID uint64, order Order) error {
		inbox, ok := inboxContracts[chainID]
		if !ok {
			return errors.New("unknown chain")
		}

		backend, err := backends.Backend(chainID)
		if err != nil {
			return err
		}

		txOpts, err := backend.BindOpts(ctx, solverAddr)
		if err != nil {
			return err
		}

		// Claim to solver address for now
		// TODO: consider claiming to hot / cold funding wallet
		tx, err := inbox.Claim(txOpts, order.ID, solverAddr)
		if err != nil {
			return errors.Wrap(err, "claim order")
		} else if _, err := backend.WaitMined(ctx, tx); err != nil {
			return errors.Wrap(err, "wait mined")
		}

		return nil
	}
}

func newFiller(
	_ netconf.ID,
	_ map[uint64]*bindings.SolveOutbox,
	_ ethbackend.Backends,
	_, _ common.Address,
) func(ctx context.Context, srcChainID uint64, order Order) error {
	return func(_ context.Context, _ uint64, _ Order) error {
		// TODO
		return nil
	}
}

func detectCustomError(custom error) string {
	contracts := map[string]*bind.MetaData{
		"inbox":      bindings.SolveInboxMetaData,
		"outbox":     bindings.SolveOutboxMetaData,
		"mock_vault": bindings.MockVaultMetaData,
		"mock_token": bindings.MockTokenMetaData,
	}

	for name, contract := range contracts {
		abi, err := contract.GetAbi()
		if err != nil {
			return "BUG"
		}
		for n, e := range abi.Errors {
			if strings.Contains(custom.Error(), e.ID.Hex()[:10]) {
				return name + "::" + n
			}
		}
	}

	return unknown
}

func checkAllowedCall(ctx context.Context, outbox *bindings.SolveOutbox, call bindings.SolveCall) error {
	callOpts := &bind.CallOpts{Context: ctx}

	if len(call.Data) < 4 {
		return errors.New("invalid call data")
	}

	callMethodID, err := cast.Array4(call.Data[:4])
	if err != nil {
		return err
	}

	allowed, err := outbox.AllowedCalls(callOpts, call.Target, callMethodID)
	if err != nil {
		return errors.Wrap(err, "get allowed calls")
	} else if !allowed {
		return errors.New("call not allowed")
	}

	return nil
}

func approveOutboxSpend(ctx context.Context, expense bindings.ISolverNetTokenExpense, backend *ethbackend.Backend, solverAddr, outboxAddr common.Address) error {
	addr, err := cast.EthAddress(expense.Token[:])
	if err != nil {
		return errors.Wrap(err, "cast token address")
	}

	token, err := bindings.NewIERC20(addr, backend)
	if err != nil {
		return errors.Wrap(err, "new token")
	}

	isApproved := func() (bool, error) {
		allowance, err := token.Allowance(&bind.CallOpts{Context: ctx}, solverAddr, outboxAddr)
		if err != nil {
			return false, errors.Wrap(err, "get allowance")
		}

		return new(big.Int).Sub(allowance, expense.Amount).Sign() >= 0, nil
	}

	if approved, err := isApproved(); err != nil {
		return err
	} else if approved {
		return nil
	}

	txOpts, err := backend.BindOpts(ctx, solverAddr)
	if err != nil {
		return err
	}

	tx, err := token.Approve(txOpts, outboxAddr, umath.MaxUint256)
	if err != nil {
		return errors.Wrap(err, "approve token")
	} else if _, err := backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
	}

	if approved, err := isApproved(); err != nil {
		return err
	} else if !approved {
		return errors.New("approve failed")
	}

	return nil
}

func newRejector(
	inboxContracts map[uint64]*bindings.SolverNetInbox,
	backends ethbackend.Backends,
	solverAddr common.Address,
) func(ctx context.Context, chainID uint64, order Order, reason rejectReason) error {
	return func(ctx context.Context, chainID uint64, order Order, reason rejectReason) error {
		inbox, ok := inboxContracts[chainID]
		if !ok {
			return errors.New("unknown chain")
		}

		backend, err := backends.Backend(chainID)
		if err != nil {
			return err
		}

		// Ensure latest on-chain order is still pending
		if latest, err := inbox.GetOrder(&bind.CallOpts{Context: ctx}, order.ID); err != nil {
			return errors.Wrap(err, "get order")
		} else if latest.State.Status != statusPending {
			return errors.New("order status not pending anymore")
		}

		txOpts, err := backend.BindOpts(ctx, solverAddr)
		if err != nil {
			return err
		}

		tx, err := inbox.Reject(txOpts, order.ID, uint8(reason))
		if err != nil {
			return errors.Wrap(err, "reject order")
		} else if _, err := backend.WaitMined(ctx, tx); err != nil {
			return errors.Wrap(err, "wait mined")
		}

		return nil
	}
}

func newAcceptor(
	inboxContracts map[uint64]*bindings.SolverNetInbox,
	backends ethbackend.Backends,
	solverAddr common.Address,
) func(ctx context.Context, chainID uint64, order Order) error {
	return func(ctx context.Context, chainID uint64, order Order) error {
		inbox, ok := inboxContracts[chainID]
		if !ok {
			return errors.New("unknown chain")
		}

		backend, err := backends.Backend(chainID)
		if err != nil {
			return err
		}

		txOpts, err := backend.BindOpts(ctx, solverAddr)
		if err != nil {
			return err
		}

		tx, err := inbox.Accept(txOpts, order.ID)
		if err != nil {
			return errors.Wrap(err, "accept order")
		} else if _, err := backend.WaitMined(ctx, tx); err != nil {
			return errors.Wrap(err, "wait mined")
		}

		return nil
	}
}

func newIDParser(inboxContracts map[uint64]*bindings.SolverNetInbox) func(chainID uint64, log types.Log) (OrderID, error) {
	return func(chainID uint64, log types.Log) (OrderID, error) {
		inbox, ok := inboxContracts[chainID]
		if !ok {
			return OrderID{}, errors.New("unknown chain")
		}

		event, ok := eventsByTopic[log.Topics[0]]
		if !ok {
			return OrderID{}, errors.New("unknown event")
		}

		return event.ParseID(inbox.SolverNetInboxFilterer, log)
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

		return Order{
			ID:         o.Resolved.OrderId,
			Resolved:   o.Resolved,
			Status:     o.State.Status,
			AcceptedBy: o.State.AcceptedBy,
			History:    o.History,
		}, true, nil
	}
}
