// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/earthquake"
	"entdemo/ent/geometry"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GeometryCreate is the builder for creating a Geometry entity.
type GeometryCreate struct {
	config
	mutation *GeometryMutation
	hooks    []Hook
}

// SetEarthquakeID sets the "earthquake_id" field.
func (gc *GeometryCreate) SetEarthquakeID(i int32) *GeometryCreate {
	gc.mutation.SetEarthquakeID(i)
	return gc
}

// SetNillableEarthquakeID sets the "earthquake_id" field if the given value is not nil.
func (gc *GeometryCreate) SetNillableEarthquakeID(i *int32) *GeometryCreate {
	if i != nil {
		gc.SetEarthquakeID(*i)
	}
	return gc
}

// SetLongitude sets the "longitude" field.
func (gc *GeometryCreate) SetLongitude(f float64) *GeometryCreate {
	gc.mutation.SetLongitude(f)
	return gc
}

// SetLatitude sets the "latitude" field.
func (gc *GeometryCreate) SetLatitude(f float64) *GeometryCreate {
	gc.mutation.SetLatitude(f)
	return gc
}

// SetDepth sets the "depth" field.
func (gc *GeometryCreate) SetDepth(f float64) *GeometryCreate {
	gc.mutation.SetDepth(f)
	return gc
}

// SetPlace sets the "place" field.
func (gc *GeometryCreate) SetPlace(s string) *GeometryCreate {
	gc.mutation.SetPlace(s)
	return gc
}

// SetID sets the "id" field.
func (gc *GeometryCreate) SetID(i int32) *GeometryCreate {
	gc.mutation.SetID(i)
	return gc
}

// SetEarthquake sets the "earthquake" edge to the Earthquake entity.
func (gc *GeometryCreate) SetEarthquake(e *Earthquake) *GeometryCreate {
	return gc.SetEarthquakeID(e.ID)
}

// Mutation returns the GeometryMutation object of the builder.
func (gc *GeometryCreate) Mutation() *GeometryMutation {
	return gc.mutation
}

// Save creates the Geometry in the database.
func (gc *GeometryCreate) Save(ctx context.Context) (*Geometry, error) {
	return withHooks(ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GeometryCreate) SaveX(ctx context.Context) *Geometry {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GeometryCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GeometryCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GeometryCreate) check() error {
	if _, ok := gc.mutation.Longitude(); !ok {
		return &ValidationError{Name: "longitude", err: errors.New(`ent: missing required field "Geometry.longitude"`)}
	}
	if _, ok := gc.mutation.Latitude(); !ok {
		return &ValidationError{Name: "latitude", err: errors.New(`ent: missing required field "Geometry.latitude"`)}
	}
	if _, ok := gc.mutation.Depth(); !ok {
		return &ValidationError{Name: "depth", err: errors.New(`ent: missing required field "Geometry.depth"`)}
	}
	if _, ok := gc.mutation.Place(); !ok {
		return &ValidationError{Name: "place", err: errors.New(`ent: missing required field "Geometry.place"`)}
	}
	return nil
}

func (gc *GeometryCreate) sqlSave(ctx context.Context) (*Geometry, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	gc.mutation.id = &_node.ID
	gc.mutation.done = true
	return _node, nil
}

func (gc *GeometryCreate) createSpec() (*Geometry, *sqlgraph.CreateSpec) {
	var (
		_node = &Geometry{config: gc.config}
		_spec = sqlgraph.NewCreateSpec(geometry.Table, sqlgraph.NewFieldSpec(geometry.FieldID, field.TypeInt32))
	)
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.Longitude(); ok {
		_spec.SetField(geometry.FieldLongitude, field.TypeFloat64, value)
		_node.Longitude = value
	}
	if value, ok := gc.mutation.Latitude(); ok {
		_spec.SetField(geometry.FieldLatitude, field.TypeFloat64, value)
		_node.Latitude = value
	}
	if value, ok := gc.mutation.Depth(); ok {
		_spec.SetField(geometry.FieldDepth, field.TypeFloat64, value)
		_node.Depth = value
	}
	if value, ok := gc.mutation.Place(); ok {
		_spec.SetField(geometry.FieldPlace, field.TypeString, value)
		_node.Place = value
	}
	if nodes := gc.mutation.EarthquakeIDs(); len(nodes) > 0 {
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
		_node.EarthquakeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GeometryCreateBulk is the builder for creating many Geometry entities in bulk.
type GeometryCreateBulk struct {
	config
	err      error
	builders []*GeometryCreate
}

// Save creates the Geometry entities in the database.
func (gcb *GeometryCreateBulk) Save(ctx context.Context) ([]*Geometry, error) {
	if gcb.err != nil {
		return nil, gcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Geometry, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GeometryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GeometryCreateBulk) SaveX(ctx context.Context) []*Geometry {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GeometryCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GeometryCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}