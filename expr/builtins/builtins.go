// Builtin functions are a library of functions natively available in
// qlbridge expression evaluation although adding your own is easy.
package builtins

import (
	"sync"

	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/value"
)

var _ = u.EMPTY
var loadOnce sync.Once

const yymmTimeLayout = "0601"

func LoadAllBuiltins() { _ = "STUB: not implemented"; return }

// math

// aggregate ops

// logical

// Map

// Date/Time functions

// Casting and Type Coercion

// String Functions

// array, string

// selection

// special functions

// Hashing functions

// json

// MySQL Builtins

// uuid generates a new uuid
//
//	uuid() =>  "...."
type UuidGenerate struct{}

// Type string
func (m *UuidGenerate) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}
func (m *UuidGenerate) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func uuidGenerateEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}
