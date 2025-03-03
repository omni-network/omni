// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package app

import (
	context "context"
	ormlist "cosmossdk.io/orm/model/ormlist"
	ormtable "cosmossdk.io/orm/model/ormtable"
	ormerrors "cosmossdk.io/orm/types/ormerrors"
)

type CursorTable interface {
	Insert(ctx context.Context, cursor *Cursor) error
	Update(ctx context.Context, cursor *Cursor) error
	Save(ctx context.Context, cursor *Cursor) error
	Delete(ctx context.Context, cursor *Cursor) error
	Has(ctx context.Context, chain_id uint64, conf_level uint32) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, chain_id uint64, conf_level uint32) (*Cursor, error)
	List(ctx context.Context, prefixKey CursorIndexKey, opts ...ormlist.Option) (CursorIterator, error)
	ListRange(ctx context.Context, from, to CursorIndexKey, opts ...ormlist.Option) (CursorIterator, error)
	DeleteBy(ctx context.Context, prefixKey CursorIndexKey) error
	DeleteRange(ctx context.Context, from, to CursorIndexKey) error

	doNotImplement()
}

type CursorIterator struct {
	ormtable.Iterator
}

func (i CursorIterator) Value() (*Cursor, error) {
	var cursor Cursor
	err := i.UnmarshalMessage(&cursor)
	return &cursor, err
}

type CursorIndexKey interface {
	id() uint32
	values() []interface{}
	cursorIndexKey()
}

// primary key starting index..
type CursorPrimaryKey = CursorChainIdConfLevelIndexKey

type CursorChainIdConfLevelIndexKey struct {
	vs []interface{}
}

func (x CursorChainIdConfLevelIndexKey) id() uint32            { return 0 }
func (x CursorChainIdConfLevelIndexKey) values() []interface{} { return x.vs }
func (x CursorChainIdConfLevelIndexKey) cursorIndexKey()       {}

func (this CursorChainIdConfLevelIndexKey) WithChainId(chain_id uint64) CursorChainIdConfLevelIndexKey {
	this.vs = []interface{}{chain_id}
	return this
}

func (this CursorChainIdConfLevelIndexKey) WithChainIdConfLevel(chain_id uint64, conf_level uint32) CursorChainIdConfLevelIndexKey {
	this.vs = []interface{}{chain_id, conf_level}
	return this
}

type cursorTable struct {
	table ormtable.Table
}

func (this cursorTable) Insert(ctx context.Context, cursor *Cursor) error {
	return this.table.Insert(ctx, cursor)
}

func (this cursorTable) Update(ctx context.Context, cursor *Cursor) error {
	return this.table.Update(ctx, cursor)
}

func (this cursorTable) Save(ctx context.Context, cursor *Cursor) error {
	return this.table.Save(ctx, cursor)
}

func (this cursorTable) Delete(ctx context.Context, cursor *Cursor) error {
	return this.table.Delete(ctx, cursor)
}

func (this cursorTable) Has(ctx context.Context, chain_id uint64, conf_level uint32) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, chain_id, conf_level)
}

func (this cursorTable) Get(ctx context.Context, chain_id uint64, conf_level uint32) (*Cursor, error) {
	var cursor Cursor
	found, err := this.table.PrimaryKey().Get(ctx, &cursor, chain_id, conf_level)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &cursor, nil
}

func (this cursorTable) List(ctx context.Context, prefixKey CursorIndexKey, opts ...ormlist.Option) (CursorIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return CursorIterator{it}, err
}

func (this cursorTable) ListRange(ctx context.Context, from, to CursorIndexKey, opts ...ormlist.Option) (CursorIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return CursorIterator{it}, err
}

func (this cursorTable) DeleteBy(ctx context.Context, prefixKey CursorIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this cursorTable) DeleteRange(ctx context.Context, from, to CursorIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this cursorTable) doNotImplement() {}

var _ CursorTable = cursorTable{}

func NewCursorTable(db ormtable.Schema) (CursorTable, error) {
	table := db.GetTable(&Cursor{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Cursor{}).ProtoReflect().Descriptor().FullName()))
	}
	return cursorTable{table}, nil
}

type SolverStore interface {
	CursorTable() CursorTable

	doNotImplement()
}

type solverStore struct {
	cursor CursorTable
}

func (x solverStore) CursorTable() CursorTable {
	return x.cursor
}

func (solverStore) doNotImplement() {}

var _ SolverStore = solverStore{}

func NewSolverStore(db ormtable.Schema) (SolverStore, error) {
	cursorTable, err := NewCursorTable(db)
	if err != nil {
		return nil, err
	}

	return solverStore{
		cursorTable,
	}, nil
}
