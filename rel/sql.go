// Package rel are the AST Structures and Parsers
// for the SQL, FilterQL, and Expression dialects.
package rel

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/lex"
	"github.com/araddon/qlbridge/value"
)

var (
	// Ensure SqlSelect and cousins etc are SqlStatements
	_ SqlStatement = (*SqlSelect)(nil)
	_ SqlStatement = (*SqlInsert)(nil)
	_ SqlStatement = (*SqlUpsert)(nil)
	_ SqlStatement = (*SqlUpdate)(nil)
	_ SqlStatement = (*SqlDelete)(nil)
	_ SqlStatement = (*SqlShow)(nil)
	_ SqlStatement = (*SqlDescribe)(nil)
	_ SqlStatement = (*SqlCommand)(nil)
	_ SqlStatement = (*SqlInto)(nil)

	// sub-query statements
	_ SqlSourceStatement = (*SqlSource)(nil)

	// Statements with Columns
	_ ColumnsStatement = (*SqlSelect)(nil)

	// A select * columns
	starCols Columns
)

func init() {
	starCols = make(Columns, 1)
	starCols[0] = NewColumnFromToken(lex.Token{T: lex.TokenStar, V: "*"})
}

type (
	// ColumnsStatement is a statement interface for those statements that
	// have columns that need to be added during parse.
	ColumnsStatement interface {
		AddColumn(col Column) error
	}
	// SqlStatement interface, to define the sql statement
	// Select, Insert, Update, Delete, Command, Show, Describe etc
	SqlStatement interface {
		// string representation of Node, AST parseable back to itself
		String() string
		// Write out this statement in a specific dialect
		WriteDialect(w expr.DialectWriter)
		// SQL keyword (select, insert, etc)
		Keyword() lex.TokenType
	}

	// SqlSourceStatement interface, to define the subselect/join-types
	// Join, SubSelect, From
	SqlSourceStatement interface {
		// string representation of Node, AST parseable back to itself
		String() string
		// Write out this statement in a specific dialect
		WriteDialect(w expr.DialectWriter)
		// SQL Keyword for this statement
		Keyword() lex.TokenType
	}
)

