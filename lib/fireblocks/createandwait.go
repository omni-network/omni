package fireblocks

import (
	"context"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
)

// CreateAndWait creates a transaction and waits for it to be signed.
func (c Client) CreateAndWait(ctx context.Context, opts TransactionRequestOptions) (TransactionResponse, error) {
	resp, err := c.CreateTransaction(ctx, opts)
	if err != nil {
		return TransactionResponse{}, err
	}

	var attempt int
	queryTicker := time.NewTicker(c.cfg.QueryInterval)
	defer queryTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			return TransactionResponse{}, errors.Wrap(ctx.Err(), "context canceled")
		case <-queryTicker.C:
			resp, err := c.GetTransactionByID(ctx, resp.ID)
			if err != nil {
				return TransactionResponse{}, err
			}

			ok, err := isComplete(resp)
			if err != nil {
				return TransactionResponse{}, err
			} else if ok {
				return resp, nil
			}

			attempt++
			if attempt%c.cfg.LogFreqFactor == 0 {
				log.Warn(ctx, "Transaction not signed yet", nil,
					"attempt", attempt,
					"id", resp.ID,
					"status", resp.Status,
				)
			}
		}
	}
}

// isComplete returns true if the transaction is complete, false if still pending, or an error if it failed.
func isComplete(resp TransactionResponse) (bool, error) {
	switch resp.Status {
	case "COMPLETED":
		return true, nil
	case "CANCELED", "BLOCKED_BY_POLICY", "REJECTED", "FAILED":
		return false, errors.New("transaction failed", "status", resp.Status)
	default:
		return false, nil
	}
}
