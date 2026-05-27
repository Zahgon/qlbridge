package builtins

import (
	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/value"
)

// OneOf choose the first non-nil, non-zero, non-false fields
//
//	oneof(nil, 0, "hello") => 'hello'
type OneOf struct{}

// Type unknown
func (m *OneOf) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *OneOf) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func oneOfEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// FilterFromArgs given set of values
func FiltersFromArgs(filterVals []value.Value) []string { _ = "STUB: not implemented"; return nil }

// Filter Filter OUT Values that match specified list of match filter criteria
//
// Operates on MapValue (map[string]interface{}), StringsValue ([]string), or string
// takes N Filter Criteria
// supports Matching:      "filter**" // matches  "filter_x", "filterstuff"
//
// Filter a map of values by key to remove certain keys
//
//	filter(match("topic_"),key_to_filter, key2_to_filter)  => {"goodkey": 22}, true
//
// Filter out VALUES (not keys) from a list of []string{} for a specific value
//
//	filter(split("apples,oranges",","),"ora*")  => ["apples"], true
//
// Filter out values for single strings
//
//	filter("apples","app*")      => []string{}, true
type Filter struct{}

// Type unknown
func (m *Filter) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }

func (m *Filter) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func FilterEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// FilterMatch  Filter IN Values that match specified list of match filter criteria
//
// Operates on MapValue (map[string]interface{}), StringsValue ([]string), or string
// takes N Filter Criteria
//
// Wildcard Matching:      "abcd*" // matches  "abcd_x", "abcdstuff"
//
// Filter a map of values by key to only keep certain keys
//
//	filtermatch(match("topic_"),key_to_filter, key2_to_filter)  => {"goodkey": 22}, true
//
// Filter in VALUES (not keys) from a list of []string{} for a specific value
//
//	filtermatch(split("apples,oranges",","),"ora*")  => ["oranges"], true
//
// Filter in values for single strings
//
//	filtermatch("apples","app*")      => []string{"apple"}, true
type FilterMatch struct{}

// Type Unknown
func (m *FilterMatch) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}

func (m *FilterMatch) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func FilterMatchEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

//u.Debugf("FilterIn():  %T:%v   filters:%v", val, val, filters)
