// Package testutil a Test only package for harness to load, implement SQL tests.
package testutil

import (
	"database/sql/driver"

	// side-effect import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	verbose *bool
)

// TestingT is an interface wrapper around *testing.T so when we import
// this go dep, govendor don't import "testing"
type TestingT interface {
	Errorf(format string, args ...interface{})
}

// Setup enables -vv verbose logging or sends logs to /dev/null
// env var VERBOSELOGS=true was added to support verbose logging with alltests
func Setup() {
	_ = "STUB: not implemented"

	// QuerySpec a test harness
	return
}

type QuerySpec struct {
	Source          string // Db source
	Sql             string
	Exec            string
	HasErr          bool
	Cols            []string
	ValidateRow     func([]interface{})
	ExpectRowCt     int
	ExpectColCt     int
	RowData         interface{}
	Expect          [][]driver.Value
	ValidateRowData func()
}

// ExecSpec execute a queryspec test
func ExecSpec(t TestingT, q *QuerySpec) { _ = "STUB: not implemented"; return }

// ExecSqlSpec execute a test harness of QuerySpec
func ExecSqlSpec(t TestingT, q *QuerySpec) { _ = "STUB: not implemented"; return }

// rowVals is an []interface{} of all of the column results

// rowVals is an []interface{} of all of the column results

func TestSelect(t TestingT, sql string, expects [][]driver.Value) {
	_ = "STUB: not implemented"
	return
}

func TestExec(t TestingT, sql string) { _ = "STUB: not implemented"; return }

// TestSqlSelect tests using the database/sql driver
func TestSqlSelect(t TestingT, source, sql string, expects [][]driver.Value) {
	_ = "STUB: not implemented"
	return
}

func TestSelectErr(t TestingT, sql string, expects [][]driver.Value) {
	_ = "STUB: not implemented"
	return
}
