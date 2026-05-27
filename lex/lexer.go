// Package Lex is a Lexer for QLBridge which is more of a lex-toolkit and
// implements 4 Dialects {SQL, FilterQL, Json, Expressions}.
package lex

import (
	"os"

	u "github.com/araddon/gou"
)

var (
	// Trace is a global var to turn on tracing.  can be turned out with env
	// variable "lextrace=true"
	//
	//     export lextrace=true
	Trace bool
)

func init() {
	if t := os.Getenv("lextrace"); t != "" {
		Trace = true
	}
}

func debugf(f string, args ...interface{}) { _ = "STUB: not implemented"; return }

var (
	// SUPPORT_DURATION FEATURE FLAGS
	SUPPORT_DURATION = true
	// Identity Quoting
	//  http://stackoverflow.com/questions/1992314/what-is-the-difference-between-single-and-double-quotes-in-sql
	// you might want to set this to not include single ticks
	//  http://dev.mysql.com/doc/refman/5.7/en/string-literals.html
	//IdentityQuoting = []byte{'[', '`', '"'} // mysql ansi-ish, no single quote identities, and allowing double-quote
	IdentityQuotingWSingleQuote = []byte{'[', '`', '\''} // more ansi-ish, allow single quotes around identities
	IdentityQuoting             = []byte{'[', '`'}       // no single quote around identities bc effing mysql uses single quote for string literals
)

const (
	eof       = -1
	decDigits = "0123456789"
	hexDigits = "0123456789ABCDEF"
)

// StateFn represents the state of the lexer as a function that returns the
// next state.
type StateFn func(*Lexer) StateFn

// NamedStateFn is a StateFn which has a name for tracing debugging.
type NamedStateFn struct {
	Name    string
	StateFn StateFn
}

// NewLexer Creates a new lexer for the input string
func NewLexer(input string, dialect *Dialect) *Lexer {
	_ = "STUB: not implemented"
	// Three tokens of buffering is sufficient for all state functions.
	return nil
}

// Lexer holds the state of the lexical scanning.
//
// Holds a *Dialect* which gives much of the rules specific to this language.
//
// many-generations removed from that Based on the lexer from the "text/template" package.
// See http://www.youtube.com/watch?v=HxaD_trXwRE
type Lexer struct {
	input         string     // the string being scanned
	state         StateFn    // the next lexing function to enter
	identityRunes []byte     // List of legal identity escape bytes
	pos           int        // current position in the input
	start         int        // start position of this token
	width         int        // width of last rune read from input
	line          int        // Line we are currently on
	linepos       int        // Position of start of current line
	lastToken     Token      // last token we emitted
	tokens        chan Token // channel of scanned tokens we output on
	doubleDelim   bool       // flag for tags starting with double braces
	dialect       *Dialect   // Dialect is the syntax-rules for all statement-types of this language
	statement     *Clause    // Statement type we are lexing
	curClause     *Clause    // Current clause we are lexing, we descend, ascend, iter()
	descent       *Clause    // Clause we have descended to
	peekedWordPos int
	peekedWord    string
	lastQuoteMark byte

	// Due to nested Expressions and evaluation this allows us to descend/ascend
	// during lex, using push/pop to add and remove states needing evaluation
	stack []NamedStateFn
}

func (l *Lexer) init() {
	l.ReverseTrim()
}

// ErrMsg an error message helper which provides context of where in input string
// the error is occuring, line, column, current token info.
func (l *Lexer) ErrMsg(t Token, msg string) error { _ = "STUB: not implemented"; return nil }

// NextToken returns the next token from the input.
func (l *Lexer) NextToken() Token {
	_ = "STUB: not implemented"

	// u.Debugf("token: start=%v  pos=%v  peek5=%s", l.start, l.pos, l.PeekX(5))
	return *new(Token)
}

// Push a named StateFn onto stack.
func (l *Lexer) Push(name string, state StateFn) { _ = "STUB: not implemented"; return }

func (l *Lexer) pop() StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

// Next returns the next rune in the input
func (l *Lexer) Next() (r rune) { _ = "STUB: not implemented"; return 0 }

func (l *Lexer) skipX(ct int) { _ = "STUB: not implemented"; return }

// RawInput return the orgiginal string we are lexing.
func (l *Lexer) RawInput() string {
	_ = "STUB: not implemented"

	// Remainder SQL and other string expressions may contain more than one
	// statement such as:
	//
	//	use schema_x;  show tables;
	//
	//	set @my_var = "value"; select a,b from `users` where name = @my_var;
	return ""
}

func (l *Lexer) Remainder() (string, bool) { _ = "STUB: not implemented"; return "", false }

// Peek returns but does not consume the next rune in the input.
func (l *Lexer) Peek() rune { _ = "STUB: not implemented"; return 0 }

// PeekX grab the next x characters without consuming
func (l *Lexer) PeekX(x int) string { _ = "STUB: not implemented"; return "" }

// get single character PAST
func (l *Lexer) peekRunePast(skip int) rune { _ = "STUB: not implemented"; return 0 }

// PeekWord grab the next word (till whitespace, without consuming)
func (l *Lexer) PeekWord() string { _ = "STUB: not implemented"; return "" }

// TODO:  optimize this, this is by far the most expensive operation
// in the lexer
//    - move to some type of early bail?  ie, use Accept() wherever possible?

//skipWs += (ri - 1)

//u.Debugf("r: %v  identifier?%v", string(r), IsIdentifierRune(r))

//i += (ri - 1)

//u.Infof("hm:   '%v'", l.input[l.pos+skipWs:l.pos+i])

// regardless of being short, lets treat like word

//u.Infof("hm:   '%v'", l.input[l.pos+skipWs:l.pos+i])

