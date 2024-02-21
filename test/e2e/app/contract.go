package app

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/netman"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func StartSendingXMsgs(ctx context.Context, portals map[uint64]netman.Portal, batches ...int) <-chan error {
	log.Info(ctx, "Generating cross chain messages async", "batches", batches)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	errChan := make(chan error, 1)

	go func() {
		for i, count := range batches {
			log.Debug(ctx, "Sending xmsgs", "batch", i, "count", count)
			err := SendXMsgs(ctx, portals, count)
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
func SendXMsgs(ctx context.Context, portals map[uint64]netman.Portal, batch int) error {
	type sentTuple struct {
		TX     *ethtypes.Transaction
		SentAt uint64
	}
	allSends := make(map[uint64][]sentTuple)
	for fromChainID, from := range portals {
		nonce, err := from.Client.PendingNonceAt(ctx, from.TxOptsFrom())
		if err != nil {
			return errors.Wrap(err, "pending nonce", "chain", from.Chain.Name)
		}

		for _, to := range portals {
			if from.Chain.ID == to.Chain.ID {
				continue
			}

			for i := 0; i < batch; i++ {
				h, err := from.Client.BlockNumber(ctx)
				if err != nil {
					return errors.Wrap(err, "block number")
				}

				tx, err := xcall(ctx, from, to.Chain.ID, nonce)
				if err != nil {
					return errors.Wrap(err, "batch_offset", i)
				}

				allSends[fromChainID] = append(allSends[fromChainID], sentTuple{
					TX:     tx,
					SentAt: h,
				})
				nonce++
			}
		}
	}

	for chainID, tups := range allSends {
		portal := portals[chainID]
		for i, tup := range tups {
			receipt, err := bind.WaitMined(ctx, portal.Client, tup.TX)
			if err != nil {
				return errors.Wrap(err, "wait mined", "chain", portal.Chain.Name, "tx_index", i)
			}

			// Only log slow confirmations
			if delta := receipt.BlockNumber.Uint64() - tup.SentAt; delta > 0 {
				log.Debug(ctx, "Sent xmsg mined",
					"chain", portal.Chain.Name,
					"sent_at", tup.SentAt, "mined_at", receipt.BlockNumber.Uint64(),
					"delta", receipt.BlockNumber.Uint64()-tup.SentAt)
			}
		}
	}

	return nil
}

// xcall sends a ethereum transaction to the portal contract, triggering a xcall.
func xcall(ctx context.Context, from netman.Portal, destChainID uint64, nonce uint64) (*ethtypes.Transaction, error) {
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

	txOpts := from.TxOpts(ctx, fee)
	txOpts.Nonce = big.NewInt(int64(nonce))

	tx, err := from.Contract.Xcall(txOpts, destChainID, to, data)
	if err != nil {
		return nil, errors.Wrap(err, "xcall",
			"src_chain", from.Chain.Name,
			"dst_chain_id", destChainID,
		)
	}

	return tx, nil
}
