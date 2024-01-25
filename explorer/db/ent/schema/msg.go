package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Msg holds the schema definition for the Msg entity.
type Msg struct {
	ent.Schema
}

// Fields of the XMsg.
func (Msg) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("UUID", uuid.UUID{}).
			Default(uuid.New),
		field.Bytes("SourceMsgSender").
			MaxLen(20),
		field.Bytes("DestAddress").
			MaxLen(20),
		field.Bytes("Data"),
		field.Uint64("DestGasLimit"),
		field.Uint64("SourceChainID"),
		field.Uint64("DestChainID"),
		field.Uint64("StreamOffset"),
		field.Bytes("TxHash").
			MaxLen(32),
		field.Time("CreatedAt").
			Default(time.Now()),
	}
}

// Edges of the XMsg.
func (Msg) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Block", Block.Type).
			Ref("Msgs").
			Unique(),
	}
}