/*
// get single character
func (l *Lexer) peekXrune(x int) rune {
	if l.pos+x > len(l.input) {
		return rune(0)
	}
	return rune(l.input[l.pos+x])
}

// PeekWord2 grab the next word (till whitespace, without consuming)
func (l *Lexer) PeekWord2() string {

	skipWs := 0
	for ; skipWs < len(l.input)-l.pos; skipWs++ {
		r, _ := utf8.DecodeRuneInString(l.input[l.pos+skipWs:])
		if !unicode.IsSpace(r) {
			break
		}
	}

	word := ""
	for i := skipWs; i < len(l.input)-l.pos; i++ {
		r, _ := utf8.DecodeRuneInString(l.input[l.pos+i:])
		if unicode.IsSpace(r) || !IsIdentifierRune(r) {
			u.Infof("hm:   '%v' word='%s' %v", l.input[l.pos:l.pos+i], word, l.input[l.pos:l.pos+i] == word)
			return word
		} else {
			word = word + string(r)
		}
	}
	return word
}

// peek word, but using laxIdentifier characters
func (l *Lexer) peekLaxWord() string {
	word := ""
	for i := 0; i < len(l.input)-l.pos; i++ {
		r, _ := utf8.DecodeRuneInString(l.input[l.pos+i:])
		if !isLaxIdentifierRune(r) {
			return word
		} else {
			word = word + string(r)
		}
	}
	return word
}
// Discard skips over the pending input before this point.
func (l *Lexer) Discard() {
	l.start = l.pos
}
*/

// backup steps back one rune. Can only be called once per call of next.
func (l *Lexer) backup() {
	_ = "STUB: not implemented"

	// IsEnd have we consumed all input?
	return
}

func (l *Lexer) IsEnd() bool {
	_ = "STUB: not implemented"
	// u.Infof("isEnd? %v:%v", l.pos, len(l.input))
	return false
}

// if l.Peek() == ';' {
// 	return true
// }

// IsComment Is this a comment?
func (l *Lexer) IsComment() bool { _ = "STUB: not implemented"; return false }

// continue on, might be, check 2nd character

// Emit passes an token back to the client.
func (l *Lexer) Emit(t TokenType) { _ = "STUB: not implemented"; return }

// We are going to use 1 based indexing (not 0 based) for lines
// because humans don't think that way

// ignore skips over the pending input before this point.
func (l *Lexer) ignore() {
	_ = "STUB: not implemented"

	// ignore skips over the item
	return
}

func (l *Lexer) ignoreWord(word string) { _ = "STUB: not implemented"; return }

// accept consumes the next rune if it's from the valid set.
func (l *Lexer) accept(valid string) bool { _ = "STUB: not implemented"; return false }

// acceptRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRun(valid string) bool { _ = "STUB: not implemented"; return false }

// Returns current string not yet emitted
func (l *Lexer) current() string { _ = "STUB: not implemented"; return "" }

// ConsumeWord lets move position to consume given word
func (l *Lexer) ConsumeWord(word string) {
	_ = "STUB: not implemented"
	// pretty sure the len(word) is valid right?
	return
}

/*
// lineNumber reports which line we're on. Doing it this way
// means we don't have to worry about peek double counting.
func (l *Lexer) lineNumber() int {
	//return 1 + strings.Count(l.input[:l.pos], "\n")
	return l.line
}
// Returns remainder of input not yet lexed
func (l *Lexer) remainder() string {
	return l.input[l.start : len(l.input)-1]
}
*/

// error returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextToken.
func (l *Lexer) errorf(format string, args ...interface{}) StateFn {
	_ = "STUB: not implemented"
	return *new(StateFn)
}

// columnNumber reports which column in the current line we're on.
func (l *Lexer) columnNumber() int {
	_ = "STUB: not implemented"
	// n := strings.LastIndex(l.input[:l.pos], "\n")
	//
	//	if n == -1 {
	//		n = 0
	//	}
	//
	// return l.pos - n
	return 0
}

// SkipWhiteSpaces Skips white space characters in the input.
func (l *Lexer) SkipWhiteSpaces() { _ = "STUB: not implemented"; return }

// New line, lets keep track of line position

// SkipWhiteSpacesNewLine Skips white space characters in the input, returns bool
// for if it contained new line
func (l *Lexer) SkipWhiteSpacesNewLine() bool { _ = "STUB: not implemented"; return false }

// New line, lets keep track of line position

// Skips white space characters at end by trimming so we can recognize the end
//
//	more easily
func (l *Lexer) ReverseTrim() { _ = "STUB: not implemented"; return }

//u.Warnf("trim: '%v'", l.input[:i+1])

// Scans input and matches against the string.
// Returns true if the expected string was matched.
// expects matchTo to be a lower case string
func (l *Lexer) match(matchTo string, skip int) bool {
	_ = "STUB: not implemented"

	// u.Debugf("match(%q)  peek:%q ", matchTo, l.PeekWord())
	return false
}

//u.Debugf("match rune? %v", string(matchRune))

//u.Debugf("rune=%s n=%s   %v  %v", string(matchRune), string(nr), matchRune != nr, unicode.ToLower(nr) != matchRune)

//u.Debugf("setting done = false?, ie did not match")

// If we finished looking for the match word, and the next item is not
// whitespace, it means we failed

//u.Debugf("Found match():  %v", matchTo)

// Scans input and tries to match the expected string.
// Returns true if the expected string was matched.
// Does not advance the input if the string was not matched.
//
// NOTE:  this assumes the @val you are trying to match against is LOWER CASE
func (l *Lexer) tryMatch(matchTo string) bool {
	_ = "STUB: not implemented"

	// u.Debugf("tryMatch:  start='%v'", l.PeekWord())
	return false
}

//u.Warnf("not found:  %v:%v", string(nextRune), matchTo)

//u.Debugf("tryMatch:  good='%v'", matchTo)

// Emits an error token and terminates the scan
// by passing back a nil ponter that will be the next state
// terminating lexer.next function
func (l *Lexer) errorToken(format string, args ...interface{}) StateFn {
	_ = "STUB: not implemented"
	//fmt.Sprintf(format, args...)
	return *new(StateFn)
}

// non-consuming isExpression, expressions are defined by
//
//	starting with
//	  - negation (!)
//	  - non quoted alpha character
//	  - (   left-paren
func (l *Lexer) isExpr() bool {
	_ = "STUB: not implemented"
	// Expressions are strings not values, so quoting them means no
	return false
}

// first character of expression cannot be digit

//u.Debugf("found negation! : %v", string(r))
// Negation is possible?

// ??? paran's wrapping sub-expressions?

