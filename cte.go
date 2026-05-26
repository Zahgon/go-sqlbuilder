// Copyright 2024 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

import (
	"reflect"

	"github.com/huandu/go-clone"
)

const (
	cteMarkerInit injectionMarker = iota
	cteMarkerAfterWith
)

// With creates a new CTE builder with default flavor.
func With(tables ...*CTEQueryBuilder) *CTEBuilder { _ = "STUB: not implemented"; return nil }

// WithRecursive creates a new recursive CTE builder with default flavor.
func WithRecursive(tables ...*CTEQueryBuilder) *CTEBuilder { _ = "STUB: not implemented"; return nil }

func newCTEBuilder() *CTEBuilder { _ = "STUB: not implemented"; return nil }

// Clone returns a deep copy of CTEBuilder.
// It's useful when you want to create a base builder and clone it to build similar queries.
func (cteb *CTEBuilder) Clone() *CTEBuilder { _ = "STUB: not implemented"; return nil }

func init() {
	t := reflect.TypeOf(CTEBuilder{})
	clone.SetCustomFunc(t, func(allocator *clone.Allocator, old, new reflect.Value) {
		cloned := allocator.CloneSlowly(old)
		new.Set(cloned)

		cteb := cloned.Addr().Interface().(*CTEBuilder)
		for i, b := range cteb.queries {
			cteb.args.Replace(cteb.queryBuilderVars[i], b)
		}
	})
}

// CTEBuilder is a CTE (Common Table Expression) builder.
type CTEBuilder struct {
	recursive        bool
	queries          []*CTEQueryBuilder
	queryBuilderVars []string

	args *Args

	injection *injection
	marker    injectionMarker
}

var _ Builder = new(CTEBuilder)

// With sets the CTE name and columns.
func (cteb *CTEBuilder) With(queries ...*CTEQueryBuilder) *CTEBuilder {
	_ = "STUB: not implemented"
	return nil
}

// WithRecursive sets the CTE name and columns and turns on the RECURSIVE keyword.
func (cteb *CTEBuilder) WithRecursive(queries ...*CTEQueryBuilder) *CTEBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Select creates a new SelectBuilder to build a SELECT statement using this CTE.
func (cteb *CTEBuilder) Select(col ...string) *SelectBuilder { _ = "STUB: not implemented"; return nil }

// DeleteFrom creates a new DeleteBuilder to build a DELETE statement using this CTE.
func (cteb *CTEBuilder) DeleteFrom(table string) *DeleteBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Update creates a new UpdateBuilder to build an UPDATE statement using this CTE.
func (cteb *CTEBuilder) Update(table string) *UpdateBuilder { _ = "STUB: not implemented"; return nil }

// String returns the compiled CTE string.
func (cteb *CTEBuilder) String() string { _ = "STUB: not implemented"; return "" }

// Build returns compiled CTE string and args.
func (cteb *CTEBuilder) Build() (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// BuildWithFlavor builds a CTE with the specified flavor and initial arguments.
func (cteb *CTEBuilder) BuildWithFlavor(flavor Flavor, initialArg ...interface{}) (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// SetFlavor sets the flavor of compiled sql.
func (cteb *CTEBuilder) SetFlavor(flavor Flavor) (old Flavor) {
	_ = "STUB: not implemented"
	return *new(Flavor)
}

// Flavor returns flavor of builder
func (cteb *CTEBuilder) Flavor() Flavor {
	_ = "STUB: not implemented"
	return *

	// SQL adds an arbitrary sql to current position.
	new(Flavor)
}

func (cteb *CTEBuilder) SQL(sql string) *CTEBuilder { _ = "STUB: not implemented"; return nil }

// TableNames returns all table names in a CTE.
func (cteb *CTEBuilder) TableNames() []string { _ = "STUB: not implemented"; return nil }

// tableNamesForFrom returns a list of table names which should be automatically added to FROM clause.
// It's not public, as this feature is designed only for SelectBuilder/UpdateBuilder/DeleteBuilder right now.
func (cteb *CTEBuilder) tableNamesForFrom() []string {
	_ = "STUB: not implemented"

	// ShouldAddToTableList() unlikely returns true.
	// Count it before allocating any memory for better performance.
	return nil
}
