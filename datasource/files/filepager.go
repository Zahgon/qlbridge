package files

import (
	"github.com/araddon/qlbridge/exec"
	"github.com/araddon/qlbridge/plan"
	"github.com/araddon/qlbridge/schema"
)

var (
	// Our file-pager wraps our file-scanners to move onto next file
	_ FileReaderIterator  = (*FilePager)(nil)
	_ schema.ConnScanner  = (*FilePager)(nil)
	_ exec.ExecutorSource = (*FilePager)(nil)

	// Default file queue size to buffer by pager
	FileBufferSize = 5
)

// FilePager acts like a Partitionied Data Source Conn, wrapping underlying
// FileSource and paging through list of files and only scanning those that
// match this pagers partition
// - by default the partitionct is -1 which means no partitioning
type FilePager struct {
	rowct           int64
	table           string
	exit            chan bool
	err             error
	closed          bool
	fs              *FileSource
	readers         chan (*FileReader)
	partition       *schema.Partition
	partid          int
	Limit           int
	tbl             *schema.Table
	p               *plan.Source
	usePartitioning bool

	schema.ConnScanner
}

// NewFilePager creates default new FilePager
func NewFilePager(tableName string, fs *FileSource) *FilePager {
	_ = "STUB: not implemented"
	return nil
}

// WalkExecSource Provide ability to implement a source plan for execution
func (m *FilePager) WalkExecSource(p *plan.Source) (exec.Task, error) {
	_ = "STUB: not implemented"
	return *new(exec.Task), nil
}

// Columns part of Conn interface for providing columns for this table/conn
func (m *FilePager) Columns() []string { _ = "STUB: not implemented"; return nil }

// NextScanner provides the next scanner assuming that each scanner
// represents different file, and multiple files for single source
func (m *FilePager) NextScanner() (schema.ConnScanner, error) {
	_ = "STUB: not implemented"
	return *new(schema.ConnScanner), nil
}

// NextFile gets next file
func (m *FilePager) NextFile() (*FileReader, error) {
	_ = "STUB: not implemented"
	return nil,

		// See if exit was called
		nil
}

func (m *FilePager) RunFetcher() { _ = "STUB: not implemented"; return }

// fetcher process run in a go-routine to pre-fetch files
// assuming we should keep n in buffer
func (m *FilePager) fetcher() { _ = "STUB: not implemented"; return }

// was closed

// If has been closed

// Return to user

// this is expected, not all files are of file type
// we are looking for
// u.Warnf("no file?? %#v", o)

// u.Debugf("%p opening: partition:%v desiredpart:%v file: %q ", m, fi.Partition, m.partid, fi.Name)

// This will back-pressure after we reach our queue size

// Next iterator for next message, wraps the file Scanner, Next file abstractions
func (m *FilePager) Next() schema.Message { _ = "STUB: not implemented"; return *new(schema.Message) }

// Kind of crap api, side-effect method? uck

// Truly was last file in partition

// now that we have a new scanner, lets try again

// Close this connection/pager
func (m *FilePager) Close() error {
	_ = "STUB: not implemented"

	// close(m.exit)
	return nil
}
