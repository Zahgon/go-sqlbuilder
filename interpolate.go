// Copyright 2018 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package sqlbuilder

// mysqlInterpolate parses query and replace all "?" with encoded args.
// If there are more "?" than len(args), returns ErrMissingArgs.
// Otherwise, if there are less "?" than len(args), the redundant args are omitted.
func mysqlInterpolate(query string, args ...interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func mysqlLikeInterpolate(flavor Flavor, query string, args ...interface{}) (string, error) {
	_ = "STUB: not implemented"
	// Roughly estimate the size to avoid useless memory allocation and copy.
	return "", nil
}

// postgresqlInterpolate parses query and replace all "$*" with encoded args.
// If there are more "$*" than len(args), returns ErrMissingArgs.
// Otherwise, if there are less "$*" than len(args), the redundant args are omitted.
func postgresqlInterpolate(query string, args ...interface{}) (string, error) {
	_ = "STUB: not implemented"
	// Roughly estimate the size to avoid useless memory allocation and copy.
	return "", nil
}

// Try to find the end of dollar quote.

// A placeholder is found.

// Try to find the beginning of dollar quote.

// PostgreSQL uses two single quotes to represent one single quote.

// sqlserverInterpolate parses query and replace all "@p*" with encoded args.
// If there are more "@p*" than len(args), returns ErrMissingArgs.
// Otherwise, if there are less "@p*" than len(args), the redundant args are omitted.
func sqlserverInterpolate(query string, args ...interface{}) (string, error) {
	_ = "STUB: not implemented"
	// Roughly estimate the size to avoid useless memory allocation and copy.
	return "", nil
}

// Only parameters starting with @p or @P are interpolated.

// A placeholder is found.

// mysqlInterpolate works the same as MySQL interpolating.
func sqliteInterpolate(query string, args ...interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// cqlInterpolate works the same as MySQL interpolating.
func cqlInterpolate(query string, args ...interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func clickhouseInterpolate(query string, args ...interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func prestoInterpolate(query string, args ...interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func informixInterpolate(query string, args ...interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func dorisInterpolate(query string, args ...interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// oraclelInterpolate parses query and replace all ":*" with encoded args.
// If there are more ":*" than len(args), returns ErrMissingArgs.
// Otherwise, if there are less ":*" than len(args), the redundant args are omitted.
func oracleInterpolate(query string, args ...interface{}) (string, error) {
	_ = "STUB: not implemented"
	// Roughly estimate the size to avoid useless memory allocation and copy.
	return "", nil
}

// Try to find the end of dollar quote.

// A placeholder is found.

// Try to find the beginning of dollar quote.

// PostgreSQL uses two single quotes to represent one single quote.

func encodeValue(buf []byte, arg interface{}, flavor Flavor) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// In SQL standard, the precision of fractional seconds in time literal is up to 6 digits.
// Round up v.

// Handle typed nil values (e.g. (*string)(nil), (*time.Time)(nil))
// This check must come before fmt.Stringer check since nil pointers may implement interfaces

// Check for fmt.Stringer after nil pointer check

// Bytes() will panic if primative is an array and cannot be addressed.
// Copy all bytes to data as a fallback.

var hexDigits = [16]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}

func appendHex(buf, v []byte) []byte { _ = "STUB: not implemented"; return nil }

func quoteStringValue(buf []byte, s string, flavor Flavor) []byte {
	_ = "STUB: not implemented"
	return nil
}
