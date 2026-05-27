package exec

import (
	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/plan"
	"github.com/araddon/qlbridge/rel"
)

// Where execution of A filter to implement where clause
type Where struct {
	*TaskBase
	filter expr.Node
	sel    *rel.SqlSelect
}

// NewWhere create new Where Clause
//
//	filters vs final differ bc the Final does final column aliasing
func NewWhere(ctx *plan.Context, p *plan.Where) *Where { _ = "STUB: not implemented"; return nil }

func NewWhereFinal(ctx *plan.Context, p *plan.Where) *Where { _ = "STUB: not implemented"; return nil }

// for _, col := range p.Stmt.Columns {
// 	_, right, _ := col.LeftRight()
// 	u.Debugf("p.Stmt col: %s %#v", right, col)
// }

//u.Debugf("cols: %v", from.Columns)
//u.Infof("source: %#v", from.Source)

//u.Debugf("col: %s %#v", right, col)

//u.Debugf("found where columns: %d", len(cols))

// NewWhereFilter filters vs final differ bc the Final does final column aliasing
func NewWhereFilter(ctx *plan.Context, sql *rel.SqlSelect) *Where {
	_ = "STUB: not implemented"
	return nil
}

// NewHaving Filter
func NewHaving(ctx *plan.Context, p *plan.Having) *Where { _ = "STUB: not implemented"; return nil }

func whereFilter(filter expr.Node, task TaskRunner, cols map[string]int) MessageHandler {
	_ = "STUB: not implemented"
	return *

	//u.Debugf("prepare filter %s", filter)
	new(MessageHandler)
}

//u.Debugf("WHERE:  T:%T  body%#v", msg, msg.Body())

//u.Debugf("WHERE:  T:%T  vals:%#v", msg, mt.Vals)
//u.Debugf("cols:  %#v", cols)

//u.Debugf("WHERE: result:%v T:%T  \n\trow:%#v \n\tvals:%#v", filterValue, msg, mt, mt.Values())
//u.Debugf("cols:  %#v", cols)

//u.Debugf("msg: %#v", msgReader)
//u.Infof("evaluating: ok?%v  result=%v filter expr: '%s'", ok, filterValue.ToString(), filter.String())

//u.Debugf("Filtering out: T:%T   v:%#v", valTyped, valTyped)

//u.Debugf("about to send from where to forward: %#v", msg)
