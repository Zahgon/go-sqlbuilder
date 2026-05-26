// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

const (
	lparen = "("
	rparen = ")"
	opOR   = " OR "
	opAND  = " AND "
	opNOT  = "NOT "
)

const minIndexBase = 256

// Cond provides several helper methods to build conditions.
type Cond struct {
	Args *Args
}

// NewCond returns a new Cond.
func NewCond() *Cond { _ = "STUB: not implemented"; return nil }

// Based on the discussion in #174, users may call this method to create
// `Cond` for building various conditions, which is a misuse, but we
// cannot completely prevent this error. To facilitate users in
// identifying the issue when they make mistakes and to avoid
// unexpected stackoverflows, the base index for `Args` is
// deliberately set to a larger non-zero value here. This can
// significantly reduce the likelihood of issues and allows for
// timely error notification to users.

// Equal is used to construct the expression "field = value".
func (c *Cond) Equal(field string, value interface{}) string { _ = "STUB: not implemented"; return "" }

// E is an alias of Equal.
func (c *Cond) E(field string, value interface{}) string { _ = "STUB: not implemented"; return "" }

// EQ is an alias of Equal.
func (c *Cond) EQ(field string, value interface{}) string { _ = "STUB: not implemented"; return "" }

// NotEqual is used to construct the expression "field <> value".
func (c *Cond) NotEqual(field string, value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// NE is an alias of NotEqual.
func (c *Cond) NE(field string, value interface{}) string { _ = "STUB: not implemented"; return "" }

// NEQ is an alias of NotEqual.
func (c *Cond) NEQ(field string, value interface{}) string { _ = "STUB: not implemented"; return "" }

// GreaterThan is used to construct the expression "field > value".
func (c *Cond) GreaterThan(field string, value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// G is an alias of GreaterThan.
func (c *Cond) G(field string, value interface{}) string { _ = "STUB: not implemented"; return "" }

// GT is an alias of GreaterThan.
func (c *Cond) GT(field string, value interface{}) string { _ = "STUB: not implemented"; return "" }

// GreaterEqualThan is used to construct the expression "field >= value".
func (c *Cond) GreaterEqualThan(field string, value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// GE is an alias of GreaterEqualThan.
func (c *Cond) GE(field string, value interface{}) string { _ = "STUB: not implemented"; return "" }

// GTE is an alias of GreaterEqualThan.
func (c *Cond) GTE(field string, value interface{}) string { _ = "STUB: not implemented"; return "" }

// LessThan is used to construct the expression "field < value".
func (c *Cond) LessThan(field string, value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// L is an alias of LessThan.
func (c *Cond) L(field string, value interface{}) string { _ = "STUB: not implemented"; return "" }

// LT is an alias of LessThan.
func (c *Cond) LT(field string, value interface{}) string { _ = "STUB: not implemented"; return "" }

// LessEqualThan is used to construct the expression "field <= value".
func (c *Cond) LessEqualThan(field string, value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// LE is an alias of LessEqualThan.
func (c *Cond) LE(field string, value interface{}) string { _ = "STUB: not implemented"; return "" }

// LTE is an alias of LessEqualThan.
func (c *Cond) LTE(field string, value interface{}) string { _ = "STUB: not implemented"; return "" }

// In is used to construct the expression "field IN (value...)".
func (c *Cond) In(field string, values ...interface{}) string { _ = "STUB: not implemented"; return "" }

// Empty values means "false".

// NotIn is used to construct the expression "field NOT IN (value...)".
func (c *Cond) NotIn(field string, values ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// Empty values means "true".

// Like is used to construct the expression "field LIKE value".
func (c *Cond) Like(field string, value interface{}) string { _ = "STUB: not implemented"; return "" }

// ILike is used to construct the expression "field ILIKE value".
//
// When the database system does not support the ILIKE operator,
// the ILike method will return "LOWER(field) LIKE LOWER(value)"
// to simulate the behavior of the ILIKE operator.
func (c *Cond) ILike(field string, value interface{}) string { _ = "STUB: not implemented"; return "" }

// Use LOWER to simulate ILIKE.

// NotLike is used to construct the expression "field NOT LIKE value".
func (c *Cond) NotLike(field string, value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// NotILike is used to construct the expression "field NOT ILIKE value".
//
// When the database system does not support the ILIKE operator,
// the NotILike method will return "LOWER(field) NOT LIKE LOWER(value)"
// to simulate the behavior of the ILIKE operator.
func (c *Cond) NotILike(field string, value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// Use LOWER to simulate ILIKE.

// IsNull is used to construct the expression "field IS NULL".
func (c *Cond) IsNull(field string) string { _ = "STUB: not implemented"; return "" }

// IsNotNull is used to construct the expression "field IS NOT NULL".
func (c *Cond) IsNotNull(field string) string { _ = "STUB: not implemented"; return "" }

// Between is used to construct the expression "field BETWEEN lower AND upper".
func (c *Cond) Between(field string, lower, upper interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// NotBetween is used to construct the expression "field NOT BETWEEN lower AND upper".
func (c *Cond) NotBetween(field string, lower, upper interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// Or is used to construct the expression OR logic like "expr1 OR expr2 OR expr3".
func (c *Cond) Or(orExpr ...string) string { _ = "STUB: not implemented"; return "" }

// Ensure that there is only 1 memory allocation.

// And is used to construct the expression AND logic like "expr1 AND expr2 AND expr3".
func (c *Cond) And(andExpr ...string) string { _ = "STUB: not implemented"; return "" }

// Ensure that there is only 1 memory allocation.

// Not is used to construct the expression "NOT expr".
func (c *Cond) Not(notExpr string) string { _ = "STUB: not implemented"; return "" }

// Ensure that there is only 1 memory allocation.

// Exists is used to construct the expression "EXISTS (subquery)".
func (c *Cond) Exists(subquery interface{}) string { _ = "STUB: not implemented"; return "" }

// NotExists is used to construct the expression "NOT EXISTS (subquery)".
func (c *Cond) NotExists(subquery interface{}) string { _ = "STUB: not implemented"; return "" }

// Any is used to construct the expression "field op ANY (value...)".
func (c *Cond) Any(field, op string, values ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// Empty values means "false".

// All is used to construct the expression "field op ALL (value...)".
func (c *Cond) All(field, op string, values ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// Empty values means "false".

// Some is used to construct the expression "field op SOME (value...)".
func (c *Cond) Some(field, op string, values ...interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// Empty values means "false".

// IsDistinctFrom is used to construct the expression "field IS DISTINCT FROM value".
//
// When the database system does not support the IS DISTINCT FROM operator,
// the NotILike method will return "NOT field <=> value" for MySQL or a
// "CASE ... WHEN ... ELSE ... END" expression to simulate the behavior of
// the IS DISTINCT FROM operator.
func (c *Cond) IsDistinctFrom(field string, value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// CASE
//     WHEN field IS NULL AND value IS NULL THEN 0
//     WHEN field IS NOT NULL AND value IS NOT NULL AND field = value THEN 0
//     ELSE 1
// END = 1

// IsNotDistinctFrom is used to construct the expression "field IS NOT DISTINCT FROM value".
//
// When the database system does not support the IS NOT DISTINCT FROM operator,
// the NotILike method will return "field <=> value" for MySQL or a
// "CASE ... WHEN ... ELSE ... END" expression to simulate the behavior of
// the IS NOT DISTINCT FROM operator.
func (c *Cond) IsNotDistinctFrom(field string, value interface{}) string {
	_ = "STUB: not implemented"
	return ""
}

// CASE
//     WHEN field IS NULL AND value IS NULL THEN 1
//     WHEN field IS NOT NULL AND value IS NOT NULL AND field = value THEN 1
//     ELSE 0
// END = 1

// Var returns a placeholder for value.
func (c *Cond) Var(value interface{}) string { _ = "STUB: not implemented"; return "" }

type condBuilder struct {
	Builder func(ctx *argsCompileContext)
}

func estimateStringsBytes(strs []string) (n int) { _ = "STUB: not implemented"; return 0 }
