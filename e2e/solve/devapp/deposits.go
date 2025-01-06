package devapp

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	stypes "github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

type DepositReq struct {
	ID      stypes.ReqID   // id in inbox
	Token   common.Address // token address
	Deposit DepositArgs    // deposit args
}

// TestFlow submits deposit requests to the solve inbox and waits for them to be processed.
// It also includes a few invalid deposits, which should be rejected.
func TestFlow(ctx context.Context, network netconf.Network, endpoints xchain.RPCEndpoints) error {
	if network.ID != netconf.Devnet {
		return errors.New("only devnet")
	}

	backends, err := ethbackend.BackendsFromNetwork(network, endpoints)
	if err != nil {
		return err
	}

	const numDeposits = 3
	deposits, err := makeDevnetDeposits(ctx, network.ID, backends, numDeposits)
	if err != nil {
		return err
	}

	// Also do invalid deposits
	const numInvalid = 2
	invalids, err := makeDevnetDeposits(ctx, network.ID, backends, numInvalid, WithInvalidCall())
	if err != nil {
		return err
	}

	timeout, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	// Wait for all deposits to be completed on dest chain by solver/outbox
	toCheck := toSet(deposits)
	for {
		if timeout.Err() != nil {
			return errors.New("timeout waiting for deposits", "remaining", toCheck)
		}

		for deposit := range toCheck {
			bal, err := DepositedBalance(ctx, network.ID, backends, deposit.Deposit.OnBehalfOf)
			if err != nil {
				return err
			}

			// Assumes balance(onBehalfOf) was zero before deposit request
			// Assumes one deposit per test case onBehalfOf addr
			if bal.Cmp(deposit.Deposit.Amount) == 0 {
				log.Debug(ctx, "Deposit complete", "req_id", deposit.ID, "remaining", len(toCheck)-1)
				delete(toCheck, deposit)
			}
		}

		if len(toCheck) == 0 {
			break
		}

		time.Sleep(time.Second)
	}

	log.Debug(ctx, "All deposits fulfilled")

	// Wait for requests to be claimed by solver
	toCheck = toSet(deposits)
	for {
		if timeout.Err() != nil {
			return errors.New("timeout waiting for claims", "remaining", toCheck)
		}

		const statusClaimed = 6

		for deposit := range toCheck {
			status, err := RequestStatus(ctx, network.ID, backends, deposit.ID)
			if err != nil {
				return err
			} else if status == statusClaimed {
				log.Debug(ctx, "Deposit claimed", "req_id", deposit.ID, "remaining", len(toCheck)-1)
				delete(toCheck, deposit)
			}
		}

		if len(toCheck) == 0 {
			break
		}

		time.Sleep(time.Second)
	}

	log.Debug(ctx, "All deposits claimed")

	// Wait for invalid to be rejected
	toCheck = toSet(invalids)
	for {
		if timeout.Err() != nil {
			return errors.New("timeout waiting for invalid deposit rejections", "remaining", toCheck)
		}

		const statusRejected uint8 = 3

		for deposit := range toCheck {
			status, err := RequestStatus(ctx, network.ID, backends, deposit.ID)
			if err != nil {
				return err
			} else if status == statusRejected {
				log.Debug(ctx, "Invalid deposit rejected", "req_id", deposit.ID, "remaining", len(toCheck)-1)
				delete(toCheck, deposit)
			}
		}

		if len(toCheck) == 0 {
			break
		}

		time.Sleep(time.Second)
	}

	return nil
}

func makeDevnetDeposits(ctx context.Context, network netconf.ID, backends ethbackend.Backends, count int, opts ...CallOption) ([]DepositReq, error) {
	if network != netconf.Devnet {
		return nil, errors.New("only devnet supported")
	}

	app := MustGetApp(network)

	backend, err := backends.Backend(app.L2.ChainID)
	if err != nil {
		return nil, err
	}

	depositors, err := addRandomDepositors(count, backend)
	if err != nil {
		return nil, errors.Wrap(err, "make depositors")
	}

	depositAmount := big.NewInt(1e18)
	fundAmount := new(big.Int).Add(depositAmount, big.NewInt(params.GWei)) // Add gas

	// fund for gas
	if err := anvil.FundAccounts(ctx, backend, fundAmount, depositors...); err != nil {
		return nil, errors.Wrap(err, "fund accounts")
	}

	addrs, err := contracts.GetAddresses(ctx, netconf.Devnet)
	if err != nil {
		return nil, errors.Wrap(err, "get addresses")
	}

	var reqs []DepositReq
	for _, depositor := range depositors {
		args := DepositArgs{
			OnBehalfOf: depositor,
			Amount:     depositAmount,
		}
		req, err := RequestDeposit(ctx, network, backends, addrs.SolveInbox, args, opts...)
		if err != nil {
			return nil, errors.Wrap(err, "request deposit")
		}

		log.Error(ctx, "Deposit requested", nil,
			"req_id", req.ID, "on_behalf_of", depositor, "amount", depositAmount,
			"target", req.Token)

		reqs = append(reqs, req)
	}

	return reqs, nil
}

// DepositedBalance returns the balance of onBehalfOf in the vault.
func DepositedBalance(ctx context.Context, network netconf.ID, backends ethbackend.Backends, onBehalfOf common.Address) (*big.Int, error) {
	app := MustGetApp(network)

	backend, err := backends.Backend(app.L1.ChainID)
	if err != nil {
		return nil, errors.Wrap(err, "backend")
	}

	vault, err := bindings.NewMockVault(app.L1Vault, backend)
	if err != nil {
		return nil, errors.Wrap(err, "new mock vault")
	}

	callOpts := &bind.CallOpts{Context: ctx}

	balance, err := vault.Balances(callOpts, onBehalfOf)
	if err != nil {
		return nil, errors.Wrap(err, "get balance")
	}

	return balance, nil
}

