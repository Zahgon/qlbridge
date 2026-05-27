package plan

import (
	"github.com/araddon/qlbridge/rel"
)

// A static projection has already had its column/types defined
//
//	and doesn't need to use internal schema to find it, often internal SHOW/DESCRIBE
func NewProjectionStatic(proj *rel.Projection) *Projection { _ = "STUB: not implemented"; return nil }

// Final Projections project final select columns for result-writing
func NewProjectionFinal(ctx *Context, p *Select) (*Projection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewProjectionInProcess(stmt *rel.SqlSelect) *Projection { _ = "STUB: not implemented"; return nil }

func (m *Projection) loadLiteralProjection(ctx *Context) error {
	_ = "STUB: not implemented"

	//u.Debugf("creating plan.Projection literal %s", ctx.Stmt.String())
	return nil
}

//u.Debugf("col.As=%q  col.Expr %#v", col.As, col.Expr)

// number?

//u.Infof("number? %#v", et)

//u.Infof("type? %#v", et)

func (m *Projection) loadFinal(ctx *Context, isFinal bool) error {
	_ = "STUB: not implemented"

	//u.Debugf("creating plan.Projection final %s", m.Stmt.String())
	return nil
}

//u.Debugf("getting cols? %v   cols=%v", from.ColumnPositions())

//_, right, _ := col.LeftRight()
//u.Infof("col %s", col)

//u.Debugf("in plan final %s", col.As)

//u.Debugf("not final %s", col.As)

//u.Debugf("projection: %p add col: %v %v", m.Proj, col.As, schemaCol.Type.String())

//u.Infof("schema col not found: final?%v col: %#v InFinal?%v", isFinal, col, col.InFinalProjection())

func projectionForSourcePlan(plan *Source) error { _ = "STUB: not implemented"; return nil }

// u.Debugf("created plan.Proj  *rel.Projection %p", plan.Proj)
// Not all Execution run-times support schema.  ie, csv files and other "ad-hoc" structures
// do not have to have pre-defined data in advance, in which case the schema output
// will not be deterministic on the sql []driver.values

//u.Debugf("col: %v  star?%v", col, col.Star)

//u.Infof("col add %v for %s", schemaCol.Type.String(), col)

//u.Infof("not in final? %#v", col)

//u.Debugf("projection: %p add col: %v %v", plan.Proj, col.As, schemaCol.Type.String())

//u.Infof("star cols? %v fields: %v", plan.Tbl.FieldPositions, plan.Tbl.Fields)

//u.Infof("  add col %v  %+v", f.Name, f)

//u.Warnf("count(*) as=%v", col.As)

// A column was included in projection that does not exist in source.
// TODO:  Should we allow sources to have settings that specify wether
//  we enforce schema validation on parse?  or on execution?  many no-sql stores
//  this is fine

// Probably not string?

//u.Infof("plan.Projection %p  cols: %d", plan.Proj, len(plan.Proj.Columns))
