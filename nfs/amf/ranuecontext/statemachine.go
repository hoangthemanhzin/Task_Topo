package ranuecontext

import (
	"etrib5gc/common"
	"etrib5gc/nas"
	"etrib5gc/nfs/amf/events"
	"etrib5gc/nfs/amf/ranuecontext/authentication"
	"etrib5gc/nfs/amf/ranuecontext/identification"
	"etrib5gc/nfs/amf/ranuecontext/secmode"
	"etrib5gc/sbi/models/n2models"
	"etrib5gc/util/fsm"

	libnas "github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
)

const (
	MM_DEREGISTERED   fsm.StateType = iota //idling
	MM_REGISTERING                         //start registration
	MM_IDENTIFYING                         //UE identity identifying
	MM_AUTHENTICATING                      //authenticating the UE
	MM_SECMODING                           //security mode establishing
	MM_CONTEXTSETTING                      //setting up the UE context
	MM_REGISTERED                          //UE is registered
	MM_DEREGISTERING                       //start de-registrationg
)

var _sm *fsm.Fsm

func init() {
	transitions := fsm.Transitions{
		//start to process InitialUeMsg
		fsm.Tuple(MM_DEREGISTERED, InitEvent):   MM_DEREGISTERED,
		fsm.Tuple(MM_DEREGISTERED, PagingEvent): MM_DEREGISTERED,

		//Registration procedure is granted to start  (by UeContext)
		fsm.Tuple(MM_DEREGISTERED, RegistrationRequestEvent): MM_REGISTERING,
		fsm.Tuple(MM_REGISTERED, RegistrationRequestEvent):   MM_REGISTERING,

		//N1Msg
		fsm.Tuple(MM_DEREGISTERED, N1MsgEvent):   MM_DEREGISTERED,
		fsm.Tuple(MM_REGISTERING, N1MsgEvent):    MM_REGISTERING,
		fsm.Tuple(MM_DEREGISTERING, N1MsgEvent):  MM_DEREGISTERING,
		fsm.Tuple(MM_REGISTERED, N1MsgEvent):     MM_REGISTERED,
		fsm.Tuple(MM_IDENTIFYING, N1MsgEvent):    MM_IDENTIFYING,
		fsm.Tuple(MM_AUTHENTICATING, N1MsgEvent): MM_AUTHENTICATING,
		fsm.Tuple(MM_SECMODING, N1MsgEvent):      MM_SECMODING,
		fsm.Tuple(MM_CONTEXTSETTING, N1MsgEvent): MM_CONTEXTSETTING,

		//N1N2TransferEvent
		fsm.Tuple(MM_DEREGISTERED, N1N2TransferEvent): MM_DEREGISTERED,
		fsm.Tuple(MM_REGISTERING, N1N2TransferEvent):  MM_REGISTERING,
		//		fsm.Tuple(MM_DEREGISTERING, N1N2TransferEvent):  MM_DEREGISTERING,
		fsm.Tuple(MM_REGISTERED, N1N2TransferEvent):     MM_REGISTERED,
		fsm.Tuple(MM_IDENTIFYING, N1N2TransferEvent):    MM_IDENTIFYING,
		fsm.Tuple(MM_AUTHENTICATING, N1N2TransferEvent): MM_AUTHENTICATING,
		fsm.Tuple(MM_SECMODING, N1N2TransferEvent):      MM_SECMODING,
		fsm.Tuple(MM_CONTEXTSETTING, N1N2TransferEvent): MM_CONTEXTSETTING,

		//from registering to a common procedure
		fsm.Tuple(MM_REGISTERING, IdenEvent):    MM_IDENTIFYING,
		fsm.Tuple(MM_REGISTERING, AuthEvent):    MM_AUTHENTICATING,
		fsm.Tuple(MM_REGISTERING, SecmodeEvent): MM_SECMODING,
		fsm.Tuple(MM_REGISTERING, SetupEvent):   MM_CONTEXTSETTING,

		//moving to a next common procedure
		fsm.Tuple(MM_IDENTIFYING, AuthEvent):       MM_AUTHENTICATING,
		fsm.Tuple(MM_AUTHENTICATING, SecmodeEvent): MM_SECMODING,
		fsm.Tuple(MM_SECMODING, SetupEvent):        MM_CONTEXTSETTING,

		//completion of a common procedure
		fsm.Tuple(MM_IDENTIFYING, IdenCmplEvent):    MM_IDENTIFYING,
		fsm.Tuple(MM_AUTHENTICATING, AuthCmplEvent): MM_IDENTIFYING,
		fsm.Tuple(MM_SECMODING, SecmodeCmplEvent):   MM_SECMODING,

		//any failure, move to registered
		fsm.Tuple(MM_REGISTERING, FailEvent):    MM_DEREGISTERED,
		fsm.Tuple(MM_IDENTIFYING, FailEvent):    MM_DEREGISTERED,
		fsm.Tuple(MM_AUTHENTICATING, FailEvent): MM_DEREGISTERED,
		fsm.Tuple(MM_SECMODING, FailEvent):      MM_DEREGISTERED,
		fsm.Tuple(MM_CONTEXTSETTING, FailEvent): MM_DEREGISTERED,

		//in context setting up state
		fsm.Tuple(MM_CONTEXTSETTING, SetupEvent):    MM_CONTEXTSETTING,
		fsm.Tuple(MM_CONTEXTSETTING, IdenEvent):     MM_CONTEXTSETTING,
		fsm.Tuple(MM_CONTEXTSETTING, IdenCmplEvent): MM_CONTEXTSETTING,
		fsm.Tuple(MM_CONTEXTSETTING, DoneEvent):     MM_REGISTERED,

		fsm.Tuple(MM_REGISTERING, T3550Event): MM_REGISTERING,

		fsm.Tuple(MM_REGISTERED, N1N2TransferEvent):   MM_REGISTERED,
		fsm.Tuple(MM_DEREGISTERED, N1N2TransferEvent): MM_DEREGISTERED,

		fsm.Tuple(MM_REGISTERED, UectxRelReqEvent):  MM_REGISTERED,
		fsm.Tuple(MM_REGISTERED, PduNotifyEvent):    MM_REGISTERED,
		fsm.Tuple(MM_REGISTERED, PduModIndEvent):    MM_REGISTERED,
		fsm.Tuple(MM_REGISTERED, PduModRspEvent):    MM_REGISTERED,
		fsm.Tuple(MM_REGISTERED, CtxSetupRspEvent):  MM_REGISTERED,
		fsm.Tuple(MM_REGISTERED, CtxSetupFailEvent): MM_REGISTERED,
		fsm.Tuple(MM_REGISTERED, PduSetupRspEvent):  MM_REGISTERED,
		fsm.Tuple(MM_REGISTERED, PduRelRspEvent):    MM_REGISTERED,

		//Send Notification
		fsm.Tuple(MM_REGISTERED, NotificationEvent): MM_REGISTERED,
	}

	callbacks := fsm.Callbacks{
		MM_DEREGISTERED:   deregistered,
		MM_REGISTERING:    registering,
		MM_IDENTIFYING:    identifying,
		MM_AUTHENTICATING: authenticating,
		MM_SECMODING:      secmoding,
		MM_CONTEXTSETTING: contextsetting,
		MM_DEREGISTERING:  deregistering,
		MM_REGISTERED:     registered,
	}
	_sm = fsm.NewFsm(transitions, callbacks)

}

