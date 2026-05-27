package exec

import (
	"database/sql/driver"
	"sync"

	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/schema"
)

var (
	// Ensure our driver implements appropriate database/sql interfaces
	_ driver.Conn    = (*qlbConn)(nil)
	_ driver.Driver  = (*qlbdriver)(nil)
	_ driver.Execer  = (*qlbConn)(nil)
	_ driver.Queryer = (*qlbConn)(nil)
	_ driver.Result  = (*qlbResult)(nil)
	_ driver.Rows    = (*qlbRows)(nil)
	_ driver.Stmt    = (*qlbStmt)(nil)
	//_ driver.Tx      = (*driverConn)(nil)

	// Create an instance of our driver
	qlbd          = &qlbdriver{}
	qlbDriverOnce sync.Once

	// Runtime Schema Config as in in-mem data structure of the
	//  datasources, tables, etc.   Sources must be registered
	//  as this is not persistent
	registry = schema.DefaultRegistry()

	_ = u.EMPTY
)

const (
	MysqlTimeFormat = "2006-01-02 15:04:05.000000000"
)

func RegisterSqlDriver() { _ = "STUB: not implemented"; return }

func DisableRecover() { _ = "STUB: not implemented"; return }

// sql.Driver Interface implementation.
//
// Notes about Value return types:
//
//	Value is a value that drivers must be able to handle.
//	It is either nil or an instance of one of these types:
//
//	  int64
//	  float64
//	  bool
//	  []byte
//	  string   [*] everywhere except from Rows.Next.
//	  time.Time
type qlbdriver struct{}

// Open returns a new connection to the database.
//
// Open may return a cached connection (one previously closed), but doing so
// is unnecessary; the sql package maintains a pool of idle connections for
// efficient re-use.
//
// The returned connection is only used by one goroutine at a time.
//
// @connInfo = database/Schema name
// @connInfo = driver-connection-info
// @connInfo = sourceType://source
func (m *qlbdriver) Open(connInfo string) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

// A stateful connection to database/source
//
// Execer is an optional interface that may be implemented by a Conn.
//
//	If a Conn does not implement Execer, the sql package's DB.Exec will
//	first prepare a query, execute the statement, and then close the
//	statement.
//
// Queryer is an optional interface that may be implemented by a Conn.
//
//	If a Conn does not implement Queryer, the sql package's DB.Query will
//	first prepare a query, execute the statement, and then close the
//	statement.
type qlbConn struct {
	parallel bool   // Do we Run In Background Mode?  Default = true
	connInfo string //
	schema   *schema.Schema
}

// Exec may return ErrSkip.
//
// Execer implementation. To be used for queries that do not return any rows
// such as Create Index, Insert, Upset, Delete etc
func (m *qlbConn) Exec(query string, args []driver.Value) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

// Queryer implementation
// Query may return ErrSkip
func (m *qlbConn) Query(query string, args []driver.Value) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

// Prepare returns a prepared statement, bound to this connection.
func (m *qlbConn) Prepare(query string) (driver.Stmt, error) {
	_ = "STUB: not implemented"
	return *new(driver.Stmt), nil
}

// Close invalidates and potentially stops any current
// prepared statements and transactions, marking this
// connection as no longer in use.
//
// Because the sql package maintains a free pool of
// connections and only calls Close when there's a surplus of
// idle connections, it shouldn't be necessary for drivers to
// do their own connection caching.
func (m *qlbConn) Close() error {
	_ = "STUB: not implemented"
	// u.Debugf("sqlbConn.Close() do we need to do anything here?")
	return nil
}

// Begin starts and returns a new transaction.
func (m *qlbConn) Begin() (driver.Tx, error) {
	_ = "STUB: not implemented"
	return *new(driver.Tx), nil
}

// sql.Tx Transaction Interface implementation.
type qlbTx struct{}

func (conn *qlbTx) Commit() error { _ = "STUB: not implemented"; return nil }

func (conn *qlbTx) Rollback() error { _ = "STUB: not implemented"; return nil }

