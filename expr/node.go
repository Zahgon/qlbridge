// Expression structures, ie the  `a = b` type expression syntax
// including parser, node types, boolean logic check, functions.
package expr

import (
	"fmt"
	"reflect"
	"time"

	"github.com/araddon/qlbridge/lex"
	"github.com/araddon/qlbridge/value"
)

var (
	// ErrNotSupported indicates an error of a piece of expression syntax not
	// being supported.
	ErrNotSupported = fmt.Errorf("Not supported")
	// ErrNotImplemented an error of expression/statement syntax not being supported
	ErrNotImplemented = fmt.Errorf("Not implemented")
	// ErrUnknownCommand Unknown Command
	ErrUnknownCommand = fmt.Errorf("Unknown Command")
	// ErrInternalError Internal Error
	ErrInternalError = fmt.Errorf("Internal Error")

	// ErrNoIncluder is message saying a FilterQL included reference
	// to an include when no Includer was available to resolve
	ErrNoIncluder = fmt.Errorf("No Includer is available")
	// ErrIncludeNotFound Include Not Found
	ErrIncludeNotFound = fmt.Errorf("Include Not Found")

	// a static nil includer whose job is to return errors
	// for vm's that don't have an includer
	noIncluder = &IncludeContext{}

	// Ensure our dialect writer implements interface
	_ DialectWriter = (*defaultDialect)(nil)

	// Ensure some of our nodes implement Interfaces
	//_ NegateableNode = (*BinaryNode)(nil)
	_ NegateableNode = (*BooleanNode)(nil)
	_ NegateableNode = (*TriNode)(nil)
	_ NegateableNode = (*IncludeNode)(nil)

	// Ensure we implement interface
	_ Includer = (*IncludeContext)(nil)

	// Ensure some of our nodes implement NodeArgs
	_ NodeArgs = (*BooleanNode)(nil)
	_ NodeArgs = (*TriNode)(nil)
	_ NodeArgs = (*BinaryNode)(nil)
	_ NodeArgs = (*FuncNode)(nil)
	_ NodeArgs = (*UnaryNode)(nil)
	_ NodeArgs = (*ArrayNode)(nil)
)

type (

	// Node is a node in an expression tree, implemented
	// by different types (binary, urnary, func, identity, etc)
	//
	// qlbridge does not currently implement statements (if, for, switch, etc)
	// just expressions, and operators
	Node interface {
		// String representation of Node parseable back to itself
		String() string

		// WriteDialect Given a dialect writer write out, equivalent of String()
		// but allows different escape characters
		WriteDialect(w DialectWriter)

		// Validate Syntax validation of this expression node
		Validate() error

		// NodePb Convert this node to a Protobuf copy of it
		NodePb() *NodePb
		// FromPB Convert a protobuf presentation of node to Node.
		FromPB(*NodePb) Node

		// Expr Convert node into a simple expression syntax
		// which can be used for json respresentation
		Expr() *Expr
		// FromExpr
		FromExpr(*Expr) error

		// Equal compares deep equality of
		Equal(Node) bool

		// NodeType the String, Identity, etc
		NodeType() string
	}

	// NodeArgs is an interface for nodes which have child arguments
	NodeArgs interface {
		ChildrenArgs() []Node
	}

	// NegateableNode A negateable node requires a special type of String() function due to
	// an enclosing urnary NOT being inserted into middle of string syntax
	//
	//   <expression> [NOT] IN ("a","b")
	//   <expression> [NOT] BETWEEN <expression> AND <expression>
	//   <expression> [NOT] LIKE <expression>
	//   <expression> [NOT] CONTAINS <expression>
	//   <expression> [NOT] INTERSECTS ("a", "b")
	//
	NegateableNode interface {
		// Node the negateable nodes also implement the entire Node
		Node
		// Negated Say if this node is negateable (it may not be), If the node
		// is negateable, we may collapse an surrounding negation into here
		Negated() bool
		// ReverseNegation if Possible:  for instance:
		//   "A" NOT IN ("a","b")    =>  "A" IN ("a","b")
		ReverseNegation() bool
		// StringNegate
		StringNegate() string
		// WriteNegate write out this node into a writer
		WriteNegate(w DialectWriter)
		// Collapse Negateable nodes may be collapsed logically into new nodes
		// return this node collapsed down to simpliest form
		Collapse() Node
	}

	// EvalContext used to contain info for usage/lookup at runtime evaluation
	EvalContext interface {
		ContextReader
	}
	// EvalIncludeContext context, used to contain info for usage/lookup at runtime evaluation
	EvalIncludeContext interface {
		ContextReader
		Includer
	}

	// ContextReader is a key-value interface to read the context of message/row
	// using a  Get("key") interface.  Used by vm to evaluate messages
	ContextReader interface {
		Get(key string) (value.Value, bool)
		Row() map[string]value.Value
		Ts() time.Time
	}

	// ContextWriter For evaluation storage
	// vm writes results to this after evaluation
	ContextWriter interface {
		Put(col SchemaInfo, readCtx ContextReader, v value.Value) error
		Delete(row map[string]value.Value) error
	}

	// ContextReadWriter represents a Context which can be Read & Written.
	ContextReadWriter interface {
		ContextReader
		ContextWriter
	}

	// RowWriter for committing row ops (insert, update)
	RowWriter interface {
		// Commit the given rowInfo to persist
		Commit(rowInfo []SchemaInfo, row RowWriter) error
		// Put (persist) given Column info write single column.
		Put(col SchemaInfo, readCtx ContextReader, v value.Value) error
	}
)

