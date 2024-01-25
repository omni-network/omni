package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Chain holds the schema definition for the Chain entity.
type Chain struct {
	ent.Schema
}

// Fields of the Chain.
func (Chain) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("UUID", uuid.UUID{}).
			Default(uuid.New),
		field.Time("CreatedAt").
			Default(time.Now()),
		field.Uint64("ChainID"),
		field.String("Name"),
	}
}

// Edges of the Chain.
func (Chain) Edges() []ent.Edge {
	return []ent.Edge{}
}