// driver.Stmt Interface implementation.
//
// Stmt is a prepared statement. It is bound to a Conn and not
// used by multiple goroutines concurrently.
type qlbStmt struct {
	job   *JobExecutor
	query string
	conn  *qlbConn
}

// Close closes the statement.
//
// As of Go 1.1, a Stmt will not be closed if it's in use
// by any queries.
func (m *qlbStmt) Close() error { _ = "STUB: not implemented"; return nil }

// NumInput returns the number of placeholder parameters.
//
// If NumInput returns >= 0, the sql package will sanity check
// argument counts from callers and return errors to the caller
// before the statement's Exec or Query methods are called.
//
// NumInput may also return -1, if the driver doesn't know
// its number of placeholders. In that case, the sql package
// will not sanity check Exec or Query argument counts.
func (m *qlbStmt) NumInput() int {
	_ = "STUB: not implemented"

	// Exec executes a query that doesn't return rows, such
	// as an INSERT, UPDATE, DELETE
	return 0
}

func (m *qlbStmt) Exec(args []driver.Value) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

// Create a Job, which is Dag of Tasks that Run()

//u.Infof("in qlbdriver.Exec about to run")

//u.Debugf("After qlb driver.Run() in Exec()")

//resultWriter.ErrChan() <- err
//job.Close()

// Query executes a query that may return rows, such as a SELECT
func (m *qlbStmt) Query(args []driver.Value) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

// Create a Job, which is Dag of Tasks that Run()

// The only type of stmt that makes sense for Query is SELECT
//  and we need list of columns that requires casing

// Prepare a result writer, we manually append this task to end
// of job?

// TODO:   this can't run in parallel-buffered mode?
// how to open in go-routine and still be able to send error to rows?

//u.Debugf("Start Job.Run")

//u.Debugf("After job.Run()")

//resultWriter.ErrChan() <- err
//job.Close()

//u.Debugf("exiting Background Query")

// driver.ColumnConverter Interface implementation.
//
// ColumnConverter may be optionally implemented by driver.Stmt if the
// statement is aware of its own columns' types and can convert from
// any type to a driver Value.
//
// ColumnConverter returns a ValueConverter for the provided
// column index.  If the type of a specific column isn't known
// or shouldn't be handled specially, DefaultValueConverter
// can be returned.
func (conn *qlbStmt) ColumnConverter(idx int) driver.ValueConverter {
	_ = "STUB: not implemented"

	// driver.Rows Interface implementation.
	//
	// Rows is an iterator over an executed query's results.
	return *new(driver.ValueConverter)
}

type qlbRows struct{}

// Columns returns the names of the columns. The number of
// columns of the result is inferred from the length of the
// slice.  If a particular column name isn't known, an empty
// string should be returned for that entry.
func (conn *qlbRows) Columns() []string {
	_ = "STUB: not implemented"

	// Close closes the rows iterator.
	return nil
}

func (conn *qlbRows) Close() error { _ = "STUB: not implemented"; return nil }

// Next is called to populate the next row of data into
// the provided slice. The provided slice will be the same
// size as the Columns() are wide.
//
// The dest slice may be populated only with
// a driver Value type, but excluding string.
// All string values must be converted to []byte.
//
// Next should return io.EOF when there are no more rows.
func (conn *qlbRows) Next(dest []driver.Value) error { _ = "STUB: not implemented"; return nil }

// driver.Result Interface implementation.
//
// Result is the result of a query execution that doesn't return rows
type qlbResult struct {
	lastId   int64
	affected int64
	err      error
}

// LastInsertId returns the database's auto-generated ID
// after, for example, an INSERT into a table with primary
// key.
func (r *qlbResult) LastInsertId() (int64, error) {
	_ = "STUB: not implemented"
	return 0,

		// RowsAffected returns the number of rows affected by the
		// query.
		nil
}

func (r *qlbResult) RowsAffected() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func join(a []string) string { _ = "STUB: not implemented"; return "" }

func queryArgsConvert(query string, args []driver.Value) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// a tiny, tiny, tiny bit of string sanitization

func escapeString(txt string) string { _ = "STUB: not implemented"; return "" }

func escapeQuotes(txt string) string { _ = "STUB: not implemented"; return "" }
