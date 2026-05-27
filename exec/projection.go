package exec

import (
	"github.com/araddon/qlbridge/plan"
)

// Projection Execution Task
type Projection struct {
	*TaskBase
	closed bool
	p      *plan.Projection
}

// In Process projections are used when mapping multiple sources together
// and additional columns such as those used in Where, GroupBy etc are used
// even if they will not be used in Final projection
func NewProjection(ctx *plan.Context, p *plan.Projection) *Projection {
	_ = "STUB: not implemented"
	return nil
}

// In Process projections are used when mapping multiple sources together
//
//	and additional columns such as those used in Where, GroupBy etc are used
//	even if they will not be used in Final projection
func NewProjectionInProcess(ctx *plan.Context, p *plan.Projection) *Projection {
	_ = "STUB: not implemented"
	return nil
}

// Final Projections project final select columns for result-writing
func NewProjectionFinal(ctx *plan.Context, p *plan.Projection) *Projection {
	_ = "STUB: not implemented"
	return nil
}

// NewProjectionLimit Only provides counting/limit projection
func NewProjectionLimit(ctx *plan.Context, p *plan.Projection) *Projection {
	_ = "STUB: not implemented"
	return nil
}

func (m *Projection) drain() { _ = "STUB: not implemented"; return }

//u.Debugf("%p dropping msg %v", msg)

// Close cleans up and closes channels
func (m *Projection) Close() error {
	_ = "STUB: not implemented"
	// u.Debugf("Projection Close  alreadyclosed?%v", m.closed)
	return nil
}

//return m.TaskBase.Close()

// CloseFinal after exit, cleanup some more
func (m *Projection) CloseFinal() error {
	_ = "STUB: not implemented"
	// u.Debugf("Projection CloseFinal  alreadyclosed?%v", m.closed)
	return nil
}

//return nil

// Create handler function for evaluation (ie, field selection from tuples)
func (m *Projection) projectionEvaluator(isFinal bool) MessageHandler {
	_ = "STUB: not implemented"
	return *new(MessageHandler)
}

// If we have a projection, use that as col count

//u.Infof("got projection message: %T %#v", msg, msg.Body())

// use our custom write context for example purposes

//u.Debugf("about to project: %#v", mt)

//u.Debugf("%d  colidx:%v sidx: %v pidx:%v key:%q Expr:%v", colIdx, col.Index, col.SourceIndex, col.ParentIndex, col.Key(), col.Expr)

// Most likely scenario here is Missing Columns.
// Unlikely traditional sql, we are going to operate in both strict-schema mode
// which would error, and sparse which will not, more like no-sql.

//return fmt.Errorf("Could not evaluate if clause: %v", col.Guard.String())

//u.Debugf("if eval val:  %T:%v", ifColValue, ifColValue)

//u.Debugf("Filtering out col")

//u.Infof("star row: %#v", starRow)

//   select *, myvar, 1

//writeContext.Put(&expr.Column{As: k}, nil, value.NewValue(v))

//   select * FROM Z

//writeContext.Put(&expr.Column{As: k}, nil, value.NewValue(v))
//u.Infof("colct: %v   v:%v", colIdx, v)

// for k, v := range ctx.Session.Row() {
// 	u.Infof("%p session? %s: %v", ctx.Session, k, v.Value())
// }

//u.Debugf("%#v", col)
//u.Debugf("evaled nil? key=%v  val=%v expr:%s", col.Key(), v, col.Expr.String())
//writeContext.Put(col, mt, v)
//u.Infof("mt: %T  mt %#v", mt, mt)
//v.Value()

//u.Debugf("%d:%d row:%d evaled: %v  val=%v", colIdx, colCt, len(row), col, v.Value())
//writeContext.Put(col, mt, v)

//u.Infof("row: %#v", row)
//u.Infof("row cols: %v", colIndex)

//u.Warnf("nice, got context reader? %T", mt)

//u.Debugf("about to project: %#v", mt)

//u.Debugf("col: idx:%v sidx: %v pidx:%v key:%v   %s", col.Index, col.SourceIndex, col.ParentIndex, col.Key(), col.Expr)

//return fmt.Errorf("Could not evaluate if clause: %v", col.Guard.String())

//u.Debugf("if eval val:  %T:%v", ifColValue, ifColValue)

//u.Debugf("Filtering out col")

//writeContext.Put(&expr.Column{As: k}, nil, value.NewValue(v))

//u.Warnf("failed eval key=%v  val=%#v expr:%s   mt:%#v", col.Key(), v, col.Expr, mt.Row())

//u.Debugf("%#v", col)
//u.Debugf("evaled nil? key=%v  val=%v expr:%s", col.Key(), v, col.Expr.String())
//writeContext.Put(col, mt, v)
//u.Infof("mt: %T  mt %#v", mt, mt)
//v.Value()

//u.Debugf("evaled: key=%v  val=%v", col.Key(), v.Value())
//writeContext.Put(col, mt, v)

//u.Infof("row: %#v cols:%#v", row, colIndex)
//u.Infof("row cols: %v", colIndex)

//u.Debugf("%p Projection reaching Limit!!! rowct:%v  limit:%v", m, rowCt, limit)
// Sending nil message is a message to downstream to shutdown
// should close rest of dag as well

//u.Debugf("row:%d  completed projection for: %p %#v", rowCt, out, outMsg)

// Limit only evaluator
func (m *Projection) limitEvaluator() MessageHandler {
	_ = "STUB: not implemented"
	return *new(MessageHandler)
}

//u.Debugf("%p Projection reaching Limit!!! rowct:%v  limit:%v", m, rowCt, limit)
// Sending nil message is a message to downstream to shutdown
//m.Close()

//return false
// swallow it
