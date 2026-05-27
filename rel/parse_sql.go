package rel

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/lex"
)

var (
	// SqlKeywords the SqlKeywords list
	SqlKeywords = []string{"select", "insert", "update", "delete", "from", "where", "as", "into", "limit",
		"exists", "in", "contains", "include", "not", "and", "having", "or", "null", "group", "order",
		"offset", "include", "all", "any", "some"}
)

// ParseError type
type ParseError struct {
	error
}

// ParseSql Parses SqlStatement and returns a statement or error
// does not parse more than one statement
func ParseSql(sqlQuery string) (SqlStatement, error) {
	_ = "STUB: not implemented"
	return *new(SqlStatement), nil
}

func parseSqlResolver(sqlQuery string, fr expr.FuncResolver) (SqlStatement, error) {
	_ = "STUB: not implemented"
	return *new(SqlStatement), nil
}

// ParseSqlSelect parse a sql statement as SELECT (or else error)
func ParseSqlSelect(sqlQuery string) (*SqlSelect, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseSqlSelectResolver parse as SELECT using function resolver.
func ParseSqlSelectResolver(sqlQuery string, fr expr.FuncResolver) (*SqlSelect, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseSqlStatements into array of SQL Statements
func ParseSqlStatements(sqlQuery string) ([]SqlStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Sqlbridge generic SQL parser evaluates should be sufficient for most
// sql compatible languages
type Sqlbridge struct {
	l       *lex.Lexer
	comment string
	*SqlTokenPager
	firstToken lex.Token
	funcs      expr.FuncResolver
}

// parse the request
func (m *Sqlbridge) parse() (SqlStatement, error) {
	_ = "STUB: not implemented"
	return *new(SqlStatement), nil
}

func readComment(p expr.TokenPager) string { _ = "STUB: not implemented"; return "" }

// We are going to loop until we find the first Non-Comment Token

// skip, currently ignore these

// first non-comment token

func discardComments(m expr.TokenPager) {
	_ = "STUB: not implemented"

	// We are going to loop until we find the first Non-Comment Token
	return
}

// discard

// first non-comment token

// First keyword was SELECT, so use the SELECT parser rule-set
func (m *Sqlbridge) parseSqlSelect() (*SqlSelect, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Consume Select?

// Optional DISTINCT keyword always immediately after SELECT KW

// columns

// select @@myvar limit 1

// SPECIAL END CASE for simple selects
// SELECT last_insert_id();

// valid end

// INTO

// FROM

// WHERE

// GROUP BY

// HAVING

// ORDER BY

// LIMIT

// OFFSET

// WITH

// ALIAS

// we are good

// First keyword was INSERT, REPLACE
func (m *Sqlbridge) parseSqlInsert() (*SqlInsert, error) {
	_ = "STUB: not implemented"

	// insert into mytable (id, str) values (0, "a")
	return nil, nil
}

// Consume Insert or Replace

// INTO

// Consume INTO

// table name

// list of fields

// left paren starts lisf of values

// Consume Values keyword

// First keyword was UPDATE
func (m *Sqlbridge) parseSqlUpdate() (*SqlUpdate, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Consume UPDATE token
}

//u.Debugf("token:  %v", m.Cur())

// list of name=value pairs

// WHERE

// First keyword was UPSERT
func (m *Sqlbridge) parseSqlUpsert() (*SqlUpsert, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Consume UPSERT token

// consume Into

// Consume Set
// list of name=value pairs

// list of fields

// left paren starts lisf of values

// Consume Values keyword

//u.Debugf("found ?  %v", m.Cur())

// WHERE

// First keyword was DELETE
func (m *Sqlbridge) parseSqlDelete() (*SqlDelete, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Consume Delete
}

// from
//u.Debugf("token:  %v", m.Cur())

// table name

//u.Debugf("found table?  %v", m.Cur())

//u.Debugf("cur lex.Token: %s", m.Cur().T.String())

// we are good

// First keyword was PREPARE
func (m *Sqlbridge) parsePrepare() (*PreparedStatement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Consume Prepare

// statement name/alias
//u.Debugf("found table?  %v", m.Cur())

// from

//u.Debugf("token:  %v", m.Cur())

// we are good

// First keyword was DESCRIBE
func (m *Sqlbridge) parseDescribe() (SqlStatement, error) {
	_ = "STUB: not implemented"
	return *new(SqlStatement), nil
}

// Consume Describe

//u.Debugf("token:  %v", m.Cur())

// TODO:  make the lexer handle this

// First keyword was SHOW
func (m *Sqlbridge) parseShow() (*SqlShow, error) {
	_ = "STUB: not implemented"

	/*
		don't currently support all these
		http://dev.mysql.com/doc/refman/5.7/en/show.html

		SHOW [FULL] COLUMNS FROM tbl_name [FROM db_name] [like_or_where]
		SHOW CREATE DATABASE db_name
		SHOW CREATE TABLE tbl_name
		SHOW CREATE TRIGGER trigger_name
		SHOW CREATE VIEW view_name
		SHOW DATABASES [like_or_where]
		SHOW ENGINE engine_name {STATUS | MUTEX}
		SHOW [STORAGE] ENGINES
		SHOW INDEX FROM tbl_name [FROM db_name]
		SHOW [FULL] TABLES [FROM db_name] [like_or_where]
		SHOW TRIGGERS [FROM db_name] [like_or_where]
		SHOW [GLOBAL | SESSION] VARIABLES [like_or_where]
		SHOW [GLOBAL | SESSION | SLAVE] STATUS [like_or_where]
		SHOW WARNINGS [LIMIT [offset,] row_count]
	*/return nil, nil
}

// Consume Show

//u.Infof("cur: %v", m.Cur())

//u.Infof("scope:%q   next:%v", req.Scope, m.Cur())

// SHOW CREATE TABLE `temp_schema`.`users`

// consume create

//u.Debugf("create what %v", m.Cur())
// {TABLE | DATABASE | EVENT ...}
//u.Debugf("create which %v", m.Cur())

//u.Debugf("show %v", m.Cur())

// consume columns

//SHOW [FULL] COLUMNS {FROM | IN} tbl_name [{FROM | IN} db_name]  [LIKE 'pattern' | WHERE expr]
// | Field      | Type     | Null | Key | Default | Extra          |

// consume Tables
// SHOW [FULL] TABLES [FROM db_name] [like_or_where]

// SHOW TABLES LIKE '%'
// Consume Like

// consume where

// First keyword was SET, USE
func (m *Sqlbridge) parseCommand() (*SqlCommand, error) {
	_ = "STUB: not implemented"

	/*
		- SET CHARACTER SET utf8
		- SET NAMES utf8
	*/return nil, nil
}

// USE, SET

// USE `baseball`;

// Look for special cases for mysql weird SET syntax

//SET NAMES utf8
// consume NAMES

// consume character
// consume set

// First keyword was CREATE
func (m *Sqlbridge) parseCreate() (*SqlCreate, error) { _ = "STUB: not implemented"; return nil, nil }

// Consume CREATE token

// Consume OR

// CREATE {DATABASE|SCHEMA|TABLE|VIEW|SOURCE|CONTINUOUSVIEW} <identity>

// Grab remainder which will be SELECT (we have already lexed AS)

// [IF NOT EXISTS]

// Consume IF

// consume paren

// list of columns comma separated

// ENGINE

// just with

// just with for now

// WITH

// First keyword was DROP
func (m *Sqlbridge) parseDrop() (*SqlDrop, error) {
	_ = "STUB: not implemented"
	return nil,

		// Consume DROP token
		nil
}

// DROP TEMPORARY TABLE x

// DROP (TABLE|VIEW|SOURCE|CONTINUOUSVIEW) <identity>

// triggers, indexes

// just table

// schema

// view

// triggers, index, etc

// WITH

func (m *Sqlbridge) parseTransaction() (*SqlCommand, error) {
	_ = "STUB: not implemented"

	// rollback, commit
	return nil, nil
}

// rollback, commit

func parseColumns(m expr.TokenPager, fr expr.FuncResolver, stmt ColumnsStatement) error {
	_ = "STUB: not implemented"
	return nil
}

//u.Debug(m.Cur())

// we have a udf/functional expression column

// function canoncial names are always lowercase

//u.Debugf("udf? %T ", col.Expr)

//u.Debugf("next? %v", m.Cur())

// Value Literal

//u.Debugf("after colstart?:   %v  ", m.Cur())

// since we can loop inside switch statement

// This indicates we have come to the End of the columns

// If guard

// Hm, we need to backup here?  Parse Node went to deep?

// loop on my friend

func (m *Sqlbridge) parseFieldList() (Columns, error) {
	_ = "STUB: not implemented"
	return *new(Columns), nil
}

//u.Debug(m.Cur().String())

//u.Debugf("after colstart?:   %v  ", m.Cur())

// since we can loop inside switch statement

func (m *Sqlbridge) parseUpdateList() (map[string]*ValueColumn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//u.Debugf("col:%v    cur:%v", lastColName, m.Cur().String())

// don't need to do anything

// TODO:  this is a bug in lexer

func (m *Sqlbridge) parseValueList() ([][]*ValueColumn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//u.Debug(m.Cur().String())

// start of row

// TODO:  this is a bug in lexer

// error?

// an array of values?
// Consume the [

// don't need to do anything

func (m *Sqlbridge) parseSources(req *SqlSelect) error { _ = "STUB: not implemented"; return nil }

// consume From

//u.Debugf("parseSources %v", m.Cur())

// SELECT [columns] FROM [table] AS t1
//   INNER JOIN (select a,b,c from users WHERE d is not null) u ON u.user_id = t1.user_id

// JOIN

// Skip over As, we don't need it

// select u.name, order.date FROM user AS u INNER JOIN ....

func (m *Sqlbridge) parseSourceSubQuery(src *SqlSource) error {
	_ = "STUB: not implemented"

	// page forward off of (
	return nil
}

// SELECT * FROM (SELECT 1, 2, 3) AS t1;

// discard right paren

func (m *Sqlbridge) parseSourceTable(req *SqlSelect) error { _ = "STUB: not implemented"; return nil }

// Skip over "AS", we don't need it

func (m *Sqlbridge) parseSourceJoin(src *SqlSource) error { _ = "STUB: not implemented"; return nil }

// Optional Inner/Outer

// Consume join keyword

// SELECT [columns] FROM [table] AS t1
//   INNER JOIN (select a,b,c from users WHERE d is not null) u ON u.user_id = t1.user_id

// Name of table

func (m *Sqlbridge) parseInto(req *SqlSelect) error { _ = "STUB: not implemented"; return nil }

// Consume Into token

func (m *Sqlbridge) parseWhereSubSelect(req *SqlSelect) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *Sqlbridge) parseWhereSelect(req *SqlSelect) error { _ = "STUB: not implemented"; return nil }

// TODO this is deeply flawed, need to fix/use tokenpager
//    with rewind ability

func (m *Sqlbridge) parseWhere() (*SqlWhere, error) { _ = "STUB: not implemented"; return nil, nil }

// Consume the Where
//u.Debugf("cur: %v peek=%v", m.Cur(), m.Peek())

// We are going to Peek forward at the next 3 tokens used
// to determine which type of where clause
// x

// Check for Types of Where
//                                 t1            T2      T3     T4
//    SELECT x FROM user   WHERE user_id         IN      (      SELECT user_id from orders where ...)
//    SELECT * FROM t1     WHERE column1         =       (      SELECT column1 FROM t2);
//    select a FROM movies WHERE director        IN      (     "Quentin","copola","Bay","another")
//    select b FROM movies WHERE director        =       "bob";
//    select b FROM movies WHERE create          BETWEEN "2015" AND "2010";
//    select b from movies WHERE director        LIKE    "%bob"
// TODO:
//    SELECT * FROM t3     WHERE ROW(5*t2.s1,77) =       (      SELECT 50,11*s1 FROM t4)

//u.Infof("in parseWhere: %v", m.Cur())
// T1  ?? this might be udf?
// t2  (IN | =)
// t3 = (
//m.Next() // t4 = SELECT

func (m *Sqlbridge) parseGroupBy(req *SqlSelect) (err error) { _ = "STUB: not implemented"; return nil }

//u.Debugf("Group By? %v", m.Cur())

// we have a udf/functional expression column
//u.Infof("udf: %v", m.Cur().V)

//u.Debugf("udf? %T ", n)

//u.Debugf("next? %v", m.Cur())

//u.Warnf("?? %v", m.Cur())

// Value Literal

//u.Debugf("GroupBy after colstart?:   %v  ", m.Cur())

// since we can loop inside switch statement

//u.Debug(m.Cur())

//u.Infof("set AS=%v", col.As)

// This indicates we have come to the End of the columns

// If guard

// loop on my friend

func (m *Sqlbridge) parseHaving(req *SqlSelect) (err error) { _ = "STUB: not implemented"; return nil }

// TODO this is deeply flawed, need to fix/use tokenpager
// with rewind ability

func (m *Sqlbridge) parseOrderBy(req *SqlSelect) (err error) { _ = "STUB: not implemented"; return nil }

// Consume Order By

//u.Debugf("Order By? %v", m.Cur())

// we have a udf/functional expression column

//u.Debugf("OrderBy after colstart?:   %v  ", m.Cur())

// since we can loop inside switch statement

// This indicates we have come to the End of the columns

// loop on my friend

func (m *Sqlbridge) parseWhereDelete(req *SqlDelete) error { _ = "STUB: not implemented"; return nil }

func (m *Sqlbridge) parseCommandColumns(req *SqlCommand) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//u.Debugf("command col? %v", m.Cur())

//u.Debugf("command after colstart?:   %v  ", m.Cur())

// since we can have multiple columns

func (m *Sqlbridge) parseCreateCols() ([]*DdlColumn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*
	CREATE TABLE articles (
	  ID int(11) NOT NULL AUTO_INCREMENT,
	  Email char(150) NOT NULL DEFAULT '',
	  PRIMARY KEY (ID),
	  CONSTRAINT emails_fk FOREIGN KEY (Email) REFERENCES Emails (Email)
	)
*/

// consume )

// since we can have multiple columns

func (m *Sqlbridge) parseDdlConstraint(col *DdlColumn) error {
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


		index_type:
			USING {BTREE | HASH}
		reference_definition:
		    REFERENCES tbl_name (index_col_name,...)
		      [MATCH FULL | MATCH PARTIAL | MATCH SIMPLE]
		      [ON DELETE reference_option]
		      [ON UPDATE reference_option]

		CONSTRAINT emails_fk FOREIGN KEY (Email) REFERENCES Emails (Email) COMMENT "hello constraint"
	*/return nil
}

// [UNIQUE [KEY] | [PRIMARY] KEY]

// [index_type]
// index_type:
//    USING {BTREE | HASH}

// consume )

// consume )

// [COMMENT 'string']

// since we can have multiple columns

func (m *Sqlbridge) parseDdlColumn(col *DdlColumn) error {
	_ = "STUB: not implemented"

	/*
		http://dev.mysql.com/doc/refman/5.7/en/create-table.html

		create_definition:
		    col_name column_definition

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

		  ID int(11) NOT NULL AUTO_INCREMENT,
		  Email char(150) NOT NULL DEFAULT '',
	*/return nil
}

//u.Debugf("create col after colstart?:   %v  ", m.Cur())

// [NOT NULL | NULL]

// [DEFAULT default_value]

// Consume DEFAULT token

// [AUTO_INCREMENT]

// [UNIQUE [KEY] | [PRIMARY] KEY]

// [COMMENT 'string']

func convertIdentityToValue(n expr.Node) { _ = "STUB: not implemented"; return }

func (m *Sqlbridge) parseLimit(req *SqlSelect) error { _ = "STUB: not implemented"; return nil }

// LIMIT 0, 1000
// consume the comma

// consume "OFFSET"

func (m *Sqlbridge) parseOffset(req *SqlSelect) error { _ = "STUB: not implemented"; return nil }

// Consume "OFFSET"

func (m *Sqlbridge) parseAlias(req *SqlSelect) error { _ = "STUB: not implemented"; return nil }

func (m *Sqlbridge) isEnd() bool { _ = "STUB: not implemented"; return false }

func ParseWith(pg expr.TokenPager) (u.JsonHelper, error) {
	_ = "STUB: not implemented"
	return *new(u.JsonHelper), nil
}

// This is an optional statement

// consume WITH

// {

// name=value pairs

func (m *Sqlbridge) parseShowFromTable(req *SqlShow) error { _ = "STUB: not implemented"; return nil }

// Consume {FROM | IN}

// FROM OR IN are required for this statement

func (m *Sqlbridge) parseShowFromDatabase(req *SqlShow) error {
	_ = "STUB: not implemented"
	return nil
}

// Consume {FROM | IN}

// this is optional

func ParseJsonObject(pg expr.TokenPager, jh u.JsonHelper) error {
	_ = "STUB: not implemented"
	return nil
}

// Consume {

//u.Debug(pg.Cur())

// Consume the right }

func parseJsonKeyValue(pg expr.TokenPager, jh u.JsonHelper) error {
	_ = "STUB: not implemented"
	return nil
}

// {

// [

func ParseJsonArray(pg expr.TokenPager) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Consume [

//u.Debug(pg.Cur())

// {

// [

//u.Debugf("list after: %#v", list)

// Consume the right ]

func ParseKeyValue(pg expr.TokenPager, jh u.JsonHelper) error {
	_ = "STUB: not implemented"
	return nil
}

// whoops, we consumed too much

//u.Debugf("exit keyvalue %v", pg.Cur())

// consume equal

// consume value

// consume comma

// TokenPager is responsible for determining end of
// current tree (column, etc)
type SqlTokenPager struct {
	*expr.LexTokenPager
	lastKw lex.TokenType
}

func NewSqlTokenPager(l *lex.Lexer) *SqlTokenPager { _ = "STUB: not implemented"; return nil }

func (m *SqlTokenPager) IsEnd() bool { _ = "STUB: not implemented"; return false }

func (m *SqlTokenPager) ClauseEnd() bool { _ = "STUB: not implemented"; return false }
