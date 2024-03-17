// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/earthquake"
	"entdemo/ent/report"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReportCreate is the builder for creating a Report entity.
type ReportCreate struct {
	config
	mutation *ReportMutation
	hooks    []Hook
}

// SetFelt sets the "felt" field.
func (rc *ReportCreate) SetFelt(i int32) *ReportCreate {
	rc.mutation.SetFelt(i)
	return rc
}

// SetNillableFelt sets the "felt" field if the given value is not nil.
func (rc *ReportCreate) SetNillableFelt(i *int32) *ReportCreate {
	if i != nil {
		rc.SetFelt(*i)
	}
	return rc
}

// SetCdi sets the "cdi" field.
func (rc *ReportCreate) SetCdi(f float64) *ReportCreate {
	rc.mutation.SetCdi(f)
	return rc
}

// SetNillableCdi sets the "cdi" field if the given value is not nil.
func (rc *ReportCreate) SetNillableCdi(f *float64) *ReportCreate {
	if f != nil {
		rc.SetCdi(*f)
	}
	return rc
}

// SetMmi sets the "mmi" field.
func (rc *ReportCreate) SetMmi(f float64) *ReportCreate {
	rc.mutation.SetMmi(f)
	return rc
}

// SetNillableMmi sets the "mmi" field if the given value is not nil.
func (rc *ReportCreate) SetNillableMmi(f *float64) *ReportCreate {
	if f != nil {
		rc.SetMmi(*f)
	}
	return rc
}

// SetAlert sets the "alert" field.
func (rc *ReportCreate) SetAlert(s string) *ReportCreate {
	rc.mutation.SetAlert(s)
	return rc
}

// SetNillableAlert sets the "alert" field if the given value is not nil.
func (rc *ReportCreate) SetNillableAlert(s *string) *ReportCreate {
	if s != nil {
		rc.SetAlert(*s)
	}
	return rc
}

// SetCreatedAt sets the "created_at" field.
func (rc *ReportCreate) SetCreatedAt(t time.Time) *ReportCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *ReportCreate) SetUpdatedAt(t time.Time) *ReportCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetDeletedAt sets the "deleted_at" field.
func (rc *ReportCreate) SetDeletedAt(t time.Time) *ReportCreate {
	rc.mutation.SetDeletedAt(t)
	return rc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rc *ReportCreate) SetNillableDeletedAt(t *time.Time) *ReportCreate {
	if t != nil {
		rc.SetDeletedAt(*t)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *ReportCreate) SetID(i int) *ReportCreate {
	rc.mutation.SetID(i)
	return rc
}

// AddEarthquakeIDs adds the "earthquakes" edge to the Earthquake entity by IDs.
func (rc *ReportCreate) AddEarthquakeIDs(ids ...int) *ReportCreate {
	rc.mutation.AddEarthquakeIDs(ids...)
	return rc
}

// AddEarthquakes adds the "earthquakes" edges to the Earthquake entity.
func (rc *ReportCreate) AddEarthquakes(e ...*Earthquake) *ReportCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return rc.AddEarthquakeIDs(ids...)
}

// Mutation returns the ReportMutation object of the builder.
func (rc *ReportCreate) Mutation() *ReportMutation {
	return rc.mutation
}

// Save creates the Report in the database.
func (rc *ReportCreate) Save(ctx context.Context) (*Report, error) {
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReportCreate) SaveX(ctx context.Context) *Report {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReportCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReportCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReportCreate) check() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Report.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Report.updated_at"`)}
	}
	return nil
}

func (rc *ReportCreate) sqlSave(ctx context.Context) (*Report, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ReportCreate) createSpec() (*Report, *sqlgraph.CreateSpec) {
	var (
		_node = &Report{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(report.Table, sqlgraph.NewFieldSpec(report.FieldID, field.TypeInt))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.Felt(); ok {
		_spec.SetField(report.FieldFelt, field.TypeInt32, value)
		_node.Felt = value
	}
	if value, ok := rc.mutation.Cdi(); ok {
		_spec.SetField(report.FieldCdi, field.TypeFloat64, value)
		_node.Cdi = value
	}
	if value, ok := rc.mutation.Mmi(); ok {
		_spec.SetField(report.FieldMmi, field.TypeFloat64, value)
		_node.Mmi = value
	}
	if value, ok := rc.mutation.Alert(); ok {
		_spec.SetField(report.FieldAlert, field.TypeString, value)
		_node.Alert = value
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.SetField(report.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.SetField(report.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.DeletedAt(); ok {
		_spec.SetField(report.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if nodes := rc.mutation.EarthquakesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   report.EarthquakesTable,
			Columns: []string{report.EarthquakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ReportCreateBulk is the builder for creating many Report entities in bulk.
type ReportCreateBulk struct {
	config
	err      error
	builders []*ReportCreate
}

// Save creates the Report entities in the database.
func (rcb *ReportCreateBulk) Save(ctx context.Context) ([]*Report, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Report, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReportMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReportCreateBulk) SaveX(ctx context.Context) []*Report {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReportCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReportCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
