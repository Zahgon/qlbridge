package files

import (
	"sync"

	"github.com/lytics/cloudstorage"

	"github.com/araddon/qlbridge/schema"
)

var (
	// Ensure we implement Source for our file source storage
	_ schema.Source = (*storeSource)(nil)

	// Connection Interfaces
	_ schema.Conn        = (*storeSource)(nil)
	_ schema.ConnScanner = (*storeSource)(nil)
)

// storeSource DataSource for reading lists of files/names/metadata of files
// from the cloudstorage Store
//
// - readers:      s3, gcs, local-fs
type storeSource struct {
	f        *FileSource
	table    string
	tbl      *schema.Table
	exit     <-chan bool
	iter     cloudstorage.ObjectIterator
	complete bool
	err      error
	rowct    uint64
	mu       sync.Mutex
}

// newStoreSource reader
func newStoreSource(table string, fs *FileSource) (*storeSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *storeSource) Init()                      { _ = "STUB: not implemented"; return }
func (m *storeSource) Setup(*schema.Schema) error { _ = "STUB: not implemented"; return nil }
func (m *storeSource) Tables() []string           { _ = "STUB: not implemented"; return nil }
func (m *storeSource) Columns() []string          { _ = "STUB: not implemented"; return nil }
func (m *storeSource) CreateIterator() schema.Iterator {
	_ = "STUB: not implemented"
	return *new(schema.Iterator)
}
func (m *storeSource) Table(tableName string) (*schema.Table, error) {
	_ = "STUB: not implemented"
	// u.Debugf("Table(%q), tbl nil?%v", tableName, m.tbl == nil)
	return nil, nil
}

func (m *storeSource) loadTable() error { _ = "STUB: not implemented"; return nil }

// u.Debugf("storeSource.loadTable(%q)", m.table)

func (m *storeSource) Open(connInfo string) (schema.Conn, error) {
	_ = "STUB: not implemented"

	// u.Debugf("Open(%q)", connInfo)
	// Make a copy of itself
	return *new(schema.Conn), nil
}

func (m *storeSource) Close() error { _ = "STUB: not implemented"; return nil }

func (m *storeSource) Next() schema.Message { _ = "STUB: not implemented"; return *new(schema.Message) }

// Should we Retry?
