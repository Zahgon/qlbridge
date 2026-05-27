package vm

import (
	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/rel"
)

var (
	// a static nil includer whose job is to return errors
	// for vm's that don't have an includer
	noIncluder = &expr.IncludeContext{}
)

type filterql struct {
	expr.EvalContext
	expr.Includer
}

// EvalFilerSelect evaluates a FilterSelect statement from read, into write context
//
// @writeContext = Write results of projection
// @readContext  = Message input, ie evaluate for Where/Filter clause
func EvalFilterSelect(sel *rel.FilterSelect, writeContext expr.ContextWriter, readContext expr.EvalContext) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

// Check and see if we are where Guarded, which would discard the entire message

// filter out this col

// Matches executes a FilterQL statement against an evaluation context
// returning true if the context matches.
func MatchesInc(inc expr.Includer, cr expr.EvalContext, stmt *rel.FilterStatement) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

// Matches executes a FilterQL statement against an evaluation context
// returning true if the context matches.
func Matches(cr expr.EvalContext, stmt *rel.FilterStatement) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

// MatchesExpr executes a expr.Node expression against an evaluation context
// returning true if the context matches.
func MatchesExpr(cr expr.EvalContext, node expr.Node) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func matchesExpr(cr expr.EvalContext, n expr.Node, depth int) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}
