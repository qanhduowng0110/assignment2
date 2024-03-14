// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/earthquake"
	"entdemo/ent/eventtype"
	"entdemo/ent/predicate"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EventTypeUpdate is the builder for updating EventType entities.
type EventTypeUpdate struct {
	config
	hooks    []Hook
	mutation *EventTypeMutation
}

// Where appends a list predicates to the EventTypeUpdate builder.
func (etu *EventTypeUpdate) Where(ps ...predicate.EventType) *EventTypeUpdate {
	etu.mutation.Where(ps...)
	return etu
}

// SetEarthquakeID sets the "earthquake_id" field.
func (etu *EventTypeUpdate) SetEarthquakeID(i int32) *EventTypeUpdate {
	etu.mutation.SetEarthquakeID(i)
	return etu
}

// SetNillableEarthquakeID sets the "earthquake_id" field if the given value is not nil.
func (etu *EventTypeUpdate) SetNillableEarthquakeID(i *int32) *EventTypeUpdate {
	if i != nil {
		etu.SetEarthquakeID(*i)
	}
	return etu
}

// ClearEarthquakeID clears the value of the "earthquake_id" field.
func (etu *EventTypeUpdate) ClearEarthquakeID() *EventTypeUpdate {
	etu.mutation.ClearEarthquakeID()
	return etu
}

// SetTypes sets the "types" field.
func (etu *EventTypeUpdate) SetTypes(s string) *EventTypeUpdate {
	etu.mutation.SetTypes(s)
	return etu
}

// SetNillableTypes sets the "types" field if the given value is not nil.
func (etu *EventTypeUpdate) SetNillableTypes(s *string) *EventTypeUpdate {
	if s != nil {
		etu.SetTypes(*s)
	}
	return etu
}

// ClearTypes clears the value of the "types" field.
func (etu *EventTypeUpdate) ClearTypes() *EventTypeUpdate {
	etu.mutation.ClearTypes()
	return etu
}

// SetEarthquake sets the "earthquake" edge to the Earthquake entity.
func (etu *EventTypeUpdate) SetEarthquake(e *Earthquake) *EventTypeUpdate {
	return etu.SetEarthquakeID(e.ID)
}

// Mutation returns the EventTypeMutation object of the builder.
func (etu *EventTypeUpdate) Mutation() *EventTypeMutation {
	return etu.mutation
}

