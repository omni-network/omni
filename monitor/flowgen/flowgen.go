package flowgen

import (
	"context"
	"fmt"
	"math/big"
	"time"

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

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

func Start(
	ctx context.Context,
	xprovider xchain.Provider,
	network netconf.Network,
	rpcEndpoints xchain.RPCEndpoints,
	keyPath string,
) error {
	var jobs []*types.Job

	// Add jobs according to the current network ID
	switch network.ID {
	case netconf.Devnet:
		// Instantiate bridging of native ETH on a devnet
		job, err := bridging.NewJob(
			netconf.Devnet,
			evmchain.IDMockL1,
			evmchain.IDMockL2,
			eoa.RoleMonitor,
			common.Address{}, // native ETH
			big.NewInt(0).Mul(util.MilliEther, big.NewInt(2)), // 0.002 ETH
		)
		if err != nil {
			return err
		}

		jobs = append(jobs, job)

	default:
	}

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

	for _, job := range jobs {
		go func() {
			ticker := time.NewTicker(job.Cadence)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if err := run(ctx, xprovider, backends, job); err != nil {
						log.Error(ctx, fmt.Sprintf("Flowgen job %s failed", job.Name), err)
					}
				}
			}
		}()
	}

	return nil
}

// run runs a job exactly once.
func run(ctx context.Context, xprovider xchain.Provider, backends ethbackend.Backends, j *types.Job) error {
	log.Debug(ctx, "Flowgen: running job", "name", j.Name)

	// We fetch the latest finalized height of the source chain before we open the order.
	// Once we open the order, we start fetching all event logs starting above this
	// finalized height to retrieve the order status.
	height, err := xprovider.ChainVersionHeight(ctx, xchain.ChainVersion{
		ID:        j.SrcChain,
		ConfLevel: xchain.ConfFinalized,
	})
	if err != nil {
		return errors.Wrap(err, "chain height")
	}

	id, err := solvernet.OpenOrder(ctx, j.Network, evmchain.IDMockL1, backends, j.Owner, j.OrderData)
	if err != nil {
		return errors.Wrap(err, "open order")
	}

	log.Debug(ctx, "Flowgen order opened", "id", id)

	status, err := waitForFinalStatus(ctx, xprovider, j, id, height+1)
	if err != nil {
		return errors.Wrap(err, "wait for status")
	}
	log.Info(ctx, "Flowgen order finalized", "id", id, "status", status)

	return nil
}

// waitForFinalStatus fetches all event logs starting from the provided height and
// tries to find a final status of the specified order and returns it. Since we assume that
// all orders will eventually be rejected, closed or claimed, the function never terminates.
func waitForFinalStatus(
	ctx context.Context,
	xprovider xchain.Provider,
	j *types.Job,
	orderID solvernet.OrderID,
	height uint64,
) (solvernet.OrderStatus, error) {
	addrs, err := contracts.GetAddresses(ctx, j.Network)
	if err != nil {
		panic(err)
	}

	statusChan := make(chan solvernet.OrderStatus, 1)
	errChan := make(chan error, 1)

	proc := func(_ context.Context, _ uint64, logs []ethtypes.Log) error {
		for _, l := range logs {
			event, ok := solvernet.EventByTopic(l.Topics[0])
			if !ok {
				return errors.New("unknown event", "topic", l.Topics[0])
			}

			id := solvernet.OrderID(l.Topics[1])
			if id != orderID {
				continue
			}

			switch event.Status {
			case solvernet.StatusInvalid, solvernet.StatusRejected, solvernet.StatusClosed, solvernet.StatusClaimed:
				statusChan <- event.Status
			default:
			}
		}

		return nil
	}

	go func() {
		req := xchain.EventLogsReq{
			ChainID:       j.SrcChain,
			ConfLevel:     xchain.ConfLatest,
			Height:        height,
			FilterAddress: addrs.SolverNetInbox,
			FilterTopics:  solvernet.AllEventTopics(),
		}

		err := xprovider.StreamEventLogs(ctx, req, proc)
		if err != nil {
			errChan <- err
		}
	}()

	select {
	case status := <-statusChan:
		return status, nil
	case err := <-errChan:
		return solvernet.StatusInvalid, errors.Wrap(err, "stream event logs")
	case <-ctx.Done():
		return solvernet.StatusInvalid, errors.Wrap(ctx.Err(), "context done")
	}
}