type (

	// Expr represents single part of an Expression, it is a generic AST structure
	// that can be built in tree structure and JSON serialized to represent full AST
	// as json.
	Expr struct {
		// The `Op` (aka token), and Args expressions are non
		// nil if it is an expression
		Op   string  `json:"op,omitempty"`
		Args []*Expr `json:"args,omitempty"`

		// If op is 0, and args nil then exactly one of these should be set
		Identity string `json:"ident,omitempty"`
		Value    string `json:"val,omitempty"`
		// Really would like to use these instead of un-typed guesses above
		// if we desire serialization into string representation that is fine
		// Int      int64
		// Float    float64
		// Bool     bool
	}

	// Func Describes a function expression which wraps and allows native go functions
	// to be called in expression vm
	Func struct {
		Name       string        // name of func, lower-cased
		Aggregate  bool          // is this aggregate func?
		CustomFunc               // CustomFunc Is dynamic function that can be registered
		Eval       EvaluatorFunc // The memoized evaluation function
	}

	// FuncNode holds a Func, which desribes a go Function as
	// well as fulfilling the Pos, String() etc for a Node
	FuncNode struct {
		Name    string        // Name of func
		F       Func          // The actual function that this AST maps to
		Eval    EvaluatorFunc // the evaluator function
		Missing bool
		Args    []Node // Arguments are them-selves nodes
	}

	// IdentityNode will look up a value out of a env bag also identities of
	// sql objects (tables, columns, etc) we often need to rewrite these as in
	// sql it is `table.column`
	IdentityNode struct {
		Quote    byte
		Text     string
		original string
		escaped  string
		left     string
		right    string
	}
	// IdentityNodes is a list of identities
	IdentityNodes []*IdentityNode

	// StringNode holds a value literal, quotes not included
	StringNode struct {
		Quote       byte
		Text        string
		noQuote     bool
		needsEscape bool // Does Text contain Quote value?
	}

	// NullNode is a simple NULL type node
	NullNode struct{}

	// NumberNode holds a number: signed or unsigned integer or float.
	// The value is parsed and stored under all the types that can represent the value.
	// This simulates in a small amount of code the behavior of Go's ideal constants.
	NumberNode struct {
		IsInt   bool    // Number has an integer value.
		IsFloat bool    // Number has a floating-point value.
		Int64   int64   // The integer value.
		Float64 float64 // The floating-point value.
		Text    string  // The original textual representation from the input.
	}

	// ValueNode holds a value.Value type
	// value.Values can be strings, numbers, arrays, objects, etc
	ValueNode struct {
		Value value.Value
		rv    reflect.Value
	}

	// BinaryNode is x op y, two nodes (left, right) and an operator
	// operators can be a variety of:
	//    +, -, *, %, /, LIKE, CONTAINS, INTERSECTS
	// Also, parenthesis may wrap these
	BinaryNode struct {
		negated  bool
		Paren    bool
		Args     []Node
		Operator lex.Token
	}

	// BooleanNode is   n nodes and an operator
	// operators can be only AND/OR
	BooleanNode struct {
		negated  bool
		Args     []Node
		Operator lex.Token
	}

	// TriNode 3 part expression such as
	//    ARG1 Between ARG2 AND ARG3
	TriNode struct {
		negated  bool
		Args     []Node
		Operator lex.Token
	}

	// UnaryNode negates a single node argument
	//
	//    (  not <expression>  |   !<expression> )
	//
	//    !eq(5,6)
	//    !true
	//    !(true OR false)
	//    !toint(now())
	UnaryNode struct {
		Arg      Node
		Operator lex.Token
	}

	// IncludeNode references a named node
	//
	//   (  ! INCLUDE <identity>  |  INCLUDE <identity> | NOT INCLUDE <identity> )
	//
	IncludeNode struct {
		negated    bool
		inlineExpr Node // a non-pointer copy of the referred to include, itself resolved
		ExprNode   Node // The expression of the referred to include
		Identity   *IdentityNode
		Operator   lex.Token
	}

	// ArrayNode for holding multiple similar elements
	//    arg0 IN (arg1,arg2.....)
	//    5 in (1,2,3,4)
	ArrayNode struct {
		wraptype string //  (   or [
		Args     []Node
	}
)

