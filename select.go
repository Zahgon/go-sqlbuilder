// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

import (
	"reflect"

	"github.com/huandu/go-clone"
)

const (
	selectMarkerInit injectionMarker = iota
	selectMarkerAfterWith
	selectMarkerAfterSelect
	selectMarkerAfterFrom
	selectMarkerAfterJoin
	selectMarkerAfterWhere
	selectMarkerAfterGroupBy
	selectMarkerAfterOrderBy
	selectMarkerAfterLimit
	selectMarkerAfterFor
)

// JoinOption is the option in JOIN.
type JoinOption string

// Join options.
const (
	FullJoin       JoinOption = "FULL"
	FullOuterJoin  JoinOption = "FULL OUTER"
	InnerJoin      JoinOption = "INNER"
	LeftJoin       JoinOption = "LEFT"
	LeftOuterJoin  JoinOption = "LEFT OUTER"
	RightJoin      JoinOption = "RIGHT"
	RightOuterJoin JoinOption = "RIGHT OUTER"
)

// NewSelectBuilder creates a new SELECT builder.
func NewSelectBuilder() *SelectBuilder { _ = "STUB: not implemented"; return nil }

func newSelectBuilder() *SelectBuilder { _ = "STUB: not implemented"; return nil }

// Clone returns a deep copy of SelectBuilder.
// It's useful when you want to create a base builder and clone it to build similar queries.
func (sb *SelectBuilder) Clone() *SelectBuilder { _ = "STUB: not implemented"; return nil }

func init() {
	t := reflect.TypeOf(SelectBuilder{})
	clone.SetCustomFunc(t, func(allocator *clone.Allocator, old, new reflect.Value) {
		cloned := allocator.CloneSlowly(old)
		new.Set(cloned)

		sb := cloned.Addr().Interface().(*SelectBuilder)
		sb.args.Replace(sb.whereClauseExpr, sb.whereClauseProxy)
		sb.args.Replace(sb.cteBuilderVar, sb.cteBuilder)
	})
}

// SelectBuilder is a builder to build SELECT.
type SelectBuilder struct {
	*WhereClause
	Cond

	whereClauseProxy *whereClauseProxy
	whereClauseExpr  string

	cteBuilderVar string
	cteBuilder    *CTEBuilder

	distinct    bool
	tables      []string
	selectCols  []string
	joinOptions []JoinOption
	joinTables  []string
	joinExprs   [][]string
	havingExprs []string
	groupByCols []string
	orderByCols []string
	order       string
	limitVar    string
	offsetVar   string
	forWhat     string

	args *Args

	injection *injection
	marker    injectionMarker
}

var _ Builder = new(SelectBuilder)

// Select sets columns in SELECT.
func Select(col ...string) *SelectBuilder { _ = "STUB: not implemented"; return nil }

// TableNames returns all table names in this SELECT statement.
func (sb *SelectBuilder) TableNames() []string { _ = "STUB: not implemented"; return nil }

// With sets WITH clause (the Common Table Expression) before SELECT.
func (sb *SelectBuilder) With(builder *CTEBuilder) *SelectBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Select sets columns in SELECT.
func (sb *SelectBuilder) Select(col ...string) *SelectBuilder {
	_ = "STUB: not implemented"
	return nil
}

// SelectMore adds more columns in SELECT.
func (sb *SelectBuilder) SelectMore(col ...string) *SelectBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Distinct marks this SELECT as DISTINCT.
func (sb *SelectBuilder) Distinct() *SelectBuilder { _ = "STUB: not implemented"; return nil }

// From sets table names in SELECT.
func (sb *SelectBuilder) From(table ...string) *SelectBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Join sets expressions of JOIN in SELECT.
//
// It builds a JOIN expression like
//
//	JOIN table ON onExpr[0] AND onExpr[1] ...
func (sb *SelectBuilder) Join(table string, onExpr ...string) *SelectBuilder {
	_ = "STUB: not implemented"
	return nil
}

// JoinWithOption sets expressions of JOIN with an option.
//
// It builds a JOIN expression like
//
//	option JOIN table ON onExpr[0] AND onExpr[1] ...
//
// Here is a list of supported options.
//   - FullJoin: FULL JOIN
//   - FullOuterJoin: FULL OUTER JOIN
//   - InnerJoin: INNER JOIN
//   - LeftJoin: LEFT JOIN
//   - LeftOuterJoin: LEFT OUTER JOIN
//   - RightJoin: RIGHT JOIN
//   - RightOuterJoin: RIGHT OUTER JOIN
func (sb *SelectBuilder) JoinWithOption(option JoinOption, table string, onExpr ...string) *SelectBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Where adds expressions to the WHERE clause in SELECT.
//
// Multiple calls to Where will join expressions with AND.
// To reset the WHERE clause, set the WhereClause field to nil.
func (sb *SelectBuilder) Where(andExpr ...string) *SelectBuilder {
	_ = "STUB: not implemented"
	return nil
}

// AddWhereClause adds all clauses in the whereClause to SELECT.
func (sb *SelectBuilder) AddWhereClause(whereClause *WhereClause) *SelectBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Having sets expressions of HAVING in SELECT.
func (sb *SelectBuilder) Having(andExpr ...string) *SelectBuilder {
	_ = "STUB: not implemented"
	return nil
}

