package plan

var (
	// Ensure our default planner meets Planner interface.
	_ Planner = (*PlannerDefault)(nil)
)

// PlannerDefault is implementation of Planner that creates a dag of plan.Tasks
// that will be turned into execution plan by executor.  This is a simple
// planner but can be over-ridden by providing a Planner that will
// supercede any single or more visit methods.
// - stateful, specific to a single request
type PlannerDefault struct {
	Planner  Planner
	Ctx      *Context
	distinct bool
	children []Task
}

// NewPlanner creates a new default planner with context.
func NewPlanner(ctx *Context) *PlannerDefault { _ = "STUB: not implemented"; return nil }

// WalkPreparedStatement not implemented
func (m *PlannerDefault) WalkPreparedStatement(p *PreparedStatement) error {
	_ = "STUB: not implemented"
	return nil
}

// WalkCommand walks the command statement
func (m *PlannerDefault) WalkCommand(p *Command) error { _ = "STUB: not implemented"; return nil }

// WalkDrop walks the draop statement
func (m *PlannerDefault) WalkDrop(p *Drop) error { _ = "STUB: not implemented"; return nil }

// WalkCreate walk a Create Plan to create the dag of tasks for Create.
func (m *PlannerDefault) WalkCreate(p *Create) error { _ = "STUB: not implemented"; return nil }

// WalkAlter walk a ALTER Plan to create the dag of tasks forAlter.
func (m *PlannerDefault) WalkAlter(p *Alter) error { _ = "STUB: not implemented"; return nil }
