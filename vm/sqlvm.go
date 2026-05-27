package vm

import (
	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/rel"
)

// EvalSql Is a partial SQL statement evaluator (that doesn't get it all right).  See
// exec package for full sql evaluator.  Can be used to evaluate a read context and write
// results to write context.  This does not project columns prior to running WHERE.
//
// @writeContext = Write out results of projection
// @readContext  = Message to evaluate does it match where clause?  if so proceed to projection
func EvalSql(sel *rel.SqlSelect, writeContext expr.ContextWriter, readContext expr.EvalContext) (bool, error) {
	_ = "STUB: not implemented"

	// Check and see if we are where Guarded, which would discard the entire message
	return false, nil
}

// ok, continue

// filter out this col

// filter out

// Write out the result of the evaluation
