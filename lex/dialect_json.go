package lex

// NewJsonLexer Creates a new json dialect lexer for the input string.
func NewJsonLexer(input string) *Lexer { _ = "STUB: not implemented"; return nil }

var (
	jsonDialectStatement = []*Clause{
		{Token: TokenNil, Lexer: LexJson},
	}
	// JsonDialect, is a json lexer
	//
	//    ["hello","world"]
	//    {"name":"bob","apples":["honeycrisp","fuji"]}
	//
	JsonDialect *Dialect = &Dialect{
		Statements: []*Clause{
			{Token: TokenNil, Clauses: jsonDialectStatement},
		},
	}
)