// ClearEarthquake clears the "earthquake" edge to the Earthquake entity.
func (etu *EventTypeUpdate) ClearEarthquake() *EventTypeUpdate {
	etu.mutation.ClearEarthquake()
	return etu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (etu *EventTypeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, etu.sqlSave, etu.mutation, etu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (etu *EventTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := etu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (etu *EventTypeUpdate) Exec(ctx context.Context) error {
	_, err := etu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etu *EventTypeUpdate) ExecX(ctx context.Context) {
	if err := etu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (etu *EventTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(eventtype.Table, eventtype.Columns, sqlgraph.NewFieldSpec(eventtype.FieldID, field.TypeInt32))
	if ps := etu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := etu.mutation.Types(); ok {
		_spec.SetField(eventtype.FieldTypes, field.TypeString, value)
	}
	if etu.mutation.TypesCleared() {
		_spec.ClearField(eventtype.FieldTypes, field.TypeString)
	}
	if etu.mutation.EarthquakeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventtype.EarthquakeTable,
			Columns: []string{eventtype.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := etu.mutation.EarthquakeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventtype.EarthquakeTable,
			Columns: []string{eventtype.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, etu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	etu.mutation.done = true
	return n, nil
}

// EventTypeUpdateOne is the builder for updating a single EventType entity.
type EventTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventTypeMutation
}

// SetEarthquakeID sets the "earthquake_id" field.
func (etuo *EventTypeUpdateOne) SetEarthquakeID(i int32) *EventTypeUpdateOne {
	etuo.mutation.SetEarthquakeID(i)
	return etuo
}

// SetNillableEarthquakeID sets the "earthquake_id" field if the given value is not nil.
func (etuo *EventTypeUpdateOne) SetNillableEarthquakeID(i *int32) *EventTypeUpdateOne {
	if i != nil {
		etuo.SetEarthquakeID(*i)
	}
	return etuo
}

// ClearEarthquakeID clears the value of the "earthquake_id" field.
func (etuo *EventTypeUpdateOne) ClearEarthquakeID() *EventTypeUpdateOne {
	etuo.mutation.ClearEarthquakeID()
	return etuo
}

// SetTypes sets the "types" field.
func (etuo *EventTypeUpdateOne) SetTypes(s string) *EventTypeUpdateOne {
	etuo.mutation.SetTypes(s)
	return etuo
}

// SetNillableTypes sets the "types" field if the given value is not nil.
func (etuo *EventTypeUpdateOne) SetNillableTypes(s *string) *EventTypeUpdateOne {
	if s != nil {
		etuo.SetTypes(*s)
	}
	return etuo
}

// ClearTypes clears the value of the "types" field.
func (etuo *EventTypeUpdateOne) ClearTypes() *EventTypeUpdateOne {
	etuo.mutation.ClearTypes()
	return etuo
}

// SetEarthquake sets the "earthquake" edge to the Earthquake entity.
func (etuo *EventTypeUpdateOne) SetEarthquake(e *Earthquake) *EventTypeUpdateOne {
	return etuo.SetEarthquakeID(e.ID)
}

// Mutation returns the EventTypeMutation object of the builder.
func (etuo *EventTypeUpdateOne) Mutation() *EventTypeMutation {
	return etuo.mutation
}

// ClearEarthquake clears the "earthquake" edge to the Earthquake entity.
func (etuo *EventTypeUpdateOne) ClearEarthquake() *EventTypeUpdateOne {
	etuo.mutation.ClearEarthquake()
	return etuo
}

// Where appends a list predicates to the EventTypeUpdate builder.
func (etuo *EventTypeUpdateOne) Where(ps ...predicate.EventType) *EventTypeUpdateOne {
	etuo.mutation.Where(ps...)
	return etuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (etuo *EventTypeUpdateOne) Select(field string, fields ...string) *EventTypeUpdateOne {
	etuo.fields = append([]string{field}, fields...)
	return etuo
}

// Save executes the query and returns the updated EventType entity.
func (etuo *EventTypeUpdateOne) Save(ctx context.Context) (*EventType, error) {
	return withHooks(ctx, etuo.sqlSave, etuo.mutation, etuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (etuo *EventTypeUpdateOne) SaveX(ctx context.Context) *EventType {
	node, err := etuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (etuo *EventTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := etuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etuo *EventTypeUpdateOne) ExecX(ctx context.Context) {
	if err := etuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (etuo *EventTypeUpdateOne) sqlSave(ctx context.Context) (_node *EventType, err error) {
	_spec := sqlgraph.NewUpdateSpec(eventtype.Table, eventtype.Columns, sqlgraph.NewFieldSpec(eventtype.FieldID, field.TypeInt32))
	id, ok := etuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EventType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := etuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, eventtype.FieldID)
		for _, f := range fields {
			if !eventtype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != eventtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := etuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := etuo.mutation.Types(); ok {
		_spec.SetField(eventtype.FieldTypes, field.TypeString, value)
	}
	if etuo.mutation.TypesCleared() {
		_spec.ClearField(eventtype.FieldTypes, field.TypeString)
	}
	if etuo.mutation.EarthquakeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventtype.EarthquakeTable,
			Columns: []string{eventtype.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := etuo.mutation.EarthquakeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   eventtype.EarthquakeTable,
			Columns: []string{eventtype.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EventType{config: etuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, etuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	etuo.mutation.done = true
	return _node, nil
}