const (
	N1MsgEvent fsm.EventType = fsm.ExitEvent + iota + 1 //N1Msg in UlNasTransport
	InitEvent                                           //receive InitUeMsg
	NasNonDeliveryEvent
	UectxRelReqEvent
	PduNotifyEvent
	PduModIndEvent
	N1N2TransferEvent
	PagingEvent
	NotificationEvent
	RegistrationRequestEvent //UeContext grants a start of the registration procedure
	PduModRspEvent
	PduRelRspEvent
	PduSetupRspEvent
	CtxSetupRspEvent
	CtxSetupFailEvent

	IdenEvent    //start identification
	AuthEvent    //start authentication
	SecmodeEvent //start security context establishment
	SetupEvent   //enter seting up ue context during registration

	IdenCmplEvent    //identification procedure completes
	AuthCmplEvent    //authentication procedure completes
	SecmodeCmplEvent //security mode establishment procedure completes

	FailEvent
	DoneEvent

	T3550Event // registration/service accept expires
	T3513Event //paging expires
	T3522Event //deregistration request expires
	T3565Event //notification expires

)

func deregistered(state fsm.State, event fsm.EventType, args interface{}) {
	ranue := state.(*RanUe)

	switch event {
	case fsm.EntryEvent:
		ranue.regctx = nil
		ranue.registered = false
		ranue.ue.HandleEvent(&common.EventData{
			EvType:  events.REGISTRATION_CMPL,
			Content: ranue,
		})
	case InitEvent:
		dat, _ := args.(*events.InitUeContextData)
		//Note: a rejection will be sent if anything goes wrong
		ranue.handleInitUeContext(dat)
	case PagingEvent:
	default:
	}
}

