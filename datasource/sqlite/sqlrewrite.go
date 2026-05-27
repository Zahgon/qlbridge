package sqlite

import (
	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/rel"
	"github.com/araddon/qlbridge/value"
)

type rewrite struct {
	sel           *rel.SqlSelect
	result        *rel.SqlSelect
	needsPolyFill bool // do we request that features be polyfilled?
}

func newRewriter(stmt *rel.SqlSelect) *rewrite { _ = "STUB: not implemented"; return nil }

// WalkSourceSelect An interface implemented by this connection allowing the planner
// to push down as much logic into mongo as possible
func (m *rewrite) rewrite() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Evaluate the Select columns make sure we can pass them down or polyfill

// this should be same right?

// eval() returns ( value, isOk, isIdentity )
func (m *rewrite) eval(arg expr.Node) (value.Value, bool, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false, false
}

// Aggregations from the <select_list>
//
//	SELECT <select_list> FROM ... WHERE
func (m *rewrite) walkSelectList() error { _ = "STUB: not implemented"; return nil }

//u.Debugf("i=%d of %d  %v %#v ", i, len(m.sel.Columns), col.Key(), col)

// case *expr.NumberNode:
// 	return nil, value.NewNumberValue(curNode.Float64), nil
// case *expr.BinaryNode:
// 	return m.walkBinary(curNode)
// case *expr.TriNode: // Between
// 	return m.walkTri(curNode)
// case *expr.UnaryNode:
// 	return m.walkUnary(curNode)

// All Func Nodes are Aggregates?

//u.Debugf("esm: %v:%v", col.As, esm)
//u.Debugf(curNode.String())
// case *expr.ArrayNode:
// 	return m.walkArrayNode(curNode)
// case *expr.IdentityNode:
// 	return nil, value.NewStringValue(curNode.Text), nil
// case *expr.StringNode:
// 	return nil, value.NewStringValue(curNode.Text), nil

//u.Debugf("likely a projection, not agg T:%T  %v", curNode, curNode)

//panic("Unrecognized node type")

// Group By Clause:  Mongo is a little weird where they move the
// group by expressions INTO the aggregation clause:
//
//	operation(field) FROM x GROUP BY x,y,z
//
//	db.article.aggregate([{"$group":{_id: "$author", count: {"$sum":1}}}]);
func (m *rewrite) walkGroupBy() error { _ = "STUB: not implemented"; return nil }

//fld := strings.Replace(expr.FindFirstIdentity(col.Expr), ".", "", -1)

// expressions when used ast part of <select_list>
func (m *rewrite) selectFunc(cur expr.Node) (expr.Node, error) {
	_ = "STUB: not implemented"
	return *new(expr.Node), nil
}

// case *expr.NumberNode:
// 	return nil, value.NewNumberValue(curNode.Float64), nil
// case *expr.BinaryNode:
// 	return m.walkBinary(curNode)
// case *expr.TriNode: // Between
// 	return m.walkTri(curNode)
// case *expr.UnaryNode:
// 	//return m.walkUnary(curNode)
// 	u.Warnf("not implemented: %#v", curNode)

// case *expr.ArrayNode:
// 	return m.walkArrayNode(curNode)
// case *expr.IdentityNode:
// 	return nil, value.NewStringValue(curNode.Text), nil
// case *expr.StringNode:
// 	return nil, value.NewStringValue(curNode.Text), nil

//panic("Unrecognized node type")

// Walk() an expression, and its logic to create an appropriately
// nested bson document for mongo queries if possible.
//
// - if can't express logic we need to allow qlbridge to poly-fill
func (m *rewrite) walkNode(cur expr.Node) (expr.Node, error) {
	_ = "STUB: not implemented"
	//u.Debugf("WalkNode: %#v", cur)
	return *new(expr.Node), nil
}

// Between

//return m.walkUnary(curNode)

