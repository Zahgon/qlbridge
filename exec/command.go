package exec

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/plan"
	"github.com/araddon/qlbridge/rel"
)

var (
	_ = u.EMPTY

	// Ensure that we implement the Task Runner interface
	_ TaskRunner = (*Command)(nil)
)

// Command is executeable task for SET SQL commands
type Command struct {
	*TaskBase
	p *plan.Command
}

// NewCommand creates new command exec task
func NewCommand(ctx *plan.Context, p *plan.Command) *Command { _ = "STUB: not implemented"; return nil }

// Close Command
func (m *Command) Close() error { _ = "STUB: not implemented"; return nil }

// Run Command
func (m *Command) Run() error {
	_ = "STUB: not implemented"
	// defer m.Ctx.Recover()
	return nil
}

func (m *Command) runSet() error { _ = "STUB: not implemented"; return nil }

//u.Debugf("running set? %v", m.p.Stmt.String())

// for k, v := range m.Ctx.Session.Row() {
// 	u.Infof("%p session? %s: %v", m.Ctx.Session, k, v.Value())
// }

func evalSetExpression(col *rel.CommandColumn, ctx expr.ContextReadWriter, arg expr.Node) error {
	_ = "STUB: not implemented"
	return nil
}

//u.Infof(`writeContext.Put("%v",%v)`, col.Key(), rhv.Value())

// Special statements

// http://dev.mysql.com/doc/refman/5.7/en/charset-connection.html
// hm, no idea what to do
/*
	SET character_set_client = charset_name;
	SET character_set_results = charset_name;
	SET character_set_connection = charset_name;
*/
