// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/earthquake"
	"entdemo/ent/geometry"
	"entdemo/ent/predicate"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GeometryUpdate is the builder for updating Geometry entities.
type GeometryUpdate struct {
	config
	hooks    []Hook
	mutation *GeometryMutation
}

// Where appends a list predicates to the GeometryUpdate builder.
func (gu *GeometryUpdate) Where(ps ...predicate.Geometry) *GeometryUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetEarthquakeID sets the "earthquake_id" field.
func (gu *GeometryUpdate) SetEarthquakeID(i int32) *GeometryUpdate {
	gu.mutation.SetEarthquakeID(i)
	return gu
}

// SetNillableEarthquakeID sets the "earthquake_id" field if the given value is not nil.
func (gu *GeometryUpdate) SetNillableEarthquakeID(i *int32) *GeometryUpdate {
	if i != nil {
		gu.SetEarthquakeID(*i)
	}
	return gu
}

// ClearEarthquakeID clears the value of the "earthquake_id" field.
func (gu *GeometryUpdate) ClearEarthquakeID() *GeometryUpdate {
	gu.mutation.ClearEarthquakeID()
	return gu
}

// SetLongitude sets the "longitude" field.
func (gu *GeometryUpdate) SetLongitude(f float64) *GeometryUpdate {
	gu.mutation.ResetLongitude()
	gu.mutation.SetLongitude(f)
	return gu
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (gu *GeometryUpdate) SetNillableLongitude(f *float64) *GeometryUpdate {
	if f != nil {
		gu.SetLongitude(*f)
	}
	return gu
}

// AddLongitude adds f to the "longitude" field.
func (gu *GeometryUpdate) AddLongitude(f float64) *GeometryUpdate {
	gu.mutation.AddLongitude(f)
	return gu
}

// SetLatitude sets the "latitude" field.
func (gu *GeometryUpdate) SetLatitude(f float64) *GeometryUpdate {
	gu.mutation.ResetLatitude()
	gu.mutation.SetLatitude(f)
	return gu
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (gu *GeometryUpdate) SetNillableLatitude(f *float64) *GeometryUpdate {
	if f != nil {
		gu.SetLatitude(*f)
	}
	return gu
}

// AddLatitude adds f to the "latitude" field.
func (gu *GeometryUpdate) AddLatitude(f float64) *GeometryUpdate {
	gu.mutation.AddLatitude(f)
	return gu
}

// SetDepth sets the "depth" field.
func (gu *GeometryUpdate) SetDepth(f float64) *GeometryUpdate {
	gu.mutation.ResetDepth()
	gu.mutation.SetDepth(f)
	return gu
}

// SetNillableDepth sets the "depth" field if the given value is not nil.
func (gu *GeometryUpdate) SetNillableDepth(f *float64) *GeometryUpdate {
	if f != nil {
		gu.SetDepth(*f)
	}
	return gu
}

// AddDepth adds f to the "depth" field.
func (gu *GeometryUpdate) AddDepth(f float64) *GeometryUpdate {
	gu.mutation.AddDepth(f)
	return gu
}

// SetPlace sets the "place" field.
func (gu *GeometryUpdate) SetPlace(s string) *GeometryUpdate {
	gu.mutation.SetPlace(s)
	return gu
}

// SetNillablePlace sets the "place" field if the given value is not nil.
func (gu *GeometryUpdate) SetNillablePlace(s *string) *GeometryUpdate {
	if s != nil {
		gu.SetPlace(*s)
	}
	return gu
}

// SetEarthquake sets the "earthquake" edge to the Earthquake entity.
func (gu *GeometryUpdate) SetEarthquake(e *Earthquake) *GeometryUpdate {
	return gu.SetEarthquakeID(e.ID)
}

// Mutation returns the GeometryMutation object of the builder.
func (gu *GeometryUpdate) Mutation() *GeometryMutation {
	return gu.mutation
}

// ClearEarthquake clears the "earthquake" edge to the Earthquake entity.
func (gu *GeometryUpdate) ClearEarthquake() *GeometryUpdate {
	gu.mutation.ClearEarthquake()
	return gu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GeometryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GeometryUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GeometryUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GeometryUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GeometryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(geometry.Table, geometry.Columns, sqlgraph.NewFieldSpec(geometry.FieldID, field.TypeInt32))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Longitude(); ok {
		_spec.SetField(geometry.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := gu.mutation.AddedLongitude(); ok {
		_spec.AddField(geometry.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := gu.mutation.Latitude(); ok {
		_spec.SetField(geometry.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := gu.mutation.AddedLatitude(); ok {
		_spec.AddField(geometry.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := gu.mutation.Depth(); ok {
		_spec.SetField(geometry.FieldDepth, field.TypeFloat64, value)
	}
	if value, ok := gu.mutation.AddedDepth(); ok {
		_spec.AddField(geometry.FieldDepth, field.TypeFloat64, value)
	}
	if value, ok := gu.mutation.Place(); ok {
		_spec.SetField(geometry.FieldPlace, field.TypeString, value)
	}
	if gu.mutation.EarthquakeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   geometry.EarthquakeTable,
			Columns: []string{geometry.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.EarthquakeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   geometry.EarthquakeTable,
			Columns: []string{geometry.EarthquakeColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{geometry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GeometryUpdateOne is the builder for updating a single Geometry entity.
type GeometryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GeometryMutation
}

// SetEarthquakeID sets the "earthquake_id" field.
func (guo *GeometryUpdateOne) SetEarthquakeID(i int32) *GeometryUpdateOne {
	guo.mutation.SetEarthquakeID(i)
	return guo
}

// SetNillableEarthquakeID sets the "earthquake_id" field if the given value is not nil.
func (guo *GeometryUpdateOne) SetNillableEarthquakeID(i *int32) *GeometryUpdateOne {
	if i != nil {
		guo.SetEarthquakeID(*i)
	}
	return guo
}

// ClearEarthquakeID clears the value of the "earthquake_id" field.
func (guo *GeometryUpdateOne) ClearEarthquakeID() *GeometryUpdateOne {
	guo.mutation.ClearEarthquakeID()
	return guo
}

// SetLongitude sets the "longitude" field.
func (guo *GeometryUpdateOne) SetLongitude(f float64) *GeometryUpdateOne {
	guo.mutation.ResetLongitude()
	guo.mutation.SetLongitude(f)
	return guo
}

// SetNillableLongitude sets the "longitude" field if the given value is not nil.
func (guo *GeometryUpdateOne) SetNillableLongitude(f *float64) *GeometryUpdateOne {
	if f != nil {
		guo.SetLongitude(*f)
	}
	return guo
}

// AddLongitude adds f to the "longitude" field.
func (guo *GeometryUpdateOne) AddLongitude(f float64) *GeometryUpdateOne {
	guo.mutation.AddLongitude(f)
	return guo
}

// SetLatitude sets the "latitude" field.
func (guo *GeometryUpdateOne) SetLatitude(f float64) *GeometryUpdateOne {
	guo.mutation.ResetLatitude()
	guo.mutation.SetLatitude(f)
	return guo
}

// SetNillableLatitude sets the "latitude" field if the given value is not nil.
func (guo *GeometryUpdateOne) SetNillableLatitude(f *float64) *GeometryUpdateOne {
	if f != nil {
		guo.SetLatitude(*f)
	}
	return guo
}

// AddLatitude adds f to the "latitude" field.
func (guo *GeometryUpdateOne) AddLatitude(f float64) *GeometryUpdateOne {
	guo.mutation.AddLatitude(f)
	return guo
}

// SetDepth sets the "depth" field.
func (guo *GeometryUpdateOne) SetDepth(f float64) *GeometryUpdateOne {
	guo.mutation.ResetDepth()
	guo.mutation.SetDepth(f)
	return guo
}

// SetNillableDepth sets the "depth" field if the given value is not nil.
func (guo *GeometryUpdateOne) SetNillableDepth(f *float64) *GeometryUpdateOne {
	if f != nil {
		guo.SetDepth(*f)
	}
	return guo
}

// AddDepth adds f to the "depth" field.
func (guo *GeometryUpdateOne) AddDepth(f float64) *GeometryUpdateOne {
	guo.mutation.AddDepth(f)
	return guo
}

// SetPlace sets the "place" field.
func (guo *GeometryUpdateOne) SetPlace(s string) *GeometryUpdateOne {
	guo.mutation.SetPlace(s)
	return guo
}

// SetNillablePlace sets the "place" field if the given value is not nil.
func (guo *GeometryUpdateOne) SetNillablePlace(s *string) *GeometryUpdateOne {
	if s != nil {
		guo.SetPlace(*s)
	}
	return guo
}

// SetEarthquake sets the "earthquake" edge to the Earthquake entity.
func (guo *GeometryUpdateOne) SetEarthquake(e *Earthquake) *GeometryUpdateOne {
	return guo.SetEarthquakeID(e.ID)
}

// Mutation returns the GeometryMutation object of the builder.
func (guo *GeometryUpdateOne) Mutation() *GeometryMutation {
	return guo.mutation
}

// ClearEarthquake clears the "earthquake" edge to the Earthquake entity.
func (guo *GeometryUpdateOne) ClearEarthquake() *GeometryUpdateOne {
	guo.mutation.ClearEarthquake()
	return guo
}

// Where appends a list predicates to the GeometryUpdate builder.
func (guo *GeometryUpdateOne) Where(ps ...predicate.Geometry) *GeometryUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GeometryUpdateOne) Select(field string, fields ...string) *GeometryUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Geometry entity.
func (guo *GeometryUpdateOne) Save(ctx context.Context) (*Geometry, error) {
	return withHooks(ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GeometryUpdateOne) SaveX(ctx context.Context) *Geometry {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GeometryUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GeometryUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GeometryUpdateOne) sqlSave(ctx context.Context) (_node *Geometry, err error) {
	_spec := sqlgraph.NewUpdateSpec(geometry.Table, geometry.Columns, sqlgraph.NewFieldSpec(geometry.FieldID, field.TypeInt32))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Geometry.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, geometry.FieldID)
		for _, f := range fields {
			if !geometry.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != geometry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.Longitude(); ok {
		_spec.SetField(geometry.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := guo.mutation.AddedLongitude(); ok {
		_spec.AddField(geometry.FieldLongitude, field.TypeFloat64, value)
	}
	if value, ok := guo.mutation.Latitude(); ok {
		_spec.SetField(geometry.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := guo.mutation.AddedLatitude(); ok {
		_spec.AddField(geometry.FieldLatitude, field.TypeFloat64, value)
	}
	if value, ok := guo.mutation.Depth(); ok {
		_spec.SetField(geometry.FieldDepth, field.TypeFloat64, value)
	}
	if value, ok := guo.mutation.AddedDepth(); ok {
		_spec.AddField(geometry.FieldDepth, field.TypeFloat64, value)
	}
	if value, ok := guo.mutation.Place(); ok {
		_spec.SetField(geometry.FieldPlace, field.TypeString, value)
	}
	if guo.mutation.EarthquakeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   geometry.EarthquakeTable,
			Columns: []string{geometry.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.EarthquakeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   geometry.EarthquakeTable,
			Columns: []string{geometry.EarthquakeColumn},
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
	_node = &Geometry{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{geometry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}