package app

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/netman"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func StartSendingXMsgs(ctx context.Context, portals map[uint64]netman.Portal, batches ...int) <-chan error {
	log.Info(ctx, "Generating cross chain messages async", "batches", batches)
	errChan := make(chan error, 1)
	go func() {
		for _, count := range batches {
			err := SendXMsgs(ctx, portals, count)
			if ctx.Err() != nil {
				errChan <- ctx.Err()
				return
			} else if err != nil {
				errChan <- errors.Wrap(err, "send xmsgs")
				return
			}
		}
		errChan <- nil
	}()

	return errChan
}

// SendXMsgs sends <count> xmsgs from every chain to every other chain, then waits for them to be mined.
func SendXMsgs(ctx context.Context, portals map[uint64]netman.Portal, batch int) error {
	allTxs := make(map[uint64][]*ethtypes.Transaction)
	for fromChainID, from := range portals {
		for _, to := range portals {
			if from.Chain.ID == to.Chain.ID {
				continue
			}

			for i := 0; i < batch; i++ {
				tx, err := xcall(ctx, from, to.Chain.ID)
				if err != nil {
					return err
				}
				allTxs[fromChainID] = append(allTxs[fromChainID], tx)
			}
		}
	}

	for chainID, txs := range allTxs {
		portal := portals[chainID]
		for i, tx := range txs {
			if err := waitMined(ctx, portal.Client, tx); err != nil {
				return errors.Wrap(err, "wait mined", "chain", portal.Chain.Name, "tx_index", i)
			}
		}
	}

	return nil
}

// xcall sends a ethereum transaction to the portal contract, triggering a xcall.
func xcall(ctx context.Context, from netman.Portal, destChainID uint64) (*ethtypes.Transaction, error) {
	// TODO: use calls to actual contracts
	var data []byte = nil
	to := common.Address{}

	fee, err := from.Contract.FeeFor(&bind.CallOpts{}, destChainID, data)
	if err != nil {
		return nil, errors.Wrap(err, "feeFor",
			"source_chain", from.Chain.ID,
			"dest_chain", destChainID,
		)
	}

	txOpts := from.TxOpts(ctx)
	txOpts.Value = fee

	tx, err := from.Contract.Xcall(txOpts, destChainID, to, data)
	if err != nil {
		return nil, errors.Wrap(err, "xcall",
			"source_chain", from.Chain.ID,
			"dest_chain", destChainID,
		)
	}

	return tx, nil
}

func waitMined(ctx context.Context, ethCl *ethclient.Client, tx *ethtypes.Transaction) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	backoff := expbackoff.New(ctx, expbackoff.WithFastConfig())
	for ctx.Err() == nil {
		_, pending, err := ethCl.TransactionByHash(ctx, tx.Hash())
		if err != nil {
			return errors.Wrap(err, "tx by hash")
		}
		if !pending {
			return nil
		}
		backoff()
	}

	return errors.New("timeout waiting for tx to be mined")
}
