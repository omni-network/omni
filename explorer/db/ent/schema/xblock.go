package schema

import (
	"entgo.io/ent"
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
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New),
	}
}

// Edges of the XBlock.
func (XBlock) Edges() []ent.Edge {
	return []ent.Edge{}
}
