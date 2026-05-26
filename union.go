// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

import (
	"reflect"

	"github.com/huandu/go-clone"
)

const (
	unionDistinct = " UNION " // Default union type is DISTINCT.
	unionAll      = " UNION ALL "
)

const (
	unionMarkerInit injectionMarker = iota
	unionMarkerAfterUnion
	unionMarkerAfterOrderBy
	unionMarkerAfterLimit
)

// NewUnionBuilder creates a new UNION builder.
func NewUnionBuilder() *UnionBuilder { _ = "STUB: not implemented"; return nil }

func newUnionBuilder() *UnionBuilder { _ = "STUB: not implemented"; return nil }

// Clone returns a deep copy of UnionBuilder.
// It's useful when you want to create a base builder and clone it to build similar queries.
func (ub *UnionBuilder) Clone() *UnionBuilder { _ = "STUB: not implemented"; return nil }

func init() {
	t := reflect.TypeOf(UnionBuilder{})
	clone.SetCustomFunc(t, func(allocator *clone.Allocator, old, new reflect.Value) {
		cloned := allocator.CloneSlowly(old)
		new.Set(cloned)

		ub := cloned.Addr().Interface().(*UnionBuilder)
		for i, b := range ub.builders {
			ub.args.Replace(ub.builderVars[i], b)
		}
	})
}

// UnionBuilder is a builder to build UNION.
type UnionBuilder struct {
	opt         string
	orderByCols []string
	order       string
	limitVar    string
	offsetVar   string

	builders    []Builder
	builderVars []string

	args *Args

	injection *injection
	marker    injectionMarker
}

var _ Builder = new(UnionBuilder)

// Union unions all builders together using UNION operator.
func Union(builders ...Builder) *UnionBuilder { _ = "STUB: not implemented"; return nil }

// Union unions all builders together using UNION operator.
func (ub *UnionBuilder) Union(builders ...Builder) *UnionBuilder {
	_ = "STUB: not implemented"
	return nil
}

// UnionAll unions all builders together using UNION ALL operator.
func UnionAll(builders ...Builder) *UnionBuilder { _ = "STUB: not implemented"; return nil }

// UnionAll unions all builders together using UNION ALL operator.
func (ub *UnionBuilder) UnionAll(builders ...Builder) *UnionBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (ub *UnionBuilder) union(opt string, builders ...Builder) *UnionBuilder {
	_ = "STUB: not implemented"
	return nil
}

// OrderBy sets columns of ORDER BY in SELECT.
//
// It's recommended to use OrderByAsc or OrderByDesc instead for better support of multiple ORDER BY columns with different directions.
// OrderBy combined with Asc/Desc only supports a single direction for all columns.
func (ub *UnionBuilder) OrderBy(col ...string) *UnionBuilder { _ = "STUB: not implemented"; return nil }

// OrderByAsc sets a column of ORDER BY in SELECT with ASC order.
// It supports chaining multiple calls to add multiple ORDER BY columns with different directions.
//
//	ub.OrderByAsc("name").OrderByDesc("id")
//	// Generates: ORDER BY name ASC, id DESC
func (ub *UnionBuilder) OrderByAsc(col string) *UnionBuilder { _ = "STUB: not implemented"; return nil }

// OrderByDesc sets a column of ORDER BY in SELECT with DESC order.
// It supports chaining multiple calls to add multiple ORDER BY columns with different directions.
//
//	ub.OrderByDesc("id").OrderByAsc("name")
//	// Generates: ORDER BY id DESC, name ASC
func (ub *UnionBuilder) OrderByDesc(col string) *UnionBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Asc sets order of ORDER BY to ASC.
//
// Deprecated: Use OrderByAsc instead. Asc only supports a single direction for all ORDER BY columns.
func (ub *UnionBuilder) Asc() *UnionBuilder { _ = "STUB: not implemented"; return nil }

// Desc sets order of ORDER BY to DESC.
//
// Deprecated: Use OrderByDesc instead. Desc only supports a single direction for all ORDER BY columns.
func (ub *UnionBuilder) Desc() *UnionBuilder { _ = "STUB: not implemented"; return nil }

// Limit sets the LIMIT in SELECT.
func (ub *UnionBuilder) Limit(limit int) *UnionBuilder { _ = "STUB: not implemented"; return nil }

// Offset sets the LIMIT offset in SELECT.
func (ub *UnionBuilder) Offset(offset int) *UnionBuilder { _ = "STUB: not implemented"; return nil }

// String returns the compiled SELECT string.
func (ub *UnionBuilder) String() string { _ = "STUB: not implemented"; return "" }

// Build returns compiled SELECT string and args.
// They can be used in `DB#Query` of package `database/sql` directly.
func (ub *UnionBuilder) Build() (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// BuildWithFlavor returns compiled SELECT string and args with flavor and initial args.
// They can be used in `DB#Query` of package `database/sql` directly.
func (ub *UnionBuilder) BuildWithFlavor(flavor Flavor, initialArg ...interface{}) (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// There might be a hidden constraint in Presto requiring offset to be set before limit.
// The select statement documentation (https://prestodb.io/docs/current/sql/select.html)
// puts offset before limit, and Trino, which is based on Presto, seems
// to require this specific order.

// If ORDER BY is not set, sort column #1 by default.
// It's required to make OFFSET...FETCH work.

// It's required to make OFFSET...FETCH work.

// [SKIP N] FIRST M
// M must be greater than 0

// #192: Doris doesn't support ? in OFFSET and LIMIT.

// SetFlavor sets the flavor of compiled sql.
func (ub *UnionBuilder) SetFlavor(flavor Flavor) (old Flavor) {
	_ = "STUB: not implemented"
	return *new(Flavor)
}

// Flavor returns flavor of builder
func (ub *UnionBuilder) Flavor() Flavor {
	_ = "STUB: not implemented"
	return *

	// Var returns a placeholder for value.
	new(Flavor)
}

func (ub *UnionBuilder) Var(arg interface{}) string { _ = "STUB: not implemented"; return "" }

// SQL adds an arbitrary sql to current position.
func (ub *UnionBuilder) SQL(sql string) *UnionBuilder { _ = "STUB: not implemented"; return nil }
