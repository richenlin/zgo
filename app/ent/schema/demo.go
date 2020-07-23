package schema

/*
  demo
*/
import "github.com/facebookincubator/ent"

// Demo holds the schema definition for the Demo entity.
type Demo struct {
	ent.Schema
}

// Fields of the Demo.
func (Demo) Fields() []ent.Field {
	return nil
}

// Edges of the Demo.
func (Demo) Edges() []ent.Edge {
	return nil
}
