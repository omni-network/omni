package usdt0

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/layerzero"

	"github.com/ethereum/go-ethereum/common"

	ormv1alpha1 "cosmossdk.io/api/cosmos/orm/v1alpha1"
	"cosmossdk.io/core/store"
	"cosmossdk.io/orm/model/ormdb"
	"cosmossdk.io/orm/types/ormerrors"
	db "github.com/cosmos/cosmos-db"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// DB provides access to USDT0 storage.
type DB struct {
	mu       sync.Mutex
	msgTable MsgSendUSDT0Table
}

// NewDB returns a new USDT0 database instance.
func NewDB(db db.DB) (*DB, error) {
	schema := &ormv1alpha1.ModuleSchemaDescriptor{SchemaFile: []*ormv1alpha1.ModuleSchemaDescriptor_FileEntry{
		{Id: 1, ProtoFileName: File_lib_usdt0_db_proto.Path()},
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
		msgTable: dbStore.MsgSendUSDT0Table(),
	}, nil
}

// InsertMsg stores a new message.
func (db *DB) InsertMsg(ctx context.Context, msg MsgSend) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	proto := msgToProto(msg)
	proto.CreatedAt = timestamppb.Now()

	if err := db.msgTable.Insert(ctx, proto); err != nil {
		return errors.Wrap(err, "insert msg")
	}

	return nil
}

// SetMsgStatus updates the status of a message by transaction hash.
func (db *DB) SetMsgStatus(ctx context.Context, txHash common.Hash, status layerzero.MsgStatus) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	stored, ok, err := db.getMsgUnsafe(ctx, txHash)
	if err != nil {
		return errors.Wrap(err, "get msg")
	} else if !ok {
		return errors.New("msg not found")
	}

	stored.Status = int32(statusToProto(status))

	if err := db.msgTable.Save(ctx, stored); err != nil {
		return errors.Wrap(err, "save msg")
	}

	return nil
}

// MsgFilter is a function that filters messages.
type MsgFilter func(msg MsgSend) bool

// GetMsgs returns all messages that match the provided filters.
func (db *DB) GetMsgs(ctx context.Context, filters ...MsgFilter) ([]MsgSend, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	iter, err := db.msgTable.List(ctx, MsgSendUSDT0TxHashIndexKey{})
	if err != nil {
		return nil, errors.Wrap(err, "list msg")
	}
	defer iter.Close()

	msgs := make([]MsgSend, 0)
	for iter.Next() {
		proto, err := iter.Value()
		if err != nil {
			return nil, errors.Wrap(err, "iter value")
		}

		msg, err := msgFromProto(proto)
		if err != nil {
			return nil, errors.Wrap(err, "from proto")
		}

		matches := true
		for _, filter := range filters {
			if !filter(msg) {
				matches = false
				break
			}
		}

		if matches {
			msgs = append(msgs, msg)
		}
	}

	return msgs, nil
}

// DeleteMsg deletes a message by transaction hash.
func (db *DB) DeleteMsg(ctx context.Context, txHash common.Hash) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	msg := &MsgSendUSDT0{
		TxHash: txHash[:],
	}

	if err := db.msgTable.Delete(ctx, msg); err != nil {
		return errors.Wrap(err, "delete msg")
	}

	return nil
}

// getMsgUnsafe retrieves a message by tx hash without locking the mutex.
func (db *DB) getMsgUnsafe(ctx context.Context, txHash common.Hash) (*MsgSendUSDT0, bool, error) {
	proto, err := db.msgTable.Get(ctx, txHash[:])
	if ormerrors.IsNotFound(err) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, errors.Wrap(err, "get msg")
	}

	return proto, true, nil
}

type dbStoreService struct {
	db.DB
}

func (db dbStoreService) OpenKVStore(context.Context) store.KVStore {
	return db
}

