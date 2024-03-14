// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/apirequestlog"
	"entdemo/ent/predicate"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// APIRequestLogUpdate is the builder for updating APIRequestLog entities.
type APIRequestLogUpdate struct {
	config
	hooks    []Hook
	mutation *APIRequestLogMutation
}

// Where appends a list predicates to the APIRequestLogUpdate builder.
func (arlu *APIRequestLogUpdate) Where(ps ...predicate.APIRequestLog) *APIRequestLogUpdate {
	arlu.mutation.Where(ps...)
	return arlu
}

// SetRequestDatetime sets the "request_datetime" field.
func (arlu *APIRequestLogUpdate) SetRequestDatetime(t time.Time) *APIRequestLogUpdate {
	arlu.mutation.SetRequestDatetime(t)
	return arlu
}

// SetNillableRequestDatetime sets the "request_datetime" field if the given value is not nil.
func (arlu *APIRequestLogUpdate) SetNillableRequestDatetime(t *time.Time) *APIRequestLogUpdate {
	if t != nil {
		arlu.SetRequestDatetime(*t)
	}
	return arlu
}

// SetRequestParameters sets the "request_parameters" field.
func (arlu *APIRequestLogUpdate) SetRequestParameters(m map[string]interface{}) *APIRequestLogUpdate {
	arlu.mutation.SetRequestParameters(m)
	return arlu
}

// SetRequestBody sets the "request_body" field.
func (arlu *APIRequestLogUpdate) SetRequestBody(m map[string]interface{}) *APIRequestLogUpdate {
	arlu.mutation.SetRequestBody(m)
	return arlu
}

// SetRequestHeaders sets the "request_headers" field.
func (arlu *APIRequestLogUpdate) SetRequestHeaders(m map[string]interface{}) *APIRequestLogUpdate {
	arlu.mutation.SetRequestHeaders(m)
	return arlu
}

// SetRequestMetadata sets the "request_metadata" field.
func (arlu *APIRequestLogUpdate) SetRequestMetadata(m map[string]interface{}) *APIRequestLogUpdate {
	arlu.mutation.SetRequestMetadata(m)
	return arlu
}

