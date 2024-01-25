package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// XBlock holds the schema definition for the XBlock entity.
type Chain struct {
	ent.Schema
}

// Fields of the XBlock.
func (Chain) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("UUID", uuid.UUID{}).
			Default(uuid.New),
		field.Time("CreatedAt").
			Default(time.Now()),
		field.Uint64("ChainId"),
		field.String("Name"),
	}
}

// Edges of the XBlock.
func (Chain) Edges() []ent.Edge {
	return []ent.Edge{}
}
