package schema

import (
	"time"

	"entgo.io/ent/schema/index"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// XProviderCursor holds the schema definition for the XProviderCursor entity.
type XProviderCursor struct {
	ent.Schema
}

// Fields of the XProviderCursor.
func (XProviderCursor) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Uint64("chain_id").Unique(),
		field.Uint64("height"),
		field.Uint64("offset"),
		field.Time("created_at").Default(time.Now()),
		field.Time("updated_at").Default(time.Now()),
	}
}

func (XProviderCursor) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("chain_id").Unique(), // Unique cursor per ChainID.
	}
}

// Edges of the XProviderCursor.
func (XProviderCursor) Edges() []ent.Edge {
	return []ent.Edge{}
}
