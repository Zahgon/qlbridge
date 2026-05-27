package gentypes

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/rel"
	"github.com/araddon/qlbridge/schema"
	"github.com/araddon/qlbridge/value"
)

var (
	_ = u.EMPTY
	// lets make sure this local interface SchemaColumns
	// matches the SourceTableColumn from schema
	_ schema.SourceTableColumn = (*fsc)(nil)
	_ SchemaColumns            = (*fsc)(nil)
)

type (
	// FilterValidate interface Will walk a filter statement validating columns, types
	// against underlying Schema.
	FilterValidate func(fs *rel.FilterStatement) error

	// SchemaColumns provides info on fields/columns to help the generator
	// understand how to map Columns to Underlying es fields
	SchemaColumns interface {
		// Underlying data type of column
		Column(col string) (value.ValueType, bool)
		// ColumnInfo of a FilterStatement column explains this column
		// and how to map to Elasticsearch field or false if the field
		// doesn't exist.
		ColumnInfo(col string) (*FieldType, bool)
	}
	// FieldType Describes a field's usage within Elasticsearch
	// - is it nested? which changes query semantics
	// - prefix for nested object values
	FieldType struct {
		Field    string // Field Name
		Prefix   string // .f, .b, .i, .v for nested object types
		Path     string // mapstr_fieldname ,etc, prefixed
		Type     value.ValueType
		TypeName string
	}
	// Payload is the top Level Request to Elasticsearch
	Payload struct {
		Size   *int                   `json:"size,omitempty"`
		Filter interface{}            `json:"filter,omitempty"`
		Fields []string               `json:"fields,omitempty"`
		Sort   []map[string]SortOrder `json:"sort,omitempty"`
	}
	// SortOder of the es query request
	SortOrder struct {
		Order string `json:"order"`
	}

	fsc struct{}
)

func (m *fsc) Column(col string) (value.ValueType, bool) {
	_ = "STUB: not implemented"
	return *new(value.ValueType), false
}
func (m *fsc) ColumnInfo(col string) (*FieldType, bool) {
	_ = "STUB: not implemented"

	// Numeric returns true if field type has numeric values.
	return nil, false
}

func (f *FieldType) Numeric() bool { _ = "STUB: not implemented"; return false }

// If a nested field with numeric values it's numeric

// Nothing else is numeric

func (f *FieldType) Nested() bool   { _ = "STUB: not implemented"; return false }
func (f *FieldType) String() string { _ = "STUB: not implemented"; return "" }

func (f *FieldType) PathAndPrefix(val string) string { _ = "STUB: not implemented"; return "" }

func (f *FieldType) PrefixAndValue(val interface{}) (string, interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *Payload) SortAsc(field string) { _ = "STUB: not implemented"; return }

func (p *Payload) SortDesc(field string) { _ = "STUB: not implemented"; return }

// For Fields declared as map[string]type  (type  = int, string, time, bool, value)
// in lql, we need to determine which nested key/value combo to search for
func ValueAndPrefix(val interface{}) (interface{}, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// Most values come through as strings

// Default to strings
