// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"entdemo/ent/earthquake"
	"entdemo/ent/geometry"
	"entdemo/ent/location"
	"entdemo/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GeometryQuery is the builder for querying Geometry entities.
type GeometryQuery struct {
	config
	ctx             *QueryContext
	order           []geometry.OrderOption
	inters          []Interceptor
	predicates      []predicate.Geometry
	withEarthquakes *EarthquakeQuery
	withLocation    *LocationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GeometryQuery builder.
func (gq *GeometryQuery) Where(ps ...predicate.Geometry) *GeometryQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit the number of records to be returned by this query.
func (gq *GeometryQuery) Limit(limit int) *GeometryQuery {
	gq.ctx.Limit = &limit
	return gq
}

// Offset to start from.
func (gq *GeometryQuery) Offset(offset int) *GeometryQuery {
	gq.ctx.Offset = &offset
	return gq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gq *GeometryQuery) Unique(unique bool) *GeometryQuery {
	gq.ctx.Unique = &unique
	return gq
}

// Order specifies how the records should be ordered.
func (gq *GeometryQuery) Order(o ...geometry.OrderOption) *GeometryQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// QueryEarthquakes chains the current query on the "earthquakes" edge.
func (gq *GeometryQuery) QueryEarthquakes() *EarthquakeQuery {
	query := (&EarthquakeClient{config: gq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(geometry.Table, geometry.FieldID, selector),
			sqlgraph.To(earthquake.Table, earthquake.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, geometry.EarthquakesTable, geometry.EarthquakesColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLocation chains the current query on the "location" edge.
func (gq *GeometryQuery) QueryLocation() *LocationQuery {
	query := (&LocationClient{config: gq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(geometry.Table, geometry.FieldID, selector),
			sqlgraph.To(location.Table, location.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, geometry.LocationTable, geometry.LocationColumn),
		)
		fromU = sqlgraph.SetNeighbors(gq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Geometry entity from the query.
// Returns a *NotFoundError when no Geometry was found.
func (gq *GeometryQuery) First(ctx context.Context) (*Geometry, error) {
	nodes, err := gq.Limit(1).All(setContextOp(ctx, gq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{geometry.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GeometryQuery) FirstX(ctx context.Context) *Geometry {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Geometry ID from the query.
// Returns a *NotFoundError when no Geometry ID was found.
func (gq *GeometryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(1).IDs(setContextOp(ctx, gq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{geometry.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GeometryQuery) FirstIDX(ctx context.Context) int {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Geometry entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Geometry entity is found.
// Returns a *NotFoundError when no Geometry entities are found.
func (gq *GeometryQuery) Only(ctx context.Context) (*Geometry, error) {
	nodes, err := gq.Limit(2).All(setContextOp(ctx, gq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{geometry.Label}
	default:
		return nil, &NotSingularError{geometry.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GeometryQuery) OnlyX(ctx context.Context) *Geometry {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Geometry ID in the query.
// Returns a *NotSingularError when more than one Geometry ID is found.
// Returns a *NotFoundError when no entities are found.
func (gq *GeometryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(2).IDs(setContextOp(ctx, gq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{geometry.Label}
	default:
		err = &NotSingularError{geometry.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GeometryQuery) OnlyIDX(ctx context.Context) int {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Geometries.
func (gq *GeometryQuery) All(ctx context.Context) ([]*Geometry, error) {
	ctx = setContextOp(ctx, gq.ctx, "All")
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Geometry, *GeometryQuery]()
	return withInterceptors[[]*Geometry](ctx, gq, qr, gq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gq *GeometryQuery) AllX(ctx context.Context) []*Geometry {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Geometry IDs.
func (gq *GeometryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if gq.ctx.Unique == nil && gq.path != nil {
		gq.Unique(true)
	}
	ctx = setContextOp(ctx, gq.ctx, "IDs")
	if err = gq.Select(geometry.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GeometryQuery) IDsX(ctx context.Context) []int {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GeometryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gq.ctx, "Count")
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gq, querierCount[*GeometryQuery](), gq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GeometryQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GeometryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gq.ctx, "Exist")
	switch _, err := gq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GeometryQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GeometryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GeometryQuery) Clone() *GeometryQuery {
	if gq == nil {
		return nil
	}
	return &GeometryQuery{
		config:          gq.config,
		ctx:             gq.ctx.Clone(),
		order:           append([]geometry.OrderOption{}, gq.order...),
		inters:          append([]Interceptor{}, gq.inters...),
		predicates:      append([]predicate.Geometry{}, gq.predicates...),
		withEarthquakes: gq.withEarthquakes.Clone(),
		withLocation:    gq.withLocation.Clone(),
		// clone intermediate query.
		sql:  gq.sql.Clone(),
		path: gq.path,
	}
}

// WithEarthquakes tells the query-builder to eager-load the nodes that are connected to
// the "earthquakes" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GeometryQuery) WithEarthquakes(opts ...func(*EarthquakeQuery)) *GeometryQuery {
	query := (&EarthquakeClient{config: gq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gq.withEarthquakes = query
	return gq
}

// WithLocation tells the query-builder to eager-load the nodes that are connected to
// the "location" edge. The optional arguments are used to configure the query builder of the edge.
func (gq *GeometryQuery) WithLocation(opts ...func(*LocationQuery)) *GeometryQuery {
	query := (&LocationClient{config: gq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	gq.withLocation = query
	return gq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		LocationID int `json:"location_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Geometry.Query().
//		GroupBy(geometry.FieldLocationID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (gq *GeometryQuery) GroupBy(field string, fields ...string) *GeometryGroupBy {
	gq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GeometryGroupBy{build: gq}
	grbuild.flds = &gq.ctx.Fields
	grbuild.label = geometry.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		LocationID int `json:"location_id,omitempty"`
//	}
//
//	client.Geometry.Query().
//		Select(geometry.FieldLocationID).
//		Scan(ctx, &v)
func (gq *GeometryQuery) Select(fields ...string) *GeometrySelect {
	gq.ctx.Fields = append(gq.ctx.Fields, fields...)
	sbuild := &GeometrySelect{GeometryQuery: gq}
	sbuild.label = geometry.Label
	sbuild.flds, sbuild.scan = &gq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GeometrySelect configured with the given aggregations.
func (gq *GeometryQuery) Aggregate(fns ...AggregateFunc) *GeometrySelect {
	return gq.Select().Aggregate(fns...)
}

func (gq *GeometryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gq); err != nil {
				return err
			}
		}
	}
	for _, f := range gq.ctx.Fields {
		if !geometry.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.sql = prev
	}
	return nil
}

func (gq *GeometryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Geometry, error) {
	var (
		nodes       = []*Geometry{}
		_spec       = gq.querySpec()
		loadedTypes = [2]bool{
			gq.withEarthquakes != nil,
			gq.withLocation != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Geometry).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Geometry{config: gq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := gq.withEarthquakes; query != nil {
		if err := gq.loadEarthquakes(ctx, query, nodes,
			func(n *Geometry) { n.Edges.Earthquakes = []*Earthquake{} },
			func(n *Geometry, e *Earthquake) { n.Edges.Earthquakes = append(n.Edges.Earthquakes, e) }); err != nil {
			return nil, err
		}
	}
	if query := gq.withLocation; query != nil {
		if err := gq.loadLocation(ctx, query, nodes, nil,
			func(n *Geometry, e *Location) { n.Edges.Location = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (gq *GeometryQuery) loadEarthquakes(ctx context.Context, query *EarthquakeQuery, nodes []*Geometry, init func(*Geometry), assign func(*Geometry, *Earthquake)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Geometry)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(earthquake.FieldGeoID)
	}
	query.Where(predicate.Earthquake(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(geometry.EarthquakesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.GeoID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "geo_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (gq *GeometryQuery) loadLocation(ctx context.Context, query *LocationQuery, nodes []*Geometry, init func(*Geometry), assign func(*Geometry, *Location)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Geometry)
	for i := range nodes {
		fk := nodes[i].LocationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(location.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "location_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (gq *GeometryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	_spec.Node.Columns = gq.ctx.Fields
	if len(gq.ctx.Fields) > 0 {
		_spec.Unique = gq.ctx.Unique != nil && *gq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GeometryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(geometry.Table, geometry.Columns, sqlgraph.NewFieldSpec(geometry.FieldID, field.TypeInt))
	_spec.From = gq.sql
	if unique := gq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gq.path != nil {
		_spec.Unique = true
	}
	if fields := gq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, geometry.FieldID)
		for i := range fields {
			if fields[i] != geometry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if gq.withLocation != nil {
			_spec.Node.AddColumnOnce(geometry.FieldLocationID)
		}
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gq *GeometryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(geometry.Table)
	columns := gq.ctx.Fields
	if len(columns) == 0 {
		columns = geometry.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gq.ctx.Unique != nil && *gq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector)
	}
	if offset := gq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GeometryGroupBy is the group-by builder for Geometry entities.
type GeometryGroupBy struct {
	selector
	build *GeometryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GeometryGroupBy) Aggregate(fns ...AggregateFunc) *GeometryGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the selector query and scans the result into the given value.
func (ggb *GeometryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ggb.build.ctx, "GroupBy")
	if err := ggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GeometryQuery, *GeometryGroupBy](ctx, ggb.build, ggb, ggb.build.inters, v)
}

func (ggb *GeometryGroupBy) sqlScan(ctx context.Context, root *GeometryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ggb.fns))
	for _, fn := range ggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ggb.flds)+len(ggb.fns))
		for _, f := range *ggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GeometrySelect is the builder for selecting fields of Geometry entities.
type GeometrySelect struct {
	*GeometryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gs *GeometrySelect) Aggregate(fns ...AggregateFunc) *GeometrySelect {
	gs.fns = append(gs.fns, fns...)
	return gs
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GeometrySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gs.ctx, "Select")
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GeometryQuery, *GeometrySelect](ctx, gs.GeometryQuery, gs, gs.inters, v)
}

func (gs *GeometrySelect) sqlScan(ctx context.Context, root *GeometryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gs.fns))
	for _, fn := range gs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