// statusToProto converts a layerzero.MsgStatus to a MsgStatus enum value.
func statusToProto(status layerzero.MsgStatus) MsgStatus {
	switch status {
	case layerzero.MsgStatusConfirming:
		return MsgStatus_MSG_STATUS_CONFIRMING
	case layerzero.MsgStatusInFlight:
		return MsgStatus_MSG_STATUS_INFLIGHT
	case layerzero.MsgStatusDelivered:
		return MsgStatus_MSG_STATUS_DELIVERED
	case layerzero.MsgStatusFailed:
		return MsgStatus_MSG_STATUS_FAILED
	case layerzero.MsgStatusPayloadStored:
		return MsgStatus_MSG_STATUS_PAYLOAD_STORED
	default:
		return MsgStatus_MSG_STATUS_UNKNOWN
	}
}

// statusFromProto converts a MsgStatus enum value to a layerzero.MsgStatus.
func statusFromProto(status MsgStatus) layerzero.MsgStatus {
	switch status {
	case MsgStatus_MSG_STATUS_CONFIRMING:
		return layerzero.MsgStatusConfirming
	case MsgStatus_MSG_STATUS_INFLIGHT:
		return layerzero.MsgStatusInFlight
	case MsgStatus_MSG_STATUS_DELIVERED:
		return layerzero.MsgStatusDelivered
	case MsgStatus_MSG_STATUS_FAILED:
		return layerzero.MsgStatusFailed
	case MsgStatus_MSG_STATUS_PAYLOAD_STORED:
		return layerzero.MsgStatusPayloadStored
	default:
		return layerzero.MsgStatusUnknown
	}
}

// msgToProto converts a MsgSend to a MsgSendUSDT0 proto.
func msgToProto(msg MsgSend) *MsgSendUSDT0 {
	return &MsgSendUSDT0{
		TxHash:      msg.TxHash[:],
		BlockHeight: msg.BlockHeight,
		SrcChainId:  msg.SrcChainID,
		DestChainId: msg.DestChainID,
		Amount:      msg.Amount.Bytes(),
		Status:      int32(statusToProto(msg.Status)),
	}
}

// msgFromProto converts a MsgSendUSDT0 proto to a MsgSend.
func msgFromProto(msg *MsgSendUSDT0) (MsgSend, error) {
	txHash, err := cast.EthHash(msg.GetTxHash())
	if err != nil {
		return MsgSend{}, errors.Wrap(err, "cast tx hash")
	}

	return MsgSend{
		TxHash:      txHash,
		BlockHeight: msg.GetBlockHeight(),
		SrcChainID:  msg.GetSrcChainId(),
		DestChainID: msg.GetDestChainId(),
		Amount:      new(big.Int).SetBytes(msg.GetAmount()),
		Status:      statusFromProto(MsgStatus(msg.GetStatus())),
	}, nil
}

// FilterMsgByDest returns a filter that matches messages with the given destination chain ID.
func FilterMsgByDest(chainID uint64) MsgFilter {
	return func(msg MsgSend) bool {
		return msg.DestChainID == chainID
	}
}

// FilterMsgBySrc returns a filter that matches messages with the given source chain ID.
func FilterMsgBySrc(chainID uint64) MsgFilter {
	return func(msg MsgSend) bool {
		return msg.SrcChainID == chainID
	}
}

// FilterMsgByStatus returns a filter that matches messages with any of the given statuses.
func FilterMsgByStatus(statuses ...layerzero.MsgStatus) MsgFilter {
	set := make(map[layerzero.MsgStatus]bool, len(statuses))
	for _, s := range statuses {
		set[s] = true
	}

	return func(msg MsgSend) bool {
		_, ok := set[msg.Status]
		return ok
	}
}

// GetMsgCreatedAt retrieves the created at timestamp of a message by tx hash.
func (db *DB) GetMsgCreatedAt(ctx context.Context, txHash common.Hash) (time.Time, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	msg, ok, err := db.getMsgUnsafe(ctx, txHash)
	if err != nil {
		return time.Time{}, errors.Wrap(err, "get msg")
	} else if !ok {
		return time.Time{}, errors.New("msg not found")
	}

	return msg.GetCreatedAt().AsTime(), nil
}