type (
	// PreparedStatement Prepared/Aliased SQL statement
	PreparedStatement struct {
		Alias     string
		Statement SqlStatement
	}
	// SqlSelect SQL Select statement
	SqlSelect struct {
		Db        string       // If provided a use "dbname"
		Raw       string       // full original raw statement
		Star      bool         // for select * from ...
		Distinct  bool         // Distinct flag?
		Columns   Columns      // An array (ordered) list of columns
		From      []*SqlSource // From, Join
		Into      *SqlInto     // Into "table"
		Where     *SqlWhere    // Expr Node, or *SqlSelect
		Having    expr.Node    // Filter results
		GroupBy   Columns
		OrderBy   Columns
		Limit     int
		Offset    int
		Alias     string       // Non-Standard sql, alias/name of sql another way of expression Prepared Statement
		With      u.JsonHelper // Non-Standard SQL for properties/config info, similar to Cassandra with, purse json
		proj      *Projection  // Projected fields
		isAgg     bool         // is this an aggregate query?  has group-by, or aggregate selector expressions (count, cardinality etc)
		finalized bool         // have we already finalized, ie formalized left/right aliases
		schemaqry bool         // is this a schema qry?  ie select @@max_packet etc

		// Memoized sql, we assume this is an immuteable struct so if this is populated use it
		pb            *SqlStatementPb
		fingerprintid int64
	}
	// SqlSource is a table name, sub-query, or join as used in
	// SELECT <columns> FROM <SQLSOURCE>
	//  - SELECT .. FROM table_name
	//  - SELECT .. from (select a,b,c from tableb)
	//  - SELECT .. FROM tablex INNER JOIN ...
	SqlSource struct {
		final       bool               // has this been finalized?
		alias       string             // either the short table name or full
		cols        map[string]*Column // Un-aliased columns, ie "x.y" -> "y"
		colIndex    map[string]int     // Key(alias) to index in []driver.Value positions
		joinNodes   []expr.Node        // x.y = q.y AND x.z = q.z  --- []Node{Identity{x},Identity{z}}
		Source      *SqlSelect         // Sql Select Source query, written by Rewrite
		Raw         string             // Raw Partial Query
		Name        string             // From Name (optional, empty if join, subselect)
		Alias       string             // From name aliased
		Schema      string             //  FROM `schema`.`table`
		Op          lex.TokenType      // In, =, ON
		LeftOrRight lex.TokenType      // Left, Right
		JoinType    lex.TokenType      // INNER, OUTER
		JoinExpr    expr.Node          // Join expression       x.y = q.y
		SubQuery    *SqlSelect         // optional, Join/SubSelect statement

		// Plan Hints, move to a dedicated planner
		Seekable bool
		// Memoized sql, we assume this is an immuteable struct so if this is populated use it
		pb *SqlSourcePb
	}
	// SqlWhere WHERE is select stmt, or set of expressions
	// - WHERE x in (select name from q)
	// - WHERE x = y
	// - WHERE x = y AND z = q
	// - WHERE tolower(x) IN (select name from q)
	SqlWhere struct {
		// Either Op + Source exists
		Op     lex.TokenType // (In|=|ON)  for Select Clauses operators
		Source *SqlSelect    // IN (SELECT a,b,c from z)

		// OR expr but not both
		Expr expr.Node // x = y AND q > 5
	}
	// SqlInsert SQL Insert Statement
	SqlInsert struct {
		kw      lex.TokenType    // Insert, Replace
		Table   string           // table name
		Columns Columns          // Column Names
		Rows    [][]*ValueColumn // Values to insert
		Select  *SqlSelect       //
	}
	// SqlUpsert SQL Upsert Statement
	SqlUpsert struct {
		Columns Columns
		Rows    [][]*ValueColumn
		Values  map[string]*ValueColumn
		Where   *SqlWhere
		Table   string
	}
	// SqlUpdate SQL Update Statement
	SqlUpdate struct {
		Values map[string]*ValueColumn
		Where  *SqlWhere
		Table  string
	}
	// SqlDelete SQL Delete Statement
	SqlDelete struct {
		Table string
		Where *SqlWhere
		Limit int
	}
	// SqlShow SQL SHOW Statement
	SqlShow struct {
		Raw        string // full raw statement
		Db         string // Database/Schema name
		Full       bool   // SHOW FULL TABLE FROM
		Scope      string // {FULL, GLOBAL, SESSION}
		ShowType   string // object type, [tables, columns, etc]
		From       string // `table`   or `schema`.`table`
		Identity   string // `table`   or `schema`.`table`
		Create     bool
		CreateWhat string
		Where      expr.Node
		Like       expr.Node
	}
	// SQL Describe statement
	SqlDescribe struct {
		Raw      string    // full original raw statement
		Identity string    // Describe
		Tok      lex.Token // Explain, Describe, Desc
		Stmt     SqlStatement
	}
	// SqlInto   INTO statement   (select a,b,c from y INTO z)
	SqlInto struct {
		Table string
	}
	// SqlCommand is admin command such as "SET", "USE"
	SqlCommand struct {
		kw       lex.TokenType  // SET or USE
		Columns  CommandColumns // can have multiple columns in command
		Identity string         //
		Value    expr.Node      //
	}
	// SqlCreate SQL CREATE statement
	SqlCreate struct {
		Raw         string       // full original raw statement
		Identity    string       // identity of table, view, etc
		Tok         lex.Token    // CREATE [TABLE,VIEW,CONTINUOUSVIEW,TRIGGER] etc
		OrReplace   bool         // OR REPLACE
		IfNotExists bool         // IF NOT EXISTS
		Cols        []*DdlColumn // columns
		Engine      map[string]interface{}
		With        u.JsonHelper
		Select      *SqlSelect
	}
	// SqlDrop SQL DROP statement
	SqlDrop struct {
		Raw      string    // full original raw statement
		Identity string    // identity of table, view, etc
		Temp     bool      // Temp?
		Tok      lex.Token // DROP [TEMP] [TABLE,VIEW,CONTINUOUSVIEW,TRIGGER] etc
		With     u.JsonHelper
	}
	// SqlAlter SQL ALTER statement
	SqlAlter struct {
		Raw      string       // full original raw statement
		Identity string       // identity to alter
		Tok      lex.Token    // ALTER [TABLE,VIEW,CONTINUOUSVIEW,TRIGGER] etc
		Cols     []*DdlColumn // columns
	}
	// Columns List of Columns in SELECT [columns]
	Columns []*Column
	// Column represents the Column as expressed in a [SELECT]
	// expression
	Column struct {
		sourceQuoteByte byte      // quote mark?   [ or ` etc
		asQuoteByte     byte      // quote mark   [ or `
		originalAs      string    // original as string
		left            string    // users.col_name   = "users"
		right           string    // users.first_name = "first_name"
		isLiteral       bool      // is this a literal column?
		ParentIndex     int       // slice idx position in parent query cols
		Index           int       // slice idx position in original query cols
		SourceIndex     int       // slice idx position in source []driver.Value
		SourceField     string    // field name of underlying field
		SourceOriginal  string    // field name of underlying field without the "left.right" parse
		As              string    // As field, auto-populate the Field Name if exists
		Comment         string    // optional in-line comments
		Order           string    // (ASC | DESC)
		Star            bool      // *
		Agg             bool      // aggregate function column?   count(*), avg(x) etc
		Expr            expr.Node // Expression, optional, often Identity.Node
		Guard           expr.Node // column If guard, non-standard sql column guard
	}
	// ValueColumn List of Value columns in INSERT into TABLE (colnames) VALUES (valuecolumns)
	ValueColumn struct {
		Value value.Value
		Expr  expr.Node
	}
	// DdlColumn represents the Data Definition Column
	DdlColumn struct {
		Kw            lex.TokenType // initial keyword (identity for normal, constraint, primary)
		Null          bool          // Do we support NULL?
		AutoIncrement bool          // auto increment
		IndexType     string        // index_type
		IndexCols     []string      // index_col_name
		RefTable      string        // refererence table
		RefCols       []string      // ref cols
		Default       expr.Node     // Default value
		DataType      string        // data type
		DataTypeSize  int           // Data Type Size:    varchar(2000)
		DataTypeArgs  []expr.Node   // data type args
		Key           lex.TokenType // UNIQUE | PRIMARY
		Name          string        // name
		Comment       string        // optional in-line comments
		Expr          expr.Node     // Expression, optional, often Identity.Node but could be composite key
	}
	// ResultColumns List of ResultColumns used to describe projection response columns
	ResultColumns []*ResultColumn
	// Result Column used in projection
	ResultColumn struct {
		Final  bool            // Is this part of final projection (ie, response)
		Name   string          // Original path/name for query field
		ColPos int             // Ordinal position in sql (or partial sql) statement
		Col    *Column         // the original sql column
		Star   bool            // Was this a select * ??
		As     string          // aliased
		Type   value.ValueType // Data Type
	}
	// Projection describes the results to expect from sql statement
	// ie the ResultColumns for a result-set
	Projection struct {
		Distinct bool
		Final    bool // Is this a Final Projection? or intermiediate?
		colNames map[string]struct{}
		Columns  ResultColumns
		// Memoized pb, we assume this is an immuteable struct so if this is populated use it
		pb *ProjectionPb
	}
	// CommandColumns SQL commands such as:
	//     set autocommit
	//     SET @@local.sort_buffer_size=10000;
	//     USE myschema;
	CommandColumns []*CommandColumn
	// CommandColumn is single column such as "autocommit"
	CommandColumn struct {
		Expr expr.Node // column expression
		Name string    // Original path/name for command field
	}
)

