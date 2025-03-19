package flowgen

import (
	"context"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/monitor/flowgen/bridging"
	"github.com/omni-network/omni/monitor/flowgen/symbiotic"
	"github.com/omni-network/omni/monitor/flowgen/types"
	stypes "github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

func Start(
	ctx context.Context,
	network netconf.Network,
	rpcEndpoints xchain.RPCEndpoints,
	keyPath string,
) error {
	if keyPath == "" {
		return errors.New("private key is required")
	}

	privKey, err := ethcrypto.LoadECDSA(keyPath)
	if err != nil {
		return errors.Wrap(err, "load key", "path", keyPath)
	}

	backends, err := ethbackend.BackendsFromNetwork(network, rpcEndpoints, privKey)
	if err != nil {
		return errors.Wrap(err, "backends")
	}

	owner := eoa.MustAddress(network.ID, eoa.RoleFlowgen)

	return startWithBackends(ctx, network, backends, owner)
}

func startWithBackends(
	ctx context.Context,
	network netconf.Network,
	backends ethbackend.Backends,
	owner common.Address,
) error {
	var jobs []types.Job

	result, err := bridging.Jobs(network.ID, backends, owner)
	if err != nil {
		return errors.Wrap(err, "bridge jobs")
	}
	jobs = append(jobs, result...)

	result, err = symbiotic.Jobs(ctx, backends, network.ID, owner)
	if err != nil {
		return errors.Wrap(err, "symbiotic jobs")
	}
	jobs = append(jobs, result...)

	for _, job := range jobs {
		go func() {
			timer := time.NewTimer(0)
			defer timer.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-timer.C:
					jobsTotal.Inc()
					runCtx := log.WithCtx(ctx, "job", job.Name)
					if err := run(runCtx, job); err != nil {
						log.Warn(runCtx, "Flowgen: job failed (will retry)", err)
						jobsFailed.Inc()
					}
					timer.Reset(job.Cadence)
				}
			}
		}()
	}

	return nil
}

// run runs a job exactly once.
func run(ctx context.Context, job types.Job) error {
	log.Debug(ctx, "Flowgen: running job")

	orderID, ok, err := job.OpenOrderFunc(ctx)
	if err != nil {
		return errors.Wrap(err, "open order")
	}

	if !ok {
		return nil
	}

	ctx = log.WithCtx(ctx, "order_id", orderID)

	log.Debug(ctx, "Flowgen: order opened")

	if err := awaitClaimed(ctx, job, orderID); err != nil {
		return errors.Wrap(err, "await claimed")
	}

	log.Info(ctx, "Flowgen: order claimed")

	return nil
}

// awaitClaimed blocks until the order is claimed.
func awaitClaimed(ctx context.Context, job types.Job, orderID solvernet.OrderID) error {
	addrs, err := contracts.GetAddresses(ctx, job.NetworkID)
	if err != nil {
		return errors.New("contract addresses")
	}

	inbox, err := bindings.NewSolverNetInbox(addrs.SolverNetInbox, job.SrcChainBackend)
	if err != nil {
		return errors.Wrap(err, "create inbox contract")
	}

	var checks uint16
	const logFreq = 20
	for {
		order, err := inbox.GetOrder(&bind.CallOpts{Context: ctx}, orderID)
		if err != nil {
			return errors.Wrap(err, "get order")
		}

		status := solvernet.OrderStatus(order.State.Status)
		reason := stypes.RejectReason(order.State.RejectReason)

		switch status {
		case solvernet.StatusClaimed:
			return nil
		case solvernet.StatusRejected:
			return errors.New("order rejected", "reason", reason)
		case solvernet.StatusClosed, solvernet.StatusInvalid:
			return errors.New("unexpected order status", "status", status)
		default:
			checks++
			if checks%logFreq == 0 {
				log.Debug(ctx, "Flowgen: order in flight", "status", status)
			}
		}

		time.Sleep(5 * time.Second)
	}
}
