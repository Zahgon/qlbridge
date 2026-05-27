package esgen

import (
	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/generators/elasticsearch/gentypes"
	"github.com/araddon/qlbridge/value"
)

func exprValueType(s gentypes.SchemaColumns, n expr.Node) value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}

// scalar returns a JSONable representation of a scalar node type for use in ES
// filters.
//
// Does not support Null.
func scalar(node expr.Node) (interface{}, bool) { _ = "STUB: not implemented"; return nil, false }

// ES supports string encoded ints

// Make sure this is a scalar value node

func fieldType(s gentypes.SchemaColumns, n expr.Node) (*gentypes.FieldType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: This shotgun approach sucks, see https://github.com/araddon/qlbridge/issues/159

//left, right, _ := expr.LeftRight(ident.Text)
//u.Debugf("left:%q right:%q isNamespaced?%v   key=%v", left, right, ident.HasLeftRight(), ident.OriginalText())

// This is legacy, we stupidly used to allow this:
//
//   `key_name.field value` -> "key_name", "field value"
//
// check if key is left.right

// Nested field lookup

func fieldValueType(s gentypes.SchemaColumns, n expr.Node) (value.ValueType, error) {
	_ = "STUB: not implemented"
	return *new(value.ValueType), nil
}

// TODO: This shotgun approach sucks, see https://github.com/araddon/qlbridge/issues/159

//left, right, _ := expr.LeftRight(ident.Text)
//u.Debugf("left:%q right:%q isNamespaced?%v   key=%v", left, right, ident.HasLeftRight(), ident.OriginalText())

// This is legacy, we stupidly used to allow this:
//
//   `key_name.field value` -> "key_name", "field value"
//
// check if key is left.right

// Nested field lookup
