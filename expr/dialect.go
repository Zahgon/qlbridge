package expr

import (
	"bytes"
	"io"

	"github.com/araddon/qlbridge/value"
)

type (
	// DialectWriter Defines interface to allow different dialects
	// to have different escape characters.  IE, given an AST structure
	// allow the writer to control the string output.  Allows translation between
	// different dialects (different escapes) as well as allows normalization
	// indentation, tabs, spacing, etc.
	// - postgres:  literal-escape = ' identity = "
	// - mysql:     literal-escape = " identity = `
	// - cql:       literal-escape = ' identity = `
	// - bigquery:  literal-escape = " identity = []
	DialectWriter interface {
		io.Writer
		Len() int
		WriteLiteral(string)
		WriteIdentity(string)
		WriteLeftRightIdentity(string, string)
		WriteIdentityQuote(string, byte)
		WriteNumber(string)
		WriteNull()
		WriteValue(v value.Value)
		String() string
	}
	// Default Dialect writer uses mysql escaping rules literals=" identity=`
	defaultDialect struct {
		bytes.Buffer
		Null           string
		LiteralQuote   byte
		IdentityQuote  byte
		stripNamespace bool
	}
	// Json or String Dialect writer uses json escaping rules literals=\
	jsonDialect struct {
		*defaultDialect
	}
	// finterprinter, ie ? substitution
	fingerprintDialect struct {
		DialectWriter
		replace string
	}
	// Keyword writer
	keywordDialect struct {
		*defaultDialect
		kw map[string]struct{}
	}
)

// NewDialectWriter creates a writer that is custom literal and identity
// escape characters
func NewDialectWriter(l, i byte) DialectWriter {
	_ = "STUB: not implemented"
	return *new(DialectWriter)
}

// NewJSONDialectWriter escape literal " with \"
func NewJSONDialectWriter() DialectWriter { _ = "STUB: not implemented"; return *new(DialectWriter) }

// NewDefaultWriter uses mysql escaping rules literals=" identity=`
func NewDefaultWriter() DialectWriter { _ = "STUB: not implemented"; return *new(DialectWriter) }

// NewDefaultNoNamspaceWriter uses mysql escaping rules literals=" identity=`
// Strip namespaces so that 'users.first_name` becomes `first_name` (strip users.)
func NewDefaultNoNamspaceWriter() DialectWriter {
	_ = "STUB: not implemented"
	return *new(DialectWriter)
}

// WriteLiteral writes literal with escapes if needed
func (w *defaultDialect) WriteLiteral(l string) { _ = "STUB: not implemented"; return }

// WriteIdentity writes identity with escaping if needed
func (w *defaultDialect) WriteIdentity(i string) { _ = "STUB: not implemented"; return }

// WriteLeftRightIdentity writes identity with escaping if needed
func (w *defaultDialect) WriteLeftRightIdentity(l, r string) { _ = "STUB: not implemented"; return }

// `user`.`email`   type namespacing, may need to be escaped differently

// WriteIdentityQuote write out an identity using given quote character
func (w *defaultDialect) WriteIdentityQuote(i string, quote byte) {
	_ = "STUB: not implemented"
	return
}

func (w *defaultDialect) WriteNumber(n string) { _ = "STUB: not implemented"; return }

func (w *defaultDialect) WriteNull() { _ = "STUB: not implemented"; return }

func (w *defaultDialect) WriteValue(v value.Value) { _ = "STUB: not implemented"; return }

// If you don't want json, then over-ride this WriteValue

// If you don't want json, then over-ride this WriteValue

// WriteLiteral writes literal and escapes " with \"
func (w *jsonDialect) WriteLiteral(l string) { _ = "STUB: not implemented"; return }

func NewKeywordDialect(kw []string) DialectWriter {
	_ = "STUB: not implemented"
	return *new(DialectWriter)
}

func (w *keywordDialect) WriteIdentity(id string) { _ = "STUB: not implemented"; return }

func NewFingerPrinter() DialectWriter { _ = "STUB: not implemented"; return *new(DialectWriter) }

func NewFingerPrintWriter(replace string, w DialectWriter) DialectWriter {
	_ = "STUB: not implemented"
	return *new(DialectWriter)
}

func (w *fingerprintDialect) WriteLiteral(l string) { _ = "STUB: not implemented"; return }

func (w *fingerprintDialect) WriteNumber(n string) { _ = "STUB: not implemented"; return }

func (w *fingerprintDialect) WriteIdentity(id string) { _ = "STUB: not implemented"; return }
