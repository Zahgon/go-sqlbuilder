// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

import (
	"reflect"
)

// Escape replaces `$` with `$$` in ident.
func Escape(ident string) string { _ = "STUB: not implemented"; return "" }

// EscapeAll replaces `$` with `$$` in all strings of ident.
func EscapeAll(ident ...string) []string { _ = "STUB: not implemented"; return nil }

// Flatten recursively extracts values in slices and returns
// a flattened []interface{} with all values.
// If slices is not a slice, return `[]interface{}{slices}`.
func Flatten(slices interface{}) (flattened []interface{}) { _ = "STUB: not implemented"; return nil }

func flatten(v reflect.Value) (elem interface{}, flattened []interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}

type rawArgs struct {
	expr string
}

// Raw marks the expr as a raw value which will not be added to args.
func Raw(expr string) interface{} { _ = "STUB: not implemented"; return nil }

type listArgs struct {
	args    []interface{}
	isTuple bool
}

// List marks arg as a list of data.
// If arg is `[]int{1, 2, 3}`, it will be compiled to `?, ?, ?` with args `[1 2 3]`.
func List(arg interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// Tuple wraps values into a tuple and can be used as a single value.
func Tuple(values ...interface{}) interface{} { _ = "STUB: not implemented"; return nil }

// TupleNames joins names with tuple format.
// The names is not escaped. Use `EscapeAll` to escape them if necessary.
func TupleNames(names ...string) string { _ = "STUB: not implemented"; return "" }

type namedArgs struct {
	name string
	arg  interface{}
}

// Named creates a named argument.
// Unlike `sql.Named`, this named argument works only with `Build` or `BuildNamed` for convenience
// and will be replaced to a `?` after `Compile`.
func Named(name string, arg interface{}) interface{} { _ = "STUB: not implemented"; return nil }
