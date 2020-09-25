// 🙅🏻🙅🏻🙅🏻 自动生成的代码，尽量不要修改

package model

import (
	"app/model/authuser"
	"app/model/predicate"
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// AuthUserQuery is the builder for querying AuthUser entities.
type AuthUserQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.AuthUser
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (auq *AuthUserQuery) Where(ps ...predicate.AuthUser) *AuthUserQuery {
	auq.predicates = append(auq.predicates, ps...)
	return auq
}

// Limit adds a limit step to the query.
func (auq *AuthUserQuery) Limit(limit int) *AuthUserQuery {
	auq.limit = &limit
	return auq
}

// Offset adds an offset step to the query.
func (auq *AuthUserQuery) Offset(offset int) *AuthUserQuery {
	auq.offset = &offset
	return auq
}

// Order adds an order step to the query.
func (auq *AuthUserQuery) Order(o ...OrderFunc) *AuthUserQuery {
	auq.order = append(auq.order, o...)
	return auq
}

// First returns the first AuthUser entity in the query. Returns *NotFoundError when no authuser was found.
func (auq *AuthUserQuery) First(ctx context.Context) (*AuthUser, error) {
	nodes, err := auq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (auq *AuthUserQuery) FirstX(ctx context.Context) *AuthUser {
	node, err := auq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AuthUser id in the query. Returns *NotFoundError when no id was found.
func (auq *AuthUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = auq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authuser.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (auq *AuthUserQuery) FirstXID(ctx context.Context) int {
	id, err := auq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only AuthUser entity in the query, returns an error if not exactly one entity was returned.
func (auq *AuthUserQuery) Only(ctx context.Context) (*AuthUser, error) {
	nodes, err := auq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authuser.Label}
	default:
		return nil, &NotSingularError{authuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (auq *AuthUserQuery) OnlyX(ctx context.Context) *AuthUser {
	node, err := auq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only AuthUser id in the query, returns an error if not exactly one id was returned.
func (auq *AuthUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = auq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authuser.Label}
	default:
		err = &NotSingularError{authuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (auq *AuthUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := auq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AuthUsers.
func (auq *AuthUserQuery) All(ctx context.Context) ([]*AuthUser, error) {
	if err := auq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return auq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (auq *AuthUserQuery) AllX(ctx context.Context) []*AuthUser {
	nodes, err := auq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AuthUser ids.
func (auq *AuthUserQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := auq.Select(authuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (auq *AuthUserQuery) IDsX(ctx context.Context) []int {
	ids, err := auq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (auq *AuthUserQuery) Count(ctx context.Context) (int, error) {
	if err := auq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return auq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (auq *AuthUserQuery) CountX(ctx context.Context) int {
	count, err := auq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (auq *AuthUserQuery) Exist(ctx context.Context) (bool, error) {
	if err := auq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return auq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (auq *AuthUserQuery) ExistX(ctx context.Context) bool {
	exist, err := auq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (auq *AuthUserQuery) Clone() *AuthUserQuery {
	return &AuthUserQuery{
		config:     auq.config,
		limit:      auq.limit,
		offset:     auq.offset,
		order:      append([]OrderFunc{}, auq.order...),
		unique:     append([]string{}, auq.unique...),
		predicates: append([]predicate.AuthUser{}, auq.predicates...),
		// clone intermediate query.
		sql:  auq.sql.Clone(),
		path: auq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AuthUser.Query().
//		GroupBy(authuser.FieldUsername).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
//
func (auq *AuthUserQuery) GroupBy(field string, fields ...string) *AuthUserGroupBy {
	group := &AuthUserGroupBy{config: auq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := auq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return auq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//	}
//
//	client.AuthUser.Query().
//		Select(authuser.FieldUsername).
//		Scan(ctx, &v)
//
func (auq *AuthUserQuery) Select(field string, fields ...string) *AuthUserSelect {
	selector := &AuthUserSelect{config: auq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := auq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return auq.sqlQuery(), nil
	}
	return selector
}

func (auq *AuthUserQuery) prepareQuery(ctx context.Context) error {
	if auq.path != nil {
		prev, err := auq.path(ctx)
		if err != nil {
			return err
		}
		auq.sql = prev
	}
	return nil
}

func (auq *AuthUserQuery) sqlAll(ctx context.Context) ([]*AuthUser, error) {
	var (
		nodes = []*AuthUser{}
		_spec = auq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &AuthUser{config: auq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("model: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, auq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (auq *AuthUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := auq.querySpec()
	return sqlgraph.CountNodes(ctx, auq.driver, _spec)
}

func (auq *AuthUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := auq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("model: check existence: %v", err)
	}
	return n > 0, nil
}

func (auq *AuthUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authuser.Table,
			Columns: authuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authuser.FieldID,
			},
		},
		From:   auq.sql,
		Unique: true,
	}
	if ps := auq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := auq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := auq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := auq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, authuser.ValidColumn)
			}
		}
	}
	return _spec
}

func (auq *AuthUserQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(auq.driver.Dialect())
	t1 := builder.Table(authuser.Table)
	selector := builder.Select(t1.Columns(authuser.Columns...)...).From(t1)
	if auq.sql != nil {
		selector = auq.sql
		selector.Select(selector.Columns(authuser.Columns...)...)
	}
	for _, p := range auq.predicates {
		p(selector)
	}
	for _, p := range auq.order {
		p(selector, authuser.ValidColumn)
	}
	if offset := auq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := auq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuthUserGroupBy is the builder for group-by AuthUser entities.
type AuthUserGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (augb *AuthUserGroupBy) Aggregate(fns ...AggregateFunc) *AuthUserGroupBy {
	augb.fns = append(augb.fns, fns...)
	return augb
}

// Scan applies the group-by query and scan the result into the given value.
func (augb *AuthUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := augb.path(ctx)
	if err != nil {
		return err
	}
	augb.sql = query
	return augb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (augb *AuthUserGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := augb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (augb *AuthUserGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(augb.fields) > 1 {
		return nil, errors.New("model: AuthUserGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := augb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (augb *AuthUserGroupBy) StringsX(ctx context.Context) []string {
	v, err := augb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (augb *AuthUserGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = augb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authuser.Label}
	default:
		err = fmt.Errorf("model: AuthUserGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (augb *AuthUserGroupBy) StringX(ctx context.Context) string {
	v, err := augb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (augb *AuthUserGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(augb.fields) > 1 {
		return nil, errors.New("model: AuthUserGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := augb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (augb *AuthUserGroupBy) IntsX(ctx context.Context) []int {
	v, err := augb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (augb *AuthUserGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = augb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authuser.Label}
	default:
		err = fmt.Errorf("model: AuthUserGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (augb *AuthUserGroupBy) IntX(ctx context.Context) int {
	v, err := augb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (augb *AuthUserGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(augb.fields) > 1 {
		return nil, errors.New("model: AuthUserGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := augb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (augb *AuthUserGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := augb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (augb *AuthUserGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = augb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authuser.Label}
	default:
		err = fmt.Errorf("model: AuthUserGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (augb *AuthUserGroupBy) Float64X(ctx context.Context) float64 {
	v, err := augb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (augb *AuthUserGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(augb.fields) > 1 {
		return nil, errors.New("model: AuthUserGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := augb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (augb *AuthUserGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := augb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (augb *AuthUserGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = augb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authuser.Label}
	default:
		err = fmt.Errorf("model: AuthUserGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (augb *AuthUserGroupBy) BoolX(ctx context.Context) bool {
	v, err := augb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (augb *AuthUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range augb.fields {
		if !authuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := augb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := augb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (augb *AuthUserGroupBy) sqlQuery() *sql.Selector {
	selector := augb.sql
	columns := make([]string, 0, len(augb.fields)+len(augb.fns))
	columns = append(columns, augb.fields...)
	for _, fn := range augb.fns {
		columns = append(columns, fn(selector, authuser.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(augb.fields...)
}

// AuthUserSelect is the builder for select fields of AuthUser entities.
type AuthUserSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (aus *AuthUserSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := aus.path(ctx)
	if err != nil {
		return err
	}
	aus.sql = query
	return aus.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aus *AuthUserSelect) ScanX(ctx context.Context, v interface{}) {
	if err := aus.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (aus *AuthUserSelect) Strings(ctx context.Context) ([]string, error) {
	if len(aus.fields) > 1 {
		return nil, errors.New("model: AuthUserSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := aus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aus *AuthUserSelect) StringsX(ctx context.Context) []string {
	v, err := aus.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (aus *AuthUserSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aus.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authuser.Label}
	default:
		err = fmt.Errorf("model: AuthUserSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aus *AuthUserSelect) StringX(ctx context.Context) string {
	v, err := aus.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (aus *AuthUserSelect) Ints(ctx context.Context) ([]int, error) {
	if len(aus.fields) > 1 {
		return nil, errors.New("model: AuthUserSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := aus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aus *AuthUserSelect) IntsX(ctx context.Context) []int {
	v, err := aus.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (aus *AuthUserSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aus.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authuser.Label}
	default:
		err = fmt.Errorf("model: AuthUserSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aus *AuthUserSelect) IntX(ctx context.Context) int {
	v, err := aus.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (aus *AuthUserSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(aus.fields) > 1 {
		return nil, errors.New("model: AuthUserSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := aus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aus *AuthUserSelect) Float64sX(ctx context.Context) []float64 {
	v, err := aus.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (aus *AuthUserSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aus.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authuser.Label}
	default:
		err = fmt.Errorf("model: AuthUserSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aus *AuthUserSelect) Float64X(ctx context.Context) float64 {
	v, err := aus.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (aus *AuthUserSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(aus.fields) > 1 {
		return nil, errors.New("model: AuthUserSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := aus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aus *AuthUserSelect) BoolsX(ctx context.Context) []bool {
	v, err := aus.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (aus *AuthUserSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aus.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authuser.Label}
	default:
		err = fmt.Errorf("model: AuthUserSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aus *AuthUserSelect) BoolX(ctx context.Context) bool {
	v, err := aus.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aus *AuthUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range aus.fields {
		if !authuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := aus.sqlQuery().Query()
	if err := aus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (aus *AuthUserSelect) sqlQuery() sql.Querier {
	selector := aus.sql
	selector.Select(selector.Columns(aus.fields...)...)
	return selector
}
