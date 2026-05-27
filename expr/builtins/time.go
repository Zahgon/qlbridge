package builtins

import (
	u "github.com/araddon/gou"
	"github.com/araddon/qlbridge/expr"
	"github.com/araddon/qlbridge/value"
)

var _ = u.EMPTY

// Now Get current time of Message (message time stamp) or else choose current
// server time if none is available in message context
type Now struct{}

// Type time
func (m *Now) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }

func (m *Now) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func nowEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Yy Get year in integer from field, must be able to convert to date
//
//	yy()                 =>  15, true    // assuming it is 2015
//	yy("2014-03-01")     =>  14, true
type Yy struct{}

// Type integer
func (m *Yy) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Yy) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func yearEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// mm Get month as integer from date
//
// @optional timestamp (if not, gets from context reader)
//
//	mm()                =>  01, true  /// assuming message ts = jan 1
//	mm("2014-03-17")    =>  03, true
type Mm struct{}

// Type integer
func (m *Mm) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *Mm) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func monthEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// yymm convert date to 4 digit string from argument if supplied, else uses message context ts
//
//	yymm() => "1707", true
//	yymm("2016/07/04") => "1607", true
type YyMm struct{}

// Type string
func (m *YyMm) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *YyMm) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func yymmEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// DayOfWeek day of week [0-6]
//
//	dayofweek() => 3, true
//	dayofweek("2016/07/04") => 5, true
type DayOfWeek struct{}

// Type int
func (m *DayOfWeek) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *DayOfWeek) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func dayOfWeekEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// hour of week [0-167]
type HourOfWeek struct{}

// Type int
func (m *HourOfWeek) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}
func (m *HourOfWeek) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func hourOfWeekEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// hourofday hour of day [0-23]
//
//	hourofday(field)
//	hourofday()  // Uses message time
type HourOfDay struct{}

// Type integer
func (m *HourOfDay) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *HourOfDay) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func hourOfDayEval(ctx expr.EvalContext, vals []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// totimestamp:   convert to date, then to unix Seconds
//
//	totimestamp() => int, true
//	totimestamp("Apr 7, 2014 4:58:55 PM") => 1396889935
type ToTimestamp struct{}

// Type integer
func (m *ToTimestamp) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}
func (m *ToTimestamp) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func toTimestampEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// todate:   convert to Date
//
//	// uses lytics/datemath
//	todate("now-3m")
//
//	// uses araddon/dateparse util to recognize formats
//	todate(field)
//
//	// first parameter is the layout/format
//	todate("01/02/2006", field )
type ToDate struct{}

// Type time
func (m *ToDate) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *ToDate) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func toDateEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Is date math

//u.Infof("hello  layout=%v  time=%v", formatStr, dateStr)

// todatein:   convert to Date with timezon
//
//	// uses lytics/datemath
//	todate("now-3m", "America/Los_Angeles")
//
//	// uses araddon/dateparse util to recognize formats
//	todate(field, "America/Los_Angeles")
type ToDateIn struct{}

// Type time
func (m *ToDateIn) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *ToDateIn) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

// Is date math

// Return the Evaluator

// We are going to correct back to UTC. so all fields are in UTC.

// TimeSeconds time in Seconds, parses a variety of formats looking for seconds
// See github.com/araddon/dateparse for formats supported on date parsing
//
//	seconds("M10:30")      =>  630
//	seconds("M100:30")     =>  6030
//	seconds("00:30")       =>  30
//	seconds("30")          =>  30
//	seconds(30)            =>  30
//	seconds("2015/07/04")  =>  1435968000
type TimeSeconds struct{}

// Type number
func (m *TimeSeconds) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}
func (m *TimeSeconds) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func timeSecondsEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// First, lets try to treat it as a time/date and
// then extract unix seconds

// Since that didn't work, lets look for a variety of seconds/minutes type
// pseudo standards
//    M10:30
//     10:30
//    100:30
//

// UnixDateTruncFunc converts a value.Value to a unix timestamp string. This is used for the BigQuery export
// since value.TimeValue returns a unix timestamp with milliseconds (ie "1438445529707") by default.
// This gets displayed in BigQuery as "47547-01-24 10:49:05 UTC", because they expect seconds instead of milliseconds.
// This function "truncates" the unix timestamp to seconds to the form "1438445529.707"
// i.e the milliseconds are placed after the decimal place.
// Inspired by the "DATE_TRUNC" function used in PostgresQL and RedShift:
// http://www.postgresql.org/docs/8.1/static/functions-datetime.html#FUNCTIONS-DATETIME-TRUNC
//
//	unixtrunc("1438445529707") --> "1438445529"
//	unixtrunc("1438445529707", "seconds") --> "1438445529.707"
//	unixtrunc("1438445529707123456", "secondsmicro") --> "1438445529.707123"
type TimeTrunc struct{ precision string }

// Type string
func (m *TimeTrunc) Type() value.ValueType { _ = "STUB: not implemented"; return *new(value.ValueType) }
func (m *TimeTrunc) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func timeTruncEvalOne(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// If the value is of type "TimeValue", return the Unix representation.

// Otherwise use date parse any

func timeTruncEvalTwo(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"
	return *new(value.Value), false
}

// Get full Unix timestamps w/ microseconds and milliseconds.

// If not a TimeValue, convert to a TimeValue and get Unix w/ milliseconds.

// Look at the seconds argument to determine the truncation.

// If seconds: add milliseconds after the decimal place.

// Otherwise return the Unix ts w/ milliseconds.

// StrFromTime extraces certain parts from a time, similar to Python's StrfTime
// See http://strftime.org/ for Strftime directives.
//
//	strftime("2015/07/04", "%B")      => "July"
//	strftime("2015/07/04", "%B:%d")   => "July:4"
//	strftime("1257894000", "%p")      => "PM"
type StrFromTime struct{}

// Type string
func (m *StrFromTime) Type() value.ValueType {
	_ = "STUB: not implemented"
	return *new(value.ValueType)
}
func (m *StrFromTime) Validate(n *expr.FuncNode) (expr.EvaluatorFunc, error) {
	_ = "STUB: not implemented"
	return *new(expr.EvaluatorFunc), nil
}

func strFromTimeEval(ctx expr.EvalContext, args []value.Value) (value.Value, bool) {
	_ = "STUB: not implemented"

	// if we have 2 items, the first is the time string
	// and the second is the format string.
	// Use leekchan/timeutil package
	return *new(value.Value), false
}
