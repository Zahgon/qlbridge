package plan

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/schema"
)

var (
	_ = u.EMPTY
)

func (m *PlannerDefault) WalkInto(p *Into) error { _ = "STUB: not implemented"; return nil }

func upsertSource(ctx *Context, table string) (schema.ConnUpsert, error) {
	_ = "STUB: not implemented"
	return *new(schema.ConnUpsert), nil
}

//return nil, err

func (m *PlannerDefault) WalkInsert(p *Insert) error { _ = "STUB: not implemented"; return nil }

func (m *PlannerDefault) WalkUpdate(p *Update) error { _ = "STUB: not implemented"; return nil }

func (m *PlannerDefault) WalkUpsert(p *Upsert) error { _ = "STUB: not implemented"; return nil }

func (m *PlannerDefault) WalkDelete(p *Delete) error { _ = "STUB: not implemented"; return nil }

//return nil, err
