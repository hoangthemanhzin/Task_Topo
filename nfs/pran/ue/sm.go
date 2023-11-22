package ue

import (
	"etrib5gc/common"
	amfran "etrib5gc/sbi/amf/ran"
	"etrib5gc/sbi/models/n2models"
	"etrib5gc/util/fsm"
	"fmt"
)

const (
	CM_IDLE fsm.StateType = iota
	CM_SEARCHING
	CM_CONNECTED
)
const (
	SearchAmfEvent fsm.EventType = fsm.ExitEvent + iota + 1
	FoundAmfEvent                //authentication is triggerred
	NgapEvent
	SbiEvent
	FailEvent
	CloseEvent
	EndOfLifeEvent
)

var _sm *fsm.Fsm

func init() {
	transitions := fsm.Transitions{
		fsm.Tuple(CM_IDLE, NgapEvent):      CM_IDLE,
		fsm.Tuple(CM_SEARCHING, NgapEvent): CM_SEARCHING,
		fsm.Tuple(CM_CONNECTED, NgapEvent): CM_CONNECTED,

		fsm.Tuple(CM_IDLE, EndOfLifeEvent):      CM_IDLE,
		fsm.Tuple(CM_SEARCHING, EndOfLifeEvent): CM_SEARCHING,

		fsm.Tuple(CM_IDLE, CloseEvent):      CM_IDLE,
		fsm.Tuple(CM_SEARCHING, CloseEvent): CM_SEARCHING,
		fsm.Tuple(CM_CONNECTED, CloseEvent): CM_CONNECTED,

		fsm.Tuple(CM_SEARCHING, SbiEvent): CM_SEARCHING,
		fsm.Tuple(CM_CONNECTED, SbiEvent): CM_CONNECTED,

		fsm.Tuple(CM_IDLE, SearchAmfEvent):     CM_SEARCHING,
		fsm.Tuple(CM_IDLE, FoundAmfEvent):      CM_CONNECTED,
		fsm.Tuple(CM_SEARCHING, FoundAmfEvent): CM_CONNECTED,

		fsm.Tuple(CM_SEARCHING, FailEvent): CM_IDLE,
		fsm.Tuple(CM_CONNECTED, FailEvent): CM_IDLE,
	}

	callbacks := fsm.Callbacks{
		CM_IDLE:      cm_idle,
		CM_CONNECTED: cm_connected,
		CM_SEARCHING: cm_searching,
	}
	_sm = fsm.NewFsm(transitions, callbacks)
}

func cm_idle(state fsm.State, event fsm.EventType, args interface{}) {
	uectx, _ := state.(*UeContext)
	switch event {
	case fsm.EntryEvent:
		//enter this state due to a failure
		if err, ok := args.(error); ok {
			uectx.Tracef("Back to idle state with en error: %s", err.Error())
		}
		uectx.cu.RemoveUe(uectx)

	case EndOfLifeEvent:
		//it takes too long, kill me
		uectx.cu.RemoveUe(uectx)

	case CloseEvent:
		uectx.cu.RemoveUe(uectx)

	case NgapEvent:
		dat, _ := args.(*common.EventData) //must never fail
		if dat.EvType == NGAP_INIT_UE {
			msg, _ := dat.Content.(*n2models.InitUeContextRequest)
			msg.RanUeId = uectx.cuNgapId
			uectx.initmsg = msg
			if uectx.amfcli == nil { //no AMF
				if err := uectx.getDefaultAmf(); err == nil {
					uectx.Info("Default AMF found")
					uectx.sendEvent(SearchAmfEvent, msg)
				} else {
					uectx.Errorf("Fail to get a default AMF: %s", err.Error())
				}
			} else { //AMF found
				uectx.sendEvent(FoundAmfEvent, msg)
			}
		}
	default:
	}
}

