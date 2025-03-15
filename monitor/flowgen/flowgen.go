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
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/monitor/flowgen/bridging"
	"github.com/omni-network/omni/monitor/flowgen/types"
	stypes "github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

	var jobs []types.Job

	result, err := bridgeJobs(network.ID)
	if err != nil {
		return err
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
					if err := run(log.WithCtx(ctx, "job", job.Name), backends, job); err != nil {
						log.Warn(ctx, "Flowgen: job failed (will retry)", err)
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
func run(ctx context.Context, backends ethbackend.Backends, j types.Job) error {
	log.Debug(ctx, "Flowgen: running job")

	orderID, err := solvernet.OpenOrder(ctx, j.Network, j.SrcChain, backends, j.Owner, j.OrderData)
	if err != nil {
		return errors.Wrap(err, "open order")
	}

	ctx = log.WithCtx(ctx, "order_id", orderID)

	log.Debug(ctx, "Flowgen: order opened")

	if err := awaitClaimed(ctx, backends, j, orderID); err != nil {
		return errors.Wrap(err, "await claimed")
	}

	log.Info(ctx, "Flowgen: order claimed")

	return nil
}

// awaitClaimed blocks until the order is claimed.
// It returns an.
func awaitClaimed(
	ctx context.Context,
	backends ethbackend.Backends,
	j types.Job,
	orderID solvernet.OrderID,
) error {
	addrs, err := contracts.GetAddresses(ctx, j.Network)
	if err != nil {
		return err
	}

	backend, err := backends.Backend(j.SrcChain)
	if err != nil {
		return errors.Wrap(err, "get backend")
	}

	inbox, err := bindings.NewSolverNetInbox(addrs.SolverNetInbox, backend)
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

func bridgeJobs(network netconf.ID) ([]types.Job, error) {
	type balanced struct {
		From uint64
		To   uint64
	}

	b, ok := map[netconf.ID]balanced{
		netconf.Devnet:  {evmchain.IDMockL1, evmchain.IDMockL2},
		netconf.Staging: {evmchain.IDBaseSepolia, evmchain.IDOpSepolia},
		netconf.Omega:   {evmchain.IDOpSepolia, evmchain.IDArbSepolia},
		netconf.Mainnet: {evmchain.IDOptimism, evmchain.IDArbitrumOne},
	}[network]
	if !ok {
		return nil, nil
	}

	// Bridging of native ETH
	amount := umath.EtherToWei(0.02) // 0.02 ETH

	job1, err := bridging.NewJob(network, b.From, b.To, eoa.RoleFlowgen, amount)
	if err != nil {
		return nil, err
	}

	job2, err := bridging.NewJob(network, b.To, b.From, eoa.RoleFlowgen, amount)
	if err != nil {
		return nil, err
	}

	return []types.Job{job1, job2}, nil
}
