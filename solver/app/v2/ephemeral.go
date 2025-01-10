//nolint:unused // This package is a work in progress.
package appv2

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
func maybeStartLoadGen(ctx context.Context, cfg Config, network netconf.ID, backends ethbackend.Backends) error {
	if !network.IsEphemeral() || cfg.LoadGenPrivKey == "" {
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
	network netconf.ID,
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
	network netconf.ID,
	backends ethbackend.Backends,
	depositor common.Address,
) error {
	app, err := devapp.GetApp(network)
	if err != nil {
		return err
	}

	addrs, err := contracts.GetAddresses(ctx, network)
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

	// Get target vault balance before deposit
	startBal, err := devapp.DepositedBalance(ctx, network, backends, depositor)
	if err != nil {
		return err
	}
	expectedBal := new(big.Int).Add(startBal, depositAmount)

	req, err := devapp.RequestDeposit(ctx, network, backends, addrs.SolveInbox, devapp.DepositArgs{
		OnBehalfOf: depositor,
		Amount:     depositAmount,
	})
	if err != nil {
		return errors.Wrap(err, "request deposits")
	}

	ctx = log.WithCtx(ctx, "req_id", req.ID)
	log.Debug(ctx, "Loadgen requested deposit to devapp")

	// Wait up to 1 hour for deposit to complete and claimed (fulfill xmsg uses ConfFinalized)
	ctx, cancel := context.WithTimeout(ctx, time.Hour)
	defer cancel()

	t0 := time.Now()
	for {
		actualBal, err := devapp.DepositedBalance(ctx, network, backends, req.Deposit.OnBehalfOf)
		if err != nil {
			return err
		} else if actualBal.Cmp(expectedBal) == 0 {
			log.Debug(ctx, "Loadgen deposit to devapp complete", "duration", time.Since(t0))
			break
		}

		time.Sleep(time.Second)
		if ctx.Err() != nil {
			return errors.New("timeout waiting for deposit",
				"req_id", req.ID,
				"start_bal", startBal,
				"expected_bal", expectedBal,
				"actual_bal", actualBal,
			)
		}
	}

	for {
		status, err := devapp.RequestStatus(ctx, network, backends, req.ID)
		if err != nil {
			return err
		} else if status == statusClaimed {
			log.Debug(ctx, "Loadgen deposit to devapp claimed", "duration", time.Since(t0))
			break
		}

		time.Sleep(time.Second * 5)
		if ctx.Err() != nil {
			return errors.New("timeout waiting for deposit", "req_id", req.ID)
		}
	}

	return nil
}
