package files

import (
	"github.com/lytics/cloudstorage"

	"github.com/araddon/qlbridge/schema"
)

var (
	// ensuure our csv handler implements FileHandler interface
	_ FileHandler = (*csvFiles)(nil)
)

func init() {
	RegisterFileHandler("csv", &csvFiles{})
}

// the built in csv filehandler
type csvFiles struct {
	appendcols []string
}

func (m *csvFiles) Init(store FileStore, ss *schema.Schema) error {
	_ = "STUB: not implemented"
	return nil
}
func (m *csvFiles) FileAppendColumns() []string { _ = "STUB: not implemented"; return nil }
func (m *csvFiles) File(path string, obj cloudstorage.Object) *FileInfo {
	_ = "STUB: not implemented"
	return nil
}

func (m *csvFiles) Scanner(store cloudstorage.StoreReader, fr *FileReader) (schema.ConnScanner, error) {
	_ = "STUB: not implemented"
	return *new(schema.ConnScanner), nil
}
