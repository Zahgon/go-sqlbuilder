// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/huandu/go-clone"
)

// Args stores arguments associated with a SQL.
type Args struct {
	// The default flavor used by `Args#Compile`
	Flavor Flavor

	indexBase    int
	argValues    *valueStore
	namedArgs    map[string]int
	sqlNamedArgs map[string]int
	onlyNamed    bool
}

func init() {
	// Predefine some $n args to avoid additional memory allocation.
	predefinedArgs = make([]string, 0, maxPredefinedArgs)

	for i := 0; i < maxPredefinedArgs; i++ {
		predefinedArgs = append(predefinedArgs, fmt.Sprintf("$%v", i))
	}
}

const maxPredefinedArgs = 64

var predefinedArgs []string

// Add adds an arg to Args and returns a placeholder.
func (args *Args) Add(arg interface{}) string { _ = "STUB: not implemented"; return "" }

func (args *Args) add(arg interface{}) int { _ = "STUB: not implemented"; return 0 }

// Find out the real arg and add it to args.

// Replace replaces the placeholder with arg.
//
// The placeholder must be the value returned by `Add`, e.g. "$1".
// If the placeholder is not found, this method does nothing.
func (args *Args) Replace(placeholder string, arg interface{}) { _ = "STUB: not implemented"; return }

// Compile compiles builder's format to standard sql and returns associated args.
//
// The format string uses a special syntax to represent arguments.
//
//	$? refers successive arguments passed in the call. It works similar as `%v` in `fmt.Sprintf`.
//	$0 $1 ... $n refers nth-argument passed in the call. Next $? will use arguments n+1.
//	${name} refers a named argument created by `Named` with `name`.
//	$$ is a "$" string.
func (args *Args) Compile(format string, initialValue ...interface{}) (query string, values []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// CompileWithFlavor compiles builder's format to standard sql with flavor and returns associated args.
//
// See doc for `Compile` to learn details.
func (args *Args) CompileWithFlavor(format string, flavor Flavor, initialValue ...interface{}) (query string, values []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// Treat the $ at the end of format is a normal $ rune.

// For unknown $ expression format, treat it as a normal $ rune.

// Value returns the value of the arg.
// The arg must be the value returned by `Add`.
func (args *Args) Value(arg string) interface{} { _ = "STUB: not implemented"; return nil }

func (args *Args) compileNamed(ctx *argsCompileContext, format string) string {
	_ = "STUB: not implemented"
	return ""
}

// Nothing.

// Invalid $ format. Ignore it.

func (args *Args) compileDigits(ctx *argsCompileContext, format string, offset int) (string, int) {
	_ = "STUB: not implemented"
	return "", 0
}

// Nothing.

func (args *Args) compileSuccessive(ctx *argsCompileContext, format string, offset int) (string, int) {
	_ = "STUB: not implemented"
	return "", 0
}

func (args *Args) mergeSQLNamedArgs(ctx *argsCompileContext) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// Add all named args to values.
// Remove duplicated named args in this step.

// Stabilize the sequence to make it easier to write test cases.

func parseNamedArgs(initialValue []interface{}) (values []interface{}, namedValues []sql.NamedArg) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sql.NamedArgs must be placed at the end of the initial value.

type argsCompileContext struct {
	*stringBuilder

	Flavor    Flavor
	Values    []interface{}
	NamedArgs []sql.NamedArg
}

func (ctx *argsCompileContext) WriteValue(arg interface{}) { _ = "STUB: not implemented"; return }

// Add all values to ctx.
// Named args must be located at the end of values.

func (ctx *argsCompileContext) WriteValues(values []interface{}, sep string) {
	_ = "STUB: not implemented"
	return
}

type valueStore struct {
	Values []interface{}
}

func init() {
	// The values in valueStore should be shadow-copied to avoid unnecessary cost.
	t := reflect.TypeOf(valueStore{})
	clone.SetCustomFunc(t, func(allocator *clone.Allocator, old, new reflect.Value) {
		values := old.FieldByName("Values")
		newValues := allocator.Clone(values)
		new.FieldByName("Values").Set(newValues)
	})
}

func (as *valueStore) Len() int { _ = "STUB: not implemented"; return 0 }

// Add adds an arg to argsValues and returns its index.
func (as *valueStore) Add(arg interface{}) int { _ = "STUB: not implemented"; return 0 }

// Set sets the arg value by index.
func (as *valueStore) Set(index int, arg interface{}) { _ = "STUB: not implemented"; return }

// Load returns the arg value by index.
// Returns nil if index is out of range or as itself is nil.
func (as *valueStore) Load(index int) interface{} { _ = "STUB: not implemented"; return nil }
