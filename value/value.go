// Value package defines the core value types (string, int, etc) for the
// qlbridge package, mostly used to provide common interfaces instead
// of reflection for virtual machine.
package value

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
)

var (
	nilStruct   *emptyStruct
	EmptyStruct = struct{}{}

	NilValueVal         = NewNilValue()
	BoolValueTrue       = BoolValue{v: true}
	BoolValueFalse      = BoolValue{v: false}
	NumberNaNValue      = NewNumberValue(math.NaN())
	EmptyStringValue    = NewStringValue("")
	EmptyStringsValue   = NewStringsValue(nil)
	EmptyMapValue       = NewMapValue(nil)
	EmptyMapStringValue = NewMapStringValue(make(map[string]string))
	EmptyMapIntValue    = NewMapIntValue(make(map[string]int64))
	EmptyMapNumberValue = NewMapNumberValue(make(map[string]float64))
	EmptyMapTimeValue   = NewMapTimeValue(make(map[string]time.Time))
	EmptyMapBoolValue   = NewMapBoolValue(make(map[string]bool))
	NilStructValue      = NewStructValue(nilStruct)
	TimeZeroValue       = NewTimeValue(time.Time{})
	ErrValue            = NewErrorValue(fmt.Errorf(""))

	_ Value = (StringValue)(EmptyStringValue)

	// force some types to implement interfaces
	_ Slice = (*StringsValue)(nil)
	_ Slice = (*SliceValue)(nil)
	_ Map   = (MapValue)(EmptyMapValue)
	_ Map   = (MapIntValue)(EmptyMapIntValue)
	_ Map   = (MapStringValue)(EmptyMapStringValue)
	_ Map   = (MapNumberValue)(EmptyMapNumberValue)
	_ Map   = (MapTimeValue)(EmptyMapTimeValue)
	_ Map   = (MapBoolValue)(EmptyMapBoolValue)
)

// This is the DataType system, ie string, int, etc
type ValueType uint8

const (
	// Enum values for Type system, DO NOT CHANGE the numbers, do not use iota
	NilType            ValueType = 0
	ErrorType          ValueType = 1
	UnknownType        ValueType = 2
	ValueInterfaceType ValueType = 3 // Is of type Value Interface, ie unknown
	NumberType         ValueType = 10
	IntType            ValueType = 11
	BoolType           ValueType = 12
	TimeType           ValueType = 13
	ByteSliceType      ValueType = 14
	StringType         ValueType = 20
	StringsType        ValueType = 21
	MapValueType       ValueType = 30
	MapIntType         ValueType = 31
	MapStringType      ValueType = 32
	MapNumberType      ValueType = 33
	MapBoolType        ValueType = 34
	MapTimeType        ValueType = 35
	SliceValueType     ValueType = 40
	StructType         ValueType = 50
	JsonType           ValueType = 51
)

func (m ValueType) String() string { _ = "STUB: not implemented"; return "" }

func (m ValueType) IsMap() bool { _ = "STUB: not implemented"; return false }

func (m ValueType) IsSlice() bool { _ = "STUB: not implemented"; return false }

func (m ValueType) IsNumeric() bool { _ = "STUB: not implemented"; return false }

type emptyStruct struct{}

type (
	Value interface {
		// Is this a nil/empty?
		// empty string counts as nil, empty slices/maps, nil structs.
		Nil() bool
		// Is this an error, or unable to evaluate from Vm?
		Err() bool
		Value() interface{}
		ToString() string
		Type() ValueType
	}
	// Certain types are Numeric (Ints, Time, Number)
	NumericValue interface {
		Float() float64
		Int() int64
	}
	// Slices can always return a []Value representation and is meant to be used
	// when iterating over all items in a non-scalar value. Maps return their keys
	// as a slice.
	Slice interface {
		SliceValue() []Value
		Len() int
		json.Marshaler
	}
	// Map interface
	Map interface {
		json.Marshaler
		Len() int
		MapValue() MapValue
		Get(key string) (Value, bool)
	}
)

