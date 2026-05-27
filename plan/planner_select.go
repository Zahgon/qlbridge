package plan

import (
	"github.com/araddon/qlbridge/rel"
	"github.com/araddon/qlbridge/schema"
)

func needsFinalProjection(s *rel.SqlSelect) bool { _ = "STUB: not implemented"; return false }

// Where?

// WalkSelect walk a select statement filling out plan.
func (m *PlannerDefault) WalkSelect(p *Select) error {
	_ = "STUB: not implemented"

	// u.Debugf("VisitSelect ctx:%p  %+v", p.Ctx, p.Stmt)
	return nil
}

// TODO:   move to a Finalize() in query parser/planner

// Need to rewrite the From statement to ensure all fields necessary to support
//  joins, wheres, etc exist but is standalone query

// now fold into previous task

// fold this source into previous

//u.Debugf("got task: %T", lastSource)

// SELECT id from article WHERE id in (select article_id from comments where comment_ct > 50);

//u.Debugf("Adding aggregate/group by? %#v", m.Planner)

//u.Infof("Projection:  %T:%p   %T:%p", proj, proj, proj.Proj, proj.Proj)

//u.Debugf("m.Ctx: %p m.Ctx.Projection:    %T:%p", m.Ctx, m.Ctx.Projection, m.Ctx.Projection)

// WalkProjectionFinal walk the select plan to create final projection.
func (m *PlannerDefault) WalkProjectionFinal(p *Select) error {
	_ = "STUB: not implemented"
	// Add a Final Projection to choose the columns for results
	return nil
}

//u.Infof("Projection:  %T:%p   %T:%p", proj, proj, proj.Proj, proj.Proj)

// Not entirely sure we should be over-writing the projection?

// Build Column Name to Position index for given *source* (from) used to interpret
// positional []driver.Value args, mutate the *from* itself to hold this map
func buildColIndex(colSchema schema.ConnColumns, p *Source) error {
	_ = "STUB: not implemented"
	return nil
}

// WalkSourceSelect is a single source select
func (m *PlannerDefault) WalkSourceSelect(p *Source) error { _ = "STUB: not implemented"; return nil }

//u.Debugf("%p VisitSubselect from.source = %q", p, p.Stmt.Source)

//u.Debugf("%p VisitSubselect from=%q", p, p)

// All of this is plan info, ie needs JoinKey

// We need to build a ColIndex of source column/select/projection column
//u.Debugf("datasource? %#v", p.Conn)

// this is fine

// Can do our own planning

// Add a Non-Final Projection to choose the columns for results

// WalkProjectionSource non final projection (ie, per from).
func (m *PlannerDefault) WalkProjectionSource(p *Source) error {
	_ = "STUB: not implemented"
	// Add a Non-Final Projection to choose the columns for results
	//u.Debugf("exec.projection: %p job.proj: %p added  %s", p, m.Ctx.Projection, p.Stmt.String())
	return nil
}

//u.Debugf("source projection: %p added  %s", proj, p.Stmt.Source.String())

// WalkLiteralQuery Handle Literal queries such as "SELECT 1, @var;"
func (m *PlannerDefault) WalkLiteralQuery(p *Select) error {
	_ = "STUB: not implemented"
	// u.Debugf("WalkLiteralQuery %+v", p.Stmt)
	// Must project and possibly where
	return nil
}

// the reason this is wrong is that the Source task gets
// added in the WalkProjectionFinal below and the Where would need to be in the
// middle of the Source -> Where -> Projection tasks

//u.Debugf("m.Ctx: %p  m.Ctx.Projection.Proj:%p ", m.Ctx, m.Ctx.Projection.Proj)
