package exec

import (
	"github.com/araddon/qlbridge/plan"
)

var (
	// JobBuilder implements JobRunner
	_ JobRunner = (*JobExecutor)(nil)

	// Ensure that we implement the plan.Planner interface for our job
	_ Executor = (*JobExecutor)(nil)
	//_ plan.SourcePlanner = (*SourceBuilder)(nil)
)

// JobExecutor translates a Sql Statement into a Execution DAG of tasks
// using the Planner, Executor supplied.  This package implements default
// executor and uses the default Planner from plan.  This will create a single
// node dag of Tasks.
type JobExecutor struct {
	Planner  plan.Planner
	Executor Executor
	RootTask TaskRunner
	Ctx      *plan.Context
	distinct bool
	children []Task
}

// NewExecutor creates a new Job Executor.
func NewExecutor(ctx *plan.Context, planner plan.Planner) *JobExecutor {
	_ = "STUB: not implemented"
	return nil
}

// BuildSqlJob given a plan context (query statement, +context) create
// a JobExecutor and error if we can't.
func BuildSqlJob(ctx *plan.Context) (*JobExecutor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BuildSqlJobPlanned Create Job made up of sub-tasks in DAG that is the
// plan for execution of this query/job.
func BuildSqlJobPlanned(planner plan.Planner, executor Executor, ctx *plan.Context) (Task, error) {
	_ = "STUB: not implemented"

	//u.Debugf("build: %q", ctx.Raw)
	return *new(Task), nil
}

// NewTask create new task (from current context).
func (m *JobExecutor) NewTask(p plan.Task) Task { _ = "STUB: not implemented"; return *new(Task) }

// WalkPlan Main Entry point to take a Plan, and convert into Execution DAG
func (m *JobExecutor) WalkPlan(p plan.Task) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

// DDL

// WalkPreparedStatement not implemented
func (m *JobExecutor) WalkPreparedStatement(p *plan.PreparedStatement) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

// DML

// WalkSelect create dag of plan Select.
func (m *JobExecutor) WalkSelect(p *plan.Select) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

func (m *JobExecutor) WalkUpsert(p *plan.Upsert) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

func (m *JobExecutor) WalkInsert(p *plan.Insert) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

func (m *JobExecutor) WalkUpdate(p *plan.Update) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

func (m *JobExecutor) WalkDelete(p *plan.Delete) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

func (m *JobExecutor) WalkSource(p *plan.Source) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

//u.Debugf("setting p.Conn %p %T", p.Conn, p.Conn)

func (m *JobExecutor) WalkSourceExec(p *plan.Source) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

func (m *JobExecutor) WalkWhere(p *plan.Where) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

func (m *JobExecutor) WalkHaving(p *plan.Having) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

func (m *JobExecutor) WalkGroupBy(p *plan.GroupBy) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

func (m *JobExecutor) WalkOrder(p *plan.Order) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

func (m *JobExecutor) WalkProjection(p *plan.Projection) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

func (m *JobExecutor) WalkJoin(p *plan.JoinMerge) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

//u.Debugf("join.Left: %#v    \nright:%#v", p.Left, p.Right)

func (m *JobExecutor) WalkJoinKey(p *plan.JoinKey) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

func (m *JobExecutor) WalkPlanAll(p plan.Task) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

//u.Debugf("sequential?%v  parallel?%v", p.IsSequential(), p.IsParallel())

func (m *JobExecutor) WalkPlanTask(p plan.Task) (Task, error) {
	_ = "STUB: not implemented"
	//u.Debugf("WalkPlanTask: %p  %T", p, p)
	return *new(Task), nil
}

// Other Statements

// WalkCommand walk Commands such as SET.
func (m *JobExecutor) WalkCommand(p *plan.Command) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

// DDL Operations

// WalkCreate walks the Create plan.
func (m *JobExecutor) WalkCreate(p *plan.Create) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

// WalkDrop walks the Drop plan.
func (m *JobExecutor) WalkDrop(p *plan.Drop) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

// WalkAlter walks the Alter plan.
func (m *JobExecutor) WalkAlter(p *plan.Alter) (Task, error) {
	_ = "STUB: not implemented"
	return *new(Task), nil
}

// WalkChildren walk dag of plan tasks creating execution tasks
func (m *JobExecutor) WalkChildren(p plan.Task, root Task) error {
	_ = "STUB: not implemented"
	return nil
}

//u.Debugf("parent: %T  walk child %p %T  %#v", p, t, t, p.Children())

// Setup this dag of tasks
func (m *JobExecutor) Setup() error { _ = "STUB: not implemented"; return nil }

// Run this task
func (m *JobExecutor) Run() error { _ = "STUB: not implemented"; return nil }

// Close the normal close of root task
func (m *JobExecutor) Close() error { _ = "STUB: not implemented"; return nil }

// The drain is the last out channel, on last task
func (m *JobExecutor) DrainChan() MessageChan { _ = "STUB: not implemented"; return *new(MessageChan) }
