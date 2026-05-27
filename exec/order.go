package exec

import (
	"github.com/araddon/qlbridge/datasource"
	"github.com/araddon/qlbridge/plan"
)

// Order
type Order struct {
	*TaskBase
	p          *plan.Order
	complete   chan bool
	closed     bool
	isComplete bool
}

// NewORder create new order by exec task
func NewOrder(ctx *plan.Context, p *plan.Order) *Order { _ = "STUB: not implemented"; return nil }

func (m *Order) Close() error { _ = "STUB: not implemented"; return nil }

// what should this be?

//u.Infof("%p group by final Close() waiting for complete", m)

//u.Warnf("%p got groupbyfinal complete", m)

func (m *Order) Run() error { _ = "STUB: not implemented"; return nil }

// are are going to hold entire row in memory while we are calculating
//  so obviously not scalable.

//u.Debugf("NICE, got closed channel shutdown")

// We are going to use VM Engine to create a value for each statement in group by
//  then join each value together to create a unique key.

//u.Debugf("msgtype:%T  key:%q for-expr:%s", sdm, key, col.Expr)

// Is this an error?
//u.Warnf("no key?  %s for %+v", col.Expr, sdm)

//u.Warnf("no col.expr? %#v", col)

//u.Infof("found key:%s for %+v", key, sdm)

//u.Debugf("got %s:%v msgs", key, vals)

type msgkey struct {
	keys []string
	msg  *datasource.SqlDriverMessageMap
}
type OrderMessages struct {
	l      []*msgkey
	invert []bool
}

func NewOrderMessages(p *plan.Order) *OrderMessages { _ = "STUB: not implemented"; return nil }

//u.Debugf("invert?  %s ORDER %v", col.Expr, col.Order)

func (m *OrderMessages) Len() int { _ = "STUB: not implemented"; return 0 }

func (m *OrderMessages) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (m *OrderMessages) Swap(i, j int) { _ = "STUB: not implemented"; return }