// Includer defines an interface used for resolving INCLUDE clauses into a
// Include reference. Implementations should return an error if the name cannot
// be resolved.
type Includer interface {
	Include(name string) (Node, error)
}

// IncludeContext A ContextReader that implements Include interface.
type IncludeContext struct {
	ContextReader
}

// NewIncludeContext a new IncludeContext from contextreader.
func NewIncludeContext(cr ContextReader) *IncludeContext { _ = "STUB: not implemented"; return nil }

// Include interface not implemented.
func (*IncludeContext) Include(name string) (Node, error) {
	_ = "STUB: not implemented"
	return *

	// FindFirstIdentity Recursively descend down a node looking for first Identity Field
	//
	//	min(year)                 == year
	//	eq(min(item), max(month)) == item
	//	eq(min(user.last_name), max(month)) == user.last_name
	new(Node), nil
}

func FindFirstIdentity(node Node) string { _ = "STUB: not implemented"; return "" }

// FindAllIdentityField Recursively descend down a node looking for all Identity Fields
//
//	min(year)                 == {year}
//	eq(min(user.name), max(month)) == {user.name, month}
func FindAllIdentityField(node Node) []string { _ = "STUB: not implemented"; return nil }

// FindAllLeftIdentityFields Recursively descend down a node looking for all
// LEFT Identity Fields
//
//	min(year)                 == {year}
//	eq(min(user.name), max(month)) == {user, month}
func FindAllLeftIdentityFields(node Node) []string { _ = "STUB: not implemented"; return nil }

// FindAllIdentities gets all identity
func FindAllIdentities(node Node) IdentityNodes {
	_ = "STUB: not implemented"
	return *new(IdentityNodes)
}

func findIdentities(node Node, l IdentityNodes) IdentityNodes {
	_ = "STUB: not implemented"
	return *new(IdentityNodes)
}

// FilterSpecialIdentities given a list of identities, filter out
// special identities such as "null", "*", "match_all"
func FilterSpecialIdentities(l []string) []string { _ = "STUB: not implemented"; return nil }

// skip

// Strings get all identity strings
func (m IdentityNodes) Strings() []string { _ = "STUB: not implemented"; return nil }

// LeftStrings get all Left Identity fields.
func (m IdentityNodes) LeftStrings() []string { _ = "STUB: not implemented"; return nil }

// FindIdentityName Recursively walk a node looking for first Identity Field
// and combine with outermost expression to create an alias
//
//	min(year)                 => "min_year"
//	eq(min(year), max(month)) =>  "eq_year
//	EXISTS url                =>  "exists_url"
func FindIdentityName(depth int, node Node, prefix string) string {
	_ = "STUB: not implemented"
	return ""
}

// use the name of function

// ValueTypeFromNode Infer Value type from Node
func ValueTypeFromNode(n Node) value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}

// Identity types will draw type from context.

// NewFuncNode create new Function Expression Node.
func NewFuncNode(name string, f Func) *FuncNode { _ = "STUB: not implemented"; return nil }

func (m *FuncNode) append(arg Node) { _ = "STUB: not implemented"; return }

func (m *FuncNode) NodeType() string { _ = "STUB: not implemented"; return "" }
func (m *FuncNode) String() string   { _ = "STUB: not implemented"; return "" }

func (m *FuncNode) WriteDialect(w DialectWriter) { _ = "STUB: not implemented"; return }

func (m *FuncNode) Validate() error { _ = "STUB: not implemented"; return nil }

// Nice new style function

func (m *FuncNode) ChildrenArgs() []Node { _ = "STUB: not implemented"; return nil }

func (m *FuncNode) NodePb() *NodePb { _ = "STUB: not implemented"; return nil }

func (m *FuncNode) FromPB(n *NodePb) Node { _ = "STUB: not implemented"; return *new(Node) }

// Panic?

// Expr convert the FuncNode to Expr
func (m *FuncNode) Expr() *Expr { _ = "STUB: not implemented"; return nil }

func (m *FuncNode) FromExpr(e *Expr) error { _ = "STUB: not implemented"; return nil }

func (m *FuncNode) Equal(n Node) bool { _ = "STUB: not implemented"; return false }

// NewNumberStr is a little weird in that this Node accepts string @text
// and uses go to parse into Int, AND Float.
func NewNumberStr(text string) (*NumberNode, error) { _ = "STUB: not implemented"; return nil, nil }

func NewNumber(fv float64) (*NumberNode, error) { _ = "STUB: not implemented"; return nil, nil }

