package plan

import (
	"math/rand"
	"time"

	"golang.org/x/net/context"

	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/rel"
	"github.com/araddon/qlbridge/schema"
)

// NextIdFunc is the id generation function to give statements
// their own id
type NextIdFunc func() uint64

// NextId is the global next id generation function
var NextId NextIdFunc

var rs = rand.New(rand.NewSource(time.Now().UnixNano()))

func init() {
	NextId = mathRandId
}

func mathRandId() uint64 { _ = "STUB: not implemented"; return 0 }

// Context for plan of a Relational task has info about the query
// projection, schema, function resolvers necessary to plan this statement.
// - may be transported across network boundaries to particpate in dag of tasks
// - holds references to in-mem data structures for schema
// - holds references to original statement
// - holds task specific state for errors, ids, etc (net.context)
// - manages Recover() - to persist/transport state
type Context struct {

	// Stateful Fields that are transported to participate across network/nodes
	context.Context                  // go context for cancelation in plan
	SchemaName      string           // schema name to load schema with
	id              uint64           // unique id per request
	fingerprint     uint64           // not unique per statement, used for getting prepared plans
	Raw             string           // Raw sql statement
	Stmt            rel.SqlStatement // Original Statement
	Projection      *Projection      // Projection for this context optional

	// Local in-memory helpers not transported across network
	Session expr.ContextReadWriter // Session for this connection
	Schema  *schema.Schema         // this schema for this connection
	Funcs   expr.FuncResolver      // Local/Dialect specific functions

	// From configuration
	DisableRecover bool

	// Local State
	Errors     []error
	errRecover interface{}
}

// NewContext plan context
func NewContext(query string) *Context { _ = "STUB: not implemented"; return nil }

func NewContextFromPb(pb *ContextPb) *Context { _ = "STUB: not implemented"; return nil }

// called by go routines/tasks to ensure any recovery panics are captured
func (m *Context) Recover() { _ = "STUB: not implemented"; return }

func (m *Context) init() {
	if m.id == 0 {
		if m.Schema != nil {
			m.SchemaName = m.Schema.Name
		}
		if ss, ok := m.Stmt.(*rel.SqlSelect); ok {
			m.fingerprint = uint64(ss.FingerPrintID())
		}
		m.id = NextId()
	}
}

// called by go routines/tasks to ensure any recovery panics are captured
func (m *Context) ToPB() *ContextPb { _ = "STUB: not implemented"; return nil }

func (m *Context) Equal(c *Context) bool { _ = "STUB: not implemented"; return false }
