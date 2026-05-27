package builtins

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/value"
)

var _ = u.EMPTY

// Not urnary negation function
//
//	not(eq(5,5)) => false, true
//	not(eq("false")) => false, true
type Not struct{}

// Type bool
func (m *Not) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }

func (m *Not) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func notEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Equal function?  returns true if items are equal
//
//	// given context   {"name":"wil","event":"stuff", "int4": 4}
//
//	eq(int4,5)  => false
type Eq struct{}

// Type bool
func (m *Eq) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }

func (m *Eq) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func equalEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Ne Not Equal function?  returns true if items are equal
//
//	// given   {"5s":"5","item4":4,"item4s":"4"}
//
//	ne(`5s`,5) => true, true
//	ne(`not_a_field`,5) => false, true
//	ne(`item4s`,5) => false, true
//	ne(`item4`,5) => false, true
type Ne struct{}

// Type bool
func (m *Ne) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Ne) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func notEqualEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Gt GreaterThan is left hand > right hand.
// Must be able to convert items to Floats.
//
//	gt(5,6)  => true, true
type Gt struct{}

// Type bool
func (m *Gt) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Gt) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func greaterThanEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Ge GreaterThan or Equal func. Must be able to convert items to Floats.
type Ge struct{}

// Type bool
func (m *Ge) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Ge) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func greatherThanOrEqualEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Le Less Than or Equal. Must be able to convert items to Floats.
type Le struct{}

// Type bool
func (m *Le) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Le) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func lessThanOrEqualEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Lt Less Than Must be able to convert items to Floats
//
//	lt(5, 6)  => true
type Lt struct{}

// Type bool
func (m *Lt) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }

func (m *Lt) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func lessThanEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Exists Answers True/False if the field exists and is non null
//
//	exists(real_field) => true
//	exists("value") => true
//	exists("") => false
//	exists(empty_field) => false
//	exists(2) => true
//	exists(todate(date_field)) => true
type Exists struct{}

// Type bool
func (m *Exists) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Exists) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func existsEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// case *expr.IdentityNode:
// 	_, ok := ctx.Get(node.Text)
// 	if ok {
// 		return value.BoolValueTrue, true
// 	}
// 	return value.BoolValueFalse, true
// case *expr.StringNode:
// 	_, ok := ctx.Get(node.Text)
// 	if ok {
// 		return value.BoolValueTrue, true
// 	}
// 	return value.BoolValueFalse, true

// Any Answers True/False if any of the arguments evaluate to truish (javascripty)
// type definintion of true
//
// Rules for if True:
//
//	int != 0
//	string != ""
//	boolean natively supported true/false
//	time != time.IsZero()
//
// Examples:
//
//	any(item,item2)  => true, true
//	any(not_field)   => false, true
type Any struct{}

// Type bool
func (m *Any) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }

func (m *Any) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func anyEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// continue

// All Answers True/False if all of the arguments evaluate to truish (javascripty)
// type definintion of true.  Non-Nil, non-Error, values.
//
// Rules for if True:
//
//	int != 0
//	string != ""
//	boolean natively supported true/false
//	time != time.IsZero()
//
// Examples:
//
//	all("hello",2, true) => true
//	all("hello",0,true)  => false
//	all("",2, true)      => false
type All struct{}

func allEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

func (m *All) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

// Type is BoolType for All function
func (m *All) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
