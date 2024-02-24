package app

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/netman"
	"github.com/omni-network/omni/test/e2e/send"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func StartSendingXMsgs(ctx context.Context, netman netman.Manager, sender send.Sender, batches ...int) <-chan error {
	log.Info(ctx, "Generating cross chain messages async", "batches", batches)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	errChan := make(chan error, 1)

	go func() {
		for i, count := range batches {
			log.Debug(ctx, "Sending xmsgs", "batch", i, "count", count)
			err := SendXMsgs(ctx, netman, sender, count)
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
func SendXMsgs(ctx context.Context, netman netman.Manager, sender send.Sender, count int) error {
	type sentTuple struct {
		ChainID uint64
		TX      *ethtypes.Transaction
		SentAt  uint64
		Err     error
	}

	allSends := make(chan sentTuple)
	var expect int
	for fromChainID, from := range netman.Portals() {
		for _, to := range netman.Portals() {
			if from.Chain.ID == to.Chain.ID {
				continue
			}

			for i := 0; i < count; i++ {
				expect++
				// Send async so whole batch included in same block. Important for testing.
				go func() {
					txOpts, backend, err := sender.BindOpts(ctx, fromChainID)
					if err != nil {
						allSends <- sentTuple{ChainID: from.Chain.ID, Err: errors.Wrap(err, "deploy opts")}
						return
					}

					h, err := backend.BlockNumber(ctx)
					if err != nil {
						allSends <- sentTuple{ChainID: from.Chain.ID, Err: errors.Wrap(err, "block number")}
						return
					}

					tx, err := xcall(txOpts, from, to.Chain.ID)
					allSends <- sentTuple{
						ChainID: from.Chain.ID,
						TX:      tx,
						SentAt:  h,
						Err:     err,
					}
				}()
			}
		}
	}

	// Wait all batches to get mined.
	var i int
	for tup := range allSends {
		name := netman.Portals()[tup.ChainID].Chain.Name

		if tup.Err != nil {
			return errors.Wrap(tup.Err, "send xmsg", "chain", name)
		}

		_, backend, err := sender.BindOpts(ctx, tup.ChainID)
		if err != nil {
			return errors.Wrap(err, "deploy opts")
		}

		receipt, err := bind.WaitMined(ctx, backend, tup.TX)
		if err != nil {
			return errors.Wrap(err, "wait mined", "chain", name, "tx_index", i)
		}

		// Only log slow confirmations
		if delta := receipt.BlockNumber.Uint64() - tup.SentAt; delta > 2 {
			log.Debug(ctx, "Sent xmsg mined (slow)",
				"chain", name,
				"sent_at", tup.SentAt, "mined_at", receipt.BlockNumber.Uint64(),
				"delta", receipt.BlockNumber.Uint64()-tup.SentAt)
		}

		i++
		if expect == i {
			break
		}
	}

	return nil
}

// xcall sends a ethereum transaction to the portal contract, triggering a xcall.
func xcall(txOpts *bind.TransactOpts, from netman.Portal, destChainID uint64) (*ethtypes.Transaction, error) {
	// TODO: use calls to actual contracts
	var data []byte
	to := common.Address{}

	fee, err := from.Contract.FeeFor(&bind.CallOpts{}, destChainID, data)
	if err != nil {
		return nil, errors.Wrap(err, "feeFor",
			"src_chain", from.Chain.Name,
			"dst_chain_id", destChainID,
		)
	}

	txOpts.Value = fee

	tx, err := from.Contract.Xcall(txOpts, destChainID, to, data)
	if err != nil {
		return nil, errors.Wrap(err, "xcall",
			"src_chain", from.Chain.Name,
			"dst_chain_id", destChainID,
		)
	}

	return tx, nil
}
