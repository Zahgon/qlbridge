package exec

import (
	"encoding/gob"

	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/plan"
	"github.com/araddon/qlbridge/rel"
	"github.com/araddon/qlbridge/value"
)

var (
	_ = u.EMPTY

	// Ensure that we implement the Task Runner interface
	_ TaskRunner = (*GroupBy)(nil)
)

func init() {
	gob.Register(AggPartial{})
}

// Group by a Sql Group By task which creates a hashable key from row
// commposed of key = {each,value,of,column,in,groupby}
//
// A very stupid naive parallel groupby holds values in memory.  This
// is a toy implementation that is only useful for small cardinality
// group-bys, small number of rows.
type GroupBy struct {
	*TaskBase
	closed bool
	p      *plan.GroupBy
}

func NewGroupBy(ctx *plan.Context, p *plan.GroupBy) *GroupBy { _ = "STUB: not implemented"; return nil }

// GroupByFinal a Sql Group By Operator finalizer for partials.  IE, if group
// by is a distributed task, then this is the reducer for sub-tasks.
type GroupByFinal struct {
	*TaskBase
	p          *plan.GroupBy
	complete   chan bool
	closed     bool
	isComplete bool
}

// NewGroupByFinal creates the group-by-finalizer task.
func NewGroupByFinal(ctx *plan.Context, p *plan.GroupBy) *GroupByFinal {
	_ = "STUB: not implemented"
	return nil
}

// Run runs this group by tasks, standard task interface.
func (m *GroupBy) Run() error { _ = "STUB: not implemented"; return nil }

// are are going to hold entire row in memory while we are calculating
//  so obviously not scalable.

// We are going to use VM Engine to create a value for each statement in group by
// then join each value together to create a unique key.

//u.Debugf("got %s:%v msgs", k, len(v))

//u.Debugf("col: idx:%v sidx: %v pidx:%v key:%v   %s", col.Index, col.SourceIndex, col.ParentIndex, col.Key(), col.Expr)

//u.Infof("mt: %T  mm %#v", mm, mm)

//u.Debugf("evaled nil? key=%v  val=%v expr:%s", col.Key(), v, col.Expr.String())
//u.Infof("mt: %T  mm %#v", mm, mm)

//u.Debugf("evaled: key=%v  val=%v", col.Key(), v.Value())

//u.Debugf("agg result: %#v  %v", row[i], row[i])

// Partial results, append key at end?  shouldn't be able to be fit in message itself?

//u.Debugf("GroupBy output row? key:%s %#v", key, row)

//u.Debugf("row: %v  cols:%v", row, colIndex)

// Run group-by-final Runs standard task interface.
func (m *GroupByFinal) Run() error { _ = "STUB: not implemented"; return nil }

//u.Debugf("GroupByFinal, got closed channel shutdown")

//u.Infof("got gbfinal message %#v", msg)

//u.Infof("found key:%s for %#v", key, mt.Vals)

//u.Debugf("got %s:%v msgs", key, vals)

//u.Debugf("col: idx:%v sidx: %v pidx:%v key:%v   %s", col.Index, col.SourceIndex, col.ParentIndex, col.Key(), col.Expr)

//u.Debugf("evaled: key=%v  val=%v", col.Key(), v.Value())

//u.Debugf("agg result: %#v  %v", row[i], row[i])

//u.Debugf("GroupBy output row? %v", row)

// Close the task, channels, cleanup.
func (m *GroupBy) Close() error { _ = "STUB: not implemented"; return nil }

// Close the task and cleanup.  Trys to wait for the downstream
// reducer tasks to complete after flushing messages.
func (m *GroupByFinal) Close() error { _ = "STUB: not implemented"; return nil }

//u.Infof("%p group by final Close() waiting for complete", m)

//u.Warnf("%p got groupbyfinal complete", m)

// AggPartial is a struct to represent the partial aggregation
// that will be reduced on finalizer.  IE, for consistent-hash based
// group-bys calculated across multiple nodes this holds info that
// needs to be further calculated it only represents this hash.
type AggPartial struct {
	Ct int64
	N  float64
}

type AggFunc func(v value.Value)
type resultFunc func() interface{}
type Aggregator interface {
	Do(v value.Value)
	Result() interface{}
	Reset()
	Merge(*AggPartial)
}
type agg struct {
	do     AggFunc
	result resultFunc
}
type groupByFunc struct {
	last interface{}
}

func (m *groupByFunc) Do(v value.Value)    { _ = "STUB: not implemented"; return }
func (m *groupByFunc) Result() interface{} { _ = "STUB: not implemented"; return nil }
func (m *groupByFunc) Reset()              { _ = "STUB: not implemented"; return }
func (m *groupByFunc) Merge(a *AggPartial) { _ = "STUB: not implemented"; return }
func NewGroupByValue(col *rel.Column) Aggregator {
	_ = "STUB: not implemented"
	return *new(Aggregator)
}

type sum struct {
	partial bool
	ct      int64
	n       float64
}

func (m *sum) Do(v value.Value) { _ = "STUB: not implemented"; return }

func (m *sum) Result() interface{} { _ = "STUB: not implemented"; return nil }

func (m *sum) Reset()              { _ = "STUB: not implemented"; return }
func (m *sum) Merge(a *AggPartial) { _ = "STUB: not implemented"; return }

func NewSum(col *rel.Column, partial bool) Aggregator {
	_ = "STUB: not implemented"
	return *new(Aggregator)
}

type avg struct {
	partial bool
	ct      int64
	n       float64
}

func (m *avg) Do(v value.Value) { _ = "STUB: not implemented"; return }

func (m *avg) Result() interface{} { _ = "STUB: not implemented"; return nil }

func (m *avg) Reset()              { _ = "STUB: not implemented"; return }
func (m *avg) Merge(a *AggPartial) { _ = "STUB: not implemented"; return }

func NewAvg(col *rel.Column, partial bool) Aggregator {
	_ = "STUB: not implemented"
	return *new(Aggregator)
}

type count struct {
	n int64
}

func (m *count) Do(v value.Value) { _ = "STUB: not implemented"; return }

func (m *count) Result() interface{} { _ = "STUB: not implemented"; return nil }

func (m *count) Reset()              { _ = "STUB: not implemented"; return }
func (m *count) Merge(a *AggPartial) { _ = "STUB: not implemented"; return }

func NewCount(col *rel.Column) Aggregator { _ = "STUB: not implemented"; return *new(Aggregator) }

func buildAggs(p *plan.GroupBy) ([]Aggregator, error) { _ = "STUB: not implemented"; return nil, nil }

// simple Non Aggregate Value  gb.As == col.AS
//   SELECT domain, count(*) FROM users GROUP BY domain;

// aliased column
// SELECT `users`.`name` AS usernames FROM `users` GROUP BY `users`.`name`
//   gb.String() == "`users`.`name`"  && col.Expr.String() == "`users`.`name`"

// Since we made it here, it is an aggregate func
//  move to a registry of some kind to allow extension

// TODO:  extract to a UDF Registry Similar to builtins

// expression logic?

// We can have a naked group by which basically means distinct? should have been caught above
