package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// XReceipt holds the schema definition for the XReceipt entity.
type XReceipt struct {
	ent.Schema
}

// Fields of the XBlock.
func (XReceipt) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("UUID", uuid.UUID{}).
			Default(uuid.New),
		field.Uint64("GasUsed"),
		field.Bool("Success"),
		field.Bytes("RelayerAddress").
			MaxLen(20),
		field.String("XMsgUUID").
			NotEmpty(),
		field.Uint64("SourceChainID"),
		field.Uint64("DestChainID"),
		field.Uint64("StreamOffset"),
		field.Bytes("TxHash").
			MaxLen(32),
		field.Time("CreatedAt").
			Default(time.Now()),
	}
}

// Edges of the XReceipt.
func (XReceipt) Edges() []ent.Edge {
	return []ent.Edge{}
}
