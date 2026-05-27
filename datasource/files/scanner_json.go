package files

import (
	"github.com/lytics/cloudstorage"

	"github.com/araddon/qlbridge/datasource"
	"github.com/araddon/qlbridge/schema"
)

var (
	// ensuure our json handler implements FileHandler interface
	_ FileHandler = (*jsonHandler)(nil)
)

func init() {
	RegisterFileHandler("json", &jsonHandler{})
}

// the built in json filehandler
type jsonHandler struct {
	parser datasource.FileLineHandler
}

// the built in json filehandler
type jsonHandlerTables struct {
	tables []string
	FileHandler
}

// NewJsonHandler creates a json file handler for paging new-line
// delimited rows of json file
func NewJsonHandler(lh datasource.FileLineHandler) FileHandler {
	_ = "STUB: not implemented"
	return *

	// NewJsonHandler creates a json file handler for paging new-line
	// delimited rows of json file
	new(FileHandler)
}

func NewJsonHandlerTables(lh datasource.FileLineHandler, tables []string) FileHandler {
	_ = "STUB: not implemented"
	return *new(FileHandler)
}

func (m *jsonHandler) Init(store FileStore, ss *schema.Schema) error {
	_ = "STUB: not implemented"
	return nil
}
func (m *jsonHandler) FileAppendColumns() []string { _ = "STUB: not implemented"; return nil }
func (m *jsonHandler) File(path string, obj cloudstorage.Object) *FileInfo {
	_ = "STUB: not implemented"
	return nil
}

func (m *jsonHandler) Scanner(store cloudstorage.StoreReader, fr *FileReader) (schema.ConnScanner, error) {
	_ = "STUB: not implemented"
	return *new(schema.ConnScanner), nil
}

func (m *jsonHandlerTables) Tables() []string { _ = "STUB: not implemented"; return nil }
