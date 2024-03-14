// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/predicate"
	"entdemo/ent/schemamigration"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SchemaMigrationUpdate is the builder for updating SchemaMigration entities.
type SchemaMigrationUpdate struct {
	config
	hooks    []Hook
	mutation *SchemaMigrationMutation
}

// Where appends a list predicates to the SchemaMigrationUpdate builder.
func (smu *SchemaMigrationUpdate) Where(ps ...predicate.SchemaMigration) *SchemaMigrationUpdate {
	smu.mutation.Where(ps...)
	return smu
}

// SetDirty sets the "dirty" field.
func (smu *SchemaMigrationUpdate) SetDirty(b bool) *SchemaMigrationUpdate {
	smu.mutation.SetDirty(b)
	return smu
}

// SetNillableDirty sets the "dirty" field if the given value is not nil.
func (smu *SchemaMigrationUpdate) SetNillableDirty(b *bool) *SchemaMigrationUpdate {
	if b != nil {
		smu.SetDirty(*b)
	}
	return smu
}

// Mutation returns the SchemaMigrationMutation object of the builder.
func (smu *SchemaMigrationUpdate) Mutation() *SchemaMigrationMutation {
	return smu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (smu *SchemaMigrationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, smu.sqlSave, smu.mutation, smu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (smu *SchemaMigrationUpdate) SaveX(ctx context.Context) int {
	affected, err := smu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (smu *SchemaMigrationUpdate) Exec(ctx context.Context) error {
	_, err := smu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smu *SchemaMigrationUpdate) ExecX(ctx context.Context) {
	if err := smu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (smu *SchemaMigrationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(schemamigration.Table, schemamigration.Columns, sqlgraph.NewFieldSpec(schemamigration.FieldID, field.TypeInt))
	if ps := smu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := smu.mutation.Dirty(); ok {
		_spec.SetField(schemamigration.FieldDirty, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, smu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{schemamigration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	smu.mutation.done = true
	return n, nil
}

// SchemaMigrationUpdateOne is the builder for updating a single SchemaMigration entity.
type SchemaMigrationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SchemaMigrationMutation
}

// SetDirty sets the "dirty" field.
func (smuo *SchemaMigrationUpdateOne) SetDirty(b bool) *SchemaMigrationUpdateOne {
	smuo.mutation.SetDirty(b)
	return smuo
}

// SetNillableDirty sets the "dirty" field if the given value is not nil.
func (smuo *SchemaMigrationUpdateOne) SetNillableDirty(b *bool) *SchemaMigrationUpdateOne {
	if b != nil {
		smuo.SetDirty(*b)
	}
	return smuo
}

// Mutation returns the SchemaMigrationMutation object of the builder.
func (smuo *SchemaMigrationUpdateOne) Mutation() *SchemaMigrationMutation {
	return smuo.mutation
}

// Where appends a list predicates to the SchemaMigrationUpdate builder.
func (smuo *SchemaMigrationUpdateOne) Where(ps ...predicate.SchemaMigration) *SchemaMigrationUpdateOne {
	smuo.mutation.Where(ps...)
	return smuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (smuo *SchemaMigrationUpdateOne) Select(field string, fields ...string) *SchemaMigrationUpdateOne {
	smuo.fields = append([]string{field}, fields...)
	return smuo
}

// Save executes the query and returns the updated SchemaMigration entity.
func (smuo *SchemaMigrationUpdateOne) Save(ctx context.Context) (*SchemaMigration, error) {
	return withHooks(ctx, smuo.sqlSave, smuo.mutation, smuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (smuo *SchemaMigrationUpdateOne) SaveX(ctx context.Context) *SchemaMigration {
	node, err := smuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (smuo *SchemaMigrationUpdateOne) Exec(ctx context.Context) error {
	_, err := smuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smuo *SchemaMigrationUpdateOne) ExecX(ctx context.Context) {
	if err := smuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (smuo *SchemaMigrationUpdateOne) sqlSave(ctx context.Context) (_node *SchemaMigration, err error) {
	_spec := sqlgraph.NewUpdateSpec(schemamigration.Table, schemamigration.Columns, sqlgraph.NewFieldSpec(schemamigration.FieldID, field.TypeInt))
	id, ok := smuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SchemaMigration.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := smuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, schemamigration.FieldID)
		for _, f := range fields {
			if !schemamigration.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != schemamigration.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := smuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := smuo.mutation.Dirty(); ok {
		_spec.SetField(schemamigration.FieldDirty, field.TypeBool, value)
	}
	_node = &SchemaMigration{config: smuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, smuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{schemamigration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	smuo.mutation.done = true
	return _node, nil
}