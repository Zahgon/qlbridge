package builtins

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/value"
)

var _ = u.EMPTY

// len length of array types
//
//	len([1,2,3])     =>  3, true
//	len(not_a_field)   =>  -- NilInt, false
type Length struct{}

// Type is IntType
func (m *Length) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Length) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func lenEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// ArrayIndex  array.index choose the nth element of an array
//
//	// given context input of
//	"items" = [1,2,3]
//
//	array.index(items, 1)     =>  1, true
//	array.index(items, 5)     =>  nil, false
//	array.index(items, -1)    =>  3, true
type ArrayIndex struct{}

// Type unknown - returns single value from SliceValue array
func (m *ArrayIndex) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}
func (m *ArrayIndex) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func arrayIndexEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// array.slice  slice element m -> n of a slice.  First arg must be a slice.
//
//	// given context of
//	"items" = [1,2,3,4,5]
//
//	array.slice(items, 1, 3)     =>  [2,3], true
//	array.slice(items, 2)        =>  [3,4,5], true
//	array.slice(items, -2)       =>  [4,5], true
type ArraySlice struct{}

// Type Unknown for Array Slice
func (m *ArraySlice) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *

	// Validate must be at least 2 args, max of 3
	new(value.ValueType)
}

func (m *ArraySlice) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func arraySliceEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// array.slice(item, start)

// array.slice(item, start, end)

// array.slice(item, start)

// array.slice(item, start, end)

// Map Create a map from two values.   If the right side value is nil
// then does not evaluate.
//
//	map(left, right)    => map[string]value{left:right}
type MapFunc struct{}

// Type is MapValueType
func (m *MapFunc) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }

func (m *MapFunc) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func mapEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// What should the map function be if lh is slice/map?

// MapTime()    Create a map[string]time of each key
//
//	maptime(field)    => map[string]time{field_value:message_timestamp}
//	maptime(field, timestamp) => map[string]time{field_value:timestamp}
type MapTime struct{}

// Type MapTime
func (m *MapTime) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *MapTime) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func mapTimeEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Match a simple pattern match on KEYS (not values) and build a map of all matched values.
// Matched portion is replaced with empty string.
// - May pass as many match strings as you want.
// - Must match on Prefix of key.
//
//	given input context of:
//	   {"score_value":24,"event_click":true, "tag_apple": "apple", "label_orange": "orange"}
//
//	   match("score_") => {"value":24}
//	   match("amount_") => false
//	   match("event_") => {"click":true}
//	   match("label_","tag_") => {"apple":"apple","orange":"orange"}
type Match struct{}

// Type is MapValueType
func (m *Match) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }

func (m *Match) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func matchEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Iterate through every value in Context

// MapKeys:  Take a map and extract array of keys
//
//	//given input:
//	{"tag.1":"news","tag.2":"sports"}
//
//	mapkeys(match("tag.")) => []string{"news","sports"}
type MapKeys struct{}

// Type []string aka strings
func (m *MapKeys) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *MapKeys) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func mapKeysEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// MapValues:  Take a map and extract array of values
//
//	// given input:
//	{"tag.1":"news","tag.2":"sports"}
//
//	mapvalue(match("tag.")) => []string{"1","2"}
type MapValues struct{}

// Type strings aka []string
func (m *MapValues) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *MapValues) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func mapValuesEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// MapInvert:  Take a map and invert key/values
//
//	// given input:
//	tags = {"1":"news","2":"sports"}
//
//	mapinvert(tags) => map[string]string{"news":"1","sports":"2"}
type MapInvert struct{}

// Type MapValue
func (m *MapInvert) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }

func (m *MapInvert) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func mapInvertEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}
