package value

import (
	"fmt"
	"strings"
	"time"

	u "github.com/araddon/gou"
)

var (
	_ = u.EMPTY

	ErrConversion             = fmt.Errorf("Error converting type")
	ErrConversionNotSupported = fmt.Errorf("Unsupported conversion")
)

// ValueTypeFromString take a string value and infer valuetype
// Will infer based on the following rules:
// - If parseable as int, will be int
// - if not above, and parse bool, is bool
// - if not above, and parse float, float
// - if not above, and parse date, date
// - else string
func ValueTypeFromString(val string) ValueType { _ = "STUB: not implemented"; return *new(ValueType) }

// ValueTypeFromStringAll take a string value and infer valuetype
// adding the valid type JSON preferred over raw string.
//
// Will infer based on the following rules:
// - If parseable as int, will be int
// - if not above, and parse bool, is bool
// - if not above, and parse float, float
// - if not above, and parse date, date
// - if not above, and appears to be json (doesn't have to be valid)
// - else string
func ValueTypeFromStringAll(val string) ValueType {
	_ = "STUB: not implemented"
	return *new(ValueType)
}

// Cast a value to given value type
func Cast(valType ValueType, val Value) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

// Equal function compares equality after detecting type.
// error if it could not evaluate
func Equal(l, r Value) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ValueToString convert all scalar values to their go string.
func ValueToString(val Value) (string, bool) { _ = "STUB: not implemented"; return "", false }

// This is controversial, if we are demanding a "ToString"
// should we:
// 1)  take first?
// 2)  append comma separated?
// 3)  error?
//
// Our answer is that we are demanding a scalar string
// and going to take first.  calling function would have had to do type detection
// if they wanted something else.

// ValueToStrings convert all scalar values to their go []string.
func ValueToStrings(val Value) ([]string, bool) { _ = "STUB: not implemented"; return nil, false }

// is this boolean string?
func IsBool(sv string) bool { _ = "STUB: not implemented"; return false }

func BoolStringVal(sv string) bool { _ = "STUB: not implemented"; return false }

// ValueToBool Convert a value type to a bool if possible
func ValueToBool(val Value) (bool, bool) { _ = "STUB: not implemented"; return false, false }

// ValueToFloat64 Convert a value type to a float64 if possible
func ValueToFloat64(val Value) (float64, bool) { _ = "STUB: not implemented"; return 0, false }

// Should we co-erce bools to 0/1?

// ValueToInt Convert a value type to a int if possible
func ValueToInt(val Value) (int, bool) { _ = "STUB: not implemented"; return 0, false }

// ValueToInt64 Convert a value type to a int64 if possible
func ValueToInt64(val Value) (int64, bool) { _ = "STUB: not implemented"; return 0, false }

// StringToTimeAnchor Convert a string type to a time if possible.
// If "now-3d" then use date-anchoring ie if prefix = 'now'.
func StringToTimeAnchor(val string, anchor time.Time) (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

// Is date math

// ValueToTime Convert a value type to a time if possible
func ValueToTime(val Value) (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

// ValueToTimeAnchor given a value, and a time anchor, conver to time.
// use "now-3d" anchoring if has prefix "now".
func ValueToTimeAnchor(val Value, anchor time.Time) (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

//u.Warnf("un-handled type to time? %#v", val)

// StringToFloat64 converts a string to a float
// includes replacement of $ and other monetary format identifiers.
// May return math.NaN
func StringToFloat64(s string) (float64, bool) { _ = "STUB: not implemented"; return 0, false }

// Some strings we are trying to convert into Numbers are messy
// $3.12 etc, lets replace them and retry conversion again
var intStrReplacer = strings.NewReplacer("$", "", ",", "", "£", "", "€", "", " ", "")

func convertStringToInt64(depth int, s string) (int64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// So, we are going to TRUNCATE, ie round down

func marshalFloat(n float64) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
