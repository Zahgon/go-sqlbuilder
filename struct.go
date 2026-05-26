// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

import (
	"database/sql/driver"
	"reflect"
	"regexp"
)

var (
	// DBTag is the struct tag to describe the name for a field in struct.
	DBTag = "db"

	// FieldTag is the struct tag to describe the tag name for a field in struct.
	// Use "," to separate different tags.
	FieldTag = "fieldtag"

	// FieldOpt is the options for a struct field.
	// Supported options include withquote, omitempty, expand, and noexpand.
	// As db column can contain "," in theory, field options should be provided in a separated tag.
	FieldOpt = "fieldopt"

	// FieldAs is the column alias (AS) for a struct field.
	FieldAs = "fieldas"

	// NoExpand changes the default behavior for tagged nested struct fields.
	// When true, tagged nested structs stay as a single column unless fieldopt:"expand" is set.
	// Set it before first use of a Struct if that Struct should use the non-expanding default.
	NoExpand = false
)

const (
	fieldOptWithQuote = "withquote"
	fieldOptOmitEmpty = "omitempty"
	fieldOptExpand    = "expand"
	fieldOptNoExpand  = "noexpand"

	optName   = "optName"
	optParams = "optParams"
)

var optRegex = regexp.MustCompile(`(?P<` + optName + `>\w+)(\((?P<` + optParams + `>.*)\))?`)

var typeOfSQLDriverValuer = reflect.TypeOf((*driver.Valuer)(nil)).Elem()

// Struct represents a struct type.
//
// All methods in Struct are thread-safe.
// We can define a global variable to hold a Struct and use it in any goroutine.
type Struct struct {
	Flavor Flavor

	structType         reflect.Type
	structFieldsParser structFieldsParser
	withTags           []string
	withoutTags        []string
}

var emptyStruct Struct

// NewStruct analyzes type information in structValue
// and creates a new Struct with all structValue fields.
// If structValue is not a struct, NewStruct returns a dummy Struct.
func NewStruct(structValue interface{}) *Struct { _ = "STUB: not implemented"; return nil }

// For sets the default flavor of s and returns a shadow copy of s.
// The original s.Flavor is not changed.
func (s *Struct) For(flavor Flavor) *Struct { _ = "STUB: not implemented"; return nil }

// WithFieldMapper returns a new Struct based on s with custom field mapper.
// The original s is not changed.
func (s *Struct) WithFieldMapper(mapper FieldMapperFunc) *Struct {
	_ = "STUB: not implemented"
	return nil
}

// WithTag sets included tag(s) for all builder methods.
// For instance, calling s.WithTag("tag").SelectFrom("t") is to select all fields tagged with "tag" from table "t".
//
// If multiple tags are provided, fields tagged with any of them are included.
// That is, s.WithTag("tag1", "tag2").SelectFrom("t") is to select all fields tagged with "tag1" or "tag2" from table "t".
func (s *Struct) WithTag(tags ...string) *Struct { _ = "STUB: not implemented"; return nil }

func (s *Struct) mergeWithTags(with []string) { _ = "STUB: not implemented"; return }

// Merge with tags.

// WithoutTag sets excluded tag(s) for all builder methods.
// For instance, calling s.WithoutTag("tag").SelectFrom("t") is to select all fields except those tagged with "tag" from table "t".
//
// If multiple tags are provided, fields tagged with any of them are excluded.
// That is, s.WithoutTag("tag1", "tag2").SelectFrom("t") is to exclude any field tagged with "tag1" or "tag2" from table "t".
func (s *Struct) WithoutTag(tags ...string) *Struct { _ = "STUB: not implemented"; return nil }

func (s *Struct) mergeWithoutTags(without []string) { _ = "STUB: not implemented"; return }

// Merge without tags.

// Filter out useless tags in s.withTags.

// Update with and without tags.

func hasTag(tags []string, tag string) bool { _ = "STUB: not implemented"; return false }

func removeDuplicatedTags(tags []string) []string { _ = "STUB: not implemented"; return nil }

// Unlikely to find any duplicates.

// SelectFrom creates a new `SelectBuilder` with table name.
// By default, all exported fields of the s are listed as columns in SELECT.
//
// Caller is responsible to set WHERE condition to find right record.
func (s *Struct) SelectFrom(table string) *SelectBuilder { _ = "STUB: not implemented"; return nil }

// SelectFromForTag creates a new `SelectBuilder` with table name for a specified tag.
// By default, all fields of the s tagged with tag are listed as columns in SELECT.
//
// Caller is responsible to set WHERE condition to find right record.
//
// Deprecated: It's recommended to use s.WithTag(tag).SelectFrom(...) instead of calling this method.
// The former one is more readable and can be chained with other methods.
func (s *Struct) SelectFromForTag(table string, tag string) (sb *SelectBuilder) {
	_ = "STUB: not implemented"
	return nil
}