func NewSqlDialect() expr.DialectWriter { _ = "STUB: not implemented"; return *new(expr.DialectWriter) }

func NewProjection() *Projection { _ = "STUB: not implemented"; return nil }

func NewResultColumn(as string, ordinal int, col *Column, valtype value.ValueType) *ResultColumn {
	_ = "STUB: not implemented"
	return nil
}

func NewSqlSelect() *SqlSelect { _ = "STUB: not implemented"; return nil }

func NewSqlInsert() *SqlInsert { _ = "STUB: not implemented"; return nil }

func NewSqlUpdate() *SqlUpdate { _ = "STUB: not implemented"; return nil }

func NewSqlUpsert() *SqlUpsert { _ = "STUB: not implemented"; return nil }

func NewSqlDelete() *SqlDelete { _ = "STUB: not implemented"; return nil }

func NewPreparedStatement() *PreparedStatement { _ = "STUB: not implemented"; return nil }

func NewSqlCreate() *SqlCreate { _ = "STUB: not implemented"; return nil }

func NewSqlDrop() *SqlDrop { _ = "STUB: not implemented"; return nil }

func NewSqlInto(table string) *SqlInto { _ = "STUB: not implemented"; return nil }

func NewSqlSource(table string) *SqlSource { _ = "STUB: not implemented"; return nil }

