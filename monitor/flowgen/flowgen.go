package flowgen

import (
	"context"
	"math/big"
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
	"github.com/omni-network/omni/lib/xchain"
	"github.com/omni-network/omni/monitor/flowgen/bridging"
	"github.com/omni-network/omni/monitor/flowgen/types"
	"github.com/omni-network/omni/monitor/flowgen/util"

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

	var jobs []types.Job

	switch network.ID {
	case netconf.Devnet:
		// Bridging of native ETH on a devnet
		job, err := bridging.NewJob(
			netconf.Devnet,
			evmchain.IDMockL1,
			evmchain.IDMockL2,
			eoa.RoleFlowgen,
			common.Address{}, // native ETH
			big.NewInt(0).Mul(util.MilliEther, big.NewInt(2)), // 0.002 ETH
		)
		if err != nil {
			return err
		}
		jobs = append(jobs, job)

		job, err = bridging.NewJob(
			netconf.Devnet,
			evmchain.IDMockL2,
			evmchain.IDMockL1,
			eoa.RoleFlowgen,
			common.Address{}, // native ETH
			big.NewInt(0).Mul(util.MilliEther, big.NewInt(2)), // 0.002 ETH
		)
		if err != nil {
			return err
		}
		jobs = append(jobs, job)

	case netconf.Omega:
		// Bridging of native ETH on omega
		job, err := bridging.NewJob(
			netconf.Omega,
			evmchain.IDBaseSepolia,
			evmchain.IDOpSepolia,
			eoa.RoleFlowgen,
			common.Address{}, // native ETH
			big.NewInt(0).Mul(util.MilliEther, big.NewInt(2)), // 0.002 ETH
		)
		if err != nil {
			return err
		}
		jobs = append(jobs, job)

		job, err = bridging.NewJob(
			netconf.Omega,
			evmchain.IDOpSepolia,
			evmchain.IDBaseSepolia,
			eoa.RoleFlowgen,
			common.Address{}, // native ETH
			big.NewInt(0).Mul(util.MilliEther, big.NewInt(2)), // 0.002 ETH
		)
		if err != nil {
			return err
		}
		jobs = append(jobs, job)

	default:
	}

	for _, job := range jobs {
		go func() {
			ticker := time.NewTicker(job.Cadence)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if err := run(log.WithCtx(ctx, "job", job.Name), backends, job); err != nil {
						log.Error(ctx, "Flowgen: job failed (will retry)", err)
					}
				}
			}
		}()
	}

	return nil
}

// run runs a job exactly once.
func run(ctx context.Context, backends ethbackend.Backends, j types.Job) error {
	log.Debug(ctx, "Flowgen: running job")

	id, err := solvernet.OpenOrder(ctx, j.Network, j.SrcChain, backends, j.Owner, j.OrderData)
	if err != nil {
		return errors.Wrap(err, "open order")
	}

	ctx = log.WithCtx(ctx, "id", id)

	log.Debug(ctx, "Flowgen: order opened")

	status, err := waitForFinalStatus(ctx, backends, j, id)
	if err != nil {
		return errors.Wrap(err, "wait for status")
	}
	log.Info(ctx, "Flowgen: order finalized", "status", status)

	return nil
}

// waitForFinalStatus monitors the specified order id for the final status and return it. Since we
// assume that all orders will eventually be rejected, closed or claimed, the function never terminates.
func waitForFinalStatus(
	ctx context.Context,
	backends ethbackend.Backends,
	j types.Job,
	orderID solvernet.OrderID,
) (solvernet.OrderStatus, error) {
	addrs, err := contracts.GetAddresses(ctx, j.Network)
	if err != nil {
		panic(err)
	}

	backend, err := backends.Backend(j.SrcChain)
	if err != nil {
		return solvernet.StatusInvalid, errors.Wrap(err, "get backend")
	}

	inbox, err := bindings.NewSolverNetInbox(addrs.SolverNetInbox, backend)
	if err != nil {
		return solvernet.StatusInvalid, errors.Wrap(err, "create inbox contract")
	}

	for {
		latest, err := inbox.GetOrder(&bind.CallOpts{Context: ctx}, orderID)
		if err != nil {
			return solvernet.StatusInvalid, errors.Wrap(err, "get order")
		}

		status := solvernet.OrderStatus(latest.State.Status)

		switch status {
		case solvernet.StatusInvalid, solvernet.StatusRejected, solvernet.StatusClosed, solvernet.StatusClaimed:
			return status, nil
		default:
			log.Debug(ctx, "Flowgen: order in flight", "status", status)
		}

		time.Sleep(time.Second)
	}
}
