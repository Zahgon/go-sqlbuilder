// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

import (
	"reflect"

	"github.com/huandu/go-clone"
)

const (
	deleteMarkerInit injectionMarker = iota
	deleteMarkerAfterWith
	deleteMarkerAfterDeleteFrom
	deleteMarkerAfterWhere
	deleteMarkerAfterOrderBy
	deleteMarkerAfterLimit
	deleteMarkerAfterReturning
)

// NewDeleteBuilder creates a new DELETE builder.
func NewDeleteBuilder() *DeleteBuilder { _ = "STUB: not implemented"; return nil }

func newDeleteBuilder() *DeleteBuilder { _ = "STUB: not implemented"; return nil }

// Clone returns a deep copy of DeleteBuilder.
// It's useful when you want to create a base builder and clone it to build similar queries.
func (db *DeleteBuilder) Clone() *DeleteBuilder { _ = "STUB: not implemented"; return nil }

func init() {
	t := reflect.TypeOf(DeleteBuilder{})
	clone.SetCustomFunc(t, func(allocator *clone.Allocator, old, new reflect.Value) {
		cloned := allocator.CloneSlowly(old)
		new.Set(cloned)

		db := cloned.Addr().Interface().(*DeleteBuilder)
		db.args.Replace(db.whereClauseExpr, db.whereClauseProxy)
		db.args.Replace(db.cteBuilderVar, db.cteBuilder)
	})
}

// DeleteBuilder is a builder to build DELETE.
type DeleteBuilder struct {
	*WhereClause
	Cond

	whereClauseProxy *whereClauseProxy
	whereClauseExpr  string

	cteBuilderVar string
	cteBuilder    *CTEBuilder

	tables      []string
	orderByCols []string
	order       string
	limitVar    string
	returning   []string

	args *Args

	injection *injection
	marker    injectionMarker
}

var _ Builder = new(DeleteBuilder)

// DeleteFrom sets table name in DELETE.
func DeleteFrom(table ...string) *DeleteBuilder { _ = "STUB: not implemented"; return nil }

// With sets WITH clause (the Common Table Expression) before DELETE.
func (db *DeleteBuilder) With(builder *CTEBuilder) *DeleteBuilder {
	_ = "STUB: not implemented"
	return nil
}

// DeleteFrom sets table name in DELETE.
func (db *DeleteBuilder) DeleteFrom(table ...string) *DeleteBuilder {
	_ = "STUB: not implemented"
	return nil
}

// TableNames returns all table names in this DELETE statement.
func (db *DeleteBuilder) TableNames() []string { _ = "STUB: not implemented"; return nil }

// Where adds expressions to the WHERE clause in DELETE.
//
// Multiple calls to Where will join expressions with AND.
// To reset the WHERE clause, set the WhereClause field to nil.
func (db *DeleteBuilder) Where(andExpr ...string) *DeleteBuilder {
	_ = "STUB: not implemented"
	return nil
}

// AddWhereClause adds all clauses in the whereClause to SELECT.
func (db *DeleteBuilder) AddWhereClause(whereClause *WhereClause) *DeleteBuilder {
	_ = "STUB: not implemented"
	return nil
}

// OrderBy sets columns of ORDER BY in DELETE.
//
// It's recommended to use OrderByAsc or OrderByDesc instead for better support of multiple ORDER BY columns with different directions.
// OrderBy combined with Asc/Desc only supports a single direction for all columns.
func (db *DeleteBuilder) OrderBy(col ...string) *DeleteBuilder {
	_ = "STUB: not implemented"
	return nil
}

// OrderByAsc sets a column of ORDER BY in DELETE with ASC order.
// It supports chaining multiple calls to add multiple ORDER BY columns with different directions.
//
//	db.OrderByAsc("name").OrderByDesc("id")
//	// Generates: ORDER BY name ASC, id DESC
func (db *DeleteBuilder) OrderByAsc(col string) *DeleteBuilder {
	_ = "STUB: not implemented"
	return nil
}

// OrderByDesc sets a column of ORDER BY in DELETE with DESC order.
// It supports chaining multiple calls to add multiple ORDER BY columns with different directions.
//
//	db.OrderByDesc("id").OrderByAsc("name")
//	// Generates: ORDER BY id DESC, name ASC
func (db *DeleteBuilder) OrderByDesc(col string) *DeleteBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Asc sets order of ORDER BY to ASC.
//
// Deprecated: Use OrderByAsc instead. Asc only supports a single direction for all ORDER BY columns.
func (db *DeleteBuilder) Asc() *DeleteBuilder { _ = "STUB: not implemented"; return nil }

// Desc sets order of ORDER BY to DESC.
//
// Deprecated: Use OrderByDesc instead. Desc only supports a single direction for all ORDER BY columns.
func (db *DeleteBuilder) Desc() *DeleteBuilder { _ = "STUB: not implemented"; return nil }

// Limit sets the LIMIT in DELETE.
func (db *DeleteBuilder) Limit(limit int) *DeleteBuilder { _ = "STUB: not implemented"; return nil }

// Returning sets returning columns.
// For DBMS that doesn't support RETURNING, e.g. MySQL, it will be ignored.
func (db *DeleteBuilder) Returning(col ...string) *DeleteBuilder {
	_ = "STUB: not implemented"
	return nil
}

// String returns the compiled DELETE string.
func (db *DeleteBuilder) String() string { _ = "STUB: not implemented"; return "" }

// Build returns compiled DELETE string and args.
// They can be used in `DB#Query` of package `database/sql` directly.
func (db *DeleteBuilder) Build() (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// BuildWithFlavor returns compiled DELETE string and args with flavor and initial args.
// They can be used in `DB#Query` of package `database/sql` directly.
func (db *DeleteBuilder) BuildWithFlavor(flavor Flavor, initialArg ...interface{}) (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// SetFlavor sets the flavor of compiled sql.
func (db *DeleteBuilder) SetFlavor(flavor Flavor) (old Flavor) {
	_ = "STUB: not implemented"
	return *new(Flavor)
}

// Flavor returns flavor of builder
func (db *DeleteBuilder) Flavor() Flavor {
	_ = "STUB: not implemented"
	return *

	// SQL adds an arbitrary sql to current position.
	new(Flavor)
}

func (db *DeleteBuilder) SQL(sql string) *DeleteBuilder { _ = "STUB: not implemented"; return nil }