func registering(state fsm.State, event fsm.EventType, args interface{}) {
	ranue := state.(*RanUe)
	switch event {
	case fsm.EntryEvent:
		//ready to start registration
		if ranue.hasValidSecmode() {
			ranue.sendEvent(SetupEvent, ranue.regctx)
		} else if ranue.isAuthenticated() {
			ranue.sendEvent(SecmodeEvent, nil)
		} else if ranue.hasId() {
			ranue.sendEvent(AuthEvent, nil)
		} else {
			ranue.sendEvent(IdenEvent, nil)
		}
	case N1N2TransferEvent:
		job, _ := args.(*common.AsyncJob)
		ranue.handleN1N2Transfer(job)
	default:
	}
}
func identifying(state fsm.State, event fsm.EventType, args interface{}) {
	ranue := state.(*RanUe)
	switch event {
	case fsm.EntryEvent:
		ranue.Infof("Start identification procedure")
		idtype := nasMessage.MobileIdentity5GSTypeSuci
		ranue.activeproc = identification.New(ranue, idtype, func(proc *identification.IdProc) {
			ranue.sendEvent(IdenCmplEvent, proc)
		})

	case N1MsgEvent:
		msg, _ := args.(*libnas.GmmMessage)
		ranue.activeproc.Handle(msg)

	case IdenCmplEvent:
		proc, _ := args.(*identification.IdProc)
		if err := proc.GetError(); err != nil {
			ranue.Errorf(err.Error())
			ranue.rejectRegistration(0)
		} else if err := ranue.ue.UpdateId(proc.Id()); err != nil {
			ranue.Errorf(err.Error())
			ranue.rejectRegistration(0)
		} else {
			//UeContext has been updated with the new identification
			if ranue.hasId() {
				//start authentication
				ranue.sendEvent(AuthEvent, nil)
			} else {
				ranue.Errorf("UE still does not have an identity")
				//still has no Id, strange case
				//need to reject registration
				ranue.rejectRegistration(0)
			}
		}

	case N1N2TransferEvent:
		job, _ := args.(*common.AsyncJob)
		ranue.handleN1N2Transfer(job)

	case fsm.ExitEvent:
		ranue.activeproc = nil
	default:
	}
}

