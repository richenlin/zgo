package schema

/*
  demo
  ID该字段内置于架构中，不需要声明。
  在基于 SQL 的数据库中，其类型默认为数据库中自动递增
*/
import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Demo holds the schema definition for the Demo entity.
type Demo struct {
	ent.Schema
}

// Fields of the Demo.
func (Demo) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").Unique(),
		field.String("name").Unique(),
		field.String("demo"),
		field.Int("status").Min(1).Max(2),
		field.String("creator"),
		field.String("updator"),
		field.Time("created_at"),
		field.Time("upadted_at"),
	}
}

// Edges of the Demo.
func (Demo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Demo.Type).
			From("parent").
			Unique(),
		// 以下内容效果相同
		// edge.To("children", Demo.Type),
		// edge.From("parent", Demo.Type).
		// 	Ref("children").
		// 	Unique(),
	}
}
