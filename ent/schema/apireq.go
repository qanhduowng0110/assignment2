// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Apireq struct {
	ent.Schema
}

func (Apireq) Fields() []ent.Field {
	return []ent.Field{field.Int("id"), 
	field.Time("req_time"), 
	field.JSON("req_param", struct{}{}), 
	field.JSON("req_body", struct{}{}).Optional(), 
	field.JSON("req_headers", struct{}{}).Optional(), 
	field.JSON("req_metadata", struct{}{}).Optional(), 
	field.Time("created_at"), 
	field.Time("updated_at"), 
	field.Time("deleted_at").Optional()}

}
func (Apireq) Edges() []ent.Edge {
	return nil
}
func (Apireq) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "apireq"}}
}
