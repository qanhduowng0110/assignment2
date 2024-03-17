// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/apireq"
	"entdemo/ent/predicate"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ApireqUpdate is the builder for updating Apireq entities.
type ApireqUpdate struct {
	config
	hooks    []Hook
	mutation *ApireqMutation
}

// Where appends a list predicates to the ApireqUpdate builder.
func (au *ApireqUpdate) Where(ps ...predicate.Apireq) *ApireqUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetReqTime sets the "req_time" field.
func (au *ApireqUpdate) SetReqTime(t time.Time) *ApireqUpdate {
	au.mutation.SetReqTime(t)
	return au
}

// SetNillableReqTime sets the "req_time" field if the given value is not nil.
func (au *ApireqUpdate) SetNillableReqTime(t *time.Time) *ApireqUpdate {
	if t != nil {
		au.SetReqTime(*t)
	}
	return au
}

// SetReqParam sets the "req_param" field.
func (au *ApireqUpdate) SetReqParam(s struct{}) *ApireqUpdate {
	au.mutation.SetReqParam(s)
	return au
}

// SetNillableReqParam sets the "req_param" field if the given value is not nil.
func (au *ApireqUpdate) SetNillableReqParam(s *struct{}) *ApireqUpdate {
	if s != nil {
		au.SetReqParam(*s)
	}
	return au
}

// SetReqBody sets the "req_body" field.
func (au *ApireqUpdate) SetReqBody(s struct{}) *ApireqUpdate {
	au.mutation.SetReqBody(s)
	return au
}

// SetNillableReqBody sets the "req_body" field if the given value is not nil.
func (au *ApireqUpdate) SetNillableReqBody(s *struct{}) *ApireqUpdate {
	if s != nil {
		au.SetReqBody(*s)
	}
	return au
}

// ClearReqBody clears the value of the "req_body" field.
func (au *ApireqUpdate) ClearReqBody() *ApireqUpdate {
	au.mutation.ClearReqBody()
	return au
}

// SetReqHeaders sets the "req_headers" field.
func (au *ApireqUpdate) SetReqHeaders(s struct{}) *ApireqUpdate {
	au.mutation.SetReqHeaders(s)
	return au
}

// SetNillableReqHeaders sets the "req_headers" field if the given value is not nil.
func (au *ApireqUpdate) SetNillableReqHeaders(s *struct{}) *ApireqUpdate {
	if s != nil {
		au.SetReqHeaders(*s)
	}
	return au
}

// ClearReqHeaders clears the value of the "req_headers" field.
func (au *ApireqUpdate) ClearReqHeaders() *ApireqUpdate {
	au.mutation.ClearReqHeaders()
	return au
}

// SetReqMetadata sets the "req_metadata" field.
func (au *ApireqUpdate) SetReqMetadata(s struct{}) *ApireqUpdate {
	au.mutation.SetReqMetadata(s)
	return au
}

// SetNillableReqMetadata sets the "req_metadata" field if the given value is not nil.
func (au *ApireqUpdate) SetNillableReqMetadata(s *struct{}) *ApireqUpdate {
	if s != nil {
		au.SetReqMetadata(*s)
	}
	return au
}

// ClearReqMetadata clears the value of the "req_metadata" field.
func (au *ApireqUpdate) ClearReqMetadata() *ApireqUpdate {
	au.mutation.ClearReqMetadata()
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *ApireqUpdate) SetCreatedAt(t time.Time) *ApireqUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *ApireqUpdate) SetNillableCreatedAt(t *time.Time) *ApireqUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *ApireqUpdate) SetUpdatedAt(t time.Time) *ApireqUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (au *ApireqUpdate) SetNillableUpdatedAt(t *time.Time) *ApireqUpdate {
	if t != nil {
		au.SetUpdatedAt(*t)
	}
	return au
}

