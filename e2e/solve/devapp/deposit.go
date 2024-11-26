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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type DepositReq struct {
	ID      [32]byte       // id in inbox
	Token   common.Address // token address
	Deposit DepositArgs    // deposit args
}

func RequestDeposits(ctx context.Context, backends ethbackend.Backends) ([]DepositReq, error) {
	app := GetApp()

	backend, err := backends.Backend(app.L2.ChainID)
	if err != nil {
		return nil, err
	}

	const numDeposits = 3

	depositors, err := addRandomDepositors(numDeposits, backend)
	if err != nil {
		return nil, errors.Wrap(err, "make depositors")
	}

	// fund for gas
	if err := anvil.FundAccounts(ctx, backend, big.NewInt(1e18), depositors...); err != nil {
		return nil, errors.Wrap(err, "fund accounts")
	}

	addrs, err := contracts.GetAddresses(ctx, netconf.Devnet)
	if err != nil {
		return nil, errors.Wrap(err, "get addresses")
	}

	reqs, err := requestDeposits(ctx, backend, addrs.SolveInbox, depositors)
	if err != nil {
		return nil, errors.Wrap(err, "request deposits")
	}

	return reqs, nil
}

func IsDeposited(ctx context.Context, backends ethbackend.Backends, req DepositReq) (bool, error) {
	app := GetApp()

	backend, err := backends.Backend(app.L1.ChainID)
	if err != nil {
		return false, errors.Wrap(err, "backend")
	}

	vault, err := bindings.NewMockVault(app.L1Vault, backend)
	if err != nil {
		return false, errors.Wrap(err, "new mock vault")
	}

	callOpts := &bind.CallOpts{Context: ctx}

	balance, err := vault.Balances(callOpts, req.Deposit.OnBehalfOf)
	if err != nil {
		return false, errors.Wrap(err, "get balance")
	}

	// assumes balance(onBehalfOf) was zero before deposit request
	// assumes one deposit per test case onBehalfOf addr
	return balance.Cmp(req.Deposit.Amount) == 0, nil
}

// TestFlow submits deposit requests to the solve inbox and waits for them to be processed.
func TestFlow(ctx context.Context, network netconf.Network, endpoints xchain.RPCEndpoints) error {
	backends, err := ethbackend.BackendsFromNetwork(network, endpoints)
	if err != nil {
		return err
	}

	deposits, err := RequestDeposits(ctx, backends)
	if err != nil {
		return err
	}

	timeout, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	// Wait for all deposits to be completed on dest chain by solver/outbox
	toCheck := toSet(deposits)
	for {
		if timeout.Err() != nil {
			return errors.New("timeout waiting for deposits")
		}

		for deposit := range toCheck {
			ok, err := IsDeposited(ctx, backends, deposit)
			if err != nil {
				return err
			} else if ok {
				log.Debug(ctx, "Deposit complete", "remaining", len(toCheck)-1)
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
			return errors.New("timeout waiting for claims")
		}

		const statusClaimed = 8

		for deposit := range toCheck {
			status, err := GetDepositStatus(ctx, backends, deposit)
			if err != nil {
				return err
			} else if status == statusClaimed {
				log.Debug(ctx, "Deposit claimed", "remaining", len(toCheck)-1)
				delete(toCheck, deposit)
			}
		}

		if len(toCheck) == 0 {
			break
		}

		time.Sleep(time.Second)
	}

	log.Debug(ctx, "All deposits claimed")

	return nil
}

func GetDepositStatus(ctx context.Context, backends ethbackend.Backends, deposit DepositReq) (uint8, error) {
	app := GetApp()

	backend, err := backends.Backend(app.L2.ChainID)
	if err != nil {
		return 0, errors.Wrap(err, "backend")
	}

	addrs, err := contracts.GetAddresses(ctx, netconf.Devnet)
	if err != nil {
		return 0, errors.Wrap(err, "get addresses")
	}

	inbox, err := bindings.NewSolveInbox(addrs.SolveInbox, backend)
	if err != nil {
		return 0, errors.Wrap(err, "new mock vault")
	}

	callOpts := &bind.CallOpts{Context: ctx}

	req, err := inbox.GetRequest(callOpts, deposit.ID)
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

func requestDeposits(ctx context.Context, backend *ethbackend.Backend, inbox common.Address, depositors []common.Address) ([]DepositReq, error) {
	var reqs []DepositReq
	for _, depositor := range depositors {
		deposit := DepositArgs{
			OnBehalfOf: depositor,
			Amount:     big.NewInt(1e18),
		}

		if err := mintAndApprove(ctx, backend, inbox, deposit); err != nil {
			return nil, errors.Wrap(err, "mint and approve")
		}

		req, err := requestAtInbox(ctx, backend, inbox, deposit)
		if err != nil {
			return nil, errors.Wrap(err, "request at inbox")
		}

		reqs = append(reqs, req)
	}

	return reqs, nil
}

func requestAtInbox(ctx context.Context, backend *ethbackend.Backend, addr common.Address, deposit DepositArgs) (DepositReq, error) {
	app := GetApp()

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

	tx, err := inbox.Request(txOpts,
		bindings.SolveCall{
			DestChainId: app.L1.ChainID,
			Target:      app.L1Vault,
			Value:       new(big.Int), // 0 native
			Data:        data,
		},
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

func mintAndApprove(ctx context.Context, backend *ethbackend.Backend, inbox common.Address, deposit DepositArgs) error {
	app := GetApp()

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

func parseReqID(inbox bindings.SolveInboxFilterer, logs []*types.Log) ([32]byte, bool) {
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
