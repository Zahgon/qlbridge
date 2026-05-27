package exec

import (
	"github.com/araddon/qlbridge/plan"
)

var (
	// Ensure that we implement the Task Runner interface
	_ TaskRunner = (*Create)(nil)
	_ TaskRunner = (*Drop)(nil)
	_ TaskRunner = (*Alter)(nil)
)

type (
	// Create is executeable task for SQL Create, Alter, Schema, Source etc.
	Create struct {
		*TaskBase
		p *plan.Create
	}
	// Drop is executeable task for SQL DROP.
	Drop struct {
		*TaskBase
		p *plan.Drop
	}
	// Alter is executeable task for SQL ALTER.
	Alter struct {
		*TaskBase
		p *plan.Alter
	}
)

// NewCreate creates new create exec task
func NewCreate(ctx *plan.Context, p *plan.Create) *Create { _ = "STUB: not implemented"; return nil }

// Close Create
func (m *Create) Close() error { _ = "STUB: not implemented"; return nil }

// Run Create
func (m *Create) Run() error { _ = "STUB: not implemented"; return nil }

/*
	// "sub_schema_name" will create a new child schema called "sub_schema_name"
	// that is added to "existing_schema_name"
	// of source type elasticsearch
	CREATE source sub_schema_name WITH {
	  "type":"elasticsearch",
	  "schema":"existing_schema_name",
	  "settings" : {
	     "apikey":"GET_YOUR_API_KEY"
	  }
	};
*/
// If we specify a parent schema to add this child schema to

// NewDrop creates new drop exec task.
func NewDrop(ctx *plan.Context, p *plan.Drop) *Drop { _ = "STUB: not implemented"; return nil }

// Close Drop
func (m *Drop) Close() error { _ = "STUB: not implemented"; return nil }

// Run Drop
func (m *Drop) Run() error { _ = "STUB: not implemented"; return nil }

// NewAlter creates new ALTER exec task.
func NewAlter(ctx *plan.Context, p *plan.Alter) *Alter { _ = "STUB: not implemented"; return nil }

// Close Alter
func (m *Alter) Close() error { _ = "STUB: not implemented"; return nil }

// Run Alter
func (m *Alter) Run() error { _ = "STUB: not implemented"; return nil }