// Tri Nodes expressions:
//
//	<expression> [NOT] BETWEEN <expression> AND <expression>
func (m *rewrite) walkFilterTri(node *expr.TriNode) (expr.Node, error) {
	_ = "STUB: not implemented"

	/*
		arg1val, aok, _ := m.eval(node.Args[0])
		if !aok {
			return nil, fmt.Errorf("Could not evaluate args: %v", node.String())
		}
		arg2val, bok := vm.Eval(nil, node.Args[1])
		arg3val, cok := vm.Eval(nil, node.Args[2])

		switch node.Operator.T {
		case lex.TokenBetween:
			u.Warnf("between? %T", arg2val.Value())
		default:
			u.Warnf("not implemented ")
		}
	*/return *new(expr.Node), nil
}

// Array Nodes expressions:
//
//	year IN (1990,1992)  =>
func (m *rewrite) walkArrayNode(node *expr.ArrayNode) (expr.Node, error) {
	_ = "STUB: not implemented"

	// Binary Node:   operations for >, >=, <, <=, =, !=, AND, OR, Like, IN
	//
	//	x = y             =>   db.users.find({field: {"$eq": value}})
	//	x != y            =>   db.inventory.find( { qty: { $ne: 20 } } )
	//
	//	x like "list%"    =>   db.users.find( { user_id: /^list/ } )
	//	x like "%list%"   =>   db.users.find( { user_id: /bc/ } )
	//	x IN [a,b,c]      =>   db.users.find( { user_id: {"$in":[a,b,c] } } )
	return *new(expr.Node), nil
}

func (m *rewrite) walkFilterBinary(node *expr.BinaryNode) (expr.Node, error) {
	_ = "STUB: not implemented"

	// If we have to recurse deeper for AND, OR operators
	return *new(expr.Node), nil
}

//u.Errorf("we found something wrong werener")
//return expr.NewIdentityNodeVal(fmt.Sprintf("%s IS NOT NULL", node.Args[0])), nil

//u.Warnf("rh %#v", node.Args[1])

// case lex.TokenLogicOr:
// 	lh, err := m.walkNode(node.Args[0])
// 	rh, err2 := m.walkNode(node.Args[1])
// 	if err != nil || err2 != nil {
// 		u.Errorf("could not get children nodes? %v %v %v", err, err2, node)
// 		return nil, fmt.Errorf("could not evaluate: %v", node.String())
// 	}
// 	node.Args[0] = lh
// 	node.Args[1] = rh
// 	return node, nil

//u.Debugf("walkBinary: %v  l:%v  r:%v  %T  %T", node, lhval, rhval, lhval, rhval)

// db.inventory.find( { qty: { $ne: 20 } } )

// db.inventory.find( { qty: { $lte: 20 } } )

// db.inventory.find( { qty: { $lt: 20 } } )

// db.inventory.find( { qty: { $gte: 20 } } )

// db.inventory.find( { qty: { $gt: 20 } } )

// { $text: { $search: <string>, $language: <string> } }
// { <field>: { $regex: /pattern/, $options: '<options>' } }

// switch vt := node.Args[1].(type) {
// case value.SliceValue:
// default:
// 	u.Warnf("not implemented type %#v", rhval)
// }

// Take an expression func, ensure we don't do runtime-checking (as the function)
// doesn't really exist, then map that function to a mongo operation
//
//	exists(fieldname)
//	regex(fieldname,value)
func (m *rewrite) walkFilterFunc(node *expr.FuncNode) (expr.Node, error) {
	_ = "STUB: not implemented"
	return *new(expr.Node), nil
}

// Take an expression func, ensure we don't do runtime-checking (as the function)
// doesn't really exist, then map that function to an Mongo Aggregation/MapReduce function
//
//	min, max, avg, sum, cardinality, terms
//
// Single Value Aggregates:
//
//	min, max, avg, sum, cardinality, count
//
// MultiValue aggregates:
//
//	terms, ??
func (m *rewrite) walkProjectionFunc(node *expr.FuncNode) (expr.Node, error) {
	_ = "STUB: not implemented"
	return *new(expr.Node), nil
}

func eval(cur expr.Node) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

//u.Errorf("unrecognized T:%T  %v", cur, cur)
