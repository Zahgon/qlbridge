package exec

import (
	"database/sql/driver"

	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/plan"
	"github.com/araddon/qlbridge/schema"
)

const (
	MaxAllowedPacket = 1024 * 1024
)

var (
	_ = u.EMPTY

	// ensure our resultwriter implements database/sql/driver `driver.Rows`
	_ driver.Rows = (*ResultWriter)(nil)

	// Ensure that we implement the Task Runner interface
	// required for usage as tasks in Executor
	_ TaskRunner = (*ResultExecWriter)(nil)
	_ TaskRunner = (*ResultWriter)(nil)
	_ TaskRunner = (*ResultBuffer)(nil)
)

type (
	// ResultExecWriter for writing tasks results
	ResultExecWriter struct {
		*TaskBase
		closed       bool
		err          error
		rowsAffected int64
		lastInsertID int64
	}
	// ResultWriter for writing tasks results
	ResultWriter struct {
		*TaskBase
		closed bool
		cols   []string
	}
	// ResultBuffer for writing tasks results
	ResultBuffer struct {
		*TaskBase
		closed bool
		cols   []string
	}
)

// NewResultExecWriter a result writer for exect task
func NewResultExecWriter(ctx *plan.Context) *ResultExecWriter {
	_ = "STUB: not implemented"
	return nil
}

// Signal to quit

// NewResultWriter for a plan
func NewResultWriter(ctx *plan.Context) *ResultWriter { _ = "STUB: not implemented"; return nil }

// NewResultRows a resultwriter
func NewResultRows(ctx *plan.Context, cols []string) *ResultWriter {
	_ = "STUB: not implemented"
	return nil
}

// NewResultBuffer create a result buffer to write temp tasks into results.
func NewResultBuffer(ctx *plan.Context, writeTo *[]schema.Message) *ResultBuffer {
	_ = "STUB: not implemented"
	return nil
}

//u.Infof("write to msgs: %v", len(*writeTo))

// Result of exec task
func (m *ResultExecWriter) Result() driver.Result {
	_ = "STUB: not implemented"
	return *new(driver.Result)
}

// Copy exec task
func (m *ResultExecWriter) Copy() *ResultExecWriter { _ = "STUB: not implemented"; return nil }

// Close exect task
func (m *ResultExecWriter) Close() error {
	_ = "STUB: not implemented"
	// u.Debugf("%p ResultExecWriter.Close()???? already closed?%v", m, m.closed)
	return nil
}

// Copy result writter
func (m *ResultWriter) Copy() *ResultWriter { _ = "STUB: not implemented"; return nil }

// Close ResultWriter
func (m *ResultWriter) Close() error {
	_ = "STUB: not implemented"
	// u.Debugf("%p ResultWriter.Close()???? already closed?%v", m, m.closed)
	return nil
}

// Copy the result buffer
func (m *ResultBuffer) Copy() *ResultBuffer { _ = "STUB: not implemented"; return nil }

// Close the ResultBuffer
func (m *ResultBuffer) Close() error {
	_ = "STUB: not implemented"
	// u.Debugf("%p ResultBuffer.Close()???? already closed?%v", m, m.closed)
	return nil
}

// Next his is implementation of the sql/driver Rows() Next() interface
func (m *ResultWriter) Next(dest []driver.Value) error { _ = "STUB: not implemented"; return nil }

// Run For ResultWriter, since we are are not paging through messages
// using this mesage channel, instead using Next() as defined by sql/driver
// we don't read the input channel, just watch stop channels
func (m *ResultWriter) Run() error { _ = "STUB: not implemented"; return nil }

// closing output channels is the signal to stop

// u.Debugf("%p got resultwriter.Run() sigquit?", m)

// Columns list of column names
func (m *ResultWriter) Columns() []string { _ = "STUB: not implemented"; return nil }

func resultWrite(m *ResultWriter) MessageHandler {
	_ = "STUB: not implemented"
	return *new(MessageHandler)
}

// if _, ok := msg.Body().(expr.ContextReader); !ok {
// 	u.Errorf("could not convert to message reader: %T", msg.Body())
// }

func msgToRow(msg schema.Message, cols []string, dest []driver.Value) error {
	_ = "STUB: not implemented"

	//u.Debugf("msg? %v  %T \n%p %v", msg, msg, dest, dest)
	return nil
}

/*
	case *datasource.ContextUrlValues:
		for i, key := range cols {
			if val, ok := mt.Get(key); ok && !val.Nil() {
				dest[i] = val.Value()
				//u.Infof("key=%v   val=%v", key, val)
			} else {
				u.Warnf("missing value? %v %T %v", key, val.Value(), val.Value())
			}
		}
		//u.Debugf("got msg in row result writer: %#v", mt)

	case *datasource.ContextSimple:
		for i, key := range cols {
			//u.Debugf("key=%v mt = nil? %v", key, mt)
			if val, ok := mt.Get(key); ok && val != nil && !val.Nil() {
				dest[i] = val.Value()
				//u.Infof("key=%v   val=%v", key, val)
			} else if val == nil {
				u.Errorf("could not evaluate? %v  %#v", key, mt)
			} else {
				u.Warnf("missing value? %v %T %v", key, val.Value(), val.Value())
			}
		}
		//u.Debugf("got msg in row result writer: %#v", dest)
*/

//u.Debugf("key=%v %T %v", key, val, val)

//u.Infof("key=%v   val=%v", key, val)

//u.Debugf("got msg in row result writer: %#v", dest)