func NewSqlWhere(where expr.Node) *SqlWhere { _ = "STUB: not implemented"; return nil }

func NewColumnFromToken(tok lex.Token) *Column { _ = "STUB: not implemented"; return nil }

//v = expr.IdentityMaybeQuote(tok.Quote, v)

func NewColumnValue(tok lex.Token) *Column { _ = "STUB: not implemented"; return nil }

func NewColumn(col string) *Column { _ = "STUB: not implemented"; return nil }

// The source column name
func (m *ResultColumn) SourceName() string { _ = "STUB: not implemented"; return "" }

func (m *ResultColumn) Equal(s *ResultColumn) bool { _ = "STUB: not implemented"; return false }

//u.Warnf("Not Equal?   %T  vs %T", m.Col, s.Col)
//u.Warnf("t!=t:   \n\t%#v\n\t%#v", m.Col, s.Col)

func resultColumnFromPb(pb *ResultColumnPb) *ResultColumn { _ = "STUB: not implemented"; return nil }

func resultColumnToPb(m *ResultColumn) *ResultColumnPb { _ = "STUB: not implemented"; return nil }

func (m *Projection) AddColumnShort(colName string, vt value.ValueType) {
	_ = "STUB: not implemented"
	//colName = strings.ToLower(colName)
	// if _, exists := m.colNames[colName]; exists {
	// 	return
	// }
	//u.Infof("adding column %s to %v", colName, m.colNames)
	//m.colNames[colName] = struct{}{}
	return
}

func (m *Projection) AddColumn(col *Column, vt value.ValueType) {
	_ = "STUB: not implemented"
	// colName := strings.ToLower(col.As)
	//
	//	if _, exists := m.colNames[colName]; exists {
	//		return
	//	}
	//
	// m.colNames[colName] = struct{}{}
	return
}

func (m *Projection) Equal(s *Projection) bool { _ = "STUB: not implemented"; return false }

//u.Warnf("Not Equal?   %T  vs %T", c, s.Columns[i])
//u.Warnf("t!=t:   \n\t%#v \n\t!= %#v", c, s.Columns[i])

func (m *Projection) FromPB(pb *ProjectionPb) *Projection { _ = "STUB: not implemented"; return nil }

func (m *Projection) ToPB() *ProjectionPb { _ = "STUB: not implemented"; return nil }

func ProjectionFromPb(pb *ProjectionPb) *Projection { _ = "STUB: not implemented"; return nil }

func projectionToPb(m *Projection) *ProjectionPb { _ = "STUB: not implemented"; return nil }

func (m *Columns) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }

func (m *Columns) String() string { _ = "STUB: not implemented"; return "" }

func (m *Columns) FieldNames() []string { _ = "STUB: not implemented"; return nil }

func (m *Columns) UnAliasedFieldNames() []string { _ = "STUB: not implemented"; return nil }

func (m *Columns) AliasedFieldNames() []string { _ = "STUB: not implemented"; return nil }

func (m *Columns) ByName(name string) (*Column, bool) {
	_ = "STUB: not implemented"
	return nil,

		// u.Debugf("col.SourceField='%s' key()='%s' As='%s' ", col.SourceField, col.Key(), col.As)
		false
}

func (m *Columns) ByAs(as string) (*Column, bool) { _ = "STUB: not implemented"; return nil, false }

func (m Columns) Equal(cols Columns) bool { _ = "STUB: not implemented"; return false }

