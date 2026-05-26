// Copyright 2024 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

import (
	"reflect"

	"github.com/huandu/go-clone"
)

const (
	cteQueryMarkerInit injectionMarker = iota
	cteQueryMarkerAfterTable
	cteQueryMarkerAfterAs
)

// CTETable creates a new CTE query builder with default flavor, marking it as a table.
//
// The resulting CTE query can be used in a `SelectBuilder“, where its table name will be
// automatically included in the FROM clause.
func CTETable(name string, cols ...string) *CTEQueryBuilder { _ = "STUB: not implemented"; return nil }

// CTEQuery creates a new CTE query builder with default flavor.
func CTEQuery(name string, cols ...string) *CTEQueryBuilder { _ = "STUB: not implemented"; return nil }

func newCTEQueryBuilder() *CTEQueryBuilder { _ = "STUB: not implemented"; return nil }

// Clone returns a deep copy of CTEQueryBuilder.
// It's useful when you want to create a base builder and clone it to build similar queries.
func (ctetb *CTEQueryBuilder) Clone() *CTEQueryBuilder { _ = "STUB: not implemented"; return nil }

func init() {
	t := reflect.TypeOf(CTEQueryBuilder{})
	clone.SetCustomFunc(t, func(allocator *clone.Allocator, old, new reflect.Value) {
		cloned := allocator.CloneSlowly(old)
		new.Set(cloned)

		ctetb := cloned.Addr().Interface().(*CTEQueryBuilder)
		ctetb.args.Replace(ctetb.builderVar, ctetb.builder)
	})
}

// CTEQueryBuilder is a builder to build one table in CTE (Common Table Expression).
type CTEQueryBuilder struct {
	name       string
	cols       []string
	builder    Builder
	builderVar string

	// if true, this query's table name will be automatically added to the table list
	// in FROM clause of SELECT statement.
	autoAddToTableList bool

	args *Args

	injection *injection
	marker    injectionMarker
}

var _ Builder = new(CTEQueryBuilder)

// CTETableBuilder is an alias of CTEQueryBuilder for backward compatibility.
//
// Deprecated: use CTEQueryBuilder instead.
type CTETableBuilder = CTEQueryBuilder

// Table sets the table name and columns in a CTE table.
func (ctetb *CTEQueryBuilder) Table(name string, cols ...string) *CTEQueryBuilder {
	_ = "STUB: not implemented"
	return nil
}

// As sets the builder to select data.
func (ctetb *CTEQueryBuilder) As(builder Builder) *CTEQueryBuilder {
	_ = "STUB: not implemented"
	return nil
}

// AddToTableList sets flag to add table name to table list in FROM clause of SELECT statement.
func (ctetb *CTEQueryBuilder) AddToTableList() *CTEQueryBuilder {
	_ = "STUB: not implemented"
	return nil
}

// ShouldAddToTableList returns flag to add table name to table list in FROM clause of SELECT statement.
func (ctetb *CTEQueryBuilder) ShouldAddToTableList() bool { _ = "STUB: not implemented"; return false }

// String returns the compiled CTE string.
func (ctetb *CTEQueryBuilder) String() string { _ = "STUB: not implemented"; return "" }

// Build returns compiled CTE string and args.
func (ctetb *CTEQueryBuilder) Build() (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// BuildWithFlavor builds a CTE with the specified flavor and initial arguments.
func (ctetb *CTEQueryBuilder) BuildWithFlavor(flavor Flavor, initialArg ...interface{}) (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// SetFlavor sets the flavor of compiled sql.
func (ctetb *CTEQueryBuilder) SetFlavor(flavor Flavor) (old Flavor) {
	_ = "STUB: not implemented"
	return *new(Flavor)
}

// Flavor returns flavor of builder
func (ctetb *CTEQueryBuilder) Flavor() Flavor {
	_ = "STUB: not implemented"
	return *

	// SQL adds an arbitrary sql to current position.
	new(Flavor)
}

func (ctetb *CTEQueryBuilder) SQL(sql string) *CTEQueryBuilder {
	_ = "STUB: not implemented"
	return nil
}

// TableName returns the CTE table name.
func (ctetb *CTEQueryBuilder) TableName() string { _ = "STUB: not implemented"; return "" }