func authenticating(state fsm.State, event fsm.EventType, args interface{}) {
	ranue := state.(*RanUe)
	switch event {
	case fsm.EntryEvent:
		ranue.Infof("Start authentication  procedure")
		ranue.activeproc = authentication.New(ranue, func(proc *authentication.AuthProc) {
			ranue.sendEvent(AuthCmplEvent, proc)
		})

	case AuthCmplEvent:
		proc, _ := args.(*authentication.AuthProc)
		if err := proc.GetError(); err != nil {
			ranue.Errorf(err.Error())
			ranue.rejectRegistration(0)
		} else {
			//Update registration context with authenticated information
			authctx := ranue.regctx.AuthCtx()
			authctx.Kamf, authctx.Supi, authctx.Eap, authctx.Success = proc.AuthInfo()
			if ranue.isAuthenticated() {
				ranue.startSecmode()
			} else {
				//still not authenticated, strange case
				//need to send a rejection
				ranue.rejectRegistration(0)
			}
		}
	case N1MsgEvent:
		msg, _ := args.(*libnas.GmmMessage)
		ranue.activeproc.Handle(msg)

	case N1N2TransferEvent:
		job, _ := args.(*common.AsyncJob)
		ranue.handleN1N2Transfer(job)

	case fsm.ExitEvent:
		ranue.activeproc = nil
	default:
	}
}

func secmoding(state fsm.State, event fsm.EventType, args interface{}) {
	ranue := state.(*RanUe)
	switch event {
	case fsm.EntryEvent:
		ranue.startSecmode()

	case SecmodeCmplEvent:
		proc, _ := args.(*secmode.SecmodeProc)
		if err := proc.GetError(); err != nil {
			ranue.Errorf(err.Error())
			ranue.rejectRegistration(0)
		} else {
			if content := proc.NasContainer(); len(content) > 0 {
				if nasmsg, err := nas.Decode(ranue, content); err == nil {
					ranue.Debug("NAS container decoded")
					if err = ranue.regctx.UpdateMsg(nasmsg.GmmMessage); err != nil {
						ranue.Errorf("Update re-transmissed NasMsg failed", err.Error())
						ranue.rejectRegistration(0)
						return
					}
				} else {
					ranue.Errorf("Decode NasContainer failed: %s", err.Error())
					ranue.rejectRegistration(0)
					return
				}
			}
			//must have a valid secmode now
			ranue.ue.HandleEvent(&common.EventData{
				EvType:  events.UPDATE_SECMODE,
				Content: ranue.secctx,
			})
			ranue.sendEvent(SetupEvent, nil)
		}

	case N1MsgEvent:
		msg, _ := args.(*libnas.GmmMessage)
		ranue.activeproc.Handle(msg)

	case N1N2TransferEvent:
		job, _ := args.(*common.AsyncJob)
		ranue.handleN1N2Transfer(job)

	case fsm.ExitEvent:
		ranue.activeproc = nil
	default:
	}
}
func contextsetting(state fsm.State, event fsm.EventType, args interface{}) {
	ranue := state.(*RanUe)
	switch event {
	case fsm.EntryEvent:
		fallthrough
	case SetupEvent:
		if pei := ranue.ue.Pei(); len(pei) == 0 {
			//get pei
			ranue.sendEvent(IdenEvent, nasMessage.MobileIdentity5GSTypeImei)
			return
		}

		//cause := nasMessage.Cause5GMMImplicitlyDeregistered
		if msg := ranue.regctx.ServiceRequest(); msg != nil {
			ranue.handleService(msg)
		} else {
			msg := ranue.regctx.RegistrationRequest() //can't be nil

			switch ranue.regctx.RegType() {
			case nasMessage.RegistrationType5GSInitialRegistration:
				ranue.Debugf("RegistrationType: Initial Registration")
				ranue.handleInitialRegistration(msg)
			case nasMessage.RegistrationType5GSMobilityRegistrationUpdating:
				fallthrough
			case nasMessage.RegistrationType5GSPeriodicRegistrationUpdating:
				ranue.Debugf("RegistrationType: Periodic Registration Updating")
				ranue.handleMobilityAndPeriodicRegistrationUpdating(msg)
			//case nasMessage.RegistrationType5GSEmergencyRegistration:
			//case nasMessage.RegistrationType5GSReserved:
			//		ranue.Debugf("RegistrationType: Reserved")
			default:
				ranue.Debugf("RegistrationType: unknown or not supported")
				//just testing
				ranue.rejectRegistration(nasMessage.Cause5GMMMessageTypeNonExistentOrNotImplemented)
			}
		}
	case IdenEvent:
		idtype, _ := args.(uint8)
		ranue.activeproc = identification.New(ranue, idtype, func(proc *identification.IdProc) {
			ranue.sendEvent(IdenCmplEvent, proc)
		})
	case IdenCmplEvent:
		ranue.activeproc = nil
		proc, _ := args.(*identification.IdProc)
		var err error
		if err = proc.GetError(); err == nil {
			if err = ranue.ue.UpdateId(proc.Id()); err == nil {
				ranue.sendEvent(SetupEvent, nil)
			}
		}
		if err != nil {
			ranue.Errorf("Can't get Pei: %s", err.Error())
			ranue.rejectRegistration(0)
		}

	case CtxSetupRspEvent:
		msg, _ := args.(*n2models.InitCtxSetupRsp)
		ranue.handleInitCtxSetupRsp(msg)

	case CtxSetupFailEvent:
		msg, _ := args.(*n2models.InitCtxSetupFailure)
		ranue.handleInitCtxSetupFail(msg)

	case N1MsgEvent:
		msg := args.(*libnas.GmmMessage)
		if msg.GetMessageType() == libnas.MsgTypeIdentityResponse && ranue.activeproc != nil {
			ranue.activeproc.Handle(msg)
		}

	case N1N2TransferEvent:
		job, _ := args.(*common.AsyncJob)
		ranue.handleN1N2Transfer(job)

	default:
	}
}

