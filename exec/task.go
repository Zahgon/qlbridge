package exec

import (
	"sync"

	"github.com/araddon/qlbridge/plan"
)

const (
	// ItemDefaultChannelSize default channel buffer for task's
	ItemDefaultChannelSize = 50
)

// TaskBase Base executeable task that implements Task interface, embedded
// into other channel based task runners
type TaskBase struct {
	sync.Mutex
	Ctx      *plan.Context
	Name     string
	Handler  MessageHandler
	depth    int
	setup    bool
	closed   bool
	hasquit  bool
	msgInCh  MessageChan
	msgOutCh MessageChan
	errCh    ErrChan
	sigCh    SigChan // notify of quit/stop
	errors   []error
}

func NewTaskBase(ctx *plan.Context) *TaskBase { _ = "STUB: not implemented"; return nil }

// All Tasks Get output channels by default, but NOT input

func (m *TaskBase) Children() []Task      { _ = "STUB: not implemented"; return nil }
func (m *TaskBase) Setup(depth int) error { _ = "STUB: not implemented"; return nil }

//u.Debugf("setup() %s %T in:%p  out:%p", m.TaskType, m, m.msgInCh, m.msgOutCh)

func (m *TaskBase) Add(task Task) error          { _ = "STUB: not implemented"; return nil }
func (m *TaskBase) AddPlan(task plan.Task) error { _ = "STUB: not implemented"; return nil }

func (m *TaskBase) MessageIn() MessageChan       { _ = "STUB: not implemented"; return *new(MessageChan) }
func (m *TaskBase) MessageOut() MessageChan      { _ = "STUB: not implemented"; return *new(MessageChan) }
func (m *TaskBase) MessageInSet(ch MessageChan)  { _ = "STUB: not implemented"; return }
func (m *TaskBase) MessageOutSet(ch MessageChan) { _ = "STUB: not implemented"; return }
func (m *TaskBase) ErrChan() ErrChan             { _ = "STUB: not implemented"; return *new(ErrChan) }
func (m *TaskBase) SigChan() SigChan             { _ = "STUB: not implemented"; return *new(SigChan) }
func (m *TaskBase) Quit()                        { _ = "STUB: not implemented"; return }

func (m *TaskBase) Close() error {
	_ = "STUB: not implemented"
	// u.Debugf("%p start Close()", m)
	return nil
}

//u.Debugf("%p finished Close()", m)

func (m *TaskBase) CloseFinal() error { _ = "STUB: not implemented"; return nil }

func MakeHandler(task TaskRunner) MessageHandler {
	_ = "STUB: not implemented"
	return *new(MessageHandler)
}

func (m *TaskBase) Run() error {
	_ = "STUB: not implemented"
	// Our context can recover panics, save error msg
	return nil
}

// closing output channels is the signal to stop
//u.Debugf("close taskbase: ch:%p    %v", m.msgOutCh, m.Type())

//u.Debugf("TaskBase: %T inchan", m)

// Either of the SigQuit, or error channel will
//  cause breaking out of message channels below

//m.errors = append(m.errors, err)

// Signal, ie quit etc
//u.Debugf("got taskbase signal")

//

//u.Debugf("sending to handler: %T  %+v", msg, msg)

//u.Debugf("msg in closed shutting down")

//u.Warnf("exiting")

// On Task stepper we don't Run it, rather use a
//
//	Next() explicit call from end user
type TaskStepper struct {
	*TaskBase
}

func NewTaskStepper(ctx *plan.Context) *TaskStepper { _ = "STUB: not implemented"; return nil }

func (m *TaskStepper) Run() error { _ = "STUB: not implemented"; return nil }

// Our context can recover panics, save error msg
// closing output channels is the signal to stop