func (s *Struct) selectFromWithTags(table string, with, without []string) (sb *SelectBuilder) {
	_ = "STUB: not implemented"
	return nil
}

func parseTableAlias(table string) string { _ = "STUB: not implemented"; return "" }

// Update creates a new `UpdateBuilder` with table name.
// By default, all exported fields of the s is assigned in UPDATE with the field values from value.
// If value's type is not the same as that of s, Update returns a dummy `UpdateBuilder` with table name.
//
// Caller is responsible to set WHERE condition to match right record.
func (s *Struct) Update(table string, value interface{}) *UpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// UpdateForTag creates a new `UpdateBuilder` with table name.
// By default, all fields of the s tagged with tag is assigned in UPDATE with the field values from value.
// If value's type is not the same as that of s, UpdateForTag returns a dummy `UpdateBuilder` with table name.
//
// Caller is responsible to set WHERE condition to match right record.
//
// Deprecated: It's recommended to use s.WithTag(tag).Update(...) instead of calling this method.
// The former one is more readable and can be chained with other methods.
func (s *Struct) UpdateForTag(table string, tag string, value interface{}) *UpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *Struct) updateWithTags(table string, with, without []string, value interface{}) *UpdateBuilder {
	_ = "STUB: not implemented"
	return nil
}

// InsertInto creates a new `InsertBuilder` with table name using verb INSERT INTO.
// By default, all exported fields of s are set as columns by calling `InsertBuilder#Cols`,
// and value is added as a list of values by calling `InsertBuilder#Values`.
//
// InsertInto never returns any error.
// If the type of any item in value is not expected, it will be ignored.
// If value is an empty slice, `InsertBuilder#Values` will not be called.
func (s *Struct) InsertInto(table string, value ...interface{}) *InsertBuilder {
	_ = "STUB: not implemented"
	return nil
}

// InsertIgnoreInto creates a new `InsertBuilder` with table name using verb INSERT IGNORE INTO.
// By default, all exported fields of s are set as columns by calling `InsertBuilder#Cols`,
// and value is added as a list of values by calling `InsertBuilder#Values`.
//
// InsertIgnoreInto never returns any error.
// If the type of any item in value is not expected, it will be ignored.
// If value is an empty slice, `InsertBuilder#Values` will not be called.
func (s *Struct) InsertIgnoreInto(table string, value ...interface{}) *InsertBuilder {
	_ = "STUB: not implemented"
	return nil
}

// ReplaceInto creates a new `InsertBuilder` with table name using verb REPLACE INTO.
// By default, all exported fields of s are set as columns by calling `InsertBuilder#Cols`,
// and value is added as a list of values by calling `InsertBuilder#Values`.
//
// ReplaceInto never returns any error.
// If the type of any item in value is not expected, it will be ignored.
// If value is an empty slice, `InsertBuilder#Values` will not be called.
func (s *Struct) ReplaceInto(table string, value ...interface{}) *InsertBuilder {
	_ = "STUB: not implemented"
	return nil
}

