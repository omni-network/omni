package app

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/e2e/solve/devapp"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

// maybeStartLoadGen starts the load generator on ephemeral networks.
func maybeStartLoadGen(ctx context.Context, cfg Config, network netconf.Network, backends ethbackend.Backends) error {
	if !network.ID.IsEphemeral() || cfg.LoadGenPrivKey == "" {
		return nil
	}

	privKey, err := ethcrypto.LoadECDSA(cfg.LoadGenPrivKey)
	if err != nil {
		return errors.Wrap(err, "load loadgen private key")
	}

	addr, err := backends.AddAccount(privKey)
	if err != nil {
		return errors.Wrap(err, "add loadgen account")
	}

	log.Info(ctx, "Starting ephemeral network load generation", "depositor", addr)

	go ephemeralLoadGenForever(ctx, network, backends, addr)

	return nil
}

func ephemeralLoadGenForever(
	ctx context.Context,
	network netconf.Network,
	backends ethbackend.Backends,
	depositor common.Address,
) {
	period := time.Minute * 5
	timer := time.NewTimer(0) // Tick immediately, then tick every period
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			if err := depositDevAppOnce(ctx, network, backends, depositor); err != nil {
				log.Warn(ctx, "Depositing to devapp failed (will retry)", err)
			}
			timer.Reset(period)
		}
	}
}

func depositDevAppOnce(
	ctx context.Context,
	network netconf.Network,
	backends ethbackend.Backends,
	depositor common.Address,
) error {
	app, err := devapp.GetApp(network.ID)
	if err != nil {
		return err
	}

	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		return err
	}

	backend, err := backends.Backend(app.L2.ChainID)
	if err != nil {
		return err
	}

	depositAmount := big.NewInt(params.GWei) // Deposit 1 gwei into devapp mock vault

	// Ensure depositor has enough balance to deposit
	if bal, err := backend.BalanceAt(ctx, depositor, nil); err != nil {
		return errors.Wrap(err, "get depositor balance")
	} else if bal.Cmp(depositAmount) <= 0 {
		return errors.New("depositor balance too low", "balance", bal, "required", depositAmount)
	}

	reqs, err := devapp.RequestDeposits(ctx, backend, addrs.SolveInbox, devapp.DepositArgs{
		OnBehalfOf: depositor,
		Amount:     depositAmount,
	})
	if err != nil {
		return errors.Wrap(err, "request deposits")
	} else if len(reqs) != 1 {
		return errors.New("expected 1 deposit request", "got", len(reqs))
	}
	req := reqs[0]
	ctx = log.WithCtx(ctx, "req_id", reqIDOffset(req.ID))
	log.Debug(ctx, "Loadgen requested deposit to devapp")

	// Wait up to 1min for deposit to complete
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	t0 := time.Now()
	for {
		if ok, err := devapp.IsDeposited(ctx, backends, req); err != nil {
			return err
		} else if ok {
			log.Debug(ctx, "Loadgen deposit to devapp complete", "duration", time.Since(t0))
			break
		}

		time.Sleep(time.Second)
		if ctx.Err() != nil {
			return errors.New("timeout waiting for deposit")
		}
	}

	return nil
}
