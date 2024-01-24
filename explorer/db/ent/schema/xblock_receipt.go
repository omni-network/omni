package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// XBlockReceipt holds the schema definition for the XBlockReceipt entity.
type XBlockReceipt struct {
	ent.Schema
}

// Fields of the XBlockReceipt.
func (XBlockReceipt) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("UUID", uuid.UUID{}).
			Default(uuid.New),
		field.String("XBlockUUID"),
		field.String("XReceiptUUID"),
		field.Time("CreatedAt").
			Default(time.Now()),
	}
}

// Edges of the XBlockReceipt.
func (XBlockReceipt) Edges() []ent.Edge {
	return []ent.Edge{}
}