// RequestStatus returns the status of a request in the inbox.
func RequestStatus(ctx context.Context, network netconf.ID, backends ethbackend.Backends, reqID [32]byte) (uint8, error) {
	app := MustGetApp(network)

	backend, err := backends.Backend(app.L2.ChainID)
	if err != nil {
		return 0, errors.Wrap(err, "backend")
	}

	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return 0, errors.Wrap(err, "get addresses")
	}

	inbox, err := bindings.NewSolveInbox(addrs.SolveInbox, backend)
	if err != nil {
		return 0, errors.Wrap(err, "new mock vault")
	}

	callOpts := &bind.CallOpts{Context: ctx}

	req, err := inbox.GetRequest(callOpts, reqID)
	if err != nil {
		return 0, errors.Wrap(err, "get balance")
	}

	return req.Status, nil
}

// addRandomDepositors adds n random depositors privkeys to the backend.
// It returns the addresses of the added depositors.
func addRandomDepositors(n int, backend *ethbackend.Backend) ([]common.Address, error) {
	var depositors []common.Address
	for range n {
		pk, err := crypto.GenerateKey()
		if err != nil {
			return nil, errors.Wrap(err, "generate key")
		}

		depositor, err := backend.AddAccount(pk)
		if err != nil {
			return nil, errors.Wrap(err, "add account")
		}

		depositors = append(depositors, depositor)
	}

	return depositors, nil
}

func RequestDeposit(ctx context.Context, network netconf.ID, backends ethbackend.Backends, inbox common.Address, deposit DepositArgs, opts ...CallOption) (DepositReq, error) {
	app := MustGetApp(network)
	backend, err := backends.Backend(app.L2.ChainID)
	if err != nil {
		return DepositReq{}, errors.Wrap(err, "l2 backend")
	}

	if err := mintAndApprove(ctx, app, backend, inbox, deposit); err != nil {
		return DepositReq{}, errors.Wrap(err, "mint and approve")
	}

	req, err := requestAtInbox(ctx, app, backend, inbox, deposit, opts...)
	if err != nil {
		return DepositReq{}, errors.Wrap(err, "request at inbox")
	}

	return req, nil
}

type CallOption func(*bindings.SolveCall)

// WithInvalidCall returns an option that sets the call target to an invalid address.
func WithInvalidCall() CallOption {
	return func(call *bindings.SolveCall) {
		call.Target = common.Address{1}
	}
}

func requestAtInbox(ctx context.Context, app App, backend *ethbackend.Backend, addr common.Address, deposit DepositArgs, opts ...CallOption) (DepositReq, error) {
	inbox, err := bindings.NewSolveInbox(addr, backend)
	if err != nil {
		return DepositReq{}, errors.Wrap(err, "new solve inbox")
	}

	txOpts, err := backend.BindOpts(ctx, deposit.OnBehalfOf)
	if err != nil {
		return DepositReq{}, errors.Wrap(err, "bind opts")
	}

	data, err := packDeposit(deposit)
	if err != nil {
		return DepositReq{}, errors.Wrap(err, "pack deposit")
	}

	call := bindings.SolveCall{
		ChainId: app.L1.ChainID,
		Target:  app.L1Vault,
		Value:   new(big.Int), // 0 native
		Data:    data,
	}

	for _, opt := range opts {
		opt(&call)
	}

	tx, err := inbox.Request(txOpts, call,
		[]bindings.SolveTokenDeposit{{
			Token:  app.L2Token,
			Amount: deposit.Amount,
		}},
	)
	if err != nil {
		return DepositReq{}, errors.Wrap(err, "request deposit")
	}

	rec, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return DepositReq{}, errors.Wrap(err, "wait mined request")
	}

	id, ok := parseReqID(inbox.SolveInboxFilterer, rec.Logs)
	if !ok {
		return DepositReq{}, errors.New("parse req id")
	}

	return DepositReq{
		ID:      id,
		Token:   app.L2Token,
		Deposit: deposit,
	}, nil
}

func mintAndApprove(ctx context.Context, app App, backend *ethbackend.Backend, inbox common.Address, deposit DepositArgs) error {
	txOpts, err := backend.BindOpts(ctx, deposit.OnBehalfOf)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	token, err := bindings.NewMockToken(app.L2Token, backend)
	if err != nil {
		return errors.Wrap(err, "new mock token")
	}

	// mint tokens
	tx, err := token.Mint(txOpts, deposit.OnBehalfOf, deposit.Amount)
	if err != nil {
		return errors.Wrap(err, "mint tokens")
	}

	if _, err := backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
	}

	// approve inbox
	tx, err = token.Approve(txOpts, inbox, deposit.Amount)
	if err != nil {
		return errors.Wrap(err, "approve inbox")
	}

	if _, err := backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined mint")
	}

	return nil
}

func parseReqID(inbox bindings.SolveInboxFilterer, logs []*types.Log) (stypes.ReqID, bool) {
	for _, log := range logs {
		e, err := inbox.ParseRequested(*log)
		if err == nil {
			return e.Id, true
		}
	}

	return [32]byte{}, false
}

func packDeposit(args DepositArgs) ([]byte, error) {
	calldata, err := vaultABI.Pack("deposit", args.OnBehalfOf, args.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "pack deposit call data")
	}

	return calldata, nil
}

func toSet[T comparable](slice []T) map[T]bool {
	set := make(map[T]bool)
	for _, v := range slice {
		set[v] = true
	}

	return set
}