func (m *Column) Key() string { _ = "STUB: not implemented"; return "" }

func (m *Column) String() string { _ = "STUB: not implemented"; return "" }

func (m *Column) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }

// Is this a select count(*) column
func (m *Column) CountStar() bool { _ = "STUB: not implemented"; return false }

func (m *Column) InFinalProjection() bool { _ = "STUB: not implemented"; return false }

func (m *Column) IsLiteral() bool { _ = "STUB: not implemented"; return false }

// count(*)
// now()
// tolower(field_name)

func (m *Column) IsLiteralOrFunc() bool { _ = "STUB: not implemented"; return false }

// count(*)
// now()
// tolower(field_name)

// What about NULL?

func (m *Column) Asc() bool { _ = "STUB: not implemented"; return false }

func (m *Column) Equal(c *Column) bool { _ = "STUB: not implemented"; return false }

// CopyRewrite Create a new copy of this column for rewrite purposes removing alias
func (m *Column) CopyRewrite(alias string) *Column { _ = "STUB: not implemented"; return nil }

//u.Warnf("in rewrite:  Alias:'%s'  '%s'.'%s'  sourcefield:'%v'", alias, left, right, m.SourceField)

// Copy - deep copy, shared nothing
func (m *Column) Copy() *Column { _ = "STUB: not implemented"; return nil }

func (m *Column) ToPB() *ColumnPb { _ = "STUB: not implemented"; return nil }

func columnFromPb(c *ColumnPb) *Column { _ = "STUB: not implemented"; return nil }

// Return left, right values if is of form   `table.column` and
// also return true/false for if it even has left/right
func (m *Column) LeftRight() (string, string, bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func (m *PreparedStatement) Keyword() lex.TokenType {
	_ = "STUB: not implemented"
	return *new(lex.TokenType)
}
func (m *PreparedStatement) String() string { _ = "STUB: not implemented"; return "" }

func (m *PreparedStatement) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }

func (m *SqlSelect) Keyword() lex.TokenType             { _ = "STUB: not implemented"; return *new(lex.TokenType) }
func (m *SqlSelect) SystemQry() bool                    { _ = "STUB: not implemented"; return false }
func (m *SqlSelect) SetSystemQry()                      { _ = "STUB: not implemented"; return }
func (m *SqlSelect) IsLiteral() bool                    { _ = "STUB: not implemented"; return false }
func (m *SqlSelect) FromPB(spb *SqlSelectPb) *SqlSelect { _ = "STUB: not implemented"; return nil }

func (m *SqlSelect) ToPbStatement() *SqlStatementPb { _ = "STUB: not implemented"; return nil }

func (m *SqlSelect) ToPB() *SqlSelectPb { _ = "STUB: not implemented"; return nil }

func (m *SqlSelect) Copy() *SqlSelect { _ = "STUB: not implemented"; return nil }

// SqlSelectToPb Given a select statement lets convert it into a PB statement
func SqlSelectToPb(m *SqlSelect) *SqlSelectPb { _ = "STUB: not implemented"; return nil }

func sqlSelectToPbDepth(m *SqlSelect, depth int) *SqlSelectPb {
	_ = "STUB: not implemented"
	//u.Debugf("SqlSelectToPb %d? %p", depth, m)
	return nil
}

func (m *SqlSelect) Equal(ss SqlStatement) bool { _ = "STUB: not implemented"; return false }

// SqlSelectFromPb take a protobuf select struct and conver to SqlSelect
func SqlSelectFromPb(pb *SqlSelectPb) *SqlSelect { _ = "STUB: not implemented"; return nil }

func (m *SqlSelect) IsAggQuery() bool { _ = "STUB: not implemented"; return false }

func (m *SqlSelect) String() string { _ = "STUB: not implemented"; return "" }

func (m *SqlSelect) writeDialectDepth(depth int, w expr.DialectWriter) {
	_ = "STUB: not implemented"
	return
}

func (m *SqlSelect) FingerPrintID() int64 { _ = "STUB: not implemented"; return 0 }

func (m *SqlSelect) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }

