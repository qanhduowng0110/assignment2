// Code generated by ent, DO NOT EDIT.

package eventtype

import (
	"entdemo/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.EventType {
	return predicate.EventType(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.EventType {
	return predicate.EventType(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.EventType {
	return predicate.EventType(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.EventType {
	return predicate.EventType(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.EventType {
	return predicate.EventType(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.EventType {
	return predicate.EventType(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.EventType {
	return predicate.EventType(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.EventType {
	return predicate.EventType(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.EventType {
	return predicate.EventType(sql.FieldLTE(FieldID, id))
}

// EarthquakeID applies equality check predicate on the "earthquake_id" field. It's identical to EarthquakeIDEQ.
func EarthquakeID(v int32) predicate.EventType {
	return predicate.EventType(sql.FieldEQ(FieldEarthquakeID, v))
}

// Types applies equality check predicate on the "types" field. It's identical to TypesEQ.
func Types(v string) predicate.EventType {
	return predicate.EventType(sql.FieldEQ(FieldTypes, v))
}

// EarthquakeIDEQ applies the EQ predicate on the "earthquake_id" field.
func EarthquakeIDEQ(v int32) predicate.EventType {
	return predicate.EventType(sql.FieldEQ(FieldEarthquakeID, v))
}

// EarthquakeIDNEQ applies the NEQ predicate on the "earthquake_id" field.
func EarthquakeIDNEQ(v int32) predicate.EventType {
	return predicate.EventType(sql.FieldNEQ(FieldEarthquakeID, v))
}

// EarthquakeIDIn applies the In predicate on the "earthquake_id" field.
func EarthquakeIDIn(vs ...int32) predicate.EventType {
	return predicate.EventType(sql.FieldIn(FieldEarthquakeID, vs...))
}

// EarthquakeIDNotIn applies the NotIn predicate on the "earthquake_id" field.
func EarthquakeIDNotIn(vs ...int32) predicate.EventType {
	return predicate.EventType(sql.FieldNotIn(FieldEarthquakeID, vs...))
}

// EarthquakeIDIsNil applies the IsNil predicate on the "earthquake_id" field.
func EarthquakeIDIsNil() predicate.EventType {
	return predicate.EventType(sql.FieldIsNull(FieldEarthquakeID))
}

// EarthquakeIDNotNil applies the NotNil predicate on the "earthquake_id" field.
func EarthquakeIDNotNil() predicate.EventType {
	return predicate.EventType(sql.FieldNotNull(FieldEarthquakeID))
}

// TypesEQ applies the EQ predicate on the "types" field.
func TypesEQ(v string) predicate.EventType {
	return predicate.EventType(sql.FieldEQ(FieldTypes, v))
}

// TypesNEQ applies the NEQ predicate on the "types" field.
func TypesNEQ(v string) predicate.EventType {
	return predicate.EventType(sql.FieldNEQ(FieldTypes, v))
}

// TypesIn applies the In predicate on the "types" field.
func TypesIn(vs ...string) predicate.EventType {
	return predicate.EventType(sql.FieldIn(FieldTypes, vs...))
}

// TypesNotIn applies the NotIn predicate on the "types" field.
func TypesNotIn(vs ...string) predicate.EventType {
	return predicate.EventType(sql.FieldNotIn(FieldTypes, vs...))
}

// TypesGT applies the GT predicate on the "types" field.
func TypesGT(v string) predicate.EventType {
	return predicate.EventType(sql.FieldGT(FieldTypes, v))
}

// TypesGTE applies the GTE predicate on the "types" field.
func TypesGTE(v string) predicate.EventType {
	return predicate.EventType(sql.FieldGTE(FieldTypes, v))
}

// TypesLT applies the LT predicate on the "types" field.
func TypesLT(v string) predicate.EventType {
	return predicate.EventType(sql.FieldLT(FieldTypes, v))
}

// TypesLTE applies the LTE predicate on the "types" field.
func TypesLTE(v string) predicate.EventType {
	return predicate.EventType(sql.FieldLTE(FieldTypes, v))
}

// TypesContains applies the Contains predicate on the "types" field.
func TypesContains(v string) predicate.EventType {
	return predicate.EventType(sql.FieldContains(FieldTypes, v))
}

// TypesHasPrefix applies the HasPrefix predicate on the "types" field.
func TypesHasPrefix(v string) predicate.EventType {
	return predicate.EventType(sql.FieldHasPrefix(FieldTypes, v))
}

// TypesHasSuffix applies the HasSuffix predicate on the "types" field.
func TypesHasSuffix(v string) predicate.EventType {
	return predicate.EventType(sql.FieldHasSuffix(FieldTypes, v))
}

// TypesIsNil applies the IsNil predicate on the "types" field.
func TypesIsNil() predicate.EventType {
	return predicate.EventType(sql.FieldIsNull(FieldTypes))
}

// TypesNotNil applies the NotNil predicate on the "types" field.
func TypesNotNil() predicate.EventType {
	return predicate.EventType(sql.FieldNotNull(FieldTypes))
}

// TypesEqualFold applies the EqualFold predicate on the "types" field.
func TypesEqualFold(v string) predicate.EventType {
	return predicate.EventType(sql.FieldEqualFold(FieldTypes, v))
}

// TypesContainsFold applies the ContainsFold predicate on the "types" field.
func TypesContainsFold(v string) predicate.EventType {
	return predicate.EventType(sql.FieldContainsFold(FieldTypes, v))
}

// HasEarthquake applies the HasEdge predicate on the "earthquake" edge.
func HasEarthquake() predicate.EventType {
	return predicate.EventType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EarthquakeTable, EarthquakeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEarthquakeWith applies the HasEdge predicate on the "earthquake" edge with a given conditions (other predicates).
func HasEarthquakeWith(preds ...predicate.Earthquake) predicate.EventType {
	return predicate.EventType(func(s *sql.Selector) {
		step := newEarthquakeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EventType) predicate.EventType {
	return predicate.EventType(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EventType) predicate.EventType {
	return predicate.EventType(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EventType) predicate.EventType {
	return predicate.EventType(sql.NotPredicates(p))
}
