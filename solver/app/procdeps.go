package app

import (
	"context"
	"math/big"
	"strings"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"
	stypes "github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// procDeps abstracts dependencies for the event processor allowed simplified testing.
type procDeps struct {
	ParseID      func(chainID uint64, log types.Log) (stypes.ReqID, error)
	GetRequest   func(ctx context.Context, chainID uint64, id stypes.ReqID) (bindings.SolveRequest, bool, error)
	ShouldReject func(ctx context.Context, chainID uint64, req bindings.SolveRequest) (rejectReason, bool, error)
	SetCursor    func(ctx context.Context, chainID uint64, height uint64) error

	Accept  func(ctx context.Context, chainID uint64, req bindings.SolveRequest) error
	Reject  func(ctx context.Context, chainID uint64, req bindings.SolveRequest, reason rejectReason) error
	Fulfill func(ctx context.Context, chainID uint64, req bindings.SolveRequest) error
	Claim   func(ctx context.Context, chainID uint64, req bindings.SolveRequest) error

	// Monitoring helpers
	TargetName func(bindings.SolveRequest) string
	ChainName  func(chainID uint64) string
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
	solverAddr, outboxAddr common.Address,
) func(ctx context.Context, srcChainID uint64, req bindings.SolveRequest) error {
	return func(ctx context.Context, srcChainID uint64, req bindings.SolveRequest) error {
		destChainID := req.Call.ChainId // Fulfilling happens on destination chain
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

		if ok, err := outbox.DidFulfill(callOpts, req.Id, srcChainID, req.Call); err != nil {
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

		for _, prereq := range prereqs {
			if err := approveOutboxSpend(ctx, prereq, backend, solverAddr, outboxAddr); err != nil {
				return errors.Wrap(err, "approve outbox spend")
			}

			if err := checkAllowedCall(ctx, outbox, req.Call); err != nil {
				return errors.Wrap(err, "check allowed call")
			}
		}

		if err := target.LogCall(ctx, req.Call); err != nil {
			return errors.Wrap(err, "debug call")
		}

		// xcall fee
		fee, err := outbox.FulfillFee(callOpts, srcChainID)
		if err != nil {
			return errors.Wrap(err, "get fulfill fee")
		}

		txOpts.Value = fee
		tx, err := outbox.Fulfill(txOpts, req.Id, srcChainID, req.Call, prereqs)
		if err != nil {
			return errors.Wrap(err, "fulfill request", "custom", detectCustomError(err))
		} else if _, err := backend.WaitMined(ctx, tx); err != nil {
			return errors.Wrap(err, "wait mined")
		}

		if ok, err := outbox.DidFulfill(callOpts, req.Id, srcChainID, req.Call); err != nil {
			return errors.Wrap(err, "did fulfill")
		} else if !ok {
			return errors.New("fulfill failed [BUG]")
		}

		return nil
	}
}

func detectCustomError(custom error) string {
	contracts := map[string]*bind.MetaData{
		"inbox":      bindings.SolveInboxMetaData,
		"outbox":     bindings.SolveOutboxMetaData,
		"mock_vault": bindings.MockVaultMetaData,
		"mock_erc20": bindings.MockERC20MetaData,
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

func approveOutboxSpend(ctx context.Context, prereq bindings.SolveTokenPrereq, backend *ethbackend.Backend, solverAddr, outboxAddr common.Address) error {
	token, err := bindings.NewIERC20(prereq.Token, backend)
	if err != nil {
		return errors.Wrap(err, "new token")
	}

	isApproved := func() (bool, error) {
		allowance, err := token.Allowance(&bind.CallOpts{Context: ctx}, solverAddr, outboxAddr)
		if err != nil {
			return false, errors.Wrap(err, "get allowance")
		}

		return new(big.Int).Sub(allowance, prereq.Amount).Sign() >= 0, nil
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

		// Ensure latest on-chain request is still pending
		if latest, err := inbox.GetRequest(&bind.CallOpts{Context: ctx}, req.Id); err != nil {
			return errors.Wrap(err, "get request")
		} else if latest.Status != statusPending {
			return errors.New("request status not pending anymore")
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

func newIDParser(inboxContracts map[uint64]*bindings.SolveInbox) func(chainID uint64, log types.Log) (stypes.ReqID, error) {
	return func(chainID uint64, log types.Log) (stypes.ReqID, error) {
		inbox, ok := inboxContracts[chainID]
		if !ok {
			return stypes.ReqID{}, errors.New("unknown chain")
		}

		event, ok := eventsByTopic[log.Topics[0]]
		if !ok {
			return stypes.ReqID{}, errors.New("unknown event")
		}

		return event.ParseID(inbox.SolveInboxFilterer, log)
	}
}

func newRequestGetter(inboxContracts map[uint64]*bindings.SolveInbox) func(ctx context.Context, chainID uint64, id stypes.ReqID) (bindings.SolveRequest, bool, error) {
	return func(ctx context.Context, chainID uint64, id stypes.ReqID) (bindings.SolveRequest, bool, error) {
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
