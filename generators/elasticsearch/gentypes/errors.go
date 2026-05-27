package gentypes

// MissingFieldErrors are returned when a segment can't be evaluated due to a
// referenced field missing from a schema.
type MissingFieldError struct {
	Field string
}

// MissingField creates a new MissingFieldError for the given field.
func MissingField(field string) *MissingFieldError { _ = "STUB: not implemented"; return nil }

func (m *MissingFieldError) Reason() string { _ = "STUB: not implemented"; return "" }
func (m *MissingFieldError) Status() int    { _ = "STUB: not implemented"; return 0 }

func (m *MissingFieldError) Error() string { _ = "STUB: not implemented"; return "" }
