// Plan package converts the AST (expr package) into a plan, which is a DAG
// of tasks that comprise that plan, the planner is pluggable.
// The plan tasks are converted to executeable plan in exec.
package plan

import (
	"database/sql/driver"
	"fmt"

	u "github.com/araddon/gou"
	"github.com/golang/protobuf/proto"

	"github.com/araddon/qlbridge/rel"
	"github.com/araddon/qlbridge/schema"
)

var (
	// ErrNotImplemented is plan specific error for not implemented
	ErrNotImplemented = fmt.Errorf("QLBridge.plan: not implemented")
	// ErrNoDataSource no datasource/type found
	ErrNoDataSource = fmt.Errorf("QLBridge.plan: No datasource found")
	// ErrNoPlan no plan
	ErrNoPlan = fmt.Errorf("No Plan")

	// Ensure our tasks implement Task Interface
	_ Task = (*PreparedStatement)(nil)
	_ Task = (*Select)(nil)
	_ Task = (*Insert)(nil)
	_ Task = (*Upsert)(nil)
	_ Task = (*Update)(nil)
	_ Task = (*Delete)(nil)
	_ Task = (*Command)(nil)
	_ Task = (*Create)(nil)
	_ Task = (*Projection)(nil)
	_ Task = (*Source)(nil)
	_ Task = (*Into)(nil)
	_ Task = (*Where)(nil)
	_ Task = (*Having)(nil)
	_ Task = (*GroupBy)(nil)
	_ Task = (*Order)(nil)
	_ Task = (*JoinMerge)(nil)
	_ Task = (*JoinKey)(nil)

	// Force any plan that participates in a Select to implement Proto
	//  which allows us to serialize and distribute to multiple nodes.
	_ Proto = (*Select)(nil)
)

type (
	// SchemaLoader func interface for loading schema.
	SchemaLoader func(name string) (*schema.Schema, error)
	// SelectTask interface to check equality
	SelectTask interface {
		Equal(Task) bool
	}
	// Proto interface to ensure implements protobuf marshalling.
	Proto interface {
		proto.Marshaler
		proto.Unmarshaler
	}

	// Task interface allows different portions of distributed
	// plans (where, group-by, source-scan, project) to have
	// its own planner.  Output is a DAG of tasks to be given
	// to executor.
	// - may be parallel or sequential
	// - must be serializeable to participate in cross network tasks
	Task interface {
		// Walk, give a planner to this task to allow
		// Task to call appropriate parts of planner.
		Walk(p Planner) error

		// Children tasks of this, this task may be participating
		// in parents.
		Children() []Task
		// Add a child to this dag
		Add(Task) error
		IsSequential() bool
		SetSequential()
		IsParallel() bool
		SetParallel()
		Equal(Task) bool
		ToPb() (*PlanPb, error)
	}

	// Planner interface for planners.  Planners take a statement
	// and walk the statement to create a DAG of tasks representing
	// necessary sub-tasks to fulfil statement.
	// implementations of planners:
	// - qlbridge/exec package implements a non-distributed query-planner
	// - dataux/planner implements a distributed query-planner
	Planner interface {
		// DML Statements
		WalkSelect(p *Select) error
		WalkInsert(p *Insert) error
		WalkUpsert(p *Upsert) error
		WalkUpdate(p *Update) error
		WalkDelete(p *Delete) error
		WalkInto(p *Into) error
		WalkSourceSelect(p *Source) error
		WalkProjectionSource(p *Source) error
		WalkProjectionFinal(p *Select) error

		// Other Statements
		WalkPreparedStatement(p *PreparedStatement) error
		WalkCommand(p *Command) error

		// DDL operations
		WalkCreate(p *Create) error
		WalkDrop(p *Drop) error
		WalkAlter(p *Alter) error
	}

	// SourcePlanner Sources can often do their own planning for sub-select statements
	// ie mysql can do its own (select, projection) mongo, es can as well
	// - provide interface to allow passing down select planning to source
	SourcePlanner interface {
		// given our request statement, turn that into a plan.Task.
		WalkSourceSelect(pl Planner, s *Source) (Task, error)
	}
)

