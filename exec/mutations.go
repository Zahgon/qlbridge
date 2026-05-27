package exec

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/plan"
	"github.com/araddon/qlbridge/rel"
	"github.com/araddon/qlbridge/schema"
)

var (
	_ = u.EMPTY

	_ TaskRunner = (*Upsert)(nil)
	_ TaskRunner = (*DeletionTask)(nil)
	_ TaskRunner = (*DeletionScanner)(nil)
)

type (
	// Upsert task for insert, update, upsert
	Upsert struct {
		*TaskBase
		closed  bool
		insert  *rel.SqlInsert
		update  *rel.SqlUpdate
		upsert  *rel.SqlUpsert
		db      schema.ConnUpsert
		dbpatch schema.ConnPatchWhere
	}
	// Delete task for sources that natively support delete
	DeletionTask struct {
		*TaskBase
		closed  bool
		p       *plan.Delete
		sql     *rel.SqlDelete
		db      schema.ConnDeletion
		deleted int
	}
	// Delete scanner if we don't have a seek operation on this source
	DeletionScanner struct {
		*DeletionTask
	}
)

// An insert to write to data source
func NewInsert(ctx *plan.Context, p *plan.Insert) *Upsert { _ = "STUB: not implemented"; return nil }

func NewUpdate(ctx *plan.Context, p *plan.Update) *Upsert { _ = "STUB: not implemented"; return nil }

func NewUpsert(ctx *plan.Context, p *plan.Upsert) *Upsert { _ = "STUB: not implemented"; return nil }

// An inserter to write to data source
func NewDelete(ctx *plan.Context, p *plan.Delete) *DeletionTask {
	_ = "STUB: not implemented"
	return nil
}

func (m *Upsert) Close() error { _ = "STUB: not implemented"; return nil }

func (m *Upsert) Run() error { _ = "STUB: not implemented"; return nil }

// status?

func (m *Upsert) updateValues() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// fall through

// TODO: qlbridge#13  Need a way of expressing which layer (here, db) this expr should run in?
//  - ie, run in backend datasource?   or here?  translate the expr to native language

//u.Debugf("key:%v col: %v   vals:%v", key, valcol, valmap[key])

// if our backend source supports Where-Patches, ie update multiple

// TODO:   If it does not implement Where Patch then we need to do a poly fill
//      Do we have to recognize if the Where is on a primary key?
// - for sources/queries that can't do partial updates we need to do a read first

// Create a key from Where

func (m *Upsert) insertRows(rows [][]*rel.ValueColumn) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *DeletionTask) Close() error { _ = "STUB: not implemented"; return nil }

func (m *DeletionTask) Run() error { _ = "STUB: not implemented"; return nil }

func (m *DeletionScanner) Run() error { _ = "STUB: not implemented"; return nil }