type (
	NumberValue struct {
		v float64
	}
	IntValue struct {
		v int64
	}
	BoolValue struct {
		v bool
	}
	StringValue struct {
		v string
	}
	TimeValue struct {
		v time.Time
	}
	StringsValue struct {
		v []string
	}
	ByteSliceValue struct {
		v []byte
	}
	SliceValue struct {
		v []Value
	}
	MapValue struct {
		v map[string]Value
	}
	MapIntValue struct {
		v map[string]int64
	}
	MapNumberValue struct {
		v map[string]float64
	}
	MapStringValue struct {
		v map[string]string
	}
	MapBoolValue struct {
		v map[string]bool
	}
	MapTimeValue struct {
		v map[string]time.Time
	}
	StructValue struct {
		v interface{}
	}
	JsonValue struct {
		v json.RawMessage
	}
	ErrorValue struct {
		v error
	}
	NilValue struct{}
)

// ValueFromString Given a string, convert to valuetype
func ValueFromString(vt string) ValueType { _ = "STUB: not implemented"; return *new(ValueType) }

// NewValue creates a new Value type from a native Go value.
//
// Defaults to StructValue for unknown types.
func NewValue(goVal interface{}) Value { _ = "STUB: not implemented"; return *new(Value) }

// should we return Nil?
// if val == "null" || val == "NULL" {}

// case []uint8:
// 	return NewByteSliceValue([]byte(val))

func NewNumberValue(v float64) NumberValue { _ = "STUB: not implemented"; return *new(NumberValue) }

func NewNumberNil() NumberValue { _ = "STUB: not implemented"; return *new(NumberValue) }

func (m NumberValue) Nil() bool                    { _ = "STUB: not implemented"; return false }
func (m NumberValue) Err() bool                    { _ = "STUB: not implemented"; return false }
func (m NumberValue) Type() ValueType              { _ = "STUB: not implemented"; return *new(ValueType) }
func (m NumberValue) Value() interface{}           { _ = "STUB: not implemented"; return nil }
func (m NumberValue) Val() float64                 { _ = "STUB: not implemented"; return 0 }
func (m NumberValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m NumberValue) ToString() string             { _ = "STUB: not implemented"; return "" }
func (m NumberValue) Float() float64               { _ = "STUB: not implemented"; return 0 }
func (m NumberValue) Int() int64                   { _ = "STUB: not implemented"; return 0 }

func NewIntValue(v int64) IntValue { _ = "STUB: not implemented"; return *new(IntValue) }

func NewIntNil() IntValue { _ = "STUB: not implemented"; return *new(IntValue) }

func (m IntValue) Nil() bool                    { _ = "STUB: not implemented"; return false }
func (m IntValue) Err() bool                    { _ = "STUB: not implemented"; return false }
func (m IntValue) Type() ValueType              { _ = "STUB: not implemented"; return *new(ValueType) }
func (m IntValue) Value() interface{}           { _ = "STUB: not implemented"; return nil }
func (m IntValue) Val() int64                   { _ = "STUB: not implemented"; return 0 }
func (m IntValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m IntValue) NumberValue() NumberValue     { _ = "STUB: not implemented"; return *new(NumberValue) }
func (m IntValue) ToString() string             { _ = "STUB: not implemented"; return "" }

func (m IntValue) Float() float64 { _ = "STUB: not implemented"; return 0 }
func (m IntValue) Int() int64     { _ = "STUB: not implemented"; return 0 }

func NewBoolValue(v bool) BoolValue { _ = "STUB: not implemented"; return *new(BoolValue) }

func (m BoolValue) Nil() bool                    { _ = "STUB: not implemented"; return false }
func (m BoolValue) Err() bool                    { _ = "STUB: not implemented"; return false }
func (m BoolValue) Type() ValueType              { _ = "STUB: not implemented"; return *new(ValueType) }
func (m BoolValue) Value() interface{}           { _ = "STUB: not implemented"; return nil }
func (m BoolValue) Val() bool                    { _ = "STUB: not implemented"; return false }
func (m BoolValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m BoolValue) ToString() string             { _ = "STUB: not implemented"; return "" }

func NewStringValue(v string) StringValue { _ = "STUB: not implemented"; return *new(StringValue) }

