package topo

import (
	"etrib5gc/pfcp"
	"etrib5gc/util/fsm"
)

const (
	UPF_NONASSOCIATED fsm.StateType = iota
	UPF_CONNECTED
	UPF_CONNECTING
)

var _sm *fsm.Fsm

const (
	ConnectedEvent     fsm.EventType = fsm.ExitEvent + iota + 1 //any message is received
	StartEvent                                                  //start association
	AssociatedEvent                                             //has association
	FailEvent                                                   //any communication failure
	DisassociatedEvent                                          //disassociated (either remote or local)
	SendHbEvent                                                 //to send a heartbeat event
)

func init() {
	transitions := fsm.Transitions{
		fsm.Tuple(UPF_NONASSOCIATED, AssociatedEvent): UPF_CONNECTED,
		fsm.Tuple(UPF_NONASSOCIATED, StartEvent):      UPF_NONASSOCIATED,

		fsm.Tuple(UPF_CONNECTING, ConnectedEvent):     UPF_CONNECTING,
		fsm.Tuple(UPF_CONNECTING, AssociatedEvent):    UPF_CONNECTED,
		fsm.Tuple(UPF_CONNECTING, SendHbEvent):        UPF_CONNECTING,
		fsm.Tuple(UPF_CONNECTING, DisassociatedEvent): UPF_NONASSOCIATED,

		fsm.Tuple(UPF_CONNECTED, ConnectedEvent):     UPF_CONNECTED,
		fsm.Tuple(UPF_CONNECTED, SendHbEvent):        UPF_CONNECTED,
		fsm.Tuple(UPF_CONNECTED, FailEvent):          UPF_CONNECTING,
		fsm.Tuple(UPF_CONNECTED, DisassociatedEvent): UPF_NONASSOCIATED,
		//add more transitions
	}

	callbacks := fsm.Callbacks{
		UPF_NONASSOCIATED: nonassociated, //not associated
		UPF_CONNECTING:    connecting,    //associated, temporarily out of reach
		UPF_CONNECTED:     connected,     //associated and connected
	}
	_sm = fsm.NewFsm(transitions, callbacks)
}

func nonassociated(state fsm.State, event fsm.EventType, args interface{}) {
	node := state.(*topoNode)
	switch event {
	case fsm.EntryEvent:
		log.Tracef("%s enter NONASSOCIATED", node.id)
		node.assotimer.Start()
	case StartEvent:
		pfcp := args.(*pfcp.Pfcp)
		//send association
		if err := node.sendAssociationRequest(pfcp); err == nil {
			node.sendEvent(AssociatedEvent, nil)
			node.assotimer.Stop()
		} else {
			//log.Infof("%s Try sending an association; return: %s", node.id, err.Error())
			node.assotimer.Start()
		}
	}
}

func connected(state fsm.State, event fsm.EventType, args interface{}) {
	node := state.(*topoNode)
	switch event {
	case fsm.EntryEvent:
		log.Infof("UPF [%s] is CONNECTED", node.id)
		node.setActive(true)
		fallthrough
	case ConnectedEvent:
		//reset timer
		node.hbtry = 0
		node.hbtimer.Start()

	case SendHbEvent:
		pfcp := args.(*pfcp.Pfcp)
		//send a heartbeat
		if node.sendHeartbeat(pfcp) != nil {
			log.Warnf("%s fails sending heartbeat %d", node.id, node.hbtry)
			node.sendEvent(FailEvent, nil)
		} else {
			node.hbtimer.Start()
		}
	case fsm.ExitEvent:
		node.setActive(false)
	}
}

func connecting(state fsm.State, event fsm.EventType, args interface{}) {
	node := state.(*topoNode)
	switch event {
	case fsm.EntryEvent:
		log.Tracef("%s is in CONNECTING; starts connecting to upf", node.id)
		node.rectimer.Start()
	case SendHbEvent:
		log.Tracef("%s try recovering %d times", node.id, node.hbtry)
		pfcp := args.(*pfcp.Pfcp)
		//send a heartbeat
		if node.sendHeartbeat(pfcp) == nil {
			node.sendEvent(ConnectedEvent, nil)
		} else if node.hbtry >= MAX_HEARTBEAT_TRY {
			node.sendEvent(DisassociatedEvent, nil)
		} else { //try to send heartbeat again later
			node.rectimer.Start()
		}
	case ConnectedEvent:
		//will transit to UPF_CONNECTED
	case fsm.ExitEvent:
		node.rectimer.Stop()
	}

}
