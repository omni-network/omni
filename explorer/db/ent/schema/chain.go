package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Chain holds the schema definition for the Chain entity.
type Chain struct {
	ent.Schema
}

// Fields of the Chain.
func (Chain) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("chain_id").Unique(),
		field.Time("created_at").Default(time.Now()),
		field.String("name"),
	}
}

// Edges of the Chain.
func (Chain) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Chain) Indexes() []ent.Index {
	return []ent.Index{}
}