func (m StringValue) Nil() bool                    { _ = "STUB: not implemented"; return false }
func (m StringValue) Err() bool                    { _ = "STUB: not implemented"; return false }
func (m StringValue) Type() ValueType              { _ = "STUB: not implemented"; return *new(ValueType) }
func (m StringValue) Value() interface{}           { _ = "STUB: not implemented"; return nil }
func (m StringValue) Val() string                  { _ = "STUB: not implemented"; return "" }
func (m StringValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m StringValue) NumberValue() NumberValue     { _ = "STUB: not implemented"; return *new(NumberValue) }

func (m StringValue) StringsValue() StringsValue {
	_ = "STUB: not implemented"
	return *new(StringsValue)
}
func (m StringValue) ToString() string { _ = "STUB: not implemented"; return "" }

func (m StringValue) IntValue() IntValue { _ = "STUB: not implemented"; return *new(IntValue) }

func NewStringsValue(v []string) StringsValue { _ = "STUB: not implemented"; return *new(StringsValue) }

func (m StringsValue) Nil() bool                    { _ = "STUB: not implemented"; return false }
func (m StringsValue) Err() bool                    { _ = "STUB: not implemented"; return false }
func (m StringsValue) Type() ValueType              { _ = "STUB: not implemented"; return *new(ValueType) }
func (m StringsValue) Value() interface{}           { _ = "STUB: not implemented"; return nil }
func (m StringsValue) Val() []string                { _ = "STUB: not implemented"; return nil }
func (m *StringsValue) Append(sv string)            { _ = "STUB: not implemented"; return }
func (m StringsValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m StringsValue) Len() int                     { _ = "STUB: not implemented"; return 0 }
func (m StringsValue) NumberValue() NumberValue {
	_ = "STUB: not implemented"
	return *new(NumberValue)
}

func (m StringsValue) IntValue() IntValue {
	_ = "STUB: not implemented"
	// Im not confident this is valid?   array first element?
	return *new(IntValue)
}

func (m StringsValue) ToString() string         { _ = "STUB: not implemented"; return "" }
func (m StringsValue) Strings() []string        { _ = "STUB: not implemented"; return nil }
func (m StringsValue) Set() map[string]struct{} { _ = "STUB: not implemented"; return nil }

func (m StringsValue) SliceValue() []Value { _ = "STUB: not implemented"; return nil }

func NewByteSliceValue(v []byte) ByteSliceValue {
	_ = "STUB: not implemented"
	return *new(ByteSliceValue)
}

func (m ByteSliceValue) Nil() bool                    { _ = "STUB: not implemented"; return false }
func (m ByteSliceValue) Err() bool                    { _ = "STUB: not implemented"; return false }
func (m ByteSliceValue) Type() ValueType              { _ = "STUB: not implemented"; return *new(ValueType) }
func (m ByteSliceValue) Value() interface{}           { _ = "STUB: not implemented"; return nil }
func (m ByteSliceValue) Val() []byte                  { _ = "STUB: not implemented"; return nil }
func (m ByteSliceValue) ToString() string             { _ = "STUB: not implemented"; return "" }
func (m ByteSliceValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m ByteSliceValue) Len() int                     { _ = "STUB: not implemented"; return 0 }

func NewSliceValues(v []Value) SliceValue { _ = "STUB: not implemented"; return *new(SliceValue) }

func NewSliceValuesNative(iv []interface{}) SliceValue {
	_ = "STUB: not implemented"
	return *new(SliceValue)
}

func (m SliceValue) Nil() bool          { _ = "STUB: not implemented"; return false }
func (m SliceValue) Err() bool          { _ = "STUB: not implemented"; return false }
func (m SliceValue) Type() ValueType    { _ = "STUB: not implemented"; return *new(ValueType) }
func (m SliceValue) Value() interface{} { _ = "STUB: not implemented"; return nil }
func (m SliceValue) Val() []Value       { _ = "STUB: not implemented"; return nil }
func (m SliceValue) ToString() string   { _ = "STUB: not implemented"; return "" }