type (
	// PlanBase holds dag of child tasks
	PlanBase struct {
		parallel bool   // parallel or sequential?
		RootTask Task   // Root task
		tasks    []Task // Children tasks
	}
	// PreparedStatement plan
	PreparedStatement struct {
		*PlanBase
		Stmt *rel.PreparedStatement
	}
	// Select plan
	Select struct {
		*PlanBase
		Ctx      *Context
		From     []*Source
		Stmt     *rel.SqlSelect
		ChildDag bool
		pbplan   *PlanPb
	}
	// Insert plan
	Insert struct {
		*PlanBase
		Stmt   *rel.SqlInsert
		Source schema.ConnUpsert
	}
	// Upsert task (not official sql) for sql Upsert.
	Upsert struct {
		*PlanBase
		Stmt   *rel.SqlUpsert
		Source schema.ConnUpsert
	}
	// Update plan for sql Update statements.
	Update struct {
		*PlanBase
		Stmt   *rel.SqlUpdate
		Source schema.ConnUpsert
	}
	// Delete plan for sql DELETE where
	Delete struct {
		*PlanBase
		Stmt   *rel.SqlDelete
		Source schema.ConnDeletion
	}
	// Command for sql commands like SET.
	Command struct {
		*PlanBase
		Ctx  *Context
		Stmt *rel.SqlCommand
	}
	// Projection holds original query for column info and schema/field types
	Projection struct {
		*PlanBase
		Final bool // Is this final projection or not?
		P     *Select
		Stmt  *rel.SqlSelect
		Proj  *rel.Projection
	}
	// Source defines a source Within a Select query, it optionally has multiple
	// sources such as sub-select, join, etc this is the plan for a each source
	Source struct {
		*PlanBase
		pbplan *PlanPb

		// Request Information, if cross-node distributed query must be serialized
		*SourcePb
		Stmt     *rel.SqlSource  // The sub-query statement (may have been rewritten)
		Proj     *rel.Projection // projection for this sub-query
		ExecPlan Proto           // If SourceExec has a plan?
		Custom   u.JsonHelper    // Source specific context info

		// Schema and underlying Source provider info, not serialized or transported
		ctx        *Context       // query context, shared across all parts of this request
		DataSource schema.Source  // The data source for this From
		Conn       schema.Conn    // Connection for this source, only for this source/task
		Schema     *schema.Schema // Schema for this source/from
		Tbl        *schema.Table  // Table schema for this From
		Static     []driver.Value // this is static data source
		Cols       []string
	}
	// Into Select INTO table
	Into struct {
		*PlanBase
		Stmt *rel.SqlInto
	}
	// GroupBy clause plan
	GroupBy struct {
		*PlanBase
		Stmt    *rel.SqlSelect
		Partial bool
	}
	// Order By clause
	Order struct {
		*PlanBase
		Stmt *rel.SqlSelect
	}
	// Where pre-aggregation filter
	Where struct {
		*PlanBase
		Final bool
		Stmt  *rel.SqlSelect
	}
	// Having post-aggregation filter plan.
	Having struct {
		*PlanBase
		Stmt *rel.SqlSelect
	}
	// JoinMerge 2 source/input tasks for join
	JoinMerge struct {
		*PlanBase
		Left      Task
		Right     Task
		LeftFrom  *rel.SqlSource
		RightFrom *rel.SqlSource
		ColIndex  map[string]int
	}
	// JoinKey plan
	JoinKey struct {
		*PlanBase
		Source *Source
	}

	// DDL Tasks

	// Create plan for CREATE {SCHEMA|SOURCE|DATABASE}
	Create struct {
		*PlanBase
		Ctx  *Context
		Stmt *rel.SqlCreate
	}
	// Drop plan for DROP {SCHEMA|SOURCE|DATABASE}
	Drop struct {
		*PlanBase
		Ctx  *Context
		Stmt *rel.SqlDrop
	}
	// Alter plan for ALTER {TABLE|COLUMN}
	Alter struct {
		*PlanBase
		Ctx  *Context
		Stmt *rel.SqlAlter
	}
)

// WalkStmt Walk given statement for given Planner to produce a query plan
// which is a plan.Task and children, ie a DAG of tasks
func WalkStmt(ctx *Context, stmt rel.SqlStatement, planner Planner) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

