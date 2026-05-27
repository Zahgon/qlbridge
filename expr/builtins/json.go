package builtins

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/value"
)

var _ = u.EMPTY

// JsonPath jmespath json parser http://jmespath.org/
//
//	json_field = `[{"name":"n1","ct":8,"b":true, "tags":["a","b"]},{"name":"n2","ct":10,"b": false, "tags":["a","b"]}]`
//
//	json.jmespath(json_field, "[?name == 'n1'].name | [0]")  =>  "n1"
type JsonPath struct{}

func (m *JsonPath) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *JsonPath) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

// if syntaxError, ok := err.(jmespath.SyntaxError); ok {
// 	u.Warnf("%s\n%s\n", syntaxError, syntaxError.HighlightLocation())
// }

func jsonPathEval(expression string) expr.EvaluatorFunc {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc)
}

// Validate that this is valid json?
