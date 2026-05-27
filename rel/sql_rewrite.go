package rel

import (
	"github.com/araddon/qlbridge/expr"
)

// RewriteSelect We are removing Column Aliases "user_id as uid"
// as well as functions - used when we are going to defer projection, aggs
func RewriteSelect(m *SqlSelect) { _ = "STUB: not implemented"; return }

// RewriteSqlSource this Source to act as a stand-alone query to backend
// @parentStmt = the parent statement that this a partial source to
func RewriteSqlSource(m *SqlSource, parentStmt *SqlSelect) *SqlSelect {
	_ = "STUB: not implemented"
	return nil
}

// Rewrite this SqlSource for the given parent, ie
//   1)  find the column names we need to request from source including those used in join/where
//   2)  rewrite the where for this partial query
//   3)  any columns in join expression that are not equal between
//          sides should be aliased towards the left-hand join portion
//   4)  if we need different sort for our join algo?

// Was not left/right qualified, so use as is?  or is this an error?
//  what is official sql grammar on this?

// TODO:
//  - rewrite the Sort
//  - rewrite the group-by

// We need to check each participant in the Join for possible
// columns which need to be re-written

// We also need to create an expression used for evaluating
// the values of Join "Keys"

func rewriteIntoProjection(sel *SqlSelect, m Columns) { _ = "STUB: not implemented"; return }

// u.Infof("source=%-15s as=%-15s exprT:%T expr=%s  star:%v", c.As, c.SourceField, c.Expr, c.Expr, c.Star)

func addIntoProjection(sel *SqlSelect, newCols []string) { _ = "STUB: not implemented"; return }

// already in projection

func rewriteWhere(stmt *SqlSelect, from *SqlSource, node expr.Node, cols Columns) (expr.Node, Columns) {
	_ = "STUB: not implemented"
	//u.Debugf("rewrite where %s", node)
	return *new(expr.Node), *new(Columns)
}

//u.Debugf("rewriteWhere  from.Name:%v l:%v  r:%v", from.alias, left, right)

//u.Warnf("nice, found it! in = %v  cols:%d", in, len(cols))

//u.Warnf("what to do? source:%v    %v", from.alias, nt.String())

//u.Debugf("returning original: %s", nt)

//u.Infof("binaryNode  T:%v", nt.Operator.T.String())

//u.Warnf("n1=%#v  n2=%#v    %#v", n1, n2, nt)

//u.Debugf("n1=%#v  n2=%#v    %#v", n1, n2, nt)

// } else if n1 != nil {
// 	return n1
// } else if n2 != nil {
// 	return n2

//u.Warnf("n1=%#v  n2=%#v    %#v", n1, n2, nt)

//u.Warnf("un-implemented op: %#v", nt)

//u.Warnf("nil?? %T  %s  %#v", node, node, node)

func joinNodesForFrom(stmt *SqlSelect, from *SqlSource, node expr.Node, depth int) expr.Node {
	_ = "STUB: not implemented"
	return *new(expr.Node)
}

//u.Debugf("joinNodesForFrom  from.Name:%v l:%v  r:%v", from.alias, left, right)

//u.Debugf("%d nice, found it! identnode=%q fromnode:%q", depth, identNode.String(), nt.String())

// This is for other side of join, ignore
//u.Warnf("what to do? source:%v    %v", from.alias, nt.String())

//u.Warnf("skipping? %v", nt.String())

//u.Warnf("%v  try join from func node: %v", depth, nt.String())

// What???
//u.Infof("error, from:%q   arg:%q", from.String(), arg.String())

//u.Infof("adding func: %s", fn.String())

//u.Infof("%v binaryNode  %v", depth, nt.String())

//u.Debugf("%d neither nil:  n1=%v  n2=%v    %q", depth, n1, n2, nt.String())
//return &BinaryNode{Operator: nt.Operator, Args: [2]Node{n1, n2}}

//u.Debugf("%d n1 not nil: n1=%v  n2=%v    %q", depth, n1, n2, nt.String())

//u.Debugf("%d n2 not nil n1=%v  n2=%v    %q", depth, n1, n2, nt.String())

//u.Warnf("%d n1=%#v  n2=%#v    %#v", depth, n1, n2, nt)

//u.Debugf("%d neither nil:  n1=%v  n2=%v    %q", depth, n1, n2, nt.String())
//return &BinaryNode{Operator: nt.Operator, Args: [2]Node{n1, n2}}

//u.Debugf("%d n1 not nil: n1=%v  n2=%v    %q", depth, n1, n2, nt.String())
// 	return n1

//u.Infof("adding node: %s", n1.String())

//u.Debugf("%d  n2 not nil n1=%v  n2=%v    %q", depth, n1, n2, nt.String())

//u.Infof("adding node: %s", n1.String())

// 	return n2

//u.Warnf("n1=%#v  n2=%#v    %#v", n1, n2, nt)

// We need to find all columns used in the given Node (where/join expression)
//
//	to ensure we have those columns in projection for sub-queries
func columnsFromJoin(from *SqlSource, node expr.Node, cols Columns) Columns {
	_ = "STUB: not implemented"
	return *new(Columns)
}

//u.Debugf("columnsFromJoin()  T:%T  node=%q", node, node.String())

//u.Debugf("from.Name:%v AS %v   Joinnode l:%v  r:%v    %#v", from.Name, from.alias, left, right, nt)
//u.Warnf("check cols against join expr arg: %#v", nt)

//u.Debugf("left='%s'  colLeft='%s' right='%s'  %#v", left, colLeft, colRight,  col)
//u.Debugf("col:  From %s AS '%s'   '%s'.'%s'  JoinExpr: '%v'.'%v' col:%#v", from.Name, from.alias, colLeft, colRight, left, right, col)

//u.Infof("columnsFromJoin from.Name:%v l:%v  r:%v", from.alias, left, right)

//u.Warnf("not? from.Name:%v l:%v  r:%v   col: P:%p %#v", from.alias, left, right, col, col)

//u.Debugf("columnsFromJoin from.Name:%v l:%v  r:%v", from.alias, left, right)

// if -1, we don't need in parent index

//u.Warnf("added col %s idx:%d pidx:%v", right, newCol.Index, newCol.Index)

//u.Warnf("columnsFromJoin func node: %s", nt.String())

// Remove any aliases
func rewriteNode(from *SqlSource, node expr.Node) expr.Node {
	_ = "STUB: not implemented"
	return *new(expr.Node)
}

//u.Debugf("rewriteNode from.Name:%v l:%v  r:%v", from.alias, left, right)

//u.Warnf("nice, found it! in = %v", in)

//u.Warnf("skipping? %v", nt.String())

// What???
