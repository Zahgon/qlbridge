package vm

import (
	"fmt"

	"github.com/araddon/qlbridge/datasource"
	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/lex"
)

var (
	SqlEvalError = fmt.Errorf("Could not evaluate sql statement")
)

// SqlVm vm is a vm for parsing, evaluating a
type SqlVm struct {
	Statement expr.SqlStatement
	Keyword   lex.TokenType
	sel       *expr.SqlSelect
	ins       *expr.SqlInsert
	del       *expr.SqlDelete
}

// SqlVm parsers a sql query into columns, where guards, etc
func NewSqlVm(sqlText string) (*SqlVm, error) { _ = "STUB: not implemented"; return nil, nil }

// Execute applies a parse expression to the specified context's
//
//	writeContext in the case of sql query is similar to a recordset for selects,
//	  or for delete, insert, update it is like the storage layer
func (m *SqlVm) Execute(writeContext datasource.ContextWriter, readContext datasource.ContextReader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Execute applies a dml sql select expression to the specified context's
//
//	writeContext in the case of sql query is similar to a recordset for selects,
//	  or for delete, insert, update it is like the storage layer
func (m *SqlVm) ExecuteSelect(writeContext datasource.ContextWriter, readContext datasource.ContextReader) (err error) {
	_ = "STUB: not implemented"
	//defer errRecover(&err)
	return nil
}

// Check and see if we are where Guarded

//u.Debugf("Has a Where:  %v", m.Request.Where.Root.StringAST())

//u.Debugf("Matched where: %v", whereValue)

// TODO:  evaluate if guard

//u.Debugf("tree.Root: as?%v %#v", col.As, col.Tree.Root)

//writeContext.Put()

func (m *SqlVm) ExecuteInsert(writeContext datasource.RowWriter) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//u.Debugf("tree.Root: i, as, val:  %v %v %v", i, col.As, row[i])

//v, ok := s.Walk(col.Tree.Root)

func (m *SqlVm) ExecuteDelete(writeContext datasource.ContextWriter, readContext datasource.ContextReader) (err error) {
	_ = "STUB: not implemented"
	//defer errRecover(&err)
	return nil
}

// Check and see if we are where Guarded

// //u.Debugf("tree.Root: as?%v %#v", col.As, col.Tree.Root)
// v, ok := s.Walk(col.Tree.Root)
// if ok {
// 	writeContext.Put(col, readContext, v)
// }
