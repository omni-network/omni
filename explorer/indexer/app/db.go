package app

import (
	"context"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/xchain"
)

func newCallback(client *ent.Client) xchain.ProviderCallback {
	return func(ctx context.Context, block xchain.Block) error {
		var err error
		tx, err := client.BeginTx(ctx, nil)
		if err != nil {
			return errors.Wrap(err, "begin transaction")
		}
		defer func() {
			if err == nil {
				return
			}
			err = tx.Rollback()
			if err != nil {
				log.Error(ctx, "Rollback failed", err)
			}
			log.Info(ctx, "Rolledback transaction")
		}()

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

		if err := tx.Commit(); err != nil {
			return errors.Wrap(err, "commit transaction")
		}

		log.Info(ctx, "Inserted xblock",
			"msgs", len(block.Msgs),
			"receipts", len(block.Receipts),
		)

		return nil
	}
}

func insertBlock(ctx context.Context, tx *ent.Tx, block xchain.Block) (*ent.Block, error) {
	b, err := tx.Block.Create().
		SetBlockHeight(block.BlockHeight).
		SetBlockHash(block.BlockHash[:]).
		SetSourceChainID(block.SourceChainID).
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
			Save(ctx)
		if err != nil {
			return errors.Wrap(err, "inserting message")
		}
	}

	return nil
}
