// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/omni-network/omni/explorer/db/ent/xprovidercursor"
)

// XProviderCursorCreate is the builder for creating a XProviderCursor entity.
type XProviderCursorCreate struct {
	config
	mutation *XProviderCursorMutation
	hooks    []Hook
}

// SetUUID sets the "UUID" field.
func (xcc *XProviderCursorCreate) SetUUID(u uuid.UUID) *XProviderCursorCreate {
	xcc.mutation.SetUUID(u)
	return xcc
}

// SetNillableUUID sets the "UUID" field if the given value is not nil.
func (xcc *XProviderCursorCreate) SetNillableUUID(u *uuid.UUID) *XProviderCursorCreate {
	if u != nil {
		xcc.SetUUID(*u)
	}
	return xcc
}

// SetChainID sets the "ChainID" field.
func (xcc *XProviderCursorCreate) SetChainID(u uint64) *XProviderCursorCreate {
	xcc.mutation.SetChainID(u)
	return xcc
}

// SetHeight sets the "Height" field.
func (xcc *XProviderCursorCreate) SetHeight(u uint64) *XProviderCursorCreate {
	xcc.mutation.SetHeight(u)
	return xcc
}

// SetCreatedAt sets the "CreatedAt" field.
func (xcc *XProviderCursorCreate) SetCreatedAt(t time.Time) *XProviderCursorCreate {
	xcc.mutation.SetCreatedAt(t)
	return xcc
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (xcc *XProviderCursorCreate) SetNillableCreatedAt(t *time.Time) *XProviderCursorCreate {
	if t != nil {
		xcc.SetCreatedAt(*t)
	}
	return xcc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (xcc *XProviderCursorCreate) SetUpdatedAt(t time.Time) *XProviderCursorCreate {
	xcc.mutation.SetUpdatedAt(t)
	return xcc
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (xcc *XProviderCursorCreate) SetNillableUpdatedAt(t *time.Time) *XProviderCursorCreate {
	if t != nil {
		xcc.SetUpdatedAt(*t)
	}
	return xcc
}

// Mutation returns the XProviderCursorMutation object of the builder.
func (xcc *XProviderCursorCreate) Mutation() *XProviderCursorMutation {
	return xcc.mutation
}

// Save creates the XProviderCursor in the database.
func (xcc *XProviderCursorCreate) Save(ctx context.Context) (*XProviderCursor, error) {
	xcc.defaults()
	return withHooks(ctx, xcc.sqlSave, xcc.mutation, xcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (xcc *XProviderCursorCreate) SaveX(ctx context.Context) *XProviderCursor {
	v, err := xcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (xcc *XProviderCursorCreate) Exec(ctx context.Context) error {
	_, err := xcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (xcc *XProviderCursorCreate) ExecX(ctx context.Context) {
	if err := xcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (xcc *XProviderCursorCreate) defaults() {
	if _, ok := xcc.mutation.UUID(); !ok {
		v := xprovidercursor.DefaultUUID()
		xcc.mutation.SetUUID(v)
	}
	if _, ok := xcc.mutation.CreatedAt(); !ok {
		v := xprovidercursor.DefaultCreatedAt
		xcc.mutation.SetCreatedAt(v)
	}
	if _, ok := xcc.mutation.UpdatedAt(); !ok {
		v := xprovidercursor.DefaultUpdatedAt
		xcc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (xcc *XProviderCursorCreate) check() error {
	if _, ok := xcc.mutation.UUID(); !ok {
		return &ValidationError{Name: "UUID", err: errors.New(`ent: missing required field "XProviderCursor.UUID"`)}
	}
	if _, ok := xcc.mutation.ChainID(); !ok {
		return &ValidationError{Name: "ChainID", err: errors.New(`ent: missing required field "XProviderCursor.ChainID"`)}
	}
	if _, ok := xcc.mutation.Height(); !ok {
		return &ValidationError{Name: "Height", err: errors.New(`ent: missing required field "XProviderCursor.Height"`)}
	}
	if _, ok := xcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`ent: missing required field "XProviderCursor.CreatedAt"`)}
	}
	if _, ok := xcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`ent: missing required field "XProviderCursor.UpdatedAt"`)}
	}
	return nil
}

func (xcc *XProviderCursorCreate) sqlSave(ctx context.Context) (*XProviderCursor, error) {
	if err := xcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := xcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, xcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	xcc.mutation.id = &_node.ID
	xcc.mutation.done = true
	return _node, nil
}

func (xcc *XProviderCursorCreate) createSpec() (*XProviderCursor, *sqlgraph.CreateSpec) {
	var (
		_node = &XProviderCursor{config: xcc.config}
		_spec = sqlgraph.NewCreateSpec(xprovidercursor.Table, sqlgraph.NewFieldSpec(xprovidercursor.FieldID, field.TypeInt))
	)
	if value, ok := xcc.mutation.UUID(); ok {
		_spec.SetField(xprovidercursor.FieldUUID, field.TypeUUID, value)
		_node.UUID = value
	}
	if value, ok := xcc.mutation.ChainID(); ok {
		_spec.SetField(xprovidercursor.FieldChainID, field.TypeUint64, value)
		_node.ChainID = value
	}
	if value, ok := xcc.mutation.Height(); ok {
		_spec.SetField(xprovidercursor.FieldHeight, field.TypeUint64, value)
		_node.Height = value
	}
	if value, ok := xcc.mutation.CreatedAt(); ok {
		_spec.SetField(xprovidercursor.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := xcc.mutation.UpdatedAt(); ok {
		_spec.SetField(xprovidercursor.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// XProviderCursorCreateBulk is the builder for creating many XProviderCursor entities in bulk.
type XProviderCursorCreateBulk struct {
	config
	err      error
	builders []*XProviderCursorCreate
}

// Save creates the XProviderCursor entities in the database.
func (xccb *XProviderCursorCreateBulk) Save(ctx context.Context) ([]*XProviderCursor, error) {
	if xccb.err != nil {
		return nil, xccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(xccb.builders))
	nodes := make([]*XProviderCursor, len(xccb.builders))
	mutators := make([]Mutator, len(xccb.builders))
	for i := range xccb.builders {
		func(i int, root context.Context) {
			builder := xccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*XProviderCursorMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, xccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, xccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, xccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (xccb *XProviderCursorCreateBulk) SaveX(ctx context.Context) []*XProviderCursor {
	v, err := xccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (xccb *XProviderCursorCreateBulk) Exec(ctx context.Context) error {
	_, err := xccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (xccb *XProviderCursorCreateBulk) ExecX(ctx context.Context) {
	if err := xccb.Exec(ctx); err != nil {
		panic(err)
	}
}
