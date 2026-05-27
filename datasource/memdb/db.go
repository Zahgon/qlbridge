// Memdb package implements a Qlbridge Datasource in-memory implemenation
// using the hashicorp go-memdb (immuteable radix tree's).
// Qlbridge Exec allows key-value datasources to have full SQL functionality.
package memdb

import (
	"database/sql/driver"

	"github.com/hashicorp/go-memdb"
	"golang.org/x/net/context"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/schema"
)

const (
	sourceType = "memdb"
)

var (
	// Ensure our MemDB implements schema.Source
	_ schema.Source = (*MemDb)(nil)

	// Ensure our dbConn implements variety of Connection interfaces.
	_ schema.Conn         = (*dbConn)(nil)
	_ schema.ConnColumns  = (*dbConn)(nil)
	_ schema.ConnScanner  = (*dbConn)(nil)
	_ schema.ConnUpsert   = (*dbConn)(nil)
	_ schema.ConnDeletion = (*dbConn)(nil)
	_ schema.ConnSeeker   = (*dbConn)(nil)
)

// MemDb implements qlbridge `Source` to allow in-memory native go data
// to have a Schema and implement and be operated on by Sql Statements.
type MemDb struct {
	exit           chan bool
	*schema.Schema                 // schema
	tbl            *schema.Table   // schema table
	indexes        []*schema.Index // index descriptions
	primaryIndex   string
	db             *memdb.MemDB
	max            int
}
type dbConn struct {
	md     *MemDb
	db     *memdb.MemDB
	txn    *memdb.Txn
	result memdb.ResultIterator
}

// NewMemDbData creates a MemDb with given indexes, columns, and values
func NewMemDbData(name string, data [][]driver.Value, cols []string) (*MemDb, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Insert initial values

// we are going to look at ~10 rows to create schema for it

// NewMemDb creates a MemDb with given indexes, columns
func NewMemDb(name string, cols []string) (*MemDb, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewMemDbForSchema creates a MemDb with given indexes, columns
func NewMemDbForSchema(name string, cols []string) (*MemDb, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Init initilize this db
func (m *MemDb) Init() {
	_ = "STUB: not implemented"

	// Setup this db with parent schema.
	return
}

func (m *MemDb) Setup(*schema.Schema) error {
	_ = "STUB: not implemented"

	// Open a Conn for this source @table name
	return nil
}

func (m *MemDb) Open(table string) (schema.Conn, error) {
	_ = "STUB: not implemented"
	return *

	// Table by name
	new(schema.Conn), nil
}

func (m *MemDb) Table(table string) (*schema.Table, error) {
	_ = "STUB: not implemented"

	// Close this source
	return nil, nil
}

func (m *MemDb) Close() error { _ = "STUB: not implemented"; return nil }

// Tables list, should be single table
func (m *MemDb) Tables() []string { _ = "STUB: not implemented"; return nil }

func (m *MemDb) buildDefaultIndexes() { _ = "STUB: not implemented"; return }

//u.Debugf("no index provided creating on %q", m.tbl.Columns()[0])

// First ensure we have one primary index

//func (m *MemDb) SetColumns(cols []string)                  { m.tbl.SetColumns(cols) }

func newDbConn(mdb *MemDb) *dbConn { _ = "STUB: not implemented"; return nil }

func (m *dbConn) Columns() []string    { _ = "STUB: not implemented"; return nil }
func (m *dbConn) Close() error         { _ = "STUB: not implemented"; return nil }
func (m *dbConn) Next() schema.Message { _ = "STUB: not implemented"; return *new(schema.Message) }

// Put interface for allowing this to accept writes via ConnUpsert.Put()
func (m *dbConn) Put(ctx context.Context, key schema.Key, row interface{}) (schema.Key, error) {
	_ = "STUB: not implemented"
	return *new(schema.Key), nil
}

func (m *dbConn) putValues(txn *memdb.Txn, row []driver.Value) (schema.Key, error) {
	_ = "STUB: not implemented"
	return *new(schema.Key), nil
}

func (m *dbConn) PutMulti(ctx context.Context, keys []schema.Key, objs interface{}) ([]schema.Key, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *dbConn) Get(key driver.Value) (schema.Message, error) {
	_ = "STUB: not implemented"
	return *new(schema.Message), nil
}

// noop

// Should not found be an error?

// Interface for Deletion
func (m *dbConn) Delete(key driver.Value) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Delete using a Where Expression
func (m *dbConn) DeleteExpression(p interface{}, where expr.Node) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//this means do NOT delete

// Delete!

// ??

// Doesn't match, so don't delete
