package influxql

import (
	"github.com/araddon/qlbridge/lex"
)

/*

Parser for InfluxDB ql

*/

// Parses string query
func Parse(query string) (*Ast, error) { _ = "STUB: not implemented"; return nil, nil }

// parser evaluateslex.Tokens
type Parser struct {
	l              *lex.Lexer
	qryText        string
	initialKeyword lex.Token
	curToken       lex.Token
}

// parse the request
func (m *Parser) parse() (*Ast, error) { _ = "STUB: not implemented"; return nil, nil }

// Now, find First Keyword

func (m *Parser) initialComment() string { _ = "STUB: not implemented"; return "" }

// We are going to loop until we find the first Non-Commentlex.Token

// skip, currently ignore these

// first non-commentlex.Token

// First keyword was SELECT, so use the SELECT parser rule-set
func (m *Parser) parseSelect(comment string) (*Ast, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// we have already parsed SELECT lex.Token to get here, so this should be first col

// * mark as star?

// FROM - required

// table/metric

// Where is optional

// limit is optional
// if err := m.parseLimit(&selStmt); err != nil {
// 	return nil, err
// }

// we are finished, nice!

func (m *Parser) parseColumns(stmt *SelectStmt) error { _ = "STUB: not implemented"; return nil }

func (m *Parser) parseWhere(stmt *SelectStmt) error {
	_ = "STUB: not implemented"

	// Where is Optional, if we didn't use a where statement return
	return nil
}
