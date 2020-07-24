package schema

import "github.com/facebookincubator/ent"

// Resource holds the schema definition for the Resource entity.
type Resource struct {
	ent.Schema
}

// Fields of the Resource.
func (Resource) Fields() []ent.Field {
	return nil
}

// Edges of the Resource.
func (Resource) Edges() []ent.Edge {
	return nil
}
