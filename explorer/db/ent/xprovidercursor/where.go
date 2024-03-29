// Code generated by ent, DO NOT EDIT.

package xprovidercursor

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/omni-network/omni/explorer/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "UUID" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldEQ(FieldUUID, v))
}

// ChainID applies equality check predicate on the "ChainID" field. It's identical to ChainIDEQ.
func ChainID(v uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldEQ(FieldChainID, v))
}

// Height applies equality check predicate on the "Height" field. It's identical to HeightEQ.
func Height(v uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldEQ(FieldHeight, v))
}

// CreatedAt applies equality check predicate on the "CreatedAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "UpdatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldEQ(FieldUpdatedAt, v))
}

// UUIDEQ applies the EQ predicate on the "UUID" field.
func UUIDEQ(v uuid.UUID) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "UUID" field.
func UUIDNEQ(v uuid.UUID) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "UUID" field.
func UUIDIn(vs ...uuid.UUID) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "UUID" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "UUID" field.
func UUIDGT(v uuid.UUID) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "UUID" field.
func UUIDGTE(v uuid.UUID) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "UUID" field.
func UUIDLT(v uuid.UUID) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "UUID" field.
func UUIDLTE(v uuid.UUID) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldLTE(FieldUUID, v))
}

// ChainIDEQ applies the EQ predicate on the "ChainID" field.
func ChainIDEQ(v uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldEQ(FieldChainID, v))
}

// ChainIDNEQ applies the NEQ predicate on the "ChainID" field.
func ChainIDNEQ(v uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldNEQ(FieldChainID, v))
}

// ChainIDIn applies the In predicate on the "ChainID" field.
func ChainIDIn(vs ...uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldIn(FieldChainID, vs...))
}

// ChainIDNotIn applies the NotIn predicate on the "ChainID" field.
func ChainIDNotIn(vs ...uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldNotIn(FieldChainID, vs...))
}

// ChainIDGT applies the GT predicate on the "ChainID" field.
func ChainIDGT(v uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldGT(FieldChainID, v))
}

// ChainIDGTE applies the GTE predicate on the "ChainID" field.
func ChainIDGTE(v uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldGTE(FieldChainID, v))
}

// ChainIDLT applies the LT predicate on the "ChainID" field.
func ChainIDLT(v uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldLT(FieldChainID, v))
}

// ChainIDLTE applies the LTE predicate on the "ChainID" field.
func ChainIDLTE(v uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldLTE(FieldChainID, v))
}

// HeightEQ applies the EQ predicate on the "Height" field.
func HeightEQ(v uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldEQ(FieldHeight, v))
}

// HeightNEQ applies the NEQ predicate on the "Height" field.
func HeightNEQ(v uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldNEQ(FieldHeight, v))
}

// HeightIn applies the In predicate on the "Height" field.
func HeightIn(vs ...uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldIn(FieldHeight, vs...))
}

// HeightNotIn applies the NotIn predicate on the "Height" field.
func HeightNotIn(vs ...uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldNotIn(FieldHeight, vs...))
}

// HeightGT applies the GT predicate on the "Height" field.
func HeightGT(v uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldGT(FieldHeight, v))
}

// HeightGTE applies the GTE predicate on the "Height" field.
func HeightGTE(v uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldGTE(FieldHeight, v))
}

// HeightLT applies the LT predicate on the "Height" field.
func HeightLT(v uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldLT(FieldHeight, v))
}

// HeightLTE applies the LTE predicate on the "Height" field.
func HeightLTE(v uint64) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldLTE(FieldHeight, v))
}

// CreatedAtEQ applies the EQ predicate on the "CreatedAt" field.
func CreatedAtEQ(v time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "CreatedAt" field.
func CreatedAtNEQ(v time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "CreatedAt" field.
func CreatedAtIn(vs ...time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "CreatedAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "CreatedAt" field.
func CreatedAtGT(v time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "CreatedAt" field.
func CreatedAtGTE(v time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "CreatedAt" field.
func CreatedAtLT(v time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "CreatedAt" field.
func CreatedAtLTE(v time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "UpdatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "UpdatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "UpdatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "UpdatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "UpdatedAt" field.
func UpdatedAtGT(v time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "UpdatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "UpdatedAt" field.
func UpdatedAtLT(v time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "UpdatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.XProviderCursor) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.XProviderCursor) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.XProviderCursor) predicate.XProviderCursor {
	return predicate.XProviderCursor(sql.NotPredicates(p))
}