// buildColsAndValuesForTag uses ib to set exported fields tagged with tag as columns
// and add value as a list of values.
func (s *Struct) buildColsAndValuesForTag(ib *InsertBuilder, with, without []string, value ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Try to filter out nil values if possible.

// If all values are nil in a column, ignore the column completely.

// InsertIntoForTag creates a new `InsertBuilder` with table name using verb INSERT INTO.
// By default, exported fields tagged with tag are set as columns by calling `InsertBuilder#Cols`,
// and value is added as a list of values by calling `InsertBuilder#Values`.
//
// InsertIntoForTag never returns any error.
// If the type of any item in value is not expected, it will be ignored.
// If value is an empty slice, `InsertBuilder#Values` will not be called.
//
// Deprecated: It's recommended to use s.WithTag(tag).InsertInto(...) instead of calling this method.
// The former one is more readable and can be chained with other methods.
func (s *Struct) InsertIntoForTag(table string, tag string, value ...interface{}) *InsertBuilder {
	_ = "STUB: not implemented"
	return nil
}

// InsertIgnoreIntoForTag creates a new `InsertBuilder` with table name using verb INSERT IGNORE INTO.
// By default, exported fields tagged with tag are set as columns by calling `InsertBuilder#Cols`,
// and value is added as a list of values by calling `InsertBuilder#Values`.
//
// InsertIgnoreIntoForTag never returns any error.
// If the type of any item in value is not expected, it will be ignored.
// If value is an empty slice, `InsertBuilder#Values` will not be called.
//
// Deprecated: It's recommended to use s.WithTag(tag).InsertIgnoreInto(...) instead of calling this method.
// The former one is more readable and can be chained with other methods.
func (s *Struct) InsertIgnoreIntoForTag(table string, tag string, value ...interface{}) *InsertBuilder {
	_ = "STUB: not implemented"
	return nil
}

// ReplaceIntoForTag creates a new `InsertBuilder` with table name using verb REPLACE INTO.
// By default, exported fields tagged with tag are set as columns by calling `InsertBuilder#Cols`,
// and value is added as a list of values by calling `InsertBuilder#Values`.
//
// ReplaceIntoForTag never returns any error.
// If the type of any item in value is not expected, it will be ignored.
// If value is an empty slice, `InsertBuilder#Values` will not be called.
//
// Deprecated: It's recommended to use s.WithTag(tag).ReplaceInto(...) instead of calling this method.
// The former one is more readable and can be chained with other methods.
func (s *Struct) ReplaceIntoForTag(table string, tag string, value ...interface{}) *InsertBuilder {
	_ = "STUB: not implemented"
	return nil
}

// DeleteFrom creates a new `DeleteBuilder` with table name.
//
// Caller is responsible to set WHERE condition to match right record.
func (s *Struct) DeleteFrom(table string) *DeleteBuilder { _ = "STUB: not implemented"; return nil }

// Addr takes address of all exported fields of the s from the st.
// The returned result can be used in `Row#Scan` directly.
func (s *Struct) Addr(st interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

// AddrForTag takes address of all fields of the s tagged with tag from the st.
// The returned value can be used in `Row#Scan` directly.
//
// If tag is not defined in s in advance, returns nil.
//
// Deprecated: It's recommended to use s.WithTag(tag).Addr(...) instead of calling this method.
// The former one is more readable and can be chained with other methods.
func (s *Struct) AddrForTag(tag string, st interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (s *Struct) addrWithTags(with, without []string, st interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// AddrWithCols takes address of all columns defined in cols from the st.
// The returned value can be used in `Row#Scan` directly.
func (s *Struct) AddrWithCols(cols []string, st interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (s *Struct) addrWithFields(fields []*structField, st interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Columns returns column names of s for all exported struct fields.
func (s *Struct) Columns() []string { _ = "STUB: not implemented"; return nil }

// ColumnsForTag returns column names of the s tagged with tag.
//
// Deprecated: It's recommended to use s.WithTag(tag).Columns(...) instead of calling this method.
// The former one is more readable and can be chained with other methods.
func (s *Struct) ColumnsForTag(tag string) (cols []string) { _ = "STUB: not implemented"; return nil }

func (s *Struct) columnsWithTags(with, without []string) (cols []string) {
	_ = "STUB: not implemented"
	return nil
}

// Values returns a shadow copy of all exported fields in st.
func (s *Struct) Values(st interface{}) []interface{} { _ = "STUB: not implemented"; return nil }

// ValuesForTag returns a shadow copy of all fields tagged with tag in st.
//
// Deprecated: It's recommended to use s.WithTag(tag).Values(...) instead of calling this method.
// The former one is more readable and can be chained with other methods.
func (s *Struct) ValuesForTag(tag string, value interface{}) (values []interface{}) {
	_ = "STUB: not implemented"
	return nil
}

func (s *Struct) valuesWithTags(with, without []string, value interface{}) (values []interface{}) {
	_ = "STUB: not implemented"
	return nil
}

// ForeachRead foreach tags.
func (s *Struct) ForeachRead(trans func(dbtag string, isQuoted bool, field reflect.StructField)) {
	_ = "STUB: not implemented"
	return
}

func (s *Struct) foreachReadWithTags(with, without []string, trans func(dbtag string, isQuoted bool, field reflect.StructField)) {
	_ = "STUB: not implemented"
	return
}

// ForeachWrite foreach tags.
func (s *Struct) ForeachWrite(trans func(dbtag string, isQuoted bool, field reflect.StructField)) {
	_ = "STUB: not implemented"
	return
}

func (s *Struct) foreachWriteWithTags(with, without []string, trans func(dbtag string, isQuoted bool, field reflect.StructField)) {
	_ = "STUB: not implemented"
	return
}

func dereferencedType(t reflect.Type) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func dereferencedValue(v reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func dereferencedFieldValue(v reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func fieldByIndex(v reflect.Value, index []int, allocate bool) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

// isEmptyValue checks if v is zero.
// Following code is borrowed from `IsZero` method in `reflect.Value` since Go 1.13.
func isEmptyValue(v reflect.Value) bool { _ = "STUB: not implemented"; return false }
