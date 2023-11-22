package secmode

import (
	"etrib5gc/util/fsm"

	"github.com/free5gc/nas/nasMessage"
)

const (
	SECMODE_IDLE    fsm.StateType = iota
	SECMODE_WAITING               //waiting to complete
)
const (
	StartEvent fsm.EventType = fsm.ExitEvent + iota + 1
	T3560Event
	SecmodeCompleteEvent
	SecmodeRejectEvent
	DoneEvent
)

var _sm *fsm.Fsm

func init() {
	transitions := fsm.Transitions{
		fsm.Tuple(SECMODE_IDLE, StartEvent):              SECMODE_WAITING,
		fsm.Tuple(SECMODE_WAITING, SecmodeCompleteEvent): SECMODE_WAITING,
		fsm.Tuple(SECMODE_WAITING, SecmodeRejectEvent):   SECMODE_WAITING,
		fsm.Tuple(SECMODE_WAITING, T3560Event):           SECMODE_WAITING,
		fsm.Tuple(SECMODE_WAITING, DoneEvent):            SECMODE_IDLE,
	}

	callbacks := fsm.Callbacks{
		SECMODE_IDLE:    idle,
		SECMODE_WAITING: waiting,
	}
	//Note: make sure that transitions and states are well-defined. The program
	//will panic if an error is returned
	_sm = fsm.NewFsm(transitions, callbacks)
}

func idle(state fsm.State, event fsm.EventType, args interface{}) {
	proc := state.(*SecmodeProc)
	switch event {
	case fsm.EntryEvent:
		if args != nil {
			proc.err = args.(error)
		}
		proc.onDone(proc)
	}

}
func waiting(state fsm.State, event fsm.EventType, args interface{}) {
	proc := state.(*SecmodeProc)

	switch event {
	case fsm.EntryEvent:
		proc.sendSecmodeCommand()

	case T3560Event:
		proc.Info("Receiver a T3560 timer expiring event")
		proc.handleT3560()
	case SecmodeCompleteEvent:
		msg, _ := args.(*nasMessage.SecurityModeComplete)
		proc.handleSecurityModeComplete(msg)

	case SecmodeRejectEvent:
		msg, _ := args.(*nasMessage.SecurityModeReject)
		proc.handleSecurityModeReject(msg)
	}
}
