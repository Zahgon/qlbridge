package expr

import (
	"sync"

	"github.com/araddon/qlbridge/value"
)

var (
	// The global function registry
	funcReg = NewFuncRegistry()
)

type (
	// EvaluatorFunc defines the evaluator func which may be stateful (or not) for
	// evaluating custom functions
	EvaluatorFunc func(ctx EvalContext, args []value.Value) (value.Value, bool)
	// CustomFunc allows custom functions to be added for run-time evaluation
	CustomFunc interface {
		// Type Define the Return Type of this function, or use value.Value for unknown.
		Type() value.ValueType
		// Validate is parse time syntax and type evaluation.  Also returns the evaluation
		// function.
		Validate(n *FuncNode) (EvaluatorFunc, error)
	}
	// AggFunc allows custom functions to specify if they provide aggregation
	AggFunc interface {
		IsAgg() bool
	}
	// FuncResolver is a function resolution interface that allows
	// local/namespaced function resolution.
	FuncResolver interface {
		FuncGet(name string) (Func, bool)
	}

	// FuncRegistry contains lists of functions for different scope/run-time evaluation contexts.
	FuncRegistry struct {
		mu    sync.RWMutex
		funcs map[string]Func
		aggs  map[string]struct{}
	}
)

// EmptyEvalFunc a no-op evaluation function for use in
func EmptyEvalFunc(ctx EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// NewFuncRegistry create a new function registry. By default their is a
// global one, but you can have local function registries as well.
func NewFuncRegistry() *FuncRegistry { _ = "STUB: not implemented"; return nil }

// Add a name/function to registry
func (m *FuncRegistry) Add(name string, fn CustomFunc) { _ = "STUB: not implemented"; return }

// FuncGet gets a function from registry if it exists.
func (m *FuncRegistry) FuncGet(name string) (Func, bool) {
	_ = "STUB: not implemented"
	return *new(Func), false
}

// FuncAdd Global add Functions to the VM func registry occurs here.
func FuncAdd(name string, fn CustomFunc) { _ = "STUB: not implemented"; return }
