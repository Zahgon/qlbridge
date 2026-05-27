package exec

import (
	"database/sql/driver"
)

// Create a multiple error type
type errList []error

func (e *errList) append(err error) { _ = "STUB: not implemented"; return }

func (e errList) error() error { _ = "STUB: not implemented"; return nil }

func (e errList) Error() string { _ = "STUB: not implemented"; return "" }

func params(args []driver.Value) []interface{} { _ = "STUB: not implemented"; return nil }
