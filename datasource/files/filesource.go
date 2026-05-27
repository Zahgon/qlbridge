// Package files implements Cloud Files logic for getting, reading, and converting
// files into databases.   It reads cloud(or local) files, gets lists of tables,
// and can scan through them using distributed query engine.
package files

import (
	"fmt"
	"time"

	u "github.com/araddon/gou"
	"github.com/lytics/cloudstorage"

	"github.com/araddon/qlbridge/schema"
)

var (
	// ensure we implement interfaces
	_ schema.Source = (*FileSource)(nil)

	schemaRefreshInterval = time.Minute * 5
)

const (
	// SourceType is the registered Source name in the qlbridge source registry
	SourceType = "cloudstore"
)

func init() {
	// We need to register our DataSource provider here
	schema.RegisterSourceType(SourceType, NewFileSource())
}

// FileReaderIterator defines a file source that can page through files
// getting next file from partition
type FileReaderIterator interface {
	// NextFile returns io.EOF on last file
	NextFile() (*FileReader, error)
}

type Partitioner func(uint64, *FileInfo) int

func SipPartitioner(partitionCt uint64, fi *FileInfo) int { _ = "STUB: not implemented"; return 0 }

// FileSource Source for reading files, and scanning them allowing
// the contents to be treated as a database, like doing a full
// table scan in mysql.  But, you can partition across files.
//
//   - readers:      gcs, local-fs
//   - tablesource:  translate lists of files into tables.  Normally we would have
//     multiple files per table (ie partitioned, per-day, etc)
//   - scanners:     responsible for file-specific
//   - files table:  a "table" of all the files from this cloud source
type FileSource struct {
	ss             *schema.Schema
	lastLoad       time.Time
	store          cloudstorage.StoreReader
	fh             FileHandler
	fdbcols        []string
	fdbcolidx      map[string]int
	fdb            schema.Source
	filesTable     string
	tablenames     []string
	tableSchemas   map[string]*schema.Table
	tables         map[string]*FileTable
	path           string
	tablePerFolder bool
	fileType       string // csv, json, proto, customname
	Partitioner    string // random, ??  (date, keyed?)
	partitionFunc  Partitioner
	partitionCt    uint64
}

// NewFileSource provides a singleton manager for a particular
// Source Schema, and File-Handler to read/manage all files from
// a source such as gcs folder x, s3 folder y
func NewFileSource() *FileSource { _ = "STUB: not implemented"; return nil }

func (m *FileSource) Init() {
	_ = "STUB: not implemented"

	// Setup the filesource with schema info
	return
}

func (m *FileSource) Setup(ss *schema.Schema) error { _ = "STUB: not implemented"; return nil }

// Open a connection to given table, partition of Source interface
func (m *FileSource) Open(tableName string) (schema.Conn, error) {
	_ = "STUB: not implemented"
	//u.Debugf("Open(%q)", tableName)
	return *new(schema.Conn), nil
}

// Close this File Source manager
func (m *FileSource) Close() error {
	_ = "STUB: not implemented"

	// Tables for this file-source
	return nil
}

func (m *FileSource) Tables() []string { _ = "STUB: not implemented"; return nil }
func (m *FileSource) init() error {
	if m.store == nil {

		// u.Debugf("File init %v", string(m.ss.Conf.Settings.PrettyJson()))

		conf := m.ss.Conf.Settings
		if tablePath := conf.String("path"); tablePath != "" {
			m.path = tablePath
		}
		if fileType := conf.String("format"); fileType != "" {
			m.fileType = fileType
		} else {
			m.fileType = "csv"
		}
		if partitioner := conf.String("partitioner"); partitioner != "" {
			m.Partitioner = partitioner
		}

		store, err := FileStoreLoader(m.ss)
		if err != nil {
			u.Errorf("Could not create filestore for source %s err=%v", m.ss.Name, err)
			return err
		}
		m.store = store

		fileHandler, exists := scannerGet(m.fileType)
		if !exists || fileHandler == nil {
			return fmt.Errorf("Could not find scanner for filetype %q", m.fileType)
		}
		if err := fileHandler.Init(store, m.ss); err != nil {
			u.Errorf("Could not create filehandler for %s type=%q err=%v", m.ss.Name, m.fileType, err)
			return err
		}
		m.fh = fileHandler
		// u.Debugf("got fh: %T", m.fh)

		// ensure any additional columns are added
		m.fdbcols = append(FileColumns, m.fh.FileAppendColumns()...)

		m.fdbcolidx = make(map[string]int, len(m.fdbcols))
		for i, col := range m.fdbcols {
			m.fdbcolidx[col] = i
		}

		m.findTables()

		m.filesTable = fmt.Sprintf("%s_files", m.ss.Name)
		m.tablenames = append(m.tablenames, m.filesTable)

		// We are going to create a DB/Store to be allow the
		// entire list of files to be shown as a meta-table
		db, err := newStoreSource(m.filesTable, m)
		if err != nil {
			u.Errorf("could not create db %v", err)
			return err
		}
		m.fdb = db
	}
	return nil
}

func (m *FileSource) File(o cloudstorage.Object) *FileInfo { _ = "STUB: not implemented"; return nil }

// u.Debugf("ignoring file, path:%v  %q  is nil", m.path, o.Name())

//u.Debugf("File(%q)  path=%q", o.Name(), m.path)

func (m *FileSource) findTables() error {
	_ = "STUB: not implemented"

	// FileHandlers may optionally provide their own
	// list of files, as deciphering folder, file structure
	// isn't always obvious
	return nil
}

// u.Debugf("from path=%q  folders: %v  err=%v", m.path, folders, err)

func (m *FileSource) findTablesFromFileNames() error { _ = "STUB: not implemented"; return nil }

// If has been closed

// u.Debugf("File %s", fi)

// Table satisfys SourceSchema interface to get table schema for given table
func (m *FileSource) Table(tableName string) (*schema.Table, error) {
	_ = "STUB: not implemented"

	//u.Debugf("Table(%q) path:%v  %#v", tableName, m.path, m.ss.Conf)
	// We have a special table that is the list of all files
	return nil, nil
}

// Check cache for this table

// Its possible that the file handle implements schema handling

// Source doesn't implement Schema Handling so we are going to get
//  a scanner and introspect some rows

//u.Debugf("%p Table(%q) cols=%v", m, tableName, t.Columns())

func (m *FileSource) buildTable(tableName string) (*schema.Table, error) {
	_ = "STUB: not implemented"

	// Since we don't have a table schema, lets create one via introspection
	//u.Debugf("introspecting file-table %q for schema type=%q path=%s", tableName, m.fileType, m.path)
	return nil, nil
}

// we are going to look at ~10 rows to create schema for it

//u.Infof("built table %v %v", tableName, t.Columns())

func (m *FileSource) createPager(tableName string, partition, limit int) (*FilePager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
