// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

// WhereClause is a Builder for WHERE clause.
// All builders which support `WHERE` clause have an anonymous `WhereClause` field,
// in which the conditions are stored.
//
// WhereClause can be shared among multiple builders.
// However, it is not thread-safe.
type WhereClause struct {
	flavor  Flavor
	clauses []clause
}

var _ Builder = new(WhereClause)

// NewWhereClause creates a new WhereClause.
func NewWhereClause() *WhereClause { _ = "STUB: not implemented"; return nil }

// CopyWhereClause creates a copy of the whereClause.
func CopyWhereClause(whereClause *WhereClause) *WhereClause { _ = "STUB: not implemented"; return nil }

type clause struct {
	args     *Args
	andExprs []string
}

func (c *clause) Build(flavor Flavor, initialArg ...interface{}) (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// whereClauseProxy is a proxy for WhereClause.
// It's useful when the WhereClause in a build can be changed.
type whereClauseProxy struct {
	*WhereClause
}

var _ Builder = new(whereClauseProxy)

// BuildWithFlavor builds a WHERE clause with the specified flavor and initial arguments.
func (wc *WhereClause) BuildWithFlavor(flavor Flavor, initialArg ...interface{}) (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// Build returns compiled WHERE clause string and args.
func (wc *WhereClause) Build() (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// SetFlavor sets the flavor of compiled sql.
// When the WhereClause belongs to a builder, the flavor of the builder will be used when building SQL.
func (wc *WhereClause) SetFlavor(flavor Flavor) (old Flavor) {
	_ = "STUB: not implemented"
	return *new(Flavor)
}

// Flavor returns flavor of clause
func (wc *WhereClause) Flavor() Flavor {
	_ = "STUB: not implemented"

	// AddWhereExpr adds an AND expression to WHERE clause with the specified arguments.
	return *new(Flavor)
}

func (wc *WhereClause) AddWhereExpr(args *Args, andExpr ...string) *WhereClause {
	_ = "STUB: not implemented"
	return nil
}

// Merge with last clause if possible.

// AddWhereClause adds all clauses in the whereClause to the wc.
func (wc *WhereClause) AddWhereClause(whereClause *WhereClause) *WhereClause {
	_ = "STUB: not implemented"
	return nil
}
