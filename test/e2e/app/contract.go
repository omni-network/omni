package app

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/netman"
	"github.com/omni-network/omni/test/e2e/send"

	"github.com/ethereum/go-ethereum/params"
)

func StartSendingXMsgs(ctx context.Context, portals map[uint64]netman.Portal, txManager send.Manager, batches ...int) <-chan error {
	log.Info(ctx, "Generating cross chain messages async", "batches", batches)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	errChan := make(chan error, 1)

	go func() {
		for i, count := range batches {
			log.Info(ctx, "Sending xmsgs", "batch", i, "count", count)
			err := SendXMsgs(ctx, portals, txManager, count)
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
func SendXMsgs(ctx context.Context, portals map[uint64]netman.Portal, txManager send.Manager, batch int) error {
	for fromChainID, from := range portals {
		for _, to := range portals {
			if from.Chain.ID == to.Chain.ID {
				continue
			}

			args := send.XCallArgs{
				DestChainID: to.Chain.ID,
				Address:     to.DeployInfo.PortalAddress,
				Data:        []byte{},
				GasLimit:    params.TxGas,
			}
			value := big.NewInt(params.GWei)
			log.Info(ctx, "Sending xcall", "from", from.Chain.Name, "to", to.Chain.Name, "value", value, "gasLimit", args.GasLimit, "destChainID", args.DestChainID, "address", args.Address.String())
			receipt, err := txManager.SendXCall(ctx, args, value, fromChainID, to.DeployInfo.PortalAddress)
			if err != nil {
				return errors.Wrap(err, "send xcall", "from", from.Chain.Name, "to", to.Chain.Name)
			}
			log.Info(ctx, "Receipt", "status", receipt.Status, "gas_used", receipt.GasUsed, "tx_hash", receipt.TxHash.String())
		}
	}

	return nil
}