func deregistering(state fsm.State, event fsm.EventType, args interface{}) {
	ranue := state.(*RanUe)
	switch event {
	case fsm.EntryEvent:
	case N1MsgEvent:
		msg := args.(*libnas.GmmMessage)

		switch msg.GetMessageType() {
		case libnas.MsgTypeULNASTransport:
			ranue.handleUlNasTransport(msg.ULNASTransport)
		default:
			//warning
		}

	case T3522Event:
		ranue.Trace("Degistration expires")
		//TODO: re-send deregistration request
	default:
	}

}

func registered(state fsm.State, event fsm.EventType, args interface{}) {
	ranue := state.(*RanUe)

	switch event {
	case fsm.EntryEvent:
		//TODO: add logic ro re-send acceptance if there is no response
		ranue.registered = true
		if ranue.regctx.RegistrationRequest() != nil {
			ranue.sendAcceptance4Registration()
		} else {
			ranue.sendAcceptance4Registration()
		}
	case N1MsgEvent:
		msg := args.(*libnas.GmmMessage)

		switch msg.GetMessageType() {
		case libnas.MsgTypeULNASTransport:
			ranue.handleUlNasTransport(msg.ULNASTransport)
		default:
			//warning
		}
	case T3565Event:
		ranue.Trace("Notification expires")
		//TODO: re-send notification

	case T3550Event:
		ranue.Trace("RegistrationAccept expires")
		//TODO: resend - registration accept

	case N1N2TransferEvent:
		job, _ := args.(*common.AsyncJob)
		ranue.handleN1N2Transfer(job)

	case NotificationEvent:
		job, _ := args.(*common.AsyncJob)
		ranue.handleNotificationCommand(job)

	case PduModRspEvent:
		msg, _ := args.(*n2models.PduSessResModRsp)
		ranue.handlePduSessResModRsp(msg)

	case PduRelRspEvent:
		msg, _ := args.(*n2models.PduSessResRelRsp)
		ranue.handlePduSessResRelRsp(msg)

	case PduSetupRspEvent:
		msg, _ := args.(*n2models.PduSessResSetRsp)
		ranue.handlePduSessResSetRsp(msg)

	case PduNotifyEvent:
		msg, _ := args.(*n2models.PduSessResNot)
		ranue.handlePduSessResNot(msg)

	case PduModIndEvent:
		msg, _ := args.(*n2models.PduSessResModInd)
		ranue.handlePduSessResModInd(msg)
	case fsm.ExitEvent:
		ranue.regctx = nil
	default:
	}
}