func (m *SliceValue) Append(v Value)              { _ = "STUB: not implemented"; return }
func (m SliceValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m SliceValue) Len() int                     { _ = "STUB: not implemented"; return 0 }
func (m SliceValue) SliceValue() []Value          { _ = "STUB: not implemented"; return nil }
func (m SliceValue) Values() []interface{}        { _ = "STUB: not implemented"; return nil }

func NewMapValue(v map[string]interface{}) MapValue {
	_ = "STUB: not implemented"
	return *new(MapValue)
}

func (m MapValue) Nil() bool                    { _ = "STUB: not implemented"; return false }
func (m MapValue) Err() bool                    { _ = "STUB: not implemented"; return false }
func (m MapValue) Type() ValueType              { _ = "STUB: not implemented"; return *new(ValueType) }
func (m MapValue) Value() interface{}           { _ = "STUB: not implemented"; return nil }
func (m MapValue) Val() map[string]Value        { _ = "STUB: not implemented"; return nil }
func (m MapValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m MapValue) ToString() string             { _ = "STUB: not implemented"; return "" }
func (m MapValue) Len() int                     { _ = "STUB: not implemented"; return 0 }
func (m MapValue) MapInt() map[string]int64     { _ = "STUB: not implemented"; return nil }

func (m MapValue) MapFloat() map[string]float64 { _ = "STUB: not implemented"; return nil }

func (m MapValue) MapString() map[string]string { _ = "STUB: not implemented"; return nil }

func (m MapValue) MapValue() MapValue { _ = "STUB: not implemented"; return *new(MapValue) }

func (m MapValue) MapTime() MapTimeValue { _ = "STUB: not implemented"; return *new(MapTimeValue) }

func (m MapValue) Get(key string) (Value, bool) {
	_ = "STUB: not implemented"
	return *new(Value), false
}

func NewMapStringValue(v map[string]string) MapStringValue {
	_ = "STUB: not implemented"
	return *new(MapStringValue)
}

func (m MapStringValue) Nil() bool                    { _ = "STUB: not implemented"; return false }
func (m MapStringValue) Err() bool                    { _ = "STUB: not implemented"; return false }
func (m MapStringValue) Type() ValueType              { _ = "STUB: not implemented"; return *new(ValueType) }
func (m MapStringValue) Value() interface{}           { _ = "STUB: not implemented"; return nil }
func (m MapStringValue) Val() map[string]string       { _ = "STUB: not implemented"; return nil }
func (m MapStringValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m MapStringValue) ToString() string             { _ = "STUB: not implemented"; return "" }
func (m MapStringValue) Len() int                     { _ = "STUB: not implemented"; return 0 }
func (m MapStringValue) MapBool() MapBoolValue {
	_ = "STUB: not implemented"
	return *new(MapBoolValue)
}

func (m MapStringValue) MapInt() MapIntValue { _ = "STUB: not implemented"; return *new(MapIntValue) }

func (m MapStringValue) MapNumber() MapNumberValue {
	_ = "STUB: not implemented"
	return *new(MapNumberValue)
}

func (m MapStringValue) MapValue() MapValue { _ = "STUB: not implemented"; return *new(MapValue) }

func (m MapStringValue) Get(key string) (Value, bool) {
	_ = "STUB: not implemented"
	return *new(Value), false
}

func (m MapStringValue) SliceValue() []Value { _ = "STUB: not implemented"; return nil }

func NewMapIntValue(v map[string]int64) MapIntValue {
	_ = "STUB: not implemented"
	return *new(MapIntValue)
}

func (m MapIntValue) Nil() bool                    { _ = "STUB: not implemented"; return false }
func (m MapIntValue) Err() bool                    { _ = "STUB: not implemented"; return false }
func (m MapIntValue) Type() ValueType              { _ = "STUB: not implemented"; return *new(ValueType) }
func (m MapIntValue) Value() interface{}           { _ = "STUB: not implemented"; return nil }
func (m MapIntValue) Val() map[string]int64        { _ = "STUB: not implemented"; return nil }
func (m MapIntValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m MapIntValue) ToString() string             { _ = "STUB: not implemented"; return "" }
func (m MapIntValue) Len() int                     { _ = "STUB: not implemented"; return 0 }
func (m MapIntValue) MapInt() map[string]int64     { _ = "STUB: not implemented"; return nil }
func (m MapIntValue) MapFloat() map[string]float64 { _ = "STUB: not implemented"; return nil }

