package datasource

import (
	"database/sql/driver"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/plan"
)

const (
	// Default Max Allowed packets for connections
	MaxAllowedPacket = 4194304
)

// http://dev.mysql.com/doc/refman/5.6/en/server-system-variables.html
var mysqlGlobalVars *ContextSimple = NewMySqlGlobalVars()

func RowsForSession(ctx *plan.Context) [][]driver.Value { _ = "STUB: not implemented"; return nil }

func NewMySqlSessionVars() expr.ContextReadWriter {
	_ = "STUB: not implemented"
	return *new(expr.ContextReadWriter)
}

func NewMySqlGlobalVars() *ContextSimple { _ = "STUB: not implemented"; return nil }

//ctx.Data["@@session.auto_increment_increment"] = value.NewBoolValue(true)
