package indexer

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/types/ormerrors"
	db "github.com/cosmos/cosmos-db"
)

// emptyBlockCursorUpdate defines the Nth block to update cursors for empty blocks
// This avoids needing to catchup a lot on startup for quite streams.
const emptyBlockCursorUpdate = 100

// unknown is the string used for unknown values.
const unknown = "unknown"

var confLevel = xchain.ConfFinalized

// Start streams goroutines that streams xblocks and indexes xmsgs vs xreceipt metrics.
func Start(
	ctx context.Context,
	network netconf.Network,
	xprov xchain.Provider,
	db db.DB,
) error {
	xdapps, err := newXDapps(ctx, network)
	if err != nil {
		return err
	}

	indexer, err := newIndexer(db, xprov, network.StreamName, xdapps)
	if err != nil {
		return errors.Wrap(err, "create indexer")
	}

	cursors, err := indexer.cursors(ctx)
	if err != nil {
		return err
	}

	for _, chain := range network.Chains {
		req := xchain.ProviderRequest{
			ChainID:   chain.ID,
			ConfLevel: confLevel,
			Height:    cursors[xchain.ChainVersion{ID: chain.ID, ConfLevel: confLevel}],
		}
		if err := xprov.StreamAsync(ctx, req, indexer.index); err != nil {
			return err
		}
	}

	go deleteForever(ctx, indexer)

	return nil
}

// newXDapps returns a map of known xdapps (xmsg source contracts) for the given network.
func newXDapps(ctx context.Context, network netconf.Network) (map[common.Address]string, error) {
	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		return nil, err
	}

	// TODO(corver): Add more addresses here

	return map[common.Address]string{
		addrs.SolverNetOutbox: "solver",
	}, nil
}

// newIndexer creates a new indexer using the provided DB.
func newIndexer(
	db db.DB,
	xprov xchain.Provider,
	streamNamer func(xchain.StreamID) string,
	xdapps map[common.Address]string,
) (*indexer, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_monitor_xmonitor_indexer_indexer_proto.Path()},
	}}

	storeSvc := dbStoreService{DB: db}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: storeSvc})
	if err != nil {
		return nil, errors.Wrap(err, "create ormdb module db")
	}

	dbStore, err := NewIndexerStore(modDB)
	if err != nil {
		return nil, errors.Wrap(err, "create store")
	}

	return &indexer{
		xprov:        xprov,
		streamNamer:  streamNamer,
		blockTable:   dbStore.BlockTable(),
		msgLinkTable: dbStore.MsgLinkTable(),
		cursorTable:  dbStore.CursorTable(),
		sampleFunc:   instrumentSample,
		xdapps:       xdapps,
	}, nil
}

// indexer indexes xchain blocks and messages.
type indexer struct {
	mu           sync.RWMutex
	xprov        xchain.Provider
	blockTable   BlockTable
	msgLinkTable MsgLinkTable
	cursorTable  CursorTable
	streamNamer  func(xchain.StreamID) string
	xdapps       map[common.Address]string
	sampleFunc   func(sample)
}

// cursors returns the indexed block height for each chain.
func (i *indexer) cursors(ctx context.Context) (map[xchain.ChainVersion]uint64, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	iter, err := i.cursorTable.List(ctx, CursorPrimaryKey{})
	if err != nil {
		return nil, errors.Wrap(err, "list cursors")
	}
	defer iter.Close()

	resp := make(map[xchain.ChainVersion]uint64)
	for iter.Next() {
		cursor, err := iter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "get cursor value")
		}

		resp[xchain.ChainVersion{
			ID:        cursor.GetChainId(),
			ConfLevel: xchain.ConfLevel(cursor.GetConfLevel()),
		}] = cursor.GetBlockHeight()
	}

	return resp, nil
}

// delete deletes all blocks (and msg links) that have been fully indexed.
func (i *indexer) delete(ctx context.Context) ([]xchain.BlockHeader, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	// Iterate over all blocks
	blockIter, err := i.blockTable.List(ctx, BlockPrimaryKey{})
	if err != nil {
		return nil, errors.Wrap(err, "list blocks")
	}
	defer blockIter.Close()

	var deleted []xchain.BlockHeader

	for blockIter.Next() {
		blockDB, err := blockIter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "get block value")
		}

		block, err := blockDB.XChainBlock()
		if err != nil {
			return nil, err
		}

		var links []*MsgLink
		indexed := true

		// Ensure all messages have been matched to receipts
		for _, msg := range block.Msgs {
			link, exist, err := i.getLink(ctx, msg.MsgID)
			if err != nil {
				return nil, err
			} else if !exist {
				// Previously deleted, count as matched
				continue
			}

			if link.GetMsgBlockId() == 0 || link.GetReceiptBlockId() == 0 {
				indexed = false
				break
			}

			links = append(links, link)
		}
		if !indexed {
			continue
		}

		// Ensure all receipts have been matched to messages
		for _, receipt := range block.Receipts {
			link, exist, err := i.getLink(ctx, receipt.MsgID)
			if err != nil {
				return nil, err
			} else if !exist {
				// Previously deleted, count as matched
				continue
			}

			if link.GetMsgBlockId() == 0 || link.GetReceiptBlockId() == 0 {
				indexed = false
				break
			}

			links = append(links, link)
		}
		if !indexed {
			continue
		}

		// All receipts and messages of the block has been matched/indexed, delete it.

		if err := i.blockTable.Delete(ctx, blockDB); err != nil {
			return nil, errors.Wrap(err, "delete block")
		}

		for _, link := range links {
			if err := i.msgLinkTable.Delete(ctx, link); err != nil {
				return nil, errors.Wrap(err, "delete block")
			}
		}

		deleted = append(deleted, block.BlockHeader)
	}

	return deleted, nil
}