// Finalize this Query plan by preparing sub-sources
//
//	ie we need to rewrite some things into sub-statements
//	- we need to share the join expression across sources
func (m *SqlSelect) Finalize() error { _ = "STUB: not implemented"; return nil }

func (m *SqlSelect) UnAliasedColumns() map[string]*Column { _ = "STUB: not implemented"; return nil }

func (m *SqlSelect) AliasedColumns() map[string]*Column { _ = "STUB: not implemented"; return nil }

//u.Debugf("aliasing: key():%-15q  As:%-15q   %-15q", col.Key(), col.As, col.String())

func (m *SqlSelect) ColIndexes() map[string]int { _ = "STUB: not implemented"; return nil }

//u.Debugf("aliasing: key():%-15q  As:%-15q   %-15q", col.Key(), col.As, col.String())

func (m *SqlSelect) AddColumn(colArg Column) error { _ = "STUB: not implemented"; return nil }

// Is this a select count(*) FROM ...   query?
func (m *SqlSelect) CountStar() bool { _ = "STUB: not implemented"; return false }

// Rewrite take current SqlSelect statement and re-write it
func (m *SqlSelect) Rewrite() { _ = "STUB: not implemented"; return }

// RewriteAsRawSelect We are removing Column Aliases "user_id as uid"
// as well as functions - used when we are going to defer projection, aggs
func (m *SqlSelect) RewriteAsRawSelect() { _ = "STUB: not implemented"; return }

func (m *SqlSource) IsLiteral() bool        { _ = "STUB: not implemented"; return false }
func (m *SqlSource) Keyword() lex.TokenType { _ = "STUB: not implemented"; return *new(lex.TokenType) }
func (m *SqlSource) SourceName() string     { _ = "STUB: not implemented"; return "" }

func (m *SqlSource) String() string { _ = "STUB: not implemented"; return "" }

func (m *SqlSource) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }

func (m *SqlSource) writeDialectDepth(depth int, w expr.DialectWriter) {
	_ = "STUB: not implemented"
	return
}

//   Jointype                Op
//  INNER JOIN orders AS o 	ON

// inner/outer

func (m *SqlSource) BuildColIndex(colNames []string) error { _ = "STUB: not implemented"; return nil }

// how many columns were added due to *

//u.Debugf("col.Key():%v  sourceField:%v  colName:%v", col.Key(), col.SourceField, colName)
//&&
//u.Debugf("build col:  idx=%d  key=%-15q as=%-15q col=%-15s sourcidx:%d", len(m.colIndex), col.Key(), col.As, col.String(), colIdx)

// Rewrite this Source to act as a stand-alone query to backend
// @parentStmt = the parent statement that this a partial source to
func (m *SqlSource) Rewrite(parentStmt *SqlSelect) *SqlSelect {
	_ = "STUB: not implemented"
	return nil
}

func (m *SqlSource) findFromAliases() (string, string) { _ = "STUB: not implemented"; return "", "" }

// Get a list of Un-Aliased Columns, ie columns with column
//
//	names that have NOT yet been aliased
func (m *SqlSource) UnAliasedColumns() map[string]*Column {
	_ = "STUB: not implemented"
	// u.Warnf("un-aliased %d", len(m.Source.Columns))
	return nil
}

//u.Debugf("aliasing: l:%q r:%q hasLeft?%v", left, right, hasLeft)

// Get a list of Column names to position
func (m *SqlSource) ColumnPositions() map[string]int { _ = "STUB: not implemented"; return nil }

//u.Debugf("aliasing: l:%v r:%v ok?%v", left, right, ok)

// We need to be able to rewrite statements to convert a stmt such as:
//
//	FROM users AS u
//	    INNER JOIN orders AS o
//	    ON u.user_id = o.user_id
//
// So that we can evaluate the Join Key on left/right
// in this case, it is simple, just
//
//	=>   user_id
//
// or this one:
//
//			FROM users AS u
//				INNER JOIN orders AS o
//				ON LOWER(u.email) = LOWER(o.email)
//
//	   =>  LOWER(user_id)
func (m *SqlSource) JoinNodes() []expr.Node { _ = "STUB: not implemented"; return nil }

func (m *SqlSource) Finalize() error { _ = "STUB: not implemented"; return nil }

