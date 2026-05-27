package datasource

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	u "github.com/araddon/gou"
)

var (
	// ErrNotDate an error for trying to corece/convert to Time a field that is not a time.
	ErrNotDate = errors.New("Unable to conver to time value")
)

// These are data-types that implement the database/sq interface for Scan() for custom types

type (
	// TimeValue Convert a string/bytes to time.Time by parsing the string
	// with a wide variety of different date formats that are supported
	// in http://godoc.org/github.com/araddon/dateparse
	TimeValue time.Time

	// StringArray Convert json to array of strings
	StringArray []string

	// JsonWrapper json data
	JsonWrapper json.RawMessage

	// JsonHelperScannable expects map json's (not array) map[string]interface
	JsonHelperScannable u.JsonHelper
)

func (m *TimeValue) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *TimeValue) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (m TimeValue) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

func (m *TimeValue) Time() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (m *TimeValue) Scan(src interface{}) error { _ = "STUB: not implemented"; return nil }

func (m *JsonWrapper) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"

	// UnmarshalJSON bytes into this typed struct
	return nil, nil
}

func (m *JsonWrapper) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// Value This is the go sql/driver interface we need to implement to allow
// conversion back forth
func (m JsonWrapper) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

func (m *JsonWrapper) Scan(src interface{}) error { _ = "STUB: not implemented"; return nil }

func (m *JsonWrapper) Unmarshal(v interface{}) error { _ = "STUB: not implemented"; return nil }

func (m *JsonHelperScannable) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalJSON bytes into this typed struct
func (m *JsonHelperScannable) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Value This is the go sql/driver interface we need to implement to allow
// conversion back forth
func (m JsonHelperScannable) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// Scan the database/sql interface for scanning sql byte vals into this
// typed structure.
func (m *JsonHelperScannable) Scan(src interface{}) error { _ = "STUB: not implemented"; return nil }

func (m *StringArray) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *StringArray) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// Value convert string to json values
func (m StringArray) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}

// Scan the database/sql interface for scanning sql byte vals into this
// typed structure.
func (m *StringArray) Scan(src interface{}) error { _ = "STUB: not implemented"; return nil }