// updateCursor updates the cursor to the provided chain to the provided height.
func (i *indexer) updateCursor(ctx context.Context, block xchain.Block) error {
	err := i.cursorTable.Save(ctx, &Cursor{
		ChainId:     block.ChainID,
		ConfLevel:   uint32(confLevel),
		BlockHeight: block.BlockHeight,
	})
	if err != nil {
		return errors.Wrap(err, "save cursor")
	}

	return nil
}

// index indexes the given block.
func (i *indexer) index(ctx context.Context, block xchain.Block) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	// Skip empty blocks
	if len(block.Msgs) == 0 && len(block.Receipts) == 0 {
		// Update cursor for every Nth empty block
		if block.BlockHeight%emptyBlockCursorUpdate == 0 {
			if err := i.updateCursor(ctx, block); err != nil {
				return err
			}
		}

		return nil
	}

	// Marshal block (we don't store all block fields explicitly)
	bz, err := json.Marshal(block) //nolint:musttag // TODO: Rather use protobuf
	if err != nil {
		return errors.Wrap(err, "marshal block")
	}

	// Insert block
	id, err := i.blockTable.InsertReturningId(ctx, &Block{
		ChainId:     block.ChainID,
		BlockHeight: block.BlockHeight,
		BlockHash:   block.BlockHash.Bytes(),
		BlockJson:   bz,
	})
	if errors.Is(err, ormerrors.UniqueKeyViolation) { // Idempotent
		existing, err := i.blockTable.GetByChainIdBlockHeightBlockHash(ctx, block.ChainID, block.BlockHeight, block.BlockHash.Bytes())
		if err != nil {
			return errors.Wrap(err, "get existing block")
		}
		id = existing.GetId()
	} else if err != nil {
		return errors.Wrap(err, "insert block")
	}

	// Upsert msg links
	for _, msg := range block.Msgs {
		if evmchain.IsDisabled(msg.DestChainID) {
			continue // Ignore msgs to disabled chains
		}

		link, _, err := i.getLink(ctx, msg.MsgID)
		if err != nil {
			return err
		}
		if link.GetMsgBlockId() != 0 && link.GetMsgBlockId() != id {
			return errors.New("mismatching msg block id [BUG]",
				"msg_id", msg.MsgID,
				"got", link.GetMsgBlockId(),
				"want", id,
			)
		}
		link.MsgBlockId = id
		if err := i.msgLinkTable.Save(ctx, link); err != nil {
			return errors.Wrap(err, "save msg link")
		}

		// Maybe instrument if both msg and receipt are indexed
		if link.GetMsgBlockId() != 0 && link.GetReceiptBlockId() != 0 {
			if err := i.instrumentMsg(ctx, link); err != nil {
				return err
			}
		}
	}

	// Upsert receipt links
	for _, receipt := range block.Receipts {
		if evmchain.IsDisabled(receipt.SourceChainID) {
			continue // Ignore receipts from disabled chains
		}

		link, _, err := i.getLink(ctx, receipt.MsgID)
		if err != nil {
			return err
		}
		if link.GetReceiptBlockId() != 0 && link.GetReceiptBlockId() != id {
			return errors.New("mismatching receipt block id [BUG]",
				"msg_id", receipt.MsgID,
				"got", link.GetReceiptBlockId(),
				"want", id,
			)
		}
		link.ReceiptBlockId = id
		if err := i.msgLinkTable.Save(ctx, link); err != nil {
			return errors.Wrap(err, "save msg link")
		}

		// Maybe instrument
		if link.GetMsgBlockId() != 0 && link.GetReceiptBlockId() != 0 {
			if err := i.instrumentMsg(ctx, link); err != nil {
				return err
			}
		}
	}

	return i.updateCursor(ctx, block) // Update cursor since we are done with this block
}

// getLink returns the msg link for the given id or a new one.
func (i *indexer) getLink(ctx context.Context, id xchain.MsgID) (*MsgLink, bool, error) {
	hash := id.Hash()
	link, err := i.msgLinkTable.Get(ctx, hash.Bytes())
	if ormerrors.IsNotFound(err) {
		return &MsgLink{
			IdHash: hash.Bytes(),
		}, false, nil
	} else if err != nil {
		return nil, false, errors.Wrap(err, "get msg link")
	}

	return link, true, nil
}

