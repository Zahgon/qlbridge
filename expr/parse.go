package expr

import (
	"os"

	u "github.com/araddon/gou"
	"github.com/araddon/qlbridge/lex"
	"github.com/araddon/qlbridge/value"
)

var (
	_     = u.EMPTY
	Trace bool
	eoft  = lex.Token{T: lex.TokenEOF}
)

func init() {
	if t := os.Getenv("exprtrace"); t != "" {
		Trace = true
	}
}

func debugf(depth int, f string, args ...interface{}) { _ = "STUB: not implemented"; return }

// We have a default Dialect, which is the "Language" or rule-set of ql
var DefaultDialect *lex.Dialect = lex.LogicalExpressionDialect

// TokenPager wraps a Lexer, and implements the Logic to determine what is
// the end of this particular clause.  Lexer's are stateless, while
// tokenpager implements state ontop of pager and allows forward/back etc
type TokenPager interface {
	Peek() lex.Token
	Next() lex.Token
	Cur() lex.Token
	Backup()
	IsEnd() bool
	ClauseEnd() bool
	Lexer() *lex.Lexer
	ErrMsg(msg string) error
}

// SchemaInfo is interface for a Column type
type SchemaInfo interface {
	Key() string
}

// SchemaInfoString implements schemaInfo Key()
type SchemaInfoString string

func (m SchemaInfoString) Key() string {
	_ = "STUB: not implemented"

	// TokenPager is responsible for determining end of
	// current tree (column, etc)
	return ""
}

type LexTokenPager struct {
	done   bool
	tokens []lex.Token // list of all the tokens
	cursor int
	lex    *lex.Lexer
}

func NewLexTokenPager(lex *lex.Lexer) *LexTokenPager { _ = "STUB: not implemented"; return nil }

func (m *LexTokenPager) ErrMsg(msg string) error { _ = "STUB: not implemented"; return nil }

func (m *LexTokenPager) lexNext() { _ = "STUB: not implemented"; return }

// Next returns the current token and advances cursor to next one
func (m *LexTokenPager) Next() lex.Token { _ = "STUB: not implemented"; return *new(lex.Token) }

//u.Warnf("Next() CRAP? increment cursor: %v of %v %v", m.cursor, len(m.tokens))

// Returns the current token, does not advance
func (m *LexTokenPager) Cur() lex.Token { _ = "STUB: not implemented"; return *new(lex.Token) }

//u.Warnf("Next() CRAP? increment cursor: %v of %v %v", m.cursor, len(m.tokens), m.cursor < len(m.tokens))

// IsEnd determines if pager is at end of statement
func (m *LexTokenPager) IsEnd() bool {
	_ = "STUB: not implemented"

	// ClauseEnd are we at end of clause
	return false
}

func (m *LexTokenPager) ClauseEnd() bool {
	_ = "STUB: not implemented"

	// Lexer get the underlying lexer
	return false
}

func (m *LexTokenPager) Lexer() *lex.Lexer {
	_ = "STUB: not implemented"

	// backup backs the input stream up one token.
	return nil
}

func (m *LexTokenPager) Backup() { _ = "STUB: not implemented"; return }

// Peek returns but does not consume the next token.
func (m *LexTokenPager) Peek() lex.Token { _ = "STUB: not implemented"; return *new(lex.Token) }

// Tree is the representation of a single parsed expression
type tree struct {
	funcCheck  bool // should we resolve function existence at parse time?
	boolean    bool // Stateful flag for in mid of boolean expressions
	TokenPager      // pager for grabbing next tokens, backup(), recognizing end
	fr         FuncResolver
}

func newTree(pager TokenPager) *tree { _ = "STUB: not implemented"; return nil }

func newTreeFuncs(pager TokenPager, fr FuncResolver) *tree { _ = "STUB: not implemented"; return nil }

// ParseExpression parse a single Expression, returning an Expression Node
//
//	ParseExpression("5 * toint(item_name)")
func ParseExpression(expressionText string) (Node, error) {
	_ = "STUB: not implemented"
	return *new(Node), nil
}

