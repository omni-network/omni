package db

import (
	"context"
	"math/big"
	"sync"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/cctp/types"
	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/common"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/types/ormerrors"
	db "github.com/cosmos/cosmos-db"
)

// DB provides access to CCTP storage.
type DB struct {
	mu          sync.Mutex
	msgTable    MsgSendUSDCTable
	cursorTable CursorTable
}

// New returns a new database instance.
func New(db db.DB) (*DB, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_lib_cctp_db_db_proto.Path()},
	}}

	storeSvc := dbStoreService{DB: db}

	modDB, err := ormdb.NewModuleDB(schema, ormdb.ModuleDBOptions{KVStoreService: storeSvc})
	if err != nil {
		return nil, errors.Wrap(err, "create ormdb module db")
	}

	dbStore, err := NewDbStore(modDB)
	if err != nil {
		return nil, errors.Wrap(err, "create store")
	}

	return &DB{
		msgTable:    dbStore.MsgSendUSDCTable(),
		cursorTable: dbStore.CursorTable(),
	}, nil
}

// SetCursor stores a new cursor for a chain.
func (db *DB) SetCursor(ctx context.Context, chainID uint64, height uint64) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if err := db.cursorTable.Save(ctx, &Cursor{
		ChainId:     chainID,
		BlockHeight: height,
	}); err != nil {
		return errors.Wrap(err, "insert cursor")
	}

	return nil
}

// GetCursor retrieves a cursor by chain ID.
func (db *DB) GetCursor(ctx context.Context, chainID uint64) (uint64, bool, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	cursor, err := db.cursorTable.Get(ctx, chainID)
	if ormerrors.IsNotFound(err) {
		return 0, false, nil
	} else if err != nil {
		return 0, false, errors.Wrap(err, "get cursor")
	}

	return cursor.GetBlockHeight(), true, nil
}

// InsertMsg stores a new message.
func (db *DB) InsertMsg(ctx context.Context, msg types.MsgSendUSDC) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if err := msg.Validate(); err != nil {
		return errors.Wrap(err, "validate msg")
	}

	if err := db.msgTable.Insert(ctx, msgToProto(msg)); err != nil {
		return errors.Wrap(err, "insert msg")
	}

	return nil
}

// SetMsg replaces an existing message with a new one (by tx hash).
func (db *DB) SetMsg(ctx context.Context, msg types.MsgSendUSDC) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if err := msg.Validate(); err != nil {
		return errors.Wrap(err, "validate msg")
	}

	if err := db.msgTable.Save(ctx, msgToProto(msg)); err != nil {
		return errors.Wrap(err, "save msg")
	}

	return nil
}

// GetMsg retrieves a message by tx hash.
func (db *DB) GetMsg(ctx context.Context, txHash common.Hash) (types.MsgSendUSDC, bool, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	proto, err := db.msgTable.Get(ctx, txHash[:])
	if ormerrors.IsNotFound(err) {
		return types.MsgSendUSDC{}, false, nil
	} else if err != nil {
		return types.MsgSendUSDC{}, false, errors.Wrap(err, "get msg")
	}

	msg, err := msgFromProto(proto)
	if err != nil {
		return types.MsgSendUSDC{}, true, errors.Wrap(err, "from proto")
	}

	return msg, true, nil
}

// MsgFilter contains optional filters for GetMsgsBy.
type MsgFilter struct {
	DestChainID uint64          // Filter by destination chain ID.
	Status      types.MsgStatus // Filter by message status.
}

// GetMsgs returns all messages.
func (db *DB) GetMsgs(ctx context.Context) ([]types.MsgSendUSDC, error) {
	return db.GetMsgsBy(ctx, MsgFilter{})
}

// GetMsgsBy returns messages filtered by the provided filter.
func (db *DB) GetMsgsBy(ctx context.Context, filter MsgFilter) ([]types.MsgSendUSDC, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	iter, err := db.msgTable.List(ctx, MsgSendUSDCTxHashIndexKey{})
	if err != nil {
		return nil, errors.Wrap(err, "list msg")
	}
	defer iter.Close()

	msgs := make([]types.MsgSendUSDC, 0)
	for iter.Next() {
		proto, err := iter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "iter value")
		}

		msg, err := msgFromProto(proto)
		if err != nil {
			return nil, errors.Wrap(err, "from proto")
		}

		// Filter by DestChainID
		if filter.DestChainID != 0 && msg.DestChainID != filter.DestChainID {
			continue
		}

		// Filter by Status
		if filter.Status != types.MsgStatusUnknown && msg.Status != filter.Status {
			continue
		}

		msgs = append(msgs, msg)
	}

	return msgs, nil
}

// DeleteMsg deletes a message by tx hash.
// A msg should be deleted after it's confirmed received.
func (db *DB) DeleteMsg(ctx context.Context, txHash common.Hash) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	msg := &MsgSendUSDC{
		TxHash: txHash[:],
	}

	if err := db.msgTable.Delete(ctx, msg); err != nil {
		return errors.Wrap(err, "delete msg")
	}

	return nil
}

// HasMsg checks if a message exists in the database by its tx hash.
func (db *DB) HasMsg(ctx context.Context, txHash common.Hash) (bool, error) {
	_, ok, err := db.GetMsg(ctx, txHash)
	if err != nil {
		return false, errors.Wrap(err, "get msg")
	}

	return ok, nil
}

type dbStoreService struct {
	db.DB
}

func (db dbStoreService) OpenKVStore(context.Context) store.KVStore {
	return db
}

func msgToProto(msg types.MsgSendUSDC) *MsgSendUSDC {
	return &MsgSendUSDC{
		TxHash:       msg.TxHash[:],
		BlockHeight:  msg.BlockHeight,
		MessageHash:  msg.MessageHash[:],
		SrcChainId:   msg.SrcChainID,
		DestChainId:  msg.DestChainID,
		Amount:       msg.Amount.Bytes(),
		MessageBytes: msg.MessageBytes,
		Recipient:    msg.Recipient[:],
		Status:       int32(msg.Status),
	}
}

func msgFromProto(msg *MsgSendUSDC) (types.MsgSendUSDC, error) {
	msgHash, err := cast.EthHash(msg.GetMessageHash())
	if err != nil {
		return types.MsgSendUSDC{}, errors.Wrap(err, "cast msg hash")
	}

	txHash, err := cast.EthHash(msg.GetTxHash())
	if err != nil {
		return types.MsgSendUSDC{}, errors.Wrap(err, "cast tx hash")
	}

	recipient, err := cast.EthAddress(msg.GetRecipient())
	if err != nil {
		return types.MsgSendUSDC{}, errors.Wrap(err, "cast recipient")
	}

	return types.MsgSendUSDC{
		TxHash:       txHash,
		BlockHeight:  msg.GetBlockHeight(),
		MessageHash:  msgHash,
		SrcChainID:   msg.GetSrcChainId(),
		DestChainID:  msg.GetDestChainId(),
		Amount:       new(big.Int).SetBytes(msg.GetAmount()),
		MessageBytes: msg.GetMessageBytes(),
		Recipient:    recipient,
		Status:       types.MsgStatus(msg.GetStatus()),
	}, nil
}
