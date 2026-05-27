package memdb

import (
	"database/sql/driver"

	u "github.com/araddon/gou"
	"github.com/hashicorp/go-memdb"

	"github.com/araddon/qlbridge/schema"
)

var (
	_ = u.EMPTY
	// Indexes
	_ memdb.Indexer = (*indexWrapper)(nil)
)

func makeId(dv driver.Value) uint64 { _ = "STUB: not implemented"; return 0 }

//by := append(make([]byte,0,8), byte(r), byte(r>>8), byte(r>>16), byte(r>>24), byte(r>>32), byte(r>>40), byte(r>>48), byte(r>>56))

// Wrap the index so we can operate on rows
type indexWrapper struct {
	t *schema.Table
	*schema.Index
}

func (s *indexWrapper) FromObject(obj interface{}) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// Add the null character as a terminator

// Add the null character as a terminator

func (s *indexWrapper) FromArgs(args ...interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Add the null character as a terminator

func makeMemDbSchema(m *MemDb) *memdb.DBSchema { _ = "STUB: not implemented"; return nil }

/*
	{
		"id": &memdb.IndexSchema{
			Name:    "id",
			Unique:  true,
			Indexer: &memdb.StringFieldIndex{Field: "ID"},
		},
		"foo": &memdb.IndexSchema{
			Name:    "foo",
			Indexer: &memdb.StringFieldIndex{Field: "Foo"},
		},
	},
*/
