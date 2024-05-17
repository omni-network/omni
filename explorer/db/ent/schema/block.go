package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Block holds the schema definition for the Block entity.
type Block struct {
	ent.Schema
}

// Fields of the XBlock.
func (Block) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("hash").MaxLen(32),
		field.Uint64("chain_id"),
		field.Uint64("height"),
		field.Uint64("offset"),
		field.Time("timestamp").Default(time.Now()),
		field.Time("created_at").Default(time.Now()),
	}
}

// Indexes of the Block.
func (Block) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("chain_id", "hash").Unique(),
		index.Fields("chain_id", "height").Unique(),
	}
}

// Edges of the Block.
func (Block) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("msgs", Msg.Type),
		edge.To("receipts", Receipt.Type),
	}
}