// Expressions are terminated by either a parenthesis
// never by spaces

// else isAlNumOrPeriod so keep looking

// non-consuming check to see if we are about to find next keyword
func (l *Lexer) isNextKeyword(peekWord string) bool { _ = "STUB: not implemented"; return false }

//u.Debugf("isNextKeyword?  '%s'   len:%v", kwMaybe, len(l.statement.Clauses))

//u.Infof("clause: %s", clause)

//u.Warnf("returning, not keyword")

//clause = l.statement.Clauses[i]
//u.Infof("clause: %+v", clause)
//u.Debugf("clause next keyword?    peek=%s cname=%q keyword=%v multi?%v children?%v", kwMaybe, clause.Name, clause.keyword, clause.multiWord, len(clause.Clauses))

//u.Infof("return true:  %v", strings.ToLower(l.PeekX(len(clause.fullWord))))

// TODO:  allow clauses to reserve keywords, or sub-clause

//u.Warnf("doing true: %v", kwMaybe)

// non-consuming isIdentity
// Identities are non-numeric string values that are not quoted
func (l *Lexer) isIdentity() bool {
	_ = "STUB: not implemented"
	// Identity are strings not values
	return false
}

// This character [ is a little special
// as it is going to look to see if the 2nd character is
// valid identity character so ie alpha/numeric

// are these always identities?  or do we need
// to also check first identifier?
// peek2 := l.PeekX(2)
// if len(peek2) == 2 {
// 	return isIdentifierFirstRune(rune(peek2[1]))
// }

// Uses the identity escaping/quote characters
func (l *Lexer) isIdentityQuoteMark(r rune) bool { _ = "STUB: not implemented"; return false }

/*
// LexMatchSkip matches expected tokentype emitting the token on success
// and returning passed state function.
func (l *Lexer) LexMatchSkip(tok TokenType, skip int, fn StateFn) StateFn {
	//u.Debugf("lexMatch   t=%s peek=%s", tok, l.PeekWord())
	if l.match(tok.String(), skip) {
		//u.Debugf("found match: %s   %v", tok, fn)
		l.Emit(tok)
		return fn
	}
	u.Error("unexpected token", tok)
	return l.errorToken("Unexpected token:" + l.current())
}
*/

// lexer to match expected value returns with args of
//
//	@matchState state function if match
//	if no match, return nil
func (l *Lexer) lexIfMatch(tok TokenType, matchState StateFn) StateFn {
	_ = "STUB: not implemented"
	return *new(StateFn)
}

// current clause state function, used for repeated clauses
func (l *Lexer) clauseState() StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

var emptyLexFn = func(*Lexer) StateFn { u.Debugf("empty statefun"); return nil }

// LexMatchClosure matches expected tokentype emitting the token on success
// and returning passed state function.
func LexMatchClosure(tok TokenType, nextFn StateFn) StateFn {
	_ = "STUB: not implemented"
	return *new(StateFn)
}

//u.Debugf("%p lexMatch   t=%s peek=%s", l, tok, l.PeekWord())

//u.Debugf("found match: %s   %v", tok, nextFn)

// State functions ------------------------------------------------------------

