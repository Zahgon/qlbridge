package vm

import (
	"time"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/lex"
	"github.com/araddon/qlbridge/value"
)

// DateConverter can help inspect a boolean expression to determine if there is
// date-math in it.  If there is datemath, can calculate the time boundary
// where the expression may possibly change from true to false.
// - Must be boolean expression
// - Only calculates the first boundary
// - Only calculates POSSIBLE boundary, given complex logic (ors etc) may NOT change.
type DateConverter struct {
	HasDateMath bool      // Does this have date math in it all?
	Node        expr.Node // The expression we are extracting datemath from
	TimeStrings []string  // List of each extracted timestring
	bt          time.Time // The possible boundary time when expression flips true/false
	at          time.Time // The Time to use as "now" or reference point
	ctx         expr.EvalIncludeContext
	err         error
}

// NewDateConverter takes a node expression
func NewDateConverter(ctx expr.EvalIncludeContext, n expr.Node) (*DateConverter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DateConverter) addBoundary(bt time.Time) { _ = "STUB: not implemented"; return }

func (d *DateConverter) addValue(lhv value.Value, op lex.TokenType, val string) {
	_ = "STUB: not implemented"
	return
}

// Given Anchor Time At calculate Relative Time Rt

// Ct = Comparison time, left hand side of expression
// At = Anchor Time
// Rt = Relative time result of Anchor Time offset by datemath "now-3d"
// Bt = Boundary time = calculated time at which expression will change boolean expression value

// none of these are supported operators for finding boundary

// 1) ----------- Ct --------------     Rt < Ct
//        Rt                            Ct > "now+-1d" = true but will be true when at + (ct - rt)
//         ------Bt
//
// 2) ------------- Ct ------------     Ct < Rt
//                        Rt            Ct > "now+-1d" = false, and will always be false
//

// Is false, and always will be false no candidates

// 3) ------ Ct -------------------     Ct < Rt
//              Rt                      Ct < "now+-1d" = true (and always will be)
//
// 4) ----------- Ct --------------     Rt < Ct
//     At----Rt                         Ct < "now+-1d" = true, but will be in true when at + (ct - rt)
//           Bt---|
//

// Is true, and always will be true no candidates

// Boundary given all the date-maths in this node find the boundary time where
// this expression possibly will change boolean value.
// If no boundaries exist, returns time.Time{} (zero time)
func (d *DateConverter) Boundary() time.Time {
	_ = "STUB: not implemented"

	// Determine if this expression node uses datemath (ie, "now-4h")
	return *new(time.Time)
}

func (d *DateConverter) findDateMath(node expr.Node) { _ = "STUB: not implemented"; return }

// If left side is datemath   "now-3d" < ident then re-write to have ident on left

// Reverse equation to put identity on left side
// "now-1d" < last_visit    =>   "last_visit" > "now-1d"

// lex.TokenEqual, lex.TokenEqualEqual, lex.TokenNE:
// none of these are supported operators for finding boundary

// Scalar/Literal values cannot be datemath, must be binary-expression
