// Copyright 2023 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

import (
	"io"
	"strings"
)

type stringBuilder struct {
	builder *strings.Builder
}

var _ io.Writer = new(stringBuilder)

func newStringBuilder() *stringBuilder { _ = "STUB: not implemented"; return nil }

// WriteLeadingString writes s to internal buffer.
// If it's not the first time to write the string, a blank (" ") will be written before s.
func (sb *stringBuilder) WriteLeadingString(s string) { _ = "STUB: not implemented"; return }

func (sb *stringBuilder) WriteString(s string) { _ = "STUB: not implemented"; return }

func (sb *stringBuilder) WriteStrings(ss []string, sep string) { _ = "STUB: not implemented"; return }

func (sb *stringBuilder) WriteStringsPrefixed(prefix string, ss []string, sep string) {
	_ = "STUB: not implemented"
	return
}

func (sb *stringBuilder) WriteRune(r rune) { _ = "STUB: not implemented"; return }

func (sb *stringBuilder) Write(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (sb *stringBuilder) String() string { _ = "STUB: not implemented"; return "" }

func (sb *stringBuilder) Reset() { _ = "STUB: not implemented"; return }

func (sb *stringBuilder) Grow(n int) {
	_ = "STUB: not implemented"

	// filterEmptyStrings removes empty strings from ss.
	// As ss rarely contains empty strings, filterEmptyStrings tries to avoid allocation if possible.
	return
}

func filterEmptyStrings(ss []string) []string { _ = "STUB: not implemented"; return nil }