//u.Warnf("finalize sqlsource: %v", len(m.Columns))

func (m *SqlSource) FromPB(n *SqlSourcePb) *SqlSource { _ = "STUB: not implemented"; return nil }

func (m *SqlSource) ToPB() *SqlSourcePb { _ = "STUB: not implemented"; return nil }

func (m *SqlSource) Equal(s *SqlSource) bool { _ = "STUB: not implemented"; return false }

func sqlSourceToPb(m *SqlSource) *SqlSourcePb { _ = "STUB: not implemented"; return nil }

// We get into recursive hell if we don't bail
// but need to go stich in source?

//u.Warnf("about to descend? %p", m.Source)

func SqlSourceFromPb(pb *SqlSourcePb) *SqlSource { _ = "STUB: not implemented"; return nil }

//u.Debugf("no source for SqlSource? %+v", pb)

func (m *SqlWhere) Keyword() lex.TokenType { _ = "STUB: not implemented"; return *new(lex.TokenType) }
func (m *SqlWhere) writeDialectDepth(depth int, w expr.DialectWriter) {
	_ = "STUB: not implemented"
	return
}

// Op = subselect or in etc
//  SELECT ... WHERE IN (SELECT ...)

func (m *SqlWhere) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }
func (m *SqlWhere) String() string                    { _ = "STUB: not implemented"; return "" }

func (m *SqlWhere) Equal(s *SqlWhere) bool { _ = "STUB: not implemented"; return false }

func SqlWhereToPb(m *SqlWhere) *SqlWherePb { _ = "STUB: not implemented"; return nil }

func SqlWhereFromPb(pb *SqlWherePb) *SqlWhere { _ = "STUB: not implemented"; return nil }

func (m *SqlInto) Keyword() lex.TokenType            { _ = "STUB: not implemented"; return *new(lex.TokenType) }
func (m *SqlInto) String() string                    { _ = "STUB: not implemented"; return "" }
func (m *SqlInto) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }
func (m *SqlInto) Equal(s *SqlInto) bool             { _ = "STUB: not implemented"; return false }

func (m *SqlInsert) Keyword() lex.TokenType            { _ = "STUB: not implemented"; return *new(lex.TokenType) }
func (m *SqlInsert) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }

// Value is not nil

func (m *SqlInsert) String() string { _ = "STUB: not implemented"; return "" }

// RewriteAsPrepareable rewite the insert as a ? substituteable query
//
//	INSERT INTO user (name) VALUES ("wonder-woman") ->
//	   INSERT INTO user (name) VALUES (?)
func (m *SqlInsert) RewriteAsPrepareable(maxRows int, mark byte) string {
	_ = "STUB: not implemented"
	return ""
}

func (m *SqlInsert) ColumnNames() []string { _ = "STUB: not implemented"; return nil }

func (m *SqlUpsert) Keyword() lex.TokenType            { _ = "STUB: not implemented"; return *new(lex.TokenType) }
func (m *SqlUpsert) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }
func (m *SqlUpsert) String() string                    { _ = "STUB: not implemented"; return "" }
func (m *SqlUpsert) SqlSelect() *SqlSelect             { _ = "STUB: not implemented"; return nil }

func (m *SqlUpdate) Keyword() lex.TokenType            { _ = "STUB: not implemented"; return *new(lex.TokenType) }
func (m *SqlUpdate) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }

func (m *SqlUpdate) String() string { _ = "STUB: not implemented"; return "" }

func (m *SqlUpdate) SqlSelect() *SqlSelect { _ = "STUB: not implemented"; return nil }

func sqlSelectFromWhere(from string, where *SqlWhere) *SqlSelect {
	_ = "STUB: not implemented"
	return nil
}

func (m *SqlDelete) Keyword() lex.TokenType            { _ = "STUB: not implemented"; return *new(lex.TokenType) }
func (m *SqlDelete) String() string                    { _ = "STUB: not implemented"; return "" }
func (m *SqlDelete) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }

func (m *SqlDelete) SqlSelect() *SqlSelect { _ = "STUB: not implemented"; return nil }

