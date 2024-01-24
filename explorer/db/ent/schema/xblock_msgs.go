package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// XBlockMsgs holds the schema definition for the XBlockMsgs entity.
type XBlockMsgs struct {
	ent.Schema
}

// Fields of the XBlockMsgs.
func (XBlockMsgs) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("UUID", uuid.UUID{}).
			Default(uuid.New),
		field.String("XBlockUUID").
			NotEmpty(),
		field.String("XMsgUUID").
			NotEmpty(),
		field.Bytes("SourceMsgSender").
			MaxLen(20),
		field.Bytes("DestAddress").
			MaxLen(20),
		field.Bytes("Data"),
		field.Uint64("DestGasLimit"),
		field.Bytes("TxHash").
			MaxLen(32),
		field.Time("CreatedAt").
			Default(time.Now()),
	}
}

// Edges of the XBlockMsgs.
func (XBlockMsgs) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("XBlock", XBlock.Type).
			Ref("Msgs").
			Unique(),
	}
}

// XBlockMsg represents a immutable cross-chain message in a block (one to many relationship).
