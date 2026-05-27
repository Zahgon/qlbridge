package datasource

import (
	"database/sql/driver"
	"time"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/schema"
	"github.com/araddon/qlbridge/value"
)

var (
	// Ensure our ContextReaders implement interface
	// context-readers hold "State" for evaluation in vm.
	_ expr.ContextReader = (*ContextSimple)(nil)
	_ expr.ContextReader = (*SqlDriverMessageMap)(nil)
	// Context Writers hold write state of vm
	_ expr.ContextWriter = (*ContextSimple)(nil)
	// Message is passed between tasks/actors across distributed
	// boundaries.   May be context reader/writer.
	_ schema.Message = (*ContextSimple)(nil)
	_ schema.Message = (*SqlDriverMessage)(nil)
	_ schema.Message = (*SqlDriverMessageMap)(nil)
)

// MessageConversion convert values of type schema.Message.
func MessageConversion(vals []interface{}) []schema.Message { _ = "STUB: not implemented"; return nil }

type (
	// SqlDriverMessage context message of values.
	SqlDriverMessage struct {
		Vals  []driver.Value
		IdVal uint64
	}
	// SqlDriverMessageMap Context message with column/position info.
	SqlDriverMessageMap struct {
		Vals     []driver.Value // Values
		ColIndex map[string]int // Map of column names to ordinal position in vals
		IdVal    uint64         // id()
		keyVal   string         // key   Non Hashed Key Value
	}
	ContextSimple struct {
		Data        map[string]value.Value
		ts          time.Time
		cursor      int
		keyval      uint64
		namespacing bool
	}
	NestedContextReader struct {
		readers []expr.ContextReader
		writer  expr.ContextWriter
		ts      time.Time
	}
	NamespacedContextReader struct {
		basereader expr.ContextReader
		namespace  string
	}
)

func NewSqlDriverMessage(id uint64, row []driver.Value) *SqlDriverMessage {
	_ = "STUB: not implemented"
	return nil
}

func (m *SqlDriverMessage) Id() uint64        { _ = "STUB: not implemented"; return 0 }
func (m *SqlDriverMessage) Body() interface{} { _ = "STUB: not implemented"; return nil }
func (m *SqlDriverMessage) ToMsgMap(colidx map[string]int) *SqlDriverMessageMap {
	_ = "STUB: not implemented"
	return nil
}

func NewSqlDriverMessageMapEmpty() *SqlDriverMessageMap { _ = "STUB: not implemented"; return nil }

func NewSqlDriverMessageMap(id uint64, row []driver.Value, colindex map[string]int) *SqlDriverMessageMap {
	_ = "STUB: not implemented"
	return nil
}

func NewSqlDriverMessageMapVals(id uint64, row []driver.Value, cols []string) *SqlDriverMessageMap {
	_ = "STUB: not implemented"
	return nil
}

func NewSqlDriverMessageMapCtx(id uint64, ctx expr.ContextReader, colindex map[string]int) *SqlDriverMessageMap {
	_ = "STUB: not implemented"
	return nil
}

func (m *SqlDriverMessageMap) Id() uint64 { _ = "STUB: not implemented"; return 0 }
func (m *SqlDriverMessageMap) Key() driver.Value {
	_ = "STUB: not implemented"
	return *new(driver.Value)
}
func (m *SqlDriverMessageMap) SetKey(key string)       { _ = "STUB: not implemented"; return }
func (m *SqlDriverMessageMap) SetKeyHashed(key string) { _ = "STUB: not implemented"; return }

func (m *SqlDriverMessageMap) Body() interface{}         { _ = "STUB: not implemented"; return nil }
func (m *SqlDriverMessageMap) Values() []driver.Value    { _ = "STUB: not implemented"; return nil }
func (m *SqlDriverMessageMap) SetRow(row []driver.Value) { _ = "STUB: not implemented"; return }
func (m *SqlDriverMessageMap) Ts() time.Time             { _ = "STUB: not implemented"; return *new(time.Time) }
func (m *SqlDriverMessageMap) Get(key string) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

