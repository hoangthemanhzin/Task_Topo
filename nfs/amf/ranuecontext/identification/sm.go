package identification

import (
	"etrib5gc/util/fsm"

	"github.com/free5gc/nas/nasMessage"
)

const (
	PROC_IDLE fsm.StateType = iota
	PROC_WAITING
)
const (
	StartEvent      fsm.EventType = fsm.ExitEvent + iota + 1
	IdResponseEvent               //receive IdentityResponse
	DoneEvent                     //identification completes
	T3570Event
)

var _sm *fsm.Fsm

func init() {
	transitions := fsm.Transitions{
		fsm.Tuple(PROC_IDLE, StartEvent):         PROC_WAITING,
		fsm.Tuple(PROC_WAITING, IdResponseEvent): PROC_WAITING,
		fsm.Tuple(PROC_WAITING, T3570Event):      PROC_WAITING,
		fsm.Tuple(PROC_WAITING, DoneEvent):       PROC_IDLE,
	}

	callbacks := fsm.Callbacks{
		PROC_IDLE:    idle,
		PROC_WAITING: waiting,
	}
	_sm = fsm.NewFsm(transitions, callbacks)
}

func idle(state fsm.State, event fsm.EventType, args interface{}) {
	proc := state.(*IdProc)
	switch event {
	case fsm.EntryEvent:
		//identification procedure has ended
		if args != nil {
			proc.err, _ = args.(error)
		}
		proc.onDone(proc)
	}
}

func waiting(state fsm.State, event fsm.EventType, args interface{}) {
	proc := state.(*IdProc)
	switch event {
	case fsm.EntryEvent:
		proc.failurecount = 0
		proc.request()
	case IdResponseEvent:
		msg := args.(*nasMessage.IdentityResponse)
		proc.handleResponse(msg)

	case T3570Event:
		proc.handleT3570()
	}

}