// Parser panics on unexpected syntax, convert this into an err

// MustParse parse a single Expression, returning an Expression Node
// and panics if it cannot be parsed
//
//	MustParse("5 * toint(item_name)")
func MustParse(expressionText string) Node { _ = "STUB: not implemented"; return *new(Node) }

// Parse a single Expression, returning an Expression Node
//
// @fr = function registry with any additional functions
//
//	ParseExprWithFuncs("5 * toint(item_name)", funcRegistry)
func ParseExprWithFuncs(p TokenPager, fr FuncResolver) (Node, error) {
	_ = "STUB: not implemented"
	return *

	// Parser panics on unexpected syntax, convert this into an err
	new(Node), nil
}

// Parse a single Expression, returning an Expression Node
//
// @pager = Token Pager
func ParsePager(pager TokenPager) (Node, error) {
	_ = "STUB: not implemented"

	// Parser panics on unexpected syntax, convert this into an err
	return *new(Node), nil
}

// errorf formats the error and terminates processing.
func (t *tree) errorf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// error terminates processing.
func (t *tree) error(err error) { _ = "STUB: not implemented"; return }

// expect verifies the current token and guarantees it has the required type
func (t *tree) expect(expected lex.TokenType, context string) lex.Token {
	_ = "STUB: not implemented"
	return *new(lex.Token)
}

// expectOneOf consumes the next token and guarantees it has one of the required types.
func (t *tree) expectOneOf(expected1, expected2 lex.TokenType, context string) lex.Token {
	_ = "STUB: not implemented"
	return *new(lex.Token)
}

// unexpected complains about the token and terminates processing.
func (t *tree) unexpected(token lex.Token, msg string) { _ = "STUB: not implemented"; return }

// recover is the handler that turns panics into returns from the top level of Parse.
func (t *tree) recover(errp *error) { _ = "STUB: not implemented"; return }

// parse take the tokens and recursively build into Node
func (t *tree) parse() (_ Node, err error) { _ = "STUB: not implemented"; return *new(Node), nil }

/*

General overview of Recursive Descent Parsing

https://www.engr.mun.ca/~theo/Misc/exp_parsing.htm

Operator Predence planner during parse phase:
  when we parse and build our node-sub-node structures we need to plan
  the precedence rules, we use a recursion tree to build this

http://dev.mysql.com/doc/refman/5.0/en/operator-precedence.html
https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Operator_Precedence
http://www.postgresql.org/docs/9.4/static/sql-syntax-lexical.html#SQL-PRECEDENCE

TODO:
 - if/else, case, for
 - call stack & vars
--------------------------------------
O -> A {( "||" | OR  ) A}
A -> C {( "&&" | AND ) C}
C -> P {( "==" | "!=" | ">" | ">=" | "<" | "<=" | "LIKE" | "IN" | "CONTAINS" | "INTERSECTS") P}
P -> M {( "+" | "-" ) M}
M -> F {( "*" | "/" ) F}
F -> v | "(" O ")" | "!" v | "-" O | "NOT" C | "EXISTS" v | "IS" O | "AND (" O ")" | "OR (" O ")"
v -> value | Func | "INCLUDE" <identity>
Func -> <identity> "(" value {"," value} ")"
value -> number | "string" | O | <identity>



Recursion:  We recurse so the LAST to evaluate is the highest (parent, then or)
   ie the deepest we get in recursion tree is the first to be evaluated

0	Value's
1	Unary + - arithmetic operators
2	* / arithmetic operators
3	Binary + - arithmetic operators, || character operators
4	All comparison operators
5	NOT logical operator
6	AND logical operator
7	OR logical operator
8	Paren's


*/

// expr:
func (t *tree) O(depth int) Node { _ = "STUB: not implemented"; return *new(Node) }

// consume --
// consume comment after --

// these are indicators of End of Current Clause, so we can return

func (t *tree) A(depth int) Node { _ = "STUB: not implemented"; return *new(Node) }

