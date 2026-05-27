package builtins

import (
	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/value"
)

// Contains does first arg string contain 2nd arg?
//
//	contains("alabama","red") => false
type Contains struct{}

// Type is Bool
func (m *Contains) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Contains) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func containsEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// TODO:  this should be false, false?
//        need to ensure doesn't break downstream

// LowerCase take a string and lowercase it. must be able to convert to string.
//
//	string.lowercase("HELLO") => "hello", true
type LowerCase struct{}

// Type string
func (m *LowerCase) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }

func (m *LowerCase) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func lowerCaseEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// UpperCase take a string and uppercase it. must be able to convert to string.
//
//	string.uppercase("hello") => "HELLO", true
type UpperCase struct{}

// Type string
func (m *UpperCase) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }

func (m *UpperCase) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func upperCaseEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// TitleCase take a string and uppercase it. must be able to convert to string.
//
//	string.uppercase("hello") => "HELLO", true
type TitleCase struct{}

// Type string
func (m *TitleCase) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }

func (m *TitleCase) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func titleCaseEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Split a string with given separator
//
//	split("apples,oranges", ",") => []string{"apples","oranges"}
type Split struct{}

// Type is Strings
func (m *Split) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Split) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func splitEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// StringIndex a string, removing leading/trailing whitespace
//
//	string.index("apples, oranges ", ",") => 6
//	string.index("apples, oranges ", "X") => -1, false
type StringIndex struct{}

func (m *StringIndex) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}
func (m *StringIndex) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func stringIndexEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// SubString from a given string, use integers to describe the start, [stop]
// of substring to extract.
//
//	string.substr("apples, oranges ", 0, 3) => "app", true
//	string.substr("apple", 3)               => "le", true
//	string.substr("apple", 30, 500)         => nil, false
type SubString struct{}

func (m *SubString) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *SubString) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func subStringEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Strip a string, removing leading/trailing whitespace
//
//	strip(split("apples, oranges ",",")) => {"apples", "oranges"}
//	strip("apples ")                     => "apples"
type Strip struct{}

// type is Unknown (string, or []string)
func (m *Strip) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Strip) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func stripEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Replace a string(s).  Replace occurences of 2nd arg In first with 3rd.
// 3rd arg "what to replace with" is optional
//
//	replace("/blog/index.html", "/blog","")  =>  /index.html
//	replace("/blog/index.html", "/blog")  =>  /index.html
//	replace("/blog/index.html", "/blog/archive/","/blog")  =>  /blog/index.html
//	replace(item, "M")
type Replace struct{}

func (m *Replace) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Replace) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func replaceEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Join items together (string concatenation)
//
//	join("apples","oranges",",")   => "apples,oranges"
//	join(["apples","oranges"],",") => "apples,oranges"
//	join("apples","oranges","")    => "applesoranges"
type Join struct{}

// Type is string
func (m *Join) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Join) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func joinEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// HasPrefix string evaluation to see if string begins with
//
//	hasprefix("apples","ap")   => true
//	hasprefix("apples","o")   => false
type HasPrefix struct{}

// Type bool
func (m *HasPrefix) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *HasPrefix) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func hasPrefixEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// HasSuffix string evaluation to see if string ends with
//
//	hassuffix("apples","es")   => true
//	hassuffix("apples","e")   => false
type HasSuffix struct{}

// Type bool
func (m *HasSuffix) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *HasSuffix) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func hasSuffixEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}
