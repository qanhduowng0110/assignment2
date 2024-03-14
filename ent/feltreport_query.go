// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"entdemo/ent/earthquake"
	"entdemo/ent/feltreport"
	"entdemo/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FeltReportQuery is the builder for querying FeltReport entities.
type FeltReportQuery struct {
	config
	ctx            *QueryContext
	order          []feltreport.OrderOption
	inters         []Interceptor
	predicates     []predicate.FeltReport
	withEarthquake *EarthquakeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FeltReportQuery builder.
func (frq *FeltReportQuery) Where(ps ...predicate.FeltReport) *FeltReportQuery {
	frq.predicates = append(frq.predicates, ps...)
	return frq
}

// Limit the number of records to be returned by this query.
func (frq *FeltReportQuery) Limit(limit int) *FeltReportQuery {
	frq.ctx.Limit = &limit
	return frq
}

// Offset to start from.
func (frq *FeltReportQuery) Offset(offset int) *FeltReportQuery {
	frq.ctx.Offset = &offset
	return frq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (frq *FeltReportQuery) Unique(unique bool) *FeltReportQuery {
	frq.ctx.Unique = &unique
	return frq
}

// Order specifies how the records should be ordered.
func (frq *FeltReportQuery) Order(o ...feltreport.OrderOption) *FeltReportQuery {
	frq.order = append(frq.order, o...)
	return frq
}

// QueryEarthquake chains the current query on the "earthquake" edge.
func (frq *FeltReportQuery) QueryEarthquake() *EarthquakeQuery {
	query := (&EarthquakeClient{config: frq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := frq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := frq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(feltreport.Table, feltreport.FieldID, selector),
			sqlgraph.To(earthquake.Table, earthquake.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, feltreport.EarthquakeTable, feltreport.EarthquakeColumn),
		)
		fromU = sqlgraph.SetNeighbors(frq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FeltReport entity from the query.
// Returns a *NotFoundError when no FeltReport was found.
func (frq *FeltReportQuery) First(ctx context.Context) (*FeltReport, error) {
	nodes, err := frq.Limit(1).All(setContextOp(ctx, frq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{feltreport.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (frq *FeltReportQuery) FirstX(ctx context.Context) *FeltReport {
	node, err := frq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FeltReport ID from the query.
// Returns a *NotFoundError when no FeltReport ID was found.
func (frq *FeltReportQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = frq.Limit(1).IDs(setContextOp(ctx, frq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{feltreport.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (frq *FeltReportQuery) FirstIDX(ctx context.Context) int32 {
	id, err := frq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FeltReport entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FeltReport entity is found.
// Returns a *NotFoundError when no FeltReport entities are found.
func (frq *FeltReportQuery) Only(ctx context.Context) (*FeltReport, error) {
	nodes, err := frq.Limit(2).All(setContextOp(ctx, frq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{feltreport.Label}
	default:
		return nil, &NotSingularError{feltreport.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (frq *FeltReportQuery) OnlyX(ctx context.Context) *FeltReport {
	node, err := frq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FeltReport ID in the query.
// Returns a *NotSingularError when more than one FeltReport ID is found.
// Returns a *NotFoundError when no entities are found.
func (frq *FeltReportQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = frq.Limit(2).IDs(setContextOp(ctx, frq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{feltreport.Label}
	default:
		err = &NotSingularError{feltreport.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (frq *FeltReportQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := frq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FeltReports.
func (frq *FeltReportQuery) All(ctx context.Context) ([]*FeltReport, error) {
	ctx = setContextOp(ctx, frq.ctx, "All")
	if err := frq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FeltReport, *FeltReportQuery]()
	return withInterceptors[[]*FeltReport](ctx, frq, qr, frq.inters)
}

// AllX is like All, but panics if an error occurs.
func (frq *FeltReportQuery) AllX(ctx context.Context) []*FeltReport {
	nodes, err := frq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FeltReport IDs.
func (frq *FeltReportQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if frq.ctx.Unique == nil && frq.path != nil {
		frq.Unique(true)
	}
	ctx = setContextOp(ctx, frq.ctx, "IDs")
	if err = frq.Select(feltreport.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (frq *FeltReportQuery) IDsX(ctx context.Context) []int32 {
	ids, err := frq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (frq *FeltReportQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, frq.ctx, "Count")
	if err := frq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, frq, querierCount[*FeltReportQuery](), frq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (frq *FeltReportQuery) CountX(ctx context.Context) int {
	count, err := frq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (frq *FeltReportQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, frq.ctx, "Exist")
	switch _, err := frq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (frq *FeltReportQuery) ExistX(ctx context.Context) bool {
	exist, err := frq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FeltReportQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (frq *FeltReportQuery) Clone() *FeltReportQuery {
	if frq == nil {
		return nil
	}
	return &FeltReportQuery{
		config:         frq.config,
		ctx:            frq.ctx.Clone(),
		order:          append([]feltreport.OrderOption{}, frq.order...),
		inters:         append([]Interceptor{}, frq.inters...),
		predicates:     append([]predicate.FeltReport{}, frq.predicates...),
		withEarthquake: frq.withEarthquake.Clone(),
		// clone intermediate query.
		sql:  frq.sql.Clone(),
		path: frq.path,
	}
}

// WithEarthquake tells the query-builder to eager-load the nodes that are connected to
// the "earthquake" edge. The optional arguments are used to configure the query builder of the edge.
func (frq *FeltReportQuery) WithEarthquake(opts ...func(*EarthquakeQuery)) *FeltReportQuery {
	query := (&EarthquakeClient{config: frq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	frq.withEarthquake = query
	return frq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EarthquakeID int32 `json:"earthquake_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FeltReport.Query().
//		GroupBy(feltreport.FieldEarthquakeID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (frq *FeltReportQuery) GroupBy(field string, fields ...string) *FeltReportGroupBy {
	frq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FeltReportGroupBy{build: frq}
	grbuild.flds = &frq.ctx.Fields
	grbuild.label = feltreport.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EarthquakeID int32 `json:"earthquake_id,omitempty"`
//	}
//
//	client.FeltReport.Query().
//		Select(feltreport.FieldEarthquakeID).
//		Scan(ctx, &v)
func (frq *FeltReportQuery) Select(fields ...string) *FeltReportSelect {
	frq.ctx.Fields = append(frq.ctx.Fields, fields...)
	sbuild := &FeltReportSelect{FeltReportQuery: frq}
	sbuild.label = feltreport.Label
	sbuild.flds, sbuild.scan = &frq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FeltReportSelect configured with the given aggregations.
func (frq *FeltReportQuery) Aggregate(fns ...AggregateFunc) *FeltReportSelect {
	return frq.Select().Aggregate(fns...)
}

func (frq *FeltReportQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range frq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, frq); err != nil {
				return err
			}
		}
	}
	for _, f := range frq.ctx.Fields {
		if !feltreport.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if frq.path != nil {
		prev, err := frq.path(ctx)
		if err != nil {
			return err
		}
		frq.sql = prev
	}
	return nil
}

func (frq *FeltReportQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FeltReport, error) {
	var (
		nodes       = []*FeltReport{}
		_spec       = frq.querySpec()
		loadedTypes = [1]bool{
			frq.withEarthquake != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FeltReport).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FeltReport{config: frq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, frq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := frq.withEarthquake; query != nil {
		if err := frq.loadEarthquake(ctx, query, nodes, nil,
			func(n *FeltReport, e *Earthquake) { n.Edges.Earthquake = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (frq *FeltReportQuery) loadEarthquake(ctx context.Context, query *EarthquakeQuery, nodes []*FeltReport, init func(*FeltReport), assign func(*FeltReport, *Earthquake)) error {
	ids := make([]int32, 0, len(nodes))
	nodeids := make(map[int32][]*FeltReport)
	for i := range nodes {
		fk := nodes[i].EarthquakeID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(earthquake.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "earthquake_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (frq *FeltReportQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := frq.querySpec()
	_spec.Node.Columns = frq.ctx.Fields
	if len(frq.ctx.Fields) > 0 {
		_spec.Unique = frq.ctx.Unique != nil && *frq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, frq.driver, _spec)
}

func (frq *FeltReportQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(feltreport.Table, feltreport.Columns, sqlgraph.NewFieldSpec(feltreport.FieldID, field.TypeInt32))
	_spec.From = frq.sql
	if unique := frq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if frq.path != nil {
		_spec.Unique = true
	}
	if fields := frq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, feltreport.FieldID)
		for i := range fields {
			if fields[i] != feltreport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if frq.withEarthquake != nil {
			_spec.Node.AddColumnOnce(feltreport.FieldEarthquakeID)
		}
	}
	if ps := frq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := frq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := frq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := frq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (frq *FeltReportQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(frq.driver.Dialect())
	t1 := builder.Table(feltreport.Table)
	columns := frq.ctx.Fields
	if len(columns) == 0 {
		columns = feltreport.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if frq.sql != nil {
		selector = frq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if frq.ctx.Unique != nil && *frq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range frq.predicates {
		p(selector)
	}
	for _, p := range frq.order {
		p(selector)
	}
	if offset := frq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := frq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FeltReportGroupBy is the group-by builder for FeltReport entities.
type FeltReportGroupBy struct {
	selector
	build *FeltReportQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (frgb *FeltReportGroupBy) Aggregate(fns ...AggregateFunc) *FeltReportGroupBy {
	frgb.fns = append(frgb.fns, fns...)
	return frgb
}

// Scan applies the selector query and scans the result into the given value.
func (frgb *FeltReportGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, frgb.build.ctx, "GroupBy")
	if err := frgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FeltReportQuery, *FeltReportGroupBy](ctx, frgb.build, frgb, frgb.build.inters, v)
}

func (frgb *FeltReportGroupBy) sqlScan(ctx context.Context, root *FeltReportQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(frgb.fns))
	for _, fn := range frgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*frgb.flds)+len(frgb.fns))
		for _, f := range *frgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*frgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := frgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FeltReportSelect is the builder for selecting fields of FeltReport entities.
type FeltReportSelect struct {
	*FeltReportQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (frs *FeltReportSelect) Aggregate(fns ...AggregateFunc) *FeltReportSelect {
	frs.fns = append(frs.fns, fns...)
	return frs
}

// Scan applies the selector query and scans the result into the given value.
func (frs *FeltReportSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, frs.ctx, "Select")
	if err := frs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FeltReportQuery, *FeltReportSelect](ctx, frs.FeltReportQuery, frs, frs.inters, v)
}

func (frs *FeltReportSelect) sqlScan(ctx context.Context, root *FeltReportQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(frs.fns))
	for _, fn := range frs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*frs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := frs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