// Find first keyword in the current queryText, then find appropriate statement in dialect.
// ie [SELECT, ALTER, CREATE, INSERT] in sql
func LexDialectForStatement(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

// ensure we have consumed all initial pre-statement comments

//u.Warnf("LexDialectForStatement peek=%s  keyword=%v ", peekWord, stmt.Token.String())

// We aren't actually going to consume anything here, just find
// the correct statement

//u.Infof("statement: %s  curClause %s", l.statement, l.curClause)

// LexStatement is the main entrypoint to lex Grammars primarily associated with QL type
// languages, which is keywords separate clauses, and have order [select .. FROM name WHERE ..]
// the keywords which are reserved serve as identifiers to stop lexing and move to next clause
// lexer
func LexStatement(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

// ensure we have consumed all comments

//u.Debugf("%p curClause %s peek: %s", l, clause, peekWord)

// Before we move onto next clause, lets check and see if we need to descend
//  into sub-clauses of current statement

//u.Infof("%p has sub-clauses  kw=%-10s peek=%-10s", l, clause.keyword, peekWord)

//u.Infof("%p has sub-clauses  kw=%-10s peek=%-10s", l, sc.keyword, peekWord)

//u.Infof("matches Sub-Clause: %-10s", sc.keyword)

// First non-optional we don't match we bail

// we only ever consume each clause once?
//u.Debugf("%p:%p stmt.clause parser?  peek=%-10q  keyword=%-10q multi?%v name=%-10q", clause.parent, clause, peekWord, clause.keyword, clause.multiWord, clause.Name)

// Set the default entry point for this keyword

//u.Debugf("dialect clause:  '%v' LexerNil?%v \n\t %s ", clause.keyword, clause.Lexer == nil, l.input)
//u.Infof("matched stmt.clause token=%-10q match %-10q  clausekw=%-10q multi?%v name=%-10q", clause.Token, peekWord, clause.keyword, clause.multiWord, clause.Name)

//u.Warnf("nil next state")

//u.Debugf("found next state")

//u.Debugf("nil lexer but matches? repeat?%v isrepeat?%v  name=%q", clause.Repeat, repeat, clause.Name)

// Before we move into child clause, lets check and see if we need to descend
//  into sub-clauses of current statement

//u.Debugf("has sub-clauses  kw=%-10s peek=%-10s", clause.keyword, peekWord)

//u.Debugf("has sub-clauses  kw=%-10s peek=%-10s", sc.keyword, peekWord)

//u.Debugf("matches Sub-Clause: kw=%-15q  %-15q", sc.keyword, sc.Name)

// First non-optional we don't match we bail

// if we didn't match

// we haven't tried to repeat yet

//u.Debugf("repeating clause: %v", clause.keyword)
// Before we move onto next clause, lets check and see if we need to descend
//  into sub-clauses of current statement

//u.Debugf("has sub-clauses  kw=%-10s peek=%-10s", clause.keyword, peekWord)

//u.Infof("has sub-clauses  kw=%-10s peek=%-10s", sc.keyword, peekWord)

//u.Debugf("matches Sub-Clause: %-10s", sc.keyword)

// First non-optional we don't match we bail

//u.Warnf("nil?  %#v ", clause.parent)

//u.Infof("moving to next parent clause: %v", clause.keyword)

// If we have consumed all clauses, we are ready to be done?
//u.Debugf("not found? word? '%s' %v", peekWord, clause)

//u.Infof("%p Run End of statement", l)

// Correctly reached EOF.

// What is this?

//u.Infof("end of statement")

// Doesn't actually lex anything, used for single token clauses
func LexEmpty(l *Lexer) StateFn {
	_ = "STUB: not implemented"

	// lex a value:   string, integer, float
	//
	// - literal strings must be quoted
	// - numerics with no period are integers
	// - numerics with period are floats
	//
	//	"stuff"    -> [string] = stuff
	//	'stuff'    -> [string] = stuff
	//	"items's with quote" -> [string] = items's with quote
	//	1.23  -> [float] = 1.23
	//	100   -> [integer] = 100
	//	["hello","world"]  -> [array] {"hello","world"}
	return *new(StateFn)
}

func LexValue(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexValue: rune=%v  peek:%v", string(rune), l.PeekX(10))

// this is a mistake and should not happen

//panic("should not have paren")

// quoted string, allows escaping

// consume the quote mark

//u.Debugf("LexValue rune=%v  end?%v  prevEscape?%v  quote=%s", string(rune), rune == eof, previousEscaped, string(firstRune))
//(rune == '\'' || rune == '"') &&

// check for escaped quote mark

// since we read lookahead after escape/quote that ends the string

// for single quote which is not part of the value

// now ignore that single quote

// at the very end

// Was escaped   \"

// if we are on next rune, then previous \ was NOT an escape, cancel

// Non-Quoted String?   Should this be a numeric?   or date or what?  duration?  what kinds are valid?
//  A:   numbers

// lets look for Booleans

// lex a regex:   first character must be a /
//
//	/^stats\./i
//	/.*/
//	/^stats.*/
func LexRegex(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

// scan looking for ending character = /

//u.Debugf("LexRegex rune=%v  end?%v  prevEscape?%v", string(rune), rune == eof, previousEscaped)

// now that we have found what appears to be end, lets see if it
// has a modifier - the i/g at end of    /^stats\./i

//u.Debugf("LexRegex rune=%v  end?%v  prevEscape?%v", string(rune), rune == eof, previousEscaped)

//return l.errorToken("expected value but got EOF")

// look for either an Expression or Identity
//
//	expressions:    Legal identity characters, terminated by (
//	identity:    legal identity characters
//
//	REPLACE(name,"stuff")
//	name
func LexExpressionOrIdentity(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

// u.Debugf("LexExpressionOrIdentity identity?%v expr?%v %v peek5='%v'", l.isIdentity(), l.isExpr(), string(l.Peek()), string(l.PeekX(5)))
// Expressions end in Parens:     LOWER(item)

// Non Expressions are Identities, or Columns
// by passing nil here, we are going to go back to Pull items off stack)

//u.Warnf("LexExpressionOrIdentity ??? -> LexValue")

// look for either an Identity or Value
func LexIdentityOrValue(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexIdentityOrValue identity?%v expr?%v %v peek5='%v'", l.isIdentity(), l.isExpr(), string(l.Peek()), string(l.PeekX(5)))
// Expressions end in Parens:     LOWER(item)

// Non Expressions are Identities, or Columns
//u.Warnf("in expr is identity? %s", l.PeekWord())
// by passing nil here, we are going to go back to Pull items off stack)

//u.Warnf("LexIdentityOrValue ??? -> LexValue")

// lex Expression looks for an expression, identified by parenthesis, may be nested
//
//	       |--expr----|
//	dostuff(name,"arg")    // the left parenthesis identifies it as Expression
//	eq(trim(name," "),"gmail.com")
func LexExpressionParens(l *Lexer) StateFn {
	_ = "STUB: not implemented"
	return *

	// first rune must be opening Parenthesis
	new(StateFn)
}

//u.Debugf("LexExpressionParens:  %v", string(firstChar))

//u.Infof("LexExpressionParens:   %v", string(firstChar))

// LexUrnaryNot NOT
func LexUrnaryNot(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

// LexParenRight:  look for end of paren, of which we have descended and consumed start
func LexParenRight(l *Lexer) StateFn {
	_ = "STUB: not implemented"

	// first rune must be closing Parenthesis
	return *new(StateFn)
}

//u.Debugf("LexParenRight:  %v  eof?%v", string(r), r == eof)

// ascend

// LexParenLeft:  look for end of paren, of which we have descended and consumed start
func LexParenLeft(l *Lexer) StateFn {
	_ = "STUB: not implemented"
	return *

	// first rune must be opening Parenthesis
	new(StateFn)
}

//u.Debugf("LexParenLeft:  %v", string(r))

//u.Infof("LexParenLeft:   %v", string(r))
// ascend

// lex expression identity keyword, does not consume parenthesis
//
//	|--expridentity---|
//	name_of_expression(name,"arg")
func lexExpressionIdentifier(l *Lexer) StateFn {
	_ = "STUB: not implemented"
	return *

	// u.Debugf("lexExpressionIdentifier identity?%v expr?%v %v:%v", l.isIdentity(), l.isExpr(), string(l.Peek()), string(l.PeekWord()))
	new(StateFn)
}

// first rune has to be valid unicode letter

//u.Warnf("lexExpressionIdentifier couldnt find expression idenity?  %v stack=%v", string(firstChar), len(l.stack))

// Now look for run of runes, where run is ended by first non-identifier character

// iterate until we find non-identifer character

// TODO:  validate identity vs next keyword?, ie ensure it is not a keyword/reserved word

// back up one character

// LexListOfArgs list of arguments, comma separated list of args which
// may be a mixture of expressions, identities, values
//
//	REPLACE(LOWER(x),"xyz")
//	REPLACE(x,"xyz")
//	COUNT(*)
//	sum( 4 * toint(age))
//	IN (a,b,c)
//	varchar(10)
//	CAST(field AS int)
//
//	(a,b,c,d)   -- For Insert statement, list of columns
func LexListOfArgs(l *Lexer) StateFn {
	_ = "STUB: not implemented"

	// as we descend into Expressions, we are going to use push/pop to
	//
	//	ascend/descend
	return *new(StateFn)
}

//u.Debugf("in LexListOfArgs:  '%s'", string(r))

//l.Emit(TokenRightParenthesis)

// Send signal to pop

//l.Push("LexParenRight", LexParenRight)

// So, not comma, * so either is Expression, Identity, Value

//u.Debugf("in LexListOfArgs:  '%s'", peekWord)
// First, lets ensure we haven't blown past into keyword?

//u.Warnf("found keyword while looking for arg? %v", string(r))

//u.Debugf("LexListOfArgs sending LexExpressionOrIdentity: %v", string(peekWord))

// LexIdentifier scans and finds named things (tables, columns)
//
//	and specifies them as TokenIdentity, uses LexIdentifierType
//
//	TODO: dialect controls escaping/quoting techniques
//
//	[name]         select [first name] from usertable;
//	'name'         select 'user' from usertable;
//	first_name     select first_name from usertable;
//	usertable      select first_name AS fname from usertable;
//	_name          select _name AS name from stuff;
var LexIdentifier = LexIdentifierOfType(TokenIdentity)
var LexTableIdentifier = LexIdentifierOfType(TokenTable)

// LexIdentifierOfType scans and finds named things (tables, columns)
//
//	supports quoted, bracket, or raw identifiers
//
//	TODO: dialect controls escaping/quoting techniques
//
//	[name]         select [first name] from usertable;
//	'name'         select 'user' from usertable;
//	`user`         select first_name from `user`;
//	first_name     select first_name from usertable;
//	usertable      select first_name AS fname from usertable;
//	_name          select _name AS name from stuff;
//	@@varname      select @@varname;
func LexIdentifierOfType(forToken TokenType) StateFn {
	_ = "STUB: not implemented"
	return *new(StateFn)
}

func lexIdentifierOfTypeNoWs(l *Lexer, shouldIgnore bool, forToken TokenType) StateFn {
	_ = "STUB: not implemented"

	// first rune has to be valid unicode letter or @@
	return *new(StateFn)
}

//u.Debugf("LexIdentifierOfType:   '%s' ='?%v peek6'%v'", string(firstChar), firstChar == '\'', l.PeekX(6))

// Fields can be bracket or single quote escaped
//  [user]
//  [email]
//  'email'
//  `email`

//u.Debugf("lex firstChar: %s  %s", string(firstChar), string(nextChar))

// Empty Identity = value?  not really an identity is it?

// Since we escaped this with a quote we lex until unescaped end?

// TODO:  escaping?

// Identity of form   [schema].[table]
//u.Warnf("%s", l.RawInput())

// Identity of form   `schema`.`table`
//u.Warnf("%s", l.RawInput())

// iterate until we find non-identifier, then make sure it is valid/end

// valid

// also valid

//u.Warnf("aborting LexIdentifier: '%v'", string(firstChar))

// iterate until we find non-identifer character

// Special case
//   content.`field name`

//u.Debugf("about to emit: %v", forToken)

// need to skip last character bc it was quoted

// pop up to parent

var LexDataTypeDefinition = LexDataType(TokenTypeDef)

// LexDataType scans and finds datatypes.
// `[]` are valid inside of data types, no escaping such as ',"
//
//	[]string       CREATE table( field []string )
//	map[string]int
//	int, string, etc
func LexDataType(forToken TokenType) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexDataType: %v", l.PeekX(5))

// Since we escaped this with a quote we allow laxIdentifier characters

//u.Infof("r=%v %v    ws=%v", string(r), r, isWhiteSpace(r))

// ok, continue

//ok, continue

// LexEndOfStatement Look for end of statement defined by either a semicolon or end of file
func LexEndOfStatement(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Warnf("END:  %s %s", string(r), l.PeekX(20))

// LexSelectClause Handle start of select statements, specifically looking for
// @@variables, *, or else we drop into <select_list>
//
//	<SELECT> :==
//	    (DISTINCT|ALL)? ( <sql_variable> | * | <select_list> ) [FROM <source_clause>]
//
//	<sql_variable> = @@stuff
func LexSelectClause(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexSelectClause  '%v'", word)

//ALL?

// DISTINCTROW?

// Look for keyword, ie something like FROM, or possibly end of statement
// consume the *
// this will skip whitespace
//u.Debugf("* ?'%v'  keyword='%v'", first, pw)

//   select * from

//u.Warnf("What is this? %v", l.PeekX(10))

//  mysql system variables start with @@

//u.Debugf("Found Sql Variable:  @@%v", word)

//u.Debugf("not found %v", first)

//u.Debugf("Found Sql Variable:  @%v", word)

// Since we did Not find anything it, start lexing normal SelectList

// Handle start of insert, Upsert statements
func LexUpsertClause(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexUpsertClause  '%v'  %v", word, l.PeekX(10))

// Handle recursive subqueries
func LexSubQuery(l *Lexer) StateFn {
	_ = "STUB: not implemented"

	// u.Debugf("LexSubQuery  '%v'", l.PeekX(10))
	return *new(StateFn)
}

/*
	TODO:   this is a hack because the LexDialect from above should be recursive,
	 	ie support sub-queries, but doesn't currently
*/

// Handle prepared statements
//
// <PREPARE_STMT> := PREPARE <identity>	FROM <string_value>
func LexPreparedStatement(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexPreparedStatement  '%v'", l.PeekX(10))

/*
	TODO:   this is a bit different from others, as after we get FROM
	 we are going to create a new lexer?  and forward over?  or reset?
*/

// Handle repeating Select List for columns
//
//	   SELECT <select_list>
//
//	   <select_list> := <select_col> [, <select_col>]*
//
//	   <select_col> :== ( <identifier> | <expression> | '*' ) [AS <identifier>] [IF <expression>] [<comment>]
//
//	Note, our Columns support a non-standard IF guard at a per column basis
func LexSelectList(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexSelectList looking for operator:  word=%q", word)

// Handle Source References ie [From table], [SubSelects], Joins
//
//	SELECT ...  FROM <sources>
//
//	<sources>      := <source> [, <join_clause> <source>]*
//	<source>       := ( <table_source> | <subselect> ) [AS <identifier>]
//	<table_source> := <identifier>
//	<join_clause>  := (INNER | LEFT | OUTER)? JOIN [ON <conditional_clause>]
//	<subselect>    := '(' <select_stmt> ')'
func LexTableReferenceFirst(l *Lexer) StateFn {
	_ = "STUB: not implemented"

	// From has already been consumed
	return *new(StateFn)
}

//u.Debugf("LexTableReferenceFirst  peek2= '%v'  isEnd?%v", l.PeekX(2), l.IsEnd())

// Cover the grouping, ie recursive/repeating nature of subqueries

// subquery

//u.Debugf("LexTableReferenceFirst looking for operator:  word=%s", word)

// nice, this is what we are looking for, let dialect take over

// are there other functions besides in?

//u.Warnf("found keyword? %v ", word)

//u.LogTracef(u.WARN, "hmmmmmmm")
//u.Debugf("LexTableReferenceFirst = '%v'", string(r))
// ensure we don't get into a recursive death spiral here?

// Since we did Not find anything, we are going to go for a Expression or Identity

// Handle Source References ie [From table], [SubSelects], Joins
//
//	SELECT ...  FROM <sources>
//
//	<sources>      := <source> [, <join_clause> <source>]*
//	<source>       := ( <table_source> | <subselect> ) [AS <identifier>]
//	<table_source> := <identifier>
//	<join_clause>  := (INNER | LEFT | OUTER)? JOIN [ON <conditional_clause>]
//	<subselect>    := '(' <select_stmt> ')'
func LexTableReferences(l *Lexer) StateFn {
	_ = "STUB: not implemented"

	// From has already been consumed
	return *new(StateFn)
}

//u.Debugf("LexTableReferences  peek2= '%v'  isEnd?%v", l.PeekX(2), l.IsEnd())

// Cover the grouping, ie recursive/repeating nature of subqueries

// subquery?

//u.Debugf("LexTableReferences looking for operator:  word=%s", word)

// TODO:  need to allow the Dialect Statements to be recursive/nested
// l.ConsumeWord("SELECT")
// l.Emit(TokenSelect)

//u.Warnf("emit from")
// l.ConsumeWord("FROM")
// l.Emit(TokenFrom)
// l.Push("LexTableReferences", LexTableReferences)
// l.Push("LexIdentifier", LexIdentifier)

//l.Push("LexTableReferences", LexTableReferences)
//l.Push("LexExpression", LexExpression)

//

// what is complete list here?

//u.Warnf("found keyword? %v ", word)

//return LexExpressionOrIdentity

//u.LogTracef(u.WARN, "hmmmmmmm")
//u.Debugf("LexTableReferences = '%v'", string(r))
// ensure we don't get into a recursive death spiral here?

// Since we did Not find anything, we are going to go for a Expression or Identity

// Handle Source References ie [From table], [SubSelects], Joins
//
//	SELECT ...  FROM <sources>
//
//	<sources>      := <source> [, <join_clause> <source>]*
//	<source>       := ( <table_source> | <subselect> ) [AS <identifier>]
//	<table_source> := <identifier>
//	<join_clause>  := (INNER | LEFT | OUTER)? JOIN [ON <conditional_clause>]
//	<subselect>    := '(' <select_stmt> ')'
func LexJoinEntry(l *Lexer) StateFn {
	_ = "STUB: not implemented"
	return *

	// u.Debugf("LexJoinEntry  peek2= '%v'  isEnd?%v", l.PeekX(2), l.IsEnd())
	new(StateFn)
}

// Cover the grouping, ie recursive/repeating nature of subqueries

// subquery?
//l.Push("LexJoinEntry", LexJoinEntry)
//return LexSelectClause

//u.Debugf("LexJoinEntry looking for operator:  word=%s", word)

// case "in":
// 	return nil

//u.Debugf("expression or identity?")

//u.LogTracef(u.WARN, "hmmmmmmm")
//u.Debugf("LexJoinEntry = '%v'", string(r))
// ensure we don't get into a recursive death spiral here?

// Since we did Not find anything, we are going to go for a Expression or Identity

// LexColumnNames Handle list of column names on insert/update statements
//
//	<insert_into> <col_names> VALUES <col_value_list>
//
//	<col_names> := '(' <identity> [, <identity>]* ')'
func LexColumnNames(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexColumnNames lr=%s  word=%q", string(r), l.PeekWord())

// Handle repeating Insert/Upsert/Update statements
//
//	<insert_into> <col_names> VALUES <col_value_list>
//	<set> <upsert_cols> VALUES <col_value_list>
//
//	<upsert_cols> := <upsert_col> [, <upsert_col>]*
//	<upsert_col> := <identity> = <expr>
//
//	<col_names> := <identity> [, <identity>]*
//	<col_value_list> := <col_value_row> [, <col_value_row>] *
//
//	<col_value_row> := '(' <expr> [, <expr>]* ')'
func LexTableColumns(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexTableColumns  r= '%v'", string(r))

//u.Debugf("looking for tablecolumns:  word=%s r=%s", word, string(r))

// TODO:  this is returning because l.clauseState()

// LexValueColumns   VALUES (a,b,c),(d,e,f);
func LexValueColumns(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexValueColumns  r= %v'", string(r))

//l.Next()
//l.Emit(TokenRightParenthesis)

// LexConditionalClause Handle logical Conditional Clause used for [WHERE, WITH, JOIN ON]
// logicaly grouped with parens and/or separated by commas or logic (AND/OR/NOT)
//
//	SELECT ... WHERE <conditional_clause>
//
//	<conditional_clause> ::= <expr> [( AND <expr> | OR <expr> | '(' <expr> ')' )]
//
//	<expr> ::= <predicatekw> '('? <expr> [, <expr>] ')'? | <func> | <subselect>
//
// SEE:  <expr> = LexExpression
func LexConditionalClause(l *Lexer) StateFn {
	_ = "STUB: not implemented"

	// u.Debugf("lexConditional: %v", l.PeekX(14))
	return *new(StateFn)
}

//l.Next()
//l.Emit(TokenRightParenthesis)
//l.Push("LexParenRight", LexParenRight)

//u.Debugf("lexConditional word: %v", word)

//u.LogThrottle(u.WARN, 5, "sure we want subQuery here? %v", word)

//u.Infof("is keyword %v", word)

//u.Debugf("go to lex expression: %v", l.PeekX(20))

// Alias for Expression
func LexColumns(l *Lexer) StateFn {
	_ = "STUB: not implemented"
	return *

	// LexLogical is a lex entry function for logical expression language (+-/> etc)
	//
	//	ie, the full logical boolean logic
	new(StateFn)
}

func LexLogical(l *Lexer) StateFn {
	_ = "STUB: not implemented"

	// u.Debugf("in lexLogical: peek: %q  end? %v", l.PeekX(5), l.IsEnd())
	return *new(StateFn)
}

// <expr>   Handle single logical expression which may be nested and  has
//
//	user defined function names that are NOT validated by lexer
//
// <expr> ::= <predicatekw> '('? <expr> [, <expr>] ')'? | <func> | <subselect>
//
//	<func> ::= <identity>'(' <expr> ')'
//	<predicatekw> ::= [NOT] (IN | INTERSECTS | CONTAINS | RANGE | LIKE | EQUALS )
//
// Examples:
//
//	(colx = y OR colb = b)
//	cola = 'a5'
//	cola != "a5", colb = "a6"
//	REPLACE(cola,"stuff") != "hello"
//	FirstName = REPLACE(LOWER(name," "))
//	cola IN (1,2,3)
//	cola LIKE "abc"
//	eq(name,"bob") AND age > 5
//	time > now() -1h
//	(4 + 5) > 10
//	reg_date BETWEEN x AND y
func LexExpression(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

// Cover the logic and grouping

// this is a logical Grouping/Ordering and must be a single
// logically valid expression

// this is a logical Grouping/Ordering
//l.Emit(TokenRightParenthesis)
// don't consume )

// comment?  or minus?

//  !=

//   <>

// x = 5 * 5

// u.Debugf("LexExpression operator:  word=%q  kw?%v", word, l.isNextKeyword(word))

// what is complete list here?

// somewhat weird edge case, not is either word not, or expression
//
//  not( x == y)        -- this is a function called not()
//  NOT (x == y)        -- this is an urnary expression NOT
//  x NOT IN ("a","b")  -- this is a binary negateable expression natural not
//  AND NOT(hasprefix(event,"gh."))

//  not( x == y)        -- this is a function called NOT
//u.Warnf("EXPR not(<expr>)  %q", string(px))

// TODO:  make this the normal, requires parser changes
// l.ConsumeWord(word)
// l.Emit(TokenNegate)
// l.Push("LexParenRight", LexParenRight)
// l.Push("LexExpression", LexExpression)
// return LexParenLeft

// THIS IS WORKING, but trying to deprecate it
// Consume word NOT

//return l.clauseState()

//  NOT (x == y)        -- this is an urnary expression
//  NOT (user_id IN ("a","b"))        -- this is an urnary expression
//u.Infof("not (<expr>)  %q   peek20:%v", string(pr), l.PeekX(20))

//  x NOT IN ("a","b")  -- this is a binary negateable expression natural not
//u.Debugf("not (in,contains,include) ?  %q", string(pr))

// this marks beginning of new related column

// ensure we don't get into a recursive death spiral here?

// Handle columnar identies with keyword appendate (ASC, DESC)
//
//	[ORDER BY] ( <identity> | <expr> ) [(ASC | DESC)]
func LexOrderByColumn(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexOrderBy  r= '%v'  %v", string(r), l.PeekX(10))

//u.Debugf("word: %v", word)

//u.Debugf("looking for operator:  word=%s", word)

// Since we did Not find anything, we are in error?

// Lex either Json or Key/Value pairs
//
//	Must start with { or [ for json
//	Start with identity for key/value pairs
func LexJsonOrKeyValue(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexJsonOrKeyValue  '%v'  %v", string(r), l.PeekX(10))

//u.Warnf("Did not find json? %v", l.PeekX(20))

// Lex Valid Json
//
//	Must start with { or [
func LexJson(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexJson  '%v'  %v", string(r), l.PeekX(10))

//u.Warnf("Did not find json? %v", l.PeekX(20))

/*
	TokenLeftBracket  TokenType = 23 // [
	TokenRightBracket TokenType = 24 // ]
	TokenLeftBrace    TokenType = 25 // {
	TokenRightBrace   TokenType = 26 // }
*/

// LexJsonValue:  Consume values, first consuming Colon
//
//	<jsonvalue> ::= ':' ( <value>, <array>, <jsonobject> ) [, ...]
func LexJsonValue(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexJsonValue  '%v'  %v", string(r), l.PeekX(10))

// recurse back up one level
//<object>

// Key's must be strings
//<array>

// Lex Valid Json Array
//
//	Must End with ]
func LexJsonArray(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexJsonArray  '%v' peek:%v", string(r), l.PeekX(10))

// consume ,

// consume {

// Key's must be strings

// consume [

// value
//u.Debugf("call lex value: %v", l.PeekX(10))

// Lex Valid Json Object
//
//	Must End with }
func LexJsonObject(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexJsonObject  '%v'  %v", string(r), l.PeekX(10))

// consume :

// consume ,

// consume {

// Key's must be strings

// consume [

// lex a string value value:
//
//	strings must be quoted
//
//	"stuff"    -> stuff
//	"items's with quote"
func LexJsonIdentity(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("in LexJsonIdentity: %v", string(rune))
// quoted string

// consume the quote mark

//u.Debugf("LexValue rune=%v  end?%v  prevEscape?%v", string(rune), rune == eof, previousEscaped)

// check for '''

// since we read lookahead after single quote that ends the string
// for lookahead

// for single quote which is not part of the value

// now ignore that single quote

// at the very end

// LexComment looks for valid comments which are any of the following
//
//	 including the in-line comment blocks
//
//	/* hello */
//	//  hello
//	-- hello
//	# hello
//	SELECT name --name is the combined first-last name
//	       , age FROM `USER` ...
func LexComment(l *Lexer) StateFn {
	_ = "STUB: not implemented"
	// u.Debugf("checking comment: '%s' ", l.input[l.pos:l.pos+2])
	// TODO:  switch statement instead of strings has prefix
	return *new(StateFn)
}

//u.Debugf("found single line comment:  // ")

//u.Debugf("found single line comment:  -- ")

//u.Debugf("found single line comment:  # ")

// A multi-line comment of format /* comment */
// it does not have to actually be multi-line, just surrounded by those comments
func LexMultilineComment(l *Lexer) StateFn {
	_ = "STUB: not implemented"
	// Consume opening "/*"
	return *new(StateFn)
}

// Consume trailing "*/"

// Comment beginning with //, # or --
func LexInlineComment(l *Lexer) StateFn {
	_ = "STUB: not implemented"

	// We are going to Find the start of the Comments
	return *new(StateFn)
}

// Should we be emitting the --, #, // ?  is that meaningful?

// Comment beginning with //, # or -- but do not emit the tag just text comment
func LexInlineCommentNoTag(l *Lexer) StateFn {
	_ = "STUB: not implemented"

	// We are going to Find the start of the Comments
	return *new(StateFn)
}

// Should we be emitting the --, #, // ?  is that meaningful?

// we have consumed it

// Consume the word

// Should we actually be consuming Whitespace? or is it meaningful?
// l.SkipWhiteSpaces()

// for {
// 	r = l.Next()
// 	if r == '\n' || r == eof {
// 		l.backup()
// 		break
// 	}
// }
// l.Emit(TokenComment)

// the text/contents of a single line comment
func lexSingleLineComment(l *Lexer) StateFn {
	_ = "STUB: not implemented"
	// Should we consume whitespace?
	// l.SkipWhiteSpaces()
	return *new(StateFn)
}

// LexNumber floats, integers, hex, exponential, signed
//
//	1.23
//	100
//	-827
//	6.02e23
//	0X1A2B,  0x1a2b, 0x1A2B.2B
//
// Floats must be in decimal and must either:
//
//   - Have digits both before and after the decimal point (both can be
//     a single 0), e.g. 0.5, -100.0, or
//   - Have a lower-case e that represents scientific notation,
//     e.g. -3e-3, 6.02e23.
//
// Integers can be:
//
//   - decimal (e.g. -827)
//   - hexadecimal (must begin with 0x and must use capital A-F, e.g. 0x1A2B)
func LexNumber(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("typ  %v   %v  %q", typ, ok, l.input[l.start:l.pos])

// Emits tokenFloat or tokenInteger.

// LexNumberOrDuration floats, integers, hex, exponential, signed
//
//	1.23
//	100
//	-827
//	6.02e23
//	0X1A2B,  0x1a2b, 0x1A2B.2B
//
// durations:   45m, 2w, 20y, 22d, 40ms, 100ms, -100ms
//
// Floats must be in decimal and must either:
//
//   - Have digits both before and after the decimal point (both can be
//     a single 0), e.g. 0.5, -100.0, or
//   - Have a lower-case e that represents scientific notation,
//     e.g. -3e-3, 6.02e23.
//
// Integers can be:
//
//   - decimal (e.g. -827)
//   - hexadecimal (must begin with 0x and must use capital A-F, e.g. 0x1A2B)
func LexNumberOrDuration(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

// LexDuration floats, integers time-durations
//
// durations:   45m, 2w, 20y, 22d, 40ms, 100ms, -100ms
func LexDuration(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

// scan for a number
//
// It returns the scanned tokenType (tokenFloat or tokenInteger) and a flag
// indicating if an error was found.
func scanNumber(l *Lexer) (typ TokenType, ok bool) {
	_ = "STUB: not implemented"
	return *new(TokenType), false
}

// scan for a number
//
// It returns the scanned tokenType (tokenFloat or tokenInteger) and a flag
// indicating if an error was found.
func scanNumericOrDuration(l *Lexer, doDuration bool) (typ TokenType, ok bool) {
	_ = "STUB: not implemented"

	// Optional leading sign.
	return *new(TokenType), false
}

//u.Debugf("scanNumericOrDuration?  '%v' peek:%v ", string(peek2), len(peek2))

// Hexadecimal.

// No signs for hexadecimals.

// Requires at least one digit.

// No dots for hexadecimals.

// Decimal

// Requires at least one digit

// Float

// Requires a digit after the dot.

// Integers can't start with 0??

// A digit is required after the scientific notation.

// duration was found

// Next thing must not be alphanumeric.

// Helpers --------------------------------------------------------------------

// is Alpha Numeric reports whether r is an alphabetic, digit, or underscore.
func isAlNum(r rune) bool { _ = "STUB: not implemented"; return false }

// is Alpha reports whether r is an alphabetic, or underscore or period
func isAlpha(r rune) bool { _ = "STUB: not implemented"; return false }

// is Alpha Numeric reports whether r is an alphabetic, digit, or underscore, or period
func isAlNumOrPeriod(r rune) bool { _ = "STUB: not implemented"; return false }

func isDigit(r rune) bool { _ = "STUB: not implemented"; return false }

func isWhiteSpace(r rune) bool { _ = "STUB: not implemented"; return false }

// IsBreak is some character such as comma, ;, etc
func IsBreak(r rune) bool { _ = "STUB: not implemented"; return false }

// Is the given rune valid in an identifier?
func isIdentCh(r rune) bool { _ = "STUB: not implemented"; return false }

// IsIdentifierRune Is this a valid identity rune?
func IsIdentifierRune(r rune) bool { _ = "STUB: not implemented"; return false }

func isIdentifierFirstRune(r rune) bool { _ = "STUB: not implemented"; return false }

// Digits can not lead identities

// are we really going to support this globally as identity?

func isLaxIdentifierRune(r rune) bool { _ = "STUB: not implemented"; return false }

func isJsonStart(r rune) bool { _ = "STUB: not implemented"; return false }

// IsValidIdentity test the given string to determine if any characters are
// not valid and therefore must be quoted
func IsValidIdentity(identity string) bool { _ = "STUB: not implemented"; return false }

func IdentityRunesOnly(identity string) bool { _ = "STUB: not implemented"; return false }
