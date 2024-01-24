package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// XBlock holds the schema definition for the XBlock entity.
type XBlock struct {
	ent.Schema
}

// Fields of the XBlock.
func (XBlock) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("UUID", uuid.UUID{}).
			Default(uuid.New),
		field.Uint64("SourceChainID"),
		field.Uint64("BlockHeight"),
		field.Bytes("BlockHash").MaxLen(32),
		field.Time("CreatedAt").
			Default(time.Now()),
	}
}

// Edges of the XBlock.
func (XBlock) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Msgs", XMsg.Type),
		edge.To("Receipts", XReceipt.Type),
	}
}
