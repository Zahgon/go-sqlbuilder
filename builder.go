// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

// Builder is a general SQL builder.
// It's used by Args to create nested SQL like the `IN` expression in
// `SELECT * FROM t1 WHERE id IN (SELECT id FROM t2)`.
type Builder interface {
	Build() (sql string, args []interface{})
	BuildWithFlavor(flavor Flavor, initialArg ...interface{}) (sql string, args []interface{})
	Flavor() Flavor
}

type compiledBuilder struct {
	args   *Args
	format string
}

var _ Builder = new(compiledBuilder)

func (cb *compiledBuilder) Build() (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

func (cb *compiledBuilder) BuildWithFlavor(flavor Flavor, initialArg ...interface{}) (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// Flavor returns flavor of builder
// Always returns DefaultFlavor
func (cb *compiledBuilder) Flavor() Flavor { _ = "STUB: not implemented"; return *new(Flavor) }

type flavoredBuilder struct {
	builder Builder
	flavor  Flavor
}

func (fb *flavoredBuilder) Build() (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

func (fb *flavoredBuilder) BuildWithFlavor(flavor Flavor, initialArg ...interface{}) (sql string, args []interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

// Flavor returns flavor of builder
func (fb *flavoredBuilder) Flavor() Flavor {
	_ = "STUB: not implemented"

	// WithFlavor creates a new Builder based on builder with a default flavor.
	return *new(Flavor)
}

func WithFlavor(builder Builder, flavor Flavor) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// Buildf creates a Builder from a format string using `fmt.Sprintf`-like syntax.
// As all arguments will be converted to a string internally, e.g. "$0",
// only `%v` and `%s` are valid.
func Buildf(format string, arg ...interface{}) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// Build creates a Builder from a format string.
// The format string uses special syntax to represent arguments.
// See doc in `Args#Compile` for syntax details.
func Build(format string, arg ...interface{}) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

// BuildNamed creates a Builder from a format string.
// The format string uses `${key}` to refer the value of named by key.
func BuildNamed(format string, named map[string]interface{}) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}