func (m *SqlDescribe) Keyword() lex.TokenType {
	_ = "STUB: not implemented"
	return *new(lex.TokenType)
}
func (m *SqlDescribe) String() string                    { _ = "STUB: not implemented"; return "" }
func (m *SqlDescribe) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }

func (m *SqlShow) Keyword() lex.TokenType            { _ = "STUB: not implemented"; return *new(lex.TokenType) }
func (m *SqlShow) String() string                    { _ = "STUB: not implemented"; return "" }
func (m *SqlShow) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }

func (m *CommandColumn) FingerPrint(r rune) string { _ = "STUB: not implemented"; return "" }
func (m *CommandColumn) String() string            { _ = "STUB: not implemented"; return "" }

func (m *CommandColumn) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }
func (m *CommandColumn) Key() string                       { _ = "STUB: not implemented"; return "" }

func (m *CommandColumns) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }
func (m *CommandColumns) String() string                    { _ = "STUB: not implemented"; return "" }

func (m *SqlCommand) Keyword() lex.TokenType            { _ = "STUB: not implemented"; return *new(lex.TokenType) }
func (m *SqlCommand) FingerPrint(r rune) string         { _ = "STUB: not implemented"; return "" }
func (m *SqlCommand) String() string                    { _ = "STUB: not implemented"; return "" }
func (m *SqlCommand) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }

func (m *SqlCreate) Keyword() lex.TokenType            { _ = "STUB: not implemented"; return *new(lex.TokenType) }
func (m *SqlCreate) FingerPrint(r rune) string         { _ = "STUB: not implemented"; return "" }
func (m *SqlCreate) String() string                    { _ = "STUB: not implemented"; return "" }
func (m *SqlCreate) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }

func (m *SqlDrop) Keyword() lex.TokenType            { _ = "STUB: not implemented"; return *new(lex.TokenType) }
func (m *SqlDrop) FingerPrint(r rune) string         { _ = "STUB: not implemented"; return "" }
func (m *SqlDrop) String() string                    { _ = "STUB: not implemented"; return "" }
func (m *SqlDrop) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }

func (m *SqlAlter) Keyword() lex.TokenType    { _ = "STUB: not implemented"; return *new(lex.TokenType) }
func (m *SqlAlter) FingerPrint(r rune) string { _ = "STUB: not implemented"; return "" }
func (m *SqlAlter) String() string            { _ = "STUB: not implemented"; return "" }
func (m *SqlAlter) WriteDialect(w expr.DialectWriter) {
	_ = "STUB: not implemented"

	// Node serialization helpers
	return
}

func tokenFromInt(iv int32) lex.Token { _ = "STUB: not implemented"; return *new(lex.Token) }

// SqlFromPb Create a sql statement from pb
func SqlFromPb(pb []byte) (SqlStatement, error) {
	_ = "STUB: not implemented"
	return *new(SqlStatement), nil
}

func statementFromPb(s *SqlStatementPb) SqlStatement {
	_ = "STUB: not implemented"
	return *new(SqlStatement)
}

func MapIntFromPb(kv []KvInt) map[string]int { _ = "STUB: not implemented"; return nil }

func ColumnsFromPb(c []*ColumnPb) Columns { _ = "STUB: not implemented"; return *new(Columns) }

func ColumnsToPb(c Columns) []*ColumnPb { _ = "STUB: not implemented"; return nil }

func optionalByte(b []byte) byte { _ = "STUB: not implemented"; return 0 }

// EqualWith compare two with helpers for equality.
func EqualWith(l, r u.JsonHelper) bool { _ = "STUB: not implemented"; return false }

// HelperString Convert a Helper into key/value string
func HelperString(w expr.DialectWriter, jh u.JsonHelper) {
	_ = "STUB: not implemented"

	// isJson := false
	//
	//	for k, v := range jh {
	//		switch lvt := lv.(type) {
	//		case int, int64, int32, string, bool, float64:
	//			//
	//		case []string, []int, []int32, []int64, []float64:
	//			//
	//		case u.JsonHelper, map[string]interface{}:
	//			isJson = true
	//			break
	//		default:
	//			u.Warnf("unhandled type comparison: %T", lv)
	//		}
	//	}
	return
}