// GroupBy sets columns of GROUP BY in SELECT.
func (sb *SelectBuilder) GroupBy(col ...string) *SelectBuilder {
	_ = "STUB: not implemented"
	return nil
}

// OrderBy sets columns of ORDER BY in SELECT.
//
// It's recommended to use OrderByAsc or OrderByDesc instead for better support of multiple ORDER BY columns with different directions.
// OrderBy combined with Asc/Desc only supports a single direction for all columns.
func (sb *SelectBuilder) OrderBy(col ...string) *SelectBuilder {
	_ = "STUB: not implemented"
	return nil
}

// OrderByAsc sets a column of ORDER BY in SELECT with ASC order.
// It supports chaining multiple calls to add multiple ORDER BY columns with different directions.
//
//	sb.OrderByAsc("name").OrderByDesc("id")
//	// Generates: ORDER BY name ASC, id DESC
func (sb *SelectBuilder) OrderByAsc(col string) *SelectBuilder {
	_ = "STUB: not implemented"
	return nil
}

// OrderByDesc sets a column of ORDER BY in SELECT with DESC order.
// It supports chaining multiple calls to add multiple ORDER BY columns with different directions.
//
//	sb.OrderByDesc("id").OrderByAsc("name")
//	// Generates: ORDER BY id DESC, name ASC
func (sb *SelectBuilder) OrderByDesc(col string) *SelectBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Asc sets order of ORDER BY to ASC.
//
// Deprecated: Use OrderByAsc instead. Asc only supports a single direction for all ORDER BY columns.
func (sb *SelectBuilder) Asc() *SelectBuilder { _ = "STUB: not implemented"; return nil }

// Desc sets order of ORDER BY to DESC.
//
// Deprecated: Use OrderByDesc instead. Desc only supports a single direction for all ORDER BY columns.
func (sb *SelectBuilder) Desc() *SelectBuilder { _ = "STUB: not implemented"; return nil }

// Limit sets the LIMIT in SELECT.
func (sb *SelectBuilder) Limit(limit int) *SelectBuilder { _ = "STUB: not implemented"; return nil }

// Offset sets the LIMIT offset in SELECT.
func (sb *SelectBuilder) Offset(offset int) *SelectBuilder { _ = "STUB: not implemented"; return nil }

// ForUpdate adds FOR UPDATE at the end of SELECT statement.
func (sb *SelectBuilder) ForUpdate() *SelectBuilder { _ = "STUB: not implemented"; return nil }

// ForShare adds FOR SHARE at the end of SELECT statement.
func (sb *SelectBuilder) ForShare() *SelectBuilder { _ = "STUB: not implemented"; return nil }

// As returns an AS expression.
func (sb *SelectBuilder) As(name, alias string) string { _ = "STUB: not implemented"; return "" }

// BuilderAs returns an AS expression wrapping a complex SQL.
// According to SQL syntax, SQL built by builder is surrounded by parens.
func (sb *SelectBuilder) BuilderAs(builder Builder, alias string) string {
	_ = "STUB: not implemented"
	return ""
}

// LateralAs returns a LATERAL derived table expression wrapping a complex SQL.
func (sb *SelectBuilder) LateralAs(builder Builder, alias string) string {
	_ = "STUB: not implemented"
	return ""
}

// NumCol returns the number of columns to select.
func (sb *SelectBuilder) NumCol() int { _ = "STUB: not implemented"; return 0 }

// String returns the compiled SELECT string.
func (sb *SelectBuilder) String() string { _ = "STUB: not implemented"; return "" }

// Build returns compiled SELECT string and args.
// They can be used in `DB#Query` of package `database/sql` directly.
func (sb *SelectBuilder) Build() (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// BuildWithFlavor returns compiled SELECT string and args with flavor and initial args.
// They can be used in `DB#Query` of package `database/sql` directly.
func (sb *SelectBuilder) BuildWithFlavor(flavor Flavor, initialArg ...interface{}) (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// There might be a hidden constraint in Presto requiring offset to be set before limit.
// The select statement documentation (https://prestodb.io/docs/current/sql/select.html)
// puts offset before limit, and Trino, which is based on Presto, seems
// to require this specific order.

// If ORDER BY is not set, sort column #1 by default.
// It's required to make OFFSET...FETCH work.

// [SKIP N] FIRST M
// M must be greater than 0

// #192: Doris doesn't support ? in OFFSET and LIMIT.

// SetFlavor sets the flavor of compiled sql.
func (sb *SelectBuilder) SetFlavor(flavor Flavor) (old Flavor) {
	_ = "STUB: not implemented"
	return *new(Flavor)
}

// Flavor returns flavor of builder
func (sb *SelectBuilder) Flavor() Flavor {
	_ = "STUB: not implemented"
	return *

	// SQL adds an arbitrary sql to current position.
	new(Flavor)
}

func (sb *SelectBuilder) SQL(sql string) *SelectBuilder { _ = "STUB: not implemented"; return nil }
