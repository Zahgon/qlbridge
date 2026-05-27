// Package mockcsv implements an in-memory csv data source for testing usage
// implemented by wrapping the mem-b-tree, loading csv data into it.  NOT
// intended for any production usages, test only.
package mockcsv

import (
	"github.com/araddon/qlbridge/datasource/membtree"
	"github.com/araddon/qlbridge/schema"
)

const (
	// SchemaName is "mockcsv"
	SchemaName = "mockcsv"
)

var (
	// Ensure this Csv Data Source implements expected interfaces
	_ schema.Source       = (*Source)(nil)
	_ schema.Alter        = (*Source)(nil)
	_ schema.Conn         = (*Table)(nil)
	_ schema.ConnUpsert   = (*Table)(nil)
	_ schema.ConnDeletion = (*Table)(nil)

	// CsvGlobal mock csv in mem store
	CsvGlobal = New()
	// Schema the mock schema
	sch *schema.Schema
)

// Schema global accessor to the mockcsv schema
func Schema() *schema.Schema { _ = "STUB: not implemented"; return nil }

// LoadTable MockCsv is used for mocking so has a global data source we can load data into
func LoadTable(schemaName, name, csvRaw string) { _ = "STUB: not implemented"; return }

// Source DataSource for testing creates an in memory b-tree per "table".
// Is not thread safe.
type Source struct {
	s             *schema.Schema
	tablenamelist []string
	tables        map[string]*membtree.StaticDataSource
	raw           map[string]string
}

// Table converts the static csv-source into a schema.Conn source
type Table struct {
	*membtree.StaticDataSource
}

// New create csv mock source.
func New() *Source { _ = "STUB: not implemented"; return nil }

// Init no-op meets interface
func (m *Source) Init() {
	_ = "STUB: not implemented"

	// Setup accept schema
	return
}

func (m *Source) Setup(s *schema.Schema) error { _ = "STUB: not implemented"; return nil }

// DropTable Drop table schema
func (m *Source) DropTable(t string) error { _ = "STUB: not implemented"; return nil }

// Open connection to given tablename.
func (m *Source) Open(tableName string) (schema.Conn, error) {
	_ = "STUB: not implemented"
	return *new(schema.Conn), nil
}

// Table get table schema for given table name.  If given table is not currently
// defined, will load, infer schema.
func (m *Source) Table(tableName string) (*schema.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Source) loadTable(tableName string) error { _ = "STUB: not implemented"; return nil }

// The expected format is that the csv data is a single string, with new-lines, etc.

// Now we are going to page through the Csv rows and Put into
// Static Data Source, ie copy into memory btree structure

// We don't know the Key

// Close csv source.
func (m *Source) Close() error {
	_ = "STUB: not implemented"

	// Tables list of tables.
	return nil
}

func (m *Source) Tables() []string { _ = "STUB: not implemented"; return nil }

// CreateTable create a csv table in this source.
func (m *Source) CreateTable(tableName, csvRaw string) { _ = "STUB: not implemented"; return }
