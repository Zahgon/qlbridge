package exec

import (
	u "github.com/araddon/gou"

	"github.com/araddon/qlbridge/plan"
)

var (
	_ = u.EMPTY

	// Ensure that we implement the plan.Tasks
	_ Task = (*TaskSequential)(nil)
)

type TaskSequential struct {
	*TaskBase
	closed  bool
	tasks   []Task
	runners []TaskRunner
}

func NewTaskSequential(ctx *plan.Context) *TaskSequential { _ = "STUB: not implemented"; return nil }

func (m *TaskSequential) PrintDag(depth int) { _ = "STUB: not implemented"; return }

func (m *TaskSequential) Close() error {
	_ = "STUB: not implemented"
	// u.Debugf("%p start Close() closed?%v", m, m.closed)
	return nil
}

//u.Debugf("%p task.Close()  %T", task, task)

func (m *TaskSequential) Setup(depth int) error {
	_ = "STUB: not implemented"
	// We don't need to setup the First(source) Input channel
	return nil
}

//u.Debugf("%d i:%d  Setup: %T p:%p", depth, i, m.runners[i], m.runners[i])

//u.Infof("%d  TaskSequential Setup  tasks len=%d", depth, len(m.tasks))

//u.Infof("%d-%d setup msgin: %T  %p", depth, i, m.runners[i], m.runners[i].MessageIn())

//u.Debugf("setup() %T in:%p  out:%p", m, m.msgInCh, m.msgOutCh)

func (m *TaskSequential) Add(task Task) error { _ = "STUB: not implemented"; return nil }

func (m *TaskSequential) Children() []Task { _ = "STUB: not implemented"; return nil }

func (m *TaskSequential) Run() (err error) {
	_ = "STUB: not implemented"
	// Our context can recover panics, save error msg
	return nil
}

//close(m.msgOutCh) // closing output channels is the signal to stop
//u.Debugf("close TaskSequential: %v", m.Type())

// Either of the SigQuit, or error channel will
//  cause breaking out of task execution below
// go func() {
// 	select {
// 	case err := <-m.errCh:
// 		u.Errorf("error on run %v", err)
// 	case <-m.sigCh:
// 		u.Warnf("%p %q got quit channel?", m, m.Name)
// 		// If we close here, we close without draining not giving messaging time
// 		// so we should????
// 		//err = m.Close()
// 		// for _, task := range m.runners {
// 		// 	task.Quit()
// 		// }
// 	}
// }()

// start tasks in reverse order, so that by time
// source starts up all downstreams have started

//u.Infof("starting task %d-%d %T in:%p  out:%p", m.depth, taskId, task, task.MessageIn(), task.MessageOut())

// TODO:  what do we do with this error?   send to error channel?

//u.Debugf("%p %q exiting taskId: %p %v %T", m, m.Name, task, taskId, task)

// Lets look for the last task to shutdown, the result-writer or projection
// will finish first on limit so we need to shutdown sources

//u.Warnf("%p got shutdown on last one, lets shutdown them all", m)

//u.Debugf("%p sending close??: %v %T", m, i, m.runners[i])

//u.Debugf("%p after close??: %v %T", m, i, m.runners[i])

// block until all tasks have finished
//u.Debugf("%p exit TaskSequential Run():  %q", m, m.Name)
