// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/earthquake"
	"entdemo/ent/featuretype"
	"entdemo/ent/ftypeearthquake"
	"entdemo/ent/predicate"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FtypeEarthquakeUpdate is the builder for updating FtypeEarthquake entities.
type FtypeEarthquakeUpdate struct {
	config
	hooks    []Hook
	mutation *FtypeEarthquakeMutation
}

// Where appends a list predicates to the FtypeEarthquakeUpdate builder.
func (feu *FtypeEarthquakeUpdate) Where(ps ...predicate.FtypeEarthquake) *FtypeEarthquakeUpdate {
	feu.mutation.Where(ps...)
	return feu
}

// SetFtID sets the "ft_id" field.
func (feu *FtypeEarthquakeUpdate) SetFtID(i int) *FtypeEarthquakeUpdate {
	feu.mutation.SetFtID(i)
	return feu
}

// SetNillableFtID sets the "ft_id" field if the given value is not nil.
func (feu *FtypeEarthquakeUpdate) SetNillableFtID(i *int) *FtypeEarthquakeUpdate {
	if i != nil {
		feu.SetFtID(*i)
	}
	return feu
}

// ClearFtID clears the value of the "ft_id" field.
func (feu *FtypeEarthquakeUpdate) ClearFtID() *FtypeEarthquakeUpdate {
	feu.mutation.ClearFtID()
	return feu
}

// SetEqID sets the "eq_id" field.
func (feu *FtypeEarthquakeUpdate) SetEqID(i int) *FtypeEarthquakeUpdate {
	feu.mutation.SetEqID(i)
	return feu
}

// SetNillableEqID sets the "eq_id" field if the given value is not nil.
func (feu *FtypeEarthquakeUpdate) SetNillableEqID(i *int) *FtypeEarthquakeUpdate {
	if i != nil {
		feu.SetEqID(*i)
	}
	return feu
}

// ClearEqID clears the value of the "eq_id" field.
func (feu *FtypeEarthquakeUpdate) ClearEqID() *FtypeEarthquakeUpdate {
	feu.mutation.ClearEqID()
	return feu
}

