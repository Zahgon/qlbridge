// sqlite implements a Qlbridge Datasource interface around sqlite.
package sqlite

import (
	"database/sql"
	"sync"

	// Import Sqlite driver
	_ "github.com/mattn/go-sqlite3"

	"github.com/araddon/qlbridge/schema"
)

const (
	// SourceType "sqlite" is the registered Source name in the qlbridge source registry
	SourceType = "sqlite"
)

func init() {
	// We need to register our DataSource provider here
	schema.RegisterSourceType(SourceType, newSourceEmtpy())
}

var (
	// Ensure our source implements Source interface
	_ schema.Source = (*Source)(nil)
	// ensure our Source implements connection features
	_ schema.Conn = (*Source)(nil)
)

// Source implements qlbridge DataSource to a sqlite file based source.
//
// Features
// - Support full predicate push down to SqlLite.
// - Support Thread-Safe wrapper around sqlite file.
type Source struct {
	exit      <-chan bool
	schema    *schema.Schema
	file      string // Local file path to sqlite db
	db        *sql.DB
	mu        sync.Mutex
	source    *Source
	qryconns  map[string]*qryconn
	tables    map[string]*schema.Table
	tblmu     sync.Mutex
	tableList []string
}

func newSourceEmtpy() schema.Source { _ = "STUB: not implemented"; return *new(schema.Source) }

// Setup this source with schema from parent.
func (m *Source) Setup(s *schema.Schema) error { _ = "STUB: not implemented"; return nil }

// It will be created if it doesn't exist.
//   "./source.enriched.db"

// SELECT * FROM dbname.sqlite_master WHERE type='table';

// if err := datasource.IntrospectTable(m.tbl, m.CreateIterator()); err != nil {
// 	u.Errorf("Could not introspect schema %v", err)
// }

// Init the source
func (m *Source) Init() {
	_ = "STUB: not implemented"

	// Open a connection, since sqlite is not threadsafe, this is locked.
	return
}

func (m *Source) Open(table string) (schema.Conn, error) {
	_ = "STUB: not implemented"
	//u.Infof("Open conn=%q", table)
	return *new(schema.Conn), nil
}

//u.Infof("after open lock")

// Table gets table schema for given table
func (m *Source) Table(table string) (*schema.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tables gets list of tables
func (m *Source) Tables() []string {
	_ = "STUB: not implemented"

	// Close this source, closing the underlying sqlite db file
	return nil
}

func (m *Source) Close() error { _ = "STUB: not implemented"; return nil }

func tableFromSQL(name, sqls string) *schema.Table { _ = "STUB: not implemented"; return nil }

//u.Debugf("%s  %v", name, sqls)

// NewFieldBase(name string, valType value.ValueType, size int, desc string)

// u.Debugf("%d  %v", i, parts)
// u.Debugf("%q", expr.IdentityTrim(parts[0]))
