// Code generated by ent, DO NOT EDIT.

package schemamigration

import (
	"entdemo/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldLTE(FieldID, id))
}

// Dirty applies equality check predicate on the "dirty" field. It's identical to DirtyEQ.
func Dirty(v bool) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldEQ(FieldDirty, v))
}

// DirtyEQ applies the EQ predicate on the "dirty" field.
func DirtyEQ(v bool) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldEQ(FieldDirty, v))
}

// DirtyNEQ applies the NEQ predicate on the "dirty" field.
func DirtyNEQ(v bool) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.FieldNEQ(FieldDirty, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SchemaMigration) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SchemaMigration) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SchemaMigration) predicate.SchemaMigration {
	return predicate.SchemaMigration(sql.NotPredicates(p))
}
