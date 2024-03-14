// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Geometry struct {
	ent.Schema
}

func (Geometry) Fields() []ent.Field {
	return []ent.Field{field.Int32("id"), field.Int32("earthquake_id").Optional(), field.Float("longitude"), field.Float("latitude"), field.Float("depth"), field.String("place")}
}
func (Geometry) Edges() []ent.Edge {
	return []ent.Edge{edge.From("earthquake", Earthquake.Type).Ref("geometries").Unique().Field("earthquake_id")}
}
func (Geometry) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "Geometry"}}
}
