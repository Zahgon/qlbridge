// Membtree implements a Datasource in-memory implemenation
// using the google btree.
package membtree

import (
	"database/sql/driver"

	"github.com/google/btree"
	"golang.org/x/net/context"

	"github.com/araddon/qlbridge/datasource"
	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/schema"
)

const (
	sourceType = "membtree"
)

var (
	// Different Features of this Static Data Source
	_ schema.Source       = (*StaticDataSource)(nil)
	_ schema.Conn         = (*StaticDataSource)(nil)
	_ schema.ConnColumns  = (*StaticDataSource)(nil)
	_ schema.ConnScanner  = (*StaticDataSource)(nil)
	_ schema.ConnSeeker   = (*StaticDataSource)(nil)
	_ schema.ConnUpsert   = (*StaticDataSource)(nil)
	_ schema.ConnDeletion = (*StaticDataSource)(nil)
)

// Key implements Key and Sort interfaces.
type Key struct {
	Id uint64
}

func NewKey(key uint64) *Key             { _ = "STUB: not implemented"; return nil }
func (m *Key) Key() driver.Value         { _ = "STUB: not implemented"; return *new(driver.Value) }
func (m *Key) Less(than btree.Item) bool { _ = "STUB: not implemented"; return false }

type DriverItem struct {
	*datasource.SqlDriverMessageMap
}

func (m *DriverItem) Less(than btree.Item) bool { _ = "STUB: not implemented"; return false }

//u.Infof("Less? %p:%p less?%v gt?%v  %v vs %v   thanT:%T", m, than, m.IdVal < it.IdVal, m.IdVal > it.IdVal, m.IdVal, it.IdVal, than)

func makeId(dv driver.Value) uint64 { _ = "STUB: not implemented"; return 0 }

// iv, err := strconv.ParseUint(string(vt), 10, 64)
// if err != nil {
// 	u.Warnf("could not create id: %v  for %v", err, dv)
// }
// return iv

// iv, err := strconv.ParseUint(vt, 10, 64)
// if err != nil {
// 	u.Warnf("could not create id: %v  for %v", err, dv)
// }
// return iv

//u.Infof("got %#v", vt)

//u.Infof("got %#v", vt)

// StaticDataSource implements qlbridge DataSource to allow in memory native go data
// to have a Schema and implement and be operated on by Sql Operations
//
// Features
// - only a single column may (and must) be identified as the "Indexed" column
// - NOT threadsafe
// - each StaticDataSource = a single Table
type StaticDataSource struct {
	exit     <-chan bool
	name     string
	tbl      *schema.Table
	indexCol int        // Which column position is indexed?  ie primary key
	cursor   btree.Item // cursor position for paging
	bt       *btree.BTree
	max      int
}

func NewStaticDataSource(name string, indexedCol int, data [][]driver.Value, cols []string) *StaticDataSource {
	_ = "STUB: not implemented"

	// This source schema is a single table
	return nil
}

// StaticDataValue is used to create a static name=value pair that matches
// DataSource interfaces
func NewStaticDataValue(name string, data interface{}) *StaticDataSource {
	_ = "STUB: not implemented"
	return nil
}

func NewStaticData(name string) *StaticDataSource { _ = "STUB: not implemented"; return nil }

func (m *StaticDataSource) Init()                      { _ = "STUB: not implemented"; return }
func (m *StaticDataSource) Setup(*schema.Schema) error { _ = "STUB: not implemented"; return nil }
func (m *StaticDataSource) Open(connInfo string) (schema.Conn, error) {
	_ = "STUB: not implemented"
	return *new(schema.Conn), nil
}
func (m *StaticDataSource) Table(table string) (*schema.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
func (m *StaticDataSource) Close() error { _ = "STUB: not implemented"; return nil }
func (m *StaticDataSource) CreateIterator() schema.Iterator {
	_ = "STUB: not implemented"
	return *new(schema.Iterator)
}
func (m *StaticDataSource) Tables() []string         { _ = "STUB: not implemented"; return nil }
func (m *StaticDataSource) Columns() []string        { _ = "STUB: not implemented"; return nil }
func (m *StaticDataSource) Length() int              { _ = "STUB: not implemented"; return 0 }
func (m *StaticDataSource) SetColumns(cols []string) { _ = "STUB: not implemented"; return }

func (m *StaticDataSource) Next() schema.Message {
	_ = "STUB: not implemented"
	//u.Infof("Next()")
	return *new(schema.Message)
}

//u.Infof("create new Ascend len=%d", m.Length())

//u.Debugf("first  item btreeP:%p itemP:%p cursorP:%p  %#v", m, item, m.cursor, item)
// stop after this

//u.Debugf("equal, return true ie continue")

//u.Debugf("found  item btreeP:%p itemP:%p cursorP:%p  %#v", m, item, m.cursor, item)
// stop after this

// if m.max > 20 {
// 	return nil
// }

//u.Debugf("reset cursor to nil  %#v", item)

//u.Infof("return item btreeP:%p itemP:%p cursorP:%p  %v %v", m, item, m.cursor, msg.Id(), msg.Values())
//u.Debugf("return? %T  %v", item, item.(*DriverItem).SqlDriverMessageMap)

//return datasource.NewSqlDriverMessageMapVals(uint64(m.cursor-1), m.data[m.cursor-1], m.cols)

// interface for Upsert.Put()
func (m *StaticDataSource) Put(ctx context.Context, key schema.Key, row interface{}) (schema.Key, error) {
	_ = "STUB: not implemented"

	//u.Infof("%p Put(),  row:%#v", m, row)
	return *new(schema.Key), nil
}

//u.Errorf("could not insert? %#v", itemResult)

//u.Debugf("%p  PUT: id:%v IdVal:%v  Id():%v vals:%#v", m, id, sdm.IdVal, sdm.Id(), rowVals)

// We need to convert the key:value to []driver.Value so
// we need to look up column index for each key, and write to vals

// TODO:   if this is a partial update, we need to look up vals

// How do we get the key?
//m.Get(key)

// Since we do not have an indexed column to work off of,
// the ideal would be to get the job builder/planner to do
// a scan with whatever info we have and feed that in?   Instead
// of us implementing our own scan?

//u.Debugf("sdm: %#v  err%v", sdm, err)

//u.Debugf("PUT: %#v", row)
//u.Infof("PUT: %v  key:%v  row:%v", id, key, row)

func (m *StaticDataSource) PutMulti(ctx context.Context, keys []schema.Key, src interface{}) ([]schema.Key, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *StaticDataSource) Get(key driver.Value) (schema.Message, error) {
	_ = "STUB: not implemented"
	return *new(schema.Message), nil
}

// Should not found be an error?

func (m *StaticDataSource) MultiGet(keys []driver.Value) ([]schema.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Interface for Deletion
func (m *StaticDataSource) Delete(key driver.Value) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//u.Warnf("could not delete: %v", key)

// DeleteExpression Delete using a Where Expression
func (m *StaticDataSource) DeleteExpression(p interface{}, where expr.Node) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//return deletedCt, fmt.Errorf("Could not evaluate where clause")

//this means do NOT delete

// Delete!

// ??

// Doesn't match, so don't delete
