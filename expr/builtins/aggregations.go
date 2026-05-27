package builtins

import (
	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/value"
)

// Avg average of values.  Note, this function DOES NOT persist state doesn't aggregate
// across multiple calls.  That would be responsibility of write context.
//
//	avg(1,2,3) => 2.0, true
//	avg("hello") => math.NaN, false
type Avg struct{}

// Type is NumberType
func (m *Avg) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Avg) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func (m *Avg) IsAgg() bool { _ = "STUB: not implemented"; return false }

func avgEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Sum function to add values. Note, this function DOES NOT persist state doesn't aggregate
// across multiple calls.  That would be responsibility of write context.
//
//	sum(1, 2, 3) => 6
//	sum(1, "horse", 3) => nan, false
type Sum struct{}

// Type is number
func (m *Sum) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *

	// IsAgg yes sum is an agg.
	new(value.ValueType)
}

func (m *Sum) IsAgg() bool { _ = "STUB: not implemented"; return false }

func (m *Sum) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func sumEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// we don't need to evaluate if nil or error

// Do we silently drop, or fail?

// Count Return int value 1 if non-nil/zero.  This should be renamed Increment
// and in general is a horrible, horrible function that needs to be replaced
// with occurrences of value, ignores the value and ensures it is non null
//
//	count(anyvalue)     =>  1, true
//	count(not_number)   =>  -- 0, false
type Count struct{}

// Type is Integer
func (m *Count) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Count) IsAgg() bool           { _ = "STUB: not implemented"; return false }

func (m *Count) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func incrementEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}