func (n *NumberNode) load() error {
	_ = "STUB: not implemented"
	// Do integer test first so we get 0x123 etc.
	return nil
}

// will fail for -0.

// If an integer extraction succeeded, promote the float.

// If a floating-point extraction succeeded, extract the int if needed.

func (m *NumberNode) NodeType() string             { _ = "STUB: not implemented"; return "" }
func (n *NumberNode) String() string               { _ = "STUB: not implemented"; return "" }
func (m *NumberNode) WriteDialect(w DialectWriter) { _ = "STUB: not implemented"; return }
func (m *NumberNode) Validate() error              { _ = "STUB: not implemented"; return nil }
func (m *NumberNode) NodePb() *NodePb              { _ = "STUB: not implemented"; return nil }

func (m *NumberNode) FromPB(n *NodePb) Node { _ = "STUB: not implemented"; return *new(Node) }

func (m *NumberNode) Expr() *Expr { _ = "STUB: not implemented"; return nil }

func (m *NumberNode) FromExpr(e *Expr) error { _ = "STUB: not implemented"; return nil }

func (m *NumberNode) Equal(n Node) bool { _ = "STUB: not implemented"; return false }

func NewStringNode(text string) *StringNode { _ = "STUB: not implemented"; return nil }

func NewStringNodeToken(t lex.Token) *StringNode { _ = "STUB: not implemented"; return nil }

func NewStringNoQuoteNode(text string) *StringNode { _ = "STUB: not implemented"; return nil }

func NewStringNeedsEscape(t lex.Token) *StringNode { _ = "STUB: not implemented"; return nil }

func (m *StringNode) NodeType() string { _ = "STUB: not implemented"; return "" }
func (m *StringNode) String() string   { _ = "STUB: not implemented"; return "" }

func (m *StringNode) WriteDialect(w DialectWriter) { _ = "STUB: not implemented"; return }

func (m *StringNode) Validate() error { _ = "STUB: not implemented"; return nil }
func (m *StringNode) NodePb() *NodePb { _ = "STUB: not implemented"; return nil }

func (m *StringNode) FromPB(n *NodePb) Node { _ = "STUB: not implemented"; return *new(Node) }

func (m *StringNode) Expr() *Expr { _ = "STUB: not implemented"; return nil }

func (m *StringNode) FromExpr(e *Expr) error { _ = "STUB: not implemented"; return nil }

func (m *StringNode) Equal(n Node) bool { _ = "STUB: not implemented"; return false }

func NewValueNode(val value.Value) *ValueNode { _ = "STUB: not implemented"; return nil }

func (m *ValueNode) NodeType() string { _ = "STUB: not implemented"; return "" }
func (m *ValueNode) IsArray() bool    { _ = "STUB: not implemented"; return false }

func (m *ValueNode) String() string { _ = "STUB: not implemented"; return "" }

func (m *ValueNode) WriteDialect(w DialectWriter) { _ = "STUB: not implemented"; return }

func (m *ValueNode) Validate() error { _ = "STUB: not implemented"; return nil }
func (m *ValueNode) NodePb() *NodePb { _ = "STUB: not implemented"; return nil }

func (m *ValueNode) FromPB(n *NodePb) Node { _ = "STUB: not implemented"; return *new(Node) }

func (m *ValueNode) Expr() *Expr { _ = "STUB: not implemented"; return nil }

func (m *ValueNode) FromExpr(e *Expr) error { _ = "STUB: not implemented"; return nil }

func (m *ValueNode) Equal(n Node) bool { _ = "STUB: not implemented"; return false }

func NewIdentityNode(tok *lex.Token) *IdentityNode { _ = "STUB: not implemented"; return nil }

func NewIdentityNodeVal(val string) *IdentityNode { _ = "STUB: not implemented"; return nil }

func (m *IdentityNode) load() {
	_ = "STUB: not implemented"

	// This is all deeply flawed, need to go fix it.  Upgrade path will
	// be sweep through and remove all usage of existing ones that used the flawed
	//
	//	`left.right value` escape syntax assuming the period is a split
	return
}

//   this came in with quote which has been stripped by lexer

//u.Debugf("branch1:  l:%q  r:%q  original:%q text:%q", m.left, m.right, m.original, m.Text)

//u.Debugf("branch2:  l:%q  r:%q  original:%q text:%q", m.left, m.right, m.original, m.Text)

//   this came in with quote which has been stripped by lexer
// m.original = fmt.Sprintf("%s%s%s", string(m.Quote), m.Text, string(m.Quote))
// m.left, m.right, _ = LeftRight(m.original)

//   this came in with quote which has been stripped by lexer

//u.Debugf("branch3:  l:%q  r:%q  original:%q text:%q", m.left, m.right, m.original, m.Text)

