package authentication

import (
	"etrib5gc/util/fsm"

	"github.com/free5gc/nas/nasMessage"
)

const (
	AUTH_IDLE fsm.StateType = iota
	AUTH_CHALLENGING
	AUTH_DONE
)
const (
	StartEvent fsm.EventType = fsm.ExitEvent + iota + 1
	AuthResponseEvent
	AuthFailureEvent
	DoneEvent //authentication completes
	T3560Event
)

var _sm *fsm.Fsm

func init() {
	transitions := fsm.Transitions{
		fsm.Tuple(AUTH_IDLE, StartEvent):               AUTH_CHALLENGING,
		fsm.Tuple(AUTH_CHALLENGING, AuthResponseEvent): AUTH_CHALLENGING,
		fsm.Tuple(AUTH_CHALLENGING, AuthFailureEvent):  AUTH_CHALLENGING,
		fsm.Tuple(AUTH_CHALLENGING, T3560Event):        AUTH_CHALLENGING,
		fsm.Tuple(AUTH_CHALLENGING, DoneEvent):         AUTH_IDLE,
	}

	callbacks := fsm.Callbacks{
		AUTH_IDLE:        idle,
		AUTH_CHALLENGING: challenging,
	}
	//Note: make sure that transitions and states are well-defined. The program
	//will panic if an error is returned
	_sm = fsm.NewFsm(transitions, callbacks)
}

func idle(state fsm.State, event fsm.EventType, args interface{}) {
	proc := state.(*AuthProc)
	switch event {
	case fsm.EntryEvent:
		if args != nil {
			proc.err = args.(error)
		}
		proc.onDone(proc)
	default:
	}
}

func challenging(state fsm.State, event fsm.EventType, args interface{}) {
	proc := state.(*AuthProc)
	switch event {
	case fsm.EntryEvent:
		//start challening
		proc.challenge()
	case AuthResponseEvent:
		msg := args.(*nasMessage.AuthenticationResponse)
		proc.handleAuthenticationResponse(msg)
	case AuthFailureEvent:
		msg := args.(*nasMessage.AuthenticationFailure)
		proc.handleAuthenticationFailure(msg)
	case T3560Event:
		proc.handleT3560()
	default:
	}
}
