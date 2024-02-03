package app

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/test/e2e/netman"

	"github.com/ethereum/go-ethereum/common"
)

func StartSendingXMsgs(ctx context.Context, portals map[uint64]netman.Portal) error {
	log.Info(ctx, "Generating cross chain messages async")
	go func() {
		for ctx.Err() == nil {
			err := SendXMsgs(ctx, portals, 3)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Error(ctx, "Failed to send xmsgs, giving up", err)
				return
			}
			time.Sleep(time.Millisecond * 1000)
		}
	}()

	return nil
}

// SendXMsgs sends one xmsg from every chain to every other chain.
func SendXMsgs(ctx context.Context, portals map[uint64]netman.Portal, count int) error {
	for _, from := range portals {
		for _, to := range portals {
			if from.Chain.ID == to.Chain.ID {
				continue
			}

			for i := 0; i < count; i++ {
				if err := xcall(ctx, from, to.Chain.ID); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// xcall sends a ethereum transaction to the portal contract, triggering a xcall.
func xcall(ctx context.Context, from netman.Portal, destChainID uint64) error {
	_, err := from.Contract.Xcall(from.TxOpts(ctx), destChainID, common.Address{}, nil)
	if err != nil {
		return errors.Wrap(err, "xcall",
			"sourc_chain", from.Chain.ID,
			"dest_chain", destChainID,
		)
	}

	return nil
}
