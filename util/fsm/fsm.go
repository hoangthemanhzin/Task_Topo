package fsm

import (
	"etrib5gc/common"
	"fmt"
)

type StateType int
type EventType int

const (
	EntryEvent EventType = iota
	ExitEvent
)

type State interface {
	CurrentState() StateType
	SetState(StateType)
}
type bareState struct {
	current StateType
}

func NewState(i StateType) State {
	return &bareState{
		current: i,
	}
}
func (s *bareState) SetState(now StateType) {
	s.current = now
}
func (s *bareState) CurrentState() StateType {
	return s.current
}

type StateEventTuple struct {
	state StateType
	event EventType
}

func Tuple(state StateType, event EventType) (tuple StateEventTuple) {
	tuple.event = event
	tuple.state = state
	return
}

type Transitions map[StateEventTuple]StateType
type CallbackFn func(State, EventType, interface{})
type Callbacks map[StateType]CallbackFn

type Fsm struct {
	transitions Transitions
	callbacks   Callbacks
}

func NewFsm(ts Transitions, clbks Callbacks) *Fsm {
	ret := &Fsm{
		transitions: ts,
		callbacks:   clbks,
	}
	knowns := make(map[StateType]bool)
	for t, _ := range ts {
		knowns[t.state] = true
	}

	for s, _ := range knowns {
		if _, ok := clbks[s]; !ok {
			panic("bad state machine configuration")
		}
	}
	return ret
}

func (fsm *Fsm) SendEvent(exec common.Executer, state State, event EventType, args interface{}) error {
	current := state.CurrentState()

	tuple := StateEventTuple{
		state: current,
		event: event,
	}

	if nextstate, ok := fsm.transitions[tuple]; ok {
		//event handler function
		fn := func() {
			fsm.callbacks[current](state, event, args)
			if current != nextstate { //there is a state transition
				fsm.callbacks[current](state, ExitEvent, args)
				state.SetState(nextstate)
				fsm.callbacks[nextstate](state, EntryEvent, args)
			}
		}
		if !exec.AddJob(fn) {
			return fmt.Errorf("Too many events have been sent to the state machine, this one is discarded")
		}
	} else {
		return fmt.Errorf("Unknown transition from state %v with event %v", current, event)
	}

	return nil
}