func (m *IdentityNode) NodeType() string { _ = "STUB: not implemented"; return "" }
func (m *IdentityNode) String() string   { _ = "STUB: not implemented"; return "" }

// What about escaping instead of replacing?

func (m *IdentityNode) WriteDialect(w DialectWriter) {
	_ = "STUB: not implemented"

	// `user`.`email`   type namespacing, may need to be escaped differently
	return
}

func (m *IdentityNode) OriginalText() string { _ = "STUB: not implemented"; return "" }

func (m *IdentityNode) Validate() error             { _ = "STUB: not implemented"; return nil }
func (m *IdentityNode) IdentityPb() *IdentityNodePb { _ = "STUB: not implemented"; return nil }

func (m *IdentityNode) NodePb() *NodePb { _ = "STUB: not implemented"; return nil }

func (m *IdentityNode) FromPB(n *NodePb) Node { _ = "STUB: not implemented"; return *new(Node) }

func (m *IdentityNode) Expr() *Expr { _ = "STUB: not implemented"; return nil }

//u.Warnf("This will NOT round-trip  l:%q  r:%q  original:%q text:%q", m.left, m.right, m.original, m.Text)

func (m *IdentityNode) FromExpr(e *Expr) error { _ = "STUB: not implemented"; return nil }

func (m *IdentityNode) IsBooleanIdentity() bool { _ = "STUB: not implemented"; return false }

func (m *IdentityNode) Bool() bool { _ = "STUB: not implemented"; return false }

func (m *IdentityNode) Equal(n Node) bool { _ = "STUB: not implemented"; return false }

// Hm, should we compare quotes or not?  Given they are dialect
// specific and don't affect logic i vote no?

// if nt.Quote != m.Quote {
// 	switch m.Quote {
// 	case '`':
// 		if nt.Quote == '\'' || nt.Quote == 0 {
// 			// ok
// 			return true
// 		}
// 	case 0:
// 		if nt.Quote == '\'' || nt.Quote == '`' {
// 			// ok
// 			return true
// 		}
// 	}

// HasLeftRight Return bool if is of form   `table.column` or `schema`.`table`
func (m *IdentityNode) HasLeftRight() bool { _ = "STUB: not implemented"; return false }

