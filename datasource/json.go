package datasource

import (
	"bufio"
	"compress/gzip"
	"io"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/schema"
)

var (
	_ schema.Source      = (*JsonSource)(nil)
	_ schema.Conn        = (*JsonSource)(nil)
	_ schema.ConnScanner = (*JsonSource)(nil)
)

type FileLineHandler func(line []byte) (schema.Message, error)

// JsonSource implements qlbridge schema DataSource, SourceConn, Scanner
// to allow new line delimited json files to be full featured databases.
// - very, very naive scanner, forward only single pass
// - can open a file with .Open()
// - not thread-safe
// - does not implement write operations
type JsonSource struct {
	table    string
	tbl      *schema.Table
	exit     <-chan bool
	complete bool
	err      error
	r        *bufio.Reader
	gz       *gzip.Reader
	rc       io.ReadCloser
	rowct    uint64
	lhSpec   FileLineHandler
	lh       FileLineHandler
	columns  []string
	colindex map[string]int
	indexCol int
	filter   expr.Node
}

// NewJsonSource reader assumes we are getting NEW LINE delimted json file
// - optionally may be gzipped
func NewJsonSource(table string, rc io.ReadCloser, exit <-chan bool, lh FileLineHandler) (*JsonSource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Gzip Files have these 2 byte prefix

//m.loadTable()

func (m *JsonSource) Init()                      { _ = "STUB: not implemented"; return }
func (m *JsonSource) Setup(*schema.Schema) error { _ = "STUB: not implemented"; return nil }
func (m *JsonSource) Tables() []string           { _ = "STUB: not implemented"; return nil }
func (m *JsonSource) Columns() []string          { _ = "STUB: not implemented"; return nil }
func (m *JsonSource) CreateIterator() schema.Iterator {
	_ = "STUB: not implemented"
	return *new(schema.Iterator)
}
func (m *JsonSource) Table(tableName string) (*schema.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *JsonSource) loadTable() error { _ = "STUB: not implemented"; return nil }

func (m *JsonSource) Open(connInfo string) (schema.Conn, error) {
	_ = "STUB: not implemented"
	return *new(schema.Conn), nil
}

func (m *JsonSource) Close() error { _ = "STUB: not implemented"; return nil }

func (m *JsonSource) Next() schema.Message { _ = "STUB: not implemented"; return *new(schema.Message) }

func (m *JsonSource) jsonDefaultLine(line []byte) (schema.Message, error) {
	_ = "STUB: not implemented"
	return *new(schema.Message), nil
}
