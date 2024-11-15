package devapp

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
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

func RequestDeposits(ctx context.Context, endpoints xchain.RPCEndpoints, backends ethbackend.Backends) ([]DepositReq, error) {
	app := GetApp()

	backend, err := backends.Backend(app.L2.ChainID)
	if err != nil {
		return nil, err
	}

	rpc, err := endpoints.ByNameOrID(app.L2.Name, app.L2.ChainID)
	if err != nil {
		return nil, err
	}

	backend = backend.Clone()

	const numDeposits = 10

	depositors, err := makeDepositors(numDeposits, backend)
	if err != nil {
		return nil, errors.Wrap(err, "make depositors")
	}

	// fund for gas
	if err := anvil.FundAccounts(ctx, rpc, big.NewInt(1e18), depositors...); err != nil {
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

func CheckDeposits(ctx context.Context, backends ethbackend.Backends, reqs []DepositReq) error {
	app := GetApp()

	backend, err := backends.Backend(app.L1.ChainID)
	if err != nil {
		return errors.Wrap(err, "backend")
	}

	vault, err := bindings.NewMockVault(app.L1Vault, backend)
	if err != nil {
		return errors.Wrap(err, "new mock vault")
	}

	callOpts := &bind.CallOpts{Context: ctx}

	for _, req := range reqs {
		balance, err := vault.Balances(callOpts, req.Deposit.OnBehalfOf)
		if err != nil {
			return errors.Wrap(err, "get balance")
		}

		// assumes balance(onBehalfOf) was zero before deposit request
		// assumes one deposit per test case onBehalfOf addr
		if balance.Cmp(req.Deposit.Amount) != 0 {
			return errors.New("missing deposit",
				"requester", req.Deposit.OnBehalfOf,
				"expected", req.Deposit.Amount,
				"actual", balance)
		}
	}

	return nil
}

func makeDepositors(n int, backend *ethbackend.Backend) ([]common.Address, error) {
	depositors := make([]common.Address, n)
	for i := 0; i < n; i++ {
		pk, err := crypto.GenerateKey()
		if err != nil {
			return nil, errors.Wrap(err, "generate key")
		}

		depositor, err := backend.AddAccount(pk)
		if err != nil {
			return nil, errors.Wrap(err, "add account")
		}

		depositors[i] = depositor
	}

	return depositors, nil
}

func requestDeposits(ctx context.Context, backend *ethbackend.Backend, inbox common.Address, depositors []common.Address) ([]DepositReq, error) {
	reqs := make([]DepositReq, 0, len(depositors))
	for i := 0; i < len(depositors); i++ {
		depositor := depositors[i]

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
	data, err := vaultDeposit.Inputs.Pack(args)
	if err != nil {
		return nil, errors.Wrap(err, "unpack data")
	}

	return data, nil
}
