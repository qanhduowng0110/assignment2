// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"entdemo/ent/earthquake"
	"entdemo/ent/predicate"
	"entdemo/ent/report"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReportQuery is the builder for querying Report entities.
type ReportQuery struct {
	config
	ctx             *QueryContext
	order           []report.OrderOption
	inters          []Interceptor
	predicates      []predicate.Report
	withEarthquakes *EarthquakeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ReportQuery builder.
func (rq *ReportQuery) Where(ps ...predicate.Report) *ReportQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *ReportQuery) Limit(limit int) *ReportQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *ReportQuery) Offset(offset int) *ReportQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *ReportQuery) Unique(unique bool) *ReportQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *ReportQuery) Order(o ...report.OrderOption) *ReportQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryEarthquakes chains the current query on the "earthquakes" edge.
func (rq *ReportQuery) QueryEarthquakes() *EarthquakeQuery {
	query := (&EarthquakeClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(report.Table, report.FieldID, selector),
			sqlgraph.To(earthquake.Table, earthquake.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, report.EarthquakesTable, report.EarthquakesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Report entity from the query.
// Returns a *NotFoundError when no Report was found.
func (rq *ReportQuery) First(ctx context.Context) (*Report, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{report.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *ReportQuery) FirstX(ctx context.Context) *Report {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Report ID from the query.
// Returns a *NotFoundError when no Report ID was found.
func (rq *ReportQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{report.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *ReportQuery) FirstIDX(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Report entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Report entity is found.
// Returns a *NotFoundError when no Report entities are found.
func (rq *ReportQuery) Only(ctx context.Context) (*Report, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{report.Label}
	default:
		return nil, &NotSingularError{report.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *ReportQuery) OnlyX(ctx context.Context) *Report {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Report ID in the query.
// Returns a *NotSingularError when more than one Report ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *ReportQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{report.Label}
	default:
		err = &NotSingularError{report.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *ReportQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Reports.
func (rq *ReportQuery) All(ctx context.Context) ([]*Report, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Report, *ReportQuery]()
	return withInterceptors[[]*Report](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *ReportQuery) AllX(ctx context.Context) []*Report {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Report IDs.
func (rq *ReportQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(report.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *ReportQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *ReportQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*ReportQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *ReportQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *ReportQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *ReportQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ReportQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *ReportQuery) Clone() *ReportQuery {
	if rq == nil {
		return nil
	}
	return &ReportQuery{
		config:          rq.config,
		ctx:             rq.ctx.Clone(),
		order:           append([]report.OrderOption{}, rq.order...),
		inters:          append([]Interceptor{}, rq.inters...),
		predicates:      append([]predicate.Report{}, rq.predicates...),
		withEarthquakes: rq.withEarthquakes.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithEarthquakes tells the query-builder to eager-load the nodes that are connected to
// the "earthquakes" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReportQuery) WithEarthquakes(opts ...func(*EarthquakeQuery)) *ReportQuery {
	query := (&EarthquakeClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withEarthquakes = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Felt int32 `json:"felt,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Report.Query().
//		GroupBy(report.FieldFelt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *ReportQuery) GroupBy(field string, fields ...string) *ReportGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ReportGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = report.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Felt int32 `json:"felt,omitempty"`
//	}
//
//	client.Report.Query().
//		Select(report.FieldFelt).
//		Scan(ctx, &v)
func (rq *ReportQuery) Select(fields ...string) *ReportSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &ReportSelect{ReportQuery: rq}
	sbuild.label = report.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ReportSelect configured with the given aggregations.
func (rq *ReportQuery) Aggregate(fns ...AggregateFunc) *ReportSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *ReportQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !report.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *ReportQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Report, error) {
	var (
		nodes       = []*Report{}
		_spec       = rq.querySpec()
		loadedTypes = [1]bool{
			rq.withEarthquakes != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Report).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Report{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withEarthquakes; query != nil {
		if err := rq.loadEarthquakes(ctx, query, nodes,
			func(n *Report) { n.Edges.Earthquakes = []*Earthquake{} },
			func(n *Report, e *Earthquake) { n.Edges.Earthquakes = append(n.Edges.Earthquakes, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *ReportQuery) loadEarthquakes(ctx context.Context, query *EarthquakeQuery, nodes []*Report, init func(*Report), assign func(*Report, *Earthquake)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Report)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(earthquake.FieldReportID)
	}
	query.Where(predicate.Earthquake(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(report.EarthquakesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ReportID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "report_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rq *ReportQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *ReportQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(report.Table, report.Columns, sqlgraph.NewFieldSpec(report.FieldID, field.TypeInt))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, report.FieldID)
		for i := range fields {
			if fields[i] != report.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *ReportQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(report.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = report.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ReportGroupBy is the group-by builder for Report entities.
type ReportGroupBy struct {
	selector
	build *ReportQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *ReportGroupBy) Aggregate(fns ...AggregateFunc) *ReportGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *ReportGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReportQuery, *ReportGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *ReportGroupBy) sqlScan(ctx context.Context, root *ReportQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ReportSelect is the builder for selecting fields of Report entities.
type ReportSelect struct {
	*ReportQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *ReportSelect) Aggregate(fns ...AggregateFunc) *ReportSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *ReportSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReportQuery, *ReportSelect](ctx, rs.ReportQuery, rs, rs.inters, v)
}

func (rs *ReportSelect) sqlScan(ctx context.Context, root *ReportQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
