package es2gen

import (
	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/generators/elasticsearch/gentypes"
)

// fieldType return the Elasticsearch field name for an identity node or an error.
func fieldType(s gentypes.SchemaColumns, n expr.Node) (*gentypes.FieldType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This shotgun approach sucks, see https://github.com/lytics/lio/issues/7565

// This is legacy crap, we stupidly used to allow this:
//  ticket to remove https://github.com/lytics/lio/issues/7565
//
//   `key_name.field value` -> "key_name", "field value"
//
// check if key is left.right

// Nested field lookup