// instrumentMsg instruments the message vs receipt metrics.
func (i *indexer) instrumentMsg(ctx context.Context, link *MsgLink) error {
	// Get stuff
	msgBlockDB, err := i.blockTable.Get(ctx, link.GetMsgBlockId())
	if errors.Is(err, ormerrors.NotFound) {
		// Block probably deleted, ignore
		return nil
	} else if err != nil {
		return errors.Wrap(err, "get msg block")
	}
	receiptBlockDB, err := i.blockTable.Get(ctx, link.GetReceiptBlockId())
	if errors.Is(err, ormerrors.NotFound) {
		// Block probably deleted, ignore
		return nil
	} else if err != nil {
		return errors.Wrap(err, "get receipt block")
	}

	msgBlock, err := msgBlockDB.XChainBlock()
	if err != nil {
		return err
	}
	receiptBlock, err := receiptBlockDB.XChainBlock()
	if err != nil {
		return err
	}

	var msg xchain.Msg
	for _, m := range msgBlock.Msgs {
		if link.IsMsg(m) {
			msg = m
		}
	}
	if msg.SourceChainID == 0 {
		return errors.New("msg not found in msg block [BUG]")
	}

	var receipt xchain.Receipt
	for _, r := range receiptBlock.Receipts {
		if link.IsReceipt(r) {
			receipt = r
		}
	}
	if receipt.SourceChainID == 0 {
		return errors.New("receipt not found in receipt block [BUG]")
	}

	override, err := isFuzzyOverride(ctx, i.xprov, receipt)
	if err != nil {
		return err
	}

	// defaults to unknown. evmchain.MetadataByID fails in fuzz tests
	srcChainName := unknown
	feeToken := unknown

	srcChain, ok := evmchain.MetadataByID(msg.SourceChainID)
	if ok {
		srcChainName = srcChain.Name

		// for now, fee token is always src chain native token
		feeToken = srcChain.NativeToken.Symbol
	}

	// Instrument sample
	s := sample{
		Stream:        i.streamNamer(msg.StreamID),
		XDApp:         i.xdapp(msg.SourceMsgSender),
		SrcChain:      srcChainName,
		FeeToken:      feeToken,
		FeeAmount:     msg.Fees,
		Latency:       receiptBlock.Timestamp.Sub(msgBlock.Timestamp),
		Success:       receipt.Success,
		ExcessGas:     umath.SubtractOrZero(msg.DestGasLimit, receipt.GasUsed),
		FuzzyOverride: override,
	}
	i.sampleFunc(s)

	log.Info(ctx, "Indexed xchain message",
		"stream", s.Stream,
		"offset", msg.StreamOffset,
		"success", s.Success,
		"gas_usage", fmt.Sprintf("%.1f%%", float64(receipt.GasUsed)/float64(msg.DestGasLimit)),
		"latency", s.Latency,
		"msg_tx", msg.TxHash,
		"receipt_tx", receipt.TxHash,
		"id_hash", strings.TrimPrefix(msg.Hash().Hex(), "0x"), // Trim 0x prefix since omniscan doesn't use it in urls.
	)

	return nil
}

// isFuzzyOverride returns true if this was a fuzzy xchain message that was
// submitted by a finalized confirmation attestation (and not by a fuzzy attestation).
func isFuzzyOverride(ctx context.Context, xprov xchain.Provider, receipt xchain.Receipt) (bool, error) {
	if receipt.ShardID.ConfLevel().IsFinalized() {
		return false, nil // Only fuzzy shards can be overridden.
	}

	sub, err := xprov.GetSubmission(ctx, receipt.DestChainID, receipt.TxHash)
	if err != nil {
		return false, errors.Wrap(err, "get submission")
	}

	return sub.AttHeader.ChainVersion.ConfLevel.IsFinalized(), nil
}

// xdapp returns the xdapp name for the given sender address or "unknown".
func (i *indexer) xdapp(sender common.Address) string {
	resp, ok := i.xdapps[sender]
	if !ok {
		return unknown
	}

	return resp
}

// deleteForever deletes indexed blocks every X minutes.
func deleteForever(ctx context.Context, i *indexer) {
	ticker := time.NewTicker(time.Minute * 10)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			deleted, err := i.delete(ctx)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Warn(ctx, "Failed to delete indexed blocks (will retry)", err)
				continue
			}

			highest := make(map[uint64]uint64)
			for _, header := range deleted {
				if header.BlockHeight > highest[header.ChainID] {
					highest[header.ChainID] = header.BlockHeight
				}
			}

			log.Debug(ctx, "Deleted indexed blocks", "count", len(deleted), "highest", highest)
		}
	}
}

// dbStoreService wraps a cosmos-db instance and provides it via OpenKVStore.
type dbStoreService struct {
	db.DB
}

func (db dbStoreService) OpenKVStore(context.Context) store.KVStore {
	return db.DB
}
