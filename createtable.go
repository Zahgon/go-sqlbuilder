// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

const (
	createTableMarkerInit injectionMarker = iota
	createTableMarkerAfterCreate
	createTableMarkerAfterDefine
	createTableMarkerAfterOption
)

// NewCreateTableBuilder creates a new CREATE TABLE builder.
func NewCreateTableBuilder() *CreateTableBuilder { _ = "STUB: not implemented"; return nil }

func newCreateTableBuilder() *CreateTableBuilder { _ = "STUB: not implemented"; return nil }

// Clone returns a deep copy of CreateTableBuilder.
// It's useful when you want to create a base builder and clone it to build similar queries.
func (ctb *CreateTableBuilder) Clone() *CreateTableBuilder { _ = "STUB: not implemented"; return nil }

// CreateTableBuilder is a builder to build CREATE TABLE.
type CreateTableBuilder struct {
	verb        string
	ifNotExists bool
	table       string
	defs        [][]string
	options     [][]string

	args *Args

	injection *injection
	marker    injectionMarker
}

var _ Builder = new(CreateTableBuilder)

// CreateTable sets the table name in CREATE TABLE.
func CreateTable(table string) *CreateTableBuilder { _ = "STUB: not implemented"; return nil }

// CreateTable sets the table name in CREATE TABLE.
func (ctb *CreateTableBuilder) CreateTable(table string) *CreateTableBuilder {
	_ = "STUB: not implemented"
	return nil
}

// CreateTempTable sets the table name and changes the verb of ctb to CREATE TEMPORARY TABLE.
func (ctb *CreateTableBuilder) CreateTempTable(table string) *CreateTableBuilder {
	_ = "STUB: not implemented"
	return nil
}

// IfNotExists adds IF NOT EXISTS before table name in CREATE TABLE.
func (ctb *CreateTableBuilder) IfNotExists() *CreateTableBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Define adds definition of a column or index in CREATE TABLE.
func (ctb *CreateTableBuilder) Define(def ...string) *CreateTableBuilder {
	_ = "STUB: not implemented"
	return nil
}

// Option adds a table option in CREATE TABLE.
func (ctb *CreateTableBuilder) Option(opt ...string) *CreateTableBuilder {
	_ = "STUB: not implemented"
	return nil
}

// NumDefine returns the number of definitions in CREATE TABLE.
func (ctb *CreateTableBuilder) NumDefine() int { _ = "STUB: not implemented"; return 0 }

// String returns the compiled INSERT string.
func (ctb *CreateTableBuilder) String() string { _ = "STUB: not implemented"; return "" }

// Build returns compiled CREATE TABLE string and args.
// They can be used in `DB#Query` of package `database/sql` directly.
func (ctb *CreateTableBuilder) Build() (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// BuildWithFlavor returns compiled CREATE TABLE string and args with flavor and initial args.
// They can be used in `DB#Query` of package `database/sql` directly.
func (ctb *CreateTableBuilder) BuildWithFlavor(flavor Flavor, initialArg ...interface{}) (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// SetFlavor sets the flavor of compiled sql.
func (ctb *CreateTableBuilder) SetFlavor(flavor Flavor) (old Flavor) {
	_ = "STUB: not implemented"
	return *new(Flavor)
}

// Flavor returns flavor of builder
func (ctb *CreateTableBuilder) Flavor() Flavor {
	_ = "STUB: not implemented"
	return *

	// Var returns a placeholder for value.
	new(Flavor)
}

func (ctb *CreateTableBuilder) Var(arg interface{}) string { _ = "STUB: not implemented"; return "" }

// SQL adds an arbitrary sql to current position.
func (ctb *CreateTableBuilder) SQL(sql string) *CreateTableBuilder {
	_ = "STUB: not implemented"
	return nil
}