// SelectPlanFromPbBytes Create a sql plan from pb.
func SelectPlanFromPbBytes(pb []byte, loader SchemaLoader) (*Select, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SelectTaskFromTaskPb create plan task for SqlSelect from pb.
func SelectTaskFromTaskPb(pb *PlanPb, ctx *Context, sel *rel.SqlSelect) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

func NewPlanBase(isParallel bool) *PlanBase { _ = "STUB: not implemented"; return nil }

func (m *PlanBase) Children() []Task    { _ = "STUB: not implemented"; return nil }
func (m *PlanBase) Add(task Task) error { _ = "STUB: not implemented"; return nil }

func (m *PlanBase) Close() error           { _ = "STUB: not implemented"; return nil }
func (m *PlanBase) Run() error             { _ = "STUB: not implemented"; return nil }
func (m *PlanBase) IsParallel() bool       { _ = "STUB: not implemented"; return false }
func (m *PlanBase) IsSequential() bool     { _ = "STUB: not implemented"; return false }
func (m *PlanBase) SetParallel()           { _ = "STUB: not implemented"; return }
func (m *PlanBase) SetSequential()         { _ = "STUB: not implemented"; return }
func (m *PlanBase) ToPb() (*PlanPb, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *PlanBase) Equal(t Task) bool          { _ = "STUB: not implemented"; return false }
func (m *PlanBase) EqualBase(p *PlanBase) bool { _ = "STUB: not implemented"; return false }

func (m *PlanBase) Walk(p Planner) error          { _ = "STUB: not implemented"; return nil }
func (m *Select) Walk(p Planner) error            { _ = "STUB: not implemented"; return nil }
func (m *PreparedStatement) Walk(p Planner) error { _ = "STUB: not implemented"; return nil }
func (m *Insert) Walk(p Planner) error            { _ = "STUB: not implemented"; return nil }
func (m *Upsert) Walk(p Planner) error            { _ = "STUB: not implemented"; return nil }
func (m *Update) Walk(p Planner) error            { _ = "STUB: not implemented"; return nil }
func (m *Delete) Walk(p Planner) error            { _ = "STUB: not implemented"; return nil }
func (m *Command) Walk(p Planner) error           { _ = "STUB: not implemented"; return nil }
func (m *Source) Walk(p Planner) error            { _ = "STUB: not implemented"; return nil }
func (m *Create) Walk(p Planner) error            { _ = "STUB: not implemented"; return nil }
func (m *Drop) Walk(p Planner) error              { _ = "STUB: not implemented"; return nil }
func (m *Alter) Walk(p Planner) error             { _ = "STUB: not implemented"; return nil }

// NewCreate creates a new Create Task plan.
func NewCreate(ctx *Context, stmt *rel.SqlCreate) *Create { _ = "STUB: not implemented"; return nil }

// NewDrop create Drop plan task.
func NewDrop(ctx *Context, stmt *rel.SqlDrop) *Drop { _ = "STUB: not implemented"; return nil }

// NewAlter create Alter plan task.
func NewAlter(ctx *Context, stmt *rel.SqlAlter) *Alter { _ = "STUB: not implemented"; return nil }

func (m *Select) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Select) MarshalTo(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (m *Select) Size() (n int) { _ = "STUB: not implemented"; return 0 }

func (m *Select) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

func (m *Select) serializeToPb() error { _ = "STUB: not implemented"; return nil }

//u.Infof("ctx %+v", m.pbplan.Select.Context)

func (m *Select) Equal(t Task) bool { _ = "STUB: not implemented"; return false }

//u.Warnf("Not Equal?   %T  vs %T", t, t2)
//u.Warnf("t!=t:   \n\t%#v \n\t%#v", t, t2)

func (m *Select) NeedsFinalProjection() bool { _ = "STUB: not implemented"; return false }

func (m *Select) IsSchemaQuery() bool {
	_ = "STUB: not implemented"
	// For Single Source statements, lets see if they are switching schema
	return false
}

//u.Debugf("schema:%q name:%q", m.From[0].Stmt.Schema, m.From[0].Stmt.Name)

func SelectFromPB(pb *PlanPb, loader SchemaLoader) (*Select, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//u.Infof("got context pb %+v", pb.Select.Context)

//u.Infof("%+v", pbt)

func SourceFromPB(pb *PlanPb, ctx *Context) (*Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//u.Debugf("custom %v", m.Custom)

// this is fine

//u.Warnf("hm  no conn, no stmt?....")
//return nil, ErrNoDataSource

// NewSource create a new plan Task for data source
func NewSource(ctx *Context, stmt *rel.SqlSource, isFinal bool) (*Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSourceStaticPlan(ctx *Context) *Source { _ = "STUB: not implemented"; return nil }

func (m *Source) Context() *Context { _ = "STUB: not implemented"; return nil }

func (m *Source) LoadConn() error {
	_ = "STUB: not implemented"

	// u.Debugf("LoadConn() nil?%v", m.Conn == nil)
	return nil
}

// Not all sources require a source, ie literal queries
// and some, information schema, or fully qualifyied schema queries
// requires schema switching

func (m *Source) IsSchemaQuery() bool { _ = "STUB: not implemented"; return false }

//u.Debugf("schema:%q name:%q", m.Stmt.Schema, m.Stmt.Name)

func (m *Source) ToPb() (*PlanPb, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Source) Equal(t Task) bool { _ = "STUB: not implemented"; return false }

func (m *Source) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Source) MarshalTo(data []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *Source) Size() (n int) { _ = "STUB: not implemented"; return 0 }

func (m *Source) Unmarshal(data []byte) error { _ = "STUB: not implemented"; return nil }

func (m *Source) serializeToPb() error { _ = "STUB: not implemented"; return nil }

func (m *Source) load() error {
	_ = "STUB: not implemented"
	// u.Debugf("source load schema=%s from=%s  %#v", m.ctx.Schema.Name, m.Stmt.SourceName(), m.Stmt)
	return nil
}

// u.Debugf("no schema found for %T  %q.%q ? err=%v", m.ctx.Schema, m.Stmt.Schema, fromName, err)

// Create a context-datasource

//u.Infof("schema=%s ds:%T  tbl:%v", m.Schema.Name, m.DataSource, tbl)

// Equal checks if two tasks are equal.
func (m *Projection) Equal(t Task) bool { _ = "STUB: not implemented"; return false }

// ToPb to protobuf.
func (m *Projection) ToPb() (*PlanPb, error) { _ = "STUB: not implemented"; return nil, nil }

// ProjectionFromPB create Projection from Protobuf.
func ProjectionFromPB(pb *PlanPb, sel *rel.SqlSelect) *Projection {
	_ = "STUB: not implemented"
	return nil
}

// NewJoinMerge A parallel join merge, uses Key() as value to merge
// two different input task/channels.
//
//	left source  ->
//	               \
//	                 --  join  -->
//	               /
//	right source ->
func NewJoinMerge(l, r Task, lf, rf *rel.SqlSource) *JoinMerge {
	_ = "STUB: not implemented"
	return nil
}

// Build an index of source to destination column indexing

//u.Debugf("left col:  idx=%d  key=%q as=%q col=%v parentidx=%v", len(m.colIndex), col.Key(), col.As, col.String(), col.ParentIndex)

//u.Debugf("left  colIndex:  %15q : idx:%d sidx:%d pidx:%d", m.leftStmt.Alias+"."+col.Key(), col.Index, col.SourceIndex, col.ParentIndex)

//u.Debugf("right col:  idx=%d  key=%q as=%q col=%v", len(m.colIndex), col.Key(), col.As, col.String())

//u.Debugf("right colIndex:  %15q : idx:%d sidx:%d pidx:%d", m.rightStmt.Alias+"."+col.Key(), col.Index, col.SourceIndex, col.ParentIndex)

// NewJoinKey creates JoinKey from Source.
func NewJoinKey(s *Source) *JoinKey { _ = "STUB: not implemented"; return nil }

// NewWhere new Where Task from SqlSelect statement.
func NewWhere(stmt *rel.SqlSelect) *Where { _ = "STUB: not implemented"; return nil }

// NewWhereFinal from SqlSelect statement.
func NewWhereFinal(stmt *rel.SqlSelect) *Where { _ = "STUB: not implemented"; return nil }

// NewHaving from SqlSelect statement.
func NewHaving(stmt *rel.SqlSelect) *Having { _ = "STUB: not implemented"; return nil }

// NewGroupBy from SqlSelect statement.
func NewGroupBy(stmt *rel.SqlSelect) *GroupBy { _ = "STUB: not implemented"; return nil }

// NewOrder from SqlSelect statement.
func NewOrder(stmt *rel.SqlSelect) *Order { _ = "STUB: not implemented"; return nil }

// Equal compares equality of two tasks.
func (m *Into) Equal(t Task) bool { _ = "STUB: not implemented"; return false }

func (m *Where) ToPb() (*PlanPb, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Where) Equal(t Task) bool { _ = "STUB: not implemented"; return false }

func WhereFromPB(pb *PlanPb) *Where { _ = "STUB: not implemented"; return nil }

func (m *Having) ToPb() (*PlanPb, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Having) Equal(t Task) bool { _ = "STUB: not implemented"; return false }

func HavingFromPB(pb *PlanPb) *Having { _ = "STUB: not implemented"; return nil }

func (m *GroupBy) ToPb() (*PlanPb, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *GroupBy) Equal(t Task) bool { _ = "STUB: not implemented"; return false }

func GroupByFromPB(pb *PlanPb) *GroupBy { _ = "STUB: not implemented"; return nil }

func (m *Order) ToPb() (*PlanPb, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Order) Equal(t Task) bool { _ = "STUB: not implemented"; return false }

func OrderFromPB(pb *PlanPb) *Order { _ = "STUB: not implemented"; return nil }

func (m *JoinMerge) Equal(t Task) bool { _ = "STUB: not implemented"; return false }

func (m *JoinKey) Equal(t Task) bool { _ = "STUB: not implemented"; return false }