//u.Debugf("could not find: %q  right=%q hasLeftRight?%v", key, right, hasLeft)

func (m *SqlDriverMessageMap) Row() map[string]value.Value { _ = "STUB: not implemented"; return nil }

func (m *SqlDriverMessageMap) Copy() *SqlDriverMessageMap { _ = "STUB: not implemented"; return nil }

// we assume? that values are immutable anyways

func NewContextSimple() *ContextSimple { _ = "STUB: not implemented"; return nil }

func NewContextSimpleData(data map[string]value.Value) *ContextSimple {
	_ = "STUB: not implemented"
	return nil
}

func NewContextSimpleNative(data map[string]interface{}) *ContextSimple {
	_ = "STUB: not implemented"
	return nil
}

func NewContextMap(data map[string]interface{}, namespacing bool) *ContextSimple {
	_ = "STUB: not implemented"
	return nil
}

func NewContextMapTs(data map[string]interface{}, namespacing bool, ts time.Time) *ContextSimple {
	_ = "STUB: not implemented"
	return nil
}

func NewContextSimpleTs(data map[string]value.Value, ts time.Time) *ContextSimple {
	_ = "STUB: not implemented"
	return nil
}

func (m *ContextSimple) SupportNamespacing()         { _ = "STUB: not implemented"; return }
func (m *ContextSimple) All() map[string]value.Value { _ = "STUB: not implemented"; return nil }
func (m *ContextSimple) Row() map[string]value.Value { _ = "STUB: not implemented"; return nil }
func (m *ContextSimple) Body() interface{}           { _ = "STUB: not implemented"; return nil }
func (m *ContextSimple) Id() uint64                  { _ = "STUB: not implemented"; return 0 }
func (m *ContextSimple) Ts() time.Time               { _ = "STUB: not implemented"; return *new(time.Time) }
func (m ContextSimple) Get(key string) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// We don't support namespacing by default?

func (m *ContextSimple) Put(col expr.SchemaInfo, rctx expr.ContextReader, v value.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *ContextSimple) Commit(rowInfo []expr.SchemaInfo, row expr.RowWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *ContextSimple) Delete(row map[string]value.Value) error {
	_ = "STUB: not implemented"

	// NewNestedContextReader provides a context reader which is a composite of ordered child readers
	// the first reader with a key will be used
	return nil
}

func NewNestedContextReader(readers []expr.ContextReader, ts time.Time) expr.ContextReader {
	_ = "STUB: not implemented"
	return *new(expr.ContextReader)
}

// NewNestedContextReader provides a context reader which is a composite of ordered child readers
// the first reader with a key will be used
func NewNestedContextReadWriter(readers []expr.ContextReader, writer expr.ContextWriter, ts time.Time) expr.ContextReadWriter {
	_ = "STUB: not implemented"
	return *new(expr.ContextReadWriter)
}

func (n *NestedContextReader) Get(key string) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

func (n *NestedContextReader) Row() map[string]value.Value { _ = "STUB: not implemented"; return nil }

// already added this key from a "higher priority" reader

func (n *NestedContextReader) Ts() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (n *NestedContextReader) Put(col expr.SchemaInfo, readCtx expr.ContextReader, v value.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *NestedContextReader) Delete(delRow map[string]value.Value) error {
	_ = "STUB: not implemented"
	return nil
}

// NewNestedContextReader provides a context reader which prefixes
// all keys with a name space.  This is useful if you have overlapping
// field names between ContextReaders within a NestedContextReader.
//
//	msg.Get("foo.key")
func NewNamespacedContextReader(basereader expr.ContextReader, namespace string) expr.ContextReader {
	_ = "STUB: not implemented"
	return *new(expr.ContextReader)
}

func (n *NamespacedContextReader) Get(key string) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

func (n *NamespacedContextReader) Row() map[string]value.Value {
	_ = "STUB: not implemented"
	return nil
}

func (n *NamespacedContextReader) Ts() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
