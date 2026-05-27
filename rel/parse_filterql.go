package rel

import (
	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/lex"
)

// FilterQLParser
type FilterQLParser struct {
	// can be a FilterStatement, FilterStatements, filterSelect, filterSelects, etc.
	// Which one is determined by which Parser func you call.
	statement string
	fs        *FilterStatement
	l         *lex.Lexer
	comment   string
	*filterTokenPager
	firstToken lex.Token
	funcs      expr.FuncResolver
}

func NewFilterParser(filter string) *FilterQLParser { _ = "STUB: not implemented"; return nil }

func NewFilterParserfuncs(filter string, funcs expr.FuncResolver) *FilterQLParser {
	_ = "STUB: not implemented"
	return nil
}

// FuncResolver sets the function resolver to use during parsing.  By default we only use the Global resolver.
// But if you set a function resolver we'll use that first and then fall back to the Global resolver.
func (f *FilterQLParser) FuncResolver(funcs expr.FuncResolver) *FilterQLParser {
	_ = "STUB: not implemented"
	return nil
}

func (f *FilterQLParser) setLexer(statement string) { _ = "STUB: not implemented"; return }

// ParseFilterQL Parses a FilterQL statement
func (f *FilterQLParser) ParseFilter() (*FilterSelect, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FilterQLParser) ParseFilters() (stmts []*FilterStatement, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FilterQLParser) ParseFilterSelects() (stmts []*FilterSelect, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseFilters Parse a list of Filter statement's from text
func ParseFilters(statement string) (stmts []*FilterStatement, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MustParseFilters(statement string) []*FilterStatement { _ = "STUB: not implemented"; return nil }

func MustParseFilter(statement string) *FilterStatement { _ = "STUB: not implemented"; return nil }

// ParseFilterQL Parses a FilterQL statement
func ParseFilterQL(filter string) (*FilterStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseFilterSelect Parse a single Select-Filter statement from text
// Select-Filters are statements of following form
//
//	"SELECT" [COLUMNS] (FILTER | WHERE) FilterExpression
//	"FILTER" FilterExpression
func ParseFilterSelect(query string) (*FilterSelect, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseFilterSelects Parse 1-n Select-Filter statements from text
// Select-Filters are statements of following form
//
//	"SELECT" [COLUMNS] (FILTER | WHERE) FilterExpression
//	"FILTER" FilterExpression
func ParseFilterSelects(statement string) (stmts []*FilterSelect, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type (
	// TokenPager is responsible for determining end of current clause
	//   An interface used to allow Parser to be neutral to dialect
	filterTokenPager struct {
		*expr.LexTokenPager
		lastKw lex.TokenType
	}
)

func newFilterTokenPager(l *lex.Lexer) *filterTokenPager { _ = "STUB: not implemented"; return nil }

func (m *filterTokenPager) IsEnd() bool { _ = "STUB: not implemented"; return false }

func (m *filterTokenPager) ClauseEnd() bool { _ = "STUB: not implemented"; return false }

// List of possible tokens that would indicate a end to the current clause

func (m *FilterQLParser) parseFilterStart() (*FilterStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *FilterQLParser) parseSelectStart() (*FilterSelect, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *FilterQLParser) initialComment() string { _ = "STUB: not implemented"; return "" }

// We are going to loop until we find the first Non-Comment Token

// skip, currently ignore these

// first non-comment token

func (m *FilterQLParser) discardNewLines() {
	_ = "STUB: not implemented"

	// We are going to loop until we find the first Non-NewLine
	return
}

// first non-comment token

func (m *FilterQLParser) discardCommentsNewLines() {
	_ = "STUB: not implemented"

	// We are going to loop until we find the first Non-Comment Token
	return
}

// skip, currently ignore these

// first non-comment token

func (m *FilterQLParser) discardComments() {
	_ = "STUB: not implemented"

	// We are going to loop until we find the first Non-Comment Token
	return
}

// skip, currently ignore these

// first non-comment token

// First keyword was SELECT, so use the SELECT parser rule-set
func (m *FilterQLParser) parseSelect() (*FilterSelect, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Consume Select

// optional FROM

// We accept either WHERE or FILTER

// one top level filter which may be nested

// one top level filter which may be nested

// LIMIT  - Optional

// WITH  - Optional

// ALIAS  - Optional

// First keyword was FILTER, so use the FILTER parser rule-set
func (m *FilterQLParser) parseFilter() (*FilterStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Consume (FILTER | WHERE )

// one top level filter which may be nested

// OPTIONAL From clause

// LIMIT - Optional

// WITH - Optional

// ALIAS - Optional

//, lex.TokenRightParenthesis

func (m *FilterQLParser) parseWhereExpr(req *FilterSelect) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *FilterQLParser) parseFirstFilters() (expr.Node, error) {
	_ = "STUB: not implemented"

	// We have 2 special cases in filterQL
	// FILTER *
	// FILTER match_all
	return *new(expr.Node), nil
}

// Consume *

// if we have match all, nothing else allowed

// if we have match all, nothing else allowed

func (m *FilterQLParser) parseLimit() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (m *FilterQLParser) parseAlias() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Consume ALIAS token

func (m *FilterQLParser) isEnd() bool { _ = "STUB: not implemented"; return false }
