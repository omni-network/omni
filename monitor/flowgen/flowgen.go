package flowgen

import (
	"context"
	"math/big"
	"slices"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/monitor/flowgen/bridging"
	"github.com/omni-network/omni/monitor/flowgen/symbiotic"
	"github.com/omni-network/omni/monitor/flowgen/types"
	sclient "github.com/omni-network/omni/solver/client"
	stypes "github.com/omni-network/omni/solver/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

func Start(
	ctx context.Context,
	network netconf.Network,
	ethClients map[uint64]ethclient.Client,
	keyPath string,
	solverAddress string,
) error {
	if network.ID == netconf.Mainnet {
		log.Info(ctx, "Skipping flowgen on mainnet")
		return nil
	}

	if keyPath == "" {
		return errors.New("private key is required")
	}

	privKey, err := ethcrypto.LoadECDSA(keyPath)
	if err != nil {
		return errors.Wrap(err, "load key", "path", keyPath)
	}

	backends, err := ethbackend.BackendsFromClients(ethClients, privKey)
	if err != nil {
		return errors.Wrap(err, "backends")
	}

	scl := sclient.New(solverAddress)

	return startWithBackends(ctx, network, backends, scl)
}

func startWithBackends(
	ctx context.Context,
	network netconf.Network,
	backends ethbackend.Backends,
	scl sclient.Client,
) error {
	j1 := bridging.Jobs(network.ID, backends, scl)

	j2, err := symbiotic.Jobs(ctx, backends, network.ID, scl)
	if err != nil {
		return errors.Wrap(err, "symbiotic jobs")
	}

	for _, job := range slices.Concat(j1, j2) {
		go func() {
			timer := time.NewTimer(0)
			defer timer.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-timer.C:
					timer.Reset(job.Cadence)
					jobsTotal.Inc()
					runCtx := log.WithCtx(ctx, "job", job.Name, "src_chain", network.ChainName(job.SrcChainID))

					ok, err := run(runCtx, network.ID, backends, job)
					if err != nil {
						log.Warn(runCtx, "Flowgen: job failed (will retry)", err)
						jobsFailed.Inc()
					}
					if !ok {
						// If the job was skipped, we retry earlier.
						retry := time.Minute
						if network.ID == netconf.Devnet {
							retry = 5 * time.Second
						}

						timer.Reset(retry)
					}
				}
			}
		}()
	}

	return nil
}

// run runs a job exactly once. It returns false if the job was skipped.
func run(ctx context.Context, network netconf.ID, backends ethbackend.Backends, job types.Job) (bool, error) {
	log.Debug(ctx, "Flowgen: running job")
	t0 := time.Now()

	results, err := job.OpenOrdersFunc(ctx)
	if err != nil {
		return false, errors.Wrap(err, "open orders")
	}

	log.Debug(ctx, "Flowgen: orders opened", "count", len(results))

	if len(results) == 0 {
		return false, nil
	}

	var token tokens.Token
	var minAmt, maxAmt *big.Int
	for i, result := range results {
		loopCtx := log.WithCtx(ctx, "order_id", result.OrderID, "index", i)

		if err := awaitClaimed(loopCtx, network, backends, job, result.OrderID); err != nil {
			return false, errors.Wrap(err, "await claimed")
		}

		amt := result.Data.Deposit.Amount
		if i == 0 {
			minAmt = amt
			maxAmt = amt

			tkn, ok := tokens.ByAddress(job.SrcChainID, result.Data.Deposit.Token)
			if !ok {
				return false, errors.New("src token not found", "address", result.Data.Deposit.Token)
			}
			token = tkn

			continue
		}

		if bi.LT(amt, minAmt) {
			minAmt = amt
		} else if bi.GT(amt, maxAmt) {
			maxAmt = amt
		}
	}

	log.Info(ctx, "Flowgen: orders claimed",
		"min_amount", token.FormatAmt(minAmt),
		"max_amount", token.FormatAmt(maxAmt),
		"count", len(results),
		"duration", time.Since(t0),
	)

	return true, nil
}

// awaitClaimed blocks until the order is claimed.
func awaitClaimed(ctx context.Context, network netconf.ID, backends ethbackend.Backends, job types.Job, orderID solvernet.OrderID) error {
	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return errors.New("contract addresses")
	}

	backend, err := backends.Backend(job.SrcChainID)
	if err != nil {
		return errors.Wrap(err, "src chain backend")
	}

	inbox, err := bindings.NewSolverNetInbox(addrs.SolverNetInbox, backend)
	if err != nil {
		return errors.Wrap(err, "create inbox contract")
	}

	var checks uint16
	const logFreq = 100
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
