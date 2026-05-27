package exec

import (
	"database/sql/driver"

	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/datasource"
	"github.com/araddon/qlbridge/plan"
	"github.com/araddon/qlbridge/rel"
	"github.com/araddon/qlbridge/schema"
)

var (
	_ = u.EMPTY

	// Ensure that we implement the Task Runner interface
	_ TaskRunner = (*JoinMerge)(nil)
)

type KeyEvaluator func(msg schema.Message) driver.Value

// Evaluate messages to create JoinKey based message, where the
//
//	Join Key (composite of each value in join expr) hashes consistently
type JoinKey struct {
	*TaskBase
	p        *plan.JoinKey
	colIndex map[string]int
}

// A JoinKey task that evaluates the compound JoinKey to allow
//
//	for parallelized join's
//
//	 source1   ->  JoinKey  ->  hash-route
//	                                       \
//	                                        --  join  -->
//	                                       /
//	 source2   ->  JoinKey  ->  hash-route
func NewJoinKey(ctx *plan.Context, p *plan.JoinKey) *JoinKey { _ = "STUB: not implemented"; return nil }

func (m *JoinKey) Run() error { _ = "STUB: not implemented"; return nil }

//u.Debugf("got signal quit")

//u.Debugf("NICE, got msg shutdown")

//u.Infof("In joinkey msg %#v", msg)

//u.Debugf("evaluating: ok?%v T:%T result=%v node '%v'", ok, joinVal, joinVal.ToString(), node.String())

//u.Infof("joinkey: %v row:%v", vals, mt)

// Scans 2 source tasks for rows, evaluate keys, use for join
type JoinMerge struct {
	*TaskBase
	leftStmt  *rel.SqlSource
	rightStmt *rel.SqlSource
	ltask     TaskRunner
	rtask     TaskRunner
	colIndex  map[string]int
}

// A very stupid naive parallel join merge, uses Key() as value to merge
//
//	two different input channels
//
//	source1   ->
//	             \
//	               --  join  -->
//	             /
//	source2   ->
//
// Distributed:
//
//	source1a  ->                |-> --  join  -->
//	source1b  -> key-hash-route |-> --  join  -->  reduce ->
//	source1n  ->                |-> --  join  -->
//	                            |-> --  join  -->
//	source2a  ->                |-> --  join  -->
//	source2b  -> key-hash-route |-> --  join  -->
//	source2n  ->                |-> --  join  -->
func NewJoinNaiveMerge(ctx *plan.Context, l, r TaskRunner, p *plan.JoinMerge) *JoinMerge {
	_ = "STUB: not implemented"
	return nil
}

func (m *JoinMerge) Run() error { _ = "STUB: not implemented"; return nil }

//u.Infof("In source Scanner msg %#v", msg)

//u.Debugf("NICE, got left shutdown")

//u.Infof("In source Scanner iter %#v", item)

//u.Debugf("NICE, got right shutdown")

//u.Info("leaving source scanner")

//u.Debugf("compare:  key:%v  left:%#v  right:%#v  rh: %#v", keyLeft, valLeft, rh[keyLeft], rh)

//u.Debugf("found match?\n\t%d left=%#v\n\t%d right=%#v", len(valLeft), valLeft, len(valRight), valRight)

//u.Debugf("msgsct: %v   msgs:%#v", len(msgs), msgs)

//outCh <- datasource.NewUrlValuesMsg(i, msg)
//u.Debugf("i:%d   msg:%#v", i, msg)

func (m *JoinMerge) mergeValueMessages(lmsgs, rmsgs []*datasource.SqlDriverMessageMap) []*datasource.SqlDriverMessageMap {
	_ = "STUB: not implemented"
	// m.leftStmt.Columns, m.rightStmt.Columns, nil
	//func mergeValuesMsgs(lmsgs, rmsgs []datasource.Message, lcols, rcols []*rel.Column, cols map[string]*rel.Column) []*datasource.SqlDriverMessageMap {
	return nil
}

//u.Infof("merge values: %v:%v", len(lcols), len(rcols))

//u.Warnf("nice SqlDriverMessageMap: %#v", lmt)

//u.Infof("out: %+v", newMsg)

func (m *JoinMerge) valIndexing(valOut, valSource []driver.Value, cols []*rel.Column) []driver.Value {
	_ = "STUB: not implemented"
	return nil
}

// Negative parent index means the parent query doesn't use this field, ie used
// as where, or join key, but not projected

//u.Infof("found: si=%v pi:%v idx:%d as=%v vals:%v len(out):%v", col.SourceIndex, col.ParentIndex, col.Index, col.As, valSource, len(valOut))
