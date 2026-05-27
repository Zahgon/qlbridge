package exec

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/plan"
)

var (
	_ = u.EMPTY

	// Ensure that we implement the Tasks
	_ Task = (*TaskParallel)(nil)
)

// A parallel set of tasks, this starts each child task and offers up
//
//	 an output channel that is a merger of each child
//
//	--> \
//	--> - ->
//	--> /
type TaskParallel struct {
	*TaskBase
	in      TaskRunner
	runners []TaskRunner
	tasks   []Task
}

func NewTaskParallel(ctx *plan.Context) *TaskParallel { _ = "STUB: not implemented"; return nil }

func (m *TaskParallel) PrintDag(depth int) { _ = "STUB: not implemented"; return }

func (m *TaskParallel) Close() error { _ = "STUB: not implemented"; return nil }

func (m *TaskParallel) Setup(depth int) error { _ = "STUB: not implemented"; return nil }

//u.Infof("parallel task in: #%d task p:%p %T  %p", i, task, task, task.MessageIn())

//u.Debugf("%d  Setup: %T", depth, m.runners[i])

func (m *TaskParallel) Add(task Task) error { _ = "STUB: not implemented"; return nil }

func (m *TaskParallel) Children() []Task { _ = "STUB: not implemented"; return nil }

func (m *TaskParallel) Run() error {
	_ = "STUB: not implemented"
	// Our context can recover panics, save error msg
	return nil
}

// TODO:  find the culprit

//u.Errorf("panic on:  %v", r)

//u.WarnT(8)
// closing output channels is the signal to stop

// Either of the SigQuit, or error channel will
//  cause breaking out of message channels below

//m.errors = append(m.errors, err)

// start tasks in reverse order, so that by time
// source starts up all downstreams have started

//u.Infof("starting task %d-%d %T in:%p  out:%p", m.depth, taskId, task, task.MessageIn(), task.MessageOut())

// TODO:  what do we do with this error?   send to error channel?

//u.Debugf("exiting taskId: %v %T", taskId, task)