func (m MapIntValue) MapValue() MapValue { _ = "STUB: not implemented"; return *new(MapValue) }

func (m MapIntValue) Get(key string) (Value, bool) {
	_ = "STUB: not implemented"
	return *new(Value), false
}

func (m MapIntValue) SliceValue() []Value { _ = "STUB: not implemented"; return nil }

func NewMapNumberValue(v map[string]float64) MapNumberValue {
	_ = "STUB: not implemented"
	return *new(MapNumberValue)
}

func (m MapNumberValue) Nil() bool                    { _ = "STUB: not implemented"; return false }
func (m MapNumberValue) Err() bool                    { _ = "STUB: not implemented"; return false }
func (m MapNumberValue) Type() ValueType              { _ = "STUB: not implemented"; return *new(ValueType) }
func (m MapNumberValue) Value() interface{}           { _ = "STUB: not implemented"; return nil }
func (m MapNumberValue) Val() map[string]float64      { _ = "STUB: not implemented"; return nil }
func (m MapNumberValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m MapNumberValue) ToString() string             { _ = "STUB: not implemented"; return "" }
func (m MapNumberValue) Len() int                     { _ = "STUB: not implemented"; return 0 }
func (m MapNumberValue) MapInt() map[string]int64     { _ = "STUB: not implemented"; return nil }

func (m MapNumberValue) MapValue() MapValue { _ = "STUB: not implemented"; return *new(MapValue) }

func (m MapNumberValue) Get(key string) (Value, bool) {
	_ = "STUB: not implemented"
	return *new(Value), false
}

func (m MapNumberValue) SliceValue() []Value { _ = "STUB: not implemented"; return nil }

func NewMapTimeValue(v map[string]time.Time) MapTimeValue {
	_ = "STUB: not implemented"
	return *new(MapTimeValue)
}

func (m MapTimeValue) Nil() bool                    { _ = "STUB: not implemented"; return false }
func (m MapTimeValue) Err() bool                    { _ = "STUB: not implemented"; return false }
func (m MapTimeValue) Type() ValueType              { _ = "STUB: not implemented"; return *new(ValueType) }
func (m MapTimeValue) Value() interface{}           { _ = "STUB: not implemented"; return nil }
func (m MapTimeValue) Val() map[string]time.Time    { _ = "STUB: not implemented"; return nil }
func (m MapTimeValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m MapTimeValue) ToString() string             { _ = "STUB: not implemented"; return "" }
func (m MapTimeValue) Len() int                     { _ = "STUB: not implemented"; return 0 }
func (m MapTimeValue) MapInt() map[string]int64     { _ = "STUB: not implemented"; return nil }

func (m MapTimeValue) MapValue() MapValue { _ = "STUB: not implemented"; return *new(MapValue) }

func (m MapTimeValue) Get(key string) (Value, bool) {
	_ = "STUB: not implemented"
	return *new(Value), false
}

func NewMapBoolValue(v map[string]bool) MapBoolValue {
	_ = "STUB: not implemented"
	return *new(MapBoolValue)
}

func (m MapBoolValue) Nil() bool                    { _ = "STUB: not implemented"; return false }
func (m MapBoolValue) Err() bool                    { _ = "STUB: not implemented"; return false }
func (m MapBoolValue) Type() ValueType              { _ = "STUB: not implemented"; return *new(ValueType) }
func (m MapBoolValue) Value() interface{}           { _ = "STUB: not implemented"; return nil }
func (m MapBoolValue) Val() map[string]bool         { _ = "STUB: not implemented"; return nil }
func (m MapBoolValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m MapBoolValue) ToString() string             { _ = "STUB: not implemented"; return "" }
func (m MapBoolValue) Len() int                     { _ = "STUB: not implemented"; return 0 }
func (m MapBoolValue) MapValue() MapValue           { _ = "STUB: not implemented"; return *new(MapValue) }

func (m MapBoolValue) Get(key string) (Value, bool) {
	_ = "STUB: not implemented"
	return *new(Value), false
}