func cm_searching(state fsm.State, event fsm.EventType, args interface{}) {
	uectx, _ := state.(*UeContext)
	switch event {
	case fsm.EntryEvent:
		//init uecontext with default AMF
		callback := uectx.cu.Callback()
		msg, _ := args.(*n2models.InitUeContextRequest)
		uectx.Info("Request InitUeContext to DAMF")
		if rsp, err := amfran.InitUeContext(uectx.amfcli, msg, callback); err != nil {
			uectx.Errorf("Request InitUeContext failed: %s", err.Error())
			uectx.sendEvent(FailEvent, err)
		} else {
			uectx.Infof("Receive AmfUeId=%d", rsp.AmfUeId)
			uectx.amfUeId = rsp.AmfUeId
		}

	case EndOfLifeEvent:
		//it takes too long, kill me
		uectx.cu.RemoveUe(uectx)

	case CloseEvent:
		uectx.cu.RemoveUe(uectx)

	case NgapEvent:
		dat, _ := args.(*common.EventData)
		switch dat.EvType {
		case NGAP_UL_NAS:
			msg, _ := dat.Content.(*n2models.UlNasTransport)
			uectx.Info("Forward UlNasTransport to  DAMF")
			if err := amfran.UlNasTransport(uectx.amfcli, uectx.AmfUeId(), msg); err != nil {
				uectx.Errorf("Forward ULNasTransport failed: %s", err.Error())
				uectx.sendEvent(FailEvent, err)
			}
		case NGAP_NAS_NON_DELIVERY:
			msg, _ := dat.Content.(*n2models.NasNonDeliveryIndication)
			uectx.Info("Forward NonDeliveryIndication to  DAMF")
			if err := amfran.NasNonDeliveryIndication(uectx.amfcli, uectx.AmfUeId(), msg); err != nil {
				uectx.Errorf("Forward NonDeliveryIndication failed: %s", err.Error())
				uectx.sendEvent(FailEvent, err)
			}

		default:
			uectx.Warnf("Receive unexpected Ngap Event: %d", dat.EvType)
		}
	case SbiEvent:
		dat, _ := args.(*common.EventData)
		job, _ := dat.Content.(*SbiJob)
		var err error
		switch dat.EvType {
		case SBI_INIT_UE_STATUS:
			jobinfo, _ := job.info.(*InitUeContextStatusJob)
			msg := jobinfo.Msg

			job.done(nil)

			if msg.Success {
				uectx.Info("First authentication step succeeded")

				//update UeContext information
				uectx.initmsg.UeCtx = msg.UeCtx
				//set AmfId
				uectx.amfid = msg.AmfId

				if err = uectx.getAmf(); err != nil {
					uectx.Errorf("Get AMF failed: %s", err.Error())
					uectx.sendEvent(FailEvent, err)
				} else {
					uectx.Info("AMF Found")
					uectx.sendEvent(FoundAmfEvent, uectx.initmsg)
				}
			} else {
				uectx.Errorf("First authentication step failed: %s", err.Error())
				if len(msg.NasPdu) > 0 {
					uectx.SendDownlinkNasTransport(&n2models.NasDlMsg{
						NasPdu: msg.NasPdu,
					})
				}
				err = fmt.Errorf(msg.Error)
				uectx.sendEvent(FailEvent, err)
			}
		case SBI_NAS_DL:
			jobinfo, _ := job.info.(*NasDlJob)
			msg := jobinfo.Msg
			if err = uectx.SendDownlinkNasTransport(msg); err != nil {
				//log.Errorf(err.Error())
				uectx.sendEvent(FailEvent, err)
			}
			job.done(err)
		default:
			err = fmt.Errorf("Unexpected Sbi event: %d", dat.EvType)
			job.done(err)
		}
	}
}

