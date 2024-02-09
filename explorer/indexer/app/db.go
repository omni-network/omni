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

		if err := insertBlockTX(ctx, tx, block); err != nil {
			if err := tx.Rollback(); err != nil { // Just log on rollback failure
				log.Error(ctx, "Rollback transaction failed", err)
			}

			return errors.Wrap(err, "insert xblock")
		}

		log.Info(ctx, "Inserted xblock",
			"msgs", len(block.Msgs),
			"receipts", len(block.Receipts),
		)

		return nil
	}
}

// insertBlockTX inserts the block as part of a tx and commits it.
// The caller should handle rollback on any error.
func insertBlockTX(ctx context.Context, tx *ent.Tx, block xchain.Block) error {
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

	if err := incrementCursor(ctx, tx, block.SourceChainID, block.BlockHeight); err != nil {
		return errors.Wrap(err, "increment cursor")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "commit transaction")
	}

	return nil
}

// incrementCursor increments the cursor for the given chainID (it ensures it matches height).
func incrementCursor(ctx context.Context, tx *ent.Tx, chainID, height uint64) error {
	cursor, ok, err := getCursor(ctx, tx.XProviderCursor, chainID)
	if err != nil {
		return errors.Wrap(err, "query cursor")
	} else if !ok {
		return errors.New("cursor not found")
	} else if cursor.Height != 0 && cursor.Height != height-1 {
		// Sanity check, we MUST insert sequentially (after 0).
		return errors.New("unexpected cursor vs block height mismatch [BUG]")
	}

	cursor.Height = height
	if _, err := tx.XProviderCursor.UpdateOne(cursor).Save(ctx); err != nil {
		return errors.Wrap(err, "update cursor")
	}

	return nil
}

// getCursor returns the current cursor for the given chainID, or false if it doesn't exist.
func getCursor(ctx context.Context, client *ent.XProviderCursorClient, chainID uint64,
) (*ent.XProviderCursor, bool, error) {
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
		SetBlockHeight(block.BlockHeight).
		SetBlockHash(block.BlockHash[:]).
		SetSourceChainID(block.SourceChainID).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "inserting block to db")
	}

	return b, nil
}

func insertMessages(ctx context.Context, tx *ent.Tx, block xchain.Block, dbBlock *ent.Block) error {
	for _, msg := range block.Msgs {
		_, err := tx.Msg.Create().
			SetBlock(dbBlock).
			SetBlockID(dbBlock.ID).
			SetData(msg.Data). // mock provider has no data? should this be nullable?
			SetDestAddress(msg.DestAddress[:]).
			SetDestChainID(msg.DestChainID).
			SetCreatedAt(time.Now()).
			// mock provider error: failed to encode args[4]: unable to encode 0xb70a309734fbef40
			// into binary format for int8 (OID 20): 13189407884695039808 is greater than maximum value for int64
			SetSourceChainID(msg.SourceChainID).
			SetDestGasLimit(msg.DestGasLimit).
			SetSourceMsgSender(msg.SourceMsgSender[:]).
			SetStreamOffset(msg.StreamOffset).
			SetTxHash(msg.TxHash[:]).
			Save(ctx)
		if err != nil {
			return errors.Wrap(err, "inserting message")
		}
	}

	return nil
}

func insertReceipts(ctx context.Context, tx *ent.Tx, block xchain.Block, dbBlock *ent.Block) error {
	for _, receipt := range block.Receipts {
		// FIXME: fix receipt schema we have some mismatch fields and missing fields
		_, err := tx.Receipt.Create().
			SetBlock(dbBlock).
			SetBlockID(dbBlock.ID).
			SetDestChainID(receipt.DestChainID).
			SetSourceChainID(receipt.SourceChainID).
			SetStreamOffset(receipt.StreamOffset).
			SetCreatedAt(time.Now()).
			Save(ctx)
		if err != nil {
			return errors.Wrap(err, "inserting message")
		}
	}

	return nil
}
