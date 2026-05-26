// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

import (
	"errors"
)

// Supported flavors.
const (
	invalidFlavor Flavor = iota

	MySQL
	PostgreSQL
	SQLite
	SQLServer
	CQL
	ClickHouse
	Presto
	Oracle
	Informix
	Doris
)

var (
	// DefaultFlavor is the default flavor for all builders.
	DefaultFlavor = MySQL
)

var (
	// ErrInterpolateNotImplemented means the method or feature is not implemented right now.
	ErrInterpolateNotImplemented = errors.New("go-sqlbuilder: interpolation for this flavor is not implemented")

	// ErrInterpolateMissingArgs means there are some args missing in query, so it's not possible to
	// prepare a query with such args.
	ErrInterpolateMissingArgs = errors.New("go-sqlbuilder: not enough args when interpolating")

	// ErrInterpolateUnsupportedArgs means that some types of the args are not supported.
	ErrInterpolateUnsupportedArgs = errors.New("go-sqlbuilder: unsupported args when interpolating")
)

// Flavor is the flag to control the format of compiled sql.
type Flavor int

// String returns the name of f.
func (f Flavor) String() string { _ = "STUB: not implemented"; return "" }

// Interpolate parses sql returned by `Args#Compile` or `Builder`,
// and interpolate args to replace placeholders in the sql.
//
// If there are some args missing in sql, e.g. the number of placeholders are larger than len(args),
// returns ErrMissingArgs error.
func (f Flavor) Interpolate(sql string, args []interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// NewCreateTableBuilder creates a new CREATE TABLE builder with flavor.
func (f Flavor) NewCreateTableBuilder() *CreateTableBuilder { _ = "STUB: not implemented"; return nil }

// NewDeleteBuilder creates a new DELETE builder with flavor.
func (f Flavor) NewDeleteBuilder() *DeleteBuilder { _ = "STUB: not implemented"; return nil }

// NewInsertBuilder creates a new INSERT builder with flavor.
func (f Flavor) NewInsertBuilder() *InsertBuilder { _ = "STUB: not implemented"; return nil }

// NewSelectBuilder creates a new SELECT builder with flavor.
func (f Flavor) NewSelectBuilder() *SelectBuilder { _ = "STUB: not implemented"; return nil }

// NewUpdateBuilder creates a new UPDATE builder with flavor.
func (f Flavor) NewUpdateBuilder() *UpdateBuilder { _ = "STUB: not implemented"; return nil }

// NewUnionBuilder creates a new UNION builder with flavor.
func (f Flavor) NewUnionBuilder() *UnionBuilder { _ = "STUB: not implemented"; return nil }

// NewCTEBuilder creates a new CTE builder with flavor.
func (f Flavor) NewCTEBuilder() *CTEBuilder { _ = "STUB: not implemented"; return nil }

// NewCTETableBuilder creates a new CTE table builder with flavor.
func (f Flavor) NewCTEQueryBuilder() *CTEQueryBuilder { _ = "STUB: not implemented"; return nil }

// Quote adds quote for name to make sure the name can be used safely
// as table name or field name.
//
//   - For MySQL, use back quote (`) to quote name;
//   - For PostgreSQL, SQL Server and SQLite, use double quote (") to quote name.
func (f Flavor) Quote(name string) string { _ = "STUB: not implemented"; return "" }

// PrepareInsertIgnore prepares the insert builder to build insert ignore SQL statement based on the sql flavor
func (f Flavor) PrepareInsertIgnore(table string, ib *InsertBuilder) {
	_ = "STUB: not implemented"
	return
}

// see https://www.postgresql.org/docs/current/sql-insert.html

// add sql statement at the end after values, i.e. INSERT INTO ... ON CONFLICT DO NOTHING

// see https://www.sqlite.org/lang_insert.html

// All other databases do not support insert ignore

// panic if the db flavor is not supported

// Set the table and reset the marker right after insert into
