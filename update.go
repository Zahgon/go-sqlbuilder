// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

import (
	"reflect"

	"github.com/huandu/go-clone"
)

const (
	updateMarkerInit injectionMarker = iota
	updateMarkerAfterWith
	updateMarkerAfterUpdate
	updateMarkerAfterSet
	updateMarkerAfterFrom
	updateMarkerAfterWhere
	updateMarkerAfterOrderBy
	updateMarkerAfterLimit
	updateMarkerAfterReturning
)

// NewUpdateBuilder creates a new UPDATE builder.
func NewUpdateBuilder() *UpdateBuilder { _ = "STUB: not implemented"; return nil }

func newUpdateBuilder() *UpdateBuilder { _ = "STUB: not implemented"; return nil }

// Clone returns a deep copy of UpdateBuilder.
// It's useful when you want to create a base builder and clone it to build similar queries.
func (ub *UpdateBuilder) Clone() *UpdateBuilder { _ = "STUB: not implemented"; return nil }

func init() {
	t := reflect.TypeOf(UpdateBuilder{})
	clone.SetCustomFunc(t, func(allocator *clone.Allocator, old, new reflect.Value) {
		cloned := allocator.CloneSlowly(old)
		new.Set(cloned)

		ub := cloned.Addr().Interface().(*UpdateBuilder)
		ub.args.Replace(ub.whereClauseExpr, ub.whereClauseProxy)
		ub.args.Replace(ub.cteBuilderVar, ub.cteBuilder)
	})
}

// UpdateBuilder is a builder to build UPDATE.
type UpdateBuilder struct {
	*WhereClause
	Cond

	whereClauseProxy *whereClauseProxy
	whereClauseExpr  string

	cteBuilderVar string
	cteBuilder    *CTEBuilder

	tables      []string
	fromTables  []string
	assignments []string
	orderByCols []string
	order       string
	limitVar    string
	returning   []string

	args *Args

	injection *injection
	marker    injectionMarker
}

var _ Builder = new(UpdateBuilder)

// Update sets table name in UPDATE.
func Update(table ...string) *UpdateBuilder { _ = "STUB: not implemented"; return nil }

// With sets WITH clause (the Common Table Expression) before UPDATE.
func (ub *UpdateBuilder) With(builder *CTEBuilder) *UpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Update sets table name in UPDATE.
func (ub *UpdateBuilder) Update(table ...string) *UpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// TableNames returns all table names in this UPDATE statement.
func (ub *UpdateBuilder) TableNames() (tableNames []string) { _ = "STUB: not implemented"; return nil }

// Set sets the assignments in SET.
func (ub *UpdateBuilder) Set(assignment ...string) *UpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// SetMore appends the assignments in SET.
func (ub *UpdateBuilder) SetMore(assignment ...string) *UpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// From sets table names of FROM in UPDATE.
func (ub *UpdateBuilder) From(table ...string) *UpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Where adds expressions to the WHERE clause in UPDATE.
//
// Multiple calls to Where will join expressions with AND.
// To reset the WHERE clause, set the WhereClause field to nil.
func (ub *UpdateBuilder) Where(andExpr ...string) *UpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// AddWhereClause adds all clauses in the whereClause to SELECT.
func (ub *UpdateBuilder) AddWhereClause(whereClause *WhereClause) *UpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Assign represents SET "field = value" in UPDATE.
func (ub *UpdateBuilder) Assign(field string, value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// Incr represents SET "field = field + 1" in UPDATE.
func (ub *UpdateBuilder) Incr(field string) string { _ = "STUB: not implemented"; return "" }

// Decr represents SET "field = field - 1" in UPDATE.
func (ub *UpdateBuilder) Decr(field string) string { _ = "STUB: not implemented"; return "" }

// Add represents SET "field = field + value" in UPDATE.
func (ub *UpdateBuilder) Add(field string, value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// Sub represents SET "field = field - value" in UPDATE.
func (ub *UpdateBuilder) Sub(field string, value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// Mul represents SET "field = field * value" in UPDATE.
func (ub *UpdateBuilder) Mul(field string, value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// Div represents SET "field = field / value" in UPDATE.
func (ub *UpdateBuilder) Div(field string, value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// OrderBy sets columns of ORDER BY in UPDATE.
//
// It's recommended to use OrderByAsc or OrderByDesc instead for better support of multiple ORDER BY columns with different directions.
// OrderBy combined with Asc/Desc only supports a single direction for all columns.
func (ub *UpdateBuilder) OrderBy(col ...string) *UpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// OrderByAsc sets a column of ORDER BY in UPDATE with ASC order.
// It supports chaining multiple calls to add multiple ORDER BY columns with different directions.
//
//	ub.OrderByAsc("name").OrderByDesc("id")
//	// Generates: ORDER BY name ASC, id DESC
func (ub *UpdateBuilder) OrderByAsc(col string) *UpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// OrderByDesc sets a column of ORDER BY in UPDATE with DESC order.
// It supports chaining multiple calls to add multiple ORDER BY columns with different directions.
//
//	ub.OrderByDesc("id").OrderByAsc("name")
//	// Generates: ORDER BY id DESC, name ASC
func (ub *UpdateBuilder) OrderByDesc(col string) *UpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Asc sets order of ORDER BY to ASC.
//
// Deprecated: Use OrderByAsc instead. Asc only supports a single direction for all ORDER BY columns.
func (ub *UpdateBuilder) Asc() *UpdateBuilder { _ = "STUB: not implemented"; return nil }

// Desc sets order of ORDER BY to DESC.
//
// Deprecated: Use OrderByDesc instead. Desc only supports a single direction for all ORDER BY columns.
func (ub *UpdateBuilder) Desc() *UpdateBuilder { _ = "STUB: not implemented"; return nil }

// Limit sets the LIMIT in UPDATE.
func (ub *UpdateBuilder) Limit(limit int) *UpdateBuilder { _ = "STUB: not implemented"; return nil }

// Returning sets returning columns.
// For DBMS that doesn't support RETURNING, e.g. MySQL, it will be ignored.
func (ub *UpdateBuilder) Returning(col ...string) *UpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// NumAssignment returns the number of assignments to update.
func (ub *UpdateBuilder) NumAssignment() int { _ = "STUB: not implemented"; return 0 }

// String returns the compiled UPDATE string.
func (ub *UpdateBuilder) String() string { _ = "STUB: not implemented"; return "" }

// Build returns compiled UPDATE string and args.
// They can be used in `DB#Query` of package `database/sql` directly.
func (ub *UpdateBuilder) Build() (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// BuildWithFlavor returns compiled UPDATE string and args with flavor and initial args.
// They can be used in `DB#Query` of package `database/sql` directly.
func (ub *UpdateBuilder) BuildWithFlavor(flavor Flavor, initialArg ...interface{}) (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// CTE table names should be written after UPDATE keyword in MySQL.

// For ISO SQL, CTE table names should be written after FROM keyword.

// SetFlavor sets the flavor of compiled sql.
func (ub *UpdateBuilder) SetFlavor(flavor Flavor) (old Flavor) {
	_ = "STUB: not implemented"
	return *new(Flavor)
}

// Flavor returns flavor of builder
func (ub *UpdateBuilder) Flavor() Flavor {
	_ = "STUB: not implemented"
	return *

	// SQL adds an arbitrary sql to current position.
	new(Flavor)
}

func (ub *UpdateBuilder) SQL(sql string) *UpdateBuilder { _ = "STUB: not implemented"; return nil }
