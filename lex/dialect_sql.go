package lex

var (

	// SqlDialect is a SQL dialect
	//
	//    SELECT
	//    UPDATE
	//    INSERT
	//    UPSERT
	//    DELETE
	//
	//    SHOW idenity;
	//    DESCRIBE identity;
	//    PREPARE
	//
	// ddl
	//    ALTER
	//    CREATE (TABLE|VIEW|CONTINUOUSVIEW|SOURCE)
	//
	//  TODO:
	//      CREATE
	//      VIEW
	SqlDialect *Dialect = &Dialect{
		Statements: []*Clause{
			{Token: TokenPrepare, Clauses: SqlPrepare},
			{Token: TokenSelect, Clauses: SqlSelect},
			{Token: TokenUpdate, Clauses: SqlUpdate},
			{Token: TokenUpsert, Clauses: SqlUpsert},
			{Token: TokenInsert, Clauses: SqlInsert},
			{Token: TokenDelete, Clauses: SqlDelete},
			{Token: TokenCreate, Clauses: SqlCreate},
			{Token: TokenDrop, Clauses: SqlDrop},
			{Token: TokenAlter, Clauses: SqlAlter},
			{Token: TokenDescribe, Clauses: SqlDescribe},
			{Token: TokenExplain, Clauses: SqlExplain},
			{Token: TokenDesc, Clauses: SqlDescribeAlt},
			{Token: TokenShow, Clauses: SqlShow},
			{Token: TokenSet, Clauses: SqlSet},
			{Token: TokenUse, Clauses: SqlUse},
			{Token: TokenRollback, Clauses: SqlRollback},
			{Token: TokenCommit, Clauses: SqlCommit},
		},
	}
	// SqlSelect Select statement.
	SqlSelect = []*Clause{
		{Token: TokenSelect, Lexer: LexSelectClause, Name: "sqlSelect.Select"},
		{Token: TokenInto, Lexer: LexInto, Optional: true, Name: "sqlSelect.INTO"},
		{Token: TokenFrom, Lexer: LexTableReferenceFirst, Optional: true, Repeat: false, Clauses: fromSource, Name: "sqlSelect.From"},
		{KeywordMatcher: sourceMatch, Optional: true, Repeat: true, Clauses: moreSources, Name: "sqlSelect.sources"},
		{Token: TokenWhere, Lexer: LexConditionalClause, Optional: true, Clauses: whereQuery, Name: "sqlSelect.where"},
		{Token: TokenGroupBy, Lexer: LexColumns, Optional: true, Name: "sqlSelect.groupby"},
		{Token: TokenHaving, Lexer: LexConditionalClause, Optional: true, Name: "sqlSelect.having"},
		{Token: TokenOrderBy, Lexer: LexOrderByColumn, Optional: true, Name: "sqlSelect.orderby"},
		{Token: TokenLimit, Lexer: LexLimit, Optional: true, Name: "sqlSelect.limit"},
		{Token: TokenOffset, Lexer: LexNumber, Optional: true, Name: "sqlSelect.offset"},
		{Token: TokenWith, Lexer: LexJsonOrKeyValue, Optional: true, Name: "sqlSelect.with"},
		{Token: TokenAlias, Lexer: LexIdentifier, Optional: true, Name: "sqlSelect.alias"},
		{Token: TokenEOF, Lexer: LexEndOfStatement, Optional: false, Name: "sqlSelect.eos"},
	}
	fromSource = []*Clause{
		{KeywordMatcher: sourceMatch, Lexer: LexTableReferenceFirst, Name: "fromSource.matcher"},
		{Token: TokenSelect, Lexer: LexSelectClause, Name: "fromSource.Select"},
		{Token: TokenFrom, Lexer: LexTableReferenceFirst, Optional: true, Repeat: true, Name: "fromSource.From"},
		{Token: TokenWhere, Lexer: LexConditionalClause, Optional: true, Name: "fromSource.Where"},
		{Token: TokenHaving, Lexer: LexConditionalClause, Optional: true, Name: "fromSource.having"},
		{Token: TokenGroupBy, Lexer: LexColumns, Optional: true, Name: "fromSource.GroupBy"},
		{Token: TokenOrderBy, Lexer: LexOrderByColumn, Optional: true, Name: "fromSource.OrderBy"},
		{Token: TokenLimit, Lexer: LexLimit, Optional: true, Name: "fromSource.Limit"},
		{Token: TokenOffset, Lexer: LexNumber, Optional: true, Name: "fromSource.Offset"},
		{Token: TokenRightParenthesis, Lexer: LexEndOfSubStatement, Optional: true, Name: "fromSource.EndParen"},
		{Token: TokenAs, Lexer: LexIdentifier, Optional: true, Name: "fromSource.As"},
		{Token: TokenOn, Lexer: LexConditionalClause, Optional: true, Name: "fromSource.On"},
	}
	moreSources = []*Clause{
		{KeywordMatcher: sourceMatch, Lexer: LexJoinEntry, Name: "moreSources.JoinEntry"},
		{Token: TokenSelect, Lexer: LexSelectClause, Optional: true, Name: "moreSources.Select"},
		{Token: TokenFrom, Lexer: LexTableReferenceFirst, Optional: true, Repeat: true, Name: "moreSources.From"},
		{Token: TokenWhere, Lexer: LexConditionalClause, Optional: true, Name: "moreSources.Where"},
		{Token: TokenHaving, Lexer: LexConditionalClause, Optional: true, Name: "moreSources.Having"},
		{Token: TokenGroupBy, Lexer: LexColumns, Optional: true, Name: "moreSources.GroupBy"},
		{Token: TokenOrderBy, Lexer: LexOrderByColumn, Optional: true, Name: "moreSources.OrderBy"},
		{Token: TokenLimit, Lexer: LexLimit, Optional: true, Name: "moreSources.Limit"},
		{Token: TokenOffset, Lexer: LexNumber, Optional: true, Name: "moreSources.Offset"},
		{Token: TokenRightParenthesis, Lexer: LexEndOfSubStatement, Optional: false, Name: "moreSources.EndParen"},
		{Token: TokenAs, Lexer: LexIdentifier, Optional: true, Name: "moreSources.As"},
		{Token: TokenOn, Lexer: LexConditionalClause, Optional: true, Name: "moreSources.On"},
	}
	whereQuery = []*Clause{
		{Token: TokenSelect, Lexer: LexSelectClause, Name: "whereQuery.Select"},
		{Token: TokenFrom, Lexer: LexTableReferences, Optional: true, Repeat: true, Name: "whereQuery.From"},
		{Token: TokenWhere, Lexer: LexConditionalClause, Optional: true, Name: "whereQuery.Where"},
		{Token: TokenHaving, Lexer: LexConditionalClause, Optional: true, Name: "whereQuery.Having"},
		{Token: TokenGroupBy, Lexer: LexColumns, Optional: true, Name: "whereQuery.GroupBy"},
		{Token: TokenOrderBy, Lexer: LexOrderByColumn, Optional: true, Name: "whereQuery.OrderBy"},
		{Token: TokenLimit, Lexer: LexNumber, Optional: true, Name: "whereQuery.Limit"},
		{Token: TokenRightParenthesis, Lexer: LexEndOfSubStatement, Optional: false, Name: "whereQuery.EOS"},
	}
	// SqlUpdate update statement
	SqlUpdate = []*Clause{
		{Token: TokenUpdate, Lexer: LexIdentifierOfType(TokenTable)},
		{Token: TokenSet, Lexer: LexColumns},
		{Token: TokenWhere, Lexer: LexColumns, Optional: true},
		{Token: TokenLimit, Lexer: LexNumber, Optional: true},
		{Token: TokenWith, Lexer: LexJsonOrKeyValue, Optional: true},
	}
	// SqlUpsert sql upsert
	SqlUpsert = []*Clause{
		{Token: TokenUpsert, Lexer: LexUpsertClause, Name: "upsert.entry"},
		{Token: TokenSet, Lexer: LexTableColumns, Optional: true},
		{Token: TokenLeftParenthesis, Lexer: LexTableColumns, Optional: true},
		{Token: TokenWith, Lexer: LexJsonOrKeyValue, Optional: true},
	}
	// SqlInsert insert statement
	SqlInsert = []*Clause{
		{Token: TokenInsert, Lexer: LexUpsertClause, Name: "insert.entry"},
		{Token: TokenLeftParenthesis, Lexer: LexColumnNames, Optional: true},
		{Token: TokenSet, Lexer: LexTableColumns, Optional: true},
		{Token: TokenSelect, Optional: true, Clauses: insertSubQuery},
		{Token: TokenValues, Lexer: LexTableColumns, Optional: true},
		{Token: TokenWith, Lexer: LexJsonOrKeyValue, Optional: true},
	}
	insertSubQuery = []*Clause{
		{Token: TokenSelect, Lexer: LexSelectClause},
		{Token: TokenFrom, Lexer: LexTableReferences, Optional: true, Repeat: true},
		{Token: TokenWhere, Lexer: LexConditionalClause, Optional: true},
		{Token: TokenHaving, Lexer: LexConditionalClause, Optional: true},
		{Token: TokenGroupBy, Lexer: LexColumns, Optional: true},
		{Token: TokenOrderBy, Lexer: LexOrderByColumn, Optional: true},
		{Token: TokenLimit, Lexer: LexNumber, Optional: true},
	}
	// SqlReplace replace statement
	SqlReplace = []*Clause{
		{Token: TokenReplace, Lexer: LexEmpty},
		{Token: TokenInto, Lexer: LexIdentifierOfType(TokenTable)},
		{Token: TokenSet, Lexer: LexTableColumns, Optional: true},
		{Token: TokenLeftParenthesis, Lexer: LexTableColumns, Optional: true},
		{Token: TokenWith, Lexer: LexJsonOrKeyValue, Optional: true},
	}
	// SqlDelete delete statement
	SqlDelete = []*Clause{
		{Token: TokenDelete, Lexer: LexEmpty},
		{Token: TokenFrom, Lexer: LexIdentifierOfType(TokenTable)},
		{Token: TokenSet, Lexer: LexColumns, Optional: true},
		{Token: TokenWhere, Lexer: LexColumns, Optional: true},
		{Token: TokenLimit, Lexer: LexNumber, Optional: true},
		{Token: TokenWith, Lexer: LexJsonOrKeyValue, Optional: true},
	}
	// SqlAlter alter statement
	SqlAlter = []*Clause{
		{Token: TokenAlter, Lexer: LexEmpty},
		{Token: TokenTable, Lexer: LexIdentifier},
		{Token: TokenChange, Lexer: LexDdlAlterColumn},
		{Token: TokenWith, Lexer: LexJsonOrKeyValue, Optional: true},
	}
	// SqlCreate CREATE {SCHEMA | DATABASE | SOURCE | TABLE | VIEW | CONTINUOUSVIEW}
	SqlCreate = []*Clause{
		{Token: TokenCreate, Lexer: LexCreate},
		{Token: TokenEngine, Lexer: LexDdlTableStorage, Optional: true},
		{Token: TokenSelect, Clauses: SqlSelect, Optional: true},
		{Token: TokenWith, Lexer: LexJsonOrKeyValue, Optional: true},
	}
	// SqlDrop DROP {SCHEMA | DATABASE | SOURCE | TABLE}
	SqlDrop = []*Clause{
		{Token: TokenDrop, Lexer: LexDrop},
	}
	// SqlDescribe Describe {table,database}
	SqlDescribe = []*Clause{
		{Token: TokenDescribe, Lexer: LexColumns},
	}
	// SqlDescribeAlt alternate spelling of Describe
	SqlDescribeAlt = []*Clause{
		{Token: TokenDesc, Lexer: LexColumns},
	}
	// SqlExplain is alias of describe
	SqlExplain = []*Clause{
		{Token: TokenExplain, Lexer: LexColumns},
	}
	// SqlShow
	SqlShow = []*Clause{
		{Token: TokenShow, Lexer: LexShowClause},
		{Token: TokenWhere, Lexer: LexConditionalClause, Optional: true},
	}
	// SqlPrepare
	SqlPrepare = []*Clause{
		{Token: TokenPrepare, Lexer: LexPreparedStatement},
		{Token: TokenFrom, Lexer: LexTableReferences},
	}
	// SqlSet
	SqlSet = []*Clause{
		{Token: TokenSet, Lexer: LexColumns},
	}
	// SqlUse
	SqlUse = []*Clause{
		{Token: TokenUse, Lexer: LexIdentifier},
	}
	// SqlRollback
	SqlRollback = []*Clause{
		{Token: TokenRollback, Lexer: LexEmpty},
	}
	// SqlCommit
	SqlCommit = []*Clause{
		{Token: TokenCommit, Lexer: LexEmpty},
	}
)

