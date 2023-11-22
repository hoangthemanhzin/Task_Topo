package ue

import (
	"etrib5gc/common"
	"etrib5gc/nas"
	"etrib5gc/sbi/models/n2models"
	"etrib5gc/util/fsm"

	libnas "github.com/free5gc/nas"
)

const (
	UE_IDLE fsm.StateType = iota
	UE_AUTHENTICATING
)
const (
	AuthEvent fsm.EventType = fsm.ExitEvent + iota + 1
	SbiNasEvent
	T3560Event
	CloseEvent
)

var _sm *fsm.Fsm

func init() {
	transitions := fsm.Transitions{
		fsm.Tuple(UE_IDLE, CloseEvent):           UE_IDLE,
		fsm.Tuple(UE_AUTHENTICATING, CloseEvent): UE_AUTHENTICATING,

		fsm.Tuple(UE_IDLE, SbiNasEvent):           UE_IDLE,
		fsm.Tuple(UE_AUTHENTICATING, SbiNasEvent): UE_AUTHENTICATING,

		fsm.Tuple(UE_AUTHENTICATING, T3560Event): UE_AUTHENTICATING,

		fsm.Tuple(UE_IDLE, AuthEvent): UE_AUTHENTICATING,
	}

	callbacks := fsm.Callbacks{
		UE_IDLE:           ue_idle,
		UE_AUTHENTICATING: ue_authenticating,
	}
	_sm = fsm.NewFsm(transitions, callbacks)
}

func ue_idle(state fsm.State, event fsm.EventType, args interface{}) {
	uectx, _ := state.(*UeContext)
	switch event {
	case fsm.EntryEvent:
	case CloseEvent:
		uectx.ctx.RemoveUe(uectx)
	case SbiNasEvent:
		dat, _ := args.(*common.EventData) //must never fail
		if dat.EvType == NAS_INIT_UE {
			uectx.sendEvent(AuthEvent, nil)
		}
	default:
	}
}

func ue_authenticating(state fsm.State, event fsm.EventType, args interface{}) {
	uectx, _ := state.(*UeContext)
	switch event {
	case fsm.EntryEvent:
		if err := uectx.authenticate(); err != nil {
			uectx.Errorf("Authentication failed: %s", err.Error())
			uectx.report(err)
		}

	case CloseEvent:
		uectx.ctx.RemoveUe(uectx)

	case SbiNasEvent:
		dat, _ := args.(*common.EventData) //must never fail
		switch dat.EvType {
		case NAS_UL_TRANSPORT:
			msg, _ := dat.Content.(*n2models.UlNasTransport)
			var err error
			var nasMsg libnas.Message
			if nasMsg, err = nas.Decode(nil, msg.NasPdu); err == nil && nasMsg.AuthenticationResponse != nil {
				if err = uectx.handleAuthenticationResponse(nasMsg.AuthenticationResponse); err != nil {
					uectx.Errorf("Authentication failed: %s", err.Error())
				} else {
					uectx.Infof("Authentication succeeded")
					if err = uectx.findAmf(); err != nil {
						uectx.Errorf("Find Amf failed: %s", err.Error())
					}
				}
			} else {
				uectx.Errorf("Invalid UlNasTransport: %s", err.Error())
			}
			//report to PRAN
			uectx.report(err)
		case NAS_NON_DELIVERY:
			//TODO: handle nas non delivery indication
			//Note: UeContext will be removed anyway (due to timeout) if authentication and AMF
			//finding procedures do not finish.
			uectx.Warnf("Receive a NasNonDeliveryIndication from Ran")
		default:
			uectx.Warnf("Unexpected Nas Event %d", dat.EvType)
		}
	}
}
