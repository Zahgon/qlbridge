package influxql

import (
	"github.com/araddon/qlbridge/lex"
)

var (
	// Tokens Specific to INFLUXDB
	TokenShortDesc lex.TokenType = 1000
	TokenLongDesc  lex.TokenType = 1001
	TokenKind      lex.TokenType = 1002
)

var selectQl = []*lex.Clause{
	{Token: lex.TokenSelect, Lexer: LexColumnsInflux},
	{Token: lex.TokenFrom, Lexer: LexInfluxName},
	{Token: lex.TokenGroupBy, Lexer: lex.LexColumns, Optional: true},
	{Token: lex.TokenLimit, Lexer: lex.LexNumber, Optional: true},
	{Token: lex.TokenInto, Lexer: lex.LexExpressionOrIdentity},
	{Token: lex.TokenWhere, Lexer: lex.LexColumns, Optional: true},
}

var InfluxQlDialect *lex.Dialect = &lex.Dialect{
	Statements: []*lex.Clause{
		{Token: lex.TokenSelect, Clauses: selectQl},
	},
}

func init() {
	lex.TokenNameMap[TokenShortDesc] = &lex.TokenInfo{Description: "SHORTDESC"}
	lex.TokenNameMap[TokenLongDesc] = &lex.TokenInfo{Description: "LONGDESC"}
	lex.TokenNameMap[TokenKind] = &lex.TokenInfo{Description: "kind"}
	// OverRide the Identity Characters in QLparse
	lex.IDENTITY_CHARS = "_./-"
	lex.LoadTokenInfo()
	InfluxQlDialect.Init()
}

// Handle influx columns
//
//	SELECT
//	     valuect(item) AS stuff SHORTDESC "stuff" KIND INT
//
// Examples:
//
//	(colx = y OR colb = b)
//	cola = 'a5'p
//	cola != "a5", colb = "a6"
//	REPLACE(cola,"stuff") != "hello"
//	FirstName = REPLACE(LOWER(name," "))
//	cola IN (1,2,3)
//	cola LIKE "abc"
//	eq(name,"bob") AND age > 5
func LexColumnsInflux(l *lex.Lexer) lex.StateFn {
	_ = "STUB: not implemented"
	return *new(lex.StateFn)
}

// lex value
//
//	SIMPLE_NAME_VALUE | TABLE_NAME_VALUE | REGEX_VALUE
func LexInfluxName(l *lex.Lexer) lex.StateFn { _ = "STUB: not implemented"; return *new(lex.StateFn) }

// a regex
