package lex

var (
	// FilterStatement a FilterQL statement.
	FilterStatement = []*Clause{
		{Token: TokenFilter, Lexer: LexFilterClause, Optional: true},
		{Token: TokenFrom, Lexer: LexIdentifier, Optional: true},
		{Token: TokenLimit, Lexer: LexNumber, Optional: true},
		{Token: TokenWith, Lexer: LexJsonOrKeyValue, Optional: true},
		{Token: TokenAlias, Lexer: LexIdentifier, Optional: true},
		{Token: TokenEOF, Lexer: LexEndOfStatement, Optional: false},
	}
	// FilterSelectStatement Filter statement that also supports column projection.
	FilterSelectStatement = []*Clause{
		{Token: TokenSelect, Lexer: LexSelectClause, Optional: false},
		{Token: TokenFrom, Lexer: LexIdentifier, Optional: false},
		{Token: TokenWhere, Lexer: LexConditionalClause, Optional: true},
		{Token: TokenFilter, Lexer: LexFilterClause, Optional: true},
		{Token: TokenLimit, Lexer: LexNumber, Optional: true},
		{Token: TokenWith, Lexer: LexJsonOrKeyValue, Optional: true},
		{Token: TokenAlias, Lexer: LexIdentifier, Optional: true},
		{Token: TokenEOF, Lexer: LexEndOfStatement, Optional: false},
	}
	// FilterQLDialect is a Where Clause filtering language slightly
	// more DSL'ish than SQL Where Clause.
	FilterQLDialect *Dialect = &Dialect{
		Statements: []*Clause{
			{Token: TokenFilter, Clauses: FilterStatement},
			{Token: TokenSelect, Clauses: FilterSelectStatement},
		},
		IdentityQuoting: IdentityQuotingWSingleQuote,
	}
)

// NewFilterQLLexer creates a new lexer for the input string using FilterQLDialect
// which is dsl for where/filtering.
func NewFilterQLLexer(input string) *Lexer { _ = "STUB: not implemented"; return nil }

// LexFilterClause Handle Filter QL Main Statement
//
//	FILTER := ( <filter_bool_expr> | <filter_expr> )
//
//	<filter_bool_expr> :=  ( AND | OR ) '(' ( <filter_bool_expr> | <filter_expr> ) [, ( <filter_bool_expr> | <filter_expr> ) ] ')'
//
//	<filter_expr> :=  <expr>
//
// Examples:
//
//	FILTER
//
// /      AND (
//
//	      daysago(datefield) < 100
//	      , domain(url) == "google.com"
//	      , INCLUDE name_of_filter
//	      ,
//	      , OR (
//	          momentum > 20
//	         , propensity > 50
//	      )
//	   )
//	ALIAS myfilter
//
//	FILTER x > 7
func LexFilterClause(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }
