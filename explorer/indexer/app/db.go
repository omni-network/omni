package app

import (
	"context"
	"time"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/xprovidercursor"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"
)

// newCallback returns the indexer xprovider callback that
// inserts xblocks into the DB. It also updates cursors.
func newCallback(client *ent.Client) xchain.ProviderCallback {
	return func(ctx context.Context, block xchain.Block) error {
		tx, err := client.BeginTx(ctx, nil)
		if err != nil {
			return errors.Wrap(err, "begin transaction")
		}

		if err := InsertBlockTX(ctx, tx, block); err != nil {
			if err := tx.Rollback(); err != nil { // Just log on rollback failure
				log.Error(ctx, "Rollback transaction failed", err)
			}

			return errors.Wrap(err, "insert xblock")
		}

		log.Info(ctx, "Inserted xblock", "chain", block.SourceChainID, "offset", block.BlockOffset, "msgs", len(block.Msgs), "receipts", len(block.Receipts))

		return nil
	}
}

// InsertBlockTX inserts the block as part of a tx and commits it.
// The caller should handle rollback on any error.
func InsertBlockTX(ctx context.Context, tx *ent.Tx, block xchain.Block) error {
	insertedBlock, err := insertBlock(ctx, tx, block)
	if err != nil {
		return errors.Wrap(err, "insert block")
	}

	err = insertMessages(ctx, tx, block, insertedBlock)
	if err != nil {
		return errors.Wrap(err, "insert messages")
	}

	err = insertReceipts(ctx, tx, block, insertedBlock)
	if err != nil {
		return errors.Wrap(err, "insert receipts")
	}

	if err := incrementCursor(ctx, tx, block.SourceChainID, block.BlockHeight, block.BlockOffset); err != nil {
		return errors.Wrap(err, "increment cursor")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "commit transaction")
	}

	return nil
}

// incrementCursor increments the cursor for the given chainID (it ensures it matches height).
func incrementCursor(ctx context.Context, tx *ent.Tx, chainID, height, offset uint64) error {
	c, ok, err := cursor(ctx, tx.XProviderCursor, chainID)
	if err != nil {
		return errors.Wrap(err, "query cursor")
	} else if !ok {
		return errors.New("cursor not found")
	}
	// } else if c.Offset > 1 && c.Offset != offset-1 {
	// 	// Sanity check, we MUST insert sequentially (after 1).
	// 	return errors.New("unexpected cursor vs block offset mismatch [BUG]", "cursor_height", c.Height, "block_height", height, "cursor_offset", c.Offset, "block_offset", offset)
	// }

	if _, err := tx.XProviderCursor.UpdateOne(c).SetHeight(height).SetOffset(offset).Save(ctx); err != nil {
		return errors.Wrap(err, "update cursor")
	}

	return nil
}

// cursor returns the current cursor for the given chainID, or false if it doesn't exist.
func cursor(ctx context.Context, client *ent.XProviderCursorClient, chainID uint64) (*ent.XProviderCursor, bool, error) {
	cursor, err := client.Query().Where(xprovidercursor.ChainID(chainID)).Only(ctx)
	if ent.IsNotFound(err) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, errors.Wrap(err, "query cursor")
	}

	return cursor, true, nil
}

func insertBlock(ctx context.Context, tx *ent.Tx, block xchain.Block) (*ent.Block, error) {
	b, err := tx.Block.Create().
		SetHeight(block.BlockHeight).
		SetHash(block.BlockHash[:]).
		SetChainID(block.SourceChainID).
		SetOffset(block.BlockOffset).
		SetTimestamp(block.Timestamp).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "inserting block to db")
	}

	return b, nil
}

func insertMessages(ctx context.Context, tx *ent.Tx, block xchain.Block, dbBlock *ent.Block) error {
	for _, m := range block.Msgs {
		msg, err := tx.Msg.Create().
			SetData(m.Data).
			SetTo(m.DestAddress[:]).
			SetDestChainID(m.DestChainID).
			SetCreatedAt(time.Now()).
			SetSourceChainID(m.SourceChainID).
			SetGasLimit(m.DestGasLimit).
			SetSender(m.SourceMsgSender[:]).
			SetOffset(m.StreamOffset).
			SetTxHash(m.TxHash[:]).
			SetCreatedAt(time.Now()).
			Save(ctx)
		if err != nil {
			return errors.Wrap(err, "inserting message")
		}

		_, err = tx.Block.UpdateOne(dbBlock).AddMsgs(msg).Save(ctx)
		if err != nil {
			return errors.Wrap(err, "setting message edge to block")
		}
	}

	return nil
}

func insertReceipts(ctx context.Context, tx *ent.Tx, block xchain.Block, dbBlock *ent.Block) error {
	for _, r := range block.Receipts {
		receipt, err := tx.Receipt.Create().
			SetGasUsed(r.GasUsed).
			SetBlockHash(block.BlockHash[:]).
			SetDestChainID(r.DestChainID).
			SetSourceChainID(r.SourceChainID).
			SetOffset(r.StreamOffset).
			SetSuccess(r.Success).
			SetRelayerAddress(r.RelayerAddress.Bytes()).
			SetTxHash(r.TxHash.Bytes()).
			SetCreatedAt(time.Now()).
			Save(ctx)
		if err != nil {
			return errors.Wrap(err, "inserting message")
		}

		_, err = tx.Block.UpdateOne(dbBlock).AddReceipts(receipt).Save(ctx)
		if err != nil {
			return errors.Wrap(err, "setting receipt edge to block")
		}
	}

	return nil
}
