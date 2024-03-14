// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"entdemo/ent/associatedevent"
	"entdemo/ent/earthquake"
	"entdemo/ent/eventtype"
	"entdemo/ent/featuretype"
	"entdemo/ent/feltreport"
	"entdemo/ent/geometry"
	"entdemo/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EarthquakeQuery is the builder for querying Earthquake entities.
type EarthquakeQuery struct {
	config
	ctx                  *QueryContext
	order                []earthquake.OrderOption
	inters               []Interceptor
	predicates           []predicate.Earthquake
	withMainEvents       *AssociatedEventQuery
	withAssociatedEvents *AssociatedEventQuery
	withEventTypes       *EventTypeQuery
	withFeatureTypes     *FeatureTypeQuery
	withFeltReports      *FeltReportQuery
	withGeometries       *GeometryQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EarthquakeQuery builder.
func (eq *EarthquakeQuery) Where(ps ...predicate.Earthquake) *EarthquakeQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit the number of records to be returned by this query.
func (eq *EarthquakeQuery) Limit(limit int) *EarthquakeQuery {
	eq.ctx.Limit = &limit
	return eq
}

// Offset to start from.
func (eq *EarthquakeQuery) Offset(offset int) *EarthquakeQuery {
	eq.ctx.Offset = &offset
	return eq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eq *EarthquakeQuery) Unique(unique bool) *EarthquakeQuery {
	eq.ctx.Unique = &unique
	return eq
}

// Order specifies how the records should be ordered.
func (eq *EarthquakeQuery) Order(o ...earthquake.OrderOption) *EarthquakeQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryMainEvents chains the current query on the "main_events" edge.
func (eq *EarthquakeQuery) QueryMainEvents() *AssociatedEventQuery {
	query := (&AssociatedEventClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(earthquake.Table, earthquake.FieldID, selector),
			sqlgraph.To(associatedevent.Table, associatedevent.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, earthquake.MainEventsTable, earthquake.MainEventsColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAssociatedEvents chains the current query on the "associated_events" edge.
func (eq *EarthquakeQuery) QueryAssociatedEvents() *AssociatedEventQuery {
	query := (&AssociatedEventClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(earthquake.Table, earthquake.FieldID, selector),
			sqlgraph.To(associatedevent.Table, associatedevent.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, earthquake.AssociatedEventsTable, earthquake.AssociatedEventsColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEventTypes chains the current query on the "event_types" edge.
func (eq *EarthquakeQuery) QueryEventTypes() *EventTypeQuery {
	query := (&EventTypeClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(earthquake.Table, earthquake.FieldID, selector),
			sqlgraph.To(eventtype.Table, eventtype.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, earthquake.EventTypesTable, earthquake.EventTypesColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFeatureTypes chains the current query on the "feature_types" edge.
func (eq *EarthquakeQuery) QueryFeatureTypes() *FeatureTypeQuery {
	query := (&FeatureTypeClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(earthquake.Table, earthquake.FieldID, selector),
			sqlgraph.To(featuretype.Table, featuretype.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, earthquake.FeatureTypesTable, earthquake.FeatureTypesColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFeltReports chains the current query on the "felt_reports" edge.
func (eq *EarthquakeQuery) QueryFeltReports() *FeltReportQuery {
	query := (&FeltReportClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(earthquake.Table, earthquake.FieldID, selector),
			sqlgraph.To(feltreport.Table, feltreport.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, earthquake.FeltReportsTable, earthquake.FeltReportsColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGeometries chains the current query on the "geometries" edge.
func (eq *EarthquakeQuery) QueryGeometries() *GeometryQuery {
	query := (&GeometryClient{config: eq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(earthquake.Table, earthquake.FieldID, selector),
			sqlgraph.To(geometry.Table, geometry.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, earthquake.GeometriesTable, earthquake.GeometriesColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Earthquake entity from the query.
// Returns a *NotFoundError when no Earthquake was found.
func (eq *EarthquakeQuery) First(ctx context.Context) (*Earthquake, error) {
	nodes, err := eq.Limit(1).All(setContextOp(ctx, eq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{earthquake.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EarthquakeQuery) FirstX(ctx context.Context) *Earthquake {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Earthquake ID from the query.
// Returns a *NotFoundError when no Earthquake ID was found.
func (eq *EarthquakeQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = eq.Limit(1).IDs(setContextOp(ctx, eq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{earthquake.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *EarthquakeQuery) FirstIDX(ctx context.Context) int32 {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Earthquake entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Earthquake entity is found.
// Returns a *NotFoundError when no Earthquake entities are found.
func (eq *EarthquakeQuery) Only(ctx context.Context) (*Earthquake, error) {
	nodes, err := eq.Limit(2).All(setContextOp(ctx, eq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{earthquake.Label}
	default:
		return nil, &NotSingularError{earthquake.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EarthquakeQuery) OnlyX(ctx context.Context) *Earthquake {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Earthquake ID in the query.
// Returns a *NotSingularError when more than one Earthquake ID is found.
// Returns a *NotFoundError when no entities are found.
func (eq *EarthquakeQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = eq.Limit(2).IDs(setContextOp(ctx, eq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{earthquake.Label}
	default:
		err = &NotSingularError{earthquake.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EarthquakeQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Earthquakes.
func (eq *EarthquakeQuery) All(ctx context.Context) ([]*Earthquake, error) {
	ctx = setContextOp(ctx, eq.ctx, "All")
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Earthquake, *EarthquakeQuery]()
	return withInterceptors[[]*Earthquake](ctx, eq, qr, eq.inters)
}

// AllX is like All, but panics if an error occurs.
func (eq *EarthquakeQuery) AllX(ctx context.Context) []*Earthquake {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Earthquake IDs.
func (eq *EarthquakeQuery) IDs(ctx context.Context) (ids []int32, err error) {
	if eq.ctx.Unique == nil && eq.path != nil {
		eq.Unique(true)
	}
	ctx = setContextOp(ctx, eq.ctx, "IDs")
	if err = eq.Select(earthquake.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EarthquakeQuery) IDsX(ctx context.Context) []int32 {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EarthquakeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, eq.ctx, "Count")
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, eq, querierCount[*EarthquakeQuery](), eq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EarthquakeQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EarthquakeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, eq.ctx, "Exist")
	switch _, err := eq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EarthquakeQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EarthquakeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EarthquakeQuery) Clone() *EarthquakeQuery {
	if eq == nil {
		return nil
	}
	return &EarthquakeQuery{
		config:               eq.config,
		ctx:                  eq.ctx.Clone(),
		order:                append([]earthquake.OrderOption{}, eq.order...),
		inters:               append([]Interceptor{}, eq.inters...),
		predicates:           append([]predicate.Earthquake{}, eq.predicates...),
		withMainEvents:       eq.withMainEvents.Clone(),
		withAssociatedEvents: eq.withAssociatedEvents.Clone(),
		withEventTypes:       eq.withEventTypes.Clone(),
		withFeatureTypes:     eq.withFeatureTypes.Clone(),
		withFeltReports:      eq.withFeltReports.Clone(),
		withGeometries:       eq.withGeometries.Clone(),
		// clone intermediate query.
		sql:  eq.sql.Clone(),
		path: eq.path,
	}
}

// WithMainEvents tells the query-builder to eager-load the nodes that are connected to
// the "main_events" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EarthquakeQuery) WithMainEvents(opts ...func(*AssociatedEventQuery)) *EarthquakeQuery {
	query := (&AssociatedEventClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withMainEvents = query
	return eq
}

// WithAssociatedEvents tells the query-builder to eager-load the nodes that are connected to
// the "associated_events" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EarthquakeQuery) WithAssociatedEvents(opts ...func(*AssociatedEventQuery)) *EarthquakeQuery {
	query := (&AssociatedEventClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withAssociatedEvents = query
	return eq
}

// WithEventTypes tells the query-builder to eager-load the nodes that are connected to
// the "event_types" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EarthquakeQuery) WithEventTypes(opts ...func(*EventTypeQuery)) *EarthquakeQuery {
	query := (&EventTypeClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withEventTypes = query
	return eq
}

// WithFeatureTypes tells the query-builder to eager-load the nodes that are connected to
// the "feature_types" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EarthquakeQuery) WithFeatureTypes(opts ...func(*FeatureTypeQuery)) *EarthquakeQuery {
	query := (&FeatureTypeClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withFeatureTypes = query
	return eq
}

// WithFeltReports tells the query-builder to eager-load the nodes that are connected to
// the "felt_reports" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EarthquakeQuery) WithFeltReports(opts ...func(*FeltReportQuery)) *EarthquakeQuery {
	query := (&FeltReportClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withFeltReports = query
	return eq
}

// WithGeometries tells the query-builder to eager-load the nodes that are connected to
// the "geometries" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EarthquakeQuery) WithGeometries(opts ...func(*GeometryQuery)) *EarthquakeQuery {
	query := (&GeometryClient{config: eq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	eq.withGeometries = query
	return eq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FeatureID string `json:"feature_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Earthquake.Query().
//		GroupBy(earthquake.FieldFeatureID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (eq *EarthquakeQuery) GroupBy(field string, fields ...string) *EarthquakeGroupBy {
	eq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EarthquakeGroupBy{build: eq}
	grbuild.flds = &eq.ctx.Fields
	grbuild.label = earthquake.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FeatureID string `json:"feature_id,omitempty"`
//	}
//
//	client.Earthquake.Query().
//		Select(earthquake.FieldFeatureID).
//		Scan(ctx, &v)
func (eq *EarthquakeQuery) Select(fields ...string) *EarthquakeSelect {
	eq.ctx.Fields = append(eq.ctx.Fields, fields...)
	sbuild := &EarthquakeSelect{EarthquakeQuery: eq}
	sbuild.label = earthquake.Label
	sbuild.flds, sbuild.scan = &eq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EarthquakeSelect configured with the given aggregations.
func (eq *EarthquakeQuery) Aggregate(fns ...AggregateFunc) *EarthquakeSelect {
	return eq.Select().Aggregate(fns...)
}

func (eq *EarthquakeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range eq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, eq); err != nil {
				return err
			}
		}
	}
	for _, f := range eq.ctx.Fields {
		if !earthquake.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EarthquakeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Earthquake, error) {
	var (
		nodes       = []*Earthquake{}
		_spec       = eq.querySpec()
		loadedTypes = [6]bool{
			eq.withMainEvents != nil,
			eq.withAssociatedEvents != nil,
			eq.withEventTypes != nil,
			eq.withFeatureTypes != nil,
			eq.withFeltReports != nil,
			eq.withGeometries != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Earthquake).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Earthquake{config: eq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eq.withMainEvents; query != nil {
		if err := eq.loadMainEvents(ctx, query, nodes,
			func(n *Earthquake) { n.Edges.MainEvents = []*AssociatedEvent{} },
			func(n *Earthquake, e *AssociatedEvent) { n.Edges.MainEvents = append(n.Edges.MainEvents, e) }); err != nil {
			return nil, err
		}
	}
	if query := eq.withAssociatedEvents; query != nil {
		if err := eq.loadAssociatedEvents(ctx, query, nodes,
			func(n *Earthquake) { n.Edges.AssociatedEvents = []*AssociatedEvent{} },
			func(n *Earthquake, e *AssociatedEvent) {
				n.Edges.AssociatedEvents = append(n.Edges.AssociatedEvents, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := eq.withEventTypes; query != nil {
		if err := eq.loadEventTypes(ctx, query, nodes,
			func(n *Earthquake) { n.Edges.EventTypes = []*EventType{} },
			func(n *Earthquake, e *EventType) { n.Edges.EventTypes = append(n.Edges.EventTypes, e) }); err != nil {
			return nil, err
		}
	}
	if query := eq.withFeatureTypes; query != nil {
		if err := eq.loadFeatureTypes(ctx, query, nodes,
			func(n *Earthquake) { n.Edges.FeatureTypes = []*FeatureType{} },
			func(n *Earthquake, e *FeatureType) { n.Edges.FeatureTypes = append(n.Edges.FeatureTypes, e) }); err != nil {
			return nil, err
		}
	}
	if query := eq.withFeltReports; query != nil {
		if err := eq.loadFeltReports(ctx, query, nodes,
			func(n *Earthquake) { n.Edges.FeltReports = []*FeltReport{} },
			func(n *Earthquake, e *FeltReport) { n.Edges.FeltReports = append(n.Edges.FeltReports, e) }); err != nil {
			return nil, err
		}
	}
	if query := eq.withGeometries; query != nil {
		if err := eq.loadGeometries(ctx, query, nodes,
			func(n *Earthquake) { n.Edges.Geometries = []*Geometry{} },
			func(n *Earthquake, e *Geometry) { n.Edges.Geometries = append(n.Edges.Geometries, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eq *EarthquakeQuery) loadMainEvents(ctx context.Context, query *AssociatedEventQuery, nodes []*Earthquake, init func(*Earthquake), assign func(*Earthquake, *AssociatedEvent)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Earthquake)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(associatedevent.FieldEarthquakeID)
	}
	query.Where(predicate.AssociatedEvent(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(earthquake.MainEventsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.EarthquakeID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "earthquake_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eq *EarthquakeQuery) loadAssociatedEvents(ctx context.Context, query *AssociatedEventQuery, nodes []*Earthquake, init func(*Earthquake), assign func(*Earthquake, *AssociatedEvent)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Earthquake)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(associatedevent.FieldAssociateID)
	}
	query.Where(predicate.AssociatedEvent(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(earthquake.AssociatedEventsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AssociateID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "associate_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eq *EarthquakeQuery) loadEventTypes(ctx context.Context, query *EventTypeQuery, nodes []*Earthquake, init func(*Earthquake), assign func(*Earthquake, *EventType)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Earthquake)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(eventtype.FieldEarthquakeID)
	}
	query.Where(predicate.EventType(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(earthquake.EventTypesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.EarthquakeID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "earthquake_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eq *EarthquakeQuery) loadFeatureTypes(ctx context.Context, query *FeatureTypeQuery, nodes []*Earthquake, init func(*Earthquake), assign func(*Earthquake, *FeatureType)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Earthquake)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(featuretype.FieldEarthquakeID)
	}
	query.Where(predicate.FeatureType(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(earthquake.FeatureTypesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.EarthquakeID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "earthquake_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eq *EarthquakeQuery) loadFeltReports(ctx context.Context, query *FeltReportQuery, nodes []*Earthquake, init func(*Earthquake), assign func(*Earthquake, *FeltReport)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Earthquake)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(feltreport.FieldEarthquakeID)
	}
	query.Where(predicate.FeltReport(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(earthquake.FeltReportsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.EarthquakeID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "earthquake_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (eq *EarthquakeQuery) loadGeometries(ctx context.Context, query *GeometryQuery, nodes []*Earthquake, init func(*Earthquake), assign func(*Earthquake, *Geometry)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int32]*Earthquake)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(geometry.FieldEarthquakeID)
	}
	query.Where(predicate.Geometry(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(earthquake.GeometriesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.EarthquakeID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "earthquake_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (eq *EarthquakeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	_spec.Node.Columns = eq.ctx.Fields
	if len(eq.ctx.Fields) > 0 {
		_spec.Unique = eq.ctx.Unique != nil && *eq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EarthquakeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(earthquake.Table, earthquake.Columns, sqlgraph.NewFieldSpec(earthquake.FieldID, field.TypeInt32))
	_spec.From = eq.sql
	if unique := eq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if eq.path != nil {
		_spec.Unique = true
	}
	if fields := eq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, earthquake.FieldID)
		for i := range fields {
			if fields[i] != earthquake.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *EarthquakeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(earthquake.Table)
	columns := eq.ctx.Fields
	if len(columns) == 0 {
		columns = earthquake.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eq.ctx.Unique != nil && *eq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EarthquakeGroupBy is the group-by builder for Earthquake entities.
type EarthquakeGroupBy struct {
	selector
	build *EarthquakeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EarthquakeGroupBy) Aggregate(fns ...AggregateFunc) *EarthquakeGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the selector query and scans the result into the given value.
func (egb *EarthquakeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, egb.build.ctx, "GroupBy")
	if err := egb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EarthquakeQuery, *EarthquakeGroupBy](ctx, egb.build, egb, egb.build.inters, v)
}

func (egb *EarthquakeGroupBy) sqlScan(ctx context.Context, root *EarthquakeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(egb.fns))
	for _, fn := range egb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*egb.flds)+len(egb.fns))
		for _, f := range *egb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*egb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EarthquakeSelect is the builder for selecting fields of Earthquake entities.
type EarthquakeSelect struct {
	*EarthquakeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (es *EarthquakeSelect) Aggregate(fns ...AggregateFunc) *EarthquakeSelect {
	es.fns = append(es.fns, fns...)
	return es
}

// Scan applies the selector query and scans the result into the given value.
func (es *EarthquakeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, es.ctx, "Select")
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EarthquakeQuery, *EarthquakeSelect](ctx, es.EarthquakeQuery, es, es.inters, v)
}

func (es *EarthquakeSelect) sqlScan(ctx context.Context, root *EarthquakeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(es.fns))
	for _, fn := range es.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*es.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}