// SetCreatedAt sets the "created_at" field.
func (feu *FtypeEarthquakeUpdate) SetCreatedAt(t time.Time) *FtypeEarthquakeUpdate {
	feu.mutation.SetCreatedAt(t)
	return feu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (feu *FtypeEarthquakeUpdate) SetNillableCreatedAt(t *time.Time) *FtypeEarthquakeUpdate {
	if t != nil {
		feu.SetCreatedAt(*t)
	}
	return feu
}

// SetUpdatedAt sets the "updated_at" field.
func (feu *FtypeEarthquakeUpdate) SetUpdatedAt(t time.Time) *FtypeEarthquakeUpdate {
	feu.mutation.SetUpdatedAt(t)
	return feu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (feu *FtypeEarthquakeUpdate) SetNillableUpdatedAt(t *time.Time) *FtypeEarthquakeUpdate {
	if t != nil {
		feu.SetUpdatedAt(*t)
	}
	return feu
}

// SetDeletedAt sets the "deleted_at" field.
func (feu *FtypeEarthquakeUpdate) SetDeletedAt(t time.Time) *FtypeEarthquakeUpdate {
	feu.mutation.SetDeletedAt(t)
	return feu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (feu *FtypeEarthquakeUpdate) SetNillableDeletedAt(t *time.Time) *FtypeEarthquakeUpdate {
	if t != nil {
		feu.SetDeletedAt(*t)
	}
	return feu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (feu *FtypeEarthquakeUpdate) ClearDeletedAt() *FtypeEarthquakeUpdate {
	feu.mutation.ClearDeletedAt()
	return feu
}

// SetEarthquakeID sets the "earthquake" edge to the Earthquake entity by ID.
func (feu *FtypeEarthquakeUpdate) SetEarthquakeID(id int) *FtypeEarthquakeUpdate {
	feu.mutation.SetEarthquakeID(id)
	return feu
}

// SetNillableEarthquakeID sets the "earthquake" edge to the Earthquake entity by ID if the given value is not nil.
func (feu *FtypeEarthquakeUpdate) SetNillableEarthquakeID(id *int) *FtypeEarthquakeUpdate {
	if id != nil {
		feu = feu.SetEarthquakeID(*id)
	}
	return feu
}

// SetEarthquake sets the "earthquake" edge to the Earthquake entity.
func (feu *FtypeEarthquakeUpdate) SetEarthquake(e *Earthquake) *FtypeEarthquakeUpdate {
	return feu.SetEarthquakeID(e.ID)
}

// SetFeatureTypeID sets the "feature_type" edge to the FeatureType entity by ID.
func (feu *FtypeEarthquakeUpdate) SetFeatureTypeID(id int) *FtypeEarthquakeUpdate {
	feu.mutation.SetFeatureTypeID(id)
	return feu
}

// SetNillableFeatureTypeID sets the "feature_type" edge to the FeatureType entity by ID if the given value is not nil.
func (feu *FtypeEarthquakeUpdate) SetNillableFeatureTypeID(id *int) *FtypeEarthquakeUpdate {
	if id != nil {
		feu = feu.SetFeatureTypeID(*id)
	}
	return feu
}

// SetFeatureType sets the "feature_type" edge to the FeatureType entity.
func (feu *FtypeEarthquakeUpdate) SetFeatureType(f *FeatureType) *FtypeEarthquakeUpdate {
	return feu.SetFeatureTypeID(f.ID)
}

// Mutation returns the FtypeEarthquakeMutation object of the builder.
func (feu *FtypeEarthquakeUpdate) Mutation() *FtypeEarthquakeMutation {
	return feu.mutation
}

// ClearEarthquake clears the "earthquake" edge to the Earthquake entity.
func (feu *FtypeEarthquakeUpdate) ClearEarthquake() *FtypeEarthquakeUpdate {
	feu.mutation.ClearEarthquake()
	return feu
}

// ClearFeatureType clears the "feature_type" edge to the FeatureType entity.
func (feu *FtypeEarthquakeUpdate) ClearFeatureType() *FtypeEarthquakeUpdate {
	feu.mutation.ClearFeatureType()
	return feu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (feu *FtypeEarthquakeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, feu.sqlSave, feu.mutation, feu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (feu *FtypeEarthquakeUpdate) SaveX(ctx context.Context) int {
	affected, err := feu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (feu *FtypeEarthquakeUpdate) Exec(ctx context.Context) error {
	_, err := feu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (feu *FtypeEarthquakeUpdate) ExecX(ctx context.Context) {
	if err := feu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (feu *FtypeEarthquakeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(ftypeearthquake.Table, ftypeearthquake.Columns, sqlgraph.NewFieldSpec(ftypeearthquake.FieldID, field.TypeInt))
	if ps := feu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := feu.mutation.CreatedAt(); ok {
		_spec.SetField(ftypeearthquake.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := feu.mutation.UpdatedAt(); ok {
		_spec.SetField(ftypeearthquake.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := feu.mutation.DeletedAt(); ok {
		_spec.SetField(ftypeearthquake.FieldDeletedAt, field.TypeTime, value)
	}
	if feu.mutation.DeletedAtCleared() {
		_spec.ClearField(ftypeearthquake.FieldDeletedAt, field.TypeTime)
	}
	if feu.mutation.EarthquakeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ftypeearthquake.EarthquakeTable,
			Columns: []string{ftypeearthquake.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := feu.mutation.EarthquakeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ftypeearthquake.EarthquakeTable,
			Columns: []string{ftypeearthquake.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if feu.mutation.FeatureTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ftypeearthquake.FeatureTypeTable,
			Columns: []string{ftypeearthquake.FeatureTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(featuretype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := feu.mutation.FeatureTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ftypeearthquake.FeatureTypeTable,
			Columns: []string{ftypeearthquake.FeatureTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(featuretype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, feu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ftypeearthquake.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	feu.mutation.done = true
	return n, nil
}

// FtypeEarthquakeUpdateOne is the builder for updating a single FtypeEarthquake entity.
type FtypeEarthquakeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FtypeEarthquakeMutation
}

// SetFtID sets the "ft_id" field.
func (feuo *FtypeEarthquakeUpdateOne) SetFtID(i int) *FtypeEarthquakeUpdateOne {
	feuo.mutation.SetFtID(i)
	return feuo
}

// SetNillableFtID sets the "ft_id" field if the given value is not nil.
func (feuo *FtypeEarthquakeUpdateOne) SetNillableFtID(i *int) *FtypeEarthquakeUpdateOne {
	if i != nil {
		feuo.SetFtID(*i)
	}
	return feuo
}

// ClearFtID clears the value of the "ft_id" field.
func (feuo *FtypeEarthquakeUpdateOne) ClearFtID() *FtypeEarthquakeUpdateOne {
	feuo.mutation.ClearFtID()
	return feuo
}

// SetEqID sets the "eq_id" field.
func (feuo *FtypeEarthquakeUpdateOne) SetEqID(i int) *FtypeEarthquakeUpdateOne {
	feuo.mutation.SetEqID(i)
	return feuo
}

// SetNillableEqID sets the "eq_id" field if the given value is not nil.
func (feuo *FtypeEarthquakeUpdateOne) SetNillableEqID(i *int) *FtypeEarthquakeUpdateOne {
	if i != nil {
		feuo.SetEqID(*i)
	}
	return feuo
}

// ClearEqID clears the value of the "eq_id" field.
func (feuo *FtypeEarthquakeUpdateOne) ClearEqID() *FtypeEarthquakeUpdateOne {
	feuo.mutation.ClearEqID()
	return feuo
}

// SetCreatedAt sets the "created_at" field.
func (feuo *FtypeEarthquakeUpdateOne) SetCreatedAt(t time.Time) *FtypeEarthquakeUpdateOne {
	feuo.mutation.SetCreatedAt(t)
	return feuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (feuo *FtypeEarthquakeUpdateOne) SetNillableCreatedAt(t *time.Time) *FtypeEarthquakeUpdateOne {
	if t != nil {
		feuo.SetCreatedAt(*t)
	}
	return feuo
}

// SetUpdatedAt sets the "updated_at" field.
func (feuo *FtypeEarthquakeUpdateOne) SetUpdatedAt(t time.Time) *FtypeEarthquakeUpdateOne {
	feuo.mutation.SetUpdatedAt(t)
	return feuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (feuo *FtypeEarthquakeUpdateOne) SetNillableUpdatedAt(t *time.Time) *FtypeEarthquakeUpdateOne {
	if t != nil {
		feuo.SetUpdatedAt(*t)
	}
	return feuo
}

// SetDeletedAt sets the "deleted_at" field.
func (feuo *FtypeEarthquakeUpdateOne) SetDeletedAt(t time.Time) *FtypeEarthquakeUpdateOne {
	feuo.mutation.SetDeletedAt(t)
	return feuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (feuo *FtypeEarthquakeUpdateOne) SetNillableDeletedAt(t *time.Time) *FtypeEarthquakeUpdateOne {
	if t != nil {
		feuo.SetDeletedAt(*t)
	}
	return feuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (feuo *FtypeEarthquakeUpdateOne) ClearDeletedAt() *FtypeEarthquakeUpdateOne {
	feuo.mutation.ClearDeletedAt()
	return feuo
}

// SetEarthquakeID sets the "earthquake" edge to the Earthquake entity by ID.
func (feuo *FtypeEarthquakeUpdateOne) SetEarthquakeID(id int) *FtypeEarthquakeUpdateOne {
	feuo.mutation.SetEarthquakeID(id)
	return feuo
}

// SetNillableEarthquakeID sets the "earthquake" edge to the Earthquake entity by ID if the given value is not nil.
func (feuo *FtypeEarthquakeUpdateOne) SetNillableEarthquakeID(id *int) *FtypeEarthquakeUpdateOne {
	if id != nil {
		feuo = feuo.SetEarthquakeID(*id)
	}
	return feuo
}

// SetEarthquake sets the "earthquake" edge to the Earthquake entity.
func (feuo *FtypeEarthquakeUpdateOne) SetEarthquake(e *Earthquake) *FtypeEarthquakeUpdateOne {
	return feuo.SetEarthquakeID(e.ID)
}

// SetFeatureTypeID sets the "feature_type" edge to the FeatureType entity by ID.
func (feuo *FtypeEarthquakeUpdateOne) SetFeatureTypeID(id int) *FtypeEarthquakeUpdateOne {
	feuo.mutation.SetFeatureTypeID(id)
	return feuo
}

// SetNillableFeatureTypeID sets the "feature_type" edge to the FeatureType entity by ID if the given value is not nil.
func (feuo *FtypeEarthquakeUpdateOne) SetNillableFeatureTypeID(id *int) *FtypeEarthquakeUpdateOne {
	if id != nil {
		feuo = feuo.SetFeatureTypeID(*id)
	}
	return feuo
}

// SetFeatureType sets the "feature_type" edge to the FeatureType entity.
func (feuo *FtypeEarthquakeUpdateOne) SetFeatureType(f *FeatureType) *FtypeEarthquakeUpdateOne {
	return feuo.SetFeatureTypeID(f.ID)
}

// Mutation returns the FtypeEarthquakeMutation object of the builder.
func (feuo *FtypeEarthquakeUpdateOne) Mutation() *FtypeEarthquakeMutation {
	return feuo.mutation
}

// ClearEarthquake clears the "earthquake" edge to the Earthquake entity.
func (feuo *FtypeEarthquakeUpdateOne) ClearEarthquake() *FtypeEarthquakeUpdateOne {
	feuo.mutation.ClearEarthquake()
	return feuo
}

// ClearFeatureType clears the "feature_type" edge to the FeatureType entity.
func (feuo *FtypeEarthquakeUpdateOne) ClearFeatureType() *FtypeEarthquakeUpdateOne {
	feuo.mutation.ClearFeatureType()
	return feuo
}

// Where appends a list predicates to the FtypeEarthquakeUpdate builder.
func (feuo *FtypeEarthquakeUpdateOne) Where(ps ...predicate.FtypeEarthquake) *FtypeEarthquakeUpdateOne {
	feuo.mutation.Where(ps...)
	return feuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (feuo *FtypeEarthquakeUpdateOne) Select(field string, fields ...string) *FtypeEarthquakeUpdateOne {
	feuo.fields = append([]string{field}, fields...)
	return feuo
}

// Save executes the query and returns the updated FtypeEarthquake entity.
func (feuo *FtypeEarthquakeUpdateOne) Save(ctx context.Context) (*FtypeEarthquake, error) {
	return withHooks(ctx, feuo.sqlSave, feuo.mutation, feuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (feuo *FtypeEarthquakeUpdateOne) SaveX(ctx context.Context) *FtypeEarthquake {
	node, err := feuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (feuo *FtypeEarthquakeUpdateOne) Exec(ctx context.Context) error {
	_, err := feuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (feuo *FtypeEarthquakeUpdateOne) ExecX(ctx context.Context) {
	if err := feuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (feuo *FtypeEarthquakeUpdateOne) sqlSave(ctx context.Context) (_node *FtypeEarthquake, err error) {
	_spec := sqlgraph.NewUpdateSpec(ftypeearthquake.Table, ftypeearthquake.Columns, sqlgraph.NewFieldSpec(ftypeearthquake.FieldID, field.TypeInt))
	id, ok := feuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FtypeEarthquake.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := feuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ftypeearthquake.FieldID)
		for _, f := range fields {
			if !ftypeearthquake.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ftypeearthquake.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := feuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := feuo.mutation.CreatedAt(); ok {
		_spec.SetField(ftypeearthquake.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := feuo.mutation.UpdatedAt(); ok {
		_spec.SetField(ftypeearthquake.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := feuo.mutation.DeletedAt(); ok {
		_spec.SetField(ftypeearthquake.FieldDeletedAt, field.TypeTime, value)
	}
	if feuo.mutation.DeletedAtCleared() {
		_spec.ClearField(ftypeearthquake.FieldDeletedAt, field.TypeTime)
	}
	if feuo.mutation.EarthquakeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ftypeearthquake.EarthquakeTable,
			Columns: []string{ftypeearthquake.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := feuo.mutation.EarthquakeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ftypeearthquake.EarthquakeTable,
			Columns: []string{ftypeearthquake.EarthquakeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if feuo.mutation.FeatureTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ftypeearthquake.FeatureTypeTable,
			Columns: []string{ftypeearthquake.FeatureTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(featuretype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := feuo.mutation.FeatureTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ftypeearthquake.FeatureTypeTable,
			Columns: []string{ftypeearthquake.FeatureTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(featuretype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FtypeEarthquake{config: feuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, feuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ftypeearthquake.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	feuo.mutation.done = true
	return _node, nil
}
