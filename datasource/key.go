package datasource

import (
	"database/sql/driver"

	"github.com/araddon/qlbridge/schema"
)

var (
	_ schema.Key = (*KeyInt)(nil)
	_ schema.Key = (*KeyInt64)(nil)
	_ schema.Key = (*KeyCol)(nil)
)

// Variety of Key Types
type (
	KeyInt struct {
		Id int
	}
	KeyInt64 struct {
		Id int64
	}
	KeyCol struct {
		Name string
		Val  driver.Value
	}
)

func NewKeyInt(key int) KeyInt { _ = "STUB: not implemented"; return *new(KeyInt) }
func (m *KeyInt) Key() driver.Value {
	_ = "STUB: not implemented"
	return *

	// func (m KeyInt) Less(than Item) bool { return m.Id < than.(KeyInt).Id }
	new(driver.Value)
}

func NewKeyInt64(key int64) KeyInt64  { _ = "STUB: not implemented"; return *new(KeyInt64) }
func (m *KeyInt64) Key() driver.Value { _ = "STUB: not implemented"; return *new(driver.Value) }

func NewKeyCol(name string, val driver.Value) KeyCol {
	_ = "STUB: not implemented"
	return *new(KeyCol)
}
func (m KeyCol) Key() driver.Value { _ = "STUB: not implemented"; return *new(driver.Value) }

// Given a Where expression, lets try to create a key which
//
//	requires form    `idenity = "value"`
func KeyFromWhere(wh interface{}) schema.Key { _ = "STUB: not implemented"; return *new(schema.Key) }

// This only allows for    identity = value
// NOT:      identity = expr(identity, arg)
//

//case *expr.FuncNode:
