package datasource

import (
	"bytes"
	"database/sql/driver"

	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/plan"
	"github.com/araddon/qlbridge/schema"
	"github.com/araddon/qlbridge/value"
)

const (
	// SchemaDbSourceType is schemadb source type name
	SchemaDbSourceType = "schemadb"
)

var (
	// Ensure our schemadb implements schema.Source etc interfaces.
	_ schema.Source      = (*SchemaDb)(nil)
	_ schema.Alter       = (*SchemaDb)(nil)
	_ schema.Conn        = (*SchemaSource)(nil)
	_ schema.ConnColumns = (*SchemaSource)(nil)
	_ schema.ConnScanner = (*SchemaSource)(nil)

	// normal tables
	defaultSchemaTables = []string{"tables", "databases", "columns", "global_variables", "session_variables",
		"functions", "procedures", "engines", "status", "indexes"}
	// DialectWriterCols list of columns for dialectwriter.
	DialectWriterCols = []string{"mysql"}
	// DialectWriters list of differnt writers.
	DialectWriters = []schema.DialectWriter{&mysqlWriter{}}

	// privates
	_        = u.EMPTY
	registry *schema.Registry
)

func init() {
	schema.CreateDefaultRegistry(schema.NewApplyer(SchemaDBStoreProvider))
	registry = schema.DefaultRegistry()
}

type (
	// SchemaDb Static Schema Source, implements qlbridge DataSource to allow in-memory
	// native go data to have a Schema and implement and be operated on by Sql Operations.
	SchemaDb struct {
		exit     <-chan bool
		s        *schema.Schema
		tbls     []string
		tableMap map[string]*schema.Table
	}
	// SchemaSource type for the schemadb connection (thread-safe).
	SchemaSource struct {
		db      *SchemaDb
		tbl     *schema.Table
		ctx     *plan.Context
		session bool
		cursor  int
		rows    [][]driver.Value
	}
)

// SchemaDBStoreProvider create source for schemadb
func SchemaDBStoreProvider(s *schema.Schema) schema.Source {
	_ = "STUB: not implemented"
	return *new(schema.Source)
}

// NewSchemaDb create new db for storing schema.
func NewSchemaDb(s *schema.Schema) *SchemaDb { _ = "STUB: not implemented"; return nil }

// Init initialize
func (m *SchemaDb) Init() { _ = "STUB: not implemented"; return }

// Setup the schemadb
func (m *SchemaDb) Setup(*schema.Schema) error {
	_ = "STUB: not implemented"

	// Close down everything.
	return nil
}

func (m *SchemaDb) Close() error {
	_ = "STUB: not implemented"

	// Tables list of table names.
	return nil
}

func (m *SchemaDb) Tables() []string {
	_ = "STUB: not implemented"

	// Table get schema Table
	return nil
}

func (m *SchemaDb) Table(table string) (*schema.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Open Create a SchemaSource specific to schema object (table, database)
func (m *SchemaDb) Open(schemaObjectName string) (schema.Conn, error) {
	_ = "STUB: not implemented"
	return *new(schema.Conn), nil
}

// SetContext set the plan context
func (m *SchemaSource) SetContext(ctx *plan.Context) { _ = "STUB: not implemented"; return }

func (m *SchemaSource) Close() error                  { _ = "STUB: not implemented"; return nil }
func (m *SchemaSource) SetRows(rows [][]driver.Value) { _ = "STUB: not implemented"; return }
func (m *SchemaSource) Columns() []string             { _ = "STUB: not implemented"; return nil }
func (m *SchemaSource) Next() schema.Message {
	_ = "STUB: not implemented"
	return *new(schema.Message)
}

//u.Debugf("msg: %#v", msg)

func (m *SchemaSource) Get(key driver.Value) (schema.Message, error) {
	_ = "STUB: not implemented"
	return *new(schema.Message), nil
}

func (m *SchemaDb) DropTable(t string) error { _ = "STUB: not implemented"; return nil }

func (m *SchemaDb) inspect(table string) { _ = "STUB: not implemented"; return }

func (m *SchemaDb) tableForTable(table string) (*schema.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//u.Debugf("s:%p infoschema:%p creating schema table for %q", m.s, m.s.InfoSchema, table)

// I really don't like where/how this gets called
// needs to be in schema somewhere?

//u.Infof("found srcTable %v fields?%v", srcTbl.Columns(), len(srcTbl.Fields))

func (m *SchemaDb) tableForProcedures(table string) (*schema.Table, error) {
	_ = "STUB: not implemented"

	//table := "procedures"  // procedures, functions
	return nil, nil
}

// u.Debugf("s:%p creating schema table for %q", m.s, table)

//  SELECT Db, Name, Type, Definer, Modified, Created, Security_type, Comment,
//     character_set_client, `collation_connection`, `Database Collation` from `context`.`procedures`;")

func (m *SchemaDb) tableForEngines() (*schema.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//u.Debugf("s:%p infoschema:%p creating schema table for %q", m.s, m.is, table)

//u.Infof("found existing table %q", table)

// Engine, Support, Comment, Transactions, XA, Savepoints

func (m *SchemaDb) tableForVariables(table string) (*schema.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *SchemaDb) tableForTables() (*schema.Table, error) {
	_ = "STUB: not implemented"

	//u.Debugf("schema:%p  table create infoschema:%p  %v", m.s, m.s.InfoSchema, m.s.Tables())
	return nil, nil
}

// I really don't like where this is, needs to be in schema somewhere

//u.Debugf("%T  %s", writer, rows[i][len(rows[i])-1])

//u.Debugf("set rows: %v for tables: %v", rows, m.s.Tables())

func (m *SchemaDb) tableForIndexes() (*schema.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//u.Debugf("s:%p infoschema:%p creating schema table for %q", m.s, m.is, table)

/*
	mysql> show keys from `user` from `mysql`;
	+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
	| Table | Non_unique | Key_name | Seq_in_index | Column_name | Collation | Cardinality | Sub_part | Packed | Null | Index_type | Comment | Index_comment |
	+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
	| user  |          0 | PRIMARY  |            1 | Host        | A         |        NULL |     NULL | NULL   |      | BTREE      |         |               |
	| user  |          0 | PRIMARY  |            2 | User        | A         |           3 |     NULL | NULL   |      | BTREE      |         |               |
	+-------+------------+----------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
*/

//t.SetRows(rows)

func (m *SchemaDb) tableForDatabases() (*schema.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type mysqlWriter struct {
}

func (m *mysqlWriter) Dialect() string { _ = "STUB: not implemented"; return "" }

func (m *mysqlWriter) FieldType(t value.ValueType) string { _ = "STUB: not implemented"; return "" }

// Table Implement Dialect Specific Writers
// ie, mysql, postgres, cassandra all have different dialects
// so the Create statements are quite different

// Table output a CREATE TABLE statement using mysql dialect.
func (m *mysqlWriter) Table(tbl *schema.Table) string { _ = "STUB: not implemented"; return "" }

//u.Infof("%s tbl=%p fields? %#v fields?%v", tbl.Name, tbl, tbl.FieldMap, len(tbl.Fields))

//tblStr := fmt.Sprintf("CREATE TABLE `%s` (\n\n);", tbl.Name, strings.Join(cols, ","))
//return tblStr, nil

func mysqlWriteField(w *bytes.Buffer, fld *schema.Field) { _ = "STUB: not implemented"; return }

func MysqlValueString(t value.ValueType) string { _ = "STUB: not implemented"; return "" }
