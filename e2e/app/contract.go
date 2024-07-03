package app

import (
	"context"
	"time"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/netman"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"golang.org/x/sync/errgroup"
)

func StartSendingXMsgs(ctx context.Context, network netconf.ID, netman netman.Manager, backends ethbackend.Backends, batches ...int) <-chan error {
	log.Info(ctx, "Generating cross chain messages async", "batches", batches)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	errChan := make(chan error, 1)

	go func() {
		for i, count := range batches {
			log.Debug(ctx, "Sending xmsgs", "batch", i, "count", count)
			err := SendXMsgs(ctx, network, netman, backends, count)
			if ctx.Err() != nil {
				errChan <- ctx.Err()
				return
			} else if err != nil {
				errChan <- errors.Wrap(err, "send xmsgs", "batch", i)
				return
			}
		}
		errChan <- nil
		cancel()
	}()

	return errChan
}

// SendXMsgs sends <count> xmsgs from every chain to every other chain, then waits for them to be mined.
func SendXMsgs(ctx context.Context, network netconf.ID, netman netman.Manager, backends ethbackend.Backends, count int) error {
	sender := eoa.MustAddress(network, eoa.RoleTester)

	waiter := backends.NewWaiter()
	var eg errgroup.Group
	for _, from := range netman.Portals() {
		for _, to := range netman.Portals() {
			if from.Chain.ChainID == to.Chain.ChainID {
				continue
			}

			for i := 0; i < count; i++ {
				// Send async so whole batch included in same block. Important for testing.
				eg.Go(func() error {
					tx, err := xcall(ctx, backends, sender, from, to.Chain.ChainID)
					if err != nil {
						return errors.Wrap(err, "xcall")
					}

					waiter.Add(from.Chain.ChainID, tx)

					return nil
				})
			}
		}
	}

	// Wait for all sends to complete
	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "send xmsgs")
	}

	// Wait for all xmsgs to be mined, so next batch sent in subsequent block.
	if err := waiter.Wait(ctx); err != nil {
		return errors.Wrap(err, "wait xmsgs")
	}

	return nil
}

// xcall sends a ethereum transaction to the portal contract, triggering a xcall.
func xcall(ctx context.Context, backends ethbackend.Backends, sender common.Address, from netman.Portal, destChainID uint64,
) (*ethtypes.Transaction, error) {
	// TODO: use calls to actual contracts
	var data []byte
	to := common.HexToAddress("0x1234")
	gasLimit := uint64(100_000)

	fee, err := from.Contract.FeeFor(&bind.CallOpts{}, destChainID, data, gasLimit)
	if err != nil {
		return nil, errors.Wrap(err, "feeFor",
			"src_chain", from.Chain.Name,
			"dst_chain_id", destChainID,
		)
	}

	txOpts, _, err := backends.BindOpts(ctx, from.Chain.ChainID, sender)
	if err != nil {
		return nil, errors.Wrap(err, "bindOpts")
	}

	txOpts.Value = fee

	tx, err := from.Contract.Xcall(txOpts, destChainID, uint8(xchain.ConfFinalized), to, data, gasLimit)
	if err != nil {
		return nil, errors.Wrap(err, "xcall",
			"src_chain", from.Chain.Name,
			"dst_chain_id", destChainID,
		)
	}

	return tx, nil
}
