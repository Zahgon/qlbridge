// VM implements the virtual machine runtime evaluator
// for the SQL, FilterQL, and Expression evalutors.
package vm

import (
	"fmt"
	"time"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/lex"
	"github.com/araddon/qlbridge/value"
)

var (
	// MaxDepth acts as a guard against potentially recursive queries
	MaxDepth = 1000
	// ErrMaxDepth If we hit max depth on recursion
	ErrMaxDepth = fmt.Errorf("Recursive Evaluation Error")
	// ErrUnknownOp an unrecognized Operator in expression
	ErrUnknownOp = fmt.Errorf("expr: unknown op type")
	// ErrUnknownNodeType Unhandled Node type for expression evaluation
	ErrUnknownNodeType = fmt.Errorf("expr: unknown node type")
	// ErrExecute could not evaluate an expression
	ErrExecute = fmt.Errorf("Could not execute")
)

// EvalBaseContext base context for expression evaluation
type EvalBaseContext struct {
	expr.EvalContext
}

// Eval - Evaluate the given expression (arg Node) against the given context.
// @ctx is the evaluation context ie the variables/values which the expression will be
// evaluated against.  It may be a simple reader of  message/data or any
// object whhich implements EvalContext.
func Eval(ctx expr.EvalContext, arg expr.Node) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// creates a new Value with a nil group and given value.
func numberNodeToValue(t *expr.NumberNode) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// ResolveIncludes take an expression and resolve any includes so that
// it does not have to be resolved at runtime.  There is also a
// InlineIncludes alternative in expr pkg which actually re-writes the expression
// to remove includes and embed the expressions they refer to as part of this expression.
func ResolveIncludes(ctx expr.Includer, arg expr.Node) error { _ = "STUB: not implemented"; return nil }

func resolveIncludesDepth(ctx expr.Includer, arg expr.Node, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

// can we switch to arg.Type()

func evalBool(ctx expr.EvalContext, arg expr.Node, depth int) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func evalDepth(ctx expr.EvalContext, arg expr.Node, depth int) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// WHERE (`users.user_id` != NULL)

func resolveInclude(ctx expr.Includer, inc *expr.IncludeNode, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func walkInclude(ctx expr.EvalContext, inc *expr.IncludeNode, depth int) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

func walkBoolean(ctx expr.EvalContext, n *expr.BooleanNode, depth int) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// one of the expressions in an OR clause matched, shortcircuit true

// one of the expressions in an AND clause did not match, shortcircuit false

// no shortcircuiting, if and=true this means all expressions returned true...
// ...if and=false (OR) this means all expressions returned false.

// Binary operands:   =, ==, !=, OR, AND, >, <, >=, <=, LIKE, contains
//
//	x == y,   x = y
//	x != y
//	x OR y
//	x > y
//	x < =
func walkBinary(ctx expr.EvalContext, node *expr.BinaryNode, depth int) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

func evalBinary(ctx expr.EvalContext, node *expr.BinaryNode, depth int) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// If we could not evaluate either we can shortcut

// We don't alllow nil == nil here bc we have a NilValue type
// that we would use for that

// Else if we can only evaluate right

// Else if we can only evaluate one, we can short circuit as well

// they are technically not equal?

// Try int first

// Fallback to float

// Try int first

// Fallback to float

// Nice, both strings

// a(value) LIKE b(pattern)

// Should we evaluate strings that are non-nil to be = true?

// Lets look at first arg, all in slice must be of same type

// [x,y,z] contains str

// [a,b,c] contains int

// [x,y,z] LIKE str

// [x,y,z] contains str

// [x,y,z] LIKE str

// does nil==nil  = true ??

func walkIdentity(ctx expr.EvalContext, node *expr.IdentityNode) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

func walkUnary(ctx expr.EvalContext, node *expr.UnaryNode, depth int) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// walkTernary ternary evaluator
//
//	A   BETWEEN   B  AND C
func walkTernary(ctx expr.EvalContext, node *expr.TriNode, depth int) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// walkArray Array evaluator:  evaluate multiple values into an array
//
//	(b,c,d)
func walkArray(ctx expr.EvalContext, node *expr.ArrayNode, depth int) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// we are returning an array of evaluated nodes

// walkFunc evaluates a function
func walkFunc(ctx expr.EvalContext, node *expr.FuncNode, depth int) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

func operateNumbers(op lex.Token, av, bv value.NumberValue) value.Value {
	_ = "STUB: not implemented"
	return *new(value.Value)
}

// +

// *

// -

//

//    %
// is this even valid?   modulus on floats?

// Below here are Boolean Returns
//  ==

//  >

//r = 1

//r = 0

//  !=    or <>

// <

// >=

// <=

//  ||

//  &&

func operateStrings(op lex.Token, av, bv value.StringValue) value.Value {
	_ = "STUB: not implemented"

	//  Any other ops besides =, ==, !=, contains, like?
	return *new(value.Value)
}

//  ==

//  !=

// a(value) LIKE b(pattern)

func operateTime(op lex.TokenType, lht, rht time.Time) (value.BoolValue, bool) {
	_ = "STUB: not implemented"
	return *new(value.BoolValue), false
}

// lhexpr > rhexpr

// lhexpr >= rhexpr

// lhexpr < rhexpr

// lhexpr <= rhexpr

// LikeCompare takes two strings and evaluates them for like equality
func LikeCompare(a, b string) (value.BoolValue, bool) {
	_ = "STUB: not implemented"
	// Do we want to always do this replacement?   Or do this at parse time or config?
	return *new(value.BoolValue), false
}

func operateInts(op lex.Token, av, bv value.IntValue) value.Value {
	_ = "STUB: not implemented"
	return *new(value.Value)
}

func operateIntVals(op lex.Token, a, b int64) (value.Value, error) {
	_ = "STUB: not implemented"
	return *new(value.Value), nil
}

// +
//r = a + b

// *
//r = a * b

// -
//r = a - b

//    /
//r = a / b

//    %
//r = a / b

// Below here are Boolean Returns
//  ==, =

//  >

//  !=    or <>

// <

// >=

// <=

//  ||

//  &&

func uoperate(op string, a float64) (r float64) { _ = "STUB: not implemented"; return 0 }