// This is a Boolean Expression Not Binary

func (t *tree) C(depth int) Node { _ = "STUB: not implemented"; return *new(Node) }

func (t *tree) cInner(n Node, depth int) Node { _ = "STUB: not implemented"; return *new(Node) }

// weird syntax:    BETWEEN x AND y     AND is ignored essentially

// Right side is an array of values

// This is a special type of Binary? its 2nd argument is a array node

// consume Function Name

// Consume "INTERSECTS"

// x INTERSECTS field   where field MUST be an array

// The 2nd argument is an array node

// consume Function Name

func (t *tree) P(depth int) Node { _ = "STUB: not implemented"; return *new(Node) }

func (t *tree) M(depth int) Node { _ = "STUB: not implemented"; return *new(Node) }

// F -> v | "(" O ")" | "!" O | "-" O | "NOT" C | "EXISTS" v | "IS" O | "AND (" O ")" | "OR (" O ")"
func (t *tree) F(depth int) Node { _ = "STUB: not implemented"; return *new(Node) }

// Urnary operations

// consume NOT, !, Minus

// TODO:  this is a bug.  An old version of generator was saving these
//  NOT news INTERSECTS ("a")    which is invalid it should be
//  news NOT INTERSECTS ("a")  OR NOT (news INTERSECTS ("a"))
//
// NOT <expr> LIKE <expr>
// NOT <expr> INTERSECTS <expr>
// NOT <expr> BETWEEN <expr> AND <expr>
// NOT <expr> CONTAINS <expr>
// NOT <expr> IN <expr>
//
// NOT identity > 7

// Urnary operations:  require right side value node
// Consume "EXISTS"

// consume AND/OR

// Consume Left Paren

// Whoops, binary not boolean, there are some ambiguous ones:
// binary:   x = y OR ( stuff > 5)
// boolean:  AND (x = y, OR ( stuff > 5, x = 9))

//debugf(depth, "found boolean expression %v", n.Collapse())

func (t *tree) v(depth int) Node { _ = "STUB: not implemented"; return *new(Node) }

// consume Include

//u.Debugf("inc: %v  nxt %v", inc, nxt)

// Consume identity

// [   ie     [1,2,3] json array or static array values
// Consume the [

// consume Function Name

// Consume  (

func (t *tree) Func(depth int, funcTok lex.Token) (fn *FuncNode) {
	_ = "STUB: not implemented"
	return nil
}

// if we aren't testing for validity, make a "fake" func
// we may not be using vm, just ast
//u.Warnf("non func? %v", funcTok.V)

// Are we sure we consume?

// will panic

// Ugh, we need a way of identifying which functions get this special
// parser?

// We are not in a comma style function
//  CAST(<expression> AS <identity>)

// This really isn't correct, we probably need an OperatorNode?

// continue

// this func arg is an expression
//     toint(str_item * 5)

// get Function from Global function registry.
func (t *tree) getFunction(name string) (fn Func, ok bool) {
	_ = "STUB: not implemented"
	return *new(Func), false
}

// ArrayNode parses multi-argument array nodes aka: IN (a,b,c).
func (t *tree) ArrayNode(depth int) Node { _ = "STUB: not implemented"; return *new(Node) }

// Consume Left Paren

// Consume the Paren

// ValueArray
//
//	IN ("a","b","c")
//	["a","b","c"]
func ValueArray(depth int, pg TokenPager) (value.Value, error) {
	_ = "STUB: not implemented"
	return *new(value.Value), nil
}

// consume token

// continue

// fine, consume the comma

func nodeArray(t *tree, depth int) ([]Node, error, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

// Consume

// We are going to loop until we find the first Non-Comment Token

// Consume new line

// indicates start of new expression
// consume comma

// skip, currently ignore these

// first non-comment token

func (t *tree) discardNewLinesAndComments() {
	_ = "STUB: not implemented"

	// We are going to loop until we find the first Non-Comment Token
	return
}

// Consume new line

// skip, currently ignore these

// first non-comment token
