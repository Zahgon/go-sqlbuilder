package sqlbuilder

import (
	"database/sql"
	"reflect"
)

var typeOfSQLScanner = reflect.TypeOf((*sql.Scanner)(nil)).Elem()

type structFields struct {
	noTag  *structTaggedFields
	tagged map[string]*structTaggedFields
}

type structTaggedFields struct {
	// All columns for SELECT.
	ForRead     []*structField
	colsForRead map[string]*structField

	// All columns which can be used in INSERT and UPDATE.
	ForWrite     []*structField
	colsForWrite map[string]struct{}

	// All columns which can be used in INSERT.
	ForInsert     []*structField
	colsForInsert map[string]struct{}
}

type structField struct {
	Name     string
	Alias    string
	As       string
	Tags     []string
	IsQuoted bool
	DBTag    string
	Field    reflect.StructField
	Index    []int

	omitEmptyTags omitEmptyTagMap
}

type structFieldExpandMode uint8

const (
	structFieldExpandDefault structFieldExpandMode = iota
	structFieldExpandEnabled
	structFieldExpandDisabled
)

type structFieldOptions struct {
	isQuoted      bool
	omitEmptyTags omitEmptyTagMap
	expandMode    structFieldExpandMode
}

type structFieldsParser func() *structFields

func makeDefaultFieldsParser(t reflect.Type) structFieldsParser {
	_ = "STUB: not implemented"
	return *new(structFieldsParser)
}

func makeCustomFieldsParser(t reflect.Type, mapper FieldMapperFunc) structFieldsParser {
	_ = "STUB: not implemented"
	return *new(structFieldsParser)
}

func makeFieldsParser(t reflect.Type, mapper FieldMapperFunc, useDefault bool) structFieldsParser {
	_ = "STUB: not implemented"
	return *new(structFieldsParser)
}

func (sfs *structFields) parse(t reflect.Type, mapper FieldMapperFunc, prefix string, index []int, allowInsert bool) {
	_ = "STUB: not implemented"
	return
}

// Skip unexported fields that are not embedded structs.

// If field is an anonymous struct or pointer to struct, parse it later.

// Parse DBTag.

func makeStructField(field reflect.StructField, alias, dbtag string, mapper FieldMapperFunc, fieldOpts structFieldOptions, prefix string, index []int, fieldIndex int) *structField {
	_ = "STUB: not implemented"
	return nil
}

func parseStructFieldOptions(field reflect.StructField) structFieldOptions {
	_ = "STUB: not implemented"
	return *new(structFieldOptions)
}

func (sfs *structFields) addField(field *structField) { _ = "STUB: not implemented"; return }

func (sfs *structFields) addReadWriteField(field *structField) { _ = "STUB: not implemented"; return }

func (sfs *structFields) addInsertField(field *structField) { _ = "STUB: not implemented"; return }

func appendFieldIndex(prefix []int, index ...int) []int { _ = "STUB: not implemented"; return nil }

func shouldExpandAnonymousStructField(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func shouldExpandTaggedStructField(t reflect.Type, dbtag string, fieldOpts structFieldOptions) bool {
	_ = "STUB: not implemented"
	return false
}

func canExpandStructType(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func implementsScannerOrValuer(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func (sfs *structFields) FilterTags(with, without []string) *structTaggedFields {
	_ = "STUB: not implemented"
	return nil
}

// Simply return the tagged fields.

// Find out all with and without fields.

func (sfs *structFields) taggedFields(tag string) *structTaggedFields {
	_ = "STUB: not implemented"
	return nil
}

func makeStructTaggedFields() *structTaggedFields { _ = "STUB: not implemented"; return nil }

// Add a new field to stfs.
// If field's key exists in stfs.fields, the field is ignored.
func (stfs *structTaggedFields) Add(field *structField) { _ = "STUB: not implemented"; return }

func (stfs *structTaggedFields) AddReadWrite(field *structField) { _ = "STUB: not implemented"; return }

func (stfs *structTaggedFields) AddInsert(field *structField) { _ = "STUB: not implemented"; return }

// Cols returns the fields whose key is one of cols.
// If any column in cols doesn't exist in sfs.fields, returns nil.
func (stfs *structTaggedFields) Cols(cols []string) []*structField {
	_ = "STUB: not implemented"
	return nil
}

// Key returns the key name to identify a field.
func (sf *structField) Key() string { _ = "STUB: not implemented"; return "" }

// NameForSelect returns the name for SELECT.
func (sf *structField) NameForSelect(flavor Flavor) string { _ = "STUB: not implemented"; return "" }

// Quote the Alias in sf with flavor.
func (sf *structField) Quote(flavor Flavor) string { _ = "STUB: not implemented"; return "" }

// ShouldOmitEmpty returns true only if any one of tags is in the omitted tags map.
func (sf *structField) ShouldOmitEmpty(tags ...string) (ret bool) {
	_ = "STUB: not implemented"
	return false
}

// Always check default tag.

type omitEmptyTagMap map[string]struct{}

func getOptMatchedMap(opt string) (res map[string]string) { _ = "STUB: not implemented"; return nil }

func getTagsFromOptParams(opts string) (tags []string) { _ = "STUB: not implemented"; return nil }

func splitTags(fieldtag string) (tags []string) { _ = "STUB: not implemented"; return nil }
