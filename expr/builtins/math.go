package builtins

import (
	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/value"
)

// Sqrt square root function.  Must be able to coerce to number.
//
//	sqrt(4)            =>  2, true
//	sqrt(9)            =>  3, true
//	sqrt(not_number)   =>  0, false
type Sqrt struct{}

// Type is NumberType
func (m *Sqrt) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *

	// Validate Must have 1 arg
	new(value.ValueType)
}

func (m *Sqrt) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func sqrtEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Pow exponents, raise x to the power of y
//
//	pow(5,2)            =>  25, true
//	pow(3,2)            =>  9, true
//	pow(not_number,2)   =>  NilNumber, false
type Pow struct{}

// Type is Number
func (m *Pow) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *

	// Must have 2 arguments, both must be able to be coerced to Number
	new(value.ValueType)
}

func (m *Pow) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func powerEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}