// SetDeletedAt sets the "deleted_at" field.
func (au *ApireqUpdate) SetDeletedAt(t time.Time) *ApireqUpdate {
	au.mutation.SetDeletedAt(t)
	return au
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (au *ApireqUpdate) SetNillableDeletedAt(t *time.Time) *ApireqUpdate {
	if t != nil {
		au.SetDeletedAt(*t)
	}
	return au
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (au *ApireqUpdate) ClearDeletedAt() *ApireqUpdate {
	au.mutation.ClearDeletedAt()
	return au
}

// Mutation returns the ApireqMutation object of the builder.
func (au *ApireqUpdate) Mutation() *ApireqMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ApireqUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ApireqUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ApireqUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ApireqUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ApireqUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(apireq.Table, apireq.Columns, sqlgraph.NewFieldSpec(apireq.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.ReqTime(); ok {
		_spec.SetField(apireq.FieldReqTime, field.TypeTime, value)
	}
	if value, ok := au.mutation.ReqParam(); ok {
		_spec.SetField(apireq.FieldReqParam, field.TypeJSON, value)
	}
	if value, ok := au.mutation.ReqBody(); ok {
		_spec.SetField(apireq.FieldReqBody, field.TypeJSON, value)
	}
	if au.mutation.ReqBodyCleared() {
		_spec.ClearField(apireq.FieldReqBody, field.TypeJSON)
	}
	if value, ok := au.mutation.ReqHeaders(); ok {
		_spec.SetField(apireq.FieldReqHeaders, field.TypeJSON, value)
	}
	if au.mutation.ReqHeadersCleared() {
		_spec.ClearField(apireq.FieldReqHeaders, field.TypeJSON)
	}
	if value, ok := au.mutation.ReqMetadata(); ok {
		_spec.SetField(apireq.FieldReqMetadata, field.TypeJSON, value)
	}
	if au.mutation.ReqMetadataCleared() {
		_spec.ClearField(apireq.FieldReqMetadata, field.TypeJSON)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(apireq.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(apireq.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.DeletedAt(); ok {
		_spec.SetField(apireq.FieldDeletedAt, field.TypeTime, value)
	}
	if au.mutation.DeletedAtCleared() {
		_spec.ClearField(apireq.FieldDeletedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apireq.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ApireqUpdateOne is the builder for updating a single Apireq entity.
type ApireqUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApireqMutation
}

// SetReqTime sets the "req_time" field.
func (auo *ApireqUpdateOne) SetReqTime(t time.Time) *ApireqUpdateOne {
	auo.mutation.SetReqTime(t)
	return auo
}

// SetNillableReqTime sets the "req_time" field if the given value is not nil.
func (auo *ApireqUpdateOne) SetNillableReqTime(t *time.Time) *ApireqUpdateOne {
	if t != nil {
		auo.SetReqTime(*t)
	}
	return auo
}

// SetReqParam sets the "req_param" field.
func (auo *ApireqUpdateOne) SetReqParam(s struct{}) *ApireqUpdateOne {
	auo.mutation.SetReqParam(s)
	return auo
}

// SetNillableReqParam sets the "req_param" field if the given value is not nil.
func (auo *ApireqUpdateOne) SetNillableReqParam(s *struct{}) *ApireqUpdateOne {
	if s != nil {
		auo.SetReqParam(*s)
	}
	return auo
}

// SetReqBody sets the "req_body" field.
func (auo *ApireqUpdateOne) SetReqBody(s struct{}) *ApireqUpdateOne {
	auo.mutation.SetReqBody(s)
	return auo
}

// SetNillableReqBody sets the "req_body" field if the given value is not nil.
func (auo *ApireqUpdateOne) SetNillableReqBody(s *struct{}) *ApireqUpdateOne {
	if s != nil {
		auo.SetReqBody(*s)
	}
	return auo
}

// ClearReqBody clears the value of the "req_body" field.
func (auo *ApireqUpdateOne) ClearReqBody() *ApireqUpdateOne {
	auo.mutation.ClearReqBody()
	return auo
}

// SetReqHeaders sets the "req_headers" field.
func (auo *ApireqUpdateOne) SetReqHeaders(s struct{}) *ApireqUpdateOne {
	auo.mutation.SetReqHeaders(s)
	return auo
}

// SetNillableReqHeaders sets the "req_headers" field if the given value is not nil.
func (auo *ApireqUpdateOne) SetNillableReqHeaders(s *struct{}) *ApireqUpdateOne {
	if s != nil {
		auo.SetReqHeaders(*s)
	}
	return auo
}

// ClearReqHeaders clears the value of the "req_headers" field.
func (auo *ApireqUpdateOne) ClearReqHeaders() *ApireqUpdateOne {
	auo.mutation.ClearReqHeaders()
	return auo
}

// SetReqMetadata sets the "req_metadata" field.
func (auo *ApireqUpdateOne) SetReqMetadata(s struct{}) *ApireqUpdateOne {
	auo.mutation.SetReqMetadata(s)
	return auo
}

// SetNillableReqMetadata sets the "req_metadata" field if the given value is not nil.
func (auo *ApireqUpdateOne) SetNillableReqMetadata(s *struct{}) *ApireqUpdateOne {
	if s != nil {
		auo.SetReqMetadata(*s)
	}
	return auo
}

// ClearReqMetadata clears the value of the "req_metadata" field.
func (auo *ApireqUpdateOne) ClearReqMetadata() *ApireqUpdateOne {
	auo.mutation.ClearReqMetadata()
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *ApireqUpdateOne) SetCreatedAt(t time.Time) *ApireqUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *ApireqUpdateOne) SetNillableCreatedAt(t *time.Time) *ApireqUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *ApireqUpdateOne) SetUpdatedAt(t time.Time) *ApireqUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (auo *ApireqUpdateOne) SetNillableUpdatedAt(t *time.Time) *ApireqUpdateOne {
	if t != nil {
		auo.SetUpdatedAt(*t)
	}
	return auo
}

// SetDeletedAt sets the "deleted_at" field.
func (auo *ApireqUpdateOne) SetDeletedAt(t time.Time) *ApireqUpdateOne {
	auo.mutation.SetDeletedAt(t)
	return auo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auo *ApireqUpdateOne) SetNillableDeletedAt(t *time.Time) *ApireqUpdateOne {
	if t != nil {
		auo.SetDeletedAt(*t)
	}
	return auo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (auo *ApireqUpdateOne) ClearDeletedAt() *ApireqUpdateOne {
	auo.mutation.ClearDeletedAt()
	return auo
}

// Mutation returns the ApireqMutation object of the builder.
func (auo *ApireqUpdateOne) Mutation() *ApireqMutation {
	return auo.mutation
}

// Where appends a list predicates to the ApireqUpdate builder.
func (auo *ApireqUpdateOne) Where(ps ...predicate.Apireq) *ApireqUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ApireqUpdateOne) Select(field string, fields ...string) *ApireqUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Apireq entity.
func (auo *ApireqUpdateOne) Save(ctx context.Context) (*Apireq, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ApireqUpdateOne) SaveX(ctx context.Context) *Apireq {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ApireqUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ApireqUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ApireqUpdateOne) sqlSave(ctx context.Context) (_node *Apireq, err error) {
	_spec := sqlgraph.NewUpdateSpec(apireq.Table, apireq.Columns, sqlgraph.NewFieldSpec(apireq.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Apireq.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apireq.FieldID)
		for _, f := range fields {
			if !apireq.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != apireq.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.ReqTime(); ok {
		_spec.SetField(apireq.FieldReqTime, field.TypeTime, value)
	}
	if value, ok := auo.mutation.ReqParam(); ok {
		_spec.SetField(apireq.FieldReqParam, field.TypeJSON, value)
	}
	if value, ok := auo.mutation.ReqBody(); ok {
		_spec.SetField(apireq.FieldReqBody, field.TypeJSON, value)
	}
	if auo.mutation.ReqBodyCleared() {
		_spec.ClearField(apireq.FieldReqBody, field.TypeJSON)
	}
	if value, ok := auo.mutation.ReqHeaders(); ok {
		_spec.SetField(apireq.FieldReqHeaders, field.TypeJSON, value)
	}
	if auo.mutation.ReqHeadersCleared() {
		_spec.ClearField(apireq.FieldReqHeaders, field.TypeJSON)
	}
	if value, ok := auo.mutation.ReqMetadata(); ok {
		_spec.SetField(apireq.FieldReqMetadata, field.TypeJSON, value)
	}
	if auo.mutation.ReqMetadataCleared() {
		_spec.ClearField(apireq.FieldReqMetadata, field.TypeJSON)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(apireq.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(apireq.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.DeletedAt(); ok {
		_spec.SetField(apireq.FieldDeletedAt, field.TypeTime, value)
	}
	if auo.mutation.DeletedAtCleared() {
		_spec.ClearField(apireq.FieldDeletedAt, field.TypeTime)
	}
	_node = &Apireq{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apireq.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
