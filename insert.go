// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

const (
	insertMarkerInit injectionMarker = iota
	insertMarkerAfterInsertInto
	insertMarkerAfterCols
	insertMarkerAfterValues
	insertMarkerAfterSelect
	insertMarkerAfterReturning
)

// NewInsertBuilder creates a new INSERT builder.
func NewInsertBuilder() *InsertBuilder { _ = "STUB: not implemented"; return nil }

func newInsertBuilder() *InsertBuilder { _ = "STUB: not implemented"; return nil }

// Clone returns a deep copy of InsertBuilder.
// It's useful when you want to create a base builder and clone it to build similar queries.
func (ib *InsertBuilder) Clone() *InsertBuilder { _ = "STUB: not implemented"; return nil }

// InsertBuilder is a builder to build INSERT.
type InsertBuilder struct {
	verb      string
	table     string
	cols      []string
	values    [][]string
	returning []string

	args *Args

	injection *injection
	marker    injectionMarker

	sbHolder string
}

var _ Builder = new(InsertBuilder)

// InsertInto sets table name in INSERT.
func InsertInto(table string) *InsertBuilder { _ = "STUB: not implemented"; return nil }

// InsertInto sets table name in INSERT.
func (ib *InsertBuilder) InsertInto(table string) *InsertBuilder {
	_ = "STUB: not implemented"
	return nil
}

// InsertIgnoreInto sets table name in INSERT IGNORE.
func InsertIgnoreInto(table string) *InsertBuilder { _ = "STUB: not implemented"; return nil }

// InsertIgnoreInto sets table name in INSERT IGNORE.
func (ib *InsertBuilder) InsertIgnoreInto(table string) *InsertBuilder {
	_ = "STUB: not implemented"
	return nil
}

// ReplaceInto sets table name and changes the verb of ib to REPLACE.
// REPLACE INTO is a MySQL extension to the SQL standard.
func ReplaceInto(table string) *InsertBuilder { _ = "STUB: not implemented"; return nil }

// ReplaceInto sets table name and changes the verb of ib to REPLACE.
// REPLACE INTO is a MySQL extension to the SQL standard.
func (ib *InsertBuilder) ReplaceInto(table string) *InsertBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Cols sets columns in INSERT.
func (ib *InsertBuilder) Cols(col ...string) *InsertBuilder { _ = "STUB: not implemented"; return nil }

// Select returns a new SelectBuilder to build a SELECT statement inside the INSERT INTO.
func (isb *InsertBuilder) Select(col ...string) *SelectBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Values adds a list of values for a row in INSERT.
func (ib *InsertBuilder) Values(value ...interface{}) *InsertBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Returning sets returning columns.
// For DBMS that doesn't support RETURNING, e.g. MySQL, it will be ignored.
func (ib *InsertBuilder) Returning(col ...string) *InsertBuilder {
	_ = "STUB: not implemented"
	return nil
}

// NumValue returns the number of values to insert.
func (ib *InsertBuilder) NumValue() int { _ = "STUB: not implemented"; return 0 }

// String returns the compiled INSERT string.
func (ib *InsertBuilder) String() string { _ = "STUB: not implemented"; return "" }

// Build returns compiled INSERT string and args.
// They can be used in `DB#Query` of package `database/sql` directly.
func (ib *InsertBuilder) Build() (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// BuildWithFlavor returns compiled INSERT string and args with flavor and initial args.
// They can be used in `DB#Query` of package `database/sql` directly.
func (ib *InsertBuilder) BuildWithFlavor(flavor Flavor, initialArg ...interface{}) (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// SetFlavor sets the flavor of compiled sql.
func (ib *InsertBuilder) SetFlavor(flavor Flavor) (old Flavor) {
	_ = "STUB: not implemented"
	return *new(Flavor)
}

// Flavor returns flavor of builder
func (ib *InsertBuilder) Flavor() Flavor {
	_ = "STUB: not implemented"
	return *

	// Var returns a placeholder for value.
	new(Flavor)
}

func (ib *InsertBuilder) Var(arg interface{}) string { _ = "STUB: not implemented"; return "" }

// SQL adds an arbitrary sql to current position.
func (ib *InsertBuilder) SQL(sql string) *InsertBuilder { _ = "STUB: not implemented"; return nil }
