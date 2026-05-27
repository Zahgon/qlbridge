package builtins

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/value"
)

var _ = u.EMPTY

// ToString cast as string.  must be able to convert to string
type ToString struct{}

// Type string
func (m *ToString) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *ToString) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func toStringEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Cast type coercion, cast to an explicit type.
//
//	cast(identity AS <type>) => 5.0
//	cast(reg_date AS string) => "2014/01/12"
//
// Types:  [char, string, int, float]
type Cast struct{}

// Type one of value types
func (m *Cast) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Cast) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func castEvalNoAs(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// http://www.cheatography.com/davechild/cheat-sheets/mysql/

func castEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"

	// identity AS identity
	//  0        1    2
	return *new(value.Value), false
}

// This is enforced by parser, so no need
// if vals[2] == nil || vals[2].Nil() || vals[2].Err() {
// 	return nil, false
// }

// http://www.cheatography.com/davechild/cheat-sheets/mysql/

// ToBool cast as boolean
type ToBool struct{}

// Type bool
func (m *ToBool) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *ToBool) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func toBoolEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// ToInt Convert to Integer:   Best attempt at converting to integer.
//
//	toint("5")          => 5, true
//	toint("5.75")       => 5, true
//	toint("5,555")      => 5555, true
//	toint("$5")         => 5, true
//	toint("5,555.00")   => 5555, true
type ToInt struct{}

// Type integer
func (m *ToInt) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *ToInt) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func toIntEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Milliseconds

// ToNumber Convert to Number:   Best attempt at converting to integer
//
//	tonumber("5") => 5.0
//	tonumber("5.75") => 5.75
//	tonumber("5,555") => 5555
//	tonumber("$5") => 5.00
//	tonumber("5,555.00") => 5555
type ToNumber struct{}

// Type number
func (m *ToNumber) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *ToNumber) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func toNumberEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Unsign converts a signed int to an unsigned int represented as a string
// for positive numbers this should simply convert the number to a string
// for negative numbers, converting to a uint adds the negative value to the
// max uint value (18446744073709551615).
//
//	unsign(-32847623329847) => 18446711226086221769
//	unsign(876) => 876
//	unsign("-70") => 18446744073709551546
type Unsign struct{}

// Type number
func (u *Unsign) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (u *Unsign) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func unsignEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}