func (m MapBoolValue) SliceValue() []Value { _ = "STUB: not implemented"; return nil }

func NewStructValue(v interface{}) StructValue { _ = "STUB: not implemented"; return *new(StructValue) }

func (m StructValue) Nil() bool                    { _ = "STUB: not implemented"; return false }
func (m StructValue) Err() bool                    { _ = "STUB: not implemented"; return false }
func (m StructValue) Type() ValueType              { _ = "STUB: not implemented"; return *new(ValueType) }
func (m StructValue) Value() interface{}           { _ = "STUB: not implemented"; return nil }
func (m StructValue) Val() interface{}             { _ = "STUB: not implemented"; return nil }
func (m StructValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m StructValue) ToString() string             { _ = "STUB: not implemented"; return "" }

func NewJsonValue(v json.RawMessage) JsonValue { _ = "STUB: not implemented"; return *new(JsonValue) }

func (m JsonValue) Nil() bool                    { _ = "STUB: not implemented"; return false }
func (m JsonValue) Err() bool                    { _ = "STUB: not implemented"; return false }
func (m JsonValue) Type() ValueType              { _ = "STUB: not implemented"; return *new(ValueType) }
func (m JsonValue) Value() interface{}           { _ = "STUB: not implemented"; return nil }
func (m JsonValue) Val() interface{}             { _ = "STUB: not implemented"; return nil }
func (m JsonValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m JsonValue) ToString() string             { _ = "STUB: not implemented"; return "" }

func NewTimeValue(v time.Time) TimeValue { _ = "STUB: not implemented"; return *new(TimeValue) }

func (m TimeValue) Nil() bool                    { _ = "STUB: not implemented"; return false }
func (m TimeValue) Err() bool                    { _ = "STUB: not implemented"; return false }
func (m TimeValue) Type() ValueType              { _ = "STUB: not implemented"; return *new(ValueType) }
func (m TimeValue) Value() interface{}           { _ = "STUB: not implemented"; return nil }
func (m TimeValue) Val() time.Time               { _ = "STUB: not implemented"; return *new(time.Time) }
func (m TimeValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m TimeValue) ToString() string             { _ = "STUB: not implemented"; return "" }
func (m TimeValue) Float() float64               { _ = "STUB: not implemented"; return 0 }
func (m TimeValue) Int() int64                   { _ = "STUB: not implemented"; return 0 }
func (m TimeValue) Time() time.Time              { _ = "STUB: not implemented"; return *new(time.Time) }

func NewErrorValue(v error) ErrorValue { _ = "STUB: not implemented"; return *new(ErrorValue) }

func NewErrorValuef(v string, args ...interface{}) ErrorValue {
	_ = "STUB: not implemented"
	return *new(ErrorValue)
}

func (m ErrorValue) Nil() bool                    { _ = "STUB: not implemented"; return false }
func (m ErrorValue) Err() bool                    { _ = "STUB: not implemented"; return false }
func (m ErrorValue) Type() ValueType              { _ = "STUB: not implemented"; return *new(ValueType) }
func (m ErrorValue) Value() interface{}           { _ = "STUB: not implemented"; return nil }
func (m ErrorValue) Val() error                   { _ = "STUB: not implemented"; return nil }
func (m ErrorValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m ErrorValue) ToString() string             { _ = "STUB: not implemented"; return "" }

// ErrorValues implement Go's error interface so they can easily cross the
// VM/Go boundary.
func (m ErrorValue) Error() string { _ = "STUB: not implemented"; return "" }

func NewNilValue() NilValue { _ = "STUB: not implemented"; return *new(NilValue) }

func (m NilValue) Nil() bool                    { _ = "STUB: not implemented"; return false }
func (m NilValue) Err() bool                    { _ = "STUB: not implemented"; return false }
func (m NilValue) Type() ValueType              { _ = "STUB: not implemented"; return *new(ValueType) }
func (m NilValue) Value() interface{}           { _ = "STUB: not implemented"; return nil }
func (m NilValue) Val() interface{}             { _ = "STUB: not implemented"; return nil }
func (m NilValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (m NilValue) ToString() string             { _ = "STUB: not implemented"; return "" }
