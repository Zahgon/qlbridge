// Package sqlite implements a Qlbridge Datasource interface around sqlite
// that translates mysql syntax to sqlite.
package sqlite

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/google/btree"
	"golang.org/x/net/context"

	// Import driver for sqlite
	_ "github.com/mattn/go-sqlite3"

	"github.com/araddon/qlbridge/exec"
	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/plan"
	"github.com/araddon/qlbridge/rel"
	"github.com/araddon/qlbridge/schema"
)

var (
	// ensure our conn implements connection features
	_ schema.ConnAll      = (*qryconn)(nil)
	_ schema.ConnMutation = (*qryconn)(nil)

	// SourcePlanner interface {
	// 	// given our request statement, turn that into a plan.Task.
	// 	WalkSourceSelect(pl Planner, s *Source) (Task, error)
	// }
	_ plan.SourcePlanner = (*qryconn)(nil)
)

type (
	// qryconn is a single-query connection in order to manage
	// stateful, non-multi-threaded access to sqlite rows object.
	qryconn struct {
		*exec.TaskBase
		stmt      rel.SqlStatement
		exit      <-chan bool
		source    *Source
		tbl       *schema.Table
		ps        *plan.Source
		indexCol  int
		rows      *sql.Rows
		ct        uint64
		cols      []string
		colidx    map[string]int
		err       error
		sqlInsert string
		sqlUpdate string
	}
)

func newQueryConn(tbl *schema.Table, source *Source) *qryconn {
	_ = "STUB: not implemented"
	return nil
}

func (m *qryconn) init() {
	cols := make([]string, len(m.cols))
	vals := make([]string, len(m.cols))
	for i, col := range m.cols {
		cols[i] = expr.IdentityMaybeQuote('"', col)
		vals[i] = "?"
	}
	m.sqlInsert = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", m.tbl.Name, strings.Join(cols, ", "), strings.Join(vals, ", "))
}

// Close the qryconn.  Since sqlite is a NON-threadsafe db, this is very important
// as we actually hold a lock per-table during scans to prevent conflict.
func (m *qryconn) Close() error { _ = "STUB: not implemented"; return nil }

// CreateIterator creates an interator to page through each row in this query resultset.
// This qryconn is wrapping a sql rows object, paging through until empty.
func (m *qryconn) CreateIterator() schema.Iterator {
	_ = "STUB: not implemented"

	// Columns gets the columns used in this query.
	return *new(schema.Iterator)
}

func (m *qryconn) Columns() []string {
	_ = "STUB: not implemented"

	// func (m *qryconn) Length() int                     { return 0 }
	return nil
}

//func (m *conn) SetColumns(cols []string)                  { m.tbl.SetColumns(cols) }

// CreateMutator part of Mutator interface to allow this connection to have access
// to the full plan context to take original sql statement and pass through to sqlite.
func (m *qryconn) CreateMutator(pc interface{}) (schema.ConnMutator, error) {
	_ = "STUB: not implemented"
	return *new(schema.ConnMutator), nil
}

func (m *qryconn) Next() schema.Message { _ = "STUB: not implemented"; return *new(schema.Message) }

//vals := make([]driver.Value, len(m.cols))
//u.Infof("expecting %d cols", len(m.cols))

//cols, _ := m.rows.Columns()
//u.Debugf("sqlite result cols provides %v but expecting %d", cols, len(m.cols))

//u.Debugf("read vals: %#v", writeCols)

// This seems pretty gross, isn't there a better way to do this?

//u.Debugf("%d %s  %T %v", i, m.cols[i], col, col)

//u.Infof("return item btreeP:%p itemP:%p cursorP:%p  %v %v", m, item, m.cursor, msg.Id(), msg.Values())
//u.Debugf("return? %T  %v", item, item.(*DriverItem).SqlDriverMessageMap)

// Put interface for Upsert.Put() to do single row insert based on key.
func (m *qryconn) Put(ctx context.Context, key schema.Key, row interface{}) (schema.Key, error) {
	_ = "STUB: not implemented"

	//u.Infof("%p Put(),  row:%#v", m, row)
	return *new(schema.Key), nil
}

//u.Debugf("empty, now do insert")
//sdm := datasource.NewSqlDriverMessageMap(id, rowVals, m.tbl.FieldPositions)

//u.Debugf("%p  PUT: id:%v IdVal:%v  Id():%v vals:%#v", m, id, sdm.IdVal, sdm.Id(), rowVals)

func (m *qryconn) PutMulti(ctx context.Context, keys []schema.Key, src interface{}) ([]schema.Key, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get a single row by key.
func (m *qryconn) Get(key driver.Value) (schema.Message, error) {
	_ = "STUB: not implemented"
	return *new(schema.Message), nil
}

// Delete deletes a single row by key
func (m *qryconn) Delete(key driver.Value) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// WalkSourceSelect An interface implemented by this connection allowing the planner
// to push down as much sql logic down to sqlite.
func (m *qryconn) WalkSourceSelect(planner plan.Planner, p *plan.Source) (plan.Task, error) {
	_ = "STUB: not implemented"
	return *new(plan.Task), nil
}

//p.Complete = true

// DeleteExpression Delete using a Where Expression
func (m *qryconn) DeleteExpression(p interface{}, where expr.Node) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

/*
	_, ok := p.(*plan.Delete)
	if !ok {
		return 0, plan.ErrNoPlan
	}

	deletedKeys := make([]*Key, 0)
	m.bt.Ascend(func(a btree.Item) bool {
		di, ok := a.(*DriverItem)
		if !ok {
			u.Warnf("wat?  %T   %#v", a, a)
			return false
		}
		msgCtx := di.SqlDriverMessageMap
		whereValue, ok := vm.Eval(msgCtx, where)
		if !ok {
			u.Debugf("could not evaluate where: %v", msgCtx.Values())
			//return deletedCt, fmt.Errorf("Could not evaluate where clause")
			return true
		}
		switch whereVal := whereValue.(type) {
		case value.BoolValue:
			if whereVal.Val() == false {
				//this means do NOT delete
			} else {
				// Delete!
				indexVal := msgCtx.Values()[m.indexCol]
				deletedKeys = append(deletedKeys, NewKey(makeId(indexVal)))
			}
		case nil:
			// ??
		default:
			if whereVal.Nil() {
				// Doesn't match, so don't delete
			} else {
				u.Warnf("unknown type? %T", whereVal)
			}
		}
		return true
	})

	for _, deleteKey := range deletedKeys {
		if ct, err := m.Delete(deleteKey); err != nil {
			u.Errorf("Could not delete key: %v", deleteKey)
		} else if ct != 1 {
			u.Errorf("delete should have removed 1 key %v", deleteKey)
		}
	}
	return len(deletedKeys), nil
*/

func MakeId(dv driver.Value) uint64 { _ = "STUB: not implemented"; return 0 }

//by := append(make([]byte,0,8), byte(r), byte(r>>8), byte(r>>16), byte(r>>24), byte(r>>32), byte(r>>40), byte(r>>48), byte(r>>56))

// Key implements Key and Sort interfaces.
type Key struct {
	Id uint64
}

func NewKey(key uint64) *Key             { _ = "STUB: not implemented"; return nil }
func (m *Key) Key() driver.Value         { _ = "STUB: not implemented"; return *new(driver.Value) }
func (m *Key) Less(than btree.Item) bool { _ = "STUB: not implemented"; return false }
