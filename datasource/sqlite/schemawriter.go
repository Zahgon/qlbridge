package sqlite

import (
	"bytes"

	"github.com/araddon/qlbridge/datasource"
	"github.com/araddon/qlbridge/schema"
	"github.com/araddon/qlbridge/value"
)

var (

// normal tables
// defaultSchemaTables = []string{"tables", "databases", "columns", "global_variables", "session_variables","functions", "procedures", "engines", "status", "indexes"}
)

func init() {
	datasource.DialectWriterCols = append(datasource.DialectWriterCols, "sqlite")
	datasource.DialectWriters = append(datasource.DialectWriters, &sqliteWriter{})
}

type sqliteWriter struct {
}

func (m *sqliteWriter) Dialect() string { _ = "STUB: not implemented"; return "" }

func (m *sqliteWriter) FieldType(t value.ValueType) string { _ = "STUB: not implemented"; return "" }

// Table Implement Dialect Specific Writers
// ie, mysql, postgres, cassandra all have different dialects
// so the Create statements are quite different

// Table output a CREATE TABLE statement using mysql dialect.
func (m *sqliteWriter) Table(tbl *schema.Table) string { _ = "STUB: not implemented"; return "" }

// TableToString Table output a CREATE TABLE statement using mysql dialect.
func TableToString(tbl *schema.Table) string { _ = "STUB: not implemented"; return "" }

//u.Infof("%s tbl=%p fields? %#v fields?%v", tbl.Name, tbl, tbl.FieldMap, len(tbl.Fields))

//tblStr := fmt.Sprintf("CREATE TABLE `%s` (\n\n);", tbl.Name, strings.Join(cols, ","))
//return tblStr, nil

// WriteField write a schema.Field as string output for sqlite create statement
//
// https://www.sqlite.org/datatype3.html
func WriteField(w *bytes.Buffer, fld *schema.Field) { _ = "STUB: not implemented"; return }

/*
	NULL. The value is a NULL value.
	INTEGER. The value is a signed integer, stored in 1, 2, 3, 4, 6, or 8 bytes depending on the magnitude of the value.
	REAL. The value is a floating point value, stored as an 8-byte IEEE floating point number.
	TEXT. The value is a text string, stored using the database encoding (UTF-8, UTF-16BE or UTF-16LE).
	BLOB. The value is a blob of data, stored exactly as it was input.
*/
//deflen := fld.Length

// TypeFromString given a string, return data type
func TypeFromString(t string) value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}

// This isn't necessarily true, as integer could be bool

// ValueString convert a value.ValueType into a sqlite type descriptor
func ValueString(t value.ValueType) string { _ = "STUB: not implemented"; return "" }