// SetCreatedAt sets the "created_at" field.
func (arlu *APIRequestLogUpdate) SetCreatedAt(t time.Time) *APIRequestLogUpdate {
	arlu.mutation.SetCreatedAt(t)
	return arlu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (arlu *APIRequestLogUpdate) SetNillableCreatedAt(t *time.Time) *APIRequestLogUpdate {
	if t != nil {
		arlu.SetCreatedAt(*t)
	}
	return arlu
}

// SetUpdatedAt sets the "updated_at" field.
func (arlu *APIRequestLogUpdate) SetUpdatedAt(t time.Time) *APIRequestLogUpdate {
	arlu.mutation.SetUpdatedAt(t)
	return arlu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (arlu *APIRequestLogUpdate) SetNillableUpdatedAt(t *time.Time) *APIRequestLogUpdate {
	if t != nil {
		arlu.SetUpdatedAt(*t)
	}
	return arlu
}

// SetDeletedAt sets the "deleted_at" field.
func (arlu *APIRequestLogUpdate) SetDeletedAt(t time.Time) *APIRequestLogUpdate {
	arlu.mutation.SetDeletedAt(t)
	return arlu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (arlu *APIRequestLogUpdate) SetNillableDeletedAt(t *time.Time) *APIRequestLogUpdate {
	if t != nil {
		arlu.SetDeletedAt(*t)
	}
	return arlu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (arlu *APIRequestLogUpdate) ClearDeletedAt() *APIRequestLogUpdate {
	arlu.mutation.ClearDeletedAt()
	return arlu
}

// Mutation returns the APIRequestLogMutation object of the builder.
func (arlu *APIRequestLogUpdate) Mutation() *APIRequestLogMutation {
	return arlu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (arlu *APIRequestLogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, arlu.sqlSave, arlu.mutation, arlu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (arlu *APIRequestLogUpdate) SaveX(ctx context.Context) int {
	affected, err := arlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (arlu *APIRequestLogUpdate) Exec(ctx context.Context) error {
	_, err := arlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arlu *APIRequestLogUpdate) ExecX(ctx context.Context) {
	if err := arlu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (arlu *APIRequestLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(apirequestlog.Table, apirequestlog.Columns, sqlgraph.NewFieldSpec(apirequestlog.FieldID, field.TypeInt))
	if ps := arlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := arlu.mutation.RequestDatetime(); ok {
		_spec.SetField(apirequestlog.FieldRequestDatetime, field.TypeTime, value)
	}
	if value, ok := arlu.mutation.RequestParameters(); ok {
		_spec.SetField(apirequestlog.FieldRequestParameters, field.TypeJSON, value)
	}
	if value, ok := arlu.mutation.RequestBody(); ok {
		_spec.SetField(apirequestlog.FieldRequestBody, field.TypeJSON, value)
	}
	if value, ok := arlu.mutation.RequestHeaders(); ok {
		_spec.SetField(apirequestlog.FieldRequestHeaders, field.TypeJSON, value)
	}
	if value, ok := arlu.mutation.RequestMetadata(); ok {
		_spec.SetField(apirequestlog.FieldRequestMetadata, field.TypeJSON, value)
	}
	if value, ok := arlu.mutation.CreatedAt(); ok {
		_spec.SetField(apirequestlog.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := arlu.mutation.UpdatedAt(); ok {
		_spec.SetField(apirequestlog.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := arlu.mutation.DeletedAt(); ok {
		_spec.SetField(apirequestlog.FieldDeletedAt, field.TypeTime, value)
	}
	if arlu.mutation.DeletedAtCleared() {
		_spec.ClearField(apirequestlog.FieldDeletedAt, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, arlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apirequestlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	arlu.mutation.done = true
	return n, nil
}

// APIRequestLogUpdateOne is the builder for updating a single APIRequestLog entity.
type APIRequestLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *APIRequestLogMutation
}

// SetRequestDatetime sets the "request_datetime" field.
func (arluo *APIRequestLogUpdateOne) SetRequestDatetime(t time.Time) *APIRequestLogUpdateOne {
	arluo.mutation.SetRequestDatetime(t)
	return arluo
}

// SetNillableRequestDatetime sets the "request_datetime" field if the given value is not nil.
func (arluo *APIRequestLogUpdateOne) SetNillableRequestDatetime(t *time.Time) *APIRequestLogUpdateOne {
	if t != nil {
		arluo.SetRequestDatetime(*t)
	}
	return arluo
}

// SetRequestParameters sets the "request_parameters" field.
func (arluo *APIRequestLogUpdateOne) SetRequestParameters(m map[string]interface{}) *APIRequestLogUpdateOne {
	arluo.mutation.SetRequestParameters(m)
	return arluo
}

// SetRequestBody sets the "request_body" field.
func (arluo *APIRequestLogUpdateOne) SetRequestBody(m map[string]interface{}) *APIRequestLogUpdateOne {
	arluo.mutation.SetRequestBody(m)
	return arluo
}

// SetRequestHeaders sets the "request_headers" field.
func (arluo *APIRequestLogUpdateOne) SetRequestHeaders(m map[string]interface{}) *APIRequestLogUpdateOne {
	arluo.mutation.SetRequestHeaders(m)
	return arluo
}

// SetRequestMetadata sets the "request_metadata" field.
func (arluo *APIRequestLogUpdateOne) SetRequestMetadata(m map[string]interface{}) *APIRequestLogUpdateOne {
	arluo.mutation.SetRequestMetadata(m)
	return arluo
}

// SetCreatedAt sets the "created_at" field.
func (arluo *APIRequestLogUpdateOne) SetCreatedAt(t time.Time) *APIRequestLogUpdateOne {
	arluo.mutation.SetCreatedAt(t)
	return arluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (arluo *APIRequestLogUpdateOne) SetNillableCreatedAt(t *time.Time) *APIRequestLogUpdateOne {
	if t != nil {
		arluo.SetCreatedAt(*t)
	}
	return arluo
}

// SetUpdatedAt sets the "updated_at" field.
func (arluo *APIRequestLogUpdateOne) SetUpdatedAt(t time.Time) *APIRequestLogUpdateOne {
	arluo.mutation.SetUpdatedAt(t)
	return arluo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (arluo *APIRequestLogUpdateOne) SetNillableUpdatedAt(t *time.Time) *APIRequestLogUpdateOne {
	if t != nil {
		arluo.SetUpdatedAt(*t)
	}
	return arluo
}

// SetDeletedAt sets the "deleted_at" field.
func (arluo *APIRequestLogUpdateOne) SetDeletedAt(t time.Time) *APIRequestLogUpdateOne {
	arluo.mutation.SetDeletedAt(t)
	return arluo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (arluo *APIRequestLogUpdateOne) SetNillableDeletedAt(t *time.Time) *APIRequestLogUpdateOne {
	if t != nil {
		arluo.SetDeletedAt(*t)
	}
	return arluo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (arluo *APIRequestLogUpdateOne) ClearDeletedAt() *APIRequestLogUpdateOne {
	arluo.mutation.ClearDeletedAt()
	return arluo
}

// Mutation returns the APIRequestLogMutation object of the builder.
func (arluo *APIRequestLogUpdateOne) Mutation() *APIRequestLogMutation {
	return arluo.mutation
}

// Where appends a list predicates to the APIRequestLogUpdate builder.
func (arluo *APIRequestLogUpdateOne) Where(ps ...predicate.APIRequestLog) *APIRequestLogUpdateOne {
	arluo.mutation.Where(ps...)
	return arluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (arluo *APIRequestLogUpdateOne) Select(field string, fields ...string) *APIRequestLogUpdateOne {
	arluo.fields = append([]string{field}, fields...)
	return arluo
}

// Save executes the query and returns the updated APIRequestLog entity.
func (arluo *APIRequestLogUpdateOne) Save(ctx context.Context) (*APIRequestLog, error) {
	return withHooks(ctx, arluo.sqlSave, arluo.mutation, arluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (arluo *APIRequestLogUpdateOne) SaveX(ctx context.Context) *APIRequestLog {
	node, err := arluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (arluo *APIRequestLogUpdateOne) Exec(ctx context.Context) error {
	_, err := arluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (arluo *APIRequestLogUpdateOne) ExecX(ctx context.Context) {
	if err := arluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (arluo *APIRequestLogUpdateOne) sqlSave(ctx context.Context) (_node *APIRequestLog, err error) {
	_spec := sqlgraph.NewUpdateSpec(apirequestlog.Table, apirequestlog.Columns, sqlgraph.NewFieldSpec(apirequestlog.FieldID, field.TypeInt))
	id, ok := arluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "APIRequestLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := arluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apirequestlog.FieldID)
		for _, f := range fields {
			if !apirequestlog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != apirequestlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := arluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := arluo.mutation.RequestDatetime(); ok {
		_spec.SetField(apirequestlog.FieldRequestDatetime, field.TypeTime, value)
	}
	if value, ok := arluo.mutation.RequestParameters(); ok {
		_spec.SetField(apirequestlog.FieldRequestParameters, field.TypeJSON, value)
	}
	if value, ok := arluo.mutation.RequestBody(); ok {
		_spec.SetField(apirequestlog.FieldRequestBody, field.TypeJSON, value)
	}
	if value, ok := arluo.mutation.RequestHeaders(); ok {
		_spec.SetField(apirequestlog.FieldRequestHeaders, field.TypeJSON, value)
	}
	if value, ok := arluo.mutation.RequestMetadata(); ok {
		_spec.SetField(apirequestlog.FieldRequestMetadata, field.TypeJSON, value)
	}
	if value, ok := arluo.mutation.CreatedAt(); ok {
		_spec.SetField(apirequestlog.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := arluo.mutation.UpdatedAt(); ok {
		_spec.SetField(apirequestlog.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := arluo.mutation.DeletedAt(); ok {
		_spec.SetField(apirequestlog.FieldDeletedAt, field.TypeTime, value)
	}
	if arluo.mutation.DeletedAtCleared() {
		_spec.ClearField(apirequestlog.FieldDeletedAt, field.TypeTime)
	}
	_node = &APIRequestLog{config: arluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, arluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apirequestlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	arluo.mutation.done = true
	return _node, nil
}