// NewSqlLexer creates a new lexer for the input string using SqlDialect
// this is sql(ish) compatible parser.
func NewSqlLexer(input string) *Lexer { _ = "STUB: not implemented"; return nil }

// find any keyword that starts a source
//
//	FROM <name>
//	FROM (select ...)
//	     [(INNER | LEFT)] JOIN
func sourceMatch(c *Clause, peekWord string, l *Lexer) bool {
	_ = "STUB: not implemented"
	//u.Debugf("%p sourceMatch?   peekWord: %s", c, peekWord)
	return false
}

// LexEndOfSubStatement Look for end of statement defined by either
// a semicolon or end of file.
func LexEndOfSubStatement(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

// LexShowClause Handle show statement
//
//	SHOW [FULL] <multi_word_identifier> <identity> <like_or_where>
func LexShowClause(l *Lexer) StateFn {
	_ = "STUB: not implemented"

	/*
	   SHOW {BINARY | MASTER} LOGS
	   SHOW BINLOG EVENTS [IN 'log_name'] [FROM pos] [LIMIT [offset,] row_count]
	   SHOW CHARACTER SET [like_or_where]
	   SHOW COLLATION [like_or_where]
	   SHOW [FULL] COLUMNS FROM tbl_name [FROM db_name] [like_or_where]
	   SHOW CREATE DATABASE db_name
	   SHOW CREATE EVENT event_name
	   SHOW CREATE FUNCTION func_name
	   SHOW CREATE PROCEDURE proc_name
	   SHOW CREATE TABLE tbl_name
	   SHOW CREATE TRIGGER trigger_name
	   SHOW CREATE VIEW view_name
	   SHOW DATABASES [like_or_where]
	   SHOW ENGINE engine_name {STATUS | MUTEX}
	   SHOW [STORAGE] ENGINES
	   SHOW ERRORS [LIMIT [offset,] row_count]
	   SHOW EVENTS
	   SHOW FUNCTION CODE func_name
	   SHOW FUNCTION STATUS [like_or_where]
	   SHOW GRANTS FOR user
	   SHOW INDEX FROM tbl_name [FROM db_name]
	   SHOW MASTER STATUS
	   SHOW OPEN TABLES [FROM db_name] [like_or_where]
	   SHOW PLUGINS
	   SHOW PROCEDURE CODE proc_name
	   SHOW PROCEDURE STATUS [like_or_where]
	   SHOW PRIVILEGES
	   SHOW [FULL] PROCESSLIST
	   SHOW PROFILE [types] [FOR QUERY n] [OFFSET n] [LIMIT n]
	   SHOW PROFILES
	   SHOW SLAVE HOSTS
	   SHOW SLAVE STATUS [NONBLOCKING]
	   SHOW [GLOBAL | SESSION] STATUS [like_or_where]
	   SHOW TABLE STATUS [FROM db_name] [like_or_where]
	   SHOW [FULL] TABLES [FROM db_name] [like_or_where]
	   SHOW TRIGGERS [FROM db_name] [like_or_where]
	   SHOW [GLOBAL | SESSION] VARIABLES [like_or_where]
	   SHOW WARNINGS [LIMIT [offset,] row_count]

	   like_or_where:
	       LIKE 'pattern'
	     | WHERE expr
	*/return *new(StateFn)
}

// u.Debugf("LexShowClause  r= '%v'", string(keyWord))

// TODO:  these should not be identities but tokens?

// SHOW CREATE TABLE tbl_name

// LexInto clause
func LexInto(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexInto  r= '%v'", string(keyWord))

// LexLimit clause
//
//	LIMIT 1000 OFFSET 100
//	LIMIT 0, 1000
//	LIMIT 1000
func LexLimit(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexLimit  r= '%v'", string(keyWord))

// LexCreate allows us to lex the words after CREATE
//
//	CREATE {SCHEMA|DATABASE|SOURCE} [IF NOT EXISTS] <identity>  <WITH>
//	CREATE {TABLE} <identity> [IF NOT EXISTS] <table_spec> [WITH]
//	CREATE [OR REPLACE] {VIEW|CONTINUOUSVIEW} <identity> AS <select_statement> [WITH]
func LexCreate(l *Lexer) StateFn {
	_ = "STUB: not implemented"

	/*
		CREATE TABLE [IF NOT EXISTS] <identity> [WITH]
		CREATE SOURCE [IF NOT EXISTS] <identity> [WITH]
		CREATE [OR REPLACE] VIEW <identity> AS <select_statement> [WITH]
	*/return *new(StateFn)
}

//u.Debugf("LexCreate  r= '%v'", string(keyWord))

func lexAs(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

func lexNotExists(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("lexNotExists  r= '%v'", string(keyWord))

func lexOrReplace(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("lexOrReplace  r= '%v'", string(keyWord))

// LexDrop allows us to lex the words after DROP
//
//	DROP {DATABASE | SCHEMA} [IF EXISTS] db_name
//
//	DROP [TEMPORARY] TABLE [IF EXISTS] tbl_name [, tbl_name] [RESTRICT | CASCADE]
//
//	DROP INDEX index_name ON tbl_name
//	    [algorithm_option | lock_option] ...
func LexDrop(l *Lexer) StateFn {
	_ = "STUB: not implemented"

	/*
		DROP {DATABASE | SCHEMA} [IF EXISTS] db_name

		DROP INDEX index_name ON tbl_name
		    [algorithm_option | lock_option] ...

		algorithm_option:
		    ALGORITHM [=] {DEFAULT|INPLACE|COPY}

		lock_option:
		    LOCK [=] {DEFAULT|NONE|SHARED|EXCLUSIVE}
	*/return *new(StateFn)
}

//u.Debugf("LexCreate  r= '%v'", string(keyWord))

// LexDdlTable data definition language table
func LexDdlTable(l *Lexer) StateFn {
	_ = "STUB: not implemented"

	/*
		CREATE [TEMPORARY] TABLE [IF NOT EXISTS] tbl_name
		    (create_definition,...)
		    [table_options]
		    [partition_options]

		CREATE [TEMPORARY] TABLE [IF NOT EXISTS] tbl_name
		    [(create_definition,...)]
		    [table_options]
		    [partition_options]
		    [IGNORE | REPLACE]
		    [AS] query_expression

		CREATE [TEMPORARY] TABLE [IF NOT EXISTS] tbl_name
		    { LIKE old_tbl_name | (LIKE old_tbl_name) }

		CREATE TABLE `City` (
		  `ID` int(11) NOT NULL AUTO_INCREMENT,
		  `Name` char(35) NOT NULL DEFAULT '',
		  `CountryCode` char(3) NOT NULL DEFAULT '',
		  `District` char(20) NOT NULL DEFAULT '',
		  `Population` int(11) NOT NULL DEFAULT '0',
		  PRIMARY KEY (`ID`),
		  KEY `CountryCode` (`CountryCode`),
		  CONSTRAINT `city_ibfk_1` FOREIGN KEY (`CountryCode`)
		     REFERENCES `Country` (`Code`)
		) ENGINE=InnoDB AUTO_INCREMENT=4080 DEFAULT CHARSET=utf8
	*/return *new(StateFn)
}

//u.Debugf("LexDdlTable  r= '%v'", string(r))

// Cover the logic and grouping

// Start of columns

// end of columns

// comment?

//u.Debugf("looking table col start:  word=%s", word)

// ensure we don't get into a recursive death spiral here?

// LexDdlTableStorage data definition language column (repeated)
//
//	ENGINE=InnoDB AUTO_INCREMENT=4080 DEFAULT CHARSET=utf8
func LexDdlTableStorage(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

// LexDdlAlterColumn data definition language column alter
//
//	CHANGE col1_old col1_new varchar(10),
//	CHANGE col2_old col2_new TEXT
//	ADD col3 BIGINT AFTER col1_new
//	ADD col2 TEXT FIRST,
func LexDdlAlterColumn(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexDdlAlterColumn  r= '%v'", string(r))

// Cover the logic and grouping

// comment?

//u.Debugf("looking for operator:  word=%s", word)

// Character set is end of ddl column
// character set

// Below here are Data Types

//u.Infof("found keyword? %v ", word)

//u.LogTracef(u.WARN, "hmmmmmmm")
//u.Infof("LexDdlAlterColumn = '%v'", string(r))

// ensure we don't get into a recursive death spiral here?

// LexDdlTableColumn data definition language column (repeated)
//
//	col1_new varchar(10),
//	col2_new TEXT
func LexDdlTableColumn(l *Lexer) StateFn {
	_ = "STUB: not implemented"

	/*
		http://dev.mysql.com/doc/refman/5.7/en/create-table.html

		create_definition:
		    col_name column_definition
		  | [CONSTRAINT [symbol]] PRIMARY KEY [index_type] (index_col_name,...)
		      [index_option] ...
		  | {INDEX|KEY} [index_name] [index_type] (index_col_name,...)
		      [index_option] ...
		  | [CONSTRAINT [symbol]] UNIQUE [INDEX|KEY]
		      [index_name] [index_type] (index_col_name,...)
		      [index_option] ...
		  | {FULLTEXT|SPATIAL} [INDEX|KEY] [index_name] (index_col_name,...)
		      [index_option] ...
		  | [CONSTRAINT [symbol]] FOREIGN KEY
		      [index_name] (index_col_name,...) reference_definition
		  | CHECK (expr)

		column_definition:
		    data_type [NOT NULL | NULL] [DEFAULT default_value]
		      [AUTO_INCREMENT] [UNIQUE [KEY] | [PRIMARY] KEY]
		      [COMMENT 'string']
		      [COLUMN_FORMAT {FIXED|DYNAMIC|DEFAULT}]
		      [STORAGE {DISK|MEMORY|DEFAULT}]
		      [reference_definition]
		  | data_type [GENERATED ALWAYS] AS (expression)
		      [VIRTUAL | STORED] [UNIQUE [KEY]] [COMMENT comment]
		      [NOT NULL | NULL] [[PRIMARY] KEY]
	*/return *new(StateFn)
}

//u.Debugf("LexDdlTableColumn  r= '%v'  peek: %s", string(r), l.PeekX(20))

//

// comment?

//u.Debugf("looking for ddl col start:  word=%s", word)

// Character set is end of ddl column
// case "character": // character set
// 	cs := strings.ToLower(l.PeekX(len("character set")))
// 	if cs == "character set" {
// 		l.ConsumeWord(cs)
// 		l.Emit(TokenCharacterSet)
// 		l.Push("LexDdlTableColumn", LexDdlTableColumn)
// 		return nil
// 	}
// Below here are Data Types

// LexEngineKeyValue key value pairs
//
//	Start with identity for key/value pairs
//	supports keyword DEFAULT
//	supports non-quoted values
func LexEngineKeyValue(l *Lexer) StateFn { _ = "STUB: not implemented"; return *new(StateFn) }

//u.Debugf("LexEngineKeyValue  %q  peek= %v", word, l.PeekX(10))
