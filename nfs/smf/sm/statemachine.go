package sm

import (
	"etrib5gc/common"
	"etrib5gc/util/fsm"

	"github.com/free5gc/nas/nasMessage"
)

const (
	SM_INACTIVE fsm.StateType = iota
	SM_ACTIVATING
	SM_ACTIVE
	SM_INACTIVATING
)

var _sm *fsm.Fsm

func init() {

	transitions := fsm.Transitions{
		fsm.Tuple(SM_INACTIVE, PostSmContextsEvent): SM_INACTIVE,
		fsm.Tuple(SM_INACTIVE, SessActEvent):        SM_ACTIVATING,

		fsm.Tuple(SM_ACTIVATING, SessActCmplEvent): SM_ACTIVE,
		fsm.Tuple(SM_ACTIVATING, SessActFailEvent): SM_INACTIVATING,

		fsm.Tuple(SM_ACTIVATING, ActTimeoutEvent): SM_INACTIVATING,

		fsm.Tuple(SM_ACTIVE, SessDeactEvent): SM_INACTIVATING,

		fsm.Tuple(SM_INACTIVATING, DeactTimeoutEvent): SM_INACTIVE,

		fsm.Tuple(SM_ACTIVE, UpdateSmContextEvent):       SM_ACTIVE,
		fsm.Tuple(SM_ACTIVATING, UpdateSmContextEvent):   SM_ACTIVATING,
		fsm.Tuple(SM_INACTIVATING, UpdateSmContextEvent): SM_INACTIVATING,

		fsm.Tuple(SM_ACTIVE, ReleaseSmContextEvent):       SM_INACTIVE,
		fsm.Tuple(SM_ACTIVATING, ReleaseSmContextEvent):   SM_INACTIVE,
		fsm.Tuple(SM_INACTIVATING, ReleaseSmContextEvent): SM_INACTIVE,

		fsm.Tuple(SM_ACTIVE, CloseEvent):       SM_INACTIVE,
		fsm.Tuple(SM_ACTIVATING, CloseEvent):   SM_INACTIVE,
		fsm.Tuple(SM_INACTIVATING, CloseEvent): SM_INACTIVE,
	}

	callbacks := fsm.Callbacks{
		SM_INACTIVE:     sm_inactive,     // no activated
		SM_ACTIVE:       sm_active,       // activated
		SM_INACTIVATING: sm_inactivating, //inactivating
		SM_ACTIVATING:   sm_activating,   //activating
	}
	_sm = fsm.NewFsm(transitions, callbacks)
}

const (
	SessActEvent       fsm.EventType = fsm.ExitEvent + iota + 1 //activating session
	SessActCmplEvent                                            //session is activated
	SessActFailEvent                                            //activation failed
	SessDeactEvent                                              //deactivate session
	SessDeactCmplEvent                                          //session is deactivated
	ActTimeoutEvent
	DeactTimeoutEvent
	PostSmContextsEvent
	UpdateSmContextEvent
	ReleaseSmContextEvent
	CloseEvent
)

func sm_inactive(state fsm.State, event fsm.EventType, args interface{}) {
	smctx := state.(*SmContext)
	switch event {
	case fsm.EntryEvent:
		//re-enter INACTIVE -> remove it from the smlist
		smctx.Trace("remove the smctx %s", smctx.ref)
		smctx.ctx.RemoveSmContext(smctx)

	case PostSmContextsEvent:
		job, _ := args.(*common.AsyncJob)
		info, _ := job.Info().(*PostSmContextsJob)
		smctx.handlePostSmContexts(info)
		if info.Ersp != nil {
			smctx.ctx.RemoveSmContext(smctx)
		} //NOTE: in case there was no error, a next event is already sent
		job.Done(nil)
	}
}

func sm_activating(state fsm.State, event fsm.EventType, args interface{}) {
	smctx := state.(*SmContext)
	var err error
	switch event {
	case fsm.EntryEvent:
		//start activting timer
		smctx.Trace("Start activation timer on entering SM_ACTIVATING")
		smctx.acttimer.Start()

		//activate pfcp sessions
		if err = smctx.establishPfcpSessions(); err == nil {
			//send PduSessionEstablishmentAccept
			if err = smctx.acceptPduSessionEstablishment(); err != nil {
				//if failed, go to SM_INACTIVATING
				smctx.sendEvent(SessActFailEvent, nil)
			} else {
				//NOTE: should we wait for response to go to Active mode?
				smctx.sendEvent(SessActCmplEvent, nil)
			}
		} else {
			//send PduSessionEstablishmentReject
			smctx.rejectPduSessionEstablishment(nasMessage.Cause5GSMProtocolErrorUnspecified)
			//go to SM_INACTIVATING
			smctx.sendEvent(SessActFailEvent, nil)
		}
	case fsm.ExitEvent:
		//stop activting timer (it may already timeout)
		smctx.Trace("Stop activation timer on exiting SM_ACTIVATING")
		smctx.acttimer.Stop()

	case ReleaseSmContextEvent:
		job, _ := args.(*common.AsyncJob)
		err := smctx.releaseSmContext()
		smctx.onkill = func() {
			job.Done(err)
		}

	case UpdateSmContextEvent:
		job, _ := args.(*common.AsyncJob)
		info, _ := job.Info().(*UpdateSmContextJob)
		smctx.handleUpdateSmContext(info)
		job.Done(nil)

	case CloseEvent:
		smctx.releaseSmContext()
	}
}

func sm_active(state fsm.State, event fsm.EventType, args interface{}) {
	smctx := state.(*SmContext)
	switch event {
	case fsm.EntryEvent:

	case ReleaseSmContextEvent:
		job, _ := args.(*common.AsyncJob)
		err := smctx.releaseSmContext()
		smctx.onkill = func() {
			job.Done(err)
		}

	case UpdateSmContextEvent:
		job, _ := args.(*common.AsyncJob)
		info, _ := job.Info().(*UpdateSmContextJob)
		smctx.handleUpdateSmContext(info)
		job.Done(nil)
	case CloseEvent:
		smctx.releaseSmContext()
	}
}

func sm_inactivating(state fsm.State, event fsm.EventType, args interface{}) {
	smctx := state.(*SmContext)
	switch event {
	case fsm.EntryEvent:
		smctx.Trace("Start deactivation timer on entering SM_INACTIVATING")
		smctx.deacttimer.Start()
	case fsm.ExitEvent:
		//stop the deactivating timer (it may already timeout)
		smctx.Trace("Stop deactivation timer on exiting SM_INACTIVATING")
		smctx.deacttimer.Stop()

	case ReleaseSmContextEvent:
		job, _ := args.(*common.AsyncJob)
		err := smctx.releaseSmContext()
		smctx.onkill = func() {
			job.Done(err)
		}

	case UpdateSmContextEvent:
		job, _ := args.(*common.AsyncJob)
		info, _ := job.Info().(*UpdateSmContextJob)
		smctx.handleUpdateSmContext(info)
		job.Done(nil)

	case CloseEvent:
		smctx.releaseSmContext()
	}

}