// Return left, right values if is of form   `table.column` or `schema`.`table` and
// also return true/false for if it even has left & right syntax
func (m *IdentityNode) LeftRight() (string, string, bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func NewNull(operator lex.Token) *NullNode { _ = "STUB: not implemented"; return nil }

func (m *NullNode) NodeType() string             { _ = "STUB: not implemented"; return "" }
func (m *NullNode) String() string               { _ = "STUB: not implemented"; return "" }
func (m *NullNode) WriteDialect(w DialectWriter) { _ = "STUB: not implemented"; return }

func (m *NullNode) Validate() error { _ = "STUB: not implemented"; return nil }
func (m *NullNode) NodePb() *NodePb { _ = "STUB: not implemented"; return nil }

func (m *NullNode) FromPB(n *NodePb) Node { _ = "STUB: not implemented"; return *new(Node) }

func (m *NullNode) Expr() *Expr { _ = "STUB: not implemented"; return nil }

func (m *NullNode) FromExpr(e *Expr) error { _ = "STUB: not implemented"; return nil }

func (m *NullNode) Equal(n Node) bool { _ = "STUB: not implemented"; return false }

/*
binary_op  = "||" | "&&" | rel_op | add_op | mul_op .
rel_op     = "==" | "!=" | "<" | "<=" | ">" | ">=" .
add_op     = "+" | "-" | "|" | "^" .
mul_op     = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .

unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" .
*/

// Create a Binary node
//
//	 @operator = * + - %/ / && || = ==
//	 @operator =  and, or, "is not"
//	@lhArg, rhArg the left, right side of binary
func NewBinaryNode(operator lex.Token, lhArg, rhArg Node) *BinaryNode {
	_ = "STUB: not implemented"
	//u.Debugf("NewBinaryNode: %v %v %v", lhArg, operator, rhArg)
	return nil
}

func (m *BinaryNode) NodeType() string { _ = "STUB: not implemented"; return "" }
func (m *BinaryNode) String() string   { _ = "STUB: not implemented"; return "" }

func (m *BinaryNode) WriteDialect(w DialectWriter) { _ = "STUB: not implemented"; return }

func (m *BinaryNode) writeToString(w DialectWriter, negate string) {
	_ = "STUB: not implemented"
	return
}

/*
Negation
I wanted to do negation on Binaries, but ended up not doing for now

ie, rewrite   NOT (X == "y")   =>  X != "y"

The general problem we ran into is that we lose some fidelity in collapsing
AST that is necessary for other evaluation run-times.

logically `NOT (X > y)` is NOT THE SAME AS  `(X <= y)   due to lack of existence of X

	func (m *BinaryNode) ReverseNegation() bool {
		switch m.Operator.T {
		case lex.TokenEqualEqual:
			m.Operator.T = lex.TokenNE
			m.Operator.V = m.Operator.T.String()
		case lex.TokenNE:
			m.Operator.T = lex.TokenEqualEqual
			m.Operator.V = m.Operator.T.String()
		case lex.TokenLT:
			m.Operator.T = lex.TokenGE
			m.Operator.V = m.Operator.T.String()
		case lex.TokenLE:
			m.Operator.T = lex.TokenGT
			m.Operator.V = m.Operator.T.String()
		case lex.TokenGT:
			m.Operator.T = lex.TokenLE
			m.Operator.V = m.Operator.T.String()
		case lex.TokenGE:
			m.Operator.T = lex.TokenLT
			m.Operator.V = m.Operator.T.String()
		default:
			//u.Warnf("What, what is this?   %s", m)
			m.negated = !m.negated
			return true
		}
		return true
	}

func (m *BinaryNode) Collapse() Node { return m }
func (m *BinaryNode) Negated() bool { return m.negated }

	func (m *BinaryNode) StringNegate() string {
		w := NewDefaultWriter()
		m.WriteNegate(w)
		return w.String()
	}

	func (m *BinaryNode) WriteNegate(w DialectWriter) {
		switch m.Operator.T {
		case lex.TokenIN, lex.TokenIntersects, lex.TokenLike, lex.TokenContains:
			m.writeToString(w, "NOT ")
		default:
			m.writeToString(w, "")
		}
	}
*/
func (m *BinaryNode) Validate() error { _ = "STUB: not implemented"; return nil }

func (m *BinaryNode) ChildrenArgs() []Node { _ = "STUB: not implemented"; return nil }

func (m *BinaryNode) NodePb() *NodePb { _ = "STUB: not implemented"; return nil }

func (m *BinaryNode) FromPB(n *NodePb) Node { _ = "STUB: not implemented"; return *new(Node) }

func (m *BinaryNode) Expr() *Expr { _ = "STUB: not implemented"; return nil }

func (m *BinaryNode) FromExpr(e *Expr) error { _ = "STUB: not implemented"; return nil }

func (m *BinaryNode) Equal(n Node) bool { _ = "STUB: not implemented"; return false }

// NewBooleanNode Create a boolean node
//
//	 @operator = AND, OR
//	@args = nodes
func NewBooleanNode(operator lex.Token, args ...Node) *BooleanNode {
	_ = "STUB: not implemented"
	//u.Debugf("NewBinaryNode: %v %v %v", lhArg, operator, rhArg)
	return nil
}

func (m *BooleanNode) NodeType() string      { _ = "STUB: not implemented"; return "" }
func (m *BooleanNode) ReverseNegation() bool { _ = "STUB: not implemented"; return false }

func (m *BooleanNode) String() string { _ = "STUB: not implemented"; return "" }

func (m *BooleanNode) StringNegate() string { _ = "STUB: not implemented"; return "" }

func (m *BooleanNode) WriteNegate(w DialectWriter) { _ = "STUB: not implemented"; return }

func (m *BooleanNode) WriteDialect(w DialectWriter) { _ = "STUB: not implemented"; return }

func (m *BooleanNode) writeToString(w DialectWriter, negate string) {
	_ = "STUB: not implemented"
	return
}

func (m *BooleanNode) Collapse() Node { _ = "STUB: not implemented"; return *new(Node) }

func (m *BooleanNode) Negated() bool   { _ = "STUB: not implemented"; return false }
func (m *BooleanNode) Validate() error { _ = "STUB: not implemented"; return nil }

func (m *BooleanNode) ChildrenArgs() []Node { _ = "STUB: not implemented"; return nil }

func (m *BooleanNode) NodePb() *NodePb { _ = "STUB: not implemented"; return nil }

func (m *BooleanNode) FromPB(n *NodePb) Node { _ = "STUB: not implemented"; return *new(Node) }

func (m *BooleanNode) Expr() *Expr { _ = "STUB: not implemented"; return nil }

func (m *BooleanNode) FromExpr(e *Expr) error { _ = "STUB: not implemented"; return nil }

func (m *BooleanNode) Equal(n Node) bool { _ = "STUB: not implemented"; return false }

// Create a Tri node
//
//	@arg1 [NOT] BETWEEN @arg2 AND @arg3
func NewTriNode(operator lex.Token, arg1, arg2, arg3 Node) *TriNode {
	_ = "STUB: not implemented"
	return nil
}

func (m *TriNode) NodeType() string      { _ = "STUB: not implemented"; return "" }
func (m *TriNode) ReverseNegation() bool { _ = "STUB: not implemented"; return false }

func (m *TriNode) String() string { _ = "STUB: not implemented"; return "" }

func (m *TriNode) StringNegate() string { _ = "STUB: not implemented"; return "" }

func (m *TriNode) WriteNegate(w DialectWriter) { _ = "STUB: not implemented"; return }

func (m *TriNode) WriteDialect(w DialectWriter) { _ = "STUB: not implemented"; return }

func (m *TriNode) writeToString(w DialectWriter, negate bool) { _ = "STUB: not implemented"; return }

func (m *TriNode) Collapse() Node  { _ = "STUB: not implemented"; return *new(Node) }
func (m *TriNode) Negated() bool   { _ = "STUB: not implemented"; return false }
func (m *TriNode) Validate() error { _ = "STUB: not implemented"; return nil }

func (m *TriNode) ChildrenArgs() []Node { _ = "STUB: not implemented"; return nil }

func (m *TriNode) NodePb() *NodePb { _ = "STUB: not implemented"; return nil }

//u.Debugf("TriNode NodePb: %T", arg)

func (m *TriNode) FromPB(n *NodePb) Node { _ = "STUB: not implemented"; return *new(Node) }

func (m *TriNode) Expr() *Expr { _ = "STUB: not implemented"; return nil }

func (m *TriNode) FromExpr(e *Expr) error { _ = "STUB: not implemented"; return nil }

func (m *TriNode) Equal(n Node) bool { _ = "STUB: not implemented"; return false }

// Unary nodes
//
//	NOT <expression>
//	! <expression>
//	EXISTS <identity>
//	<identity> IS NOT NULL
func NewUnary(operator lex.Token, arg Node) Node { _ = "STUB: not implemented"; return *new(Node) }

// In the event we are adding a binary here, we might have
// rewritten a little so lets make sure its interpreted/nested coorectly

func (m *UnaryNode) NodeType() string { _ = "STUB: not implemented"; return "" }
func (m *UnaryNode) String() string   { _ = "STUB: not implemented"; return "" }

func (m *UnaryNode) WriteDialect(w DialectWriter) { _ = "STUB: not implemented"; return }

func (m *UnaryNode) Validate() error { _ = "STUB: not implemented"; return nil }

func (m *UnaryNode) ChildrenArgs() []Node { _ = "STUB: not implemented"; return nil }

func (m *UnaryNode) Collapse() Node  { _ = "STUB: not implemented"; return *new(Node) }
func (m *UnaryNode) NodePb() *NodePb { _ = "STUB: not implemented"; return nil }

func (m *UnaryNode) FromPB(n *NodePb) Node { _ = "STUB: not implemented"; return *new(Node) }

func (m *UnaryNode) Expr() *Expr { _ = "STUB: not implemented"; return nil }

func (m *UnaryNode) FromExpr(e *Expr) error { _ = "STUB: not implemented"; return nil }

func (m *UnaryNode) Equal(n Node) bool { _ = "STUB: not implemented"; return false }

// Include nodes
//
//	NOT INCLUDE <identity>
//	! INCLUDE <identity>
//	INCLUDE <identity>
func NewInclude(operator lex.Token, id *IdentityNode) *IncludeNode {
	_ = "STUB: not implemented"
	return nil
}

func (m *IncludeNode) NodeType() string { _ = "STUB: not implemented"; return "" }
func (m *IncludeNode) String() string   { _ = "STUB: not implemented"; return "" }

func (m *IncludeNode) WriteDialect(w DialectWriter) { _ = "STUB: not implemented"; return }

func (m *IncludeNode) ReverseNegation() bool { _ = "STUB: not implemented"; return false }

func (m *IncludeNode) StringNegate() string { _ = "STUB: not implemented"; return "" }

func (m *IncludeNode) WriteNegate(w DialectWriter) {
	_ = "STUB: not implemented"
	// double negation
	return
}

func (m *IncludeNode) Validate() error { _ = "STUB: not implemented"; return nil }
func (m *IncludeNode) Negated() bool   { _ = "STUB: not implemented"; return false }
func (m *IncludeNode) Collapse() Node  { _ = "STUB: not implemented"; return *new(Node) }
func (m *IncludeNode) NodePb() *NodePb { _ = "STUB: not implemented"; return nil }

func (m *IncludeNode) FromPB(n *NodePb) Node { _ = "STUB: not implemented"; return *new(Node) }

func (m *IncludeNode) Expr() *Expr { _ = "STUB: not implemented"; return nil }

func (m *IncludeNode) FromExpr(e *Expr) error { _ = "STUB: not implemented"; return nil }

func (m *IncludeNode) Equal(n Node) bool { _ = "STUB: not implemented"; return false }

// NewArrayNode Create an array of Nodes which is a valid node type for boolean IN operator
func NewArrayNode() *ArrayNode { _ = "STUB: not implemented"; return nil }

func NewArrayNodeArgs(args []Node) *ArrayNode { _ = "STUB: not implemented"; return nil }

func (m *ArrayNode) NodeType() string { _ = "STUB: not implemented"; return "" }
func (m *ArrayNode) String() string   { _ = "STUB: not implemented"; return "" }

func (m *ArrayNode) WriteDialect(w DialectWriter) { _ = "STUB: not implemented"; return }

func (m *ArrayNode) Validate() error { _ = "STUB: not implemented"; return nil }

func (m *ArrayNode) ChildrenArgs() []Node { _ = "STUB: not implemented"; return nil }

func (m *ArrayNode) Append(n Node)   { _ = "STUB: not implemented"; return }
func (m *ArrayNode) NodePb() *NodePb { _ = "STUB: not implemented"; return nil }

func (m *ArrayNode) FromPB(n *NodePb) Node { _ = "STUB: not implemented"; return *new(Node) }

func (m *ArrayNode) Expr() *Expr { _ = "STUB: not implemented"; return nil }

func (m *ArrayNode) FromExpr(e *Expr) error { _ = "STUB: not implemented"; return nil }

func (m *ArrayNode) Equal(n Node) bool { _ = "STUB: not implemented"; return false }

// Node serialization helpers
func tokenFromInt(iv int32) lex.Token { _ = "STUB: not implemented"; return *new(lex.Token) }

// NodeFromPb Create a node from pb
func NodeFromPb(pb []byte) (Node, error) { _ = "STUB: not implemented"; return *new(Node), nil }

// NodeFromNodePb Create a node from pb
func NodeFromNodePb(n *NodePb) Node { _ = "STUB: not implemented"; return *new(Node) }

func NodesFromNodesPbPtr(pb []*NodePb) []Node { _ = "STUB: not implemented"; return nil }

func NodesFromNodesPb(pb []NodePb) []Node { _ = "STUB: not implemented"; return nil }

func NodesPbFromNodes(nodes []Node) []*NodePb { _ = "STUB: not implemented"; return nil }

func NodesEqual(n1, n2 Node) bool { _ = "STUB: not implemented"; return false }

func ExprsFromNodes(nodes []Node) []*Expr { _ = "STUB: not implemented"; return nil }

func NodesFromExprs(args []*Expr) ([]Node, error) { _ = "STUB: not implemented"; return nil, nil }

func NodeFromExpr(e *Expr) (Node, error) { _ = "STUB: not implemented"; return *new(Node), nil }

/*
	// Logic, Expressions, Operators etc
	TokenMultiply:   {Kw: "*", Description: "Multiply"},
	TokenMinus:      {Kw: "-", Description: "-"},
	TokenPlus:       {Kw: "+", Description: "+"},
	TokenPlusPlus:   {Kw: "++", Description: "++"},
	TokenPlusEquals: {Kw: "+=", Description: "+="},
	TokenDivide:     {Kw: "/", Description: "Divide /"},
	TokenModulus:    {Kw: "%", Description: "Modulus %"},
	TokenEqual:      {Kw: "=", Description: "Equal"},
	TokenEqualEqual: {Kw: "==", Description: "=="},
	TokenNE:         {Kw: "!=", Description: "NE"},
	TokenGE:         {Kw: ">=", Description: "GE"},
	TokenLE:         {Kw: "<=", Description: "LE"},
	TokenGT:         {Kw: ">", Description: "GT"},
	TokenLT:         {Kw: "<", Description: "LT"},
	TokenIf:         {Kw: "if", Description: "IF"},
	TokenAnd:        {Kw: "&&", Description: "&&"},
	TokenOr:         {Kw: "||", Description: "||"},
	TokenLogicOr:    {Kw: "or", Description: "Or"},
	TokenLogicAnd:   {Kw: "and", Description: "And"},
	TokenIN:         {Kw: "in", Description: "IN"},
	TokenLike:       {Kw: "like", Description: "LIKE"},
	TokenNegate:     {Kw: "not", Description: "NOT"},
	TokenBetween:    {Kw: "between", Description: "between"},
	TokenIs:         {Kw: "is", Description: "IS"},
	TokenNull:       {Kw: "null", Description: "NULL"},
	TokenContains:   {Kw: "contains", Description: "contains"},
	TokenIntersects: {Kw: "intersects", Description: "intersects"},
*/
//e.Op = strings.ToUpper(e.Op)

// udf

// bool

// This is a special Case, it is possible its urnary
// but in general we can collapse it

// very weird special case for FILTER * where the * is an ident not op

//u.Debugf("%T  %s", n, n)

// Negateable nodes possibly can be collapsed to simpler form
