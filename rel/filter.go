package rel

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/expr"
)

var (
	_ = u.EMPTY

	// Ensure each Filter statement implement's Filter interface
	_ Filter = (*FilterStatement)(nil)
	_ Filter = (*FilterSelect)(nil)
	// Statements with Columns
	_ ColumnsStatement = (*FilterSelect)(nil)
)

type (
	// Filter interface for Filter Statements (either Filter/FilterSelect)
	Filter interface {
		String() string
	}
	// FilterSelect is a Filter but also has projected columns
	FilterSelect struct {
		*FilterStatement
		Columns Columns
	}
	// FilterStatement is a statement of type = Filter
	FilterStatement struct {
		checkedIncludes bool
		includes        []string
		Description     string       // initial pre-start comments
		Raw             string       // full original raw statement
		Filter          expr.Node    // FILTER <filter_expr>
		Where           expr.Node    // WHERE <expr> [AND <expr>] syntax
		OrderBy         Columns      // order by
		From            string       // From is optional
		Limit           int          // Limit
		Alias           string       // Non-Standard sql, alias/name of sql another way of expression Prepared Statement
		With            u.JsonHelper // Non-Standard SQL for properties/config info, similar to Cassandra with, purse json
	}
)

// NewFilterStatement Create A FilterStatement
func NewFilterStatement() *FilterStatement { _ = "STUB: not implemented"; return nil }

func (m *FilterStatement) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }

// String representation of FilterStatement
func (m *FilterStatement) String() string { _ = "STUB: not implemented"; return "" }

// FingerPrint consistent hashed int value of FingerPrint above
func (m *FilterStatement) FingerPrintID() int64 { _ = "STUB: not implemented"; return 0 }

// Includes Recurse this statement and find all includes
func (m *FilterStatement) Includes() []string { _ = "STUB: not implemented"; return nil }

func (m *FilterStatement) Equal(s *FilterStatement) bool { _ = "STUB: not implemented"; return false }

func NewFilterSelect() *FilterSelect { _ = "STUB: not implemented"; return nil }

func (m *FilterSelect) AddColumn(colArg Column) error { _ = "STUB: not implemented"; return nil }

func (m *FilterSelect) WriteDialect(w expr.DialectWriter) { _ = "STUB: not implemented"; return }

// String representation of FilterSelect
func (m *FilterSelect) String() string { _ = "STUB: not implemented"; return "" }

// FingerPrint consistent hashed int value of FingerPrint above
func (m *FilterSelect) FingerPrintID() int64 { _ = "STUB: not implemented"; return 0 }

// Equal Checks for deep equality
func (m *FilterSelect) Equal(s *FilterSelect) bool { _ = "STUB: not implemented"; return false }
