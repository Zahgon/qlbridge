// Package mockcsvtestdata is csv test data only used for tests.
package mockcsvtestdata

import (
	"sync"

	"github.com/araddon/qlbridge/plan"
	"github.com/araddon/qlbridge/schema"
)

var (
	loadData   sync.Once
	MockSchema *schema.Schema

	// TestContext is a function to create plan.Context for a given test context.
	TestContext func(query string) *plan.Context
)

func SetContextToMockCsv() { _ = "STUB: not implemented"; return }

func SchemaLoader(name string) (*schema.Schema, error) { _ = "STUB: not implemented"; return nil, nil }

func LoadTestDataOnce() { _ = "STUB: not implemented"; return }

// Load in a "csv file" into our mock data store

//reg.RefreshSchema(mockcsv.SchemaName)
