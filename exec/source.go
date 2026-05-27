package exec

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/plan"
	"github.com/araddon/qlbridge/schema"
)

var (
	_ = u.EMPTY

	// Ensure that we implement the Task Runner interface
	// to ensure this can run in exec engine
	_ TaskRunner = (*Source)(nil)
)

// RequiresContext defines a Source which requires context.
type RequiresContext interface {
	SetContext(ctx *plan.Context)
}

// Source defines a datasource execution task.  It will Scan a data source for
// rows to feed into exec dag of tasks.  The source scanner uses iter.Next()
// messages.  The source may optionally allow Predicate PushDown, that is
// use the SQL select/where to filter rows so its not a real table scan. This
// interface is called ExecutorSource.
//
// Examples of Sources:
//  1. table      -- FROM table
//  2. channels   -- FROM stream
//  3. join       -- SELECT t1.name, t2.salary
//     FROM employee AS t1
//     INNER JOIN info AS t2
//     ON t1.name = t2.name;
//  4. sub-select -- SELECT * FROM (SELECT 1, 2, 3) AS t1;
type Source struct {
	*TaskBase
	p          *plan.Source
	Scanner    schema.ConnScanner
	ExecSource ExecutorSource
	JoinKey    KeyEvaluator
	closed     bool
}

// NewSource create a scanner to read from data source
func NewSource(ctx *plan.Context, p *plan.Source) (*Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Some sources require context so we seed it here

// NewSourceScanner A scanner to read from sub-query data source (join, sub-query, static)
func NewSourceScanner(ctx *plan.Context, p *plan.Source, scanner schema.ConnScanner) *Source {
	_ = "STUB: not implemented"
	return nil
}

func (m *Source) Copy() *Source { _ = "STUB: not implemented"; return nil }

func (m *Source) closeSource() error { _ = "STUB: not implemented"; return nil }

func (m *Source) Close() error { _ = "STUB: not implemented"; return nil }

// Still need to close base right?

func (m *Source) Run() error { _ = "STUB: not implemented"; return nil }

// continue
