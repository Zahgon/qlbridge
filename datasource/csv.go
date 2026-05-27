package datasource

import (
	"compress/gzip"
	"encoding/csv"
	"io"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/schema"
)

var (
	_ schema.Source      = (*CsvDataSource)(nil)
	_ schema.Conn        = (*CsvDataSource)(nil)
	_ schema.ConnScanner = (*CsvDataSource)(nil)
)

// Csv DataSource, implements qlbridge schema DataSource, SourceConn, Scanner
//
//	to allow csv files to be full featured databases.
//	- very, very naive scanner, forward only single pass
//	- can open a file with .Open()
//	- assumes comma delimited
//	- not thread-safe
//	- does not implement write operations
type CsvDataSource struct {
	table    string
	tbl      *schema.Table
	exit     <-chan bool
	csvr     *csv.Reader
	gz       *gzip.Reader
	rc       io.ReadCloser
	rowct    uint64
	headers  []string
	colindex map[string]int
	indexCol int
	filter   expr.Node
}

// NewCsvSource reader assumes we are getting first row as headers
// - optionally may be gzipped
func NewCsvSource(table string, indexCol int, ior io.Reader, exit <-chan bool) (*CsvDataSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO:  move this compression to the file-reader not here

// allow empty fields
// if flagCsvDelimiter == "|" {
// 	m.csvr.Comma = '|'
// } else if flagCsvDelimiter == "\t" || flagCsvDelimiter == "t" {
// 	m.csvr.Comma = '\t'
// }

//u.Debugf("headers: %v", headers)

//u.Infof("csv headers: %v colIndex: %v", headers, m.colindex)

func (m *CsvDataSource) Init()                      { _ = "STUB: not implemented"; return }
func (m *CsvDataSource) Setup(*schema.Schema) error { _ = "STUB: not implemented"; return nil }
func (m *CsvDataSource) Tables() []string           { _ = "STUB: not implemented"; return nil }
func (m *CsvDataSource) Columns() []string          { _ = "STUB: not implemented"; return nil }
func (m *CsvDataSource) Table(tableName string) (*schema.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *CsvDataSource) loadTable() error { _ = "STUB: not implemented"; return nil }

func (m *CsvDataSource) Open(connInfo string) (schema.Conn, error) {
	_ = "STUB: not implemented"
	return *new(schema.Conn), nil
}

func (m *CsvDataSource) Close() error { _ = "STUB: not implemented"; return nil }

func (m *CsvDataSource) Next() schema.Message {
	_ = "STUB: not implemented"
	return *new(schema.Message)
}

//u.Debugf("headers: %#v \n\trows:  %#v", m.headers, row)
