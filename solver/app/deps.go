package app

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// procDeps abstracts dependencies for the event processor allowed simplified testing.
type procDeps struct {
	ParseID      func(chainID uint64, log types.Log) ([32]byte, error)
	GetRequest   func(ctx context.Context, chainID uint64, id [32]byte) (bindings.SolveRequest, bool, error)
	ShouldReject func(ctx context.Context, chainID uint64, req bindings.SolveRequest) (rejectReason, bool, error)
	SetCursor    func(ctx context.Context, chainID uint64, height uint64) error

	Accept  func(ctx context.Context, chainID uint64, req bindings.SolveRequest) error
	Reject  func(ctx context.Context, chainID uint64, req bindings.SolveRequest, reason rejectReason) error
	Fulfill func(ctx context.Context, chainID uint64, req bindings.SolveRequest) error
	Claim   func(ctx context.Context, chainID uint64, req bindings.SolveRequest) error
}

func newClaimer(
	inboxContracts map[uint64]*bindings.SolveInbox,
	backends ethbackend.Backends,
	solverAddr common.Address,
) func(ctx context.Context, chainID uint64, req bindings.SolveRequest) error {
	return func(ctx context.Context, chainID uint64, req bindings.SolveRequest) error {
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
		tx, err := inbox.Claim(txOpts, req.Id, solverAddr)
		if err != nil {
			return errors.Wrap(err, "claim request")
		} else if _, err := backend.WaitMined(ctx, tx); err != nil {
			return errors.Wrap(err, "wait mined")
		}

		return nil
	}
}

func newFulfiller(
	network netconf.ID,
	outboxContracts map[uint64]*bindings.SolveOutbox,
	backends ethbackend.Backends,
	solverAddr common.Address,
) func(ctx context.Context, chainID uint64, req bindings.SolveRequest) error {
	return func(ctx context.Context, chainID uint64, req bindings.SolveRequest) error {
		outbox, ok := outboxContracts[chainID]
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

		if ok, err := outbox.DidFulfill(&bind.CallOpts{Context: ctx}, req.Id, chainID, req.Call); err != nil {
			return errors.Wrap(err, "did fulfill")
		} else if ok {
			log.Info(ctx, "Skipping already fulfilled request", "req_id", req.Id)
			return nil
		}

		target, err := getTarget(network, req.Call)
		if err != nil {
			return errors.Wrap(err, "get target [BUG]")
		}

		prereqs, err := target.TokenPrereqs(req.Call)
		if err != nil {
			return errors.Wrap(err, "get token prereqs")
		}

		tx, err := outbox.Fulfill(txOpts, req.Id, chainID, req.Call, prereqs)
		if err != nil {
			return errors.Wrap(err, "fulfill request")
		} else if _, err := backend.WaitMined(ctx, tx); err != nil {
			return errors.Wrap(err, "wait mined")
		}

		return nil
	}
}

func newRejector(
	inboxContracts map[uint64]*bindings.SolveInbox,
	backends ethbackend.Backends,
	solverAddr common.Address,
) func(ctx context.Context, chainID uint64, req bindings.SolveRequest, reason rejectReason) error {
	return func(ctx context.Context, chainID uint64, req bindings.SolveRequest, reason rejectReason) error {
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

		tx, err := inbox.Reject(txOpts, req.Id, uint8(reason))
		if err != nil {
			return errors.Wrap(err, "reject request")
		} else if _, err := backend.WaitMined(ctx, tx); err != nil {
			return errors.Wrap(err, "wait mined")
		}

		return nil
	}
}

func newAcceptor(
	inboxContracts map[uint64]*bindings.SolveInbox,
	backends ethbackend.Backends,
	solverAddr common.Address,
) func(ctx context.Context, chainID uint64, req bindings.SolveRequest) error {
	return func(ctx context.Context, chainID uint64, req bindings.SolveRequest) error {
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

		tx, err := inbox.Accept(txOpts, req.Id)
		if err != nil {
			return errors.Wrap(err, "accept request")
		} else if _, err := backend.WaitMined(ctx, tx); err != nil {
			return errors.Wrap(err, "wait mined")
		}

		return nil
	}
}

func newIDParser(inboxContracts map[uint64]*bindings.SolveInbox) func(chainID uint64, log types.Log) ([32]byte, error) {
	return func(chainID uint64, log types.Log) ([32]byte, error) {
		inbox, ok := inboxContracts[chainID]
		if !ok {
			return [32]byte{}, errors.New("unknown chain")
		}

		event, ok := eventsByTopic[log.Topics[0]]
		if !ok {
			return [32]byte{}, errors.New("unknown event")
		}

		return event.ParseID(inbox.SolveInboxFilterer, log)
	}
}

func newRequestGetter(inboxContracts map[uint64]*bindings.SolveInbox) func(ctx context.Context, chainID uint64, id [32]byte) (bindings.SolveRequest, bool, error) {
	return func(ctx context.Context, chainID uint64, id [32]byte) (bindings.SolveRequest, bool, error) {
		inbox, ok := inboxContracts[chainID]
		if !ok {
			return bindings.SolveRequest{}, false, errors.New("unknown chain")
		}

		req, err := inbox.GetRequest(&bind.CallOpts{Context: ctx}, id)
		// TODO(corver): Detect not found
		if err != nil {
			return bindings.SolveRequest{}, false, errors.Wrap(err, "get request")
		}

		return req, true, nil
	}
}
