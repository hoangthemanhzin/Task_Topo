package uecontext

import (
	"etrib5gc/common"
	"etrib5gc/nas"
	"etrib5gc/nfs/amf/events"
	"etrib5gc/util/fsm"
)

const (
	MM_IDLE fsm.StateType = iota
	MM_REGISTERING
)

var _sm *fsm.Fsm

func init() {
	transitions := fsm.Transitions{
		fsm.Tuple(MM_IDLE, N1N2TransferEvent):            MM_IDLE,
		fsm.Tuple(MM_REGISTERING, N1N2TransferEvent):     MM_REGISTERING,
		fsm.Tuple(MM_IDLE, RegistrationRequestEvent):     MM_REGISTERING,
		fsm.Tuple(MM_REGISTERING, RegistrationDoneEvent): MM_IDLE,
		fsm.Tuple(MM_REGISTERING, UpdateSecmodeEvent):    MM_REGISTERING,
	}

	callbacks := fsm.Callbacks{
		MM_IDLE:        mm_idle,
		MM_REGISTERING: mm_registering,
	}
	_sm = fsm.NewFsm(transitions, callbacks)

}

const (
	RegistrationRequestEvent fsm.EventType = fsm.ExitEvent + iota + 1
	RegistrationDoneEvent                  //registration request or service request has been handled
	UpdateSecmodeEvent
	N1N2TransferEvent
)

func mm_idle(state fsm.State, event fsm.EventType, args interface{}) {
	uectx := state.(*UeContext)
	//NOTE: if UeContext is idle, we are free to replace a RanUe with a new one
	switch event {
	case fsm.EntryEvent:

	case RegistrationRequestEvent:
		ranue, _ := args.(RanUe)
		//update Ue with any authenticated information that was retrieved from
		//DAMF, and information from the request message
		ctx := ranue.RegistrationContext()
		uectx.update(ctx)

		//attach the RanUe then forward to it for handling the request
		uectx.AttachRanUe(ranue)
		ranue.HandleEvent(&common.EventData{
			EvType:  events.REGISTRATION_GRANTED,
			Content: ctx,
		})
	case N1N2TransferEvent:
		job, _ := args.(*common.AsyncJob)
		uectx.handleN1N2Transfer(job)
	default:
	}
}

func mm_registering(state fsm.State, event fsm.EventType, args interface{}) {
	uectx := state.(*UeContext)
	//NOTE: once a registration is done, the UeContext must receive a
	//notification event from the RanUe (who starts the registration)
	switch event {
	case fsm.EntryEvent:
	case RegistrationRequestEvent:
		ranue, _ := args.(RanUe)
		//reject any other registration request
		ranue.HandleEvent(&common.EventData{
			EvType:  events.REGISTRATION_REJECTED,
			Content: ranue.RegistrationContext(),
		})

	case RegistrationDoneEvent:
		uectx.Info("RanUe notify Registration Done")
	case UpdateSecmodeEvent:
		uectx.Info("Update Security Mode")
		uectx.secctx, _ = args.(*nas.SecCtx)

	case N1N2TransferEvent:
		job, _ := args.(*common.AsyncJob)
		uectx.handleN1N2Transfer(job)
	default:
	}

}