// TODO: when Ue in connected state (connecting to a
// designated AMF) is removed due to a communication failure, it may need to
// notify its partners (either the gnB or the AMF) of its failure.
func cm_connected(state fsm.State, event fsm.EventType, args interface{}) {
	uectx, _ := state.(*UeContext)
	switch event {
	case fsm.EntryEvent:
		//live forever
		uectx.alivetimer.Stop()

		//init uecontext with AMF
		callback := uectx.cu.Callback()
		uectx.Info("Request InitUeContext to AMF")
		if rsp, err := amfran.InitUeContext(uectx.amfcli, uectx.initmsg, callback); err != nil {
			uectx.Errorf("Request InitUeContext to AMF failed: %s", err.Error())
			uectx.sendEvent(FailEvent, err)
		} else {
			uectx.Infof("Receive AmfUeId=%d", rsp.AmfUeId)
			uectx.amfUeId = rsp.AmfUeId
		}

	case CloseEvent:
		uectx.cu.RemoveUe(uectx)

	case NgapEvent:
		dat, _ := args.(*common.EventData)
		switch dat.EvType {
		case NGAP_UL_NAS:
			msg, _ := dat.Content.(*n2models.UlNasTransport)
			uectx.Info("Forward UlNasTransport to  AMF")
			if err := amfran.UlNasTransport(uectx.amfcli, uectx.AmfUeId(), msg); err != nil {
				uectx.Infof("Forward UlNasTransport to  AMF failed: %s", err.Error())
				uectx.sendEvent(FailEvent, err)
			}

		case NGAP_NAS_NON_DELIVERY:
			msg, _ := dat.Content.(*n2models.NasNonDeliveryIndication)
			uectx.Info("Forward NasNonDelivery to  AMF")
			if err := amfran.NasNonDeliveryIndication(uectx.amfcli, uectx.AmfUeId(), msg); err != nil {
				uectx.Infof("Forward NasNonDelivery to  AMF failed: %s", err.Error())
				uectx.sendEvent(FailEvent, err)
			}
		case NGAP_UE_RRC_REP:
			msg, _ := dat.Content.(*n2models.RrcInactTranRep)
			uectx.Info("Forward RrcInactiveTransactionReport to  AMF")
			if err := amfran.RrcInactiveTransactionReport(uectx.amfcli, uectx.AmfUeId(), msg); err != nil {
				uectx.Infof("Forward RrcInactiveTransactionReport to  AMF failed: %s", err.Error())
				uectx.sendEvent(FailEvent, err)
			}

		case NGAP_UE_SET_RSP:
			msg, _ := dat.Content.(*n2models.InitCtxSetupRsp)
			if job := uectx.pendingjobs.get(SBI_UE_SET_REQ); job != nil {
				jobinfo, _ := job.info.(*InitCtxSetupReqJob)
				jobinfo.Rsp = msg
				job.done(nil)
			} else {
				uectx.Warnf("Orphan InitCtxSetRsp message")
			}
		case NGAP_UE_SET_FAIL:
			msg, _ := dat.Content.(*n2models.InitCtxSetupFailure)
			if job := uectx.pendingjobs.get(SBI_UE_SET_REQ); job != nil {
				jobinfo, _ := job.info.(*InitCtxSetupReqJob)
				jobinfo.Ersp = msg
				job.done(nil)
			} else {
				uectx.Warnf("Orphan InitCtxSetFailure message")
			}
		case NGAP_UE_MOD_RSP:
			msg, _ := dat.Content.(*n2models.UeCtxModRsp)
			if job := uectx.pendingjobs.get(SBI_UE_MOD_REQ); job != nil {
				jobinfo, _ := job.info.(*UeCtxModReqJob)
				jobinfo.Rsp = msg
				job.done(nil)
			} else {
				uectx.Warnf("Orphan UeCtxModRsp message")
			}
		case NGAP_UE_MOD_FAIL:
			msg, _ := dat.Content.(*n2models.UeCtxModFail)
			if job := uectx.pendingjobs.get(SBI_UE_MOD_REQ); job != nil {
				jobinfo, _ := job.info.(*UeCtxModReqJob)
				jobinfo.Ersp = msg
				job.done(nil)
			} else {
				uectx.Warnf("Orphan InitCtxSetFailure message")
			}
		case NGAP_PDU_SET_RSP:
			msg, _ := dat.Content.(*n2models.PduSessResSetRsp)
			if job := uectx.pendingjobs.get(SBI_PDU_SET_REQ); job != nil {
				jobinfo, _ := job.info.(*PduSessResSetReqJob)
				jobinfo.Rsp = msg
				job.done(nil)
			} else {
				uectx.Warnf("Orphan PduSessResSetRsp message")
			}
		case NGAP_PDU_MOD_RSP:
			msg, _ := dat.Content.(*n2models.PduSessResModRsp)
			if job := uectx.pendingjobs.get(SBI_PDU_MOD_REQ); job != nil {
				jobinfo, _ := job.info.(*PduSessResModReqJob)
				jobinfo.Rsp = msg
				job.done(nil)
			} else {
				uectx.Warnf("Orphan PduSessResModRsp message")
			}
		case NGAP_PDU_REL_RSP:
			msg, _ := dat.Content.(*n2models.PduSessResRelRsp)
			if job := uectx.pendingjobs.get(SBI_PDU_REL_CMD); job != nil {
				jobinfo, _ := job.info.(*PduSessResRelCmdJob)
				jobinfo.Rsp = msg
				job.done(nil)
			} else {
				uectx.Warnf("Orphan PduSessResRelRsp message")
			}
		case NGAP_PDU_NOT:
			msg, _ := dat.Content.(*n2models.PduSessResNot)
			uectx.Info("Forward PduSessionResourceNotification to AMF")
			if err := amfran.PduSessionResourceNotification(uectx.amfcli, uectx.AmfUeId(), msg); err != nil {
				uectx.Infof("Forward PduSessionResourceNotification to AMF failed: %s", err.Error())
				uectx.sendEvent(FailEvent, err)
			}
		case NGAP_PDU_MOD_IND:
			msg, _ := dat.Content.(*n2models.PduSessResModInd)
			uectx.Info("Forward PduSessionResourceModifyIndication to AMF")
			if err := amfran.PduSessionResourceModifyIndication(uectx.amfcli, uectx.AmfUeId(), msg); err != nil {
				uectx.Info("Forward PduSessionResourceModifyIndication to AMF failed: %s", err.Error())
				uectx.sendEvent(FailEvent, err)
			}
		default:
			uectx.Warnf("Unexpected Ngap event: %d", dat.EvType)
		}
	case SbiEvent:
		dat, _ := args.(*common.EventData)
		job, _ := dat.Content.(*SbiJob)
		var err error
		switch dat.EvType {
		case SBI_NAS_DL:
			jobinfo, _ := job.info.(*NasDlJob)
			msg := jobinfo.Msg
			if err = uectx.SendDownlinkNasTransport(msg); err != nil {
				//log.Errorf(err.Error())
				uectx.sendEvent(FailEvent, err)
			}
			job.done(err)

		case SBI_UE_SET_REQ:
			jobinfo, _ := job.info.(*InitCtxSetupReqJob)
			msg := jobinfo.Msg
			if err := uectx.SendInitialContextSetupRequest(msg); err != nil {
				uectx.sendEvent(FailEvent, err)
				job.done(err)
			} else {
				uectx.addJob(dat.EvType, job)
			}

		case SBI_UE_MOD_REQ:
			jobinfo, _ := job.info.(*UeCtxModReqJob)
			msg := jobinfo.Msg
			if err := uectx.SendUEContextModificationRequest(msg); err != nil {
				uectx.sendEvent(FailEvent, err)
				job.done(err)
			} else {
				uectx.addJob(dat.EvType, job)
			}
		case SBI_PDU_SET_REQ:
			jobinfo, _ := job.info.(*PduSessResSetReqJob)
			msg := jobinfo.Msg
			if err := uectx.SendPduSessionResourceSetupRequest(msg); err != nil {
				uectx.sendEvent(FailEvent, err)
				job.done(err)
			} else {
				uectx.addJob(dat.EvType, job)
			}

		case SBI_PDU_MOD_REQ:
			jobinfo, _ := job.info.(*PduSessResModReqJob)
			msg := jobinfo.Msg
			if err := uectx.SendPduSessionResourceModifyRequest(msg); err != nil {
				uectx.sendEvent(FailEvent, err)
				job.done(err)
			} else {
				uectx.addJob(dat.EvType, job)
			}

		case SBI_PDU_REL_CMD:
			jobinfo, _ := job.info.(*PduSessResRelCmdJob)
			msg := jobinfo.Msg
			if err := uectx.SendPduSessionResourceReleaseCommand(msg); err != nil {
				uectx.sendEvent(FailEvent, err)
				job.done(err)
			} else {
				uectx.addJob(dat.EvType, job)
			}

		default:
			err = fmt.Errorf("Unexpected Sbi event: %d", dat.EvType)
			job.done(err)
		}

	}
}
